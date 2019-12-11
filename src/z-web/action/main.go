package main

import (
	"net/http"
	"html/template"
)

type Employee struct {
	Id       int
	LastName string
	Email    string
}

func main() {
	//http.HandleFunc("/testif", testIF)
	//http.HandleFunc("/testif", testWith)
	//http.HandleFunc("/testif", testTemplate)
	http.HandleFunc("/testif", defineTemplate)
	http.ListenAndServe(":8080", nil)
}

func defineTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("define.html"))
	t.ExecuteTemplate(w, "model", "")
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	t.Execute(w, "<<我能在两个文件中显示吗？>>")
}

func testRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))

	//languages := []string{
	//	"Golang",
	//	"PHP",
	//	"Rust",
	//}

	employeeList := []Employee{}

	employee1 := Employee{
		Id:       1,
		LastName: "LINUS",
		Email:    "linus@linux.con",
	}
	employeeList = append(employeeList, employee1)

	employee2 := Employee{
		Id:       2,
		LastName: "aaron",
		Email:    "aaron@linux.con",
	}
	employeeList = append(employeeList, employee2)

	employee3 := Employee{
		Id:       3,
		LastName: "lee",
		Email:    "lee@linux.con",
	}
	employeeList = append(employeeList, employee3)

	t.Execute(w, employeeList)
}

func testWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, "王子")
}
