package client

import (
	"context"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/x/sigs"
	"github.com/pkg/errors"
	cmn "github.com/tendermint/tendermint/libs/common"
	tmpubsub "github.com/tendermint/tendermint/libs/pubsub"
	"github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// BroadcastTxSyncDefaultTimeOut timeout for sync tx broadcasting
const BroadcastTxSyncDefaultTimeOut = 15 * time.Second

var QueryNewBlockHeader = tmtypes.EventQueryNewBlockHeader

// Client is an interface to interact with weave apps
type Client interface {
	TendermintClient() client.Client
	GetUser(addr weave.Address) (*UserResponse, error)
	GetWallet(addr weave.Address) (*WalletResponse, error)
	BroadcastTx(tx weave.Tx) BroadcastTxResponse
	BroadcastTxAsync(tx weave.Tx, out chan<- BroadcastTxResponse)
	BroadcastTxSync(tx weave.Tx, timeout time.Duration) BroadcastTxResponse
	AbciQuery(path string, data []byte) (AbciResponse, error)
}

// CustomClient is a tendermint client wrapped to provide
// simple access to the data structures used in custom modle.
type CustomClient struct {
	conn client.Client
	// subscriber is a unique identifier for subscriptions
	subscriber string
}

// NewClient wraps a CustomClient around an existing
// tendermint client connection.
func NewClient(conn client.Client) *CustomClient {
	return &CustomClient{
		conn:       conn,
		subscriber: "tools-client",
	}
}

// TendermintClient returns underlying tendermint client
func (cc *CustomClient) TendermintClient() client.Client {
	return cc.conn
}

// Nonce has a client/address pair, queries for the nonce
// and caches recent nonce locally to quickly sign
type Nonce struct {
	mutex     sync.Mutex
	client    Client
	addr      weave.Address
	nonce     int64
	fromQuery bool
}

// NewNonce creates a nonce for a client / address pair.
// Call Query to force a query, Next to use cache if possible
func NewNonce(client Client, addr weave.Address) *Nonce {
	return &Nonce{client: client, addr: addr}
}

// Query always queries the blockchain for the next nonce
func (n *Nonce) Query() (int64, error) {
	user, err := n.client.GetUser(n.addr)
	if err != nil {
		return 0, err
	}
	n.mutex.Lock()
	if user != nil {
		n.nonce = user.UserData.Sequence
	} else {
		n.nonce = 0 // new account starts at 0
	}
	n.fromQuery = true
	n.mutex.Unlock()
	return n.nonce, nil
}

// Next will use a cached value if present, otherwise Query
// It will always increment by 1, assuming last nonce
// was properly used. This is designed for cases where
// you want to rapidly generate many tranasactions without
// querying the blockchain each time
func (n *Nonce) Next() (int64, error) {
	n.mutex.Lock()
	initializedFromBlockchain := !n.fromQuery && n.nonce == 0
	n.mutex.Unlock()
	if initializedFromBlockchain {
		return n.Query()
	}
	n.mutex.Lock()
	n.nonce++
	n.fromQuery = false
	result := n.nonce
	n.mutex.Unlock()
	return result, nil
}

//************ generic (weave) functionality *************//

// Status will return the raw status from the node
func (cc *CustomClient) Status() (*ctypes.ResultStatus, error) {
	return cc.conn.Status()
}

// Genesis will return the genesis directly from the node
func (cc *CustomClient) Genesis() (*tmtypes.GenesisDoc, error) {
	gen, err := cc.conn.Genesis()
	if err != nil {
		return nil, err
	}
	return gen.Genesis, nil
}

// ChainID will parse out the chainID from the status result
func (cc *CustomClient) ChainID() (string, error) {
	gen, err := cc.Genesis()
	if err != nil {
		return "", err
	}
	return gen.ChainID, nil
}

// Height will parse out the Height from the status result
func (cc *CustomClient) Height() (int64, error) {
	status, err := cc.conn.Status()
	if err != nil {
		return -1, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}

// AbciResponse contains a query result:
// a (possibly empty) list of key-value pairs, and the height
// at which it queried
type AbciResponse struct {
	// a list of key/value pairs
	Models []weave.Model
	Height int64
}

// AbciQuery calls abci query on tendermint rpc,
// verifies if it is an error or empty, and if there is
// data pulls out the ResultSets from keys and values into
// a useful AbciResponse struct
func (cc *CustomClient) AbciQuery(path string, data []byte) (AbciResponse, error) {
	var out AbciResponse

	q, err := cc.conn.ABCIQuery(path, data)
	if err != nil {
		return out, err
	}
	resp := q.Response
	if resp.IsErr() {
		return out, errors.Errorf("(%d): %s", resp.Code, resp.Log)
	}
	out.Height = resp.Height

	if len(resp.Key) == 0 {
		return out, nil
	}

	// assume there is data, parse the result sets
	var keys, vals app.ResultSet
	err = keys.Unmarshal(resp.Key)
	if err != nil {
		return out, err
	}
	err = vals.Unmarshal(resp.Value)
	if err != nil {
		return out, err
	}

	out.Models, err = app.JoinResults(&keys, &vals)
	return out, err
}

// TxSearch searches transactions using underlying tendermint client
func (cc *CustomClient) TxSearch(query string, prove bool, page, perPage int) (*ctypes.ResultTxSearch, error) {
	return cc.conn.TxSearch(query, prove, page, perPage)
}

// BroadcastTxResponse is the result of submitting a transaction.
type BroadcastTxResponse struct {
	Error    error                           // not-nil if there was an error sending
	Response *ctypes.ResultBroadcastTxCommit // not-nil if we got response from node
}

// IsError returns the error for failure if it failed,
// or null if it succeeded
func (b BroadcastTxResponse) IsError() error {
	if b.Error != nil {
		return b.Error
	}
	if b.Response.CheckTx.IsErr() {
		ctx := b.Response.CheckTx
		return errors.Errorf("CheckTx error: (%d) %s", ctx.Code, ctx.Log)
	}
	if b.Response.DeliverTx.IsErr() {
		dtx := b.Response.DeliverTx
		return errors.Errorf("CheckTx error: (%d) %s", dtx.Code, dtx.Log)
	}
	return nil
}

// BroadcastTx serializes a signed transaction and writes to the
// blockchain. It returns when the tx is committed to the
// blockchain.
//
// If you want high-performance, parallel sending, use BroadcastTxAsync
func (cc *CustomClient) BroadcastTx(tx weave.Tx) BroadcastTxResponse {
	out := make(chan BroadcastTxResponse, 1)
	defer close(out)
	go cc.BroadcastTxAsync(tx, out)
	res := <-out
	return res
}

// BroadcastTxSync brodcasts transactions synchronously
func (cc *CustomClient) BroadcastTxSync(tx weave.Tx, timeout time.Duration) BroadcastTxResponse {
	data, err := tx.Marshal()
	if err != nil {
		return BroadcastTxResponse{Error: err}
	}

	res, err := cc.conn.BroadcastTxSync(data)
	if err != nil {
		return BroadcastTxResponse{Error: err}
	}
	if res.Code != 0 {
		return BroadcastTxResponse{Error: errors.WithMessage(fmt.Errorf("CheckTx failed with code %d", res.Code), res.Log)}
	}

	// and wait for confirmation
	evt, err := cc.WaitForTxEvent(data, tmtypes.EventTx, timeout)
	if err != nil {
		return BroadcastTxResponse{Error: err}
	}

	txe, ok := evt.(tmtypes.EventDataTx)
	if !ok {
		if err != nil {
			return BroadcastTxResponse{Error: fmt.Errorf("WaitForOneEvent did not return an EventDataTx object")}
		}
	}

	return BroadcastTxResponse{
		Response: &ctypes.ResultBroadcastTxCommit{
			DeliverTx: txe.Result,
			Height:    txe.Height,
			Hash:      txe.Tx.Hash(),
		},
	}
}

// WaitForTxEvent listens for and particular event type of evtTyp to be fired
func (cc *CustomClient) WaitForTxEvent(tx tmtypes.Tx, evtTyp string, timeout time.Duration) (tmtypes.TMEventData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	query := tmtypes.EventQueryTxFor(tx)

	uuid := hex.EncodeToString(append(tx.Hash(), cmn.RandBytes(2)...))
	evts, err := cc.conn.Subscribe(ctx, uuid, query.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to subscribe")
	}

	// make sure to unregister after the test is over
	defer cc.conn.UnsubscribeAll(ctx, uuid)

	select {
	case evt := <-evts:
		return evt.Data.(tmtypes.TMEventData), nil
	case <-ctx.Done():
		return nil, errors.New("timed out waiting for event")
	}
}

// BroadcastTxAsync can be run in a goroutine and will output
// the result or error to the given channel.
// Useful if you want to send many tx in parallel
func (cc *CustomClient) BroadcastTxAsync(tx weave.Tx, out chan<- BroadcastTxResponse) {
	data, err := tx.Marshal()
	if err != nil {
		out <- BroadcastTxResponse{Error: err}
		return
	}

	// TODO: make this async, maybe adjust return value
	res, err := cc.conn.BroadcastTxCommit(data)
	msg := BroadcastTxResponse{
		Error:    err,
		Response: res,
	}
	out <- msg
}

// SubscribeHeaders queries for headers and starts a goroutine
// to typecase the events into Headers. Returns a cancel
// function. If you don't want the automatic goroutine, use
// Subscribe(QueryNewBlockHeader, out)
func (cc *CustomClient) SubscribeHeaders(out chan<- *tmtypes.Header) (func(), error) {
	query := tmtypes.EventQueryNewBlockHeader
	pipe, cancel, err := cc.Subscribe(query)
	if err != nil {
		return nil, err
	}
	go func() {
		for msg := range pipe {
			evt, ok := msg.Data.(tmtypes.EventDataNewBlockHeader)
			if !ok {
				// TODO: something else?
				panic("Unexpected event type")
			}
			out <- &evt.Header
		}
		close(out)
	}()
	return cancel, nil
}

// Subscribe will take an arbitrary query and push all events to
// the given channel. If there is no error,
// returns a cancel function that can be called to cancel
// the subscription
func (cc *CustomClient) Subscribe(query tmpubsub.Query) (<-chan ctypes.ResultEvent, func(), error) {
	ctx := context.Background()
	out, err := cc.conn.Subscribe(ctx, cc.subscriber, query.String())
	if err != nil {
		return out, nil, err
	}
	cancel := func() {
		cc.conn.Unsubscribe(ctx, cc.subscriber, query.String())
	}
	return out, cancel, nil
}

// UnsubscribeAll cancels all subscriptions
func (cc *CustomClient) UnsubscribeAll() error {
	ctx := context.Background()
	return cc.conn.UnsubscribeAll(ctx, cc.subscriber)
}

// GetWallet will return a wallet given an address
// If non wallet is present, it will return (nil, nil)
// Error codes are used when the query failed on the server
func (cc *CustomClient) GetWallet(addr weave.Address) (*WalletResponse, error) {
	// make sure we send a valid address to the server
	err := addr.Validate()
	if err != nil {
		return nil, errors.WithMessage(err, "Invalid Address")
	}

	resp, err := cc.AbciQuery("/wallets", addr)
	if err != nil {
		return nil, err
	}
	if len(resp.Models) == 0 { // empty list or nil
		return nil, nil
	}
	// assume only one result
	model := resp.Models[0]
	// make sure the return value is expected
	acct := walletKeyToAddr(model.Key)
	if !addr.Equals(acct) {
		return nil, errors.Errorf("Mismatch. Queried %s, returned %s", addr, acct)
	}
	out := WalletResponse{
		Address: acct,
		Height:  resp.Height,
	}

	// parse the value as wallet bytes
	err = out.Wallet.Unmarshal(model.Value)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// key is the address prefixed with "wallet:"
func walletKeyToAddr(key []byte) weave.Address {
	return key[5:]
}

// UserResponse is a response on a query for a User
type UserResponse struct {
	Address  weave.Address
	UserData sigs.UserData
	Height   int64
}

// GetUser will return nonce and public key registered
// for a given address if it was ever used.
// If it returns (nil, nil), then this address never signed
// a transaction before (and can use nonce = 0)
func (b *CustomClient) GetUser(addr weave.Address) (*UserResponse, error) {
	// make sure we send a valid address to the server
	err := addr.Validate()
	if err != nil {
		return nil, errors.WithMessage(err, "Invalid Address")
	}

	resp, err := b.AbciQuery("/auth", addr)
	if err != nil {
		return nil, err
	}
	if len(resp.Models) == 0 { // empty list or nil
		return nil, nil // no wallet
	}
	// assume only one result
	model := resp.Models[0]

	// make sure the return value is expected
	acct := userKeyToAddr(model.Key)
	if !addr.Equals(acct) {
		return nil, errors.Errorf("Mismatch. Queried %s, returned %s", addr, acct)
	}
	out := UserResponse{
		Address: acct,
		Height:  resp.Height,
	}

	// parse the value as wallet bytes
	err = out.UserData.Unmarshal(model.Value)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// key is the address prefixed with "sigs:"
func userKeyToAddr(key []byte) weave.Address {
	return key[5:]
}
