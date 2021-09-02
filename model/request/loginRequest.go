package request

type LoginRequest struct {
	IP       string
	Password map[string]string
}
