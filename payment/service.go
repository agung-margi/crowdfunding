package payment

type service struct {
}

type Service interface {
	GetToken() (string, error)
}
