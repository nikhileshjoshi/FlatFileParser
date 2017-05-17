#Flat File Parser

A library to parse fixed length flat files in Go. Specify the start and end location of the field in the "loc" struct tag.

In the IT industry were often face a challenge when we integrate our applications with Mainframe and AS400 systems that predominantly use fixed length flat files to read and write data. Most of the code used to read those flat files have very poor readability, and requires writing a lot of boiler plate code that is difficult to maintain.

Introducing Go Flat File Parser. This library was created to help parse the data in a neat way to read and parse data from a flat file and process it accordingly. This library helps write readable code that your future colleges won't hate to maintain.

Right now the library is designed to only load data read in a single line, ie. each line in the file is considered a distinct record.

There are lot of use cases where a record spans multiple lines, like in a shipment data flat file where, the lines that start H is the header of the shipment and the lines that start with D is the shipment details record are part of the same shipment.

```go
import (
	"fmt"
	"github.com/nikhileshjoshi/flatFileParser"
	"io/ioutil"
)

//The struct fields should start with a CAPITAL letter so that the fields are exported.
type PO struct {
	PoNumber    int64   `loc:0,6`
	StyleNumber string  `loc:6,15`
	UnitPrice   float64 `loc:15,20`
	Qty         int     `loc:10,15`
}

func main(){
  //Create an empty slice.
  var pos, strPO []PO

  //Use the DecodeFile to decode the whole file to a struct slice.
  if err := flatFileParser.DecodeFile("PO_extract.txt", &pos); err != nil {
		panic(err)
	}
	fmt.Println("POS:", pos)

  //Incase you have the flat file data in memory as a string use the Decode function to decode the data to a struct slice.
  bs, err := ioutil.ReadFile("read.txt")
	if err != nil {
		panic(err)
	}
  if err := flatFileParser.Decode(string(bs), &pos); err != nil {
		panic(err)
	}
	fmt.Println("POS:", pos)
}

```
##Todo
- [ Support multiline pull ]
- [ Support time parsing ]
- [ Write Flat Files ]
