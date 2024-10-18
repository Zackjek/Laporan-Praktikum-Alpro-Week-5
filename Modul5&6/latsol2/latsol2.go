package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Print("Masukkan jumlah kerucut: ")
    fmt.Scan(&n)

    // Loop untuk memasukkan jari-jari dan tinggi, lalu menghitung volume
    for i := 1; i <= n; i++ {
        var jariJari, tinggi float64
        fmt.Printf("Masukkan jari-jari alas kerucut ke-%d: ", i)
        fmt.Scan(&jariJari)
        fmt.Printf("Masukkan tinggi kerucut ke-%d: ", i)
        fmt.Scan(&tinggi)

        // Menghitung volume kerucut
        volume := (1.0 / 3.0) * math.Pi * math.Pow(jariJari, 2) * tinggi
        fmt.Printf("Volume kerucut ke-%d adalah: %.2f\n", i, volume)
    }
}
