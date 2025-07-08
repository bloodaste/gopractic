package main

import (
	"fmt"
	"os"
)

type bill struct {
	customer     string
	ordereditems map[string]float64
	tip          float64
}

func billing(name string ) bill {
	b := bill{
		customer:     name,
		ordereditems: map[string]float64{},
		tip: 0,
	
	}
	return b
}

func (b bill) format() string{
	fs := "your breakdowon items :\n"
	var total float64 = 0; 
	for index, v := range b.ordereditems{
		fs += fmt.Sprintf("%v: %v \n" , index , v)
		total += v
	}
	// the first value of fs is "breakdwon items "then it replace the fs with v and the index 

	fs += fmt.Sprintf("tip: %v" , b.tip)
	fs += fmt.Sprintf("Thank you %v, your total is %.2f\n", b.customer, total)
	return fs
}


func (b *bill) changetip(t float64){	 
	b.tip = t 

}

func (b bill) additems(i string, p float64){
	b.ordereditems[i] = p

}
func (b *bill) changecustomersname(n string){
	b.customer = n
}

func (b *bill) savefun(){
 data := []byte(b.format()) 

 err := os.WriteFile("bill/"+b.customer+".csv" , data, 0777)

 if err != nil{
	 panic(err)
 }

}