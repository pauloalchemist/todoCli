package internal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func GetTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Tarefa não pode ser em branco")
	}

	return s.Text(), nil
}
