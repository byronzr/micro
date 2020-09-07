package micro

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"encoding/json"
)

func Inf(msg ...interface{}) {
	//logger := log.New(os.Stdout, "\033[32mINF\033[0m ", log.LstdFlags)
	logger := log.New(os.Stdout, "\033[32mINF\033[0m ", log.LstdFlags)
	logger.Print(msg...)
}

func PPP(msg ...interface{}) {
	logger := log.New(os.Stdout, "\033[05m///\033[0m ", log.LstdFlags)
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

func DD(o ...interface{}) {
	buf := make([]byte, 0)
	for idx, s := range o {
		out, err := json.MarshalIndent(s, "", "\t")
		if err != nil {
			panic(err)
		}
		buf = append(buf, []byte(fmt.Sprintf("\n\033[33m----%02d----------------------------------------------\033[0m\n%s", idx, string(out)))...)
	}
	fmt.Println(string(buf))
}