package main

/*
#include <stdlib.h>
#include <string.h>
*/

import "C"

import (
	"encoding/json"
	"os"
	"os/exec"
	"time"
	utl "webapp/utilities"
)

func str(ch *C.char) string {
	return C.GoString(ch)
}

func ch(str string) *C.char {
	return C.CString(str)
}

type options struct {
	Url        string   `json:Url`
	Size       string   `json:Size`
	Args       []string `json:Args`
	FullScreen bool
}

//export CreateWindow
func CreateWindow(a *C.char) {

	go func() {
		args := str(a)
		e := []byte(args)
		var newOption options
		json.Unmarshal(e, &newOption)

		// fmt.Println(e)
		// fmt.Println(newOption.Url, newOption.Args)
		// fmt.Println(args)

		Url := newOption.Url
		Size := newOption.Size
		Args := newOption.Args
		FullScreen := newOption.FullScreen

		pathName := utl.LocateChrome()
		utl.CreateWindow(pathName, Url, Size, FullScreen, Args)

	}()
	StartDevServer(true)

}

func StartDevServer(start bool) {
	if start {
		cmd := exec.Command("npm", "--prefix", "myapp", "run", "dev")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		checkErrors(err)
	}
}

func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	time.Sleep(time.Second)
}
