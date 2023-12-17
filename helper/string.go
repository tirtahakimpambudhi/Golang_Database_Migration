package helper

import (
	"fmt"
	"regexp"
	"strings"
)

// InQueryPlaceholders mengganti tanda tanya dalam query dengan placeholders sesuai dengan panjang slice.
func InQueryPlaceholders(query string, length int) string {
	var qword string
	for i := 0; i < length; i++ {
		qword += "?,"
	}
	if strings.HasSuffix(qword, ",") {
		qword = qword[:len(qword)-1]
	}
	query = strings.Replace(query, "(?)", fmt.Sprintf("(%v)", qword), -1)
	return query
}

func RemoveComa(coloumn string) string {
	if strings.HasSuffix(coloumn, ",") {
		coloumn = coloumn[:len(coloumn)-1]
	}
	return coloumn
}
// SliceToInterface mengonversi slice menjadi slice dari interface{}.
func SliceToInterface(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func RemoveLimitOffset(query string) string {
	re := regexp.MustCompile(`SELECT(.*)FROM(.*)LIMIT(.*)OFFSET(.*)`)
	query = re.ReplaceAllString(query, "SELECT COUNT(*) FROM$2")

	return query
}
