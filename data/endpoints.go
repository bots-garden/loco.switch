package data

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
)

func randomChoice(min, max int) int {
	return min + rand.Intn(max-min)
}
/*
https://golang.cafe/blog/golang-random-number-generator.html
*/

func GetEndpointUrl(key_function string) (string, error) {
	// https://www.jscape.com/blog/load-balancing-algorithms
	// random load balancing
	functionEndpoints := GetFunctionEndpoints(key_function)
	if functionEndpoints == nil {
		log.Println("🔴 data.GetEndpointUrl", key_function, "no endpoint or no function")
		return "", errors.New("no endpoint or no function")
	} else {
		endpointNumber := randomChoice(0, len(functionEndpoints)) // a function can have several endpoints
		selectedEndpoint := GetFunctionEndpoints(key_function)[endpointNumber]
		
		var protocol string
		if selectedEndpoint.Https {
			protocol = "https://"
		} else {
			protocol = "http://"
		}
	
		functionUrl := protocol + selectedEndpoint.Domain + ":" +
			strconv.Itoa(selectedEndpoint.HttpPort) + selectedEndpoint.Resource
	
		return functionUrl, nil
	}

}
