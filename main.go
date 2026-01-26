package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Name       string
	IsNaughty  bool
	Attendance int
}

var classData = make(map[int]Student)

func init() {

	classData[27] = Student{Name: "Ghost", IsNaughty: true, Attendance: 85}
	classData[1] = Student{Name: "Topper", IsNaughty: false, Attendance: 95}
	classData[50] = Student{Name: "LastBencher", IsNaughty: true, Attendance: 40}
}

func getsStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classData)
}

func main() {
	http.HandleFunc("/students", getsStudents)

	println("Server starting on localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
