package vonage

import (
	"fmt"
	"github.com/DawnKosmos/vonage-go-sdk/internal/conversation"
	"github.com/DawnKosmos/vonage-go-sdk/jwt"
	"net/http"
	"testing"
)

func TestConversationUserMember(t *testing.T) {
	gen, err := jwt.NewGeneratorFromFilename("45a30037-7919-4919-a47d-2e354cce75d9", "internal/conversation/private.key")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	client := &http.Client{}

	u1, err := CreateUser(conversation.CreateUserRequest{
		Name:        "TestUser1",
		DisplayName: "Alice",
		ImageUrl:    "",
	}, client, gen)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	u2, err := CreateUser(conversation.CreateUserRequest{
		Name:        "TestUser2",
		DisplayName: "Bob",
		ImageUrl:    "",
	}, client, gen)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println("Alice", u1)
	fmt.Println("Bob", u2)

	con, err := NewConversation(conversation.CreateConversationRequest{
		Name:        "TestConversation",
		DisplayName: "Kitty Club",
		ImageUrl:    "",
		Properties:  conversation.Property{60},
	}, client, gen)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(con)

	_, err = con.AddMember(u1, gen)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	_, err = con.AddMember(u2, gen)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

}
