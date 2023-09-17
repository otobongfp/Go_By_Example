package main

import (
    "fmt"
    "time"
)

func main() {

    salary:=5000
    switch {
    case salary<=3000:
        fmt.Println(salary, "is for an entry level role")
    case salary<=7000:
        fmt.Println(salary, "is for an intermediate role")
    case salary>=7000:
        fmt.Println(salary, "is for a senior")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}