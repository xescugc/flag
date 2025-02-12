package flag

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

type AuthorizedKeys struct {
	Keys []ssh.PublicKey
	File string
}

func (f *AuthorizedKeys) UnmarshalFlag(value string) error {
	authorizedKeysBytes, err := ioutil.ReadFile(value)
	if err != nil {
		return fmt.Errorf("failed to read authorized keys: %s", err)
	}

	var authorizedKeys []ssh.PublicKey

	for {
		key, _, _, rest, err := ssh.ParseAuthorizedKey(authorizedKeysBytes)
		if err != nil {
			// there's no good error to check for here
			break
		}

		authorizedKeys = append(authorizedKeys, key)

		authorizedKeysBytes = rest
	}

	f.Keys = authorizedKeys
	f.File = value

	return nil
}

// Reload reloads the value of the Keys
func (f *AuthorizedKeys) Reload() error {
	return f.UnmarshalFlag(f.File)
}
