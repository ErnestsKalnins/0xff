package feature

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Feature struct {
	ID            uuid.UUID `json:"id"`
	ProjectID     uuid.UUID `json:"projectId"`
	TechnicalName string    `json:"technicalName"`
	DisplayName   *string   `json:"displayName,omitempty"`
	Description   *string   `json:"description,omitempty"`
	CreatedAt     int64     `json:"createdAt"`
	UpdatedAt     int64     `json:"updatedAt"`
}

type ErrFeatureNotFound struct {
	ID uuid.UUID
}

func (e ErrFeatureNotFound) Error() string {
	return fmt.Sprintf("could not find feature by id %s", e.ID)
}

type EnvironmentFeature struct {
	ID            uuid.UUID
	FeatureID     uuid.UUID `json:"id"`
	EnvironmentID uuid.UUID
	TechnicalName string  `json:"technicalName"`
	DisplayName   *string `json:"displayName,omitempty"`
	Description   *string `json:"description,omitempty"`
	State         State   `json:"state"`
	CreatedAt     int64   `json:"createdAt"`
	UpdatedAt     int64   `json:"updatedAt"`
}

// State is a sum-type of all possible Feature states.
type State interface {
	state()
}

// All possible featureState types.
const (
	stateConstant = "constant"
	stateRandom   = "random"
	stateSegment  = "segment"
)

// StateConstant is the simplest state - either the feature is on or off.
type StateConstant bool

func (StateConstant) state() {}

// StateRandom is the probability of the feature being on.
type StateRandom uint8

func (StateRandom) state() {}

// StateSegment contains the key values which have to be present for the
// feature to be on.
type StateSegment map[string]string

func (StateSegment) state() {}

// StateMarshaler wraps a State to implement un/marshaling interfaces for JSON.
type StateMarshaler struct {
	Value State
}

// MarshalJSON implements json.Marshaler.
func (m StateMarshaler) MarshalJSON() ([]byte, error) {
	var envelope struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	}
	switch v := m.Value.(type) {
	case StateConstant:
		envelope.Type = stateConstant
		envelope.Value = v
	case StateRandom:
		envelope.Type = stateRandom
		envelope.Value = v
	case StateSegment:
		envelope.Type = stateSegment
		envelope.Value = v
	default:
		return nil, fmt.Errorf("bad featureState type: %T", v)
	}
	return json.Marshal(envelope)
}

// UnmarshalJSON implements json.Unmarshaler.
func (m *StateMarshaler) UnmarshalJSON(b []byte) error {
	var envelope struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(b, &envelope); err != nil {
		return err
	}
	switch envelope.Type {
	case stateConstant:
		var b bool
		if err := json.Unmarshal(envelope.Value, &b); err != nil {
			return err
		}
		m.Value = StateConstant(b)
		return nil
	case stateRandom:
		var r uint8
		if err := json.Unmarshal(envelope.Value, &r); err != nil {
			return err
		}
		m.Value = StateRandom(r)
		return nil
	case stateSegment:
		var mm map[string]string
		if err := json.Unmarshal(envelope.Value, &mm); err != nil {
			return err
		}
		m.Value = StateSegment(mm)
		return nil
	default:
		return fmt.Errorf("bad featureState type: %s", envelope.Type)
	}
}

// Schedule is a sum-type of all possible schedules.
type Schedule interface {
	schedule()
}

// All possible schedule types.
const (
	scheduleNone     = "none"
	scheduleFrom     = "from"
	scheduleTo       = "to"
	scheduleInterval = "interval"
)

// ScheduleNone indicates that the feature is always set to its featureState.
type ScheduleNone struct{}

func (ScheduleNone) schedule() {}

// ScheduleFrom indicates that the feature is set to its featureState starting
// from the contained time.
type ScheduleFrom int64

func (ScheduleFrom) schedule() {}

// ScheduleTo indicates that the feature is set to its featureState until the
// contained time.
type ScheduleTo int64

func (ScheduleTo) schedule() {}

// ScheduleInterval indicates that the feature is set to its featureState in the
// contained time interval.
type ScheduleInterval struct {
	From int64
	To   int64
}

func (ScheduleInterval) schedule() {}
