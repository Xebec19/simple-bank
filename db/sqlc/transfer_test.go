package db

import (
	"context"
	"testing"

	"github.com/Xebec19/simple-bank/util"
	"github.com/stretchr/testify/require"
)

// createTestTransfer creates a random transfer with random accounts
func createTestTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer1, _ := testQueries.CreateTransfer(context.Background(), args)
	return transfer1
}

func TestReadAccountTransfer(t *testing.T) {
	transfer := createTestTransfer(t)
	args := ReadAccountTransferParams{
		FromAccountID: transfer.FromAccountID,
		Limit:         1,
		Offset:        0,
	}
	transfers, err := testQueries.ReadAccountTransfer(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 1)
	for _, trf := range transfers {
		require.NotEmpty(t, trf)
		require.Equal(t, trf.FromAccountID, transfer.FromAccountID)
	}
}

/*
todo test updateTransfer
create a transfer, use updateTransfer and then use
readTranfer to validate its value
*/
