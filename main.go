// Main package
package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

	"loco-switch/data"
	"loco-switch/handlers"
	"loco-switch/helpers"
	"loco-switch/models"
)

/*
TODO:

- route protection (with filter) => check how we can do it with gin
- wasm filters
- api for monitoring, health check etc...
- how to handle the logs?
*/

// TO BE USED to protect the routes
//var locoSwitchAdminToken = getEnv("LOCO_SWITCH_ADMIN_TOKEN", "")

// TODO: document (every thing)

// Logger function → check the request to dectect problems
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		c.Next()
		// after request

		// access the status we are sending
		status := c.Writer.Status()

		if status == 502 {
			var revision = "default"
			var functionName = c.Params[0].Value
			if len(c.Params) > 1 {
				revision = c.Params[1].Value
			}
			//log.Println(status, functionName, revision)

			var functionError models.FunctionError
			if data.GetFunction(functionName+":"+revision).Name != "" {
				// function is registered but does not respond
				functionError = models.FunctionError{
					Name:     functionName,
					Revision: revision,
					Error:    "connection refused",
				}
				// TODO: unregister the function after N trys
				// TODO: or restart the function
			} else {
				// function is not registered
				functionError = models.FunctionError{
					Name:     functionName,
					Revision: revision,
					Error:    "function not registered",
				}
			}

			c.JSON(status, &functionError)
		}

	}
}

func main() {

	data.LoadFunctions()

	go func() {
		for {
			data.SaveFunctions()
			time.Sleep(5 * time.Second)
		}
	}()

	httpServer := gin.Default()

	httpServer.Use(Logger())

	/*
		TODO:

			At start, check if we use wasm plugin
			if yes or no, select the appropriate handlers:
			- registerFunction or registerFunctionWithFilters
			- getFunctions or getFunctionsWithFilters
			- proxy or proxyWithFilters
			- etc ...

	*/

	locoSwitchEndPoint := helpers.GetEnv("LOCO_SWITCH_ENDPOINT", "functions")

	// 👋 we need to handle the GET too
	//Create catchall routes
	httpServer.Any("/"+locoSwitchEndPoint+"/:function_name", handlers.Proxy)
	// 🚧 work in progress
	httpServer.Any("/"+locoSwitchEndPoint+"/:function_name/:function_revision", handlers.ProxyRevision)

	// 🚧 work in progress
	httpServer.POST("/admin/"+locoSwitchEndPoint+"/registration", handlers.RegisterFunction)
	httpServer.DELETE("/admin/"+locoSwitchEndPoint+"/registration", handlers.UnRegisterFunction)

	httpServer.GET("/admin/"+locoSwitchEndPoint+"/list", handlers.GetFunctions)

	httpServer.POST("/admin/"+locoSwitchEndPoint+"/endpoint", handlers.AddEndpointToFunction)
	httpServer.DELETE("/admin/"+locoSwitchEndPoint+"/endpoint", handlers.RemoveEndpointFromFunction)

	if helpers.GetEnv("LOCO_SWITCH_CRT", "") != "" {

		httpServer.RunTLS(
			":"+helpers.GetEnv("LOCO_SWITCH_HTTPS_PORT", "4443"),
			helpers.GetEnv("LOCO_SWITCH_CRT", "certs/loco-switch.local.crt"),
			helpers.GetEnv("LOCO_SWITCH_KEY", "certs/loco-switch.local.key"),
		)
	} else {

		// Ngrok support: https://ngrok.com
		// https://ngrok.com/blog-post/ngrok-go
		if ngrok.WithAuthtokenFromEnv() != nil {
			tun, err := ngrok.Listen(context.Background(),
				config.HTTPEndpoint(),
				ngrok.WithAuthtokenFromEnv(),
			)
			if err != nil {
				log.Println("❌ Error while creating tunnel:", err)
			}

			log.Println("👋 Ngrok tunnel created:", tun.URL())

			ex, err := os.Executable()
			if err != nil {
				log.Fatal("❌ Error after creating tunnel:", err)
			}
			exPath := filepath.Dir(ex)

			f, err := os.Create(exPath + "/ngrok.url")

			if err != nil {
				log.Fatal("❌ Error when creating ngrok.url:", err)

			}

			defer f.Close()

			_, errWrite := f.WriteString(tun.URL())

			if errWrite != nil {
				log.Fatal("❌ Error when writing ngrok.url:", errWrite)
			}

			log.Println("🤚 Ngrok URL:", exPath+"/ngrok.url")

			httpServer.RunListener(tun)
		}

		httpServer.Run(":" + helpers.GetEnv("LOCO_SWITCH_HTTP_PORT", "8080"))

	}

}
