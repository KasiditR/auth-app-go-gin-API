package services

import (
	"github.com/KasiditR/auth-app-go-gin-API/config"
	"github.com/KasiditR/auth-app-go-gin-API/internal/models"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type GitHubUser struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type GitHubEmail struct {
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
}

type GitHubAccessToken struct {
	AccessToken string `json:"access_token"`
}

func GetGithubAccessToken(code string) (string, error) {
	var gitHubAccessToken GitHubAccessToken
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"client_id":     config.LoadConfig().GithubClientId,
			"client_secret": config.LoadConfig().GithubSecret,
			"code":          code,
		}).
		SetResult(&gitHubAccessToken).
		Post("https://github.com/login/oauth/access_token")

	if err != nil {
		return "", err
	}

	return gitHubAccessToken.AccessToken, nil
}

func GetGithubUserData(accessToken string) (*models.UserData, error) {
	client := resty.New()

	var user GitHubUser
	_, err := client.R().
		SetAuthToken(accessToken).
		SetHeader("Accept", "application/json").
		SetResult(&user).
		Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}

	var emails []GitHubEmail
	_, err = client.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Accept", "application/json").
		SetResult(&emails).
		Get("https://api.github.com/user/emails")

	if err != nil {
		return nil, err
	}

	var primaryEmail string
	for _, email := range emails {
		if email.Primary {
			primaryEmail = email.Email
			break
		}
	}

	nameParts := []string{}
	if user.Name != "" {
		nameParts = append(nameParts, user.Name)
	}
	firstName := ""
	lastName := ""
	if len(nameParts) > 0 {
		firstName = nameParts[0]
		if len(nameParts) > 1 {
			lastName = nameParts[1]
		}
	}
	id := strconv.Itoa(user.ID)

	userData := &models.UserData{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		FullName:  user.Name,
		Email:     primaryEmail,
		Picture:   user.AvatarURL,
	}

	return userData, nil
}
