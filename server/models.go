//go:generate kallax gen

package main

import (
	"fmt"
	"time"
	"errors"
	"reflect"
	"github.com/gin-gonic/gin"
	"github.com/fatih/structs"
	"github.com/iancoleman/strcase"

	"gopkg.in/src-d/go-kallax.v1"
)


type Model struct {
	Id uint32 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	Model
	Name string `json:"name"`
	// https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address
	Email string `gorm:"type:varchar(254);unique;not null" json:"email"`
	// Slug string `gorm:"unique;not null" json:"slug"`
	Password string `gorm:"not null"`
}

type Project struct {
	Model
	Name string `json:"name"`
	Description string `json:"description"`

	User User `json:"user"`
	UserId uint32 `json:"user_id"`
}


var UnableToBindError error = errors.New("Couldn't bind json")
func BindJSONWithTemplate(c *gin.Context, giveValue *map[string]interface{}, templateStruct interface{}) error {
	if err := c.ShouldBindJSON(giveValue); err != nil {
		c.AbortWithError(422, UnableToBindError)
		return UnableToBindError
	}
	if alignsOk, alignsFailReason := checkMapAlignsWithStruct(giveValue, templateStruct); !alignsOk {
		c.AbortWithError(422, errors.New(alignsFailReason))
		return UnableToBindError
	}

	return nil
}


func checkMapAlignsWithStruct(checkedMap *map[string]interface{}, incomingStruct interface{}) (bool, string) {
	if !structs.IsStruct(incomingStruct) {
		return false, fmt.Sprintf("bad incomingStruct: %T", incomingStruct)
	}

	checkingStruct := structs.New(incomingStruct)
	for key, value := range *checkedMap {
		camelKey := strcase.ToCamel(key)

		structField, keyIsOk := checkingStruct.FieldOk(camelKey)
		if !keyIsOk {
			return false, fmt.Sprintf("bad key: %s", camelKey)
		}
		if value == nil {
			continue
		}

		assertedNestedMap, assertedOk := value.(map[string]interface{})
		if assertedOk {
			nestedOk, nestedFailReason := checkMapAlignsWithStruct(&assertedNestedMap, structField.Value())
			if !nestedOk {
				return false, nestedFailReason
			}
		} else if valueKind := reflect.TypeOf(value).Kind(); valueKind != structField.Kind() {
			return false, fmt.Sprintf("value of %s of type %T instead of %T", camelKey, valueKind, structField.Kind())
		}
	}

	return true, ""
}
