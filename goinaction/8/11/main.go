package main 

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace     *log.Logger
	Info      *log.Logger
	Warning   *log.Logger
	Error     *log.Logger
)

func init() {
    file, err  := os.OpenFile("error.txt",
		os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("大开文件失败", err)
	}

	Trace = log.New(ioutil.Discard,
		"trace:",
		log.Ldate|log.Ltime|log.Lshortfile)
	
	Info = log.New(os.Stdout,
		"info:",
		log.Ldate|log.Ltime|log.Lshortfile)
		
	Warning = log.New(os.Stdout,
		"warning:",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"Error",
	    log.Ldate|log.Ltime|log.Lshortfile)
}

func main () {
	Trace.Println("trace:i have some")
	Info.Println("info")
	Warning.Println("wn")
	Error.Println("error")
}