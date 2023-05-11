package notification

type InputNotification struct {
	MaterialID int    `json:"material_id" binding:"required"`
	Message    string `json:"message" binding:"required"`
	SenderID   int    `json:"sender_id" binding:"required"`
	ReceiverID int    `json:"receiver_id" binding:"required"`
}
