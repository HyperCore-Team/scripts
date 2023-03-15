package main

import (
	"fmt"
	"math/big"
	"os"
	"znn-sdk-go/zenon"
)

func main() {
	args := os.Args
	keyFilePath := args[1]
	password := args[2]
	z, err := zenon.NewZenon(keyFilePath)

	if err != nil {
		zenon.CommonLogger.Error("Error while creating Zenon SDK instance", "error", err)
		return
	}

	if err := z.Start(password, "ws://127.0.0.1:35998", 0); err != nil {
		zenon.CommonLogger.Error("Error while trying to connect to node", "error", err)
		return
	}

	if err := z.Send(z.Client.TokenApi.IssueToken("LiqZTS1", "LIQZTS1", "liquidity.com", big.NewInt(30000000), big.NewInt(9007199254740991), 6, true, true, true)); err != nil {
		fmt.Println(err)
	}

	if err := z.Send(z.Client.TokenApi.IssueToken("LiqZTS2", "LIQZTS2", "liquidity.com", big.NewInt(30000000), big.NewInt(9007199254740991), 6, true, true, true)); err != nil {
		fmt.Println(err)
	}

	tokens, err := z.Client.TokenApi.GetByOwner(z.Address(), 0, 5)
	if err != nil {
		fmt.Println(err)
	}
	tokenStandards := make([]string, 0)
	znnPercentages := make([]uint32, 0)
	qsrPercentages := make([]uint32, 0)
	minAmounts := make([]*big.Int, 0)
	total := 10000
	if tokens.Count > 0 {
		percentage := total / tokens.Count
		for _, token := range tokens.List {
			tokenStandards = append(tokenStandards, token.ZenonTokenStandard.String())
			znnPercentages = append(znnPercentages, uint32(percentage))
			qsrPercentages = append(qsrPercentages, uint32(percentage))
			minAmounts = append(minAmounts, big.NewInt(100000000))
		}
	}

	//if err := z.Send(z.Client.LiquidityApi.SetTokenTupleMethod(tokenStandards, znnPercentages, qsrPercentages, minAmounts)); err != nil {
	//	fmt.Println(err)
	//}
}