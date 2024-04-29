package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// Registrar os handlers. Cada evento pode ter vários handlers.
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Essa função é atachada ao EventDispatcher criado acima (primeiro parêntese) e ela valida se o handler
// passado por parâmetro já existe para o evento (também passado por parâmetro) antes de adicionar
// o handler. Em caso de erro, ele retorna mensagem de erro informando que o handler já existe.
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
