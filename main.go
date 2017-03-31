package main

import (
	"fmt"
	"io/ioutil"
  "strings"
)

type Pairs struct {
	Id   int
	Name string
}

type configStruct struct{
  column string
  typeName string
  loc string
}

func main() {
	bs, err := ioutil.ReadFile("test.gofig")
	if err != nil {
		panic(err)
	}
  fmt.Println(string(bs))
  GetStringBetween(string(bs),"{", "}")
}

func Decode(s string, &Interface{}, config configStruct){

}

func createConfigStruct(sArr []string){
  for
}

func GetStringBetween(s string, prefix string, suffix string) {
  fmt.Println(strings.Index(s, prefix), strings.Index(s, suffix))
  i := strings.Index(s, prefix)
  e := strings.Index(s, suffix)
  fmt.Println(s[i+1:e])
  arr := strings.Split(s[i+1:e], "\n")
  fmt.Println("line count: " , len(arr))
  var ar []string
  for _, s := range arr{
    if strings.TrimSpace(s) != ""{
      ar = append(ar,s)
    }
  }
  fmt.Println("line count: " , len(ar))
}
