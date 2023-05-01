package handlers

import (
	"loco-switch/data"
	"loco-switch/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
curl -X DELETE http://localhost:8080/admin/functions/endpoint \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"orange"
}
EOF
*/
// UnRegisterFunction is a handler for /admin/functions/registration (DELETE)
func UnRegisterFunction(c *gin.Context) {

	function := models.Function{}

	if err := c.BindJSON(&function); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	functionRecord := models.FunctionRecord{
		Name:      function.Name,
		Revision:  function.Revision,
	}

	data.RemoveFunction(function.Name+":"+function.Revision)

	c.JSON(http.StatusAccepted, &functionRecord)

}

/*
curl -X DELETE http://localhost:8080/admin/functions/endpoint \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"scale",
    "httpPort":3333,
	"status":0,
    "https":false,
    "domain":"localhost",
	"resource":""
}
EOF
*/
// RemoveEndpointFromFunction is a handler for /admin/functions/endpoint (DELETE)
//   - Remove a endpoint from a function
//   - It will be used to downscale an existing function.
func RemoveEndpointFromFunction(c *gin.Context) {
	function := models.Function{}

	if err := c.BindJSON(&function); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	endpoint := models.Endpoint{
		HttpPort: function.HttpPort,
		Https:    function.Https,
		Domain:   function.Domain,
		Resource: function.Resource,
	}

	currentEndpointsFunction := data.GetFunctionEndpoints(function.Name+":"+function.Revision)

	// create a new functionRecord
	functionRecord := models.FunctionRecord{
		Name:      function.Name,
		Revision:  function.Revision,
		Status:    function.Status,
		Timestamp: time.Now(),
		Endpoints: []models.Endpoint{},
	}
	// remove the endpoint
	for _, item := range currentEndpointsFunction {
		if !(item.HttpPort == endpoint.HttpPort && item.Domain == endpoint.Domain) {
			functionRecord.Endpoints = append(functionRecord.Endpoints, item)
		} 
    }

	// rewrite (replace) the function record 
	data.SetFunction(function.Name+":"+function.Revision, functionRecord)

	c.JSON(http.StatusAccepted, &functionRecord)
}

/*
curl -X POST http://localhost:8080/admin/functions/endpoint \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"orange",
    "httpPort":2222,
    "status":0,
    "https":false,
    "domain":"localhost",
	"resource":""
}
EOF
*/

// AddEndpointToFunction is a handler for /admin/functions/endpoint (POST)
// 	- It will be used to scale an existing function
// 	- We call the same function, but it will reach a different process
func AddEndpointToFunction(c *gin.Context) {
	function := models.Function{}

	if err := c.BindJSON(&function); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	endpoint := models.Endpoint{
		HttpPort: function.HttpPort,
		Https:    function.Https,
		Domain:   function.Domain,
		Resource: function.Resource,
	}

	currentEndpointsFunction := data.GetFunctionEndpoints(function.Name+":"+function.Revision)

	// create a new functionRecord
	functionRecord := models.FunctionRecord{
		Name:      function.Name,
		Revision:  function.Revision,
		Status:    function.Status,
		Timestamp: time.Now(),
		Endpoints: currentEndpointsFunction,
	}

	functionRecord.Endpoints = append(functionRecord.Endpoints, endpoint)

	// rewrite (replace) the function record 
	data.SetFunction(function.Name+":"+function.Revision, functionRecord)

	c.JSON(http.StatusAccepted, &functionRecord)
}

/*
curl http://localhost:8080/admin/functions/list 
*/
// GetFunctions is a handler for /admin/functions/list
//  - It returns a list of all functions
func GetFunctions(c *gin.Context) {
	functionsList := data.GetFunctions()
	c.JSON(http.StatusAccepted, &functionsList)
}

/*
curl -X POST http://localhost:8080/admin/functions/registration \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"hello_js",
    "revision":"orange",
    "httpPort":4444,
    "status":0,
    "https":false,
    "domain":"localhost",
	"resource":""
}
EOF
*/
// RegisterFunction is a handler for /admin/functions/registration (POST)
// - It creates a function record in the functions map
func RegisterFunction(c *gin.Context) {
	function := models.Function{}

	if err := c.BindJSON(&function); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	endpoint := models.Endpoint{
		HttpPort: function.HttpPort,
		Https:    function.Https,
		Domain:   function.Domain,
		Resource: function.Resource,
	}

	functionRecord := models.FunctionRecord{
		Name:      function.Name,
		Revision:  function.Revision,
		Status:    function.Status,
		Timestamp: time.Now(),
		Endpoints: []models.Endpoint{endpoint},
	}

	data.SetFunction(function.Name+":"+function.Revision, functionRecord)

	c.JSON(http.StatusAccepted, &functionRecord)

}