package data

import (
	"assignment-1/models"
	"time"
)
	
var Employees = []models.Employee{
	{Nama: "Adit", Alamat: "Batam", Pekerjaan: "Backend Engineer", Alasan: "Belajar hal baru", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Nama: "Andi", Alamat: "Bandung", Pekerjaan: "Mobile Engineer", Alasan: "Belajar hal baru", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Nama: "Febri", Alamat: "Bandung", Pekerjaan: "Frontend Engineer", Alasan: "Belajar hal baru", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}
