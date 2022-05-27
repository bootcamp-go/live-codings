package main

import "fmt"

type ErrGOT struct {
	message string
	code    int
}

func NewErrGOT(message string, code int) *ErrGOT {
	return &ErrGOT{
		message: message,
		code:    code,
	}
}

func (e *ErrGOT) Error() string {
	return fmt.Sprintf("%s Error with code %d", e.message, e.code)
}

func translateToHighValyrian(word string) (string, error) {
	if word == "Cersei will..." {
		return "", NewErrGOT("No more GOT spoilers!", 1)
	} else if word == "Daenerys will ..." {
		return "", NewErrGOT("Daenerys Targaryen is immune to fire!", 2)
	}
	return word, nil
}

func main() {
	//word := "Daenerys will ..."
	word := "Cersei will..."

	t, err := translateToHighValyrian(word)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("translate: %s", t)
}
