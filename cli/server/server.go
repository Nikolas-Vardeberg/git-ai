package server

type ServerConfigS struct {
	BaseURL string `json:"base_url"`
}

// This should be the URL where the api is hosted. For local development, use localhost. For production, use the deployed URL. # Disclaimer my api route will not work for you :) 
var ServerConfig = ServerConfigS{
	BaseURL: "https://git-ai-git-timeline-niklusrooks-projects.vercel.app",
}