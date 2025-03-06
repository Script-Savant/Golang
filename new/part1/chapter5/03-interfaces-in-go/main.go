package main

import "fmt"

type Employee interface {
    GetName() string
}

type Engineer struct{
    Name string
}

func PrintDetails(e Employee){
    fmt.Println(e.GetName())
}

func (e *Engineer) GetName() string {
    return "Engineer Name: " + e.Name
}

func main(){
    engineer := &Engineer{Name: "Alex"}
    PrintDetails(engineer)
}
