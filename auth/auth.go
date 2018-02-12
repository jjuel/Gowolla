package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	baseURL         = "https://api.dwolla.com"
	tokenURL        = "https://www.dwolla.com/oauth/v2/token"
	sandboxBaseURL  = "https://api-sandbox.dwolla.com"
	sandboxTokenURL = "https://sandbox.dwolla.com/oauth/v2/token"
)

// Request maps the request to /oauth/v2/token endpoint.
type Request struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

// TokenResp maps the response for the /oauth/v2/token endpoint.
type TokenResp struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// GetToken processes the auth request and returns the TokenResp
func GetToken(data Request) (TokenResp, error) {
	var tokenResp TokenResp

	// Set up the request
	buf := bytes.NewBuffer([]byte(data.GrantType))
	req, err := http.NewRequest("POST", sandboxTokenURL, buf)
	if err != nil {
		return tokenResp, nil
	}
	req.SetBasicAuth(data.ClientID, data.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Do the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return tokenResp, fmt.Errorf("Http request to %s failed %s", req.URL, err.Error())
	}
	defer resp.Body.Close()

	//dump, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%q\n", dump)

	if sc := resp.StatusCode; sc < 200 || sc > 299 {
		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&tokenResp); err != nil {
			return tokenResp, fmt.Errorf("Decoding error response failed %v", err)
		}

		return tokenResp, nil
	}

	// Decode the body.
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&tokenResp); err != nil {
		return tokenResp, fmt.Errorf("Decoding token response failed %v", err)
	}

	return tokenResp, nil
}
