package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/frankegoesdown/chat/internal/app"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/pusher/chatkit-server-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	gw "pkg/chat/api/chat"
)

type config struct {
	DbConnectionString string `yaml:"dbConnectionString"`
	InstanceLocator    string `yaml:"instanceLocator"`
	SecretKey          string `yaml:"secretKey"`
	CoreAddress        string `yaml:"coreAddress"`
}

func (c *config) newConfig() *config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var conf config
	conf.newConfig()

	err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Println(err)
	}

	db, err := sql.Open("postgres", conf.DbConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	client, err := chatkit.NewClient(conf.InstanceLocator, conf.SecretKey)
	if err != nil {
		log.Fatal(err)
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

		err = app.CreateUser(ctx, db, client, userOptions)
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
		var message app.MultiPartMessage
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		Parts := message.Parts
		// TODO: another way to get it from client
		messageOptions := chatkit.SendMultipartMessageOptions{
			RoomID:   message.RoomID,
			SenderID: message.SenderID,
			Parts:    Parts}

		log.Println(message)
		log.Println(messageOptions)
		_, err = app.SendMultipartMessage(ctx, client, messageOptions)
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

//func run() error {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	// Register gRPC server endpoint
//	// Note: Make sure the gRPC server is running properly and accessible
//	mux := runtime.NewServeMux()
//	opts := []grpc.DialOption{grpc.WithInsecure()}
//	err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
//	if err != nil {
//		return err
//	}
//
//	// Start HTTP server (and proxy calls to gRPC server endpoint)
//	return http.ListenAndServe(":8081", mux)
//}

//func main() {
//	flag.Parse()
//	defer glog.Flush()
//
//	if err := run(); err != nil {
//		glog.Fatal(err)
//	}
//}
