package main

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func HGetClientSecret(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    
    out, err	:= getClientSecret(Esia)
    if err != nil {
	HOutHttpError(w, -2, err.Error(), r.RemoteAddr)
	return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    str, err		:= json.Marshal(out)
    if err != nil {
	HOutHttpError(w, -3, err.Error(), r.RemoteAddr)
    }

    HOutHttp(w, "", "", string(str))
}

