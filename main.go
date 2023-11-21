package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed data.json
var data string

func main() {
	sch, err := jsonschema.Compile("./schema.json")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	var v interface{}
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		log.Fatal(err)
	}

	if err = sch.Validate(v); err != nil {
		log.Fatalf("%#v", err)
	}

	fmt.Println("is valid :)")
}
