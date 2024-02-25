package assets

import(
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func ShowBanner(){

	out := exec.Command("clear")
	out.Stdout = os.Stdout
	out.Run()
	
	// Buka file
    file, err := os.Open("assets/banner.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Buat scanner untuk membaca file baris per baris
    scanner := bufio.NewScanner(file)

    // Loop melalui setiap baris file
    for scanner.Scan() {
        // Tampilkan baris
        fmt.Println(scanner.Text())
    }

    // Periksa kesalahan yang mungkin terjadi selama pemindaian
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
    }
}