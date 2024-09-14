# Z-API Go

Essa biblioteca foi desenvolvida para o uso do [z-api.io](z-api), api do whatsapp usando a linguagem go.

Status: Em desenvolvimento

## Enviando Mensagem

Este é um exemplo de envio de mensagem a partir do client.

```go
import z_api "github.com/verbeux-ai/z-api-go"
import "context"

client := z_api.NewClient(
    z_api.WithToken("<token>"),
    z_api.WithInstance("<instance>"),
    z_api.WithSecret("<secret>"),
)
response, err := client.SendTextMessage(context.Background(), &z_api.TextMessageRequest{
    Phone:   text.Phone,
    Message: txt,
})
if err != nil {
    panic(err)
}
fmt.Println(response)
```

## Escutando mensagens

Este é um exemplo de como escutar mensagens no webhook

```go
import "github.com/verbeux-ai/z-api-go/listener"

whatsappListener := listener.NewMessageListener()
whatsappListener.HandleErrors(func (err error) {
    fmt.Println("fail", err)
})

// register listeners
whatsappListener.OnMessage(func (message *listener.WebhookMessage) error {
    if message.Text != nil {
        // treat your text message here
    }
    
    return nil
})

if err := whatsappListener.ReadBodyAsync(ctx.Request().Body); err != nil {
    panic(err)
}
```

## Features disponíveis

| Funcionalidade        | Implementado |
|-------------------------------|--------------|
| Enviar Mensagem de Texto      | ✔            |
| Atualizar Webhook de Entrega  | ✔            |
| Atualizar Webhook de Recebido | ✔            |
| Obter Status                  | ✔            |
| Obter Imagem do QR Code       | ✔            |
| Desconectar Instância         | ✔            |
| Obter Conversas               | ✔            |
| Obter Tags                    | ✔            |


> Você está convidado a contribuir ao repositório!