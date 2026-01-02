package main

import "fmt"

type Student struct {
	Name       string
	IsNaughty  bool
	Attendance int
}

func main() {

	classData := make(map[int]Student)

	// Adding data
	classData[27] = Student{Name: "Ghost", IsNaughty: true, Attendance: 85}
	classData[1] = Student{Name: "Topper", IsNaughty: false, Attendance: 95}
	classData[50] = Student{Name: "LastBencher", IsNaughty: true, Attendance: 40}

	fmt.Println("THE GHOST STUDENT DATA:", classData[27].Name, classData[27].IsNaughty, classData[27].Attendance)
	fmt.Println("THE TOPPER STUDENT DATA:", classData[1].Name, classData[1].IsNaughty, classData[1].Attendance)
	fmt.Println("THE LAST BENCHER STUDENT DATA:", classData[50].Name, classData[50].IsNaughty, classData[50].Attendance)
}
