package main

import (
	x509 "github.com/fanxiong/gm-chain/mitx509"
	"path/filepath"
	"os"
	"log"
	"github.com/fanxiong/gm-chain/node"
	"fmt"
)

func main() {
	//dir info
	path:= node.DefaultDataDir()
	if path != "" {
		path=filepath.Join(path,x509.CertPath)
	}
	fmt.Println("path--->",path)
	certPath:=path
	flag,error:=x509.PathExists(certPath)
	if error!=nil{
		log.Panic("error find exist")
	}
	if !flag{
		//there is no filepath create it
		err:=os.MkdirAll(certPath,os.ModePerm)
		if err!=nil{
			log.Panic("create dir fail")
		}
	}
	//root info
	rootInfo := x509.CertInfo{
		Country: []string{"SGP"},
		Organization: []string{"MiaoA Fundation"},
		IsCA: true,
		OrganizationalUnit: []string{"M.I.T"},
		EmailAddress: []string{"contact@tnb.fund"},
		Locality: []string{"SGP"},
		Province: []string{"SGP"},
		CommonName: "M.I.T Team",
		CrtName: filepath.Join(certPath,x509.RootCrtName),
		KeyName: filepath.Join(certPath,x509.RootKeyName)}

	//exist if the root cert
	crtBool:=x509.FileExists(rootInfo.CrtName)
	keyBool:=x509.FileExists(rootInfo.KeyName)



	if(!crtBool&&!keyBool){
		error=x509.CreateRootCERT(rootInfo)
		if error!=nil{
			log.Println("create root cert success")
		}else{
			log.Println("crate cert success")
		}
	}else{
		log.Println("there is exist cert or key")
	}




}