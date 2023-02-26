package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func SerilizingJSON() {
	data := []byte(`{
		"ID":134,
		"Number":"ИЛМ-1274",
		"Year":2,
		"Students":[
			{
				"LastName":"Вещий",
				"FirstName":"Лифон",
				"MiddleName":"Вениаминович",
				"Birthday":"4апреля1970года",
				"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
				"Phone":"+7(948)709-47-24",
				"Rating":[1,2,3]
			},
			{
				"LastName":"Какач",
				"FirstName":"Кекы",
				"MiddleName":"Вениаминович",
				"Birthday":"4апреля1970года",
				"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
				"Phone":"+7(948)709-47-24",
				"Rating":[5,4,8,5,4]
			}
		]
	}`)
	r := strings.NewReader(string(data))
	data, _ = io.ReadAll(r)
	type Student struct {
		Rating []int `json:"Rating"`
	}
	type Group struct {
		// Name   string
		// Age    int
		// Status bool
		Students []Student `json:"Students"`
	}
	type AvgRatings struct {
		Avg float64 `json:"Average"`
	}

	// fmt.Println("\"beautified\" version")
	// data, err = json.MarshalIndent(s, "", "\t")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Printf("%s\n", data)

	fmt.Println("using pointer to empty struct")
	var mtGroup Group
	if err := json.Unmarshal(data, &mtGroup); err != nil {
		fmt.Println(err)
		return
	}
	amountStuds := float64(len(mtGroup.Students))
	amountRatings := 0.0
	for _, student := range mtGroup.Students {
		amountRatings += float64(len(student.Rating))
	}
	var avg AvgRatings = AvgRatings{Avg: amountRatings / amountStuds}
	marshed, _ := json.MarshalIndent(avg, "", "    ")
	fmt.Println(string(marshed))

	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var (
		src = testStruct{Name: "John Connor", Age: 35} // структура с данными
		dst = testStruct{}                             // структура без данных
		buf = new(bytes.Buffer)                        // буфер для чтения и записи
	)

	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(buf)

	fmt.Printf("buf before encode %v\n", buf)
	fmt.Printf("src before encode %+v\n", src)
	fmt.Printf("dst before encode %+v\n", dst)
	fmt.Println("-----------------------------")

	enc.Encode(src)
	fmt.Printf("buf after encode %v\n", buf)
	fmt.Printf("src after encode %+v\n", src)
	fmt.Printf("dst after encode %+v\n", dst)
	fmt.Println("-----------------------------")

	dec.Decode(&dst)
	fmt.Printf("buf after decode %v\n", buf)
	fmt.Printf("src after decode %+v\n", src)
	fmt.Printf("dst after decode %+v\n", dst)
}
