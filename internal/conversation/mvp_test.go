package conversation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DawnKosmos/vonage-go-sdk/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var urltest = "https://api.nexmo.com/v0.1/conversations"

func TestMVP(t *testing.T) {

	gen, err := jwt.NewGeneratorFromFilename("45a30037-7919-4919-a47d-2e354cce75d9", "private.key")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	tk, err := gen.GenerateToken()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	client := &http.Client{}

	by, err := json.Marshal(CreateConversationRequest{
		Name:        "TestBro",
		DisplayName: "Murrat",
		Properties:  Property{60},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	req, err := http.NewRequest("POST", urltest, bytes.NewBuffer(by))
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tk)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp.StatusCode)

	var a CreateConversationResponse

	err = processResponse(resp, &a)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(a)

}

func processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error processing response: %v", err)
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		log.Printf("Error processing response: %v", err)
		return err
	}
	return nil
}
