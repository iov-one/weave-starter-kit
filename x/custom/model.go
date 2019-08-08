package custom

import (
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
)

func init() {
	// Migration needs to be registered for every message introduced in the codec.
	// This is the convention to message versioning.
	migration.MustRegister(1, &StateIndexed{}, migration.NoModification)
	migration.MustRegister(1, &State{}, migration.NoModification)
}

var _ orm.Model = (*StateIndexed)(nil)

// SetID is a minimal implementation, useful when the ID is a separate protobuf field
func (m *StateIndexed) SetID(id []byte) error {
	m.ID = id
	return nil
}

// Validate ensures the StateIndexed fields are valid
func (m *StateIndexed) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	errs = errors.AppendField(errs, "ID", isGenID(m.ID, true))
	errs = errors.AppendField(errs, "Str", stringValidation(m.Str))
	if m.Byte == nil {
		errs = errors.AppendField(errs, "Byte", errors.ErrEmpty)
	}
	if m.InnerStateEnum != InnerStateEnum_CaseOne && m.InnerStateEnum != InnerStateEnum_CaseTwo {
		errs = errors.AppendField(errs, "InnerStateEnum", errors.ErrState)
	}
	
	return errs
}

// Copy produces a new StateIndexed clone to fulfill the Model interface
func (m *StateIndexed) Copy() orm.CloneableData {
	return &StateIndexed{
		Metadata:       m.Metadata.Copy(),
		ID:             copyBytes(m.ID),
		InnerStateEnum: m.InnerStateEnum,
		Str:            m.Str,
		Byte:           copyBytes(m.Byte),
		DeletedAt:      m.DeletedAt,
	}
}

var _ orm.Model = (*State)(nil)

// Validate ensures the State fields are valid
func (m *State) Validate() error {
	var errs error

	errs = errors.AppendField(errs, "Metadata", m.Metadata.Validate())
	if m.InnerState == nil {
		errs = errors.AppendField(errs, "InnerState", errors.ErrEmpty)
	}
	errs = errors.AppendField(errs, "Address", m.Address.Validate())
	if err := m.CreatedAt.Validate(); err != nil {
		errs = errors.AppendField(errs, "CreatedAt", m.CreatedAt.Validate())
	} else if m.CreatedAt == 0 {
		errs = errors.AppendField(errs, "CreatedAt", errors.ErrEmpty)
	}

	return errs
}

// Copy produces a new State clone to fulfill the Model interface
func (m *State) Copy() orm.CloneableData {
	return &State{
		Metadata:   m.Metadata.Copy(),
		InnerState: m.InnerState,
		Address:    copyBytes(m.Address),
		CreatedAt:  m.CreatedAt,
	}
}

// isGenID ensures that the ID is 8 byte input.
// if allowEmpty is set, we also allow empty
func isGenID(id []byte, allowEmpty bool) error {
	if len(id) == 0 {
		if allowEmpty {
			return nil
		}
		return errors.Wrap(errors.ErrEmpty, "required")
	}
	if len(id) != 8 {
		return errors.Wrap(errors.ErrInput, "must be 8 bytes")
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
