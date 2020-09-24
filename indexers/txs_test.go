package indexers

import (
	"fmt"
	"github.com/terra-project/mantle-official/test/fixtures"
	"github.com/terra-project/mantle/test"
	"sync"
	"testing"
)

func TestTxs(t *testing.T) {
	simapp, accounts := fixtures.NewTestChain(
		RegisterTxs,
	)

	fmt.Println("list of all accounts")
	for _, account := range accounts {
		fmt.Println(account.GetAddress().String())
	}

	// create random blocks
	for i := 0; i < 20; i++ {
		testBlock := test.NewBlock()
		wg := sync.WaitGroup{}
		wg.Add(len(accounts)-1)
		for j := 0; j < len(accounts)-1; j++ {
			sender := accounts[j]
			receiver := accounts[j+1]
			go AppendTxAsync(testBlock, sender, receiver, &wg)
		}

		wg.Wait()

		_ = simapp.Inject(testBlock.ToBlock())
	}

	for{}

}
