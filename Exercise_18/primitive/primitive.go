package primitive

import (
	"fmt"
	"os/exec"
	"strings"
)

//Mode means what kind of image the user want
type Mode int

//Mode defines the all modes supported by the primitive command
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

func Transform() {

}

//RunPrimitive is used to run primitive command
func RunPrimitive(inFile, outFile string, numShape int) (string, error) {
	args := (fmt.Sprintf("-i %s -o %s -n %d", inFile, outFile, numShape))
	cmd := exec.Command("primitive", strings.Fields(args)...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
