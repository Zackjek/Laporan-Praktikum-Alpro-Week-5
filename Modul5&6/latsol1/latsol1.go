package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)

    bil := 0
    for i := 1; i <= n; i++ {
        bil += i
    }

    fmt.Printf("Jumlah dari 1 sampai %d adalah %d\n", n, bil)
}
