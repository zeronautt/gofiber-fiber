package main

import (
	"testing"
)

func Test_Ctx_IP_Malformed_XForwardedFor(t *testing.T) {
	app := New()
	app.config.ProxyHeader = HeaderXForwardedFor

	tests := []struct {
		headerVal string
		expected  string
	}{
		{", 203.0.113.195", "203.0.113.195"},
		{"203.0.113.195, , 70.41.3.18", "203.0.113.195"},
		{"  ,  203.0.113.195  , 70.41.3.18", "203.0.113.195"},
		{",,", "127.0.0.1"}, // Should fallback to remote IP
	}

	for _, tt := range tests {
		ctx := &Ctx{
			app:      app,
			headers:  map[string]string{HeaderXForwardedFor: tt.headerVal},
			remoteIP: "127.0.0.1",
		}
		result := ctx.IP()
		if result != tt.expected {
			t.Errorf("For header '%s', expected IP '%s', got '%s'", tt.headerVal, tt.expected, result)
		}
	}
}
