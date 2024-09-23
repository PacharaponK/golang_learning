package main

import (
    "fmt"
)

type Person struct {
    name   string
    age    int
    job    string
    salary int
}

func main() {
    var pers1 Person
    var pers2 Person

    // Pers1 specification
    pers1.name = "Bom"
    pers1.age = 20
    pers1.job = "Programmer"
    pers1.salary = 25000

    // Pers2 specification
    pers2 = Person{
        name:   "Earth",
        age:    19,
        job:    "Developer",
        salary: 30000,
    }

    // Print Pers1 info by calling a function
    printPerson(pers1)

    // Print Pers2 info by calling a function
    printPerson(pers2)
}

func printPerson(pers Person) {
    fmt.Println("Name: ", pers.name)
    fmt.Println("Age: ", pers.age)
    fmt.Println("Job: ", pers.job)
    fmt.Println("Salary: ", pers.salary)
    fmt.Println("Performance: ", pers.getPerformance())
}

func (person Person) getPerformance() string {
    indicator := person.salary / person.age
    if indicator > 10000 {
        return "Excellent"
    } else if indicator > 5000 {
        return "Good"
    } else if indicator > 2500 {
        return "Average"
    } else if indicator > 1250 {
        return "Fair"
    } else {
        return "Poor"
    }
}
