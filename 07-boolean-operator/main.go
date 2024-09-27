package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	var lulusujian = ujian >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusAbsensi && lulusujian
	fmt.Println(lulus)
}
