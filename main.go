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
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//вообще го смотрит на байты в респонс теле и определяет тип контент тайпа(опять же, вроде бы)
	//но вот так уот можно установить данный хедер мануально
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "<h1>ZULUL</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Sorry, page not found.", http.StatusNotFound)
// 	}
// }

// делаем свой router, который будет использовать Handler интерфейс с методом ServeHTTP(ResponseWriter, *Request)
// Вопрос - а зачем этот стракт вместо функции, которая была эо этого? (pathHandler)
// как я понял - для модульности и "интерфейсности") Например, чтобы поднять n разных сервисов с РАЗНЫМИ базами данных и коннектами.
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Sorry, page not found!", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on port :3000...")
	http.ListenAndServe(":3000", router)
}

/*
Памятка:

http.Handler - interface with the ServeHTTP method.

http.HandleFunc - a function !type! that accepts same arguments
as ServeHTTP method. So it implements http.Handler interface.

type http.HandlerFunc - The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers

Так что можно законвертить нашу ф-цию pathHandler, которая принимает responsewriter и request в HandleFunc ТИП!

fmt.Println("Starting the server on port :3000...")
http.ListenAndServe(":3000", http.HandlerFunc(pathHandler)) //http.HandlerFunc(pathHandler) - type conversion/cast, а не вызов ф-ции
*/
