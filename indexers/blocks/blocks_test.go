package blocks

func noop() {

}

//
// import (
// 	"fmt"
// 	"github.com/terra-money/mantle/indexers/tx_infos"
// 	"github.com/terra-money/mantle/test/fixtures"
// 	"github.com/terra-money/mantle/utils"
// 	"github.com/terra-money/mantle-sdk/test"
// 	"sync"
// 	"testing"
// )
//
// func TestIndexBlocks(t *testing.T) {
// 	simapp, accounts := fixtures.NewTestChain(
// 		tx_infos.RegisterTxInfos,
// 		RegisterBlocks,
// 	)
//
// 	fmt.Println("list of all accounts")
// 	for _, account := range accounts {
// 		fmt.Println(account.GetAddress().String())
// 	}
//
// 	// create random blocks
// 	for i := 0; i < 20; i++ {
// 		testBlock := test.NewBlock()
// 		wg := sync.WaitGroup{}
// 		wg.Add(len(accounts)-1)
// 		for j := 0; j < len(accounts)-1; j++ {
// 			sender := accounts[j]
// 			receiver := accounts[j+1]
// 			go utils.AppendTxAsync(testBlock, sender, receiver, &wg)
// 		}
//
// 		wg.Wait()
//
// 		_ = simapp.Inject(testBlock.ToBlock())
// 	}
//
// 	for{}
// }
