package flatFileParser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"errors"
	"io/ioutil"
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
					panic(err)
				}
				setValue(&fv, a[x:y])
			}
			//fmt.Println(a)
			interfaceSlice = reflect.Append(interfaceSlice, v)
		}
	}
	//i = interfaceSlice.Interface()
	fmt.Println(i, interfaceSlice)
	interfacePtrValue.Elem().Set(interfaceSlice)
	return nil
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
			return nil, nil, err
		}
		if y, err = strconv.Atoi(arr[1]); err != nil {
			return nil, nil, err
		}
		return x, y, nil
	} else {
		str := "location (Tag) format for " + sf.Name + ", tag:" + sf.Tag + " is wrong."
		return 0, 0, errors.New(str)
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

func DecodeFile(filePath string, i interface{}) {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	Decode(string(bs), i)
	//json.Unmarshal()
	//ioutil.ReadAll(strings.NewReader(string(bs)))
}