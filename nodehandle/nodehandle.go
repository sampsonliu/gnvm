package nodehandle

import (

	// go
	"fmt"
	//"os"
	//"os/exec"
)

const (
	DIVIDE   = "\\"
	NODEHOME = "NODE_HOME_2"
	PATH     = "path"
	NODE     = "node.exe"
)

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Get current path Error: " + err.Error())
		return ""
	}
	return path
}

func Use(folderName string, global bool) {

	// get path
	path := getCurrentPath() + DIVIDE + folderName
	fmt.Println("Use node version path is " + path)

	// get node path
	nodePath := path + DIVIDE
	fmt.Println("Node.exe path is " + nodePath)

	// set env
	/*
		if err := os.Setenv(NODEHOME, nodePath); err != nil {
			fmt.Println("Set Env Error: " + err.Error())
			return
		}
		fmt.Println("NODE_HOME is " + os.Getenv(NODEHOME))

		// set path
		if err := os.Setenv("Path", os.Getenv(NODEHOME)); err != nil {
			fmt.Println("Set Env Error: " + err.Error())
			return
		}
		fmt.Println(os.Getenv("Path	"))
	*/

	/*
		cmd := exec.Command("set", NODEHOME, nodePath)
		if err := cmd.Run(); err != nil {
			fmt.Println("Exec set NODE_HOME Error: " + err.Error())
		}
		fmt.Println("NODE_HOME is " + os.Getenv(NODEHOME))
	*/

	/*
		if err := syscall.Setenv(NODEHOME, nodePath); err != nil {
			fmt.Println("Set Env Error: " + err.Error())
			return
		}
		value, _ := syscall.Getenv(NODEHOME)
		fmt.Println("NODE_HOME is " + value)
	*/

	/*
		output, err := exec.Command("gnvm2.bat").CombinedOutput()
		if err != nil {
			fmt.Println("sadfsafaf " + err.Error())
			return
		}
		fmt.Println(string(output))
	*/

	/*
		os.Setenv("FOO", "BAR")
		if err := syscall.Exec(os.Getenv("cmd"), []string{os.Getenv("cmd")}, syscall.Environ()); err != nil {
			fmt.Println("sdfdfafaf " + err.Error())
		}
		//fmt.Println(syscall.Environ())
	*/

	/*
		cmd := exec.Command("cmd", "/C", "del", "D:\\a.txt")
		if err := cmd.Run(); err != nil {
			fmt.Println("sadfsafaf " + err.Error())
		}
	*/
}
