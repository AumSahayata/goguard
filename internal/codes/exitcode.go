package codes

import "fmt"

type ExitCodeError struct {
	Code int
	Msg  string
}

func (e *ExitCodeError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}

	return fmt.Sprintf("exit with code %d", e.Code)
}
