package micro

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
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
	logger := log.New(os.Stdout, "\033[31mERR\033[0m ", log.LstdFlags)
	nmsg := []interface{}{}
	nmsg = append(nmsg, "\033[31m")
	nmsg = append(nmsg, msg...)
	nmsg = append(nmsg, "\033[0m")
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

func SS(msg ...interface{}) {
	s := debug.Stack()
	logger := log.New(os.Stdout, "\033[31m>>> ", log.LstdFlags)
	nmsg := []interface{}{}
	nmsg = append(nmsg, msg...)
	nmsg = append(nmsg, "\n"+string(s))
	nmsg = append(nmsg, "\033[0m")
	logger.Print(nmsg...)
}
