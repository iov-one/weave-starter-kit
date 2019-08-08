package custom

import (
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
)

type StateIndexedBucket struct {
	orm.ModelBucket
}

func NewStateIndexedBucket() *StateIndexedBucket {
	b := orm.NewModelBucket("stateind", &StateIndexed{}, orm.WithIDSequence(indexedStateSeq))
	return &StateIndexedBucket{
		ModelBucket: migration.NewModelBucket("mstateindexed", b),
	}
}

var indexedStateSeq = orm.NewSequence("indexedstate", "id")

type StateBucket struct {
	orm.ModelBucket
}

func NewStateBucket() *StateBucket {
	b := orm.NewModelBucket("state", &State{})
	return &StateBucket{
		ModelBucket: migration.NewModelBucket("mstate", b),
	}
}
