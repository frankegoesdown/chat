package app

import (
	"context"
	"github.com/pusher/chatkit-server-go"
	"log"
)

// CreateRoom method for room creating
func CreateRoom(ctx context.Context, client *chatkit.Client, roomOptions chatkit.CreateRoomOptions) (chatkit.Room, error) {
	randomName := RandStringBytesMaskImprSrcSB(32)
	roomOptions.Name = randomName

	// ToDo change userID for user from header
	roomOptions.CreatorID = "first"
	log.Println(roomOptions)
	room, err := client.CreateRoom(ctx, roomOptions)
	if err != nil {
		return  room, err
	}
	return room, nil
}



// DeleteRoom method for room creating
func DeleteRoom(ctx context.Context, client *chatkit.Client, roomID string) error {
	err := client.DeleteRoom(ctx, roomID)
	if err != nil {
		return  err
	}
	return nil
}
