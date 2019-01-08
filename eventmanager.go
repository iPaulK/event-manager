package eventmanager

// EventManager is event manager structure
type EventManager struct {
	Storage Storage
}

// DispatchEvent dispatches an event to all registered subscribers.
func (e *EventManager) DispatchEvent(eventName string, data interface{}) {
	subscribers := e.Storage.getSubscribers(eventName)

	for _, subscriber := range subscribers {
		subscriber.Call(&data)
	}
}

// AddEventSubscriber adds a subscriber to an event.
func (e *EventManager) AddEventSubscriber(eventName string, subscriber Subscriber) {
	e.Storage.addSubscriber(eventName, subscriber)
}

// RemoveEventSubscriber an event subscriber from the specified events.
func (e *EventManager) RemoveEventSubscriber(eventName string, subscriber Subscriber) {
	e.Storage.removeSubscriber(eventName, subscriber)
}

// NewEventManager is factory method for the event manager
func NewEventManager(
	storage Storage,
) *EventManager {
	return &EventManager{
		storage,
	}
}
