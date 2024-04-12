package create

import (
	"errors"
	"fmt"
	"os"
	"password-manager/model"
	"password-manager/security"
)

func CreateValue(pair model.KeyValue) (error) {
	fmt.Println("key: ", pair.Key)
	fmt.Println("value: ", pair.Value)
	encryptedValue := string(security.EncryptPassword([]byte(pair.Value), "tomoe"))
	f, err := os.OpenFile("password.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("error opening file")
	}

	defer f.Close()
	if _, err := f.WriteString(pair.Key + "||" + encryptedValue + "\n"); err != nil {
		return errors.New("error writing to file")
	}
	return nil

}
