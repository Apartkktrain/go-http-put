package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
  loggerCh chan string
)

func ResponseHandle(w http.ResponseWriter, r *http.Request) {
    loggerCh <- r.Method
    r.ParseForm()

    switch r.Method {
    case "PUT":
      length := r.ContentLength
      body := make([]byte, length)
      r.Body.Read(body)
      fmt.Println(string(body))
    default:
      return
    }

}

func mainsub() {
  loggerCh = make(chan string)
  log.Println("Start")
    go func() {
      http.HandleFunc("/", ResponseHandle)
      err := http.ListenAndServe(":80", nil)
      if err != nil {
          log.Fatal("ListenAndServe: ", err)
        }
      }()
      for {
              select {
                case str := <-loggerCh:
                  log.Printf("get request: = %s", str)
                }
      }

    }
