package chat

type Message struct {
	User string `json:"user"`
	Body string `json:"body"`
}

func (self *Message) String() string {
	return self.User + ": " + self.Body
}
