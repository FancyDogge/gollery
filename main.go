package main

import (
	"fmt"
	"net/http"
)

/*
<h1>ZULUL</h1> - это тело, которое отправляется в респонсе с помощью
http.ResponseWriter

С помощью него же отправляются и все остальные составляющие респонса, как я порнимаю
например сет куки или хэдеры

дефолтный статус код - 200, если ничего не указывать. Также статус код должен быть назначен в первую очередь(вроде бы)

мы можем использовать w(response writer) в fmt.Fprint(w, "<h1>ZULUL</h1>"), т.к. он реализует Write интерфейс, который принимает Fprint
*/
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>ZULUL</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on port :3000...")
	http.ListenAndServe(":3000", nil)
}
