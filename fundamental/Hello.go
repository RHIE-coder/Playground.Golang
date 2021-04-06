package main

import "fmt"

func main() {

	myMap := map[int]string{
		10: "TEN",
		5:  "FIVE",
		3:  "THREE",
	}

	fmt.Println(myMap) //map[3:THREE 5:FIVE 10:TEN]

	val, isExist := myMap[10]

	fmt.Println(val, isExist) // TEN true

	for key, val := range myMap {
		fmt.Println(key, " : ", val)
	}

	val, isExist = myMap[11]

	fmt.Println(val, isExist) //   false
	if val == "" {            //val은 ""
		fmt.Println("val은 \"\"")
	}

}
