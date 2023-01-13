package cli

import (
	"context"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListFinishedLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-finished-lottery",
		Short: "list all finishedLottery",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFinishedLotteryRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.FinishedLotteryAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowFinishedLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-finished-lottery [index]",
		Short: "shows a finishedLottery",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetFinishedLotteryRequest{
				Index: argIndex,
			}

			res, err := queryClient.FinishedLottery(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
