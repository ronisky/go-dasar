package main

import "fmt"

func main() {
	//menggunakan var, tanpa tipe daya, menggunakan perantara "="
	var firstName = "jhon"

	/*
	Ketika deklarasi, tipe data yg digunakan harus dituliskan juga. 
	Istilah lain dari konsep ini adalah "manifest typing"
	*/

	// var firstName string ="jhon"
	// var lastName string
	// lastName = "wick"

	/*
	Go juga mengadopsi konsep "type inference", 
	yaitu metode deklarasi variabel yang tipe data-nya 
	ditentukan oleh tipe data nilainya
	*/

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"
	/* deklarasi menggunakan tanda ":=",
	untuk assigment nilai harus menggunakan tanda "=" */
	// lastName = "Erone"
	//Deklarasi menggunakan ":=" hanya bisa dilakukan di dalam blk fungsi


	// Multi Variabel
	/*
	var first, second, thrid string
	first, second, thrid = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	*/

	// Variable Underscore _

	/*
	Biasanya variabel underscore sering dimanfaatkan untuk menampung 
	nilai balik fungsi yang tidak digunakan.

	Perlu diketahui, bahwa isi variabel underscore tidak dapat ditampilkan. 
	Data yang sudah masuk variabel tersebut akan hilang.

	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "jhon", "wick"
	*/

	// Deklarasi variabel menggunakan keyword "new"
	// digunakan untuk membuat variabel pointer dengan tipe data tertentu, 
	// nilai default-nya akan menyesuaikan tipe datanya.

	name := new(string)
	fmt.Println(name) //0xc000044250
	fmt.Println(*name) //""

	// Deklarasi variable menggunaan keyword "make"
	/*
		Keyword ini hanya bisa digunakan untuk pembuatan beberapa jenis variabke saja, yaitu:
		- channel
		- slice
		- map
	*/

	
	fmt.Printf("helo jhon wick!\n")
	fmt.Printf("helo %s %s!\n", firstName, lastName)
	fmt.Println("helo", firstName, lastName + "!")
}
