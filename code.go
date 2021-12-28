package main

import(
     "net/url"
)

func getClientSecret(esia esia) (rec OutURL, err error){

    rec.State	= getState()
    rec.Data	= getData()
    message	:= getMessage(esia, rec.Data, rec.State)
    
    rec.ClientSecret, err	= esia.SignMessage(message)
    if err != nil {
        return rec, err
    }
    rec.Data	= url.QueryEscape(rec.Data)
    return rec, nil
}

