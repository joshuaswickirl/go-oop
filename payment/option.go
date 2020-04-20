package payment

type Option interface {
	ProcessPayment(float32) bool
}
