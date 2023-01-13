package testutil

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func (m *MockBankKeeper) ExpectAny(context context.Context) {
	m.EXPECT().GetBalance(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any()).AnyTimes()
}
