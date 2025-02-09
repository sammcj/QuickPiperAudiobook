package iconv

import (
	"QuickPiperAudiobook/internal/binarymanagers"
	"fmt"
	"io"
	"os/exec"
)

// Remove diacritics from text so that an english voice can read it
// without explicitly speaking the diacritics and messing with speech
func RemoveDiacritics(input io.Reader) (io.Reader, error) {
	if _, err := exec.LookPath("iconv"); err != nil {
		return nil, fmt.Errorf("iconv not found in PATH: %v", err)
	}

	command := "iconv -f UTF-8 -t ASCII//TRANSLIT//IGNORE"
	output, err := binarymanagers.RunPiped(command, input)
	if err != nil {
		return nil, err
	}
	return output.Stdout, nil
}
