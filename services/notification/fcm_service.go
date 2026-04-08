package notification

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
)

// SendToDevice - kirim notifikasi ke satu device
func (f *fcmService) SendToDevice(userFCMToken string, data map[string]interface{}) error {
	if userFCMToken == "" {
		log.Println("No FCM token for user")
		return nil
	}

	title, _ := data["title"].(string)
	body, _ := data["body"].(string)
	senderID, _ := data["sender_id"].(string)
	senderName, _ := data["sender_name"].(string)
	msgType, _ := data["type"].(string)

	message := &messaging.Message{
		Token: userFCMToken,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: map[string]string{
			"sender_id":         senderID,
			"sender_name":       senderName,
			"type":              msgType,
			"current_user_id":   senderID,
			"current_user_name": senderName,
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				ChannelID:   "chat_messages",
				Color:       "#075E54",
				ClickAction: "FLUTTER_NOTIFICATION_CLICK",
			},
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Sound: "default",
				},
			},
		},
	}

	_, err := f.client.Send(context.Background(), message)
	if err != nil {
		log.Printf("Failed to send FCM: %v", err)
		return err
	}

	log.Printf("✅ Push notification sent to %s", userFCMToken)
	return nil
}

func (f *fcmService) SaveOrUpdateToken(userID, token, deviceType string) error {
	return f.repo.SaveOrUpdateToken(userID, token, deviceType)
}

func (f *fcmService) GetTokenByUserID(userID string) (string, error) {
	return f.repo.GetTokenByUserID(userID)
}

func (f *fcmService) DeleteToken(userID string) error {
	return f.repo.DeleteToken(userID)
}
