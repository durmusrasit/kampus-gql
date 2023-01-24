package loader

import "io/ioutil"

func ReadSchema(path string) (string, error) {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}
