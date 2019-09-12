package api

// Client interface defines the methods a concrete client must implement.
type Client interface {
    Get(route string, headers map[string]string, queryValues map[string]string) (*RawResponse, error)
    Post(route string, body []byte, headers map[string]string, queryValues map[string]string) (*RawResponse, error)
    AddHeader(key string, value string)
    RemoveHeader(key string)
}
