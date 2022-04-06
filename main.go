package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

const Character = "#"

var TerminalWidth int

// CheckError prints a message and the error if any
func CheckError(message string, err error) {
	if err != nil {
		println("ERROR " + message + ": " + err.Error())
		os.Exit(1)
	}
}

// Center prints a centered message in terminal
func Center(message string) {
	fmt.Println(strings.Repeat(Character, TerminalWidth))
	text := 2 + len(message)
	before := (TerminalWidth - text) / 2
	after := TerminalWidth - text - before
	fmt.Println(Character + strings.Repeat(" ", before) + message + strings.Repeat(" ", after) + Character)
	fmt.Println(strings.Repeat(Character, TerminalWidth))
}

// Title prints a title on terminal
func Title(message string) {
	after := TerminalWidth - 3 - len(message)
	fmt.Println(Character + " " + message + " " + strings.Repeat(Character, after))
}

// ResponsePrinter prints response
func ResponsePrinter(response *http.Response) error {
	Title("RESPONSE")
	httputil.DumpResponse(response, true)
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return err
	}
	fmt.Println(strings.TrimSpace(string(dump)))
	return nil
}

// RequestHandler returns a handler for proxy
func RequestHandler(address string) func(response http.ResponseWriter, request *http.Request) {
	url, err := url.Parse(address)
	CheckError("invalid redirection address", err)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ModifyResponse = ResponsePrinter
	return func(response http.ResponseWriter, request *http.Request) {
		Center("REQUEST from " + request.Host + " at " + time.Now().UTC().Format("2006-01-02T15:04:05"))
		Title("REQUEST")
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(response, "can't read body", http.StatusBadRequest)
			return
		}
		fmt.Println(strings.TrimSpace(string(dump)))
		proxy.ServeHTTP(response, request)
	}
}

func main() {
	port := flag.Int("port", 8000, "port proxy is listening")
	addr := flag.String("addr", "http://127.0.0.1:8080", "redirection address")
	flag.Parse()
	var err error
	TerminalWidth, _, err = terminal.GetSize(0)
	CheckError("getting terminal width", err)
	http.HandleFunc("/", RequestHandler(*addr))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
