package tables

import (
	// "encoding/json"

	"log"

	"github.com/gofiber/fiber/v3"

	"back/cmd/hadlers/structs"
	"back/internal/repository"
)

func GatTables(c fiber.Ctx) error {

	db := repository.Conn()
	defer db.Close()
	rows, err := db.Query("SELECT id, act_number, sender_name, sender_position, sender_organization, date, receiver_name, receiver_position,receiver_organization FROM ACTS;")
	if err != nil {
		log.Print(err)
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка при выполнении запроса"})
	}
	defer rows.Close()

	// Создаём массив для хранения результатов
	var data []structs.TableRequest

	// Сканируем строки в структуру
	for rows.Next() {
		var row structs.TableRequest

		if err := rows.Scan(&row.ID, &row.ActNumber, &row.SenderName, &row.SenderPositon, &row.SenderOrganization, &row.Date, &row.ReceiverName, &row.ReceiverPosition, &row.ReceiverOrganization); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Ошибка при обработке данных"})
		}

		data = append(data, row)
	}

	// Возвращаем JSON с результатами

	return c.JSON(data)
}
