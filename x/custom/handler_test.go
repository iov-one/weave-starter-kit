package custom

import (
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestCreateStateIndexed(t *testing.T) {

	meta := &weave.Metadata{Schema: 1}

	cases := map[string]struct {
		msg             weave.Msg
		expected        *StateIndexed
		wantCheckErrs   map[string]*errors.Error
		wantDeliverErrs map[string]*errors.Error
	}{
		"success": {
			msg: &CreateStateIndexedMsg{
				Metadata: meta,
				InnerStateEnum: InnerStateEnum_CaseOne,
				Str: "cstm_str",
				Byte: []byte{0, 1},
			},
			expected: &StateIndexed{
				Metadata: meta,
				InnerStateEnum: InnerStateEnum_CaseOne,
				Str: "cstm_str",
				Byte: []byte{0, 1},
			},
			wantCheckErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":   nil,
				"Byte":     nil,
				"DeletedAt":      nil,
			},
			wantDeliverErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"Str":   nil,
				"Byte":     nil,
				"DeletedAt":      nil,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			auth := &weavetest.Auth{}

			h := NewStateIndexedHandler(auth)
			kv := store.MemStore()
			bucket := NewStateIndexedBucket()
			migration.MustInitPkg(kv, packageName)

			tx := &weavetest.Tx{Msg: tc.msg}

			if _, err := h.Check(nil, kv, tx); err != nil {
				for field, wantErr := range tc.wantCheckErrs {
					assert.FieldError(t, err, field, wantErr)
				}
			}
			
			res, err := h.Deliver(nil, kv, tx)
			for field, wantErr := range tc.wantDeliverErrs {
				assert.FieldError(t, err, field, wantErr)
			}

			if tc.expected != nil {
				stored, err := bucket.GetStateIndexed(kv, res.Data)

				assert.Nil(t, err)
				assert.Equal(t, tc.expected, &stored)
			}
		})
	}

}
