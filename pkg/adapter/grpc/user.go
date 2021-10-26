package grpc

import (
	"19chat/pkg/user"
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrInternalServerError = errors.New("internal-server-error")
	ErrUnauthorized        = errors.New("unauthorized")
)

func DefaultUserService(userService user.UserService) UserServiceServer {
	return &userServer{
		userService,
		map[uint]*UserService_SubscribeToChatServer{},
	}
}

type userServer struct {
	service       user.UserService
	subscriptions map[uint]*UserService_SubscribeToChatServer
}

func (u *userServer) CreateUser(context context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {

	err := u.service.CreateAccount(request.Email, request.Name, request.LastName, request.Password)
	if err != nil {
		if err == user.ErrUserExists {
			return nil, user.ErrUserExists
		}
		return nil, ErrInternalServerError
	}

	return &CreateUserResponse{}, nil
}

func (u *userServer) LoginUser(context context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {

	token, err := u.service.Login(request.Email, request.Password)
	if err != nil {
		if err == user.ErrUserOrPasswordIncorrect {
			return nil, user.ErrUserOrPasswordIncorrect
		}
		fmt.Println(err)
		return nil, ErrInternalServerError
	}

	return &LoginUserResponse{AccessToken: token}, nil
}

func (u *userServer) GetUserInfo(context context.Context, request *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	userID, err := u.authorize(context)
	if err != nil {
		return nil, err
	}

	info, err := u.service.GetInfo(userID)
	if err != nil {
		return nil, ErrInternalServerError
	}

	response := GetUserInfoResponse{
		Email:    info.Email,
		Name:     info.Name,
		LastName: info.LastName,
	}

	return &response, nil
}

// Chat
func (u *userServer) SendMessage(context context.Context, message *SendMessageRequest) (*emptypb.Empty, error) {
	userID, err := u.authorize(context)
	if err != nil {
		return nil, err
	}

	sendMessage := user.Message{
		From: int64(userID),
		To:   message.To,
		Text: message.Text,
		Date: time.Now().Unix(),
	}

	err = u.service.SendMessage(sendMessage)
	if err != nil {
		return nil, err
	}

	u.sendMessageToSubscriptions(sendMessage)

	return &emptypb.Empty{}, nil
}

func (u *userServer) SubscribeToChat(empty *emptypb.Empty, subscription UserService_SubscribeToChatServer) error {
	context := subscription.Context()
	userID, err := u.authorize(context)
	if err != nil {
		return err
	}

	u.subscriptions[userID] = &subscription
	u.sendListOfUsersToAllSubscribers()

	<-context.Done()

	delete(u.subscriptions, userID)
	u.sendListOfUsersToAllSubscribers()

	return nil
}

func (u *userServer) GetMessageHistory(context context.Context, request *GetMessageHistoryRequest) (*Messages, error) {
	userID, err := u.authorize(context)
	if err != nil {
		return nil, err
	}
	toUserID := request.UserId

	messageHistory, err := u.service.GetMessageHistory(userID, uint(toUserID))
	if err != nil {
		return nil, err
	}

	messages := Messages{
		Messages: []*Message{},
	}

	for _, message := range messageHistory {
		messages.Messages = append(messages.Messages, &Message{
			From: message.From,
			To:   message.To,
			Date: message.Date,
			Text: message.Text,
		})
	}

	return &messages, nil
}

func (u *userServer) authorize(context context.Context) (uint, error) {
	headers, ok := metadata.FromIncomingContext(context)
	if !ok {
		return 0, ErrUnauthorized
	}
	token := headers["authorization"]
	if len(token) < 1 || token[0] == "" {
		return 0, ErrUnauthorized
	}
	userID, err := u.service.Authorize(token[0])
	if err != nil {
		return 0, ErrUnauthorized
	}
	return userID, nil
}

func (u *userServer) sendListOfUsersToAllSubscribers() {
	userList := UserList{
		User: []*User{},
	}

	for key := range u.subscriptions {
		userInfo, err := u.service.GetInfo(key)
		if err != nil {
			fmt.Println("Warning: cannot find user on list")
			continue
		}
		userList.User = append(userList.User, &User{
			UserId:   int64(userInfo.Model.ID),
			Name:     userInfo.Name,
			LastName: userInfo.LastName,
		})
	}

	streamMessageList := StreamMessage_List{
		List: &userList,
	}

	for _, value := range u.subscriptions {
		if value == nil {
			continue
		}
		(*value).Send(&StreamMessage{
			Body: &streamMessageList,
		})
	}
}

func (u *userServer) sendMessageToSubscriptions(message user.Message) {
	fromSubscription := u.subscriptions[uint(message.From)]
	toSubscription := u.subscriptions[uint(message.To)]

	messageBody := StreamMessage_Messages{
		Messages: &Messages{
			Messages: []*Message{
				{
					From: message.From,
					To:   message.To,
					Text: message.Text,
					Date: message.Date,
				},
			},
		},
	}

	streamMessage := StreamMessage{
		Body: &messageBody,
	}

	if fromSubscription != nil {
		(*fromSubscription).Send(&streamMessage)
	}

	if toSubscription != nil && message.From != message.To {
		(*toSubscription).Send(&streamMessage)
	}
}
