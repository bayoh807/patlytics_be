package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type resourceType struct {
	data map[string]map[string]map[string]interface{}
}

var Resource = resourceType{
	data: make(map[string]map[string]map[string]interface{}), // Initialize the nested map
}

func init() {

	Resource.data["patents"], _ = Resource.readJson("patents", "publication_number")
	Resource.data["companies"], _ = Resource.readJson("company_products", "name")

}

func (r *resourceType) GetData(title string, key ...string) interface{} {
	if v, ok := Resource.data[title]; !ok {
		return nil
	} else if len(key) == 0 || key[0] == "" {
		return v
	} else if v2, ok2 := v[key[0]]; !ok2 {
		return nil
	} else {
		return v2
	}
}
func (r *resourceType) readJson(name, key string) (map[string]map[string]interface{}, error) {
	dir, _ := os.Getwd()
	jsonFile, err := os.Open(dir + "/data/" + name + ".json")
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var itmes []map[string]interface{}
		results := map[string]map[string]interface{}{}
		json.Unmarshal([]byte(byteValue), &itmes)

		for _, item := range itmes {
			results[item[key].(string)] = item
		}

		return results, nil
	}

}
