const changeCase = require('change-case')
const pluralize = require('pluralize')
const yaml = require('js-yaml')
const fs = require('fs')

// String.format(cssBlockStringTemplate, selector, attrs.join('\n\t'))
String.format = function(format) {
	let args = Array.prototype.slice.call(arguments, 1)
	return format.replace(/{(\d+)}/g, (match, number) => args[number] != 'undefined' ? args[number] : match)
}



const modelsManifest = yaml.load(fs.readFileSync('./models.yml'))

const {
	functions = "",
	universal: rawUniversal = {},
	...rawTables
} = modelsManifest

const globalFunctions = [functions]
const universalFields = Object.entries(rawUniversal).map(processFields)
const tables = []


// 0: tableName
// 1: list of column names
// 2: name of vector
const searchTriggerBoilerplate = `CREATE TRIGGER search_update_{0}_{2}
BEFORE INSERT OR UPDATE OF {1} ON {0}
FOR EACH ROW
EXECUTE PROCEDURE tsvector_update_trigger({2}, 'pg_catalog.english', {1});
`

// `CREATE FUNCTION film_weighted_tsv_trigger() RETURNS trigger AS $$
// begin
//   new.weighted_tsv :=
//      setweight(to_tsvector('english', COALESCE(new.title,'')), 'A') ||
//      setweight(to_tsvector('english', COALESCE(new.description,'')), 'B');
//   return new;
// end
// $$ LANGUAGE plpgsql;`


for (const [tableName, fields] of Object.entries(rawTables)) {
	const processed = Object.entries(fields).map(processFields)

	tables.push({
		name: tableName,
		fields: universalFields.concat(processed),
	})
}


function processFields([fieldName, rawField]) {
	const field = { name: fieldName }

	if (typeof rawField == 'string') {
		field.type = rawField
		rawField = {}
	}
	else field.type = rawField.type

	if (rawField.type == 'tsvector') {
		rawField.triggers = String.format(
			searchTriggerBoilerplate,
			// pull a trick here to preserve the tableName one
			'{0}',
			rawField.includes.join(', '),
			fieldName,
		)

		rawField.triggers += `\nCREATE INDEX {0}_${fieldName}_idx ON {0} USING gin (${fieldName});`

		rawField.server_private = true
	}

	if (rawField.functions) globalFunctions.push(rawField.functions)
	field.triggers = rawField.triggers

	field.references = rawField.references

	field.unique = !!rawField.unique
	field.required = !!rawField.required

	const read_only = !!rawField.read_only
	field.read_only = read_only

	const server_private = !!rawField.server_private
	const private = !!rawField.private
	const no_patch = !!rawField.no_patch

	field.owner_read = !server_private
	field.public_read = !server_private && !private
	field.owner_patch = !server_private && !no_patch && !read_only

	return field
}

function genericStringifyField(field) {
	const fieldFunction = this.fieldFunctions[field.type] || this.fieldFunctions._default
	return fieldFunction.call(this, field)
}

function genericStringify(tables) {
	return tables.map(this.stringifyTable.bind(this)).join('\n\n')
}

const NO_ARGS = Symbol()
const ARRAY_ARGS = Symbol()
const RANGE_ARGS = Symbol()

// method array: [outer name, args type, goqu name, return type]
const allGoTypeMethods = [
	['As', 'string', undefined, 'AliasedExpression'],
	['Asc', NO_ARGS, undefined, 'OrderedExpression'],
	['Desc', NO_ARGS, undefined, 'OrderedExpression'],
	['Distinct', NO_ARGS, undefined, 'SqlFunctionExpression'],
]
const postgresGoTypeMap = {
	'primary': {
		goType: 'int64',
		methods: [
			['Eq'],
			['Neq'],
			['In', ARRAY_ARGS],
			['NotIn', ARRAY_ARGS],
		]
	},
	'text': {
		goType: 'string',
		methods: [
			['Eq'],
			['Neq'],
			['Gt'],
			['Gte'],
			['Lt'],
			['Lte'],
			['In', ARRAY_ARGS],
			['NotIn', ARRAY_ARGS],
			['Like'],
			['NotLike'],
			['ILike'],
			['NotILike'],
		]
	},
	'boolean': {
		goType: 'bool',
		methods: [
			['Is'],
			['True', NO_ARGS, 'IsTrue'],
			['False', NO_ARGS, 'IsFalse'],
		]
	},
	'bytea': {
		goType: '[]byte',
		reflect: {
			outer: 'Slice',
			inner: 'Uint8',
		},
		methods: [
			['Eq'],
			['Neq'],
		]
	},
	'bigint': {
		goType: 'int64',
		methods: [
			['Eq'],
			['Neq'],
			['Gt'],
			['Gte'],
			['Lt'],
			['Lte'],
			['Between', RANGE_ARGS, undefined, 'RangeExpression'],
			['NotBetween', RANGE_ARGS, undefined, 'RangeExpression'],
			['In', ARRAY_ARGS],
			['NotIn', ARRAY_ARGS],
		]
	},
	'timestamptz': {
		goType: 'time.Time',
		reflect: {
			outer: 'Struct',
		},
		methods: [
			['Eq'],
			['Neq'],
			['Gt'],
			['Gte'],
			['Lt'],
			['Lte'],
			['Between', RANGE_ARGS, undefined, 'RangeExpression'],
			['NotBetween', RANGE_ARGS, undefined, 'RangeExpression'],
			['In', ARRAY_ARGS],
			['NotIn', ARRAY_ARGS],
		]
	},
}



const boilerPlateTemplate = `
func (d *{0}Dataset) Where(expressions ...goqu.Expression) *{0}Dataset {
	return &{0}Dataset{ d.Dataset.Where(expressions...) }
}

func (d *{0}Dataset) Select(columns ...DbColumn) *{0}Dataset {
	return &{0}Dataset{ d.Dataset.Select(makeColumns(columns)...) }
}

func (d *{0}Dataset) Returning(columns ...DbColumn) *{0}Dataset {
	return &{0}Dataset{ d.Dataset.Returning(makeColumns(columns)...) }
}

func (d *{0}Dataset) Update(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Update(makeRecord(expressions))
}

func (d *{0}Dataset) Insert(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Insert(makeRecord(expressions))
}

func (d *{0}Dataset) Patch(values map[string]interface{}) *patchExec {
	var realValues = make(map[string]interface{})
	for key, value := range values {
		realValues[strcase.ToSnake(key)] = value
	}

	p := patchExec{
		d.Dataset.Update(realValues),
		validatePatch(&realValues, &{1}),
		realValues,
	}

	return &p
}`

const go = {
	decideArgsString(argType, goType) {
		if (argType === NO_ARGS) return ['', '']
		if (argType === ARRAY_ARGS) return [`val []${goType}`, 'val']
		if (argType === RANGE_ARGS) return [`startVal ${goType}, endVal ${goType}`, 'goqu.RangeVal{ Start: startVal, End: endVal }']
		return [`val ${argType}`, 'val']
	},

	makeGoquTypeName: (tableName, fieldName) => tableName + changeCase.pascal(fieldName) + 'Column',

	makeGoquTypeForField(tableName, field) {
		const { required: fieldRequired, name: fieldName, type: fieldPostgresType, read_only: fieldReadOnly } = field

		if (fieldPostgresType == 'tsvector') {
			const goquTypeEntries = []
			const goquTypeName = this.makeGoquTypeName(tableName, fieldName)

			goquTypeEntries.push(`type ${goquTypeName} struct {\n\tcolumn\n}`)
			const searchLiteral = `goqu.L(\`${tableName}.${fieldName} @@ to_tsquery('pg_catalog.english', ?)\`, val)`
			goquTypeEntries.push(`func (c *${goquTypeName}) Search(val string) goqu.LiteralExpression {\n\treturn ${searchLiteral}\n}`)

			// TODO add a "rank" function that can be used to create an order by

			return goquTypeEntries.join('\n')
		}

		const { goType: fieldGoType, methods: typeMethods } = postgresGoTypeMap[fieldPostgresType]

		const goquTypeEntries = []
		const goquTypeName = this.makeGoquTypeName(tableName, fieldName)
		goquTypeEntries.push(`type ${goquTypeName} struct {\n\tcolumn\n}`)


		if (!fieldReadOnly) {
			goquTypeEntries.push(`func (c *${goquTypeName}) Set(val ${fieldGoType}) SetExpression {\n\treturn SetExpression{ Name: "${fieldName}", Value: val }\n}`)

			if (!fieldRequired) {
				goquTypeEntries.push(`func (c *${goquTypeName}) Clear() SetExpression {\n\treturn SetExpression{ Name: "${fieldName}", Value: nil }\n}`)
			}
		}
		// if (!fieldRequired && !fieldDefault) {
		if (!fieldRequired) {
			// ['IsNull', NO_ARGS, undefined, undefined],
			goquTypeEntries.push(`func (c *${goquTypeName}) IsNull() goqu.BooleanExpression {\n\treturn c.column.i.IsNull()\n}`)

			// ['IsNotNull', NO_ARGS, undefined, undefined],
			goquTypeEntries.push(`func (c *${goquTypeName}) IsNotNull() goqu.BooleanExpression {\n\treturn c.column.i.IsNotNull()\n}`)
		}

		for (const method of allGoTypeMethods.concat(typeMethods)) {
			const [
				outerName,
				argType = fieldGoType,
				innerName = outerName,
				returnType = 'BooleanExpression',
			] = method

			const [argsString, valString] = this.decideArgsString(argType, fieldGoType)
			goquTypeEntries.push(`func (c *${goquTypeName}) ${outerName}(${argsString}) goqu.${returnType} {\n\treturn c.column.i.${innerName}(${valString})\n}`)
		}

		return goquTypeEntries.join('\n')
	},


	stringifyTable(table) {
		const returnStrings = []
		const tableName = table.name
		const pascalTableName = changeCase.pascal(tableName)


		// then its schema
		// modelSchema
		const schemaName = tableName + 'Schema'

		const fieldTypes = []
		const schemaStructFields = []
		const schemaInstanceFields = []
		const modelKinds = []
		for (const field of table.fields) {
			const fieldName = field.name
			const pascalFieldName = changeCase.pascal(fieldName)
			const goquTypeName = this.makeGoquTypeName(tableName, fieldName)

			fieldTypes.push(this.makeGoquTypeForField(tableName, field))
			fieldTypes.push("")

			schemaStructFields.push(pascalFieldName + ' ' + goquTypeName)

			const fieldNameString = `"${tableName}.${fieldName}"`
			schemaInstanceFields.push(`${pascalFieldName}: ${goquTypeName}{ column { i: goqu.I(${fieldNameString}) } },`)

			const currentFieldType = postgresGoTypeMap[field.type]
			if (field.owner_patch) {
				const {
					outer: reflectOuter = changeCase.upperCaseFirst(currentFieldType.goType),
					inner: reflectInner = 'Invalid',
				} = currentFieldType.reflect || {}

				if (reflectOuter == 'Struct')
					modelKinds.push(`"${fieldName}": NestedKind { Outer: reflect.Struct, Instance: ${currentFieldType.goType}{} },`)
				else
					modelKinds.push(`"${fieldName}": NestedKind { Outer: reflect.${reflectOuter}, Inner: reflect.${reflectInner} },`)
			}
		}
		returnStrings.push(fieldTypes.join('\n'))

		const datasetName = `${tableName}Dataset`
		returnStrings.push(`type ${datasetName} struct {\n\t*goqu.Dataset\n}`)

		returnStrings.push(`type ${schemaName} struct {\n\tTable *goqu.Dataset\n\tQuery *${datasetName}\n\t${schemaStructFields.join('\n\t')}\n}`)

		const tableCreateString = `db.From("${tableName}")`
		// Model (the schema instance) will have (and act as) a model specific safe dataset
		returnStrings.push(`var ${pascalTableName} = &${schemaName}{\n\tTable: ${tableCreateString},\n\tQuery: &${datasetName}{ ${tableCreateString} },\n\t${schemaInstanceFields.join('\n\t')}\n}`)

		const modelKindsName = tableName + 'Kinds'
		returnStrings.push(`var ${modelKindsName} = map[string]NestedKind {\n\t${modelKinds.join('\n\t')}\n}`)

		// the specific boiler plate methods that were previously in models-enhancements
		returnStrings.push(String.format(boilerPlateTemplate, tableName, modelKindsName))



		const structTypes = []

		structTypes.push(['server', table.fields.filter(field => field.type != 'tsvector')])
		for (const permissionLevel of ['owner_patch', 'owner_read', 'public_read']) {
			structTypes.push([permissionLevel, table.fields.filter(field => !!field[permissionLevel])])
		}

		for (const [permissionLevel, fields] of structTypes) {
			const structName = changeCase.pascal(permissionLevel) + changeCase.pascal(pluralize.singular(tableName))

			const fieldsString = fields.map((field) => {
				return changeCase.pascal(field.name) + ' ' + postgresGoTypeMap[field.type].goType
			}).join('\n\t')

			returnStrings.push(`type ${structName} struct {\n\t${fieldsString}\n}`)
		}

		return returnStrings.join('\n\n')
	},

	stringify: genericStringify,

	create(tables) {
		const structsFileString = [
			"package main",
			"import (",
			`\t"time"`,
			`\t"reflect"`,
			`\t"github.com/blainehansen/goqu"`,
			`\t"github.com/iancoleman/strcase"`,
			")",
			"",
			this.stringify(tables),
		].join('\n')

		fs.writeFileSync('./models.go', structsFileString)
	}
}


const postgres = {
	fieldFunctions: {
		primary: (field) => `${field.name} serial NOT NULL PRIMARY KEY`,
		_default(field) {
			let fieldString = `${field.name} ${field.type}`
			if (field.required) fieldString += ' NOT NULL'
			if (field.unique) fieldString += ' UNIQUE'
			if (field.references) fieldString += ` REFERENCES ${field.references}`
			return fieldString
		},
	},

	stringifyField: genericStringifyField,

	stringifyTable(table) {
		const tableName = table.name

		const fieldStrings = []
		const triggerStrings = []
		for (const field of table.fields) {
			fieldStrings.push(this.stringifyField(field))

			const triggers = field.triggers
			if (triggers) {
				triggerStrings.push(String.format(triggers, tableName))
			}
		}

		const stringifiedFields = fieldStrings.join(',\n\t')
		const stringifiedTriggers = triggerStrings.join('\n')
		return `CREATE TABLE ${tableName} (\n\t${stringifiedFields}\n);\n\n${stringifiedTriggers}`
	},

	stringify(tables) {
		const tablesString = tables.map(this.stringifyTable.bind(this)).join('\n\n')
		return `${globalFunctions.join('\n')}\n\nBEGIN;\n\n${tablesString}\n\nCOMMIT;\n\n`
	},

	create(tables) {
		const migrateFileString = [
			"source ../.env.sh",
			"",
			"PGPASSWORD=$DATABASE_PASSWORD psql -U $DATABASE_USER -h $SYSTEM_DATABASE_HOST $DATABASE_DB_NAME << EOF",
			"",
			this.stringify(tables),
			"EOF"
		].join('\n')

		fs.writeFileSync('./migrate.sh', migrateFileString)
	}
}

go.create(tables)
postgres.create(tables)
