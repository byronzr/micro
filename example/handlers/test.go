package handlers
import(

	"encoding/json"
	"net/http"

)

type POST struct{}
type OPTIONS struct{}

func (POST) ReportThisweek(r *http.Request) (response []byte, err error) {
	msg := "POST.ReprotThisweek."
	return json.Marshal(msg)
}

func (OPTIONS) ReportThisweek(r *http.Request) (response []byte, err error) {
	msg := "POST.ReprotThisweek."
	return json.Marshal(msg)
}