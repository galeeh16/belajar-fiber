package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// handle error message
//
// example result error:
// {
// 	field1: errorMessage1,
// 	field2: errorMessage2,
// 	fieldN: errorMessageN
// }
func HandleValidationError(err error, dataStruct interface{}) map[string]string {
	errorList := make(map[string]string)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(dataStruct).Elem().FieldByName(err.StructField())
			fieldName := field.Tag.Get("json")

			errorList[fieldName] = GetErrorMessage(err, fieldName, dataStruct)	
        }
	}

	return errorList
}

func GetErrorMessage(err validator.FieldError, fieldName string, dataStruct interface{}) string {
	f, _ := reflect.TypeOf(dataStruct).Elem().FieldByName(err.Param())
	fieldNameParam := f.Tag.Get("json")
	
	switch err.Tag() {
		case "required":
			return fmt.Sprintf("%s is required", fieldName)
		case "min":
			return fmt.Sprintf("%s must be at least %s characters", fieldName, err.Param())
		case "max":
			return fmt.Sprintf("%s must be at least %s characters", fieldName, err.Param())
		case "eqfield": 
			return fmt.Sprintf("%s and %s must match", fieldName, fieldNameParam)
		case "email": 
			return fmt.Sprintf("%s must be valid email address", fieldName)
		case "unique_user_id", "unique_contact_name": 
			return fmt.Sprintf("%s has been taken", fieldName)
		default: 
			// return "message did not mapping"
			return err.Error()
	}
}