package main

import (
    "fmt"
)

func main() {
    var dasar, pangkat int
    fmt.Print("Masukkan bilangan dasar: ")
    fmt.Scan(&dasar)
    fmt.Print("Masukkan bilangan pangkat: ")
    fmt.Scan(&pangkat)

    Hasil := 1
    for i := 0; i < pangkat; i++ {
        Hasil *= dasar
    }

    fmt.Printf("Hasil dari %d dipangkatkan dengan %d adalah %d\n", dasar, pangkat, Hasil)
}
