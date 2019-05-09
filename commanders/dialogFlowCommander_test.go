package commanders

import (
	"fmt"
	"testing"

	"github.com/arjunmahishi/Chatops/mocks"
)

func Test_sendDialog(t *testing.T) {
	resp, err := sendDialog("how are you doing")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	fmt.Println(resp.Result.Fulfillment.Speech)
}

func TestDialogFlowCommand_Execute(t *testing.T) {
	mockPayload := mocks.PayloadHandler{}

	mockPayload.On("GetMessage").Return("how are you doing")

	var dialog DialogFlowCommand
	_, err := dialog.Execute(&mockPayload)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}
