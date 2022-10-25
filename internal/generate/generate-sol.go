package generate

import (
	"fmt"
	"io/ioutil"
)

// This method will be used to create the .sol file for the contract
func GenerateSol(path string, contract string) error {
	// write the contract to the file
	err := ioutil.WriteFile(path, []byte(contract), 0644)
	if err != nil {
		return err
	}

	// write to the file
	fmt.Println("Contract written to file")
	return nil
}
