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
		{Address: "cosmos13qvqehe2yv5fxeah7363pd8w32x6fgql0jnkzs", Bet: 11, Denom: "token"},
		{Address: "cosmos13af4dg2putus363llvq0kjsrfct42ntv26gd6d", Bet: 12, Denom: "token"},
		{Address: "cosmos1rmg598ae3r62euwqq2j4tjmj9kdrjmnzjh9000", Bet: 13, Denom: "token"},
		{Address: "cosmos1t9p63h6k5s7mhz2uykn6hnjmtssc0l0lt9s7fj", Bet: 14, Denom: "token"},
		{Address: "cosmos13nt74j5t4y3hhfjx40p42hdc2p308q48uccth3", Bet: 15, Denom: "token"},
		{Address: "cosmos102qnej88ljdkl2lsgf23vauhqgmwte8jra54xq", Bet: 16, Denom: "token"},
		{Address: "cosmos1md6g8sak7zm4wjyjpht4shaa7mhcgxkhfgch6w", Bet: 17, Denom: "token"},
		{Address: "cosmos1jgwue7rjjglesw6l86ydedlsepl6ulz04hg3qp", Bet: 18, Denom: "token"},
		{Address: "cosmos1zml56sa8ngxrs0txsv5hqujjgcerr6fmr6a95t", Bet: 19, Denom: "token"},
		{Address: "cosmos1xfzyzyxy9ynmwpcpk6tn4635tarcqdfu42cxh9", Bet: 20, Denom: "token"},
	}
	got := getWinnerIndex(u)
	fmt.Println(got)
}
