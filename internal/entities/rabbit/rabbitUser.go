package rabbitEntities

type UserCreation struct {
	ID            string `json:"id"`
	ServiceKey    string `json:"service_key"`
	Exchange      string `json:"exchange,omitempty"`
	ReadExchange  string `json:"read_exchange,omitempty"`
	WriteExchange string `json:"write_exchange,omitempty"`
}
