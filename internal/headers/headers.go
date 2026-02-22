package headers

import (
	"bytes"
	"fmt"
	"strings"
)

var rn=[]byte("\r\n")

func isValidHeaderChar(b byte) bool {
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
	parts:=bytes.SplitN(fieldline,[]byte(":"),2)
	

	if len(parts)!=2 {
		return "","",fmt.Errorf("malformed fieldline ")
	}

	name:=parts[0]
	value:=bytes.TrimSpace(parts[1])

	for _, b := range name  {
		if  !isValidHeaderChar(b) {
			return "","",fmt.Errorf("invalid character in header name")	
		}
	}


	if bytes.HasSuffix(name,[]byte(" ")) {
		return "","",fmt.Errorf("malformed field name ")
	}

	lowerName:=strings.ToLower(string(name))

	return lowerName,string(value),nil
}


type Headers struct{
	headers map[string]string 
}

func NewHeaders() *Headers{
	return &Headers{
		headers: map[string]string{},
	}

}

func (h *Headers) Get(name string ) string {
	return  h.headers[strings.ToLower((name))]
}

func (h *Headers) Set(name,value string )  {
	name=strings.ToLower(name)
	if v,ok:=h.headers[name];ok {
		h.headers[name]=fmt.Sprintf("%s,%s",v,value)
	}else{
		h.headers[name]=value	
	}
}

func (h *Headers) ForEach(cb func(n,v string)) {
	for n,v :=range h.headers{
		cb(n,v)
		
	}
	
}

func (h *Headers) Parse(data []byte) ( int, bool,  error) {
    
	read :=0
	done:=false
	for {
		idx :=bytes.Index(data[read:],rn)
		if idx==-1 {
			break
		}
		//empty header
		if  idx==0 {
			done=true
			read+=len(rn)
			break
		}

		name,value,err:=ParseHeader(data[read:read+idx])
		if err!=nil {
			return 0,false,err
		}
		read+=idx +len(rn)
		h.Set(name,value)
		
		
	}
	return read,done,nil
}

