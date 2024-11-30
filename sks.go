package main

import (
	"flag"
)

const (
	SAVE   = "s"
	FILTER = "f"
	DELETE = "d"
	URL    = "h"
	VALUE  = "v"

	ROTATE = "r"

	LENGTH = "l"
)

var (
	saveFlag   = flag.String(SAVE, "", "")
	filterFlag = flag.String(FILTER, "", "")
	deleteFlag = flag.String(DELETE, "", "")
	valueFlag  = flag.String(VALUE, "", "")
	rotateFlag = flag.String(ROTATE, "", "")

	lengthFlag = flag.Int(LENGTH, 30, "")
)

var args []string

func main() {
	defer panicRecover()
	flag.Parse()

	exec(*saveFlag != "", SAVE_PASSWORD)
	exec(*deleteFlag != "", DELETE_PASSWORD)
	exec(*rotateFlag != "", ROTATE_SECRET)

	exec(true, LIST_PASSWORDS)
}

func SAVE_PASSWORD() {
	secret := sks.find(*saveFlag)
	if secret == nil {
		secret = &Secret{Name: *saveFlag, Value: generatePassword(*lengthFlag)}
		sks.Secrets = append(sks.Secrets, secret)
	}
	if *valueFlag != "" {
		secret.Value = *valueFlag
	}
	if sks.getNotes() != "" {
		secret.Notes = sks.getNotes()
	}
	sks.save()
}

func DELETE_PASSWORD() {
	secret := sks.find(*deleteFlag)
	if secret == nil {
		sks.secretNotFound()
	}
	sks.removeSecret(*deleteFlag)
	sks.save()
}

func ROTATE_SECRET() {
	secret := sks.find(*rotateFlag)
	if secret == nil {
		sks.secretNotFound()
	}
	secret.Value = generatePassword(*lengthFlag)
	sks.save()
}

func LIST_PASSWORDS() {
	var filteredSecrets []*Secret = sks.Secrets
	if *filterFlag != "" {
		filteredSecrets = sks.filter(*filterFlag)
	}
	tabprint(sks.fields())
	for _, secret := range filteredSecrets {
		tabprint(secret.format())
	}
	tabflush()
}
