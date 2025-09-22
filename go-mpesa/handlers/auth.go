package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func GetAccessToken() (string, error) {
	consumerKey := os.Getenv("MPESA_CONSUMER_KEY")
	consumerSecret := os.Getenv("MPESA_CONSUMER_SECRET")

	// Encode credentials
	credentials := consumerKey + ":" + consumerSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))

	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+encoded)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.AccessToken, nil
}
