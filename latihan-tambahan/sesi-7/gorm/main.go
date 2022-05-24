package main

import (
	"fmt"
	"sesi-7-gorm/database"
	"sesi-7-gorm/models"
	"sesi-7-gorm/repository"
	"strings"

	"gorm.io/gorm"
)

func main() {
	db := database.StartDB()

	user(db)
	product(db)
}

func user(db *gorm.DB) {
	userRepo := repository.NewUserRepo(db)
	// Create User
	user := models.User{
		Email: "adit@gmail.com",
	}

	err := userRepo.CreateUser(&user)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	fmt.Println("Created Success")

	// Get All User
	employees, err := userRepo.GetAllUsers()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for k, emp := range *employees {
		fmt.Println("user :", k+1)
		emp.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	// Get User By ID
	emp, err := userRepo.GetUserByID(3)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	emp.Print()

	// Update User By ID
	var requestUpdate = models.User{
		Email: "ucup@gmail.com",
	}
	err = userRepo.UpdateUserByID(2, &requestUpdate)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	// Delete User By ID
	var id uint = 1
	err = userRepo.DeleteUserByID(id)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	fmt.Println("Delete user with id", id, "success")
}

func product(db *gorm.DB) {
	productRepo := repository.NewProductRepo(db)
	// Create Product
	product := models.Product{
		Name: "Celana",
		Brand: "H&M",
		UserID: 4,
	}

	err := productRepo.CreateProduct(&product)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	fmt.Println("Created Product Success")

	// Get All Product
	products, err := productRepo.GetAllProducts()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for k, product := range *products {
		fmt.Println("product :", k+1)
		product.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	// Get Product By ID
	emp, err := productRepo.GetProductByID(4)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	emp.Print()

	// Update Product By ID
	var requestUpdate = models.Product{
		Name: "Kemeja",
		Brand: "Nevada",
	}
	err = productRepo.UpdateProductByID(4, &requestUpdate)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	// Delete Product By ID
	var id uint = 3
	err = productRepo.DeleteProductByID(id)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	fmt.Println("Delete product with id", id, "success")
}
