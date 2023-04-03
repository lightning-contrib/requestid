package requestid

import (
	"github.com/go-labx/lightning"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDMiddleware(t *testing.T) {
	// Create a mock request with an empty x-request-id header
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	ctx, _ := lightning.NewContext(rr, req)

	// Create a mock handler function that sets the X-Request-ID header
	handler := RequestID()
	handler(ctx)

	// Check that the X-Request-ID header was set in the response
	if rr.Header().Get("X-Request-ID") == "" {
		t.Errorf("X-Request-ID header not set")
	}
}
