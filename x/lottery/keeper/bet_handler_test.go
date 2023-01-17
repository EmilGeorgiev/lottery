package keeper_test

import (
	"context"
	"errors"
	"testing"

	keepertest "github.com/EmilGeorgiev/lottery/testutil/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery"
	"github.com/EmilGeorgiev/lottery/x/lottery/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery/testutil"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	client1 = "cosmos1lp5wt7h98ddkwytslqctw5mh78happraje998p"
	client2 = "cosmos1p3jntzl6wthpn2glf6ztvvkdcmnd2ha9xwh4wd"
	client3 = "cosmos19pzj5xkx96ax6ts380elwn05zuj69gz0ygeh2x"
	client4 = "cosmos1qaqnkc44f2q77z9r5j4mfvkxa3jm7rsn3fd7rt"
	client5 = "cosmos1458527uwucwsttdeyyafy5yt8jqj6vj0r2x2l9"
)

func setupTest(t testing.TB) (keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper, types.Lottery) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.LotteryKeeperWithMock(t, bankMock)
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

func TestPayReward_UserPlacedTheHighestBet(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bank, l := setupTest(t)
	defer ctrl.Finish()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	si := &types.SystemInfo{LotteryPool: 50}
	winnerAddress, _ := types.GetAddress(client5)
	entirePool := si.LotteryPool + l.GetSumOfAllBetsPlusFees()
	coin := sdk.NewCoin("token", sdk.NewInt(int64(entirePool)))
	bank.EXPECT().SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, winnerAddress, sdk.NewCoins(coin)).Return(nil)
	winnerIndex := int64(4)

	// Action
	actual := k.PayReward(sdkCtx, winnerIndex, si, l)

	// Assert
	expected := &types.SystemInfo{LotteryPool: 0}
	require.Equal(t, entirePool, actual)
	require.Equal(t, expected, si)
}

func TestPayReward_UserPlacedTheHighestBetButPaymentFailed(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bank, l := setupTest(t)
	defer ctrl.Finish()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	si := &types.SystemInfo{LotteryPool: 50}
	winnerAddress, _ := types.GetAddress(client5)
	entirePool := si.LotteryPool + l.GetSumOfAllBetsPlusFees()
	coin := sdk.NewCoin("token", sdk.NewInt(int64(entirePool)))
	bank.EXPECT().
		SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, winnerAddress, sdk.NewCoins(coin)).
		Return(errors.New("Oops"))
	winnerIndex := int64(4)

	// Assert
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "cannot pay reward to the winner: Oops", r)
	}()

	// Action
	_ = k.PayReward(sdkCtx, winnerIndex, si, l)
}

func TestPayReward_UserPlacedTheLowestBet(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bank, l := setupTest(t)
	defer ctrl.Finish()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	si := &types.SystemInfo{LotteryPool: 40}
	fees := len(l.EnterLotteryTxs) * types.TxFee
	coin := sdk.NewCoin("token", sdk.NewInt(int64(fees)))
	bank.EXPECT().SendCoinsFromModuleToModule(sdkCtx, types.ModuleName, types.FeeCollectorName, sdk.NewCoins(coin)).Return(nil)
	winnerIndex := int64(0)

	// Action
	actual := k.PayReward(sdkCtx, winnerIndex, si, l)

	// Assert
	expected := &types.SystemInfo{LotteryPool: 40 + l.GetSumOfAllBets()}
	require.Equal(t, uint64(0), actual)
	require.Equal(t, expected, si)
}

func TestPayReward_SendCoinsFromLotteryModuleToFeeCollectorFailed(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bank, l := setupTest(t)
	defer ctrl.Finish()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	si := &types.SystemInfo{LotteryPool: 40}
	fees := len(l.EnterLotteryTxs) * types.TxFee
	coin := sdk.NewCoin("token", sdk.NewInt(int64(fees)))
	bank.EXPECT().
		SendCoinsFromModuleToModule(sdkCtx, types.ModuleName, types.FeeCollectorName, sdk.NewCoins(coin)).
		Return(errors.New("Oops"))
	winnerIndex := int64(0)

	// Assert
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "cannot send fees to the fee collector: Oops", r)
	}()

	// Action
	k.PayReward(sdkCtx, winnerIndex, si, l)
}

func TestPayReward_UserPlacedNeitherLowestOrHighestBet(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bank, l := setupTest(t)
	defer ctrl.Finish()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	si := &types.SystemInfo{LotteryPool: 40}
	fees := len(l.EnterLotteryTxs) * types.TxFee
	coinFees := sdk.NewCoin("token", sdk.NewInt(int64(fees)))
	bank.EXPECT().
		SendCoinsFromModuleToModule(sdkCtx, types.ModuleName, types.FeeCollectorName, sdk.NewCoins(coinFees)).
		Return(nil)
	coin := sdk.NewCoin("token", sdk.NewInt(int64(l.GetSumOfAllBets())))
	winnerAddress, _ := types.GetAddress(client4)
	bank.EXPECT().
		SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, winnerAddress, sdk.NewCoins(coin)).
		Return(nil)
	winnerIndex := int64(3)

	// Action
	actual := k.PayReward(sdkCtx, winnerIndex, si, l)

	// Assert
	expected := &types.SystemInfo{LotteryPool: 40}
	require.Equal(t, uint64(15), actual)
	require.Equal(t, expected, si)
}

func TestPayReward_AllUsersPlacedTheSameBet(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bank, _ := setupTest(t)
	defer ctrl.Finish()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	si := &types.SystemInfo{LotteryPool: 40}
	l := types.Lottery{EnterLotteryTxs: []*types.EnterLotteryTx{
		{UserAddress: client1, Bet: 8, Denom: "token", Datetime: "2023-01-16 19:24:42.743045679 +0000 UTC"},
		{UserAddress: client2, Bet: 8, Denom: "token", Datetime: "2023-01-16 19:25:42.743045679 +0000 UTC"},
		{UserAddress: client3, Bet: 8, Denom: "token", Datetime: "2023-01-16 19:26:42.743045679 +0000 UTC"},
		{UserAddress: client4, Bet: 8, Denom: "token", Datetime: "2023-01-16 19:27:42.743045679 +0000 UTC"},
		{UserAddress: client5, Bet: 8, Denom: "token", Datetime: "2023-01-16 19:28:42.743045679 +0000 UTC"},
	}}
	fees := len(l.EnterLotteryTxs) * types.TxFee
	coinFees := sdk.NewCoin("token", sdk.NewInt(int64(fees)))
	bank.EXPECT().
		SendCoinsFromModuleToModule(sdkCtx, types.ModuleName, types.FeeCollectorName, sdk.NewCoins(coinFees)).
		Return(nil)
	coin := sdk.NewCoin("token", sdk.NewInt(int64(l.GetSumOfAllBets())))
	winnerAddress, _ := types.GetAddress(client4)
	bank.EXPECT().
		SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, winnerAddress, sdk.NewCoins(coin)).
		Return(nil)
	winnerIndex := int64(3)

	// Action
	actual := k.PayReward(sdkCtx, winnerIndex, si, l)

	// Assert
	expected := &types.SystemInfo{LotteryPool: 40}
	require.Equal(t, uint64(40), actual)
	require.Equal(t, expected, si)
}

func TestPayReward_UserPlacedNeitherLowestOrHighestBetButPayRewardFailed(t *testing.T) {
	// SetUp
	k, ctx, ctrl, bank, l := setupTest(t)
	defer ctrl.Finish()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	si := &types.SystemInfo{LotteryPool: 40}
	fees := len(l.EnterLotteryTxs) * types.TxFee
	coinFees := sdk.NewCoin("token", sdk.NewInt(int64(fees)))
	bank.EXPECT().
		SendCoinsFromModuleToModule(sdkCtx, types.ModuleName, types.FeeCollectorName, sdk.NewCoins(coinFees)).
		Return(nil)
	coin := sdk.NewCoin("token", sdk.NewInt(int64(l.GetSumOfAllBets())))
	winnerAddress, _ := types.GetAddress(client4)
	bank.EXPECT().
		SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, winnerAddress, sdk.NewCoins(coin)).
		Return(errors.New("oops"))
	winnerIndex := int64(3)

	// Assert
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "cannot pay reward to the winner: oops", r)
	}()

	// Action
	k.PayReward(sdkCtx, winnerIndex, si, l)
}

func TestCollectBet(t *testing.T) {
	// SetUp
	k, ctx, _, bank, _ := setupTest(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	el := types.MsgEnterLottery{
		Creator: client1,
		Bet:     5,
		Denom:   "token",
	}
	addr, _ := el.GetAddress()
	bet := sdk.NewCoin("token", sdk.NewInt(5))
	bank.EXPECT().
		SendCoinsFromAccountToModule(sdkCtx, addr, types.ModuleName, sdk.NewCoins(bet)).
		Return(nil)

	// Action
	err := k.CollectBet(sdkCtx, el)

	// Assert
	require.NoError(t, err)
}

func TestCollectBet_WhenSenderAddressIsInvalid(t *testing.T) {
	// SetUp
	k, ctx, _, _, _ := setupTest(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	el := types.MsgEnterLottery{
		Creator: "invalid",
	}

	// Action
	err := k.CollectBet(sdkCtx, el)

	// Assert
	expected := "user address is invalid: invalid: decoding bech32 failed: invalid bech32 string length 7"
	require.Equal(t, expected, err.Error())
}

func TestCollectBet_WhenSendCoinsFromAccountToModuleFailed(t *testing.T) {
	// SetUp
	k, ctx, _, bank, _ := setupTest(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	el := types.MsgEnterLottery{
		Creator: client1,
		Bet:     5,
		Denom:   "token",
	}
	addr, _ := el.GetAddress()
	bet := sdk.NewCoin("token", sdk.NewInt(5))
	bank.EXPECT().
		SendCoinsFromAccountToModule(sdkCtx, addr, types.ModuleName, sdk.NewCoins(bet)).
		Return(errors.New("oops"))

	// Action
	err := k.CollectBet(sdkCtx, el)

	// Assert
	expected := sdkerrors.Wrapf(errors.New("oops"), types.ErrCanNotPayBet.Error())
	require.Equal(t, expected.Error(), err.Error())
}

func TestRefundBet(t *testing.T) {
	// SetUp
	k, ctx, _, bank, _ := setupTest(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	el := &types.EnterLotteryTx{
		UserAddress: client1,
		Bet:         4,
		Denom:       "token",
	}
	addr, _ := types.GetAddress(el.UserAddress)
	coin := sdk.NewCoin("token", sdk.NewInt(4))
	bank.EXPECT().
		SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, addr, sdk.NewCoins(coin)).
		Return(nil)

	// Acction
	k.RefundBet(sdkCtx, el)
}

func TestRefundBetToInvalidAddress(t *testing.T) {
	// SetUp
	k, ctx, _, _, _ := setupTest(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	el := &types.EnterLotteryTx{
		UserAddress: "xxxx",
	}

	// Assert
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "user address is invalid: xxxx: decoding bech32 failed: invalid bech32 string length 4", r)
	}()

	// Acction
	k.RefundBet(sdkCtx, el)
}

func TestRefundBetFailedWhenSendRefundToTheUser(t *testing.T) {
	// SetUp
	k, ctx, _, bank, _ := setupTest(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	el := &types.EnterLotteryTx{
		UserAddress: client1,
		Bet:         4,
		Denom:       "token",
	}
	addr, _ := types.GetAddress(el.UserAddress)
	coin := sdk.NewCoin("token", sdk.NewInt(4))
	bank.EXPECT().
		SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, addr, sdk.NewCoins(coin)).
		Return(errors.New("oops"))

	// Assert
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "cannot pay reward to the winner: oops", r)
	}()

	// Acction
	k.RefundBet(sdkCtx, el)
}
