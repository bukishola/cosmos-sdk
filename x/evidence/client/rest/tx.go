package rest

import (
	"github.com/kava-labs/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, handlers []EvidenceRESTHandler) {
	// TODO: Register tx handlers.
}
