package transaction

import (
	"time"
	"tripatra-api/material"
)

const (
	StatusReceive = "receive" //transaksi telah diterima gudang (warehouse_officer)
	StatusIssue   = "issue"   //transaksi terjadi issue maka tdk di proses
	StatusUpdated = "updated" //transaksi updated materials
	StatusDeleted = "deleted" //transaksi dihapus
)

const (
	// NewWarehouse  = "new-warehouse"
	AddWarehouse  = "add-warehouse"
	TakeWarehouse = "take-warehouse"
)

type Transaction struct {
	ID                int
	MaterialID        int
	Quantity          int
	UpdatedAt         time.Time
	CreatedAt         time.Time
	Status            string
	Reason            string
	SenderID          int
	ReceiverID        int
	WarehouseCategory string

	Material material.Material
}
