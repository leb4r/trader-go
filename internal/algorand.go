package internal

// see https://developer.algorand.org/solutions/exporting-algorand-transactions-tax-reporting/

import (
	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

// LoadAlgorandAccount loads an Algorand account into the database
func LoadAlgorandAccount(accountId, indexerAddress, token string) error {
	return nil
}

func newIndexerClient(algodAddress, token string) (*indexer.Client, error) {
	client, err := common.MakeClient(algodAddress, "X-API-KEY", token)
	if err != nil {
		return nil, err
	}
	return (*indexer.Client)(client), nil
}
