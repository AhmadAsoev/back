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

	var data []structs.TableRequest

	for rows.Next() {
		var row structs.TableRequest

		if err := rows.Scan(&row.ID, &row.ActNumber, &row.SenderName, &row.SenderPositon, &row.SenderOrganization, &row.Date, &row.ReceiverName, &row.ReceiverPosition, &row.ReceiverOrganization); err != nil {
			log.Print("Error scanning row: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "Ошибка при обработке данных"})
		}

		data = append(data, row)
	}
	if err := rows.Err(); err != nil {
		log.Print("Error iterating rows: ", err)
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка при обработке данных"})
	}

	return c.JSON(data)
}
