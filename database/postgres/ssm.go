package postgres

import (
	"context"
	"database/sql"

	"github.com/CryptoLarry/ssm/ssmapi"
)

// SSMGetByTXID returns all the ssm accosiated with a TXID
func (c *Client) SSMGetByTXID(ctx context.Context, id string) ([]*ssmapi.Ssmitem, error) {

	var bs = make([]*ssmapi.Ssmitem, 0)
	err := c.db.SelectContext(ctx, &bs, `SELECT * FROM ssm WHERE txid = $1`, id)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return bs, nil

}

// SSMPost saves the SSm to DB
func (c *Client) SSMPost(ctx context.Context, i *ssmapi.Ssmitem) (string, error) {

	// Generate an ID if needed
	b := new(ssmapi.Ssmitem)
	saved := "Saved"
	err := c.db.GetContext(ctx, b, `SELECT * FROM ssm WHERE txid = $1 AND address = $2`, i.Txid, i.Address)
	if err != sql.ErrNoRows {
		return saved, ErrAlreadyExists
	} else if err != sql.ErrNoRows && err != nil {
		return saved, err
	}

	_, err = c.db.ExecContext(ctx, `
		INSERT INTO ssm (txid, address, encrypted_payload, encryption_version)
		VALUES($1, $2, $3,$4)
	`, i.Txid, i.Address, i.EncryptedPayload, i.EncryptionVersion)
	if err != nil {
		return saved, err
	}
	return saved, nil

}
