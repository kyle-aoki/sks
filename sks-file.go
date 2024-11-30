package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

type Secret struct {
	Name  string
	Value string
	Notes string
}

func (sf *SksFile) secretNotFound() {
	panic("secret not found: " + *deleteFlag)
}

func (sf *SksFile) getNotes() string {
	return strings.Join(flag.Args(), " ")
}

func (s *Secret) format() string {
	return fmt.Sprintf("%s\t%s\t%s\n", s.Name, s.Value, s.Notes)
}

type SksFile struct {
	Secrets []*Secret
}

func (sf *SksFile) fields() string {
	return fmt.Sprintf("%s\t%s\t%s\n", "NAME", "VALUE", "NOTES")
}

func (sf *SksFile) find(name string) *Secret {
	for _, secret := range sf.Secrets {
		if name == secret.Name {
			return secret
		}
	}
	return nil
}

func (sf *SksFile) removeSecret(name string) {
	var newSecrets []*Secret
	for _, secret := range sks.Secrets {
		if secret.Name == name {
			continue
		}
		newSecrets = append(newSecrets, secret)
	}
	sf.Secrets = newSecrets
}

func (sf *SksFile) filter(value string) []*Secret {
	var filteredSecrets []*Secret
	for _, secret := range sf.Secrets {
		if strings.Contains(secret.Name, value) {
			filteredSecrets = append(filteredSecrets, secret)
		}
	}
	return filteredSecrets
}

var sks *SksFile = loadSksFile()

func sksFilePath() string {
	const SksFileName = ".sks"
	homedir := must(os.UserHomeDir())
	return path.Join(homedir, SksFileName)
}

func loadSksFile() *SksFile {
	if _, err := os.Stat(sksFilePath()); os.IsNotExist(err) {
		check(os.WriteFile(sksFilePath(), []byte("{}"), 0600))
	}
	return fromJson[*SksFile](sksFilePath())
}

func (sf *SksFile) save() {
	toJson(sksFilePath(), sks)
}
