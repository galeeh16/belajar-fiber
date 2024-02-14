package postcontroller

import (
	"galih/belajar-fiber/models/entity"
	"galih/belajar-fiber/models/request"
	"galih/belajar-fiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetAllPost(c *fiber.Ctx) error {
	var posts []entity.Post

	err := utils.DB.Model(&entity.Post{}).Scan(&posts).Error 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success get posts",
		"data": &posts,
	})
}

func CreatePost(c *fiber.Ctx) error {
	req := new(request.CreatePostRequest)
	c.BodyParser(req)

	// validasi
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		d := utils.HandleValidationError(err, req)
		return c.Status(fiber.StatusBadRequest).JSON(d)
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// create post
	post := &entity.Post{
		Title: req.Title,
		Description: req.Description,
		UserID: int(claims["sub"].(float64)),
	}

	err = utils.DB.Create(&post).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success create post",
		"data": &post,
	})
}