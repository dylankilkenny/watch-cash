package bitsocket

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	b64Query := b64.StdEncoding.EncodeToString(allTransactions)
	events, err := stream("https://bitsocket.org/s/" + b64Query)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	for event := range events {
		tx := Transaction{}
		err := json.Unmarshal(event, &tx)
		if err != nil {
			fmt.Println("JSON ERROR: ", err)
		}
		fmt.Println(tx.Data)

	}
}

/*
	Thanks to:
	https://github.com/peterhellberg/sseclient/blob/master/sseclient.go
*/

// stream opens a connection to a url and recieves a stream of server sent events
func stream(url string) (events chan []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got response status code %d\n", resp.StatusCode)
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
