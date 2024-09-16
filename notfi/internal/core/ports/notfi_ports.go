package ports

type NotfiService interface {
	SendEmail() (bool, error)
}
