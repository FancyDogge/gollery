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

// Я дурак и пытался написать уже написанное
// func notFoundError(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "404 page not found", 404)
// }

// router курильщика, чтобы разобраться в происходящем
func pathHandler(w http.ResponseWriter, r *http.Request) {
	/*
		У нашего http.Request есть поле (field) - URL *url.URL
		Это поинтер на URL struct, у которого в свою очередь есть поле Path string // path (relative paths may omit leading slash)
		Поэтому мы достаем все это дело через точки .

		URL.RawPath используется вместо URL.Path только если есть кейсы, где нужно отличить, допустим URL-encoded /(в виде %2F) от обычного /
	*/
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		//return 404 если пути не существует
		// http.NotFound(w, r)           //Это мой вариант до которго я додумался из документации
		w.WriteHeader(http.StatusNotFound)      //Это, и строка ниже - подход из туториала, хз, я так понял смысл в том,
		fmt.Fprint(w, "Sorry, page not found.") //что так у нас больше кастомизации тела ответа
		// http.Error(w, "Sorry, page not found.", http.StatusNotFound) //еще один вариант из туториала
	}
}

func main() {
	//добавляем роутер курильщика в рут путь / чтобы там он уже отдавал нужный хэндлер
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on port :3000...")
	http.ListenAndServe(":3000", nil)
}
