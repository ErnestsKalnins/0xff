package feature

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type readState struct {
	ID            uuid.UUID             `json:"id"`
	TechnicalName string                `json:"technicalName"`
	DisplayName   *string               `json:"displayName,omitempty"`
	Description   *string               `json:"description,omitempty"`
	State         featureStateTransport `json:"state"`
	CreatedAt     int64                 `json:"createdAt"`
	UpdatedAt     int64                 `json:"updatedAt"`
}

type writeState struct {
	ID            uuid.UUID
	FeatureID     uuid.UUID
	EnvironmentID uuid.UUID
	State         featureState
	UpdatedAt     int64
}

// state is a sum-type of all possible feature states.
type featureState interface {
	featureState()
}

// All possible featureState types.
const (
	featureStateTypeConstant = "constant"
	featureStateTypeRandom   = "random"
	featureStateTypeSegment  = "segment"
)

// featureStateConstant is the simplest state - either the feature is on or off.
type featureStateConstant bool

func (featureStateConstant) featureState() {}

// featureStateRandom is the probability of the feature being on.
type featureStateRandom uint8

func (featureStateRandom) featureState() {}

// featureStateSegment contains the key values which have to be present for the
// feature to be on.
type featureStateSegment map[string]string

func (featureStateSegment) featureState() {}

// featureStateTransport wraps a featureState to implement un/marshaling interfaces
// for JSON, SQL, and potentially other mediums.
type featureStateTransport struct {
	value featureState
}

// MarshalJSON implements json.Marshaler.
func (f featureStateTransport) MarshalJSON() ([]byte, error) {
	var envelope struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	}
	switch v := f.value.(type) {
	case featureStateConstant:
		envelope.Type = featureStateTypeConstant
		envelope.Value = v
	case featureStateRandom:
		envelope.Type = featureStateTypeRandom
		envelope.Value = v
	case featureStateSegment:
		envelope.Type = featureStateTypeSegment
		envelope.Value = v
	default:
		return nil, fmt.Errorf("bad featureState type: %T", v)
	}
	return json.Marshal(envelope)
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *featureStateTransport) UnmarshalJSON(b []byte) error {
	var envelope struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(b, &envelope); err != nil {
		return err
	}
	switch envelope.Type {
	case featureStateTypeConstant:
		var b bool
		if err := json.Unmarshal(envelope.Value, &b); err != nil {
			return err
		}
		f.value = featureStateConstant(b)
		return nil
	case featureStateTypeRandom:
		var r uint8
		if err := json.Unmarshal(envelope.Value, &r); err != nil {
			return err
		}
		f.value = featureStateRandom(r)
		return nil
	case featureStateTypeSegment:
		var m map[string]string
		if err := json.Unmarshal(envelope.Value, &m); err != nil {
			return err
		}
		f.value = featureStateSegment(m)
		return nil
	default:
		return fmt.Errorf("bad featureState type: %s", envelope.Type)
	}
}

// Value implements driver.Valuer.
func (f featureStateTransport) Value() (driver.Value, error) {
	return f.MarshalJSON()
}

// Scan implements sql.Scanner.
func (f *featureStateTransport) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		f.value = featureStateConstant(false)
		return nil
	case []byte:
		return f.UnmarshalJSON(v)
	default:
		return fmt.Errorf("bad featureState type: %T", v)
	}
}
