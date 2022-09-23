package data

import "fmt"

type Student struct {
	Name       string
	Address    string
	Profession string
	Reason     string
}

func (s Student) PrintInfo() {
	fmt.Printf("Nama\t\t: %s\n", s.Name)
	fmt.Printf("Alamat\t\t: %s\n", s.Address)
	fmt.Printf("Pekerjaan\t: %s\n", s.Profession)
	fmt.Printf("Alasan\t\t: %s\n", s.Reason)
}
