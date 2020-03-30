package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    slice := generateSlice(4)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    bubblesort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice)
}

// generate slice untuk membuat data yang belum terurut dengan random angka
func generateSlice(size int) []int {
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}
func bubblesort(items []int) {
    var (
        n      = len(items)
        sorted = false
    )
    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
                items[i+1], items[i] = items[i], items[i+1]
                swapped = true
            }
        }
        /* jika status swapped tidak sama dengan false, ganti status sorted menjadi true, maka proses for diberhentikan*/
        if !swapped {
            sorted = true
        }
        n = n - 1 // sebelum proses for utama di berhentikan, nilai n selalu dikurangi 1
    }
}
