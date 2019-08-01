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

func TestCreateCustomStateIndexed(t *testing.T) {

	meta := &weave.Metadata{Schema: 1}

	cases := map[string]struct {
		msg             weave.Msg
		expected        *CustomStateIndexed
		wantCheckErrs   map[string]*errors.Error
		wantDeliverErrs map[string]*errors.Error
	}{
		"nil message": {
			wantCheckErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
			wantDeliverErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
		"success": {
			msg: &CreateCustomStateIndexedMsg{
				Metadata: meta,
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomString: "cstm_str",
				CustomByte: []byte{0, 1},
			},
			expected: &CustomStateIndexed{
				Metadata: meta,
				InnerStateEnum: InnerStateEnum_CaseOne,
				CustomString: "cstm_str",
				CustomByte: []byte{0, 1},
			},
			wantCheckErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
			wantDeliverErrs: map[string]*errors.Error{
				"Metadata":       nil,
				"ID":             nil,
				"InnerStateEnum": nil,
				"CustomString":   nil,
				"CustomByte":     nil,
				"DeletedAt":      nil,
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			auth := &weavetest.Auth{}

			h := NewStateIndexedHandler(auth)
			kv := store.MemStore()
			bucket := NewCustomStateIndexedBucket()
			migration.MustInitPkg(kv, packageName)

			tx := &weavetest.Tx{Msg: tc.msg}

			if _, err := h.Check(nil, kv, tx); err != nil {
				for field, wantErr := range tc.wantCheckErrs {
					assert.FieldError(t, err, field, wantErr)
				}
			}
			
			dres, err := h.Deliver(nil, kv, tx)
			for field, wantErr := range tc.wantDeliverErrs {
				assert.FieldError(t, err, field, wantErr)
			}

			if tc.expected != nil {
				var stored CustomStateIndexed
				err = bucket.One(kv, dres.Data, &stored)
				assert.Nil(t, err)
				assert.Equal(t, tc.expected, &stored)
			}
		})
	}

}
