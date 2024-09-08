package main

import (
	"encoding/json"
	"log"

	"github.com/ibiscum/go-oas-examples/3.0/recursive/recmodels"
	"github.com/spyzhov/ajson"
)

var node_obj recmodels.Node

func main() {
	body := []byte(`{"id":"aef04060-e0d5-4655-b660-9cbb728459e4","level":0,"nodes":[{"id":"503fbe63-b384-4911-be8f-0249260afce9","level":1,"relationship":"child"},{"id":"68d50b47-35b8-4ca1-a9b1-b4bc9adce6a4","level":1,"relationship":"child"},{"id":"ac16847c-1c48-437a-bbea-a7a968eece5a","level":1,"nodes":[{"id":"cc7f2cca-d0a4-473a-b40d-70c556e8b4a6","level":2,"relationship":"grandchild"},{"id":"ce6def39-1761-4435-bd58-bdb32f630295","level":2,"nodes":[],"relationship":"grandchild"}],"relationship":"child"}],"relationship":"parent"}`)

	root, err := ajson.Unmarshal(body)
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	nodes, err := root.JSONPath("$..[?(@.level)]")
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	for _, node := range nodes {
		result, err := ajson.Marshal(node)
		if err != nil {
			log.Fatalf("Error %s", err)
		}
		log.Printf("%s", string(result))

		err = json.Unmarshal(result, &node_obj)
		if err != nil {
			log.Fatalf("Error %s", err)
		}

		log.Printf("%s", *node_obj.Id)
	}

}
