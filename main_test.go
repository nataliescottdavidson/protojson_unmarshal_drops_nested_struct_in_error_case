package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"

	"example.com/main/gobindings"
)

/*
message NestedStruct {
	string bang = 1;
	int32 intButWeGetString = 2;
}

message Struct {
	string baz = 1;
	NestedStruct containsError = 2;
}
*/

func TestNestedUnmarshalWithError(t *testing.T) {

	jsonString := `{"baz": "buzz", "containsError": {"bang": "fizz", "intButWeGetString":"sillystring"}}`

	// encoding/json behavior
	var jsonOutput gobindings.Struct
	err := json.Unmarshal([]byte(jsonString), &jsonOutput)

	// error is expected
	assert.Equal(t, "json: cannot unmarshal string into Go struct field NestedStruct.containsError.intButWeGetString of type int32", err.Error())

	// but we get all the good fields still
	expected := gobindings.Struct{
		Baz:           "buzz",
		ContainsError: &gobindings.NestedStruct{Bang: "fizz"},
	}
	assert.Equal(t, expected.String(), jsonOutput.String())

	// protojson

	var protojsonOutput gobindings.Struct
	err = protojson.Unmarshal([]byte(jsonString), &protojsonOutput)
	assert.Equal(t, "proto: (line 1:71): invalid value for int32 field intButWeGetString: \"sillystring\"", err.Error())

	// we lost containsError { bang: "fizz"} !! bang didn't do anything wrong!!
	assert.Nil(t, protojsonOutput.ContainsError)

	// we still have fields outside the nested struct that contains the error though
	assert.Equal(t, "buzz", protojsonOutput.Baz)

}
