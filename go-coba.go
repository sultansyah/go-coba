package go_coba

import "fmt"

func Login(nama string) {
	if nama == "sultan" {
		fmt.Println("anda berhasil login")
	} else {
		fmt.Println("anda gagal login")
	}
}
