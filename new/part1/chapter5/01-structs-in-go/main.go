package main

import "fmt"

type Engineer struct {
    Name string
    Age int
    Project Project
}

type Project struct {
    Name string
    Priority string
    Technologies []string
}

func main(){
    println("Hello world")

    engineer := Engineer{
        Name: "Alex",
        Age: 29,
        Project: Project{
            Name: "Beginner's guide to Golang",
            Priority: "High",
            Technologies: []string {"Go", "Atom"},
        },
    }

    fmt.Printf("%+v\n", engineer)

    println("Engineer Name :", engineer.Name)
    fmt.Println("Project Technologies :", engineer.Project.Technologies)
}
