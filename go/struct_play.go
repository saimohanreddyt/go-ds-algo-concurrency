package main

import (
	"fmt"
)

type User struct {
	Name    string
	Email   string
	Age     int
	Addr    string
	MoNo    int
	ActVity string
}

func main() {

	//	syntax
	ias := User{"IAS", "ias@google.com", 22, "Chennai", 9876543210, "SoftWare"}
	fmt.Printf("%v\n", ias)
	fmt.Println("")
	//if we want to print vales along with varaible using +
	fmt.Printf("%+v\n", ias)
	fmt.Println("")
	//if we want to print specific images
	fmt.Printf("%+v\n", ias.Email)
	fmt.Println("")

	//syntax
	tias := User{"TIAS", "tias@google.com", 22, "Chennai", 9876543201, "SoftWare"}
	fmt.Printf("%+v\n", tias)
	fmt.Println("")

	//syntax
	var mias = new(User)
	mias.Name = "mias"
	mias.Email = "mias@google.com"
	mias.Age = 23
	mias.MoNo = 9876543021
	mias.ActVity = "CharttedAccoutntant"
	//syntax
	fmt.Printf("%+v\n", mias)
	fmt.Println("")
	var iasm = &User{"IASM", "iasm@google.com", 22, "Chennai", 9876540321, "SoftWare"}
	fmt.Printf("%+v\n", iasm)
}
