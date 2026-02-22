package main

import (
	"httpfromtcp/internal/request"
	"fmt"
	"log"
	"net"
)



func main() {
	
    listener,err :=net.Listen("tcp", ":42069")          //openning a tcp port taht just listen 
	if err !=nil {                                     //handle error
		log.Fatal("ERROR ","ERROR",err)
	}
	for {
		conn,err :=listener.Accept()                  //coonection get accepted for listinng 

		fmt.Printf("\t There is a New HTTP message for you  : \r\n")
	    fmt.Printf("==============================================\r\n")
	
		if err!=nil {
			log.Fatal("ERROR ","ERROR",err)
		}

		r,err:=request.RequestFromReader(conn)
		if err!=nil{
			log.Fatal("ERROR ","ERROR",err)
		}

		fmt.Printf("Request line:\n")
		fmt.Printf("- Method: %s\n",r.RequestLine.Method)		
		fmt.Printf("- Target: %s\n",r.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n",r.RequestLine.HttpVersion)
		fmt.Printf("Headers: \n")
		r.Headers.ForEach(func (n,v string)  {
			fmt.Printf("- %s: %s\n",n,v)
		})
		fmt.Printf("==============================================\r\n")

		
		
	
	}



}