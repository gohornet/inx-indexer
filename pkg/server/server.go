package server

import (
	"github.com/labstack/echo/v4"

	"github.com/iotaledger/inx-indexer/pkg/indexer"
	iotago "github.com/iotaledger/iota.go/v4"
)

const (
	APIRoute = "/api/indexer/v2"
)

type IndexerServer struct {
	Indexer                 *indexer.Indexer
	Bech32HRP               iotago.NetworkPrefix
	RestAPILimitsMaxResults int
}

func NewIndexerServer(indexer *indexer.Indexer, echo *echo.Echo, prefix iotago.NetworkPrefix, maxPageSize int) *IndexerServer {
	s := &IndexerServer{
		Indexer:                 indexer,
		Bech32HRP:               prefix,
		RestAPILimitsMaxResults: maxPageSize,
	}
	s.configureRoutes(echo.Group(APIRoute))

	return s
}
