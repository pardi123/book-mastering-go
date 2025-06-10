package main 

import (
	"log"
	"log/syslog"
)
func main(){
	sysLog, err := syslog.New(syslog.LOG_SYSLOG,"systemLog.go")
	if err != nil {
		log.Println(err)
	} else {
		log.SetOutput(sysLog)
		log.Print("everything is Fine")

	}
}