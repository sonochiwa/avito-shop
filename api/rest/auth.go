package rest

type AuthHandlerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandlerResponse struct {
	Token string `json:"token"`
}
