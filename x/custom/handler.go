package custom

import (
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/x"
)

const (
	packageName              = "custom"
	newCustomStateCost int64 = 100
)

// RegisterQuery registers buckets for querying.
func RegisterQuery(qr weave.QueryRouter) {
	NewStateIndexedBucket().Register("indexedStates", qr)
	NewStateBucket().Register("states", qr)
}

// RegisterRoutes registers handlers for message processing.
func RegisterRoutes(r weave.Registry, auth x.Authenticator) {
	r = migration.SchemaMigratingRegistry(packageName, r)

	r.Handle(&CreateStateIndexedMsg{}, NewStateIndexedHandler(auth))
	r.Handle(&CreateStateMsg{}, NewStateHandler(auth))
}

// ------------------- StateIndexed HANDLER -------------------

// StateIndexedHandler will handle creating custom indexed state buckets
type StateIndexedHandler struct {
	auth x.Authenticator
	b   *StateIndexedBucket
}

var _ weave.Handler = StateIndexedHandler{}

// NewStateIndexedHandler creates a handler
func NewStateIndexedHandler(auth x.Authenticator) weave.Handler {
	return StateIndexedHandler{
		auth: auth,
		b:   NewStateIndexedBucket(),
	}
}

// validate does all common pre-processing between Check and Deliver
func (h StateIndexedHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*CreateStateIndexedMsg, error) {
	var msg CreateStateIndexedMsg

	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, errors.Wrap(err, "load msg")
	}

	return &msg, nil
}

// Check just verifies it is properly formed and returns
// the cost of executing it.
func (h StateIndexedHandler) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	return &weave.CheckResult{GasAllocated: newCustomStateCost}, nil
}

// Deliver creates an custom state and saves if all preconditions are met
func (h StateIndexedHandler) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, err := h.validate(ctx, store, tx)
	
	if err != nil {
		return nil, err
	}

	stateIndexed := &StateIndexed{
		Metadata: &weave.Metadata{},
		InnerStateEnum: msg.InnerStateEnum,
		Str: msg.Str,
		Byte: msg.Byte,
	}

	if _, err = h.b.Create(store, stateIndexed); err != nil {
		return nil, errors.Wrap(err, "cannot store indexed state")
	}

	if err != nil {
		return nil, err
	}

	return &weave.DeliverResult{Data: stateIndexed.ID}, err
}

// ------------------- CustomState HANDLER -------------------

// StateHandler will handle creating custom state buckets
type StateHandler struct {
	auth x.Authenticator
	b   *StateBucket
}

var _ weave.Handler = StateHandler{}

// NewStateHandler creates a handler
func NewStateHandler(auth x.Authenticator) weave.Handler {
	return StateHandler{
		auth: auth,
		b:   NewStateBucket(),
	}
}

// validate does all common pre-processing between Check and Deliver
func (h StateHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*CreateStateMsg, error) {
	var msg CreateStateMsg

	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, errors.Wrap(err, "load msg")
	}

	return &msg, nil
}

// Check just verifies it is properly formed and returns
// the cost of executing it.
func (h StateHandler) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, err := h.validate(ctx, store, tx)
	if err != nil {
		return nil, err
	}

	return &weave.CheckResult{GasAllocated: newCustomStateCost}, nil
}

// Deliver creates an custom state and saves if all preconditions are met
func (h StateHandler) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, err := h.validate(ctx, store, tx)
	
	if err != nil {
		return nil, err
	}

	now := weave.AsUnixTime(time.Now())

	state := &State{
		Metadata: &weave.Metadata{},
		InnerState: msg.InnerState,
		Address: msg.Address,
		CreatedAt: now,
	}

	_, err = h.b.Put(store, nil, state)

	if err != nil {
		return nil, err
	}

	return &weave.DeliverResult{}, err
}
