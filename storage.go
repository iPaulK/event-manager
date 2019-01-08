package eventmanager

// Storage the storage interface
type Storage interface {
	getSubscribers(string) []Subscriber
	addSubscriber(string, Subscriber)
	removeSubscriber(string, Subscriber)
}

// Memory storage structure
type Memory struct {
	// Map of registered subscribers.
	subscribers map[string][]Subscriber
}

// getSubscribers gets the subscribers of a specific event or all subscribers.
func (m *Memory) getSubscribers(eventName string) []Subscriber {
	return m.subscribers[eventName]
}

// addSubscriber adds an event subscriber that listens on the specified events.
func (m *Memory) addSubscriber(eventName string, subscriber Subscriber) {
	subscribers := m.subscribers[eventName]
	m.subscribers[eventName] = append(subscribers, subscriber)
}

// removeSubscriber removes an event subscriber from the specified events.
func (m *Memory) removeSubscriber(eventName string, subscriber Subscriber) {
	var key int
	for k, v := range m.subscribers[eventName] {
		if subscriber == v {
			key = k
			break
		}
	}
	m.subscribers[eventName] = append(m.subscribers[eventName][:key], m.subscribers[eventName][1+key:]...)
}

// NewMemoryStorage is factory method for the storage structure
func NewMemoryStorage() *Memory {
	subscribers := make(map[string][]Subscriber)
	return &Memory{
		subscribers,
	}
}
