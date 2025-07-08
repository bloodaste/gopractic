package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func helperfunction(promt string, r *bufio.Reader)(string, error){
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return  strings.TrimSpace(input),err
}

func createbill() bill{
	read := bufio.NewReader(os.Stdin)
	
	name, _ := helperfunction("enter the name: ", read)

	b := billing(name)
	return b
}

func promtopsion (b *bill){
	read := bufio.NewReader(os.Stdin)
	option, _ := helperfunction("options: \na for addtiems \ni for items \nt for tip \nenter your option:", read)
	if option == "t"{
		tip, _ := helperfunction("enter tipp",read)
		tipv, err := strconv.ParseFloat(tip,64)
		if err != nil {
			fmt.Print("invalid numbers")
			promtopsion(b)
		}
		b.changetip(tipv)
		promtopsion(b)
		
	}else if option == "a"{
		namep, _ := helperfunction("additems: ",read)
		price, _ := helperfunction("the price of the item: ", read)

		 p, err := strconv.ParseFloat(price,64)
		 if err != nil{
			fmt.Println("you cant input letters in price")		
			promtopsion(b)
		 } else {
			b.additems(namep,p)
		 }
		fmt.Println("you add items")
		promtopsion(b)
	}else if option == "s" {
		fmt.Println("you want to save the fill?")	
		b.savefun()
		fmt.Println("you save the bill with", b.customer)
	}else {
	fmt.Println("not valid")
	promtopsion(b)
	}

}
func main() {

	mybill := createbill()

	promtopsion(&mybill)
	fmt.Println(mybill)
//mybill.format()
	//fmt.Println(=currentbilling)z

}