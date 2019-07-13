package app

import (
	"context"
	"github.com/pusher/chatkit-server-go"
)

// CreateUser method for user creating
func CreateUser(ctx context.Context, client *chatkit.Client, userOptions chatkit.CreateUserOptions) error {
	err := client.CreateUser(ctx, userOptions)
	if err != nil {
		return err
	}
	return nil
}


// CreateUsers method for users creating
func CreateUsers(ctx context.Context, client *chatkit.Client, userOptions []chatkit.CreateUserOptions) error {
	err := client.CreateUsers(ctx, userOptions)
	if err != nil {
		return err
	}
	return nil
}

// AddUserToRoom method for users creating
func AddUsersToRoom(ctx context.Context, client *chatkit.Client, usersToRoom UsersToRoom) error {
	err := client.AddUsersToRoom(ctx, usersToRoom.RoomID, usersToRoom.UserIDs)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUsers method for deleting user
func DeleteUser(ctx context.Context, client *chatkit.Client, userID string) error {
	err := client.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}



