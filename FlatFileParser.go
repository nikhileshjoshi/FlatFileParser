/*
MIT License

Copyright (c) 2017 Nikhilesh Joshi <reachnikhilesh@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package flatFileParser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"
	"github.com/metakeule/fmtdate"
)

func Decode(s string, i interface{}) error {
	arr := strings.Split(s, "\n")
	interfacePtrValue := reflect.ValueOf(i)
	if interfacePtrValue.Kind() != reflect.Ptr || interfacePtrValue.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(i)}
	}
	t := reflect.TypeOf(i).Elem().Elem()
	interfaceSlice := interfacePtrValue.Elem()
	//interfaceSlice := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	for _, a := range arr {
		if strings.TrimSpace(a) != "" {

			//v := reflect.ValueOf(interfaceValue)
			v := reflect.New(t).Elem()

			for i := 0; i < t.NumField(); i++ {
				fv := v.Field(i)
				//ft := t.Field(i)
				x, y, err := getLoc(t.Field(i))
				if err != nil {
					return err
				}
				err = setValue(&fv, a[x:y], t.Field(i).Tag.Get("format"))
				if err != nil{
					return err
				}
			}
			//fmt.Println(a)
			interfaceSlice = reflect.Append(interfaceSlice, v)
		}
	}
	//i = interfaceSlice.Interface()
	//fmt.Println(i, interfaceSlice)
	interfacePtrValue.Elem().Set(interfaceSlice)
	return nil
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
	var x, y int
	var err error
	arr := strings.Split(sf.Tag.Get("loc"), ",")
	if len(arr) == 2 {
		if x, err = strconv.Atoi(arr[0]); err != nil {
			return 0, 0, err
		}
		if y, err = strconv.Atoi(arr[1]); err != nil {
			return 0, 0, err
		}
		return x, y, nil
	} else {
		str := fmt.Sprint("location (Tag) format for ", sf.Name, ", tag:", sf.Tag, " is wrong.")
		return 0, 0, errors.New(str)
	}
}

func setValue(t *reflect.Value, value string, format string) error{
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.Atoi(value)
		if err != nil{
			return errors.New(fmt.Sprint("Error when converting value ", value, " Err: ", err ))
		}
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
	case reflect.Struct:
		if t.Type().String() == "time.Time" {
			t.Set(reflect.ValueOf(parseTime(value, format)))
		}

	}
	return nil
}

func parseTime(str string, format string) time.Time{

	t, err := fmtdate.Parse(format, str)
	fmt.Println("str:", str, "format:", format, "t:", t)

	if err!= nil{
		panic(err)
	}

	return t
}

func DecodeFile(filePath string, i interface{}) error {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return Decode(string(bs), i)
}
