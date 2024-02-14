package authcontroller

import (
	"galih/belajar-fiber/models/entity"
	"galih/belajar-fiber/models/request"
	"galih/belajar-fiber/models/response"
	"galih/belajar-fiber/utils"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func Login(c *fiber.Ctx) error {
	request := new(request.LoginRequest)
	c.BodyParser(request)

	validate := validator.New()
	errValidator := validate.Struct(request)

	if errValidator != nil {
		d := utils.HandleValidationError(errValidator, request)
		return c.Status(fiber.StatusBadRequest).JSON(d)
	}

	// get user 
	var user response.UserResponse
	err := utils.DB.Take(&entity.User{}, "user_id = ?", request.UserID).Scan(&user).Error

	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"endpoint": c.Path(),
			"method": c.Method(), 
			"error": err.Error(),
		}).Warn("user not found")
	
		return c.Status(401).JSON(fiber.Map{
			"message": "User or password are invalid",
			"data": nil,
		})
	}

	key := os.Getenv("JWT_SECRET")
	ttlDuration := time.Duration(1) * time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "mydomain.com",
		"sub": user.ID,
		"exp": time.Now().Add(ttlDuration).Unix(),
	})

	accessToken, _ := token.SignedString([]byte(key))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": accessToken,
		"data": &user,
	})
} 