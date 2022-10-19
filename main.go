package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX,hasSPF,sprREcord,hasDMARC, dmarcRecord\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err:= scanner.Err();err !=nil{
		log.Fatal("Error could not read from input :%v\n",err)
	}
}

func checkDomain(domain string){
	var hasMX,hasSPF,hasDMARC bool

	var spfRecord,dmarcRecord string

	mxRecords,err :=  net.LookupMX(domain)

	if err!=nil{
		log.Print("Error: %v\n",err)
	}

	if len(mxRecords)>0{
		hasMX = true
	}

	txtRecords,err := net.LookupTXT(domain)

	if err!=nil{
		log.Printf("Error: %v\n",err)
	}

	for _,record :=  range txtRecords{
		if strings.HasPrefix(record,"v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	net.loo

}