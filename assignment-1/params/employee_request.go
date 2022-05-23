package params

type EmployeeRequest struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func NewEmployeeRequest(nama, alamat, pekerjaan, alasan string) *EmployeeRequest {
	return &EmployeeRequest{
		Nama:      nama,
		Pekerjaan: pekerjaan,
		Alamat:    alamat,
		Alasan:    alasan,
	}
}
