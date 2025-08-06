# github.com/tetsuo/recaptcha

Verify Google reCAPTCHA challenge responses using the [reCAPTCHA siteverify API](https://developers.google.com/recaptcha/docs/verify).

## Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tetsuo/recaptcha"
)

func main() {
	resp, err := recaptcha.Verify(context.Background(), http.DefaultClient,
    "response-token", "your-secret-key", "user-ip")
	if err != nil {
		// handle error
	}
	if resp.Success {
		fmt.Println("reCAPTCHA verified!")
	} else {
		fmt.Println("Verification failed:", resp.ErrorCodes)
	}
}
```
