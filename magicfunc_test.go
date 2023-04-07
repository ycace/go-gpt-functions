package magic_test

import (
	"context"
	"log"
	"testing"

	"github.com/ycace/go-gpt-functions"
)

func TestMagicObject(t *testing.T) {
	ctx := context.TODO()
	magicObj := magic.NewMagicObject("YOUR_API_KEY")
	response, err := magicObj.DoMagic(ctx, "Give me a random number between 0 and 100", nil)
	if err != nil {
		t.Error("doMagic err: ", err)
	}
	log.Print("magicResponse random: ", response)

	response, err = magicObj.DoMagic(ctx, "Give me a simple user tip for server errors", nil)
	if err != nil {
		t.Error("doMagic err: ", err)
	}
	log.Print("magicResponse server err info: ", response)

	response, err = magicObj.DoMagic(ctx, "Output information about the word 'default', in json format", nil)
	if err != nil {
		t.Error("doMagic err: ", err)
	}
	log.Print("magicResponse json word info: ", response)

	response, err = magicObj.DoMagic(ctx, "Generate me a random user information with five random properties and values, in json format.", nil)
	if err != nil {
		t.Error("doMagic err: ", err)
	}
	log.Print("magicResponse user info: ", response)

	response, err = magicObj.DoMagic(ctx, "Generate me a random user information with six random properties and values, in json format.", map[string]interface{}{
		"name": "ycace",
		"age":  "18",
	})
	if err != nil {
		t.Error("doMagic err: ", err)
	}
	log.Print("magicResponse user info with input: ", response)
}