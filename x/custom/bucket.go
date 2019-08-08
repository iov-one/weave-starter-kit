package custom

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
)

type StateIndexedBucket struct {
	orm.IDGenBucket
}

func NewStateIndexedBucket() *StateIndexedBucket {
	b := migration.NewBucket(packageName, "stateind", orm.NewSimpleObj(nil, &StateIndexed{}))
	return &StateIndexedBucket{
		IDGenBucket: orm.WithSeqIDGenerator(b, "id"),
	}
}

// GetStateIndexed loads the StateIndexed for the given id. If it does not exist then ErrNotFound is returned.
func (b *StateIndexedBucket) GetStateIndexed(db weave.KVStore, id []byte) (*StateIndexed, error) {
	obj, err := b.Get(db, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load indexed state")
	}
	return asStateIndexed(obj)
}

func asStateIndexed(obj orm.Object) (*StateIndexed, error) {
	if obj == nil || obj.Value() == nil {
		return nil, errors.Wrap(errors.ErrNotFound, "unknown id")
	}
	rev, ok := obj.Value().(*StateIndexed)
	if !ok {
		return nil, errors.Wrapf(errors.ErrModel, "invalid type: %T", obj.Value())
	}
	return rev, nil
}
type StateBucket struct {
	orm.ModelBucket
}

func NewStateBucket() *StateBucket {
	b := orm.NewModelBucket("state", &State{})
	return &StateBucket{
		ModelBucket: migration.NewModelBucket("mstate", b),
	}
}
