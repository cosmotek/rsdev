package auth

import (
	"encoding/json"
	"errors"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

var ErrSessionNotFound = errors.New("session not found")

func checkAppFiles() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	return os.MkdirAll(path.Join(home, ".rsdev"), os.ModePerm)
}

func (h AuthHeader) Persist(host string) error {
	err := checkAppFiles()
	if err != nil {
		return err
	}

	sessions, err := readPersistenceFile()
	if err != nil {
		return err
	}

	sessions.Sessions[host] = h
	return sessions.writeFile()
}

func Retrieve(host string) (AuthHeader, error) {
	err := checkAppFiles()
	if err != nil {
		return AuthHeader{}, err
	}

	sessions, err := readPersistenceFile()
	if err != nil {
		return AuthHeader{}, err
	}

	session, ok := sessions.Sessions[host]
	if !ok {
		return AuthHeader{}, ErrSessionNotFound
	}

	return session, nil
}

type PersistedSessions struct {
	Sessions map[string]AuthHeader `json:"sessions"`
}

func readPersistenceFile() (PersistedSessions, error) {
	home, err := homedir.Dir()
	if err != nil {
		return PersistedSessions{}, err
	}

	file, err := os.Open(path.Join(home, ".rsdev", "sessions.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return PersistedSessions{
				Sessions: map[string]AuthHeader{},
			}, nil
		}

		return PersistedSessions{}, err
	}

	sessions := PersistedSessions{}
	err = json.NewDecoder(file).Decode(&sessions)
	if err != nil {
		return PersistedSessions{}, err
	}

	return sessions, nil
}

func (p PersistedSessions) writeFile() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	// ignore error
	os.Remove(path.Join(home, ".rsdev", "sessions.json"))

	file, err := os.Create(path.Join(home, ".rsdev", "sessions.json"))
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	return encoder.Encode(p)
}
