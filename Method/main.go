package main

import "fmt"

type Student struct {
	name string
	age  int
}

// Method
func (s Student) greeting() {
	fmt.Println("Welcome", s.name)
}

func (s Student) changeName1(name string) {
	fmt.Println("--> on changeName1, name changed to", name)
	s.name = name
}

// Method Pointer
func (s *Student) changeName2(name string) {
	fmt.Println("--> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var s1 = Student{"Cahyo Arif Andiyarto", 18}
	s1.greeting()

	var s2 = Student{"Ilham", 23}
	fmt.Println("s2 name", s2.name)

	s2.changeName1("Akbar")
	fmt.Println("s2 after changeName1", s2.name)

	s2.changeName2("Maulana")
	fmt.Println("s2 after changeName2", s2.name)

	// Pengaksesan method dari variabel objek pointer
	var s3 = &Student{"Malik", 31}
	s3.greeting()
}
