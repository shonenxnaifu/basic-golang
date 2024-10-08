package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("run application before logging")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Println("Application ended")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error on Application")
	}

	// ketika panic code ini tidak akan dieksekusi
	// tapi code lain diluar function tetap dieksekusi
	fmt.Println("Application running")
}

func main() {
	//runApplication()
	runApp(true)
	fmt.Println("runn")
}
