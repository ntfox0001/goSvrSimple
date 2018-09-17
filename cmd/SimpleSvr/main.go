package main

import (
	"flag"
	"fmt"

	log "github.com/inconshreveable/log15"
)

var (
	crashNow   string
	usePprof   string
	logFile    string
	logLevel   string
	logFormat  string
	cpuProfile string
	version    bool
)

func init() {
	flag.StringVar(&crashNow, "crashnow", "0", "crash now")
	flag.StringVar(&usePprof, "pprof", "0", "use pprof")
	flag.StringVar(&logFormat, "logformat", "term", "log format: json, fmt, term")
	flag.StringVar(&logFile, "log", "", "log file")
	flag.StringVar(&logLevel, "loglevel", "dbug", fmt.Sprintf("log level: %s, %s, %s, %s, %s", log.LvlDebug, log.LvlInfo, log.LvlWarn, log.LvlError, log.LvlCrit))
	flag.StringVar(&cpuProfile, "cpuprofile", "", "cpu profilename")
	flag.BoolVar(&version, "ver", false, "show version")
}
func main() {
	application.Application
}
