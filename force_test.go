package force

import "testing"

func TestCreateClient(t *testing.T) {
	forceAPI := NewClientTest()
	if forceAPI == nil {
		t.Fatalf("Cannot Create Test Client")
	}

}
