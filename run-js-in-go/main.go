package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func main() {
	// Membuat mesin JavaScript
	vm := otto.New()

	// Menjalankan kode JavaScript
	script := `
function add(a, b) {
	return a + b;
}
add(5, 10);
	`

	result, err := vm.Run(script)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mengambil nilai dari hasil eksekusi JavaScript
	value, _ := result.ToInteger()
	fmt.Println("Hasil penjumlahan:", value)
}
