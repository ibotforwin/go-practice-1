package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-ambdassador/src/database"
	"go-ambdassador/src/models"
	"golang.org/x/crypto/bcrypt"
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

	//Encrypt password, and the second return item is an error which we are ignoring.
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

	//Get user data from register request and fill model
	//(this is done if passwords match)
	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		Password:     password,
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "hello",
	})
}
