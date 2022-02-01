package erratum

import (
	"errors"
	"fmt"
)

// The second set of paren here indicates a named return value
func Use(opener ResourceOpener, input string) (err error) {
	var res Resource

	// Keep trying to open the resource if this is a transient error, otherwise
	// keep looping until we get another error or we get no error
	for {
		res, err = opener()
		if errors.Is(err, nil) {
			break // Opened successfully stop trying
		}

		transient := TransientError{}
		if !errors.As(err, &transient) {
			return err
		}
	}
	fmt.Println("Successfully opened the resource")
	defer res.Close()
	defer func() {
		fmt.Println("Calling recover function...")
		if r := recover(); r != nil {
			switch pObject := r.(type) {
			case string:
				fmt.Println("Handling string..")
				err = errors.New(pObject)
			case FrobError:
				fmt.Println("Handling FrobError...")
				res.Defrob(pObject.defrobTag)
				err = pObject
			case error:
				fmt.Println("Handling generic error..")
				err = pObject
			default:
				err = errors.New("Unknown panic!")
			}
		}
	}()

	// This call might panic
	res.Frob(input)

	fmt.Println("Successfully recovered")
	return err
}
