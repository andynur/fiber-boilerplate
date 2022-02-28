package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
	"unicode"
)

func ZeroPadNumber(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}

func TimestampString() string {
	now := time.Now()
	return fmt.Sprintf("%d%02d%02d%02d%02d%02d%02d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
}

func CurrentTimestampFormatted() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func RemoveAllSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func Retry(attempts int, sleep time.Duration, f func() error) error {
	type stop struct {
		error
	}

	if err := f(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, f)
		}
		return err
	}

	return nil
}

func JsonStringFormatted(jsonString string) interface{} {
	type formatMessage struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}

	var messages map[string]string
	json.Unmarshal([]byte(jsonString), &messages)

	rawAttributes := []formatMessage{}
	for key, val := range messages {
		capitalizeKey := strings.Title(strings.ReplaceAll(key, "_", " "))
		rawAttributes = append(rawAttributes, formatMessage{Label: capitalizeKey, Value: val})
	}

	return &rawAttributes
}

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				exists = true
				return
			}
		}
	}

	return
}
