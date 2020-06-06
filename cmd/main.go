//package main
package main

import (
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"log"
	"os"
	"strconv"
	"swagger_test/restapi"
	"swagger_test/restapi/operations"
	"swagger_test/service"
)

func main(){

	swaggerSpec, err := loads.Analyzed(restapi.FlatSwaggerJSON,"2.0")

	api := operations.NewSwaggerAPI(swaggerSpec)

	if err != nil{
		log.Panicln("Unable to analyze swaggerSpec", err)
	}
	svc := service.NewService()
	svc.Configure(api)

	srv:= restapi.NewServer(api)

	srv.EnabledListeners = []string{"http"}
	toint:= func(str string) int{
		i, _:= strconv.ParseInt(str, 10, 0)
		return int(i)
	}


	srv.Port = toint(os.Getenv("PORT"))
	fmt.Println(srv.Port)
	//or get port as flag
	var portFlag = flag.Int("port", 3000, "Port to run this service on")
	srv.Port = *portFlag
	defer func() {
		log.Println("shutting down...")
		srv.Shutdown()
	}()

	log.Println("starting server...")

	if err := srv.Serve(); err != nil {
		log.Panicln("server err:", err)
	}

}
