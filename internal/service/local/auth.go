package local

type authService struct {
	privateKey string
}

func NewAuthService(privateKey string) *authService {
	return &authService{privateKey: privateKey}
}
