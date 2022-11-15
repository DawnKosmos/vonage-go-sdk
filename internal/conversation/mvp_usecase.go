package conversation

/*Version 1
1.Create Conversation
2. Create Users
3. Create Members
*/

type CreateConversationRequest struct {
	Name        string   `json:"name,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	ImageUrl    string   `json:"image_url,omitempty"`
	Properties  Property `json:"properties,omitempty"`
}

type Property struct {
	Ttl int `json:"ttl,omitempty"`
}

type CreateConversationResponse struct {
	Id   string `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}

type CreateUserRequest struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}

type CreateUserResponse struct {
	Id   string `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}

type CreateMemberRequest struct {
	Action           string      `json:"action,omitempty"`
	UserId           string      `json:"user_id,omitempty"`
	MemberId         string      `json:"member_id,omitempty"`
	Channel          Channel     `json:"channel,omitempty"`
	Media            interface{} `json:"media,omitempty"`
	KnockingId       string      `json:"knocking_id,omitempty"`
	MemberIdInviting string      `json:"member_id_inviting,omitempty"`
}

type Channel struct {
	Type   string   `json:"type,omitempty"`
	LegId  string   `json:"leg_id,omitempty"`
	From   string   `json:"from,omitempty"`
	To     string   `json:"to,omitempty"`
	LegIds []string `json:"leg_ids,omitempty"`
}

type CreateMemberResponse struct {
	Id        string `json:"id,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	State     string `json:"state,omitempty"`
	Timestamp struct {
		Invited string `json:"invited,omitempty"`
		Joined  string `json:"joined,omitempty"`
		Left    string `json:"left,omitempty"`
	} `json:"timestamp,omitempty"`
	Channel   Channel `json:"channel,omitempty"`
	Href      string  `json:"href,omitempty"`
	Initiator struct {
		Joined struct {
			IsSystem bool   `json:"is_system,omitempty"`
			UserId   string `json:"user_id,omitempty"`
			MemberId string `json:"member_id,omitempty"`
		} `json:"joined,omitempty"`
	} `json:"initiator,omitempty"`
}

/*

Version 3

1. Create Conversation
2. Create User
3. Create member


type CreateConversationRequest struct {
	Name        string
	DisplayName *string
	ImageUrl    string
	Properties  ConversationProperty
	Numbers     []Number
	Callback    ConversationCallback
	Method      string
}

type ConversationProperty struct {
	Ttl        int
	Type       string
	CustomData interface{}
}

type ConversationCallback struct {
	Url       string
	EventMask string
}

type CallbackParams struct {
	ApplicationId string
	NccoUrl       string
}

type CreateConversationResponse struct {
	Id             string
	Name           string
	DisplayName    *string
	ImageUrl       string
	State          string
	SequenceNumber string
	TimeStamp      Timestamp
	Properties     ConversationProperty
	Numbers        []Number
}

type Number struct {
	Type        string
	Uri         string
	User        string
	Number      int
	ContentType string
	Extension   string
	Username    string
	Password    string
}

// User

type CreateUserRequest struct {
	Name        string
	DisplayName string
	ImageUrl    string
	Channels    interface{}
}

type CreateUserResponse struct {
	Id          string
	Name        string
	DisplayName string
	ImageUrl    string
	Properties  interface{}
	Channels    interface{}
}

type UserTimestamp struct {
	Create    string
	Updated   string
	Destroyed string
}

// Member

type CreateMemberRequest struct {
	State   string
	User    User
	Channel Channel
	//Media
	//KnockingId
	//MemberIdInviting
}
type Channel struct {
	Type   string
	LegId  string
	From   string
	To     string
	LegIds []string
}
type User struct {
	Id   string
	Name string
}

type CreateMemberResponse struct {
	Id             string
	ConversationId string
	State          string
	Timestamp      MemberTimestamp
}

type MemberTimestamp struct {
	Invited string
	Joined  string
	Left    string
}

type Initiator struct {
	Joined
}
type Joined struct {
	IsSystem bool
	UserId   string
	MemberId string
}
*/
