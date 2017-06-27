package main

import (
	"fmt"
	"github.com/nikhileshjoshi/flatFileParser"
	"io/ioutil"
	"time"
)

type Pair struct {
	Id   int       `loc:"0,3"`
	Name string    `loc:"3,6"`
	Ti   time.Time `loc:"6,14" format:"YYYYMMDD"`
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
	fmt.Println("p:", p[0], p)

	if err := flatFileParser.DecodeFile("read.txt", &r); err != nil {
		panic(err)
	}
	fmt.Println("r:", r[0])
}
