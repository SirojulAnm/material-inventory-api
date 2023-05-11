package notification

import "time"

type Notification struct {
	ID         int
	MaterialID int
	Message    string
	UpdatedAt  time.Time
	CreatedAt  time.Time
	SenderID   int
	ReceiverID int
}
