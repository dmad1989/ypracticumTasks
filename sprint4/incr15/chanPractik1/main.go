package main

import (
    "fmt"
    "time"
)

func main() {
    chIn := make(chan int, 50)
    chOut := make(chan int)

    go func() {
        // горутина берёт числа из chIn
        for {
            left := <-chIn
            right := <-chIn
            // получаем два числа из chIn и записываем их сумму в chOut
            chOut <- left + right
        }
    }()

    go func() {
        // горутина берёт числа из chOut
        for {
            s := <-chOut
            // пусть обработка значений из chOut занимает какое-то время
            time.Sleep(20 * time.Millisecond)
            if s%10 == 1 {
                fmt.Printf("%d ", s)
            }
        }
    }()
    // отправляем сто чисел в канал chIn
    for i := 0; i < 100; i++ {
        chIn <- i
    }
    fmt.Printf("# ")
    time.Sleep(1000 * time.Millisecond)
}