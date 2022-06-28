package main

import (
	"fmt"
	"go-gorm/config"
	"go-gorm/model"
	"go-gorm/repository"
	"log"
)

func main() {
	config := config.NewConfigDB()
	db := config.DbConn()
	defer config.DbClose()

	// err := db.AutoMigrate(&model.Customer{})
	// if err != nil {
	// 	panic(err)
	// }

	repo := repository.NewCustomerRepository(db)
	// Insert
	// customer := model.Customer{
	// 	Id:      uuid.New().String()[:8],
	// 	Name:    "Dzikrur R",
	// 	Address: "Sragen",
	// 	Phone:   "081229xxx",
	// 	Email:   "dzikrurrxxx@gmail.com",
	// 	Balance: 1000,
	// }
	// repo.Create(&customer)
	// Update
	customerExisting := model.Customer{
		Id:      "1945f1a2",
		Name:    "Dzikrur R",
		Address: "Sragen",
	}
	// STRUCT with
	// err := repo.Update2(&customerExisting, model.Customer{
	// 	Address: "Sragen",
	// 	Balance: 15000000,
	// })
	// MAP With
	// err := repo.Update(&customerExisting, map[string]interface{}{
	// 	"Address": "Klaten",
	// 	"Balance": 10000000,
	// 	"is_status": 1,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// Delete
	// err := repo.Delete(&customerExisting)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// FindById
	// repo.FindById(customerExisting.Id)
	fmt.Println(repo.FindById(customerExisting.Id))

	// FindAllBy
	// var customers []model.Customer
	// customers, err := repo.FindAllBy(map[string]interface{}{
	// 	"address": "klaten",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)
	// FindFirstBy
	// customers, err := repo.FindFirstBy(map[string]interface{}{
	// 	"address": "klaten",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)
	// FindBy
	customers, err := repo.FindBy("name LIKE ? AND is_status = ?", "%R%", 0)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(customers)

}
