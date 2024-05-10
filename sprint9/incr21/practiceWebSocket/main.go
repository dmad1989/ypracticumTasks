package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// обслуживает WebSocket-функция, принимающая один аргумент websocket.Conn
// в это соединение можно писать и читать из него
func wsHandle(ws *websocket.Conn) {
	for {
		var msg string
		// ожидаем запрос
		// Message — это Codec
		// Codec имеет методы Receive и Send
		// для чтения/записи websocket.Conn
		// с кодированием/декодированием
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println(err)
			return
		}
		// отправляем ответ
		err = websocket.Message.Send(ws, fmt.Sprintf("Привет, %s!", msg))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	// простым приведением типов конвертируем функцию wsHandle,
	// обслуживающую веб-сокет, в тип websocket.Handler
	// websocket.Handler может быть зарегистрирован
	// обычным образом, как и http.Handler
	http.Handle("/", websocket.Handler(wsHandle))
	http.ListenAndServe(":3000", nil)
}
