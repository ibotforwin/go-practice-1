package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go-ambdassador/src/database"
	"go-ambdassador/src/middlewares"
	"go-ambdassador/src/models"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	//If c.BodyParser generates an error, return error. If err == nil, do nothing.
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		//If the password does not match
		//Set status to 400
		c.Status(400)
		//We use fiber.Map to fill response body with the error message and
		//set the JSON of c to that response body.
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	//Get user data from register request and fill model
	//(this is done if passwords match)
	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		IsAmbassador: false,
	}
	user.SetPassword(data["password"])
	database.DB.Create(&user)

	return c.JSON(&user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	//If c.BodyParser generates an error, return error. If err == nil, do nothing.
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid credentials"})
	}

	if user.CheckPassword(data["password"]) != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "Invalid credentials"})
	}

	fmt.Println(user.Id)

	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
