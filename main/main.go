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

func main() {
	bs, err := ioutil.ReadFile("read.txt")
	if err != nil {
		panic(err)
	}

	var p, r []Pair
	flatFileParser.Decode(string(bs), &p)
	//Decode(string(bs), &p)
	//p := i.([]Pair)
	fmt.Println("p:", p[1])

	flatFileParser.DecodeFile("read.txt", &r)
	fmt.Println("r:", r[0])
}
