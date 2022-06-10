package util

import "fmt"

func GetExtFile(contentType string) (string, error) {
    switch contentType {
    case "text/csv":
        return "csv", nil
    case "text/plain":
        return "txt", nil
    default:
        return "", fmt.Errorf("file is wrong format")
    }
}