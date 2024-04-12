package read

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadValues() (passwords map[string]string,keys []string, err error ) {
	f, err := os.Open("password.txt")
		if err != nil {
		return nil,nil, errors.New("error: file not found")
	}
	defer f.Close()
	passwords = make(map[string]string)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		part := strings.Split(scanner.Text(), "||")
		key := part[0]
		keys = append(keys, key)
		if len(part) == 2 {
			passwords[part[0]] = part[1]
		}
		fmt.Println("Key: ", key)
	}
	if len(passwords) == 0 {
		return nil,nil, errors.New("error: no key value found")
	}
	return passwords,keys, nil
}
