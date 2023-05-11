package transaction

type Service interface {
	Add(input InputTransaction) (Transaction, error)
	FindByID(ID int) (Transaction, error)
	Update(ID int, status string) (Transaction, error)
	FindBySender(senderID int) ([]Transaction, error)
	FindByReceiver(receiverID int) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Add(input InputTransaction) (Transaction, error) {
	transaction := Transaction{}
	transaction.MaterialID = input.MaterialID
	transaction.Quantity = input.Quantity
	transaction.Status = input.Status
	transaction.Reason = input.Reason
	transaction.SenderID = input.SenderID
	transaction.ReceiverID = input.ReceiverID
	transaction.WarehouseCategory = input.WarehouseCategory

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) FindByID(ID int) (Transaction, error) {
	var transaction Transaction
	transaction, err := s.repository.FindByID(ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) Update(ID int, status string) (Transaction, error) {
	transaction, err := s.repository.Update(ID, status)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) FindBySender(senderID int) ([]Transaction, error) {
	transaction, err := s.repository.FindBySender(senderID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) FindByReceiver(receiverID int) ([]Transaction, error) {
	transaction, err := s.repository.FindByReceiver(receiverID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
