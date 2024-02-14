package contactcontroller

import (
	"fmt"
	"galih/belajar-fiber/models/entity"
	"galih/belajar-fiber/models/request"
	"galih/belajar-fiber/models/response"
	"galih/belajar-fiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetAllContact(c *fiber.Ctx) error {
	var contacts []response.ContactResponse

	err := utils.DB.Model(&entity.Contact{}).Scan(&contacts).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get data",
		"data":    &contacts,
	})
}

func CreateContact(c *fiber.Ctx) error {
	req := new(request.CreateContactRequest)
	c.BodyParser(req)

	fmt.Println("req", req)
	// return c.JSON(a.Error());
	// get user from jwt
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	validate := validator.New()
	validate.RegisterValidation("unique_contact_name", utils.ValidateUniqueNameContact)
	err := validate.Struct(req)

	if err != nil {
		d := utils.HandleValidationError(err, req)
		return c.Status(fiber.StatusBadRequest).JSON(d)
	}

	contact := &entity.Contact{
		Name:      req.Name,
		Address:   req.Address,
		Handphone: req.Handphone,
		Email:     req.Email,
		UserID:    uint(claims["sub"].(float64)),
	}

	err = utils.DB.Create(&contact).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	// build response
	response := &response.ContactResponse{
		ID:        contact.ID,
		Name:      contact.Name,
		Email:     contact.Email,
		Handphone: contact.Handphone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
		UserID:    contact.UserID,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success created contact",
		"data":    response,
	})
}
