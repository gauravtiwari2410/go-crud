package controller

import (
	"fiber-crud/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "invalid data",
			})
	}

	if data["name"] =""{
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier name is required",
	})}
	if data["passcode"] =""{
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier passcode is required",
	})}


//now saving cashier to db
 cashier:= model.Cashier{

	Name: data["name"],
	Passcode: data["passcode"],
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
 }
 db.DB.Create(&cashier)
 return c.Status(200).JSON(fiber.Map{
	"success" :true,
	"message" : "cashier added successfully"
	"data" : cashier,
 })

}
func EditCashier(c *fiber.Ctx) error {
	return nil
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}
func CashierList(c *fiber.Ctx) error {
	return c.SendString("cashier list")
}
func GetCashierDetails(c *fiber.Ctx) error {
	return nil
}
func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
