package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Schools ms server up and running ...")
	for _, sch := range schools {
		fmt.Printf("Name: %v\n", sch.Name)
		fmt.Printf("Email: %v\n", sch.Email)
		fmt.Printf("Grades: %v\n", sch.Grades)
		fmt.Printf("Address: %v\n", sch.AddressID)
		fmt.Printf("Parents: %v\n", sch.Parents)
		fmt.Printf("Teachers: %v\n", sch.Teachers)
		fmt.Printf("Students: %v\n", sch.Students)
	}
}
