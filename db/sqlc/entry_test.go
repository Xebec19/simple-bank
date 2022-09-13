package db

import (
	"context"
	"testing"

	"github.com/Xebec19/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createTestEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)
	args := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}
	entry, _ := testQueries.CreateEntry(context.Background(), args)
	return entry
}

func TestReadEntry(t *testing.T) {
	entry1 := createTestEntry(t)
	entry2, err := testQueries.ReadEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestUpdateEntries(t *testing.T) {
	entry1 := createTestEntry(t)
	amount1 := util.RandomMoney()
	args := UpdateEntriesParams{
		ID:     entry1.ID,
		Amount: amount1,
	}
	entry2, err := testQueries.UpdateEntries(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, entry2.Amount, amount1)
	require.Equal(t, entry1.ID, entry2.ID)
}
