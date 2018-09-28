package main

import (
	"fmt"
	"strings"
	"reflect"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/blainehansen/goqu"
)

type column struct {
	i goqu.IdentifierExpression
}

func (c column) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c column) I() goqu.IdentifierExpression {
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


type arrayArg string
const (
	DIST_ANY arrayArg = "ANY"
	DIST_ALL arrayArg = "ALL"
)

func makeStringArrayLiteral(inputValues []string) goqu.LiteralExpression {
	var b strings.Builder

	b.WriteString("ARRAY [ ")

	for i, v := range inputValues {
		if i != 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "'%s'", v)
	}

	b.WriteString(" ]")

	return goqu.L(b.String())
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



type NestedKind struct {
	Outer reflect.Kind
	Inner reflect.Kind
	Instance interface{}
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

		if schemaKind.Outer == reflect.Struct {
			if reflect.TypeOf(schemaKind.Instance) != baseValueType {
				return false
			}
		}

		valueIterable := valueKind == reflect.Array || valueKind == reflect.Slice
		schemaIterable := schemaKind.Outer == reflect.Array || schemaKind.Outer == reflect.Slice

		if valueIterable && schemaIterable {
			innerValueKind := baseValueType.Elem().Kind()

			fmt.Println(value)

			if innerValueKind == reflect.Interface {
				// return true
				iterable := reflect.ValueOf(value)

				for iterableIndex := 0; iterableIndex < iterable.Len(); iterableIndex++ {
					fmt.Println()
					item := iterable.Index(iterableIndex)
					fmt.Println(item)

					itemKind := item.Kind()
					fmt.Println(itemKind)

					// if itemKind == reflect.Struct {
					// 	for structIndex := 0; structIndex < item.NumField(); structIndex++ {
					// 		fmt.Printf("%+v\n", item.Field(structIndex))
					// 	}
					// } else if !typesMatch(itemKind, schemaKind.Inner) {
					// 	return false
					// }
				}
			} else if !typesMatch(innerValueKind, schemaKind.Inner) {
				return false
			}

		} else if !typesMatch(valueKind, schemaKind.Outer) {
			return false
		}
	}

	return true
}




type patchExec struct {
	*goqu.CrudExec
	patchValid bool
	PatchMap map[string]interface{}
}

func (e *patchExec) Exec() (bool, sql.Result, error) {
	if !e.patchValid {
		return false, nil, nil
	}

	result, err := e.CrudExec.Exec()
	return true, result, err
}

func doPatch(c *gin.Context, p *patchExec) bool {
	patchValid, result, updateError := p.Exec()

	if !patchValid {
		c.AbortWithError(422, fmt.Errorf("invalid patch arguments: %s", p.PatchMap))
		return false
	}

	if updateError != nil {
		c.AbortWithError(500, updateError)
		return false
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		c.AbortWithStatus(404)
		return false
	}

	return true
}

func doExec(c *gin.Context, e *goqu.CrudExec) bool {
	result, updateError := e.Exec()
	if updateError != nil {
		c.AbortWithError(500, updateError)
		return false
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		c.AbortWithStatus(404)
		return false
	}

	return true
}
