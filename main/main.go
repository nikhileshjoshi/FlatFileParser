package main

import (
	_ "errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"github.com/nikhileshjoshi/flatFileParser"
)

type Pair struct {
	Id   int    `loc:"0,3"`
	Name string `loc:"3,6"`
}

type configStruct struct {
	column   string
	typeName string
	start    int
	end      int
}

func main() {

	bs, err := ioutil.ReadFile("read.txt")
	if err != nil {
		panic(err)
	}

	var p []Pair
	flatFileParser.Decode(string(bs), &p)
	//Decode(string(bs), &p)
	//p := i.([]Pair)
	fmt.Println(p)

}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "FlatFile: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "FlatFile: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "FlatFile: Unmarshal(nil " + e.Type.String() + ")"
}

func getLoc(sf reflect.StructField) (int, int, error) {
	arr := strings.Split(sf.Tag.Get("loc"), ",")
	if len(arr) == 2 {
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		return x, y, nil
	} else {
		//str := "location (Tag) format for " + sf.Name + //", tag:" + sf.Tag + " is wrong."
		return 0, 0, nil //errors.New(str)
	}
}

func setValue(t *reflect.Value, value string) {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, _ := strconv.Atoi(value)
		t.SetInt(int64(i))
	case reflect.String:
		t.SetString(value)
	case reflect.Bool:
		if len(value) == 1 {
			if value == "0" {
				value = "false"
			} else {
				value = "true"
			}
		}
		b, _ := strconv.ParseBool(value)
		t.SetBool(b)
	}
}



