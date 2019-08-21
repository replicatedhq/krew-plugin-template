package main

import (
	"path"
	"io/ioutil"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)
func renderGoMod(templateContext TemplateContext) error {
	input, err := ioutil.ReadFile(path.Join("hack", "README.txt"))
        if err != nil {
		return errors.Wrap(err, "failed to read README.txt")
        }

        err = ioutil.WriteFile(path.Join("..", "README.md"), []byte(output), 0644)
        if err != nil {
                return errors.Wrap(err, "failed to write README.md")
	}
	
	return nil
}