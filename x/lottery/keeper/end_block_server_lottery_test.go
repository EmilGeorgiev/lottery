package keeper

import (
	"fmt"
	"testing"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
)

func TestGetWinner(t *testing.T) {
	u := []*types.User{
		{Address: "cosmos19xn45se48ua3jjkdrpehq3gj9a00gg5gptmr09", Bet: 1, Denom: "token"},
		{Address: "cosmos1cpdt2j7jfhaek504t56wpfv0zrq83ez8q4u335", Bet: 2, Denom: "token"},
		{Address: "cosmos1yhfshaufadllmr5p2588mu6uv7cuddzgzxfr4t", Bet: 3, Denom: "token"},
		{Address: "cosmos1j7tww47qlpp60pfjwa9cqawnju2tu347ddmn2n", Bet: 4, Denom: "token"},
		{Address: "cosmos1wq43k8mq8e6xu6xg9gg7seqc0j0m96klgq4a0y", Bet: 5, Denom: "token"},
		{Address: "cosmos1erw3hf5vg5fsa5nrzent2t6sen0jlmemtlga2t", Bet: 6, Denom: "token"},
		{Address: "cosmos1mr9sxeena7xm6fjm3zj0q3pmfah7k522cmur8p", Bet: 7, Denom: "token"},
		{Address: "cosmos1p2y92270yqggk2du3d3a2yuqgzzf9fluaey946", Bet: 8, Denom: "token"},
		{Address: "cosmos1l6ptczusekwszctm9e6sjkr9sgg3jyvcskxzt4", Bet: 9, Denom: "token"},
		{Address: "cosmos1xfzyzyxy9ynmwpcpk6tn4635tarcqdfu42cxh8", Bet: 10, Denom: "token"},
	}
	got := getWinnerIndex(u)
	fmt.Println(got)
}
