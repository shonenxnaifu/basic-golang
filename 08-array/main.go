package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90, 95, 80,
	}

	fmt.Println(values)
	fmt.Println(len(names))
	fmt.Println(len(values))

	// panjang array pada go mengikuti jumlah array pada saat deklarasi/inisialisasi
	// tidak mengikuti jumlah isi array
	var lagi [10]string
	fmt.Println(len(lagi))
}
