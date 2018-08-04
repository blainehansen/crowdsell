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

for (const [tableName, fields] of Object.entries(rawTables)) {
	const processed = Object.entries(fields).map(processFields)

	tables.push({
		name: tableName,
		fields: universalFields.concat(processed),
	})
}

function processFields([name, rawField]) {
	const field = { name }

	if (typeof rawField == 'string') {
		field.type = rawField
		rawField = {}
	}
	else field.type = rawField.type

	if (rawField.functions) globalFunctions.push(rawField.functions)
	field.triggers = rawField.triggers

	field.references = rawField.references

	field.unique = !!rawField.unique
	field.required = !!rawField.required

	const server_private = !!rawField.server_private
	const private = !!rawField.private
	const no_patch = !!rawField.no_patch

	field.owner_read = !server_private
	field.public_read = !server_private && !private
	field.owner_patch = !server_private && !no_patch

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

// method array: [outer name, args type, goqu name, return type, necessary field switches]
const allMethods = [
	['As', 'string', undefined, 'AliasedExpression'],
	['Asc', NO_ARGS, undefined, 'OrderedExpression'],
	['Desc', NO_ARGS, undefined, 'OrderedExpression'],
	['Distinct', NO_ARGS, undefined, 'SqlFunctionExpression'],
]
const postgresGoTypeMap = {
	'primary': {
		goType: 'int64',
		readOnly: true,
		methods: [
			['Get', undefined, 'Eq'],
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
			['IsNull', NO_ARGS, undefined, undefined, { required: false }],
	    ['IsNotNull', NO_ARGS, undefined, undefined, { required: false }],
		]
	},
	'boolean': {
		goType: 'bool',
		methods: [
			['Is'],
			['True', NO_ARGS, 'IsTrue'],
			['False', NO_ARGS, 'IsFalse'],
			['IsNull', NO_ARGS, undefined, undefined, { required: false }],
	    ['IsNotNull', NO_ARGS, undefined, undefined, { required: false }],
		]
	},
	'bytea': {
		goType: '[]byte',
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
			['IsNull', NO_ARGS, undefined, undefined, { required: false }],
	    ['IsNotNull', NO_ARGS, undefined, undefined, { required: false }],
		]
	},
	'timestamptz': {
		goType: 'time.Time',
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
			['IsNull', NO_ARGS, undefined, undefined, { required: false }],
	    ['IsNotNull', NO_ARGS, undefined, undefined, { required: false }],
		]
	},
}



const go = {
	decideArgsString(argType, goType) {
		if (argType === NO_ARGS) return ['', '']
		if (argType === ARRAY_ARGS) return [`val []${goType}`, 'val']
		if (argType === RANGE_ARGS) return [`startVal ${goType}, endVal ${goType}`, 'goqu.RangeVal{ Start: startVal, End: endVal }']
		return [`val ${argType}`, 'val']
	},

	makeGoquTypeName: (postgresType) => postgresType + 'Column',

	makeGoquTypes() {
		const goquTypes = []
		for (const [postgresType, { goType, methods: typeMethods, readOnly = false }] of Object.entries(postgresGoTypeMap)) {
			const methods = typeMethods.concat(allMethods)

			const goquTypeName = this.makeGoquTypeName(postgresType)

			goquTypes.push(`type ${goquTypeName} struct {\n\tc string\n\ti goqu.IdentifierExpression\n}`)
			goquTypes.push(`func (c ${goquTypeName}) Identifier() goqu.IdentifierExpression {\n\treturn c.i\n}`)

			if (!readOnly) {
				goquTypes.push(`func (c *${goquTypeName}) Set(val ${goType}) SetExpression {\n\treturn SetExpression{ Name: c.c, Value: val }\n}`)
			}

			for (const method of methods) {
				const [
					outerName,
					argType = goType,
					innerName = outerName,
					returnType = 'BooleanExpression',
				] = method

				const [argsString, valString] = this.decideArgsString(argType, goType)
				goquTypes.push(`func (c *${goquTypeName}) ${outerName}(${argsString}) goqu.${returnType} {\n\treturn c.i.${innerName}(${valString})\n}`)
			}

			goquTypes.push("")
		}

		return goquTypes.join('\n')
	},

	stringifyTable(table) {
		const returnStrings = []

		// first the generic table object
		const tableName = table.name
		const pascalTableName = changeCase.pascal(tableName)
		returnStrings.push(`var ${pascalTableName}Table = CreateTable("${tableName}")`)

		// then its schema
		const schemaName = tableName + 'Schema'

		const schemaStructFields = []
		const schemaInstanceFields = []
		for (const field of table.fields) {
			const pascalFieldName = changeCase.pascal(field.name)
			const goquTypeName = this.makeGoquTypeName(field.type)

			schemaStructFields.push(pascalFieldName + ' ' + goquTypeName)
			const fieldNameString = `"${tableName}.${field.name}"`
			schemaInstanceFields.push(`${pascalFieldName}: ${goquTypeName}{ c: ${fieldNameString}, i: goqu.I(${fieldNameString}) },`)
		}

		returnStrings.push(`type ${schemaName} struct {\n\t${schemaStructFields.join('\n\t')}\n}`)
		returnStrings.push(`var ${pascalTableName} = &${schemaName}{\n\t${schemaInstanceFields.join('\n\t')}\n}`)


		const structTypes = []

		structTypes.push(['server', table.fields.slice()])
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
			`\t"github.com/blainehansen/goqu"`,
			")",
			"",
			this.makeGoquTypes(),
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
			"source ./.env",
			"PGPASSWORD=$DATABASE_PASSWORD psql -U $DATABASE_USER -h $SYSTEM_DATABASE_HOST $DATABASE_DB_NAME << EOF",
			this.stringify(tables),
			"EOF"
		].join('\n')

		fs.writeFileSync('./migrate.sh', migrateFileString)
	}
}

go.create(tables)
postgres.create(tables)
