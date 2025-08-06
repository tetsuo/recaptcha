# recaptcha

Verify Google reCAPTCHA challenge responses using the [reCAPTCHA siteverify API](https://developers.google.com/recaptcha/docs/verify).

## Usage

```go
resp, err := recaptcha.Verify(context.Background(),
  http.DefaultClient, "response-token", "your-secret-key", "user-ip")
if err != nil {
  // handle error
}
if resp.Success {
  fmt.Println("reCAPTCHA verified!")
} else {
  fmt.Println("Verification failed:", resp.ErrorCodes)
}
```
