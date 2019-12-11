package main

import (
	"fmt"
)

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}, val interface{}) {
	d.data[key] = val
}

func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

func (d *Dictionary) Traversal(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}

	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

func main() {

	dict := NewDictionary()
	dict.Set("good1", 29)
	dict.Set("good2", 49)
	dict.Set("good3", 69)
	language := dict.Get("language")
	if language != nil {
		fmt.Println("language is", language)
	}
	dict.Traversal(func(k, v interface{}) bool {
		if v.(int) >= 40 {
			fmt.Println(k, "is expensive")
			return true
		}

		fmt.Println(k, "is cheap")
		return true
	})
}
