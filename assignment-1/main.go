package main

import (
	"assignment-1/data"
	"assignment-1/params"
	"assignment-1/repositories"
	"fmt"
	"strconv"
	"time"
)

func main() {
	var input string
	var condition = true

	req1 := params.NewEmployeeRequest("Kevin", "Jogja", "Mahasiswa", "Belajar hal baru")
	req2 := params.NewEmployeeRequest("Bobby", "Bekasi", "Mahaiswa", "Belajar hal baru")
	req3 := params.NewEmployeeRequest("Zafran", "Bandung", "Backend Developer", "Belajar hal baru")

	data.Employees = append(data.Employees, repositories.InsertEmployee(req1))
	data.Employees = append(data.Employees, repositories.InsertEmployee(req2))
	data.Employees = append(data.Employees, repositories.InsertEmployee(req3))

	for condition {
		fmt.Println("Program Assignment 1 Hacktiv8")
		fmt.Println("1. Tampilkan semua data")
		fmt.Println("2. Tampilkan berdasarkan id")
		fmt.Println("3. Masukkan Data")
		fmt.Println("4. Keluar")
		fmt.Println("Inputkan Pilihanmu :")
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println("\033[2J")
			fmt.Println("List semua data employee")
			for i, employee := range data.Employees {
				fmt.Print(i+1, ". ")
				employee.PrintAll()
			}
			fmt.Println()
			fmt.Println("Lanjutkan Program (Y/N)? :")
			fmt.Scanln(&input)
			if input == "n" || input == "N" {
				condition = false
			}
		case "2":
			var id string
			fmt.Println("\033[2J")
			fmt.Print("Masukkan Id : ")
			fmt.Scanln(&id)

			fmt.Println("Data Employee dengan id :", id)

			ind, err := strconv.ParseInt(id, 10, 32)
			if err == nil {
				data.Employees[ind-1].PrintDetail()
			}
			fmt.Println()
			fmt.Println("Lanjutkan Program (Y/N)? :")
			fmt.Scanln(&input)
			if input == "n" || input == "N" {
				condition = false
			}
		case "3":
			var nama, alamat, pekerjaan, alasan string
			fmt.Println("Input employee baru")
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan alamat : ")
			fmt.Scanln(&alamat)
			fmt.Print("Masukkan pekerjaan : ")
			fmt.Scanln(&pekerjaan)
			fmt.Print("Masukkan alasan : ")
			fmt.Scanln(&alasan)
			request := params.NewEmployeeRequest(nama, alamat, pekerjaan, alasan)
			data.Employees = append(data.Employees, repositories.InsertEmployee(request))
		default:
			condition = false
		}
	}
	fmt.Println("Program akan keluar dalam 3 detik...")
	time.Sleep(3 * time.Second)
}
