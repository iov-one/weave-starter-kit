package custom

import (
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
)

var _ orm.Model = (*CustomStateIndexed)(nil)

func (m *CustomStateIndexed) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "ID", isGenID(m.ID, true))
	errs = errors.AppendField(errs, "CustomString", customStringValidation(m.CustomString))
	if m.CustomByte == nil {
		errs = errors.Append(errs, errors.Field("CustomByte", errors.ErrEmpty, "missing custom byte"))
	}
	if m.InnerStateEnum != InnerStateEnum_CaseOne && m.InnerStateEnum != InnerStateEnum_CaseTwo {
		errs = errors.Append(errs,
			errors.Field("InnerStateEnum", errors.ErrState, "invalid inner state enum"))
	}
	// TODO add custom validation for your state fields
	return errs
}

func (m *CustomStateIndexed) Copy() orm.CloneableData {
	return &CustomStateIndexed{
		Metadata:       m.Metadata.Copy(),
		ID:             copyBytes(m.ID),
		InnerStateEnum: m.InnerStateEnum,
		CustomInt:      m.CustomInt,
		CustomString:   m.CustomString,
		CustomByte:     copyBytes(m.CustomByte),
		DeletedAt:      m.DeletedAt,
	}
}

// isGenID ensures that the ID is 8 byte input.
// if allowEmpty is set, we also allow empty
func isGenID(id []byte, allowEmpty bool) error {
	if len(id) == 0 {
		if allowEmpty {
			return nil
		}
		return errors.Wrap(errors.ErrEmpty, "missing id")
	}
	if len(id) != 8 {
		return errors.Wrap(errors.ErrInput, "id must be 8 bytes")
	}
	return nil
}

func copyBytes(in []byte) []byte {
	if in == nil {
		return nil
	}
	cpy := make([]byte, len(in))
	copy(cpy, in)
	return cpy
}
