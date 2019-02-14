package ssmapi

import (
	"context"
	"fmt"
)

// SSMStore is the persistent store of ssm secure
type SSMStore interface {
	SSMGetByTXID(context.Context, string) ([]*Ssmitem, error)
	SSMPost(context.Context, *Ssmitem) (string, error)
	//ThingDeleteByID(context.Context, string) error
	//ThingFind(context.Context) ([]*ssm, error)
}

// Ssmitem is strcut for a secure memo
type Ssmitem struct {
	Txid              string `json:"txid" db:"txid"`
	Address           string `json:"address" db:"address"`
	EncryptedPayload  string `json:"encrypted_payload" db:"encrypted_payload"`
	EncryptionVersion string `json:"encryption_version" db:"encryption_version"`
}

// String is the stringer method
func (s *Ssmitem) String() string {
	return fmt.Sprintf(`{"txid":"%s","address":"%s","EncryptedPayload":"%s","EncryptionVersion":"%s"}`, s.Txid, s.Address, s.EncryptedPayload, s.EncryptionVersion)
}
