package python

import (
	"os"
	"os/exec"
)

const pythonExecutable = "python/.venv/Scripts/python.exe"
const pythonSourcePath = "python/"

func RunPy(sourceFile string, args string) {
	cmd := exec.Command(pythonExecutable, pythonSourcePath+sourceFile, args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		println(err.Error())
	}
}
