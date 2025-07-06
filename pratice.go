package main

// func saygreeting(name string){
// 	fmt.Printf("goodmorning %v \n", name)
// }
// func goodbye(name string){
// 	fmt.Printf("goodbye %v \n", name)
// }

// func cyclename(n []string , f func(string)){
// 	for _, value := range n{
// 		f(value)
// 	}
// }

// func circleradius(num float64 ) float64{
// 	return math.Pi * num *num
// }

// func getinitial(n string) (string, string){
// 	fmt.Println("the name entered is:",n)
// 	s := strings.ToUpper(n)
// 	name := strings.Split(s, " ")

// 	 var initial  []string

// 	 for _, value := range name{
// 		initial = append(initial,value[:1])
// 	 }
// 	 if len(initial ) > 1{
// 		return initial[0], initial[1]
// 	 }else{
// 		return initial[0] , "_"
// 	 }
// }
func pratice() {
// string
	// var nameone string = "mario"
	// var nametwo = "luigi"
	// var namethree string

	// fmt.Println(nameone,nametwo, namethree)

	// nameone = "bowser"
	// nametwo = "peach"
	// fmt.Println(nameone,nametwo, namethree)

	// namefour := 4

	// fmt.Println(nameone,nametwo, namethree,namefour)

	// var age int = 30
	// var age3 = 40
	// age4 := 50

	// fmt.Println(age,age3,age4);

	// //bit and memori

	// var nums int8 = 18
	// var num  int8 = -128
	// var num7 uint8 =  25
	// fmt.Println(nums, num7, num);

	//    age := 25
	//    name := "kurt"
	// 	fmt.Print("hello \n")
	// 	fmt.Print("world")
	// 	fmt.Println("my name is ", name," my age is \n " ,age)

	// 	fmt.Printf("my name is %s my age is %d \n",name , age )
	// 	fmt.Printf("type of age is %T", name)
	// 	fmt.Printf("type of age is %0.1f \n", 10.26)

	// 	//sprtinf

	// 	var nametwo = fmt.Sprintf("my name is %s my age is %d \n",name , age )

	// 	fmt.Println(nametwo)

	//var ages [3]int = [3]int{10, 25, 30}

// 	var ages = [3]int{100, 25, 30}
// 	 name := [4]string {"kurt", "andre", "gutierrez", "mabait"}
// 	fmt.Println(ages, len(ages))
// 	fmt.Println(name,len(name))

// 	sort.Ints(ages[:])
// 	fmt.Println(ages)

// 	var scores = []int{100, 50,200,2,5 ,7, 300}
	
// 	scores = append(scores, 58)
// 	fmt.Println(scores)

// 	ranges :=  scores[1:3]
// 	rangestwo :=  scores[1:]
// 	fmt.Println(ranges)
// 	fmt.Println(rangestwo)

	
// 	sort.Ints(scores)
// 	fmt.Println(scores)
// 	sort.Strings(name[:])
// 	fmt.Println(name)
// 	gretting := "hello friend"
// fmt.Println("Output:",strings.Contains(gretting, "hello"))
// fmt.Println(strings.ReplaceAll(gretting, "hello", "hi"))
// fmt.Println(strings.ToUpper(gretting))
// fmt.Println(strings.Index(gretting,"f"))
// fmt.Println(strings.Split(gretting, "") )
// 	num := []int {30,40,50,20,5}
// 	sort.Ints(num)
// 	fmt.Println(num)
// 	index := sort.SearchInts(num, 30)
// 	fmt.Println(index)


//  	for i := 0 ; i<=len(num) ; i++{
// 	fmt.Println("the value of i:", i)

//    }
//    for index, value := range num{
// 	 fmt.Printf("possiton index %d value of the index %d \n", index, value)
// 	 num[index] += 10
 
//    }
//    fmt.Println(num)


   //	saygreeting("kurt")
//	goodbye("kurt")
	// cyclename([]string{"kurt", "andrei", "gutieerzz"}, saygreeting)

	// cyclename([]string{"kurt", "andrei", "gutieerzz"}, goodbye)	
	// rad := circleradius(210)

	// fmt.Printf(" the radius of the circle is :%0.2f" ,rad)	
	// fn ,sn:=  getinitial("kurtandrei")
	//  fmt.Printf("first name letter %s \nsecond name first letter is %s", fn,sn)
	// package main

	// import "fmt"
	
	// func updatingname(x *string) {
	// 	*x = "hasbeen change"
		
	// }
	// func updatemap(x map[string]float64 ){
	// 	x["coffe"] = 47.99
	// }
	// func main() {
	// 	soup:= map[string]float64{
	// 		"soup": 64.99,
	// 		"pie": 7.49,
	// 	}
	
	// 	fmt.Println(soup)
	// 	fmt.Println(soup["pie"])
	
	// 	for index, value := range soup{
	// 		fmt.Println(index , ":",value)
	// 	}
	
	// 	updatemap(soup)
	// 	fmt.Println(soup)
	// 	name := "kurt"
		
	// 	m := &name
		
	
		
	
	// 	fmt.Println(m)
	// 	fmt.Println("deferrence in pointers", *m)
		
	// 	fmt.Println(name)
	// 	updatingname(m)
	// 	fmt.Println(*m )
	// 	fmt.Println(name)
		
	// }
}