package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func helperfunction(promt string, r *bufio.Reader)(string, error){
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return  strings.TrimSpace(input),err
}

func createbill() bill{
	read := bufio.NewReader(os.Stdin)
	
	name, _ := helperfunction("enter the name", read)

	b := billing(name)
	return b
}

func promtopsion (b bill){
	read := bufio.NewReader(os.Stdin)
	option, _ := helperfunction("options: \na for addtiems \ni for items \nt for tip \nenter your option:", read)
	if option == "t"{
		fmt.Println("enter the amount of tip")
	
	}else if option == "a"{
		fmt.Println("add items")
	}else if option == "i" {
		fmt.Println("do you want to see the i item bill?")
	}else {
	fmt.Println("not valid")
	}

}
func main() {

	mybill := createbill()

	promtopsion(mybill)
	fmt.Println(mybill)
	//fmt.Println(currentbilling)z

}