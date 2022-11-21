package vonage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DawnKosmos/vonage-go-sdk/internal/conversation"
	"github.com/DawnKosmos/vonage-go-sdk/jwt"
	"io"
	"log"
	"net/http"
)

/*
DawnKosmos
using version 1
*/

var conversationUrl = "https://api.nexmo.com/v0.1/conversations"

type Conversation struct {
	Client *http.Client
	Id     string
	Href   string
}

/*
PseudoCode
NewConversation(Client http.Client, jwt Token)
GetConversation(conversationId)
Conversation
	AddMember
	IdInt
*/

func NewConversation(in conversation.CreateConversationRequest, client *http.Client, gen *jwt.Generator) (*Conversation, error) {
	token, err := gen.GenerateToken()
	if err != nil {
		return nil, err
	}

	by, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", conversationUrl, bytes.NewBuffer(by))
	if err != nil {
		return nil, err
	}

	resp, err := requestAuth(client, req, token)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, failedRequest(resp)
	}
	var data conversation.CreateConversationResponse
	if err := processResponse(resp, &data); err != nil {
		return nil, err
	}

	return &Conversation{
		Client: client,
		Id:     data.Id,
		Href:   data.Href,
	}, nil
}

func (c *Conversation) AddMember(userId string, gen *jwt.Generator) (*conversation.CreateMemberResponse, error) {
	token, err := gen.GenerateToken()
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(conversation.CreateMemberRequest{
		Action: "join",
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/members", conversationUrl, c.Id), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := requestAuth(c.Client, req, token)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, failedRequest(resp)
	}

	var data conversation.CreateMemberResponse
	err = processResponse(resp, &data)
	return &data, err
}

func CreateUser(input conversation.CreateUserRequest, client *http.Client, gen *jwt.Generator) (string, error) {
	token, err := gen.GenerateToken()
	if err != nil {
		return "", err
	}
	body, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.nexmo.com/v0.1/users"), bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	resp, err := requestAuth(client, req, token)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", failedRequest(resp)
	}

	var data conversation.CreateUserResponse
	if err := processResponse(resp, &data); err != nil {
		return "", err
	}

	return data.Id, nil
}

func CreateJWTToken(gen *jwt.Generator, username string) *jwt.Generator {
	gen = gen.AddPath(jwt.Path{Path: "/*/conversations/**"}).AddPath(jwt.Path{Path: "/*/image/**"}).AddPath(jwt.Path{Path: "/*/session/**"})
	return gen
}

// Help funcs
func processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
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

func failedRequest(resp *http.Response) error {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error processing response: %v", err)
		return err
	}
	return errors.New(fmt.Sprintf("StatusCode: %d, Body: %s", resp.StatusCode, string(body)))
}

func requestAuth(client *http.Client, req *http.Request, token string) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	return client.Do(req)
}
