package static

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type staticStoreType map[string]map[string]interface{}

var staticStore staticStoreType

func LoadStaticStore(jsonFile string) {
	inBuf, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(inBuf, &staticStore)
	if err != nil {
		fmt.Println(staticStore)
		panic(err)

	}

	return
}

func GetCredsForPlanID(planID string) map[string]interface{} {
	return staticStore[planID]
}
