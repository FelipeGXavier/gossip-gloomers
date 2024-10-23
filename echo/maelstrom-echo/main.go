package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("echo", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "echo_ok"

		// Send message back to the client who sent the message
		return n.Reply(msg, body)
	})

	// n.Run() read messages from STDIN and call handler associated to the type
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
