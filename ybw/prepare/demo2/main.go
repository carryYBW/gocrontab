package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//windows 系统在下载 cygwin64
	cmd := exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 5; ls -l")
	err := cmd.Run()
	fmt.Println(err)

}
