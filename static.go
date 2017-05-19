package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type staticStoreType map[string]map[string]interface{}

func StaticStoreLoad(jsonFile string) staticStoreType {
	var staticStore staticStoreType
	inBuf, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(inBuf, &staticStore)
	if err != nil {
		fmt.Println(staticStore)
		panic(err)

	}

	return staticStore
}

func StaticStoreGetCredsForPlanID(store staticStoreType, planID string) map[string]interface{} {
	return store[planID]
}
