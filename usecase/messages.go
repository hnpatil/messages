package usecase

type messagesImpl struct {
}

func NewMessages() Messages {
	return &messagesImpl{}
}
