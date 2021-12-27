package main

import (
	"fmt"
	"net/http"
)

func HOutHttp(w http.ResponseWriter, token string, oid string, json string){
    
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json		= fmt.Sprintf("{\"code\":0,\"text\":\"\",\"error\":\"\",\"data\":%s}", json)
    fmt.Fprint(w, json)
    return
}

func HOutHttpError(w http.ResponseWriter, num int, str string, ident string){

    w.WriteHeader(500)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json		:= fmt.Sprintf("{\"code\":%d,\"text\":\"%s\",\"error\":\"\",\"data\":\"\"}", num, str)
    fmt.Fprint(w, json)

    return
}
