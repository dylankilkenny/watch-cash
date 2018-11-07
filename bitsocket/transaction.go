package bitsocket

type Transaction struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	TxId string `json:"txId"`
	From []struct {
		PrevTransactionID string `json:"prevTransactionId"`
		Sender            string `json:"sender"`
	} `json:"from"`
	To []struct {
		Receiver string `json:"receiver"`
		Amount   int    `json:"amount"`
	} `json:"to"`
}
