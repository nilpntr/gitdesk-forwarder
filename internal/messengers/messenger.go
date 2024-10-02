package messengers

type Messenger interface {
	SendMessage(webhookUrl, title, description, url string) error
}

var Messengers = map[string]Messenger{
	"slack": &SlackMessenger{},
}
