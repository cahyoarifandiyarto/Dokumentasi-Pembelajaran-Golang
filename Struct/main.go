package main

import "fmt"

func main() {
	// Inisialisasi variabel struct
	var s1 = Student{}
	s1.id = 1
	s1.name = "Cahyo Arif Andiyarto"
	s1.age = 18
	s1.grade = 12

	// Variabel objek pointer
	var s2 *Student = &s1 // -> Variabel s2 mereference alamat memori pada variabel s1
	fmt.Println("student 1, name:", s1.name)
	fmt.Println("student 1, grade:", s1.grade)
	fmt.Println("student 2, name:", s2.name)

	s2.name = "Ilham"
	fmt.Println("student 1, name:", s1.name)
	fmt.Println("student 1, age:", s2.grade)
	fmt.Println("student 2, name:", s2.name)

	// Anonymous struct
	var bio = struct {
		Person
		address string
	}{}
	bio.Person = Person{1, "Ahmad", 23}
	bio.address = "Sleman"
	fmt.Println("ID:", bio.Person.id)
	fmt.Println("Nama:", bio.Person.name)
	fmt.Println("Umur:", bio.Person.age)
	fmt.Println("Domisili:", bio.address)
}

// Struct Student
type Student struct {
	grade  int
	Person // -> Embbeded Struct dari Struct Person
}

// Struct Person
type Person struct {
	id   int
	name string
	age  int
}
