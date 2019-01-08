package eventmanager

// Subscriber is subscriber interface
type Subscriber interface {
	Call(data interface{})
}
