package mira

import (
	"io/ioutil"
	"os"
	"regexp"
)

func ReadCredsFromFile(file string) Credentials {
	// Declare all regexes
	ClientId, _ := regexp.Compile(`CLIENT_ID\s*=\s*(.+)`)
	ClientSecret, _ := regexp.Compile(`CLIENT_SECRET\s*=\s*(.+)`)
	Username, _ := regexp.Compile(`USERNAME\s*=\s*(.+)`)
	Password, _ := regexp.Compile(`PASSWORD\s*=\s*(.+)`)
	UserAgent, _ := regexp.Compile(`USER_AGENT\s*=\s*(.+)`)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Credentials{}
	}
	s := string(data)
	creds := Credentials{
		ClientId.FindStringSubmatch(s)[1],
		ClientSecret.FindStringSubmatch(s)[1],
		Username.FindStringSubmatch(s)[1],
		Password.FindStringSubmatch(s)[1],
		UserAgent.FindStringSubmatch(s)[1],
	}
	return creds
}

// Assuming that they all exist. Probably a bad idea. We can
// expand it later and do a more aggressive error handling.
func ReadCredsFromEnv() Credentials {
	var clientID, clientSecret, username, password, userAgent string
	var exists bool

	clientID, exists = os.LookupEnv("CLIENT_ID")
	if !exists {
		clientID = os.Getenv("BOT_CLIENT_ID")
	}
	clientSecret, exists = os.LookupEnv("CLIENT_SECRET")
	if !exists {
		clientSecret = os.Getenv("BOT_CLIENT_SECRET")
	}
	username, exists = os.LookupEnv("USERNAME")
	if !exists {
		username = os.Getenv("BOT_USERNAME")
	}
	password, exists = os.LookupEnv("PASSWORD")
	if !exists {
		password = os.Getenv("BOT_PASSWORD")
	}
	userAgent, exists = os.LookupEnv("USER_AGENT")
	if !exists {
		userAgent = os.Getenv("BOT_USER_AGENT")
	}

	return Credentials{
		clientID,
		clientSecret,
		username,
		password,
		userAgent,
	}
}
