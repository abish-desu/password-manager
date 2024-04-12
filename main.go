package main

import (
	"bufio"
	"fmt"
	"os"
	"password-manager/create"
	"password-manager/model"
	"password-manager/read"
	"password-manager/security"
	"password-manager/update"
	"password-manager/validation"
	"strconv"
	"strings"
)

func selectKey(passwords map[string]string, keys []string) (int, string) {
	for i, key := range keys {
		fmt.Printf("%d: Key: %s, value: %s\n", i, key, passwords[key])
	}
	fmt.Printf("Enter the index of the key you want to select: ")
	reader := bufio.NewReader(os.Stdin)
	indexStr, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(indexStr))
	if err != nil || index < 0 || index >= len(keys) {
		fmt.Println("Invalid index")
		return -1, ""
	}
	selectedKey := keys[index]
	fmt.Printf("You selected Key: %s, Value: %s\n", selectedKey, passwords[selectedKey])
	return index, selectedKey
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Invalid command")
		return
	}

	switch args[1] {
	case "create":
		msg, err := validate.ValidateCreateCommand(args)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
		pair := model.KeyValue{Key: args[2], Value: args[3]}
		err = create.CreateValue(pair)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Key-Value added successfully")
		}
	case "read":
		fmt.Println("Reading Key-Value")
		passwords, _, err := read.ReadValues()
		if err != nil {
			fmt.Println(err)
			return
		}
		decrypt := len(args) > 2 && args[2] == "tomoe"
		for key, value := range passwords {
			if decrypt {
				value = string(security.DecryptPassword([]byte(value), "tomoe"))
			}
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		}

	case "update":
		passwords, keys, err := read.ReadValues()
		if err != nil {
			fmt.Println(err)
			return
		}
		index, selectedKey := selectKey(passwords, keys)
		if index == -1 {
			return
		}
		newKey, newValue, err := update.UpdateKeyValuesInPrompt()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Update the map and keys slice
		delete(passwords, selectedKey) // Remove the old key-value pair
		passwords[newKey] =  string(security.EncryptPassword([]byte(newValue), "tomoe"))   // Add the new key-value pair
		keys[index] = newKey           // Update the key in the keys sliceu
		fmt.Println(keys)
		msg, err := update.UpdateKeyValuesInTextFile(passwords, keys)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	case "delete":
		passwords, keys, err := read.ReadValues()
		if err != nil {
			fmt.Println(err)
			return
		}
		index, selectedKey := selectKey(passwords, keys)
		if index == -1 {
			return
		}
		// Update the map and keys slice
		delete(passwords, selectedKey)
		// delete the key from keys using index
		keys = append(keys[:index], keys[index+1:]...)
		fmt.Println(keys)
		msg, err := update.UpdateKeyValuesInTextFile(passwords, keys)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)

	default:
		fmt.Println("Invalid command")
	}
}
