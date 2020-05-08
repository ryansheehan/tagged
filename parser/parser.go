package parser

import (
	"strconv"
	"strings"
)

// Parse parses the text property for hash-tagged strings
func Parse(text string) map[string]interface{} {
	const prefix = "#"
	const escape = "##"
	tags := make(map[string]interface{})
	words := strings.Fields(text)

	for _, word := range words {
		if !strings.HasPrefix(word, escape) && strings.HasPrefix(word, prefix) {
			kv := strings.SplitN(strings.TrimPrefix(word, prefix), ":", 2)
			key := kv[0]
			var strval string
			var value interface{}
			if len(kv) > 1 {
				strval = kv[1]
				if i, err := strconv.Atoi(strval); err != nil {
					if f, err := strconv.ParseFloat(strval, 64); err != nil {
						value = strval
					} else {
						value = f
					}
				} else {
					value = i
				}
			} else {
				key, value = kv[0], nil
			}
			tags[key] = value
		}
	}

	return tags
}
