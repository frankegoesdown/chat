package app

import "github.com/pusher/chatkit-server-go"

// Message struct
type Message struct {
	RoomID   string `json:"room_id"`
	Text     string `json:"text"`
	SenderID string `json:"sender_id"`
}

// Room struct
type Room struct {
	RoomID string `json:"room_id"`
}

// User struct
type User struct {
	UserID string `json:"user_id"`
}

// UsersToRoom struct for add users to room
type UsersToRoom struct {
	RoomID  string   `json:"room_id"`
	UserIDs []string `json:"user_ids"`
}

// UsersToRoom struct for getting messages from room
// InitialID *uint   Starting ID of messages to retrieve
// Direction *string One of older or newer
// Limit     *uint   Number of messages to retrieve
type RoomMessages struct {
	RoomID         string                         `json:"room_id"`
	MessagesOption chatkit.GetRoomMessagesOptions `json:"messages_option"`
}
