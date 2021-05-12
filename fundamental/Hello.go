package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Str string `json:"str"`
	Num int    `json:"num"`
}

func main() {
	dataToJson := Data{"Alice", 100}

	fmt.Println(dataToJson)

	//인코딩
	jsonBytes, err := json.Marshal(dataToJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonBytes)

	//디코딩
	var jsonToData Data
	err = json.Unmarshal(jsonBytes, &jsonToData)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonToData)
}
