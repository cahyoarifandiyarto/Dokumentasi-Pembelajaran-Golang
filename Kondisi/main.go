package main

import "fmt"

func main() {
	// if, else if, else
	var nilai = 80

	if nilai == 100 {
		fmt.Println("Selamat anda mendapat nilai sempurna")
	} else if nilai > 75 {
		fmt.Println("Selamat anda lulus")
	} else {
		fmt.Println("Anda tidak lulus ujian. Belajar terus ya...")
	}

	// Variabel temporary if - else
	var point = 2

	if nilai := point * 25; nilai == 100 {
		fmt.Println("Lulus dengan nilai 100")
	} else if nilai >= 50 {
		fmt.Println("Selamat anda lulus dengan nilai minumun")
	} else {
		fmt.Println("Maaf anda tidak lulus :(")
	}

	// switch - case
	var role = "superadmin"

	switch role {
	case "admin":
		fmt.Println("Anda login sebagai admin")
	case "superadmin":
		fmt.Println("Anda login sebagai superadmin")
	default:
		fmt.Println("Anda belum login. Silahkan login")
	}
}
