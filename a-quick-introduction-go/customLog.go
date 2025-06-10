package main 

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main(){
	LOGFILE := path.Join(os.TempDir(),"mGo.log")
	fmt.Println(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "ilog",log.LstdFlags)
	iLog.Println("Hello There !")
	iLog.Println("Mastering go 4th edition !")
}