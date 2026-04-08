package notification

import (
	"context"
	"srv-api/chat/repositories/notification"

	m "github.com/srv-api/middlewares/middlewares"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FcmService interface {
	SaveOrUpdateToken(userID, token, deviceType string) error
	GetTokenByUserID(userID string) (string, error)
	DeleteToken(userID string) error
	SendToDevice(userFCMToken string, data map[string]interface{}) error
}

type fcmService struct {
	client *messaging.Client
	repo   notification.FCMRepository
	jwt    m.JWTService
}

func NewFCMService(repo notification.FCMRepository, credentialsPath string, jwtS m.JWTService) (FcmService, error) {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return &fcmService{
		client: client,
		jwt:    jwtS,
	}, nil
}
