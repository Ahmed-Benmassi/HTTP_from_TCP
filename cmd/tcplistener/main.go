package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)


func getLinesChannel(f io.ReadCloser) <-chan string {        //func that take the message and read  it  8 byte and each end line with \n we split it and somme it to a line that is readabael
	lines := make(chan string)                               
	go func() {
		
		defer close(lines)
		currentLineContents := ""
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(b[:n]) 
			parts := strings.Split(str, "\n")                    //go func take the 8 byte readed message and take the last line of the message and complete it to tha first line 
			for i := 0; i < len(parts)-1; i++ {                  // the last part of the slice that is a beggining of the next line will be the first word of he next line 
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}                                                       //example  "hello worltd go on" ist goes like this [hello wo,rld go on] 
			currentLineContents += parts[len(parts)-1]
		}
	}()
	return lines
}



func main() {
	
    listener,err :=net.Listen("tcp", ":42069")          //openning a tcp port taht just listen 
	if err !=nil {                                     //handle error
		log.Fatal("ERROR ","ERROR",err)
	}
	
	
	for {
		conn,err :=listener.Accept()                  //coonection get accepted for listinng 
		if err!=nil {
			log.Fatal("ERROR ","ERROR",err)
		}

		for line := range getLinesChannel(conn) {      //printing the message like network like file
			fmt.Printf("read: %s\n",line)
		}
		
	
	}



}