package models

import (
	"fmt"
	"time"
)

type Employee struct {
	Nama, Alamat, Pekerjaan, Alasan string
	CreatedAt, UpdatedAt            time.Time
}

func (e *Employee) PrintAll() {
	fmt.Printf("Nama\t\t: %s,\n   Alamat\t: %s,\n   Pekerjaan\t: %s,\n   Alasan\t: %s,\n   CreatedAt\t: %v,\n   UpdatedAt\t: %v\n", e.Nama, e.Alamat, e.Pekerjaan, e.Alasan, e.CreatedAt, e.UpdatedAt)
}

func (e *Employee) PrintDetail() {
	fmt.Printf("Nama\t\t: %s,\nAlamat\t\t: %s,\nPekerjaan\t: %s,\nAlasan\t\t: %s,\nCreatedAt\t: %v,\nUpdatedAt\t: %v\n", e.Nama, e.Alamat, e.Pekerjaan, e.Alasan, e.CreatedAt, e.UpdatedAt)
}
