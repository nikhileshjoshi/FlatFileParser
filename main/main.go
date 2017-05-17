package main

import (
	"fmt"
	"github.com/nikhileshjoshi/flatFileParser"
	"io/ioutil"
)

type Pair struct {
	Id   int    `loc:"0,3"`
	Name string `loc:"3,6"`
}

type Config struct {
	PoNumber    int64   `loc:0,6`
	StyleNumber string  `loc:6,15`
	UnitPrice   float64 `loc:15,20`
	Qty         int     `loc:10,15`
}

func main() {
	bs, err := ioutil.ReadFile("read.txt")
	if err != nil {
		panic(err)
	}

	var p, r []Pair
	if err := flatFileParser.Decode(string(bs), &p); err != nil {
		panic(err)
	}
	//Decode(string(bs), &p)
	//p := i.([]Pair)
	fmt.Println("p:", p[1], p)

	if err := flatFileParser.DecodeFile("read.txt", &r); err != nil {
		panic(err)
	}
	fmt.Println("r:", r[0])
}
