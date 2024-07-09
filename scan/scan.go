package scan

import (
	"bufio"
	"io"
)

func Scan(r io.Reader, f func(line string) error) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := f(scanner.Text())
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
