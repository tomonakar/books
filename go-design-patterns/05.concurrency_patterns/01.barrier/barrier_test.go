package barrier

import (
	"fmt"
	"strings"
	"testing"
)

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"https://httpbin.org/get"}

		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"https://malformed-url", "https://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		fmt.Println(result)
		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"https://httpbin.org/headers", "https://httpbin.org/User-Agent"}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log(result)

	})
}
