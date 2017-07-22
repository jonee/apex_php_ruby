package main

import (
	"github.com/apex/go-apex"
	// "github.com/kr/pretty"

	"encoding/json"

	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"

	// "net/http"

	// "fmt"
	"log"
	// "os"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {

		var err error

		var t map[string]interface{}
		err = json.Unmarshal(event, &t)
		if err != nil {
			log.Println("parameters not provided: " + err.Error())
			return nil, err
		}

		log.Println(t)

		return map[string]interface{}{"hi": "hello"}, nil

	})
}

