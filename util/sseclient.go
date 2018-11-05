package sseclient

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

/*
	Thanks to:
	https://github.com/peterhellberg/sseclient/blob/master/sseclient.go
*/

// OpenURL opens a connection to a stream of server sent events
func OpenURL(url string) (events chan []byte, err error) {

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
