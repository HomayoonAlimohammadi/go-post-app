package main

type AuthServer struct {
	jwtManager JWTManager
}

func (a *AuthServer) Login(authorName string) (string, error) {
	token, err := a.jwtManager.Generate(authorName)
	if err != nil {
		return "", err
	}

	return token, nil
}
