# Flat File Parser

A library to parse fixed length flat files in Go. Specify the start and end location of the field in the "loc" struct tag.

In the IT industry were often face a challenge when we integrate our applications with Mainframe and AS400 systems that predominantly use fixed length flat files to read and write data. Most of the code used to read those flat files have very poor readability, and requires writing a lot of boiler plate code that is difficult to maintain.

Introducing Go Flat File Parser. This library was created to help parse the data in a neat way to read and parse data from a flat file and process it accordingly. This library helps write readable code that your future colleges won't hate to maintain.

Right now the library is designed to only load data read in a single line, ie. each line in the file is considered a distinct record.

There are lot of use cases where a record spans multiple lines, like in a shipment data flat file where, the lines that start H is the header of the shipment and the lines that start with D is the shipment details record are part of the same shipment.

```go
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
	Ti   time.Time `loc:"6,14" format:"YYYYMMDD"` //Use MS excel style date formats.
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

	fmt.Println("p:", p)

	if err := flatFileParser.DecodeFile("read.txt", &r); err != nil {
		panic(err)
	}
	fmt.Println("r:", r)
}

```
## Todo
- [ Support multiline pull ]
- [ Write Flat Files ]
