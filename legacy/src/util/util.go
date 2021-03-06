package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

// JsonFileLoader is Json Load from File
func JsonFileLoader(file string, st interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return JsonLoader(&data, st)
}

// JsonLoader is Json Load from Memory
func JsonLoader(data *[]byte, st interface{}) error {
	err := json.Unmarshal(*data, st)
	if err != nil {
		return err
	}

	return nil
}

func Copy(arr interface{}) []interface{} {
	dest := make([]interface{}, 0)

	switch val := arr.(type) {
	case []interface{}:
	case []int:
		for _, v := range val {
			dest = append(dest, v)
		}
	case []string:
		for _, v := range val {
			dest = append(dest, v)
		}
	case int:
		dest = append(dest, val)
	case string:
		dest = append(dest, val)
	}

	return dest
}

func Contains(i interface{}, val interface{}) bool {
	arr := Copy(i)

	if reflect.TypeOf(arr[0]) != reflect.TypeOf(val) {
		log.Printf("missmatch Type in Contains [%s] and [%s]", reflect.TypeOf(arr[0]), reflect.TypeOf(val))
		panic(0)
	}

	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

/*
func AllOf(arr []Traits, val Traits) bool{
	for _, v := range arr{
		if v != val{
			return false
		}
	}
	return true
}

func NoneOf(arr []Traits, val Traits) bool{
	for _, v := range arr{
		if v == val{
			return false
		}
	}
	return true
}

func CountIf(arr []Traits, val Traits) int{
	n := 0
	for _, v := range arr{
		if v == val{
			n++
		}
	}
	return n
}

*/
/////////
