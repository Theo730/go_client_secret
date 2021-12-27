package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/julienschmidt/httprouter"
)

var Esia esia

func main() {

    var strInterface string
    err 	:= Init(&Esia);

    if err != nil{
	fmt.Printf("%v", err)
	os.Exit(-1)
    }

    if len(os.Args) < 4 {
	fmt.Printf("\n use: %s -h for help \n", os.Args[0])
	os.Exit(0)
    }

    router := httprouter.New()
    router.GET("/api/v1/getClientSecret", HGetClientSecret)

    if len(Esia.Interface) > 1{
	strInterface	= fmt.Sprintf("%s:%d", Esia.Interface, Esia.Port)
    }else{
	strInterface	= fmt.Sprintf(":%d", Esia.Port)
    }
    fmt.Printf("Listening HTTP\n")
    http.ListenAndServe(strInterface, router)
    os.Exit(0)
}