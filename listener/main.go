package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"

	B "github.com/dylankilkenny/go-bitdb"
	"github.com/dylankilkenny/watch-cash/listener/mail"
	"github.com/joho/godotenv"
)

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

type ApiResponse struct {
	Status int      `json:"status"`
	Error  ApiError `json:"errors"`
	Data   []User   `json:"data"`
}

type ApiError struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TxInputAmount struct {
	Amount int `json:"amount"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	version := 3
	bitdburl := "https://bitdb.network/q/"
	bitdb := B.New(version, "qrqm04uwrd0wwaguxea079h0znswwt3quuejvl6zd6", bitdburl)

	bitsocketurl := "http://192.168.0.214:4000/s/"
	bitsocket := B.NewSocket(version, bitsocketurl)

	jq := ".[] | { txId: .tx.h, from: [.in[] | { prevTransactionId: .e.h, sender: .e.a }], to: [.out[] | { receiver: .e.a?, amount: .e.v? }] }"
	events, err := bitsocket.Stream("", jq)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for event := range events {
		if event.Type != "mempool" {
			continue
		}

		tx := castTransaction(event.Data)
		fmt.Println("Data recieved:", tx)

		checkTxInputs(tx, bitdb)
		checkTxOutputs(tx)

	}
}

func checkTxInputs(tx transaction, bitdb *B.Request) {
	for _, input := range tx.From {
		users := fetchUsers(input.Sender)
		if len(users) == 0 {
			continue
		}

		for _, user := range users {
			txhash := bitdb.TxHash(input.PrevTransactionID)
			resp, err := bitdb.Request(txhash, ".[] | .out[0] | {amount: .e.v}")
			if err != nil {
				fmt.Println("BitDb Request:", err)
			}
			confirmed, _ := resp.Confirmed.(map[string]interface{})
			var amount TxInputAmount
			amount = json.Unmarshal(resp.Confirmed, TxInputAmount)
			bch := float64(confirmed["amount"].(float64)) / math.Pow10(int(8))
			sendMail(user, input.Sender, "sent", fmt.Sprintf("%f", bch), tx.TxID)
		}
	}
}

func checkTxOutputs(tx transaction) {
	for _, output := range tx.To {
		users := fetchUsers(output.Receiver)
		if len(users) == 0 {
			continue
		}
		bch := float64(output.Amount) / math.Pow10(int(8))
		for _, user := range users {
			sendMail(user, output.Receiver, "recieved", fmt.Sprintf("%f", bch), tx.TxID)
		}
	}
}

func fetchUsers(address string) []User {

	message := map[string]string{
		"address": address,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:3001/api/secret/users", bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API_KEY", os.Getenv("WATCHCASHAPIKEY"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result ApiResponse
	json.NewDecoder(resp.Body).Decode(&result)

	switch result.Status {
	case 200:
		return result.Data
	default:
		return nil
	}
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

func sendMail(user User, address, txType, amount, txID string) {
	newmail := mail.NewEmail([]string{user.Email})
	absPath, _ := filepath.Abs("mail/template.html")
	newmail.Send(absPath, map[string]string{
		"name":    user.Name,
		"address": address,
		"txType":  txType,
		"amount":  amount,
		"txID":    txID,
	})
}
