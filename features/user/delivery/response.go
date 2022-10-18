package delivery

import "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessLogin(msg string, token string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
		"token":   token,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "login":
		cnv := core.(domain.Core)
		res = LoginResponse{Username: cnv.Username, Password: cnv.Password}
	}

	return res
}
