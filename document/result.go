package document

type Response struct {
    Count int `json:"count,omitempty"`
    Hits  struct {
        Total struct {
            Value int `json:"value,omitempty"`
        } `json:"Total,omitempty"`
    } `json:"hits,omitempty"`
    Aggregations struct {
        CountBy struct {
            Buckets []struct {
                KeyAsString string `json:"key_as_string,omitempty"`
                Key         int    `json:"key,omitempty"`
                DocsCount   int    `json:"doc_count,omitempty"`
            } `json:"buckets,omitempty"`
        } `json:"count_by,omitempty"`
    } `json:"aggregations,omitempty"`
}

type TermStringBuilder struct {
    Term map[string]interface{} `json:"term,omitempty"`
}

type TermsStringBuilder struct {
    Terms map[string]interface{} `json:"terms,omitempty"`
}

type RangeStringBuilder struct {
    Range map[string]interface{} `json:"range,omitempty"`
}

type AggsCondition struct {
    CountBy struct {
        DateHistogram struct {
            Field            string `json:"field,omitempty"`
            CalendarInterval string `json:"calendar_interval,omitempty"`
        } `json:"date_histogram"`
    } `json:"count_by,omitempty"`
}

type QueryBuilder struct {
    Query struct {
        Bool struct {
            Must []interface{} `json:"must,omitempty"`
        } `json:"bool,omitempty"`
    } `json:"query,omitempty"`
    Aggs *AggsCondition `json:"aggs,omitempty"`
}