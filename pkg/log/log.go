package log

import (
	"fmt"
	"os"
)

func CheckError(err error, message string, params ...interface{}) {
	if err != nil {
		LogError(err, message, params...)
		os.Exit(1)
	}
}

func LogError(err error, message string, params ...interface{}) {
	log("ERROR: while " + message + ": %w", append(params, err)...)
}

func LogInfo(message string, params ...interface{}) {
	log("INFO: " + message, params...)
}

func LogTrace(message string, params ...interface{}) {
	log("TRACE: " + message, params...)
}

func log(message string, params ...interface{}) {
	fmt.Printf(message + "\n", params...)
}