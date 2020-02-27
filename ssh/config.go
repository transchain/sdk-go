package ssh

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

// Config is the configuration structure used by an ssh client.
type Config struct {
	User       string `yaml:"user"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	SshKeyPath string `yaml:"sshKeyPath"`
}

// GetSshClientConfig loads and parses a private key to return an ssh client config.
func (c Config) GetSshClientConfig() (*ssh.ClientConfig, error) {
	keyBuf, err := ioutil.ReadFile(c.SshKeyPath)
	if err != nil {
		return nil, err
	}
	key, err := ssh.ParsePrivateKey(keyBuf)
	return &ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}
