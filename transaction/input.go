package transaction

type InputTransaction struct {
	MaterialID        int    `json:"material_id" binding:"required"`
	Quantity          int    `json:"quantity" binding:"required"`
	Status            string `json:"status"`
	Reason            string `json:"reason"`
	SenderID          int    `json:"sender_id"`
	ReceiverID        int    `json:"receiver_id" binding:"required"`
	WarehouseCategory string `json:"warehouse_category" binding:"required"`
}

type InputApproval struct {
	Status        string `json:"status" binding:"required"`
	Reason        string `json:"reason" binding:"required"`
	TransactionID int    `json:"transaction_id" binding:"required"`
}
