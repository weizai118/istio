// Code generated by protoc-gen-validate
// source: envoy/config/overload/v2alpha/overload.proto
// DO NOT EDIT!!!

package v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on ResourceMonitor with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ResourceMonitor) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return ResourceMonitorValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceMonitorValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// ResourceMonitorValidationError is the validation error returned by
// ResourceMonitor.Validate if the designated constraints aren't met.
type ResourceMonitorValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ResourceMonitorValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResourceMonitor.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ResourceMonitorValidationError{}

// Validate checks the field values on ThresholdTrigger with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ThresholdTrigger) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetValue(); val < 0 || val > 1 {
		return ThresholdTriggerValidationError{
			Field:  "Value",
			Reason: "value must be inside range [0, 1]",
		}
	}

	return nil
}

// ThresholdTriggerValidationError is the validation error returned by
// ThresholdTrigger.Validate if the designated constraints aren't met.
type ThresholdTriggerValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ThresholdTriggerValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThresholdTrigger.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ThresholdTriggerValidationError{}

// Validate checks the field values on Trigger with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Trigger) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return TriggerValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	switch m.TriggerOneof.(type) {

	case *Trigger_Threshold:

		if v, ok := interface{}(m.GetThreshold()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TriggerValidationError{
					Field:  "Threshold",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return TriggerValidationError{
			Field:  "TriggerOneof",
			Reason: "value is required",
		}

	}

	return nil
}

// TriggerValidationError is the validation error returned by Trigger.Validate
// if the designated constraints aren't met.
type TriggerValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TriggerValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrigger.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TriggerValidationError{}

// Validate checks the field values on OverloadAction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OverloadAction) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return OverloadActionValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetTriggers()) < 1 {
		return OverloadActionValidationError{
			Field:  "Triggers",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetTriggers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverloadActionValidationError{
					Field:  fmt.Sprintf("Triggers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// OverloadActionValidationError is the validation error returned by
// OverloadAction.Validate if the designated constraints aren't met.
type OverloadActionValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e OverloadActionValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverloadAction.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = OverloadActionValidationError{}

// Validate checks the field values on OverloadManager with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OverloadManager) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRefreshInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OverloadManagerValidationError{
				Field:  "RefreshInterval",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if len(m.GetResourceMonitors()) < 1 {
		return OverloadManagerValidationError{
			Field:  "ResourceMonitors",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetResourceMonitors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverloadManagerValidationError{
					Field:  fmt.Sprintf("ResourceMonitors[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverloadManagerValidationError{
					Field:  fmt.Sprintf("Actions[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// OverloadManagerValidationError is the validation error returned by
// OverloadManager.Validate if the designated constraints aren't met.
type OverloadManagerValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e OverloadManagerValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverloadManager.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = OverloadManagerValidationError{}
