package utilities

import (
	"os/exec"
)

func CreateWindow(pathName string, u string, s string, fs bool, a []string) {
	Url := "--app=" + u
	Size := "--window-size=" + s
	Args := a
	if fs {
		fsc := "--start-fullscreen"
		array := append(Args, Size, Url, fsc)
		cmdinstance := exec.Command(pathName, array...)
		cmdinstance.Run()
	} else {
		array := append(Args, Size, Url)
		cmdinstance := exec.Command(pathName, array...)
		cmdinstance.Run()
	}
}
