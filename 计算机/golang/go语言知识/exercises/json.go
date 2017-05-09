package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := []byte(`{"action":"login","id":1234,"pwd":"1213546","devList":["1234","4567"]}`)
	var dat map[string]interface{}
	err := json.Unmarshal(str, &dat)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dat)
	fmt.Println(dat["action"])
	//fmt.Println(dat["action"].(type))
	//arrData := dat["devList"].Array()

	if arr, ok := dat["devList"].([]string); ok {
		fmt.Printf("arr:%T,%v\n", arr, arr)
	} else {
		fmt.Println(ok)
	}
	switch t := dat["devList"].(type) {
	case []interface{}:
		for i, vv := range t {
			fmt.Println(i, vv)
		}
	}
	for k, v := range dat {
		switch t := v.(type) {
		case []interface{}:
			for i, vv := range t {
				fmt.Println(i, vv)
			}
		case string:
			fmt.Println(k, v)
		}
	}

	//fmt.Println((dat["devList"].([]string))[0])
}
