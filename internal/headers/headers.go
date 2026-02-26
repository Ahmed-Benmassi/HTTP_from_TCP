package headers

import (
	"bytes"
	"fmt"
	"strings"
)

var rn=[]byte("\r\n")

func isValidHeaderChar(b byte) bool {      //allowing just those type of characters
    // Letters
    if b >= 'a' && b <= 'z' { return true }
    if b >= 'A' && b <= 'Z' { return true }

    // Digits
    if b >= '0' && b <= '9' { return true }

    // Allowed special characters
    switch b {
    case '!', '#', '$', '%', '&', '\'',
         '*', '+', '-', '.', '^', '_',
         '`', '|', '~':
        return true
    }

    return false
}



func ParseHeader(fieldline []byte)	 (string,string,error) {
	parts:=bytes.SplitN(fieldline,[]byte(":"),2)               //split the fieldline into 2 :keys and value in ":"
	

	if len(parts)!=2 {
		return "","",fmt.Errorf("malformed fieldline ")   //if parts having more or less than 2 things are (keys ;values )retuern error
	}

	name:=parts[0]
	value:=bytes.TrimSpace(parts[1])                     //remove space for comparaision

	for _, b := range name  {
		if  !isValidHeaderChar(b) {
			return "","",fmt.Errorf("invalid character in header name")	      //if character not valid do error
		}
	}


	if bytes.HasSuffix(name,[]byte(" ")) {                   //if it have " " in the end instead of ":" return error
		return "","",fmt.Errorf("malformed field name ")
	}

	lowerName:=strings.ToLower(string(name))

	return lowerName,string(value),nil
}


type Headers struct{
	headers map[string]string                   //"Host": "localhost"
}

func NewHeaders() *Headers{                      //Creates a new Headers object
	return &Headers{
		headers: map[string]string{},
	}

}

func (h *Headers) Get(name string ) (string,bool) {        //get header keys and lowercase it 
	str,ok:=h.headers[strings.ToLower((name))]
	return str,ok  
}

func (h *Headers) Replace(name,value string )  {                //replace values to already valued names or empty values name 
	name=strings.ToLower(name)
	h.headers[name]=value	
	
}

func (h *Headers) Set(name, value string) {
	if h == nil {
		return
	}

	if h.headers == nil {
		h.headers = make(map[string]string)
	}

	name = strings.ToLower(name)
	h.headers[name] = value
}

func (h *Headers) Delete(name string )  {                //deleting keys
	name=strings.ToLower(name)
	delete(h.headers,name)	
	
}

func (h *Headers) ForEach(cb func(n,v string)) {              //doing a loop with foreach that calls afunc on every key and value to do something later 
	for n,v :=range h.headers{
		cb(n,v)
		
	}
	
}

func (h *Headers) Parse(data []byte) ( int, bool,  error) {           //parse HTTP headers from raw bytes.
    
	read :=0
	done:=false
	for {
		idx :=bytes.Index(data[read:],rn)                                  //Loops over each line separated by \r\n.
		if idx==-1 {
			break
		}                                                                                                                            
		if  idx==0 {                                                        // Stops if it finds the empty line (end of headers)
			done=true
			read+=len(rn)
			break
		}

		name,value,err:=ParseHeader(data[read:read+idx])                 // Uses ParseHeader to extract name and value.
		if err!=nil {
			return 0,false,err
		}
		read+=idx +len(rn)                                               // Stores each header using Set.
		h.Replace(name,value)
		
		
	}
	return read,done,nil                                                // Returns how many bytes were read, whether done, and any error empty header
}

