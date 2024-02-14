package usercontroller

import (
	"fmt"
	"galih/belajar-fiber/models/entity"
	"galih/belajar-fiber/models/request"
	"galih/belajar-fiber/models/response"
	"galih/belajar-fiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	var users []response.UserResponse
	// var users []entity.User
	err := utils.DB.Preload("Contacts").Find(&users).Error

	fmt.Println("err", err)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error,
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get users",
		"data":    &users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	request := new(request.CreateUserRequest)
	c.BodyParser(request)

	validate := validator.New()
	validate.RegisterValidation("unique_user_id", utils.ValidatorUniqueUserID)
	errValidator := validate.Struct(request)

	if errValidator != nil {
		d := utils.HandleValidationError(errValidator, request)
		return c.Status(fiber.StatusBadRequest).JSON(d)
	}

	// hashing password
	password, errPass := utils.HashPassword(request.Password)
	if errPass != nil {
		return errPass
	}

	user := &entity.User{
		UserID:   request.UserID,
		Name:     request.Name,
		Password: password,
	}

	utils.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success create user",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	req := new(request.UpdateUserRequest)

	c.BodyParser(req)

	validate := validator.New()
	errValidator := validate.Struct(req)

	if errValidator != nil {
		d := utils.HandleValidationError(errValidator, req)
		return c.Status(fiber.StatusBadRequest).JSON(d)
	}

	var user response.UserResponse

	// find user by id
	err := utils.DB.Take(&entity.User{}, "id = ?", id).Scan(&user).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"data":    nil,
		})
	}

	// update nama user
	user.Name = req.Name
	// save
	utils.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success update user",
		"data":    &user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entity.User

	// find user by id
	err := utils.DB.Take(&user, "id = ?", id).Scan(&user).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"data":    nil,
		})
	}

	utils.DB.Delete(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    nil,
		"message": "Success delete user",
	})
}
