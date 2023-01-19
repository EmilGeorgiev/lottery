package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/EmilGeorgiev/lottery/testutil/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery"
	"github.com/EmilGeorgiev/lottery/x/lottery/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery/testutil"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupEnterLotteryTest(t testing.TB) (keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper, types.Lottery) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	accMock := testutil.NewMockAccountKeeper(ctrl)
	k, ctx := keepertest.LotteryKeeperWithMock(t, bankMock, accMock)
	lottery.InitGenesis(ctx, *k, *types.DefaultGenesis())

	l := types.Lottery{EnterLotteryTxs: []*types.EnterLotteryTx{
		{UserAddress: client1, Bet: 1, Denom: "token", Datetime: "2023-01-16 19:24:42.743045679 +0000 UTC"},
		{UserAddress: client2, Bet: 2, Denom: "token", Datetime: "2023-01-16 19:25:42.743045679 +0000 UTC"},
		{UserAddress: client3, Bet: 3, Denom: "token", Datetime: "2023-01-16 19:26:42.743045679 +0000 UTC"},
		{UserAddress: client4, Bet: 4, Denom: "token", Datetime: "2023-01-16 19:27:42.743045679 +0000 UTC"},
		{UserAddress: client5, Bet: 5, Denom: "token", Datetime: "2023-01-16 19:28:42.743045679 +0000 UTC"},
	}}

	return *k, sdk.WrapSDKContext(ctx), ctrl, bankMock, l
}

func TestEnterLottery(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bankMock, _ := setupEnterLotteryTest(t)
	defer ctrl.Finish()
	coin := sdk.NewCoin("token", sdk.NewInt(int64(10)))
	sender, _ := types.GetAddress(alice)
	bankMock.EXPECT().
		GetBalance(gomock.Any(), sender, "token").
		Return(sdk.Coin{Denom: "token", Amount: sdk.NewInt(500)})
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(ctx), sender, types.ModuleName, sdk.NewCoins(coin)).
		Return(nil)
	el := &types.MsgEnterLottery{
		Creator: alice,
		Bet:     10,
		Denom:   "token",
	}
	msgSrv := keeper.NewMsgServerImpl(k)

	// Action
	actual, err := msgSrv.EnterLottery(ctx, el)

	// Assert
	expected := &types.MsgEnterLotteryResponse{RegisteredUsers: 1}
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
