package opensearch

import (
    "golib/document"
    "io"
)

func (o *OpenSearchClient) CountDocument(index []string, bodyQuery io.Reader) (document.Response, error) {
    return o.clients.CountDocument(index, bodyQuery)
}
