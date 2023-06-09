package main

import (
	"bytes"
	"fmt"
	"github.com/xela07ax/toolsXela/tp"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	originFilePath = "/etc/resolv.conf"
	//originFilePath     = "/home/xela/Projects/back-encore-1/AutostarterResolver/dat/resolv.conf"
	fileTemplateResolv = "template.resolv"
	replaceString      = "{DATA}"
	logFileName        = "logix.log"
)

type MyLogger struct {
	logDir string
}

func (m *MyLogger) WriteLog(text string, errorType bool) {
	file, err := os.OpenFile(m.logDir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	tp.Fck(err)
	typePrefix := "INFO"
	if errorType {
		typePrefix = "ERROR"
	}
	_, err = file.Write([]byte(fmt.Sprintf("\n[%s]%s| %s", time.Now().Format("2006.01.02 15:04:05"), typePrefix, text)))
	tp.Fck(err)
}

func main() {
	fmt.Println("Hola! AutostarterResolver")
	dir, err := tp.BinDir()
	tp.Fck(err)
	logDir := filepath.Join(dir, logFileName)
	logger := &MyLogger{logDir: logDir}
	logger.WriteLog("Hola", false)
	logger.WriteLog("Starting... Step 1 resolv", false)
	// Step 1 resolv
	fTemplateData, err := ioutil.ReadFile(filepath.Join(dir, fileTemplateResolv))
	if err != nil {
		logger.WriteLog(err.Error(), true)
		tp.Fck(err)
	}

	fResolvData, err := ioutil.ReadFile(originFilePath)
	if err != nil {
		logger.WriteLog(err.Error(), true)
		tp.Fck(err)
	}

	finalResolvFile := strings.Replace(string(fTemplateData), replaceString, string(fResolvData), -1)

	fileResolv, err := os.Create(originFilePath)
	if err != nil {
		logger.WriteLog(err.Error(), true)
		tp.Fck(err)
	}
	_, err = fileResolv.Write([]byte(finalResolvFile))
	if err != nil {
		logger.WriteLog(err.Error(), true)
		tp.Fck(err)
	}
	// Step 2 nvidia power limit _ nvidia-smi -i 0 -pl 160
	logger.WriteLog("Starting... Step 2 nvidia power limit", false)

	cmd := exec.Command("nvidia-smi", "-i", "0", "-pl", "160")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	logger.WriteLog(fmt.Sprintf("out:%s| bin/bash:%s", outb.String(), errb.String()), false)
	if err != nil {
		logger.WriteLog(err.Error(), true)
		tp.Fck(err)
	}
	logger.WriteLog("Good by", false)
	fmt.Println("Good by! AutostarterResolver")
}
