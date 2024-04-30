package events

import (
	"sync"
	"time"
)

// Interface para Carregar os dados do evento
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{} // Uma interface vazia para que possa pegar diversos payloads e diversos formatos.
}

// Interface para criar o executor (handler) das operações do evento.
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup) //event é uma variável do tipo EventInterface
}

// Interface para o gerenciamento dos eventos
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error                          // Vai disparar (fire) as ações do evento
	Remove(eventName string, handler EventHandlerInterface) error // Remover evento do dispatcher - remove da fila
	Has(eventName string, handler EventHandlerInterface) bool     // Verificar se existe um event com um handler
	Clear() error                                                 // Limpar o event dipatcher, limpa todos os eventos registrados
}

/*

evento - evento criado
dispatcher -> dispatch(evento)

*/
