package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/gabriel-kimutai/ano2tes/database"
	"github.com/gabriel-kimutai/ano2tes/models"
	"github.com/gabriel-kimutai/ano2tes/views"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func User(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	return views.Render(c, views.Base(views.UserPage(user.UserName)))
}

func LoginPage(c *fiber.Ctx) error {
	return views.Render(c, views.Login())
}

func UserCreate(c *fiber.Ctx) error {
	user := new(models.User)
	password := c.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}
	user = &models.User{
		Password: string(hashedPassword),
	}
	result := database.DB.Db.Create(&user)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.Redirect("/login", 302)
}

func UserLogin(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string
		Password string
	}
	email := c.FormValue("email")
	fmt.Println(email)

	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}
	user := new(models.User)
	result := database.DB.Db.First(&user, "email = ?", input.Email)
	if result.RowsAffected <= 0 {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	if !CheckPasswordHash(user.Password, input.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.UserName,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:        "Authorization",
		Value:       tokenString,
		Expires:     time.Now().Add(time.Hour * 72),
		SessionOnly: true,
	})
	return c.Redirect("/user", 302)
}

func UserLogout(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.Redirect("/")
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}
