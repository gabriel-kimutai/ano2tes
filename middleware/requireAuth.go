package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gabriel-kimutai/ano2tes/database"
	"github.com/gabriel-kimutai/ano2tes/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *fiber.Ctx) error {
	tokenString := c.Cookies("Authorization")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		log.Fatalf("Failed to parse token: %v", err)

	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Redirect("/login", 302)
		}
		var user models.User;
		result := database.DB.Db.First(&user, claims["user_id"])
		if result.RowsAffected <= 0 {
			c.Redirect("/login", 302)
		}
		c.Locals("user", user)
		return c.Next()
	}
	return c.Redirect("/login", 302)
}