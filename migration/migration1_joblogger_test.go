package migration

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/textileio/powergate/v2/ffs"
	"github.com/textileio/powergate/v2/tests"
	"github.com/textileio/powergate/v2/util"
)

func TestV1_JobLogger(t *testing.T) {
	t.Parallel()

	ds := tests.NewTxMapDatastore()

	pre(t, ds, "testdata/v1_JobLogger.pre")

	c1, _ := util.CidFromString("QmPewMLNZEgnLxaenjo9Q5qwQwW3zHZ7Ac973UmeJ6VWHE")
	c2, _ := util.CidFromString("QmZTMaDfCMWqhUXYDnKup8ctCTyPxnriYW7G4JR8KXoX5M")
	cidOwners := map[cid.Cid][]ffs.APIID{
		c1: {ffs.APIID("ID1"), ffs.APIID("ID2")},
		c2: {ffs.APIID("ID3")},
	}
	txn, _ := ds.NewTransaction(false)
	err := migrateJobLogger(txn, cidOwners)
	require.NoError(t, err)
	require.NoError(t, txn.Commit())

	post(t, ds, "testdata/v1_JobLogger.post")
}
