package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func buildUi() {
	cmd := exec.Command("npm", "--prefix", "myapp", "run", "build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkErrors(err)
}

func clearUi() {
	file, err := ioutil.ReadFile("./www/index.html")
	checkErrors(err)
	clearBundle := strings.Replace(string(file), "/bundle.js", "bundle.js", -1)
	clearStyles := strings.Replace(clearBundle, "/style.css", "style.css", -1)
	f, err := os.Create("./www/index.html")
	checkErrors(err)
	f.Write([]byte(clearStyles))
	defer f.Close()
}

func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	buildUi()
	clearUi()
	// time.Sleep(time.Second)
}
