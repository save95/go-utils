package maker

import "testing"

func TestAppID(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(AppID("xg"))
	}
}
