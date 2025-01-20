package logger

import (
	"os"
	"log"
	"fmt"
	"time"
)

func Log(szLogFileName string, szLogString string) {
	szLogFile, err := os.OpenFile(szLogFileName, os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer szLogFile.Close()
	szLogFile.WriteString(szLogString + "\r\n")
}

func Logger(szLogString string) {
	szTodate := time.Now().Format("20060102")
	szTodateDir := szTodate + "/"
	szRootDir := "Log"
	szRootDirPath := szRootDir + "/" + szTodateDir
	szFileName := "Log_" + szTodate + ".txt"
	szLogFilePath := szRootDirPath + szFileName
	
	// check log dir exist
	if _, err := os.Stat(szRootDirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(szRootDirPath, os.ModePerm); err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	// check log file exist
	if _, err := os.Stat(szLogFilePath); err != nil {
		szFile, _ := os.Create(szLogFilePath)
		defer szFile.Close()
	}

	Log(szLogFilePath, szLogString) 
}

func DEBUG(szLogString string) {
	szNow := time.Now().Format(time.DateTime)
	szLog := szNow + " | DEBUG | " + szLogString
	fmt.Println(szLog)
	Logger(szLog)
}

func INFO(szLogString string) {
	szNow := time.Now().Format(time.DateTime)
	szLog := szNow + " | INFO | " + szLogString
	fmt.Println(szLog)
	Logger(szLog)
}

func WARN(szLogString string) {
	szNow := time.Now().Format(time.DateTime)
	szLog := szNow + " | WARN | " + szLogString
	fmt.Println(szLog)
	Logger(szLog)
}

func ERROR(szLogString string) {
	szNow := time.Now().Format(time.DateTime)
	szLog := szNow + " | ERROR | " + szLogString
	fmt.Println(szLog)
	Logger(szLog)
}