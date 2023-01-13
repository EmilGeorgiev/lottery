package cli

import (
	"strconv"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEnterLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enter-lottery [bet] [denom]",
		Short: "Broadcast message enterLottery",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBet, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argDenom := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnterLottery(
				clientCtx.GetFromAddress().String(),
				argBet,
				argDenom,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
