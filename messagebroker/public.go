package messagebroker

// NewProducer is
func NewProducer(url string) Producer {
	return &producerImpl{
		defaultProducer: producerImplInit(url),
	}
}

func NewConsumer(url string) Consumer {
	consumerHandler := ConsumerHandler{
		Topic:           "onOrderCreated",
		Channel:         "test",
		FunctionHandler: func(message []byte) {},
	}
	return &consumerImpl{
		url: url,
		consumerHandlers: []ConsumerHandler{
			consumerHandler,
		},
	}
}
