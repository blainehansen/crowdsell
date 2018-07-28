package main

// import (
// 	"fmt"
// 	"strings"
// 	// "errors"
// 	"github.com/fatih/structs"
// 	"github.com/iancoleman/strcase"
// 	"gopkg.in/src-d/go-kallax.v1"
// )

// var projectSchemaStruct = structs.New(Schema.Project)
// var userSchemaStruct = structs.New(Schema.User)

// func GetProjectPatch(patchMap map[string]interface{}, patchProject *Project) ([]kallax.SchemaField, error) {
// 	return GetPatch(patchMap, patchProject, projectSchemaStruct)
// }

// func GetUserPatch(patchMap map[string]interface{}, patchUser *User) ([]kallax.SchemaField, error) {
// 	return GetPatch(patchMap, patchUser, userSchemaStruct)
// }


// func GetPatch(patchMap map[string]interface{}, patchObject interface{}, schemaStruct *structs.Struct) ([]kallax.SchemaField, error) {
// 	if !structs.IsStruct(patchObject) {
// 		return nil, fmt.Errorf("bad patchObject: %T", patchObject)
// 	}

// 	reflectAblePatchObject := structs.New(patchObject)

// 	var patchColumns []kallax.SchemaField
// 	for mapKey, mapValue := range patchMap {
// 		camelKey := strcase.ToCamel(mapKey)
// 		if strings.Contains(camelKey, "Id") || strings.Contains(camelKey, "Slug") {
// 			return nil, fmt.Errorf("bad key: %s", camelKey)
// 		}

// 		schemaStructField, keyIsOk := schemaStruct.FieldOk(camelKey)
// 		if keyIsOk {
// 			patchColumns = append(patchColumns, schemaStructField.Value().(kallax.SchemaField))
// 			err := reflectAblePatchObject.Field(camelKey).Set(mapValue)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			return nil, fmt.Errorf("bad key: %s", camelKey)
// 		}
// 	}

// 	return patchColumns, nil
// }
