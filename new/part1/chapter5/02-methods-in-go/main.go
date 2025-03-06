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

func (e Engineer) Print(){
    println("Engineer Information")
    println("Name :", e.Name)
    println("Age :", e.Age)
    println("Current Project :", e.Project.Name)
}

func (e *Engineer) UpdateAge() {
    e.Age += 1
}

func (e *Engineer) GetProjectPriority() string{
    result := "Project Priority is : " + e.Project.Priority
    return result
}

func main(){
    fmt.Println("Methods in Go")
    engineer := Engineer{
        Name: "Alex",
        Age: 28,
        Project: Project{
            Name: "Beginner's guide to Golang",
            Priority: "High",
            Technologies: []string {"Go", "Atom"},
        },
    }

    engineer.UpdateAge()
    engineer.Print()

    fmt.Println(engineer.GetProjectPriority())
}
