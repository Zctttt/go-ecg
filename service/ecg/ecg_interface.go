package ecg

import "go-sv/models"

type Heart interface {
	Connect() [4]byte
	UserList() []byte
	Set([]byte) bool
	Get(string, string) (*models.HeartData, error)
}
