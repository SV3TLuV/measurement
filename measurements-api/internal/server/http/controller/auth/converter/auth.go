package converter

import (
	model "measurements-api/internal/model"
	model2 "measurements-api/internal/server/http/controller/auth/model"
)

func ToAuthResultViewFromService(result *model.AuthResult) *model2.AuthResultView {
	if result == nil {
		return nil
	}

	return &model2.AuthResultView{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}
}

func ToLoginDataFromRequest(request *model2.LoginRequest) *model.LoginData {
	if request == nil {
		return nil
	}

	return &model.LoginData{
		Login:    request.Login,
		Password: request.Password,
	}
}
