package authorization

import (
	"fmt"
	"log"

	"encoding/json"
	"github.com/gofiber/fiber/v3"

	"back/cmd/hadlers/structs"
	"back/internal/repository"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(c fiber.Ctx) error {
	var data structs.LoginRequest

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		log.Println("Ошибка при парсинге JSON:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Невозможно обработать JSON",
		})
	}

	db := repository.Conn()
	defer db.Close()

	var selectUserByEmail = "SELECT id, first_name, last_name, email, hash_pass, created_at, updated_at FROM USERS WHERE email=@email;"

	var userData structs.RegisterRequest

	if err := db.QueryRow(selectUserByEmail, sql.Named("email", data.Email)).Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Email, &userData.Password, &userData.CreatedAt, &userData.UpdatedAt); err != nil {
		log.Print(err)
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка при обработке данных"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Password does not match",
		})
	} else {
		token, err := createToken(userData)
		if err != nil {
			log.Fatal("Error creating token:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate token",
			})
		}

		log.Print(token)
		return c.JSON(fiber.Map{
			"token": token,
		})
	}

}

var jwtKey = []byte("mySecretKey")

func createToken(user structs.RegisterRequest) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", user.Id),
		Issuer:    "myApp",
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "Ошибка подписки токена с использование секретного ключа", err
	}

	return signedToken, nil
}
