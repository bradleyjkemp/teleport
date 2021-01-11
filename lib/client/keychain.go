// +build !darwin

package client

import "github.com/gravitational/trace"

type KeychainLocalKeyStore struct{ LocalKeyStore }

func NewKeychainLocalKeyStore(publicStore LocalKeyStore) (s *KeychainLocalKeyStore, err error) {
	return nil, trace.Errorf("Using the keychain to store keys is only supported on MacOS at the moment")
}
