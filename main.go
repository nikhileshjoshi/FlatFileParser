package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	_"errors"
)

type Pairs struct {
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
	//createConfigStruct("test.gofig")
	p := Pairs{}

	t := reflect.TypeOf(p)
	ps := reflect.ValueOf(&p)
	v := ps.Elem()

	bs, err := ioutil.ReadFile("read.txt")
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(bs), "\n")

	for _, s := range arr{
		if strings.TrimSpace(s) != "" {
			for i := 0; i < v.NumField(); i++ {
				fv := v.Field(i)
				//ft := t.Field(i)
				x, y, err := getLoc(t.Field(i))
				if err != nil{
					panic(err)
				}
				setValue(&fv, s[x:y])
			}
		}

	}
	fmt.Println(p)


}

func getLoc(sf reflect.StructField) (int, int, error){
	arr := strings.Split(sf.Tag.Get("loc"), ",")
	if len(arr)==2{
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		return x,y, nil
	}else{
		//str := "location (Tag) format for " + sf.Name + //", tag:" + sf.Tag + " is wrong."
		return 0, 0, nil//errors.New(str)
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

func DecodeFile(filePath string, i interface{}, config []configStruct) {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	Decode(string(bs), i, config)

	//ioutil.ReadAll(strings.NewReader(string(bs)))
}

func Decode(s string, i interface{}, config []configStruct) {
	//json.Unmarshal()
	arr := strings.Split(s, "\n")
	for _, a := range arr {
		fmt.Println(a)
	}
}

func createConfigStruct(fileName string) []configStruct {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	sArr := GetStringBetween(string(bs), "{", "}")

	var config []configStruct
	var start, end int
	for _, s := range sArr {
		s = strings.Replace(s, "\t", " ", -1)
		arr := strings.Split(s, " ")
		if len(arr) == 3 {
			start, end = getStartAndEnd(arr[2])
		}
		config = append(config, configStruct{arr[0], arr[1], start, end})
	}
	return config
}

func GetStringBetween(s string, prefix string, suffix string) []string {
	fmt.Println(strings.Index(s, prefix), strings.Index(s, suffix))
	i := strings.Index(s, prefix)
	e := strings.Index(s, suffix)
	fmt.Println(s[i+1 : e])
	arr := strings.Split(s[i+1:e], "\n")
	fmt.Println("line count: ", len(arr))
	var ar []string
	for _, s := range arr {
		if strings.TrimSpace(s) != "" {
			ar = append(ar, s)
		}
	}
	fmt.Println("line count: ", len(ar))
	return ar
}

func getStartAndEnd(str string) (int, int) {
	var start, end int
	str = strings.TrimSpace(str)
	if len(str) == 3 {
		start, _ = strconv.Atoi(str[0:1])
		end, _ = strconv.Atoi(str[2:3])
	} else {
		//split the string by : and pull the last string and
	}
	return start, end
}
