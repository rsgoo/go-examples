package main

import "fmt"

func main() {
	studentMap := map[string]map[string]interface{}{}

	studentMap["tom"] = map[string]interface{}{}
	studentMap["tom"]["name"] = "tom cat"
	studentMap["tom"]["age"] = "2"
	studentMap["tom"]["gender"] = "male"

	studentMap["jerry"] = map[string]interface{}{}
	studentMap["jerry"]["name"] = "jerry mouse"
	studentMap["jerry"]["age"] = "3"
	studentMap["jerry"]["gender"] = "male"

	for name, studentInfo := range studentMap {
		fmt.Println(name)
		for key, value := range studentInfo {
			fmt.Println(key, ":", value)
		}
	}

}
