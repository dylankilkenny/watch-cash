package bitsocket

// Transaction struct
type Transaction struct {
	Type string `json:"type"`
	Data data   `json:"data"`
}

type data struct {
	TxID string `json:"txId"`
	From []struct {
		PrevTransactionID string `json:"prevTransactionId"`
		Sender            string `json:"sender"`
	} `json:"from"`
	To []struct {
		Receiver string `json:"receiver"`
		Amount   int    `json:"amount"`
	} `json:"to"`
}

// Addresses returns a slice of addresses from the
// tx inputs and outputs
func (t *Transaction) Addresses() []string {
	addresses := make([]string, 0)
	for _, input := range t.Data.From {
		addresses = append(addresses, input.Sender[12:])
	}

	for _, output := range t.Data.To {
		addresses = append(addresses, output.Receiver[12:])
	}

	return addresses
}
