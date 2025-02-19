package services

import (
	"github.com/cloudscode/via-svit/database"
	"github.com/cloudscode/via-svit/models"
)

func GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Find(&orders)
	return orders, err
}

func CreateOrder(order *models.Order) error {
	_, err := database.DB.Insert(order)
	return err
}

func UpdateOrder(order *models.Order) error {
	_, err := database.DB.ID(order.ID).Update(order)
	return err
}

func DeleteOrder(id int64) error {
	_, err := database.DB.ID(id).Delete(&models.Order{})
	return err
}
