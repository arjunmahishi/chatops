package routes

import (
	"testing"

	"github.com/arjunmahishi/Chatops/mocks"
)

func Benchmark_getResponse(b *testing.B) {
	var mockPaylod mocks.PayloadHandler

	mockPaylod.On("GetMessage").Return("twilix-5xx")
	mockPaylod.On("GetSenderEmail").Return("arjun@exotel.in")
	mockPaylod.On("GetSenderName").Return("Arjun mahishi")

	for n := 0; n < b.N; n++ {
		getResponse(&mockPaylod)
	}
}
