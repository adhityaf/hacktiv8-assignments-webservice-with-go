package repositories

import (
	"assignment-1/models"
	"assignment-1/params"
	"time"
)

func InsertEmployee(req *params.EmployeeRequest) (employee models.Employee) {
	employee.Nama = req.Nama
	employee.Alamat = req.Alamat
	employee.Pekerjaan = req.Pekerjaan
	employee.Alasan = req.Alasan
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	return
}
