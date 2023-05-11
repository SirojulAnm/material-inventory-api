package notification

type Service interface {
	Add(input InputNotification) (Notification, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Add(input InputNotification) (Notification, error) {
	notification := Notification{}
	notification.MaterialID = input.MaterialID
	notification.Message = input.Message
	notification.SenderID = input.SenderID
	notification.ReceiverID = input.ReceiverID

	newNotification, err := s.repository.Save(notification)
	if err != nil {
		return newNotification, err
	}

	return newNotification, nil
}
