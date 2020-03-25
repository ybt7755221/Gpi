package service

import (
	"gpi/libraries/jwtGo"
)

type ToolService struct {
}

func (t *ToolService) GetJWTToken(params map[string]string) (string, error) {
	return jwtGo.GenerateToken(params)
}

