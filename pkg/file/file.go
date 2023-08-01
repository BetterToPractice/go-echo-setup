package file

import (
	"fmt"
	"os"
	"time"
)

func EnsureDir(fp string) error {
	return os.MkdirAll(fp, os.ModePerm)
}

func Create(name string) (*os.File, error) {
	return os.Create(name)
}

func Close(fd *os.File) error {
	return fd.Close()
}

func Remove(name string) error {
	return os.Remove(name)
}

func EnsureDirRW(dataDir string) error {
	err := EnsureDir(dataDir)
	if err != nil {
		return err
	}

	checkFile := fmt.Sprintf("%s/rw.%d", dataDir, time.Now().UnixNano())
	fd, err := Create(checkFile)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("open %s: rw permission denied", dataDir)
		}
		return err
	}

	if err := Close(fd); err != nil {
		return fmt.Errorf("close error: %s", err)
	}

	if err := Remove(checkFile); err != nil {
		return fmt.Errorf("remove error: %s", err)
	}

	return nil
}
