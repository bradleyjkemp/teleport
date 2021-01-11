// +build darwin

package client

import (
	"encoding/json"
	"fmt"

	"github.com/gravitational/trace"
	"github.com/keybase/go-keychain"
)

// KeychainLocalKeyStore stores sensitive key material in the OS keychain and delegates
// non-sensitive/public material to the wrapped LocalKeyStore (usually the filesystem)
type KeychainLocalKeyStore struct {
	*FSLocalKeyStore
}

func NewKeychainLocalKeyStore(fsKeystore *FSLocalKeyStore) (s *KeychainLocalKeyStore) {
	return &KeychainLocalKeyStore{
		fsKeystore,
	}
}

func (k KeychainLocalKeyStore) AddKey(proxy string, username string, key *Key) error {
	data, err := json.Marshal(key)
	if err != nil {
		return fmt.Errorf("failed marshalling key: %w", err)
	}
	// always delete the key (if it exists)
	_ = k.DeleteKey(proxy, username)
	return keychain.AddItem(keychain.NewGenericPassword("teleport", username, proxy, data, ""))
}

func (k KeychainLocalKeyStore) GetKey(proxy string, username string, opts ...KeyOption) (*Key, error) {
	data, err := keychain.GetGenericPassword("teleport", username, proxy, "")
	if err != nil {
		return nil, trace.Errorf("failed to get password from keychain: %w", err)
	}
	if len(data) == 0 {
		return nil, trace.NotFound("no session keys for %v in %v", username, proxy)
	}

	var key Key
	if err := json.Unmarshal(data, &key); err != nil {
		return nil, fmt.Errorf("failed to unmarshal key: %w", err)
	}

	for _, o := range opts {
		if err := o.getKey(k.KeyDir, username, &key); err != nil {
			k.log.Error(err)
			return nil, trace.Wrap(err)
		}
	}
	return &key, nil
}

func (k KeychainLocalKeyStore) DeleteKey(proxy string, username string, opts ...KeyOption) error {
	err := keychain.DeleteItem(keychain.NewGenericPassword("teleport", username, proxy, nil, ""))
	if err != nil && err != keychain.ErrorItemNotFound {
		return err
	}
	for _, o := range opts {
		if err := o.deleteKey(k.KeyDir, username); err != nil {
			return trace.Wrap(err)
		}
	}
	return nil
}

func (k KeychainLocalKeyStore) DeleteKeys() error {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService("teleport")
	query.SetMatchLimit(keychain.MatchLimitAll)
	query.SetReturnData(false)

	items, err := keychain.QueryItem(query)
	if err != nil {
		return trace.Errorf("failed to find keychain items: %w", err)
	}
	for _, item := range items {
		if err := k.DeleteKey(item.Label, item.Account); err != nil {
			return trace.Errorf("failed to delete item: %w", err)
		}
	}
	return nil
}
