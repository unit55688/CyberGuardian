package common

import (
	"bufio"
	"os"
	"strings"

	"CyberGuardian/logger"

	"gopkg.in/ini.v1"
)

func GetHosts(szHostsFile string) []string {
	if _, err := os.Stat(szHostsFile); err != nil {
		szFile, _ := os.Create(szHostsFile)
		defer szFile.Close()
		logger.Log(szHostsFile, "192.168.1.1")
		logger.Log(szHostsFile, "192.168.1.2")
	}

	szHosts := []string{}
	file, _ := os.Open(szHostsFile)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	// fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		szLines := fileScanner.Text()
		szHosts = append(szHosts, strings.TrimSpace(szLines))

	}
	return szHosts
}

func GetMailConfig(szConfigFile string) map[string]string {
	if _, err := os.Stat(szConfigFile); err != nil {
		szFile, _ := os.Create(szConfigFile)
		defer szFile.Close()
		logger.Log(szConfigFile, "[SMTP]")
		logger.Log(szConfigFile, "server = smtp.gmail.com")
		logger.Log(szConfigFile, "port = 587")
		logger.Log(szConfigFile, "sender = ")
		logger.Log(szConfigFile, "password = ")
		logger.Log(szConfigFile, "receiver = ")
		logger.Log(szConfigFile, "")
		logger.Log(szConfigFile, "[Mail]")
		logger.Log(szConfigFile, "subject = \"設備斷線告警\"")
		logger.Log(szConfigFile, "body = \"設備無回應，請檢查網絡連接！!\"")
	}

	szResult := map[string]string{}
	cfg, err := ini.Load(szConfigFile)
	if err != nil {
		logger.ERROR("無法讀取 config.cfg: " + err.Error())
	}

	szServer := cfg.Section("SMTP").Key("server").String()
	szPort := cfg.Section("SMTP").Key("port").String()
	szSender := cfg.Section("SMTP").Key("sender").String()
	szPassword := cfg.Section("SMTP").Key("password").String()
	szReceiver := cfg.Section("SMTP").Key("receiver").String()

	szSubject := cfg.Section("Mail").Key("subject").String()
	szBody := cfg.Section("Mail").Key("body").String()

	szResult["server"] = szServer
	szResult["port"] = szPort
	szResult["sender"] = szSender
	szResult["password"] = szPassword
	szResult["receiver"] = szReceiver
	szResult["subject"] = szSubject
	szResult["body"] = szBody

	return szResult
}
