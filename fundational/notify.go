package fundational

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Notify(username, msg string) {
	postmsg := struct {
		Userids string `json:"userids"`
		Notify  string `json:"notify"`
	}{
		username,
		msg,
	}
	buf, _ := json.Marshal(postmsg)
	content := string(buf)
	Inf("notify: ", content)
	reader := strings.NewReader(content)
	_, err := http.Post("http://172.17.0.1:88/api/users/notify", "application/json; charset=utf-8", reader)
	if err != nil {
		Inf("notify error: ", err.Error())
		return
	}
}
