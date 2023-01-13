package keeper_test

import (
	"context"
	keepertest "github.com/EmilGeorgiev/lottery/testutil/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery"
	"github.com/EmilGeorgiev/lottery/x/lottery/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.LotteryKeeperWithMock(t, bankMock)
	lottery.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	return server, *k, context, ctrl, bankMock
}
