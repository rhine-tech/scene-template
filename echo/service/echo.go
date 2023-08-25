package service

import "scene-template/echo"

type echoService struct {
}

func NewEchoService() echo.EchoService {
	return &echoService{}
}

func (e *echoService) Echo(message string) string {
	return "echo: " + message
}
