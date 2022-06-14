package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fatih/color"
)

func main(){
	color.Green("Start Server")
	setup()
}


func SetUpServer(){
	server := &http.Server{
		Addr:              ":80",
		Handler:           http.HandlerFunc(RequestHandler),
	}
	go StartServer(server)
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<- quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	color.Yellow("[Server]Server is stop")
}

func StartServer(server *http.Server){
	err := server.ListenAndServe()
	if err != nil{
		color.Red("HttpServerStartError")
	}
}

func RequestHandler(w http.ResponseWriter, Request *http.Request){
	var RequestURLpath string = Request.URL.Path

	if RequestURLpath == "/"{
		return	
	}
	Request.ParseMultipartForm(100)

	files := Request.Form

	color.Blue("kita")

	for _,f := range files {
		fmt.Println("Keyが1個ありました!")
		fmt.Println(f)
	}

}
