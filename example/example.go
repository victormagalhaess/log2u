package main

import (
	"os"

	logger "github.com/victormagalhaess/log2u"
)

func main() {
	log, err := logger.New(true, true, false, os.Stdout, "2006-01-02 15:04:05")
	if err != nil {
		panic("Cannot initialize logger")
	}
	log.Info("This is an info message")
	log.Success("This is a success message")
	log.Warning("This is a warning message")
	log.Error("This is an error message")
	log.Critical("This is a critical message")
	log.Debug("This is a debug message")

	log.Infof("This is an info message %s", "with a string")
	log.Successf("This is a success message %s", "with a string")
	log.Warningf("This is a warning message %s", "with a string")
	log.Errorf("This is an error message %s", "with a string")
	log.Criticalf("This is a critical message %s", "with a string")
	log.Debugf("This is a debug message %s", "with a string")

	logStack, err := logger.New(true, true, true, os.Stdout, "2006-01-02 15:04:05")
	if err != nil {
		panic("Cannot initialize logger")
	}
	logStack.Info("This is an info message with stack")
	logStack.Success("This is a success message with stack")
	logStack.Warning("This is a warning message with stack")
	logStack.Error("This is an error message with stack")
	logStack.Critical("This is a critical message with stack")
	logStack.Debug("This is a debug message with stack")

	logStack.Infof("This is an info message %s", "with a string and stack")
	logStack.Successf("This is a success message %s", "with a string and stack")
	logStack.Warningf("This is a warning message %s", "with a string and stack")
	logStack.Errorf("This is an error message %s", "with a string and stack")
	logStack.Criticalf("This is a critical message %s", "with a string and stack")
	logStack.Debugf("This is a debug message %s", "with a string and stack")
}
