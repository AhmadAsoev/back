package authorization

import (
	"fmt"
	"log"

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

	data.Email = c.Params("email")
	data.Password = c.Params("password")
	db := repository.Conn()
	defer db.Close()

	var selectUserByEmail = "SELECT id, first_name, last_name, email, hash_pass, created_at, updated_at FROM USERS WHERE email=@email;"

	var userData structs.RegisterRequest

	if err := db.QueryRow(selectUserByEmail, sql.Named("email", data.Email)).Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Email, &userData.Password, &userData.CreatedAt, &userData.UpdatedAt); err != nil {
		log.Print(err)
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка при обработке данных"})
	}

	log.Print(userData)

	err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password))
	if err != nil {
		return c.JSON("Password does not match")
	} else {
		token, err := createToken(userData)
		if err != nil {
			log.Fatal("Error creating token:", err)
		}

		return c.JSON(token)
	}

}

var jwtKey = []byte("mySecretKey")

func createToken(user structs.RegisterRequest) (string, error) {
	// Устанавливаем время истечения токена (например, 1 час)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Создаём claims для токена (данные пользователя и время истечения)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", user.Id), // ID пользователя
		Issuer:    "myApp",
		ExpiresAt: expirationTime.Unix(),
	}

	// Создаём новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "Ошибка подписки токена с использование секретного ключа", err
	}

	return signedToken, nil
}
