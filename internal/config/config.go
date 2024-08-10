package config

import "github.com/joho/godotenv"

// Load loads the environment variables from the file
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
