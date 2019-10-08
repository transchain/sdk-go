package sftp

// Client interface defines the methods a concrete client must implement.
type Client interface {
    SendFile(sourcePath string, destPathSuffix string) error
}
