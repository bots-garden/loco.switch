package data

import (
	"encoding/json"
	"io/ioutil"
	"loco-switch/helpers"
	"loco-switch/models"
	"log"
	"os"
)

//TODO: use concurrent map
var functionsMap = make(map[string] models.FunctionRecord)
var storage = helpers.GetEnv("LOCO_SWITCH_STORAGE", "functions.json")

func GetFunctionEndpoints(key_function string) []models.Endpoint {
	function := functionsMap[key_function]
	//log.Println("🎃 GetFunctionEndpoints", key_function, function)
	return function.Endpoints
}

func SetFunction(key_function string, function models.FunctionRecord) {
	functionsMap[key_function] = function
}

func GetFunction(key_function string) models.FunctionRecord {
	return functionsMap[key_function]
}

func RemoveFunction(key_function string) {
	delete(functionsMap, key_function)
}

func GetFunctions() map[string] models.FunctionRecord {
	return functionsMap
}

// TODO: handle errors correctly
func SaveFunctions() {
	// save map to json file
	jsonString, _ := json.MarshalIndent(functionsMap, "", "  ")
	_ = ioutil.WriteFile(storage, jsonString, 0644)
}

// TODO: handle errors correctly
func LoadFunctions() {
	// load map from json file
	jsonFile, err := os.Open(storage)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened", storage)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &functionsMap)

}

 
