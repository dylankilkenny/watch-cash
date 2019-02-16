package main

import (
	"encoding/json"
	"fmt"
	"math"
	"path/filepath"

	bitdb "github.com/dylankilkenny/go-bitdb"
	"github.com/dylankilkenny/watch-cash/listener/mail"
	"github.com/dylankilkenny/watch-cash/server/user"
	userModel "github.com/dylankilkenny/watch-cash/server/user/model"
)

func main() {
	version := 3
	bitsocketurl := "https://bitsocket.org/s/"
	bitdburl := "https://bitdb.network/q/"

	bitsocket := bitdb.NewSocket(version, bitsocketurl)
	bitdb := bitdb.New(version, "qrqm04uwrd0wwaguxea079h0znswwt3quuejvl6zd6", bitdburl)

	jq := ".[] | { txId: .tx.h, from: [.in[] | { prevTransactionId: .e.h, sender: \"bitcoincash:\\(.e.a)\" }], to: [.out[] | { receiver: \"bitcoincash:\\(.e.a?)\", amount: .e.v? }] }"
	events, err := bitsocket.Stream("", jq)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for event := range events {
		if event.Type != "mempool" {
			continue
		}
		tx := castTransaction(event.Data)
		fmt.Println(tx)

		for _, input := range tx.From {
			users, err := user.SubscribedUsers(input.Sender[12:])
			if err != nil {
				fmt.Println(err)
			}
			for _, user := range users {
				txhash := bitdb.TxHash(input.PrevTransactionID)
				fmt.Println(txhash)

				resp, err := bitdb.Request(txhash, ".[] | .out[0] | {amount: .e.v}")
				if err != nil {
					fmt.Println("BitDb", err)
				}
				confirmed, _ := resp.Confirmed.(map[string]interface{})
				fmt.Println(confirmed)
				bch := float64(confirmed["amount"].(float64)) / math.Pow10(int(8))
				sendMail(user, input.Sender[12:], "sent", fmt.Sprintf("%f", bch), tx.TxID)
			}
		}
		for _, output := range tx.To {
			users, err := user.SubscribedUsers(output.Receiver[12:])
			if err != nil {
				fmt.Println(err)
			}
			bch := float64(output.Amount) / math.Pow10(int(8))
			for _, user := range users {
				sendMail(user, output.Receiver[12:], "recieved", fmt.Sprintf("%f", bch), tx.TxID)
			}
		}
	}
}

type transaction struct {
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

func castTransaction(data interface{}) transaction {
	tx := transaction{}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON ERROR: ", err)
	}
	err = json.Unmarshal(dataBytes, &tx)
	if err != nil {
		fmt.Println("JSON ERROR: ", err)
	}
	return tx
}

func sendMail(user userModel.User, address, txType, amount, txID string) {
	newmail := mail.NewEmail([]string{user.Email})
	absPath, _ := filepath.Abs("mail/template.html")
	newmail.Send(absPath, map[string]string{
		"username": user.FirstName,
		"address":  address,
		"txType":   txType,
		"amount":   amount,
		"txID":     txID,
	})
}
