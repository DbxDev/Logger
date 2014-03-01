package Logger

import (
	"flag"
	"fmt"
)

var logAll bool = false
var debug *bool = flag.Bool("debug", false, "Logger in debug.")
var info *bool = flag.Bool("info", false, "Logger in info.")
var warn *bool = flag.Bool("warn", false, "Logger in warn.")
var error *bool = flag.Bool("error", false, "Logger in error.")

func Init() {
	if !flag.Parsed() {
		flag.Parse()
		Debugf("Parsing flag")
	}
	var mode string = ""
	if *debug {
		mode = "DEBUG"
	} else if *info {
		mode = "INFO"
	} else if *warn {
		mode = "WARN"
	} else if *error {
		mode = "ERROR"
	} else {
		logAll = true
		mode = "DEFAULT. Print everything."
	}
	Infof("Logger initialized to %v", mode)
}
func Debugf(format string, args ...interface{}) {
	if logAll || *debug {
		logf("[DEBUG]", format, args...)
	}
}
func Infof(format string, args ...interface{}) {
	if logAll || *debug || *info {
		logf("[INFO]", format, args...)
	}
}
func Warnf(format string, args ...interface{}) {
	if logAll || *debug || *info || *warn {
		logf("[WARN]", format, args...)
	}
}
func Errorf(format string, args ...interface{}) {
	if logAll || *debug || *info || *warn || *error {
		logf("[ERROR]", format, args...)
	}
}

func logf(suffix string, format string, args ...interface{}) {
	fmt.Printf(suffix+format+"\n", args...)
}
