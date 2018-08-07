package main

import (
	"reflect"
	"github.com/blainehansen/goqu"
)

type column struct {
	i goqu.IdentifierExpression
}

func (c column) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c *column) As(val string) goqu.AliasedExpression {
	return c.i.As(val)
}
func (c *column) Asc() goqu.OrderedExpression {
	return c.i.Asc()
}
func (c *column) Desc() goqu.OrderedExpression {
	return c.i.Desc()
}
func (c *column) Distinct() goqu.SqlFunctionExpression {
	return c.i.Distinct()
}




type DbColumn interface {
	Identifier() goqu.IdentifierExpression
}

type SetExpression struct {
	Name string
	Value interface{}
}

func makeRecord(expressions []SetExpression) goqu.Record {
	giveRecord := goqu.Record{}

	for _, expression := range expressions {
		giveRecord[expression.Name] = expression.Value
	}

	return giveRecord
}

func makeColumns(columns []DbColumn) []interface{} {
	giveColumns := make([]interface{}, len(columns))

	for i, column := range columns {
		giveColumns[i] = column.Identifier()
	}

	return giveColumns
}



type SafeDataset struct {
	*goqu.Dataset
}

func CreateTable(tableName string) *SafeDataset {
	return &SafeDataset{ db.From(tableName) }
}

func (d *SafeDataset) Where(expressions ...goqu.Expression) *SafeDataset {
	return &SafeDataset{ d.Dataset.Where(expressions...) }
}

func (d *SafeDataset) Select(columns ...DbColumn) *SafeDataset {
	return &SafeDataset{ d.Dataset.Select(makeColumns(columns)...) }
}

func (d *SafeDataset) Returning(columns ...DbColumn) *SafeDataset {
	return &SafeDataset{ d.Dataset.Returning(makeColumns(columns)...) }
}

func (d *SafeDataset) Update(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Update(makeRecord(expressions))
}

func (d *SafeDataset) Insert(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Insert(makeRecord(expressions))
}

// func (d *SafeDataset) InsertStruct(insert interface{}) *goqu.CrudExec {
// 	if !structs.IsStruct(insert) {
// 		panic("InsertStruct was not passed a struct: %T", insert)
// 	}

// 	values := make(map[string]interface{})
// 	for _, field := range structs.Fields(server) {
// 		values[field.Name()] = field.Value()
// 	}

// 	return d.Dataset.Insert(values)
// }


func (d *usersSchema) Patch(values map[string]interface{}) (*goqu.CrudExec, bool) {
	if !validatePatch(values, usersSchemaFields) {
		return nil, false
	}

	return d.SafeDataset.Dataset.Update(values), true
}

update := Users.Patch()




type NestedKind struct {
	Outer reflect.Kind
	Inner reflect.Kind
}

func typesMatch(valueKind reflect.Kind, schemaKind reflect.Kind) bool {
	switch schemaKind {
		case reflect.Int64:
			switch valueKind {
				case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64: return true
				default: return false
			}
		case reflect.Int32:
			switch valueKind {
				case reflect.Int8, reflect.Int16, reflect.Int32: return true
				default: return false
			}
		case reflect.Int16:
			switch valueKind {
				case reflect.Int8, reflect.Int16: return true
				default: return false
			}
		case reflect.Uint64:
			switch valueKind {
				case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64: return true
				default: return false
			}
		case reflect.Uint32:
			switch valueKind {
				case reflect.Uint8, reflect.Uint16, reflect.Uint32: return true
				default: return false
			}
		case reflect.Uint16:
			switch valueKind {
				case reflect.Uint8, reflect.Uint16: return true
				default: return false
			}
		default:
			return valueKind == schemaKind
	}
}

func validatePatch(values *map[string]interface{}, schema *map[string]NestedKind) bool {
	schemaAccess := *schema
	for key, value := range *values {
		schemaKind, present := schemaAccess[key]
		if !present {
			return false
		}

		baseValueType := reflect.TypeOf(value)
		valueKind := baseValueType.Kind()


		// reflect.TypeOf(t) == reflect.TypeOf((*Test)(nil)).Elem()

		if schemaKind.Outer == reflect.Invalid

		valueIterable := valueKind == reflect.Array || valueKind == reflect.Slice
		schemaIterable := schemaKind.Outer == reflect.Array || schemaKind.Outer == reflect.Slice

		if valueIterable && schemaIterable {
			innerValueKind := baseValueType.Elem().Kind()
			if !typesMatch(innerValueKind, schemaKind.Inner) {
				return false
			}

		} else if !typesMatch(valueKind, schemaKind.Outer) {
			return false
		}
	}

	return true
}

// func main() {
// 	projectSchema := map[string]NestedKind {
// 		"id": NestedKind { reflect.Int64, reflect.Invalid },
// 		"name": NestedKind { reflect.String, reflect.Invalid },
// 		"password": NestedKind { reflect.Array, reflect.Uint8 },
// 	}

// 	values := map[string]interface{} {
// 		"id": 4,
// 		"name": "stuff",
// 		"password": []byte("hello"),
// 	}

// 	fmt.Println(validatePatch(&values, &projectSchema))
// }
