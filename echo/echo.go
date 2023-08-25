package echo

type EchoService interface {
	Echo(message string) string
}
