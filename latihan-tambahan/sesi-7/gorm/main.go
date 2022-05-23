package main

import (
	"fmt"
	"sesi-7-gorm/database"
	"sesi-7-gorm/models"
	"sesi-7-gorm/repository"
	"strings"
)

func main() {
	db := database.StartDB()
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
