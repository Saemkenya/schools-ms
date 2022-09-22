package main

import (
	"encoding/json"
	"net/http"
)

var (
	schools   []School
	addresses []Address
	users     []User
	grades    []Grade
	students  []Student
)

// init func run when build is invoked
func init() {
	addresses = []Address{
		{
			1,
			"0700 000 000",
			"123",
			"00100",
			"Nairobi",
			"047",
			"Nairobi",
			"Kenya",
		},
		{
			2,
			"0711 000 000",
			"456",
			"20303",
			"Ol'Kalou",
			"018",
			"Nyandarua",
			"Kenya",
		},
	}
	users = []User{
		{
			1,
			"Saem Ndisho",
			"saem@187ca.co.ke",
			addresses[1],
		},
		{
			2,
			"Elan Wambo",
			"elan@187ca.co.ke",
			addresses[0],
		},
	}
	grades = []Grade{
		{
			1,
			"First Grade",
			"NORTH",
			students,
			1,
		},
	}
	students = []Student{
		{
			1,
			"Lloyd Ng'ang'a",
			grades[0],
			nil,
			1,
		},
	}
	schools = []School{
		{
			1,
			"Elimu Nairobi Academy",
			"elim_nai@187ca.co.ke",
			1,
			grades,
			[]User{users[1]},
			[]User{users[0]},
			students,
		},
	}

}

// Address stores all fields required for entity addresses
type Address struct {
	ID         int    `json:"id"`
	Tel        string `json:"tel"`
	PoBox      string `json:"po_box"`
	Code       string `json:"code"`
	Town       string `json:"town"`
	CountyCode string `json:"county_code"`
	CountyName string `json:"county_name"`
	Country    string `json:"country"`
}

// School defines the fields of a school entity
type School struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AddressID int       `json:"address_id"`
	Grades    []Grade   `json:"grades"`
	Teachers  []User    `json:"teachers"`
	Parents   []User    `json:"parents"`
	Students  []Student `json:"students"`
}

// User is any entity that can auth in to the system
type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

// Grade is a class representation in a shoold
type Grade struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Stream     string    `json:"stream"`
	Population []Student `json:"population"`
	SchoolID   int       `json:"school_id"`
}

// Student represents a pupil belonging to a school
// Student is an entity that can auth in to the system
type Student struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	CurrentGrade Grade   `json:"current_grade"`
	PassedGrades []Grade `json:"passed_grades"`
	SchoolID     int     `json:"school_id"`
}
