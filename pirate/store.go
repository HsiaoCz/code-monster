package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

func getOrSetKey() error {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return errors.New("error opening database")
	}
	defer db.Close()

	// Create the bucket if it doesnt exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return errors.New("error creating database bucket")
		}
		return nil
	})
	if err != nil {
		return err
	}

	var apiKey string
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("MyBucket"))
		bytes := bucket.Get([]byte("OPENAI_API_KEY"))
		apiKey = string(bytes)
		return nil
	})

	if apiKey == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your OpenAl API Key")
		input, err := reader.ReadString('\n')
		if err != nil {
			return errors.New("error reading user input")
		}
		apiKey = strings.TrimSpace(input)
		if err := db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("MyBucket"))
			if err := bucket.Put([]byte("OPENAI_API_KEY"), []byte(apiKey)); err != nil {
				return errors.New("error adding key to database")
			}
			return nil
		}); err != nil {
			return err
		}
	}

	if err := os.Setenv("OPENAI_API_KEY", apiKey); err != nil {
		return errors.New("error setting enviromentn variable")
	}

	return nil
}
