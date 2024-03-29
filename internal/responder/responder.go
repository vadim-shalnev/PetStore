package responder

import "net/http"

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		panic(err)
	}
}

func HandleSuccess(w http.ResponseWriter, data interface{}) {

}
