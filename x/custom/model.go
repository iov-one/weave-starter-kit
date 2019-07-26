package custom

import (
	"github.com/iov-one/tutorial/morm"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
)

var _ morm.Model = (*CustomStateIndexed)(nil)

func (m *CustomStateIndexed) SetID(id []byte) error {
	m.ID = id
	return nil
}

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

	if err := m.DeletedAt.Validate(); err != nil {
		errs = errors.AppendField(errs, "DeletedAt", m.DeletedAt.Validate())
	} else if m.DeletedAt == 0 {
		errs = errors.Append(errs,
			errors.Field("DeletedAt", errors.ErrEmpty, "missing created at"))
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

var _ orm.Model = (*CustomState)(nil)

func (m *CustomState) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	if m.InnerState == nil {
		errs = errors.Append(errs, errors.Field("InnerState", errors.ErrEmpty, "missing inner state"))
	}
	if err := m.CreatedAt.Validate(); err != nil {
		errs = errors.AppendField(errs, "CreatedAt", m.CreatedAt.Validate())
	} else if m.CreatedAt == 0 {
		errs = errors.Append(errs,
			errors.Field("DeletedAt", errors.ErrEmpty, "missing created at"))
	}

	return errs
}

func (m *CustomState) Copy() orm.CloneableData {
	return &CustomState{
		Metadata:      m.Metadata.Copy(),
		InnerState:    m.InnerState,
		CustomAddress: copyBytes(m.CustomAddress),
		CreatedAt:     m.CreatedAt,
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
