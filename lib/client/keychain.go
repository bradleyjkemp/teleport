// +build !darwin

package client

type KeychainLocalKeyStore struct{LocalKeyStore}

func NewKeychainLocalKeyStore(publicStore LocalKeyStore) (s *KeychainLocalKeyStore, err error) {
	return nil, fmt.Errorf("Using the keychain to store keys is only supported on MacOS at the moment")
}
