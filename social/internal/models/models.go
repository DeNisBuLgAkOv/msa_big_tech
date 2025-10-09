package models

import "time"

const (
	//заявка отправлена, ждет ответа
	FriendRequestPending = "PENDING"
	// заявка принята
	FriendRequestAccepted = "ACCEPTED"
	//заявка отклонена
	FriendRequestDeclined = "DECLINED"
)

type FriendRequest struct {
	ID         string
	FromUserID string
	ToUserID   string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
