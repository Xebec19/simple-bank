type Store struct {
	*Queries
	db *sql.DB
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{

}

func NewStore(db *sql.DB) *Store{
	return &Store{
		db: db,
		Queries: New(db)
	}
}

tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})

type TxOptions struct {
	Isolation IsolationLevel
	ReadOnly bool
}

func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx,nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v",err,rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error){}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxResult,error){
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amount
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount
		})
		if err != nil {
			return err
		}

		return nil
		// TODO: update accounts balance
	})
	return result, err
}