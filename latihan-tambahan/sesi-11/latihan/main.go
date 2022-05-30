package main

import (
	"fmt"
	"sesi-11/helpers"
)

/*
	Soal 1.
	Buatlah sebuah function untuk menampilkan
	apakah seseorang itu boleh bekerja atau tidak, dengan
	syarat :
		1. Laki laki dengan umur > 17 th < 60 th
		2. Perempuan dengan umur > 19 th < 60 th

	lalu silahkan buat function untuk unit testnya, dengan data berikut :
	func TestPria
		data :
			- Pria, umur 16th => expected false
			- Pria, umur 17th => expected false
			- Pria, umur 50th => expected true
			- Pria, umur 60th => expected false

	func TestWanita
		data :
			- Wanita, umur 16th => expected false
			- Wanita, umur 21th => expected true
			- Wanita, umur 50th => expected true
			- Wanita, umur 60th => expected false

	Soal 2
	Budi ingin bermain permainan pertarungan. Yang mana, setiap budi berhasil
	mengalahkan seorang petarung(A), maka budi akan mendapatkan kekuatannya (B)
	jadi kekuatan Budi akan terus bertambah seiring berhasil mengalahakn petarung.
	Silahkan cari berapa kekuatan terakhir budi, dengan syarat budi tidak bisa
	mengalahkan petarung yang lebih kuat dari dia.

	contoh:
		case 1:
			petarung : [5,3,9,8]
			kekuatan : [2,2,3,1]
			budi : 3
			nilaiAkhirExpected: 7

		case 2:
			petarung : [2,6,3,9]
			kekuatan : [2,2,3,5]
			budi : 2
			nilaiAkhirExpected: 14
*/
func main() {
	res := helpers.Soal2([]int{5, 3, 9, 8}, []int{2, 2, 3, 1}, 3)
	fmt.Print(res)
}
