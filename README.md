# fc-utils
Repositório da aula de manipulação de enventos.


## Anotações

Inicializando o módulo

```bash

go mod init github.com/wandermaia/fc-utils


```


## Testes

Vamos utilizar suite de testes do testify (github.com/stretchr/testify). A suite ajuda a preparar os testes para que não seja necessário ficar repetindo código de testes.

Tier Down - Uma operação que pode ser executada ao término da execução do método de testes.


## RabbitMQ

Será utilizada uma versão docker.

```bash

docker compose up -d
docker compose ps

curl http://localhost:15672

```
O diferencial do RabbitMQ é que ele não manda mensagens diretamente para a fila, ele manda para uma Exchange.
O Exchange é um tipo de roteador. Com isso, ele consegue enviar a mesma mensagem para mais de uma fila, por exemplo uma de auditoria.

O roteamento é baseado em chaves (Routing Key). 

Tipos de Exchange
    - Direct: Encaminhamento direto
    - Fanout: Não precisa de routing key. Manda para todas as filas sempre.
    - Topic: Muito parecida com a Direct, porém com mais recursos. Aceita por exemplo wildcard na key.

As conexões, tem "subconexões" que são chamadas de Channels, que acessam as filas ou publicando mensegens nas Exchanges.

autoAck é utilizado somente quando a mensagem pode ser perdida.

## Referências

Testify

https://github.com/stretchr/testify

Simlulador RabbitMQ

https://tryrabbitmq.com/

Go RabbitMQ Client Library

https://github.com/rabbitmq/amqp091-go
