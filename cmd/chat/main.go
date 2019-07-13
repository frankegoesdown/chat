package main

import (
	"context"
	"encoding/json"
	"github.com/frankegoesdown/chat/internal/app"
	"github.com/go-chi/chi"
	"github.com/pusher/chatkit-server-go"
	"log"
	"net/http"
)

func main() {
	client, err := chatkit.NewClient("v1:us1:89d2abb9-db24-4959-9721-1559acafe0cd", "49d4d308-af02-4425-90d7-be258dde616d:TNZ9zAm+TwynR0cnspwNkkDp5dadYf810LCH7s+Unsw=")
	if err != nil {
		log.Println("can't make new client")
	}
	ctx := context.Background()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/create_room", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var roomOptions chatkit.CreateRoomOptions
		err := json.NewDecoder(r.Body).Decode(&roomOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		room, err := app.CreateRoom(ctx, client, roomOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		log.Println(room)
		// ToDo: make response

	}))

	r.Post("/delete_room", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var room app.Room
		err := json.NewDecoder(r.Body).Decode(&room)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		err = app.DeleteRoom(ctx, client, room.RoomID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		// ToDo: make response

	}))

	r.Post("/add_users_to_room", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var usersToRoom app.UsersToRoom
		err := json.NewDecoder(r.Body).Decode(&usersToRoom)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		err = app.AddUsersToRoom(ctx, client, usersToRoom)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		// ToDo: make response

	}))

	r.Post("/create_user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var userOptions chatkit.CreateUserOptions
		err := json.NewDecoder(r.Body).Decode(&userOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		err = app.CreateUser(ctx, client, userOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("user successfully created"))

	}))

	r.Post("/create_users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var userOptions []chatkit.CreateUserOptions
		err := json.NewDecoder(r.Body).Decode(&userOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		err = app.CreateUsers(ctx, client, userOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("users successfully created"))

	}))

	r.Post("/delete_user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var user app.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		err = app.DeleteUser(ctx, client, user.UserID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		// ToDo: make response

	}))

	r.Get("/get_room_messages", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var user app.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		err = app.DeleteUser(ctx, client, user.UserID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		// ToDo: make response

	}))

	r.Post("/send_message", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		//var messageOptions chatkit.SendMessageOptions
		var message app.Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		// TODO: another way to get it from client
		messageOptions := chatkit.SendMessageOptions{
			RoomID:   message.RoomID,
			Text:     message.Text,
			SenderID: message.SenderID}

		_, err = app.SendMessage(ctx, client, messageOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		// ToDo: make response

	}))

	r.Post("/send_multipart_message", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		//var messageOptions chatkit.SendMessageOptions
		var message app.Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		// TODO: another way to get it from client
		messageOptions := chatkit.SendMessageOptions{
			RoomID:   message.RoomID,
			Text:     message.Text,
			SenderID: message.SenderID}

		_, err = app.SendMessage(ctx, client, messageOptions)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		// ToDo: make response

	}))

	http.ListenAndServe(":3001", r)
	//client.AddUsersToRoom()
}
