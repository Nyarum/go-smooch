package smooch

type messageType string

const (
	messageTextType messageType = "text"
)

type roleType string

const (
	appUserType  roleType = "appUser"
	appMakerType roleType = "appMaker"
)

type RequestMessage struct {
	Text string      `json:"text"`
	Role roleType    `json:"role"`
	Type messageType `json:"type"`
	Name string      `json:"name"`
}

type ResponseMessage struct {
	Message struct {
		ID        string  `json:"_id"`
		AuthorID  string  `json:"authorId"`
		Role      string  `json:"role"`
		Type      string  `json:"type"`
		Name      string  `json:"name"`
		Text      string  `json:"text"`
		AvatarURL string  `json:"avatarUrl"`
		Received  float64 `json:"received"`
	} `json:"message"`
	Conversation struct {
		ID          string `json:"_id"`
		UnreadCount int    `json:"unreadCount"`
	} `json:"conversation"`
}

func (s *Smooch) SendMessage(userID, text string) (*ResponseMessage, error) {
	requestMessage := &RequestMessage{
		Text: text,
		Role: appMakerType,
		Type: messageTextType,
		Name: "admin",
	}

	resp, err := s.requestPost("appusers/"+userID+"/messages", requestMessage)
	if err != nil {
		return nil, err
	}

	responseMessage := &ResponseMessage{}

	err = resp.ToJSON(responseMessage)
	if err != nil {
		return nil, err
	}

	return responseMessage, nil
}
