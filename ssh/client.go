package ssh

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "path/filepath"

    "github.com/pkg/sftp"
    "golang.org/x/crypto/ssh"
)

// RemoteAccess interface defines the methods a concrete remote accessor must implement.
type RemoteAccess interface {
    SendFile(sourcePath string, destPath string) error
    ExecCmd(connectPath string, cmd string) error
}

// Client is an ssh wrapper to dialog with a sftp server.
type Client struct {
    SshClientConfig *ssh.ClientConfig
    SshUrl          string
}

// Client constructor.
func NewClient(cfg *Config) (*Client, error) {
    sshUrl := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
    sshConfig, err := cfg.GetSshClientConfig()
    if err != nil {
        return nil, err
    }
    return &Client{
        SshUrl:          sshUrl,
        SshClientConfig: sshConfig,
    }, nil
}

// SendFile sends a file from a source path to the destination directory with the same name.
func (c *Client) SendFile(sourceFilePath string, destDirPath string) error {
    client, err := ssh.Dial("tcp", c.SshUrl, c.SshClientConfig)
    if err != nil {
        return err
    }
    defer func() {
        _ = client.Close()
    }()

    sftpCli, err := sftp.NewClient(client)
    if err != nil {
        return err
    }
    defer func() {
        _ = sftpCli.Close()
    }()

    // Open the source file to copy
    in, err := os.Open(sourceFilePath)
    if err != nil {
        return err
    }
    defer func() {
        _ = in.Close()
    }()

    // Create the destination folders
    if err := sftpCli.MkdirAll(destDirPath); err != nil {
        return err
    }

    _, fileName := filepath.Split(sourceFilePath)

    // Open the destination file
    out, err := sftpCli.Create(filepath.Join(destDirPath, fileName))
    if err != nil {
        return err
    }
    defer func() {
        _ = out.Close()
    }()

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }

    return nil
}

// ExecCmd will remotely execute a command.
func (c *Client) ExecCmd(connectPath string, cmd string) error {
    client, err := ssh.Dial("tcp", c.SshUrl, c.SshClientConfig)
    if err != nil {
        return err
    }
    defer func() {
        _ = client.Close()
    }()

    session, err := client.NewSession()
    if err != nil {
        return err
    }
    defer func() {
        _ = session.Close()
    }()

    var stdErrBuf bytes.Buffer
    session.Stderr = &stdErrBuf

    cmd = fmt.Sprintf("cd %s && %s", connectPath, cmd)
    err = session.Run(cmd)
    if err != nil {
        return fmt.Errorf("cmd [%s] returned [%s] with message [%s]", cmd, err, stdErrBuf.String())
    }

    return nil
}
