package conversation

// Get Conversation

type GetConversationsResponse struct {
	PageSize int
	Embedded
}

type Embedded struct {
}

type GetConversationResponse struct {
	Id             string
	Name           string
	DisplayName    *string
	ImageUrl       string
	State          string
	SequenceNumber string
	TimeStamp      Timestamp
}

type Timestamp struct {
	Created string
}

type GetConversatonProperties struct {
	Ttl        int
	Type       string
	CustomData interface{}
}
