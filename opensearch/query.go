package opensearch

import (
    "github.com/Kephas73/go-lib/document"
    "io"
)

func (o *OpenSearchClient) CountDocument(index []string, bodyQuery io.Reader) (document.Response, error) {
    return o.clients.CountDocument(index, bodyQuery)
}

func (o *OpenSearchClient) SearchDocument(index []string, bodyQuery io.Reader) (document.Response, error) {
    return o.clients.SearchDocument(index, bodyQuery)
}
