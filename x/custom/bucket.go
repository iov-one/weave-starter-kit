package custom

import (
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/migration"
)

type StateIndexedBucket struct {
	orm.ModelBucket
}

func newStateIndexedBucket() *StateIndexedBucket {
	b := orm.NewModelBucket("stateIndexed", &StateIndexed{})
	return &StateIndexedBucket{
		ModelBucket: migration.NewModelBucket("mStateIndexed", b),
	}
}

type StateBucket struct {
	orm.ModelBucket
}

func newStateBucket() *StateBucket {
	b := orm.NewModelBucket("state", &State{})
	return &StateBucket{
		ModelBucket: migration.NewModelBucket("mState", b),
	}
}
