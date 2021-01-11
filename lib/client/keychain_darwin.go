// +build darwin

package client

import (
	"crypto/x509"

	"github.com/gravitational/teleport/lib/auth"
	"golang.org/x/crypto/ssh"
)

type KeychainLocalKeyStore struct{}

func NewKeychainLocalKeyStore() (s *FSLocalKeyStore, err error) {
	return nil, nil
}

func (k KeychainLocalKeyStore) AddKey(proxy string, username string, key *Key) error {
	panic("implement me")
}

func (k KeychainLocalKeyStore) GetKey(proxy string, username string) (*Key, error) {
	panic("implement me")
}

func (k KeychainLocalKeyStore) DeleteKey(proxyHost string, username string) error {
	panic("implement me")
}

func (k KeychainLocalKeyStore) DeleteKeys() error {
	panic("implement me")
}

func (k KeychainLocalKeyStore) AddKnownHostKeys(hostname string, keys []ssh.PublicKey) error {
	panic("implement me")
}

func (k KeychainLocalKeyStore) GetKnownHostKeys(hostname string) ([]ssh.PublicKey, error) {
	panic("implement me")
}

func (k KeychainLocalKeyStore) SaveCerts(proxy string, cas []auth.TrustedCerts) error {
	panic("implement me")
}

func (k KeychainLocalKeyStore) GetCerts(proxy string) (*x509.CertPool, error) {
	panic("implement me")
}

func (k KeychainLocalKeyStore) GetCertsPEM(proxy string) ([]byte, error) {
	panic("implement me")
}

