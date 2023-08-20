package schema

type CreateUserPayload struct {
	Name string `json:"name"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type NameResponse struct {
	Name string `json:"name"`
}
