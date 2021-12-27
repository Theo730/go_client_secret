package main


func getClientSecret(esia esia) (rec OutURL, err error){

    rec.State	= getState()
    strTime	:= getData()
    message	:= getMessage(esia, strTime, rec.State)
    
    rec.ClientSecret, err	= esia.SignMessage(message)
    if err != nil {
        return rec, err
    }

    return rec, nil
}

