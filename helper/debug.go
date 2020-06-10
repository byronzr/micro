package helper

import (
	"encoding/json"
	"fmt"
)

func DD(i interface{}) {
	if content, err := json.MarshalIndent(i, "", "\t"); err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("\033[32m%s\033[0m", string(content)))
	}
}
