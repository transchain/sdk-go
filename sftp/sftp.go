package sftp

import (
    "fmt"
    "io"
    "os"
    "path/filepath"

    "github.com/pkg/sftp"
    "golang.org/x/crypto/ssh"

    tcSsh "github.com/transchain/sdk-go/ssh"
)

// PkgSftpClient is a sftp.Client wrapper to dialog with a sftp server.
type PkgSftpClient struct {
    SshClientConfig *ssh.ClientConfig
    SshUrl          string
    BasePath        string
}

// PkgSftpClient constructor.
func NewPkgSftpClient(cfg *tcSsh.Config) (*PkgSftpClient, error) {
    sshUrl := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
    sshConfig, err := cfg.GetSshClientConfig()
    if err != nil {
        return nil, err
    }
    return &PkgSftpClient{
        BasePath:        cfg.BasePath,
        SshUrl:          sshUrl,
        SshClientConfig: sshConfig,
    }, nil
}

// SendFile sends a file from a source path to the concatenated base path and dest path suffix.
func (sp *PkgSftpClient) SendFile(sourcePath string, destPathSuffix string) error {
    client, err := ssh.Dial("tcp", sp.SshUrl, sp.SshClientConfig)
    defer func() {
        _ = client.Close()
    }()
    if err != nil {
        return err
    }

    sftpCli, err := sftp.NewClient(client)
    defer func() {
        _ = sftpCli.Close()
    }()
    if err != nil {
        return err
    }

    // Open the source file to copy
    in, err := os.Open(sourcePath)
    defer func() {
        _ = in.Close()
    }()
    if err != nil {
        return err
    }

    fullDestPath := filepath.Join(sp.BasePath, destPathSuffix)

    // Create the destination folders
    if err := sftpCli.MkdirAll(fullDestPath); err != nil {
        return err
    }

    _, fileName := filepath.Split(sourcePath)

    // Open the destination file
    out, err := sftpCli.OpenFile(filepath.Join(fullDestPath, fileName), os.O_CREATE|os.O_WRONLY)
    defer func() {
        _ = out.Close()
    }()
    if err != nil {
        return err
    }

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }

    return nil
}
