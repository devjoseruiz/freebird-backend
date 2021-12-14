package models

/*LoginResponse contains the token that is returned after a successful login*/
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
