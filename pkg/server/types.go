package server

import iotago "github.com/iotaledger/iota.go/v4"

// outputsResponse defines the response of a GET outputs REST API call.
type outputsResponse struct {
	// The ledger index at which these outputs where available at.
	LedgerIndex iotago.SlotIndex `json:"ledgerIndex"`
	// The maximum count of results that are returned by the node.
	PageSize uint32 `json:"pageSize"`
	// The cursor to use for getting the next results.
	Cursor *string `json:"cursor,omitempty"`
	// The output IDs (transaction hash + output index) of the outputs on this address.
	Items []string `json:"items"`
}
