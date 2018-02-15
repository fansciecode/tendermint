package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/types"
)

// ShowNodeIDCmd dumps node's ID to the standard output.
var ShowNodeIDCmd = &cobra.Command{
	Use:   "show_node_id",
	Short: "Show this node's ID",
	Run:   showNodeID,
}

func showNodeID(cmd *cobra.Command, args []string) {
	privValidator := types.LoadOrGenPrivValidatorFS(config.PrivValidatorFile())
	fmt.Println(p2p.PubKeyToID(privValidator.PubKey))
}
