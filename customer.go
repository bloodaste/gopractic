package main

import "fmt"

type bill struct {
	customer     string
	ordereditems map[string]float64
	tip          float64
}

func billing(name string ) bill {
	b := bill{
		customer:     name,
	
	}
	return b
}

func (b bill) format() {
	fs := "your breakdowon items :\n"
	var total float64 = 0; 
	for index, v := range b.ordereditems{
		fs += fmt.Sprintf("%v: %v \n" , index , v)
		total += v
	}
	// the first value of fs is "breakdwon items "then it replace the fs with v and the index 
	total += b.tip
	fmt.Println(b.customer)
	fmt.Printf("tip: %v" , b.tip)
	fmt.Printf("Thank you %v for your order %v,with %v tip your total is %0.2f",b.customer,fs , b.tip,total)

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