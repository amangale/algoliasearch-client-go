package search_index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestReplicaDeletion(t *testing.T) {
	t.Parallel()
	_, index, indexName := cts.InitSearchClient1AndIndex(t)
	replicaIndexName := indexName + "_replica"

	await := algolia.Await()

	{
		res, err := index.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		await.Collect(res)
	}

	{
		res, err := index.SetSettings(search.Settings{
			Replicas: opt.Replicas(replicaIndexName),
		})
		require.NoError(t, err)
		await.Collect(res)
	}

	require.NoError(t, await.Wait())
	require.NoError(t, index.DeleteReplica(replicaIndexName))
}
