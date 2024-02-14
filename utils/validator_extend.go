package utils

import (
	"fmt"
	"galih/belajar-fiber/models/entity"

	"github.com/go-playground/validator/v10"
)

func ValidatorUniqueUserID(fl validator.FieldLevel) bool {
	user := &entity.User{}
	userId := fl.Field().String() // inputan jsonnya
	err := DB.Model(&entity.User{}).Where("user_id = ?", userId).Take(&user).Error

	return err != nil // return true pass
}

type ContactValidation struct {
	Name string `json:"name"`
}

func (f *ContactValidation) TableName() string {
	return "contacts"
}

func ValidateUniqueNameContact(fl validator.FieldLevel) bool {
	contact := &ContactValidation{}
	name := fl.Field().String() // inputan jsonnya
	err := DB.Where("name = ?", name).Take(&contact).Error

	fmt.Println("err", err)

	return err != nil // return true pass
}
