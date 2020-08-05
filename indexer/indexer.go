package indexer

import (
	"fmt"
	"log"
	"sync"

	"github.com/terra-project/mantle/graph"
	"github.com/terra-project/mantle/graph/generate"
	"github.com/terra-project/mantle/types"
)

type IndexerBaseInstance struct {
	indexers       []types.Indexer
	indexerOutputs [][]types.ModelType
	committer      types.GraphQLCommitter
	querier        types.GraphQLQuerier
}

func NewIndexerBaseInstance(
	indexers []types.Indexer,
	indexerOutputs [][]types.ModelType,
	querier types.GraphQLQuerier,
	committer types.GraphQLCommitter,
) *IndexerBaseInstance {
	return &IndexerBaseInstance{
		indexers:       indexers,
		indexerOutputs: indexerOutputs,
		committer:      committer,
		querier:        querier,
	}
}

func (instance *IndexerBaseInstance) RunIndexerRound() {
	// create wait group for ALL indexers
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(instance.indexers))

	for i, indexer := range instance.indexers {

		// indexer outputs are necessary for detecting self-reference
		indexerOutput := instance.indexerOutputs[i]

		runIndexer(
			&waitGroup,
			instance.committer,
			instance.querier,
			indexer,
			indexerOutput,
		)
	}

	waitGroup.Wait()
}

func runIndexer(
	wg *sync.WaitGroup,
	committer types.GraphQLCommitter,
	querier types.GraphQLQuerier,
	indexer types.Indexer,
	indexerOutput []types.ModelType,
) {
	var isolatedQuerier = createIsolatedQuerier(querier, indexerOutput)
	var isolatedCommitter = createIsolatedCommitter(committer)

	go func() {
		defer wg.Done()
		indexer(isolatedQuerier, isolatedCommitter)
	}()
}

func createIsolatedQuerier(
	querier types.GraphQLQuerier,
	indexerOutput []types.ModelType,
) types.IndexerQuerier {
	return func(query interface{}, variables types.GraphQLParams) error {
		qs := generate.GenerateQuery(query, variables)
		result := querier(qs, variables, indexerOutput)

		if result.HasErrors() {
			log.Print("IsolatedQuerier Failed. --")

			for _, err := range result.Errors {
				fmt.Println(err, err.Locations, err.Path)
			}

			return fmt.Errorf("Isolated Querier error")
		}

		return graph.UnmarshalGraphQLResult(result, query)
	}
}

func createIsolatedCommitter(committer types.GraphQLCommitter) types.IndexerCommitter {
	return func(entity interface{}) error {
		return committer(entity)
	}
}
