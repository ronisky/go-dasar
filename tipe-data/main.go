package main

import "fmt"

func main() {

	// Tipe data numerik Non-desimal
	var positiveNumber uint = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	
	// Tipe data numerik desimal
	
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	
	// Tipe data bool / boolean

	var exist bool =true
	fmt.Printf("exist? %t \n", exist)


	// Tipe data string

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var newMessage = `Nama saya "Jhon Wick".
	Salam kenal.
	Mari belajar 'Golang'.`

	fmt.Println(newMessage)

	
	// Nilai nil & Zero Value
	// nil bukan merupakan tipe data, melainkan sebuah nilai. 
	// variabel yang isi nilainya nil berarti memiliki nilai kosong

	// semua tipe data memiliki zero value (nilai default tipe data)
	/*
	- Zero value dari string adalah "" (string kosong)
	- Zero value dari bool adalah false
	- Zero value dari tipe numerik non-desimal dalah 0
	- Zero value dari tipe numerik desimal adalah 0.0

	Zero value berbeda dengan nil. Nil adalah nilai kosng, benar-benar kosong. 
	
	Ada beberapa tipe data nyang bisa di-set nilainya dengan nil, di antaranya:
	- pointer
	- tipe data fungsi
	- slice
	- map
	- channel
	- intface kosong atau interface{}
	*/
}