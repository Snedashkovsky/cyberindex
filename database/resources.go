package database

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dbtypes "github.com/forbole/bdjuno/database/types"
)

func (db *CyberDb) SaveInvestmints(
	agent string,
	amount sdk.Coin,
	resource string,
	length uint64,
	timestamp string,
	height int64,
	txHash string,
) error {
	query := `
		INSERT INTO investmints (agent, amount, resource, length, timestamp, height, transaction_hash)
		VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT DO NOTHING
	`

	coin := dbtypes.NewDbCoin(amount)
	value, err := coin.Value()
	if err != nil {
		return err
	}

	_, err = db.Sql.Exec(query,
		agent,
		value,
		resource,
		length,
		timestamp,
		height,
		txHash,
	)
	if err != nil {
		return err
	}

	return nil
}
