// +build !darwin

package client

import (
	"crypto/x509"

	"github.com/gravitational/teleport/lib/auth"
	"golang.org/x/crypto/ssh"
)

type KeychainLocalKeyStore struct{}

func NewKeychainLocalKeyStore() (s *KeychainLocalKeyStore, err error) {
	return nil, fmt.Errorf("Using the keychain to store keys is only supported on MacOS at the moment")
}

func (k KeychainLocalKeyStore) AddKey(proxy string, username string, key *Key) error {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) GetKey(proxy string, username string) (*Key, error) {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) DeleteKey(proxyHost string, username string) error {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) DeleteKeys() error {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) AddKnownHostKeys(hostname string, keys []ssh.PublicKey) error {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) GetKnownHostKeys(hostname string) ([]ssh.PublicKey, error) {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) SaveCerts(proxy string, cas []auth.TrustedCerts) error {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) GetCerts(proxy string) (*x509.CertPool, error) {
	panic("unimplemented")
}

func (k KeychainLocalKeyStore) GetCertsPEM(proxy string) ([]byte, error) {
	panic("unimplemented")
}

