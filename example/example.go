package main

import (
	"os"

	logger "github.com/victormagalhaess/log2u"
)

func main() {
	log, err := logger.New(true, true, false, true, os.Stdout, "2006-01-02 15:04:05")
	if err != nil {
		panic("Cannot initialize logger")
	}

	log.Info("This is an info message")
	log.Success("This is a success message")
	log.Warning("This is a warning message")
	log.Error("This is an error message")
	log.Critical("This is a critical message")
	log.Debug("This is a debug message")
	log.CustomAnsiPrint("This is a custom message print in dark purple", 89)

	log.Infof("This is an info message %s", "with a string")
	log.Successf("This is a success message %s", "with a string")
	log.Warningf("This is a warning message %s", "with a string")
	log.Errorf("This is an error message %s", "with a string")
	log.Criticalf("This is a critical message %s", "with a string")
	log.Debugf("This is a debug message %s", "with a string")
	log.CustomAnsiPrintf("This is a custom message print in dark purple %s", 89, "with a string")

	log.SetShouldStack(true)
	log.SetTimeFormat("02/01/2006 15:04:05")

	log.Info("This is an info message with stack")
	log.Success("This is a success message with stack")
	log.Warning("This is a warning message with stack")
	log.Error("This is an error message with stack")
	log.Critical("This is a critical message with stack")
	log.Debug("This is a debug message with stack")

	log.Infof("This is an info message %s", "with a string and stack")
	log.Successf("This is a success message %s", "with a string and stack")
	log.Warningf("This is a warning message %s", "with a string and stack")
	log.Errorf("This is an error message %s", "with a string and stack")
	log.Criticalf("This is a critical message %s", "with a string and stack")
	log.Debugf("This is a debug message %s", "with a string and stack")

	log.SetRichOutput(false)

	log.CustomAnsiPrint("This is a custom message print in itallic", logger.Italic)
	log.CustomAnsiPrint("This is a custom message print in faint", logger.Faint)
	log.CustomAnsiPrint("This is a custom message print in bold", logger.Bold)
	log.CustomAnsiPrint("This is a custom message print in reverse", logger.ReverseVideo)
	log.CustomAnsiPrint("This is a custom message print in underline", logger.Underline)
	log.CustomAnsiPrint("This is a custom message print in crossed out", logger.CrossedOut)

}
