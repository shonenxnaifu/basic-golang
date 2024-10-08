package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// data dari slice bersifat mutable
	// dapat diubah walaupun sudah diinisialisasi
	// months[5] = "diubah"

	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	
	var slice2 = months[11:]
	fmt.Println(slice2)

	// append akan membuat slice baru
	// sehingga data hasil dari append 
	// tidak merefer ke slice yang udah diinisialisasi sebelumnya (months)
	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 3, 5)

	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	newSlice[2] = "Khannedy"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice perlu mendefinisikan len dan cap dari slice yang akan dicopy
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
