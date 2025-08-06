package recaptcha

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RecaptchaResponse represents the response returned by the reCAPTCHA verification service.
type RecaptchaResponse struct {
	// Success indicates whether the reCAPTCHA verification was successful.
	Success bool `json:"success"`
	// ChallengeTimestamp is the timestamp when the challenge was solved.
	ChallengeTimestamp string `json:"challenge_ts"`
	// Hostname is the hostname of the site where the reCAPTCHA was solved.
	Hostname string `json:"hostname"`
	// ErrorCodes contains a list of error codes related to the reCAPTCHA validation, if any.
	ErrorCodes []string `json:"error-codes"`
}

const recaptchaVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

// Verify verifies a reCAPTCHA challenge response.
func Verify(ctx context.Context, c *http.Client, challengeResponse, secretKey, remoteIP string) (captchaResponse *RecaptchaResponse, err error) {
	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, http.MethodPost, recaptchaVerifyURL, strings.NewReader(url.Values{
		"secret":   {secretKey},
		"response": {challengeResponse},
		"remoteip": {remoteIP},
	}.Encode()))
	if err != nil {
		return
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var resp *http.Response
	resp, err = c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &captchaResponse)
	if err != nil {
		return nil, err
	}
	return
}
