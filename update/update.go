package update

import (
	"errors"
	"os"
	"bufio"
	"fmt"
	"strings"
	
)

func UpdateKeyValuesInPrompt() (string, string, error) {
	fmt.Println("Enter the new key and value in the format 'key value':")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("ivalid input format. please enter in the format 'key value'")
	}
	return parts[0], parts[1], nil

}
func UpdateKeyValuesInTextFile(passwords map[string]string, keys []string) (string, error) {
	file, err := os.OpenFile("password.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return "", errors.New("error opening file")
	}
	defer file.Close()
	for i, key := range keys {
		line := key + "||" + passwords[key]
		if i < len(keys)-1 {
			line += "\n"
		}
		if _, err := file.WriteString(line); err != nil {
			return "", errors.New("error writing to file")
		}
	}
	return "Successfully updated key value", nil
}
