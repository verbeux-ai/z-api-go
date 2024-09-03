package z_api

const (
	textMessageEndpoint     = "instances/%s/token/%s/send-text"               // POST
	webhookDeliveryEndpoint = "instances/%s/token/%s/update-webhook-delivery" // PUT
	webhookReceivedEndpoint = "instances/%s/token/%s/update-webhook-received" // PUT
	statusEndpoint          = "instances/%s/token/%s/status"                  // GET
	qrCodeImageEndpoint     = "instances/%s/token/%s/qr-code/image"           // GET
	disconnectEndpoint      = "instances/%s/token/%s/disconnect"              // GET
)
