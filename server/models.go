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

	// "gopkg.in/src-d/go-kallax.v1"
)


type Model struct {
	Id int64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	Model
	Name string `json:"name"`
	Email string `gorm:"type:varchar(254);unique;not null" json:"email"`
	Slug string `gorm:"unique;not null" json:"slug"`
	InternalSlug string `gorm:"unique;not null" json:"internal_slug"`
	Password []byte `gorm:"not null"`

	ProfilePhotoSlug string `json:"profile_photo_slug"`
}

type Project struct {
	Model
	Name string `json:"name"`
	Description string `json:"description"`
	Slug string `gorm:"unique;not null" json:"slug"`
	InternalSlug string `gorm:"unique;not null" json:"internal_slug"`

	User User `json:"user"`
	UserId int64 `json:"user_id"`
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
