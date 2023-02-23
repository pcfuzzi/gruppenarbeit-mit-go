package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    var answer int = rand.Intn(100) + 1
    var guess int

    fmt.Println("Willkommen zum Zahlraten-Spiel!")
    fmt.Println("Ich denke an eine Zahl zwischen 1 und 100. Kannst du sie erraten?")

    for {
        fmt.Print("Dein Tipp: ")
        fmt.Scanln(&guess)

        if guess < answer {
            fmt.Println("Dein Tipp ist zu niedrig. Versuche es nochmal.")
        } else if guess > answer {
            fmt.Println("Dein Tipp ist zu hoch. Versuche es nochmal.")
        } else {
            fmt.Println("Glückwunsch! Du hast die Zahl erraten!")
            break
        }
    }
}
