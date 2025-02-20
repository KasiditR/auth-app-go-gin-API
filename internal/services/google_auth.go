package services

import (
	"encoding/json"
	"github.com/KasiditR/auth-app-go-gin-API/internal/models"
	"github.com/go-resty/resty/v2"
	"log"
)

type GoogleTokenInfo struct {
	Sub        string `json:"sub"`
	Email      string `json:"email"`
	GiveName   string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
}

func GetGoogleUserData(idToken string) (*models.UserData, error) {
	var googleToken GoogleTokenInfo
	client := resty.New()
	resp, err := client.R().
		Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken)

	if err != nil {
		return nil, err
	}

	log.Println(string(resp.Body()))
	err = json.Unmarshal(resp.Body(), &googleToken)
	if err != nil {
		return nil, err
	}

	userData := &models.UserData{
		ID:        googleToken.Sub,
		FirstName: googleToken.GiveName,
		LastName:  googleToken.FamilyName,
		FullName:  googleToken.Name,
		Email:     googleToken.Email,
		Picture:   googleToken.Picture,
	}

	return userData, nil
}
