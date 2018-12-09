package main

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	BitDB "github.com/dylankilkenny/watch-cash/bitdb"
	"github.com/dylankilkenny/watch-cash/bitsocket/mail"
	"github.com/dylankilkenny/watch-cash/bitsocket/transaction"
	"github.com/dylankilkenny/watch-cash/server/db"
	userModel "github.com/dylankilkenny/watch-cash/server/user/model"
)

// allTransactions is a bitsocket query which will
// retrieve all incoming transactions on the BCH network
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
	fmt.Println("[bitsocket] Started")
	db.Init()
	b64Query := b64.StdEncoding.EncodeToString(allTransactions)
	events, err := stream("https://bitsocket.org/s/" + b64Query)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	for event := range events {
		tx := transaction.Transaction{}
		err := json.Unmarshal(event, &tx)
		if err != nil {
			fmt.Println("JSON ERROR: ", err)
		}
		for _, input := range tx.Data.From {
			users, err := subscribedUsers(input.Sender[12:])
			if err != nil {
				fmt.Println(err)
			}
			for _, user := range users {
				bitdb := BitDB.New(3, "qq54zc33pttdp6l8ycnnj99ahan8a2hfrygqyz0fc3")
				txhash := BitDB.TxHash{Hash: input.PrevTransactionID}
				resp, err := bitdb.Request(txhash, ".[] | .out[0] | {amount: .e.v}")
				var obj map[string]interface{}
				err = json.Unmarshal([]byte(resp), &obj)
				if err != nil {
					fmt.Println(err)
				}
				confirmed, ok := obj["c"].(map[string]interface{})
				if !ok {
					fmt.Println("Ok: ", ok)
				}
				r := mail.NewEmail([]string{user.Email})
				absPath, _ := filepath.Abs("mail/template.html")
				r.Send(absPath, map[string]string{
					"username": user.FirstName,
					"address":  input.Sender[12:],
					"txType":   "sent",
					"amount":   string(int(confirmed["amount"].(float64))),
					"txID":     tx.Data.TxID,
				})
			}
		}
		for _, output := range tx.Data.To {
			users, err := subscribedUsers(output.Receiver[12:])
			if err != nil {
				fmt.Println(err)
			}
			for _, user := range users {
				r := mail.NewEmail([]string{user.Email})
				absPath, _ := filepath.Abs("mail/template.html")
				r.Send(absPath, map[string]string{
					"username": user.FirstName,
					"address":  output.Receiver[12:],
					"txType":   "sent",
					"amount":   string(output.Amount),
					"txID":     tx.Data.TxID,
				})
			}
		}
	}

}

//	https://github.com/peterhellberg/sseclient/blob/master/sseclient.go

// stream opens a connection to a url and recieves a stream of server sent events
func stream(url string) (events chan []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got response status code %d", resp.StatusCode)
	}
	events = make(chan []byte)
	var buf bytes.Buffer
	go func() {
		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "error during resp.Body read:%s\n", err)
				close(events)
			}
			switch {
			// event data
			case bytes.HasPrefix(line, []byte("data:")):
				buf.Write(line[6:])
			// end of event
			case bytes.Equal(line, []byte("\n")):
				b := buf.Bytes()
				if bytes.HasPrefix(b, []byte("{")) {
					buf.Reset()
					events <- b
				}
			default:
				fmt.Fprintf(os.Stderr, "Error: len:%d\n%s", len(line), line)
				close(events)
			}
		}
	}()
	return events, nil
}

func subscribedUsers(address string) ([]userModel.User, error) {
	var users []userModel.User
	var db = db.GetDB()

	if err := db.Joins("JOIN addresses ON addresses.user_id = users.id").Where("addresses.address = ?", address).Find(&users).Error; err == nil {
		return users, err
	}
	return users, nil
}
