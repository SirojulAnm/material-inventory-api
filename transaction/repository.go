package transaction

import "gorm.io/gorm"

type Repository interface {
	Save(material Transaction) (Transaction, error)
	FindByID(ID int) (Transaction, error)
	Update(ID int, status string, reason string) (Transaction, error)
	FindBySender(senderID int) ([]Transaction, error)
	FindByReceiver(receiverID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Preload("Material").Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(ID int, status string, reason string) (Transaction, error) {
	var transaction Transaction
	err := r.db.Model(&transaction).Where("id = ?", ID).Updates(Transaction{Status: status, Reason: reason}).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindBySender(senderID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Order("updated_at DESC").Not("status = ?", "deleted").Preload("Material").Where("sender_id = ?", senderID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) FindByReceiver(receiverID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Order("updated_at DESC").Not("status = ?", "deleted").Preload("Material").Where("receiver_id = ?", receiverID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
