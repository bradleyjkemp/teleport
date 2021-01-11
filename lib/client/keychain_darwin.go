// +build darwin

package client

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/keybase/go-keychain"
)

// KeychainLocalKeyStore stores sensitive key material in the OS keychain and delegates
// non-sensitive/public material to the wrapped LocalKeyStore (usually the filesystem)
type KeychainLocalKeyStore struct{
	LocalKeyStore
}

func NewKeychainLocalKeyStore(publicStore LocalKeyStore) (s *KeychainLocalKeyStore, err error) {
	fmt.Fprintln(os.Stderr, "Using the keychain :)") // TODO: removeme
	return &KeychainLocalKeyStore{
		publicStore,
	}, nil
}

func (k KeychainLocalKeyStore) AddKey(proxy string, username string, key *Key) error {
	data, err := json.Marshal(key)
	if err != nil {
		return fmt.Errorf("failed marshalling key: %w", err)
	}
	return keychain.AddItem(keychain.NewGenericPassword("tsh:"+proxy, username, "", data, ""))
}

func (k KeychainLocalKeyStore) GetKey(proxy string, username string) (*Key, error) {
	data, err := keychain.GetGenericPassword("tsh:"+proxy, username, "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get password from keychain: %w", err)
	}

	var key *Key
	if err := json.Unmarshal(data, key); err != nil {
		return nil, fmt.Errorf("failed to unmarshal key: %w", err)
	}
	return key, nil
}

func (k KeychainLocalKeyStore) DeleteKey(proxy string, username string) error {
	return keychain.DeleteItem(keychain.NewGenericPassword("tsh:"+proxy, username, "", nil, ""))
}

func (k KeychainLocalKeyStore) DeleteKeys() error {
	panic("implement me")
}
