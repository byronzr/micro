package fundational

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func Inf(msg ...interface{}) {
	logger := log.New(os.Stdout, "\033[32mINF\033[0m ", log.LstdFlags)
	logger.Print(msg...)
}

func PPP(msg ...interface{}) {
	logger := log.New(os.Stdout, "/// ", log.LstdFlags)
	logger.Print(msg...)
}

func Wrn(msg ...interface{}) {
	logger := log.New(os.Stdout, "\033[33mWRN\033[0m ", log.LstdFlags)
	logger.Print(msg...)
}

func Err(msg ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logger := log.New(os.Stdout, "\033[31mERR\033[0m ", log.LstdFlags)
	nmsg := []interface{}{fmt.Sprintf("%s:%d\n", file, line)}
	nmsg = append(nmsg, msg...)
	logger.Print(nmsg...)
}
