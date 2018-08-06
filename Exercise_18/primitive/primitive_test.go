package primitive

import (
	"os"
	"testing"

	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
	m.Run()
}
func TestWithMode(t *testing.T) {
	result := WithMode(ModeCombo)
	if result == nil {
		t.Error("Expected string but got not result")
	}
}

func TestTransform(t *testing.T) {
	f, _ := os.Open("../img/ghoper.jpg")
	opts := WithMode(ModeCombo)
	_, err := Transform(f, "jpg", 1, opts)
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}

func TestTransformNegativePrimitive(t *testing.T) {
	f, _ := os.Open("../img/ghoper.jpg")
	opts := WithMode(ModeCombo)
	_, err := Transform(f, "jpg", -1, opts)
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

// func TestTransformNegativeImageReader(t *testing.T) {
// 	f, _ := os.Open("../img/ghoper.jpg")
// 	ioutil.ReadAll(f)
// 	opts := WithMode(ModeCombo)
// 	_, err := Transform(f, "jpg", 1, opts)
// 	if err == nil {
// 		t.Error("Expected error but got no error")
// 	}
// }

func TestRunPrimitive(t *testing.T) {
	args := WithMode(ModeCircle)
	_, err := runPrimitive("../img/ghoper.jpg", "../img/out.jpg", 1, args...)
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}

func TestCreateTempFile(t *testing.T) {
	_, err := createTempFile("", "txt")
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}
func TestCreateTempFileNegative(t *testing.T) {
	_, err := createTempFile("/invalid/invalid", "txt")
	if err == nil {
		t.Error("Expected error but got no error")
	}
}
