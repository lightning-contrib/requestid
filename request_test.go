package requestid

import (
	"github.com/go-labx/lightning"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefault(t *testing.T) {
	// Create a mock request with an empty x-request-id header
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	ctx, _ := lightning.NewContext(rr, req)

	// Create a mock handler function that sets the X-Request-ID header
	handler := Default()
	handler(ctx)

	requestId := rr.Header().Get("X-Request-ID")
	// Check that the X-Request-ID header was set in the response
	if requestId == "" {
		t.Errorf("X-Request-ID header not set")
	}
	if len(requestId) != 24 {
		t.Errorf("X-Request-ID size not valid")
	}
}

func TestNew(t *testing.T) {
	// Create a mock request with an empty x-request-id header
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	ctx, _ := lightning.NewContext(rr, req)

	// Create a mock handler function that sets the X-Request-ID header
	handler := New(WithSize(16), WithHeaderKey("X-Trace-ID"), WithAlphabet("1234567890"))
	handler(ctx)

	requestId := rr.Header().Get("X-Trace-ID")
	// Check that the X-Request-ID header was set in the response
	if requestId == "" {
		t.Errorf("X-Request-ID header not set")
	}
	if len(requestId) != 16 {
		t.Errorf("X-Request-ID size not valid")
	}

}
