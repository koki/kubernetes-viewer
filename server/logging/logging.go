package logging

import (
	"github.com/sirupsen/logrus"
)

type Session struct {
	ID    *int    `json:"id,omitempty"`
	Login *string `json:"login,omitempty"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type Convert struct {
	URL         string `json:"url,omitempty"`
	Unconverted string `json:"unconverted,omitempty"`

	// error parsing the Convert request body.
	RequestError string `json:"request_error,omitempty"`
}

type Login struct {
}

type Log struct {
	Session *Session `json:"session,omitempty"`
	Convert *Convert `json:"convert,omitempty"`
	Login   *Login   `json:"login,omitempty"`
	Other   *string  `json:"other,omitempty"`
	Error   *string  `json:"error,omitempty"`
}

func IntPtrAt(obj map[interface{}]interface{}, key string) *int {
	if val, ok := obj[key]; ok {
		if intPtr, ok := val.(int); ok {
			return &intPtr
		}
	}

	return nil
}

func StringPtrAt(obj map[interface{}]interface{}, key string) *string {
	if val, ok := obj[key]; ok {
		if stringPtr, ok := val.(string); ok {
			return &stringPtr
		}
	}

	return nil
}

func MkSession(sessionValues map[interface{}]interface{}) *Session {
	if sessionValues == nil {
		return nil
	}

	return &Session{
		ID:    IntPtrAt(sessionValues, "id"),
		Login: StringPtrAt(sessionValues, "login"),
		Name:  StringPtrAt(sessionValues, "name"),
		Email: StringPtrAt(sessionValues, "email"),
	}
}

func MkConvert(url, unconverted, err string) *Convert {
	return &Convert{
		URL:          url,
		Unconverted:  unconverted,
		RequestError: err,
	}
}

func MkError(err error) *string {
	if err == nil {
		return nil
	}

	x := err.Error()
	return &x
}

func MkOtherLog(msg string) Log {
	return Log{
		Other: &msg,
	}
}

func WriteLog(log Log) {
	logrus.WithField("log", log).Info()
}
