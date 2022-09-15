package main

import (
	"embed"
	"encoding/json"
	"net"
	"net/http"
	utl "webapp/utilities"
)

//go:embed www
var fs embed.FS

//go:embed build.json
var buildjs string

type options struct {
	Url        string   `json:Url`
	Size       string   `json:Size`
	Args       []string `json:Args`
	FullScreen bool
}

func readFile(file string, Url string) {

	e := []byte(file)
	var newOption options
	json.Unmarshal(e, &newOption)

	// Url := newOption.Url
	Size := newOption.Size
	Args := newOption.Args
	FullScreen := newOption.FullScreen

	pathName := utl.LocateChrome()
	utl.CreateWindow(pathName, Url, Size, FullScreen, Args)

}

func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	checkErrors(err)
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.FS(fs)))
	fileName := buildjs
	Url := "http://" + ln.Addr().String() + "/www"
	readFile(fileName, Url)
}
