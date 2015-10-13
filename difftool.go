package difftool

import (
	// "io"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
)

//Open opens difftool specified in DIFFTOOL environment variable
func Open(a, b []byte) error {
	diffTool := os.Getenv("DIFFTOOL")
	if diffTool == "" {
		return errors.New("DIFFTOOL environment variable not definded")
	}

	var fA, fB *os.File
	var err error

	if fA, err = ioutil.TempFile("", "A"); err != nil {
		return err
	}

	if fB, err = ioutil.TempFile("", "B"); err != nil {
		return err
	}

	defer os.Remove(fA.Name())
	defer os.Remove(fB.Name())

	if _, err = fA.Write(a); err != nil {
		return err
	}

	if _, err = fB.Write(b); err != nil {
		return err
	}

	fA.Close()
	fB.Close()

	if err = exec.Command(diffTool, fA.Name(), fB.Name()).Run(); err != nil {
		return err
	}

	return nil
}
