// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/quic/connection_id_generator/v3/envoy_deterministic_connection_id_generator.proto

package connection_id_generatorv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DeterministicConnectionIdGeneratorConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DeterministicConnectionIdGeneratorConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// DeterministicConnectionIdGeneratorConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// DeterministicConnectionIdGeneratorConfigMultiError, or nil if none found.
func (m *DeterministicConnectionIdGeneratorConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DeterministicConnectionIdGeneratorConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeterministicConnectionIdGeneratorConfigMultiError(errors)
	}

	return nil
}

// DeterministicConnectionIdGeneratorConfigMultiError is an error wrapping
// multiple validation errors returned by
// DeterministicConnectionIdGeneratorConfig.ValidateAll() if the designated
// constraints aren't met.
type DeterministicConnectionIdGeneratorConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeterministicConnectionIdGeneratorConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeterministicConnectionIdGeneratorConfigMultiError) AllErrors() []error { return m }

// DeterministicConnectionIdGeneratorConfigValidationError is the validation
// error returned by DeterministicConnectionIdGeneratorConfig.Validate if the
// designated constraints aren't met.
type DeterministicConnectionIdGeneratorConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeterministicConnectionIdGeneratorConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeterministicConnectionIdGeneratorConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeterministicConnectionIdGeneratorConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeterministicConnectionIdGeneratorConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeterministicConnectionIdGeneratorConfigValidationError) ErrorName() string {
	return "DeterministicConnectionIdGeneratorConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DeterministicConnectionIdGeneratorConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeterministicConnectionIdGeneratorConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeterministicConnectionIdGeneratorConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeterministicConnectionIdGeneratorConfigValidationError{}
