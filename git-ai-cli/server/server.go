package server

type ServerConfigS struct {
	BaseURL string `json:"base_url"`
}

var ServerConfig = ServerConfigS{
	BaseURL: "http://localhost:3000",
}