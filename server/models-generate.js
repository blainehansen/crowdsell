const changeCase = require('change-case')
const pluralize = require('pluralize')
const yaml = require('js-yaml')
const fs = require('fs')

const modelsManifest = yaml.load(fs.readFileSync('./models.yml'))

const tables = []

const universalFields = Object.entries(modelsManifest.universal).map(processFields)
delete modelsManifest.universal
// console.log(universalFields)

for (const [tableName, fields] of Object.entries(modelsManifest)) {
	// console.log(tableName)
	// console.log(fields)

	const processed = Object.entries(fields).map(processFields)
	// console.log(processed)

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

	field.references = rawField.references

	field.unique = !!rawField.unique
	field.required = !!rawField.required

	const server_private = !!rawField.server_private
	const private = !!rawField.private
	const read_only = !!rawField.read_only

	field.owner_read = !server_private
	field.public_read = !server_private && !private
	field.owner_write = !server_private && !read_only

	return field
}

function genericStringifyField(field) {
	const fieldFunction = this.fieldFunctions[field.type] || this.fieldFunctions._default
	return fieldFunction.call(this, field)
}

function genericStringify(tables) {
	return tables.map(this.stringifyTable.bind(this)).join('\n\n')
}

const structs = {
	typeMap: {
		'primary': 'int64',
		'text': 'string',
		'bytea': '[]byte',
		'bigint': 'int64',
		'timestamptz': 'time.Time',
	},

	fieldFunctions: {
		_default(field) {
			return changeCase.pascal(field.name) + ' ' + this.typeMap[field.type]
		}
	},

	stringifyField: genericStringifyField,

	stringifyTable(table) {
		const structTypes = []

		structTypes.push(['', table.fields.slice()])

		for (const permissionLevel of ['owner_write', 'owner_read', 'public_read']) {
			structTypes.push([permissionLevel, table.fields.filter(field => !!field[permissionLevel])])
		}
		const structStrings = structTypes.map(([permissionLevel, fields]) => {
			const structName = changeCase.pascal(permissionLevel) + changeCase.pascal(pluralize.singular(table.name))
			const fieldsString = fields.map(this.stringifyField.bind(this)).join('\n\t')
			return `type ${structName} struct {\n\t${fieldsString}\n}`
		})

		return structStrings.join('\n\n')
	},

	stringify: genericStringify,
}




// type AliasMethods interface {
// 	As(interface{}) AliasedExpression
// }

// type ComparisonMethods interface {
// 	Eq(interface{}) BooleanExpression
// 	Neq(interface{}) BooleanExpression
// 	Gt(interface{}) BooleanExpression
// 	Gte(interface{}) BooleanExpression
// 	Lt(interface{}) BooleanExpression
// 	Lte(interface{}) BooleanExpression
// }



const goqu = {
	fieldFunctions: {
		_default(field) {
			const fieldString = `goqu.I("${field.name}")`
		}
	},

	stringifyTable(table) {
		// for each field, it would be cool to create

	},

	stringify: genericStringify,
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
		const fieldsString = table.fields.map(this.stringifyField.bind(this)).join(',\n\t')
		return `CREATE TABLE ${table.name} (\n\t${fieldsString}\n)`
	},

	stringify(tables) {
		const tablesString = tables.map(this.stringifyTable.bind(this)).join('\n')
		return `BEGIN;\n\n${tablesString}\n\nEND;`
	},
}

console.log(structs.stringify(tables))
// console.log(postgres.stringify(tables))
