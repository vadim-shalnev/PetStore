package responder

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	var status int
	switch err.Error() {
	case "не удалось прочитать запрос", "не удалось дессериализировать JSON":
		status = http.StatusBadRequest
	case "не верный логин", "не верный пароль", "вы успешно вышли из сервиса":
		status = http.StatusUnauthorized
	case "ошибка в работе dadata", "ошибка запроса Select":
		status = http.StatusInternalServerError
	default:
		status = http.StatusBadRequest
	}
	http.Error(w, err.Error(), status)
}

func HandleSuccess(w http.ResponseWriter, data interface{}) {
	respJSON, err := json.Marshal(data)
	if err != nil {
		HandleError(w, err)
		return
	}
	w.Write(respJSON)
}
