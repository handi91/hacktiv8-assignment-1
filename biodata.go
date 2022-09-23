package main

import (
	"biodata/data"
	"fmt"
	"os"
	"strconv"
)

type Student = data.Student

/* Mapping absent number -> Student */
var students = make(map[int]Student, 0)

func init() {
	var studentsData = []Student{
		{
			Name:       "Handi",
			Address:    "Jalan Salak",
			Profession: "Mahasiswa",
			Reason:     "Mau nambah skill",
		},
		{
			Name:       "Didit",
			Address:    "Jalan Apel",
			Profession: "Guru SD",
			Reason:     "Bosan jadi guru",
		},
		{
			Name:       "Ica",
			Address:    "Jalan Sawo",
			Profession: "Content Creator",
			Reason:     "Pengen bisa koding",
		},
		{
			Name:       "Dimas",
			Address:    "Jalan Manggis",
			Profession: "Freelancer",
			Reason:     "Mau punya pekerjaan tetap di tech company",
		},
	}

	enrollStudent(studentsData)
}

func main() {
	absentNum, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Nomor absen tidak valid")
		return
	}

	student, isExist := getStudentByAbsentNumber(absentNum)

	if !isExist {
		fmt.Println("Data tidak ditemukan")
		return
	}

	student.PrintInfo()
}

func getStudentByAbsentNumber(absentNumber int) (student Student, isExist bool) {
	if absentNumber <= 0 || absentNumber > len(students) {
		return
	}

	student = students[absentNumber]
	isExist = true

	return
}

func enrollStudent(studentList []Student) {
	for i, v := range studentList {
		students[i+1] = v
	}
}
