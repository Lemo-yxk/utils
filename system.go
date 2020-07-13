/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-12-26 19:29
**/

package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

type system int

const System system = iota

var ch = make(chan int)

func (system system) OpenBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func (system system) Exit(code int) {
	ch <- code
}

func (system system) Block() int {
	return <-ch
}
