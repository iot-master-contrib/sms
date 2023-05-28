package internal

type Sender interface {
	Connect(id, key string) error
	Send(sign, template string, params map[string]string, phones []string) error
}
