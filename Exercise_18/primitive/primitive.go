package primitive

import (
	"fmt"
	"os/exec"
	"strings"
)

//RunPrimitive is used to run primitive command
func RunPrimitive(inFile, outFile string, numShape int) (string, error) {
	args := strings.Fields(fmt.Sprintf("-i %s -o %s -n %d", inFile, outFile, numShape))
	cmd := exec.Command("primitive", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
