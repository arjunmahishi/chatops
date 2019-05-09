package commanders

import (
	"fmt"
	"testing"
)

func TestSyncCommands(t *testing.T) {
	err := SyncCommands("./commands.json")
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, command := range CommanderList {
		fmt.Printf("\n%v\n\n\n", command)
	}
}
