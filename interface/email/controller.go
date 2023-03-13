package email

type EmailHandler interface {
	ComposeEmail() error
}
