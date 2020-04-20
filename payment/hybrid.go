package payment

type HybridAccount struct {
	CreditAccount
	CheckingAccount
}

func (h *HybridAccount) AvailableFunds() float32 {
	println("Getting total funds from hybrid account...")
	return h.CreditAccount.AvailableFunds() + h.CheckingAccount.AvailableFunds()
}
