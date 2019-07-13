package app

import (
	"context"
	"github.com/pusher/chatkit-server-go"
)

// GetRoomMessages method for room creating
func GetRoomMessages(ctx context.Context, client *chatkit.Client, roomMessages RoomMessages) ([]chatkit.Message, error) {
	messages, err := client.GetRoomMessages(ctx, roomMessages.RoomID, roomMessages.MessagesOption)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// SendMessage method for room creating
func SendMessage(ctx context.Context, client *chatkit.Client, message chatkit.SendMessageOptions) (uint, error) {
	messageID, err := client.SendMessage(ctx, message)
	if err != nil {
		return 0, err
	}
	return messageID, nil
}

// SendMultipartMessage method for room creating
func SendMultipartMessage(ctx context.Context, client *chatkit.Client, message chatkit.SendMultipartMessageOptions) (uint, error) {
	messageID, err := client.SendMultipartMessage(ctx, message)
	if err != nil {
		return 0, err
	}
	return messageID, nil
}

//// SendSimpleMessage method for room creating
//func SendSimpleMessage(ctx context.Context, client *chatkit.Client, message chatkit.SendSimpleMessageOptions) (uint, error) {
//	messageID, err := client.SendSimpleMessage(ctx, message)
//	if err != nil {
//		return 0, err
//	}
//	return messageID, nil
//}
