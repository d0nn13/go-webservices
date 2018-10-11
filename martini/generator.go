package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/ldap.v2"
)

const (
	ldapHost = "localhost"
	ldapPort = 389
)

type input struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Challenge string `json:"challenge"`
}

// Generate stuff
func Generate(body io.ReadCloser) ([]byte, error) {
	rawInput, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var input input
	err = json.Unmarshal(rawInput, &input)
	if err != nil {
		return nil, err
	}

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	if err != nil {
		return nil, err
	}
	defer l.Close()

	dn := fmt.Sprintf("cn=%s,dc=example,dc=com", input.Login)
	err = l.Bind(dn, input.Password)
	if err != nil {
		return nil, err
	}

	return []byte{'o', 'k'}, nil
}
