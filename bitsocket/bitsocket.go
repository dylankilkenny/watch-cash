package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dylankilkenny/watch-cash/models"
	"github.com/dylankilkenny/watch-cash/util"
)

var allTransactions = []byte(`
  {
	"v": 3,
	"q": {
	  "find": {
	  }
	},
	"r": {
	  "f": ".[] | { txId: .tx.h, from: [.in[] | { prevTransactionId: .e.h, sender: \"bitcoincash:\\(.e.a)\" }], to: [.out[] | { receiver: \"bitcoincash:\\(.e.a?)\", amount: .e.v? }] }"
	}
  }
`)

func main() {
	b64Query := b64.StdEncoding.EncodeToString(allTransactions)

	events, err := sseclient.OpenURL("https://bitsocket.org/s/" + b64Query)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for event := range events {
		tx := models.Transaction{}
		err := json.Unmarshal(event, &tx)
		if err != nil {
			fmt.Println("JSON ERROR: ", err)
		}
		fmt.Println(tx.Data)

	}
}
