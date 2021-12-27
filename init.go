package main

import (
    "fmt"
    "crypto/x509"
    "flag"
    "encoding/pem"
    "errors"
    "io/ioutil"
    "strings"
    "github.com/Theo730/pkcs7"
)

func Init(esia *esia)(err error){

    var keyFile, certFile string
    flag.StringVar(&esia.ClientID, "ClientID", "ISP03", " - \"Mnenomics\" in terms of ESIA ")
    flag.StringVar(&esia.Scopes, "Scopes", "fullname", " - Scopes ")
    flag.StringVar(&esia.Interface, "Interface", "", " - interface, default all")
    flag.IntVar(&esia.Port, "Port", 5959, " - port, default 5959")
    flag.StringVar(&keyFile, "Key", "keys/key.pem", " - private key file, default keys/key.pem")
    flag.StringVar(&certFile, "Cert", "keys/cert.pem", " - certificate file, default keys/cert.pem")
    flag.Parse()
    
    pkey, err 		:= ioutil.ReadFile(keyFile)
    if err != nil {
	return errors.New(fmt.Sprintf("Error open key file %v %v", err, keyFile))
    }

    blockPkey, _ 	:= pem.Decode(pkey)
    if blockPkey == nil || blockPkey.Type != "PRIVATE KEY" {
	return errors.New("failed to decode PEM block containing private key ")
    }

    key, err 		:= pkcs7.ParsePKCS8PrivateKey(blockPkey.Bytes)
    if err != nil {
	return err
    }
    esia.Key		= key

    ckey, err 		:= ioutil.ReadFile(certFile)
    if err != nil {
	return errors.New(fmt.Sprintf("Error open cert file %v %v", err, certFile))
    }

    blockCkey, _ := pem.Decode(ckey)
    if blockCkey == nil || blockCkey.Type != "CERTIFICATE" {
	return errors.New("failed to decode PEM block containing certificate")
    }

    cert, err		:= x509.ParseCertificate(blockCkey.Bytes)
    if err != nil {
	return err
    }
    esia.Cert		= cert

    esia.Scopes		= strings.ReplaceAll(esia.Scopes, ",", "")
    
    return nil
}
