package main

import (
	"encoding/json"
	"fmt"
	"log"
	"util"

	uuid "github.com/satori/go.uuid"
)

var categoryJSON = `
    {
        "server":[
            "momery",
            "processor",
            {"storage":["volume","volume1"]},
            {"xxx":[
                "yyy",
                {"kkk":["ttt","fff"]}
                ]
            }
        ]
    }
`

type devInfo struct {
	info     string
	children map[string]map[string]devInfo
}

var deviceInfo devInfo

func main() {
	log.SetFlags(log.Lshortfile)
	category := map[string]interface{}{}
	err := json.Unmarshal([]byte(categoryJSON), &category)
	if err != nil {
		log.Fatalf("json unmarshal error: % v ", err)
	}
	fmt.Println(category)
	fmt.Println(crateMap(category))
}

func crateMap(category map[string]interface{}) map[string]map[string]devInfo {
	dev := map[string]map[string]devInfo{}
	devUUID := uuid.Must(uuid.NewV4()).String()
	for k, v := range category {
		fmt.Printf("k:%s\n", k)
		var subDev map[string]map[string]devInfo
		sub, ok := v.([]interface{})
		if ok {
			subDev = map[string]map[string]devInfo{}
			for _, sv := range sub {
				if c, ok := sv.(string); ok {
					fmt.Printf("c:%s\n", c)
					devUUID := uuid.Must(uuid.NewV4()).String()
					di := devInfo{info: util.GetRandStrings(10), children: nil}
					du := map[string]devInfo{devUUID: di}
					subDev[c] = du
				} else {
					var subDevInfo map[string]map[string]devInfo
					if subCategory, ok := sv.(map[string]interface{}); ok {
						subDevInfo = crateMap(subCategory)
						for ddk, sdv := range subDevInfo {
							for kk, vv := range sdv {
								fmt.Printf("kk:%s\n", kk)
								du := map[string]devInfo{kk: vv}
								subDev[ddk] = du
							}
						}

					} else {
						log.Fatal("subCategory convert error")
					}
				}
			}

		} else {
			log.Printf("v is not []interface{} type,is %T", v)
		}
		di := devInfo{info: util.GetRandStrings(10), children: subDev}
		du := map[string]devInfo{devUUID: di}
		dev[k] = du
	}
	return dev
}
