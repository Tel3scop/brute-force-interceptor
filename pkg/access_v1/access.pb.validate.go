// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: access.proto

package access_v1

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

// Validate checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthRequestMultiError, or
// nil if none found.
func (m *AuthRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Login

	// no validation rules for Password

	// no validation rules for Ip

	if len(errors) > 0 {
		return AuthRequestMultiError(errors)
	}

	return nil
}

// AuthRequestMultiError is an error wrapping multiple validation errors
// returned by AuthRequest.ValidateAll() if the designated constraints aren't met.
type AuthRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRequestMultiError) AllErrors() []error { return m }

// AuthRequestValidationError is the validation error returned by
// AuthRequest.Validate if the designated constraints aren't met.
type AuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRequestValidationError) ErrorName() string { return "AuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRequestValidationError{}

// Validate checks the field values on AuthResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthResponseMultiError, or
// nil if none found.
func (m *AuthResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return AuthResponseMultiError(errors)
	}

	return nil
}

// AuthResponseMultiError is an error wrapping multiple validation errors
// returned by AuthResponse.ValidateAll() if the designated constraints aren't met.
type AuthResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthResponseMultiError) AllErrors() []error { return m }

// AuthResponseValidationError is the validation error returned by
// AuthResponse.Validate if the designated constraints aren't met.
type AuthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthResponseValidationError) ErrorName() string { return "AuthResponseValidationError" }

// Error satisfies the builtin error interface
func (e AuthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthResponseValidationError{}

// Validate checks the field values on ResetBucketRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResetBucketRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetBucketRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResetBucketRequestMultiError, or nil if none found.
func (m *ResetBucketRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetBucketRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Login

	// no validation rules for Ip

	if len(errors) > 0 {
		return ResetBucketRequestMultiError(errors)
	}

	return nil
}

// ResetBucketRequestMultiError is an error wrapping multiple validation errors
// returned by ResetBucketRequest.ValidateAll() if the designated constraints
// aren't met.
type ResetBucketRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetBucketRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetBucketRequestMultiError) AllErrors() []error { return m }

// ResetBucketRequestValidationError is the validation error returned by
// ResetBucketRequest.Validate if the designated constraints aren't met.
type ResetBucketRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetBucketRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetBucketRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetBucketRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetBucketRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetBucketRequestValidationError) ErrorName() string {
	return "ResetBucketRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ResetBucketRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetBucketRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetBucketRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetBucketRequestValidationError{}

// Validate checks the field values on ResetBucketResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResetBucketResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetBucketResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResetBucketResponseMultiError, or nil if none found.
func (m *ResetBucketResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetBucketResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ResetBucketResponseMultiError(errors)
	}

	return nil
}

// ResetBucketResponseMultiError is an error wrapping multiple validation
// errors returned by ResetBucketResponse.ValidateAll() if the designated
// constraints aren't met.
type ResetBucketResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetBucketResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetBucketResponseMultiError) AllErrors() []error { return m }

// ResetBucketResponseValidationError is the validation error returned by
// ResetBucketResponse.Validate if the designated constraints aren't met.
type ResetBucketResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetBucketResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetBucketResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetBucketResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetBucketResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetBucketResponseValidationError) ErrorName() string {
	return "ResetBucketResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ResetBucketResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetBucketResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetBucketResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetBucketResponseValidationError{}

// Validate checks the field values on AddToBlacklistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddToBlacklistRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddToBlacklistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddToBlacklistRequestMultiError, or nil if none found.
func (m *AddToBlacklistRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddToBlacklistRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Subnet

	if len(errors) > 0 {
		return AddToBlacklistRequestMultiError(errors)
	}

	return nil
}

// AddToBlacklistRequestMultiError is an error wrapping multiple validation
// errors returned by AddToBlacklistRequest.ValidateAll() if the designated
// constraints aren't met.
type AddToBlacklistRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddToBlacklistRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddToBlacklistRequestMultiError) AllErrors() []error { return m }

// AddToBlacklistRequestValidationError is the validation error returned by
// AddToBlacklistRequest.Validate if the designated constraints aren't met.
type AddToBlacklistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddToBlacklistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddToBlacklistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddToBlacklistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddToBlacklistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddToBlacklistRequestValidationError) ErrorName() string {
	return "AddToBlacklistRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddToBlacklistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddToBlacklistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddToBlacklistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddToBlacklistRequestValidationError{}

// Validate checks the field values on AddToBlacklistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddToBlacklistResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddToBlacklistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddToBlacklistResponseMultiError, or nil if none found.
func (m *AddToBlacklistResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddToBlacklistResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AddToBlacklistResponseMultiError(errors)
	}

	return nil
}

// AddToBlacklistResponseMultiError is an error wrapping multiple validation
// errors returned by AddToBlacklistResponse.ValidateAll() if the designated
// constraints aren't met.
type AddToBlacklistResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddToBlacklistResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddToBlacklistResponseMultiError) AllErrors() []error { return m }

// AddToBlacklistResponseValidationError is the validation error returned by
// AddToBlacklistResponse.Validate if the designated constraints aren't met.
type AddToBlacklistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddToBlacklistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddToBlacklistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddToBlacklistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddToBlacklistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddToBlacklistResponseValidationError) ErrorName() string {
	return "AddToBlacklistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddToBlacklistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddToBlacklistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddToBlacklistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddToBlacklistResponseValidationError{}

// Validate checks the field values on RemoveFromBlacklistRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveFromBlacklistRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveFromBlacklistRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveFromBlacklistRequestMultiError, or nil if none found.
func (m *RemoveFromBlacklistRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveFromBlacklistRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Subnet

	if len(errors) > 0 {
		return RemoveFromBlacklistRequestMultiError(errors)
	}

	return nil
}

// RemoveFromBlacklistRequestMultiError is an error wrapping multiple
// validation errors returned by RemoveFromBlacklistRequest.ValidateAll() if
// the designated constraints aren't met.
type RemoveFromBlacklistRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveFromBlacklistRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveFromBlacklistRequestMultiError) AllErrors() []error { return m }

// RemoveFromBlacklistRequestValidationError is the validation error returned
// by RemoveFromBlacklistRequest.Validate if the designated constraints aren't met.
type RemoveFromBlacklistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveFromBlacklistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveFromBlacklistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveFromBlacklistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveFromBlacklistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveFromBlacklistRequestValidationError) ErrorName() string {
	return "RemoveFromBlacklistRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveFromBlacklistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveFromBlacklistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveFromBlacklistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveFromBlacklistRequestValidationError{}

// Validate checks the field values on RemoveFromBlacklistResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveFromBlacklistResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveFromBlacklistResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveFromBlacklistResponseMultiError, or nil if none found.
func (m *RemoveFromBlacklistResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveFromBlacklistResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RemoveFromBlacklistResponseMultiError(errors)
	}

	return nil
}

// RemoveFromBlacklistResponseMultiError is an error wrapping multiple
// validation errors returned by RemoveFromBlacklistResponse.ValidateAll() if
// the designated constraints aren't met.
type RemoveFromBlacklistResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveFromBlacklistResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveFromBlacklistResponseMultiError) AllErrors() []error { return m }

// RemoveFromBlacklistResponseValidationError is the validation error returned
// by RemoveFromBlacklistResponse.Validate if the designated constraints
// aren't met.
type RemoveFromBlacklistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveFromBlacklistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveFromBlacklistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveFromBlacklistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveFromBlacklistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveFromBlacklistResponseValidationError) ErrorName() string {
	return "RemoveFromBlacklistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveFromBlacklistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveFromBlacklistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveFromBlacklistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveFromBlacklistResponseValidationError{}

// Validate checks the field values on AddToWhitelistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddToWhitelistRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddToWhitelistRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddToWhitelistRequestMultiError, or nil if none found.
func (m *AddToWhitelistRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddToWhitelistRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Subnet

	if len(errors) > 0 {
		return AddToWhitelistRequestMultiError(errors)
	}

	return nil
}

// AddToWhitelistRequestMultiError is an error wrapping multiple validation
// errors returned by AddToWhitelistRequest.ValidateAll() if the designated
// constraints aren't met.
type AddToWhitelistRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddToWhitelistRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddToWhitelistRequestMultiError) AllErrors() []error { return m }

// AddToWhitelistRequestValidationError is the validation error returned by
// AddToWhitelistRequest.Validate if the designated constraints aren't met.
type AddToWhitelistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddToWhitelistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddToWhitelistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddToWhitelistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddToWhitelistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddToWhitelistRequestValidationError) ErrorName() string {
	return "AddToWhitelistRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddToWhitelistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddToWhitelistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddToWhitelistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddToWhitelistRequestValidationError{}

// Validate checks the field values on AddToWhitelistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddToWhitelistResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddToWhitelistResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddToWhitelistResponseMultiError, or nil if none found.
func (m *AddToWhitelistResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddToWhitelistResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AddToWhitelistResponseMultiError(errors)
	}

	return nil
}

// AddToWhitelistResponseMultiError is an error wrapping multiple validation
// errors returned by AddToWhitelistResponse.ValidateAll() if the designated
// constraints aren't met.
type AddToWhitelistResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddToWhitelistResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddToWhitelistResponseMultiError) AllErrors() []error { return m }

// AddToWhitelistResponseValidationError is the validation error returned by
// AddToWhitelistResponse.Validate if the designated constraints aren't met.
type AddToWhitelistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddToWhitelistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddToWhitelistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddToWhitelistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddToWhitelistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddToWhitelistResponseValidationError) ErrorName() string {
	return "AddToWhitelistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddToWhitelistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddToWhitelistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddToWhitelistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddToWhitelistResponseValidationError{}

// Validate checks the field values on RemoveFromWhitelistRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveFromWhitelistRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveFromWhitelistRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveFromWhitelistRequestMultiError, or nil if none found.
func (m *RemoveFromWhitelistRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveFromWhitelistRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Subnet

	if len(errors) > 0 {
		return RemoveFromWhitelistRequestMultiError(errors)
	}

	return nil
}

// RemoveFromWhitelistRequestMultiError is an error wrapping multiple
// validation errors returned by RemoveFromWhitelistRequest.ValidateAll() if
// the designated constraints aren't met.
type RemoveFromWhitelistRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveFromWhitelistRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveFromWhitelistRequestMultiError) AllErrors() []error { return m }

// RemoveFromWhitelistRequestValidationError is the validation error returned
// by RemoveFromWhitelistRequest.Validate if the designated constraints aren't met.
type RemoveFromWhitelistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveFromWhitelistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveFromWhitelistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveFromWhitelistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveFromWhitelistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveFromWhitelistRequestValidationError) ErrorName() string {
	return "RemoveFromWhitelistRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveFromWhitelistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveFromWhitelistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveFromWhitelistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveFromWhitelistRequestValidationError{}

// Validate checks the field values on RemoveFromWhitelistResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveFromWhitelistResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveFromWhitelistResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveFromWhitelistResponseMultiError, or nil if none found.
func (m *RemoveFromWhitelistResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveFromWhitelistResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RemoveFromWhitelistResponseMultiError(errors)
	}

	return nil
}

// RemoveFromWhitelistResponseMultiError is an error wrapping multiple
// validation errors returned by RemoveFromWhitelistResponse.ValidateAll() if
// the designated constraints aren't met.
type RemoveFromWhitelistResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveFromWhitelistResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveFromWhitelistResponseMultiError) AllErrors() []error { return m }

// RemoveFromWhitelistResponseValidationError is the validation error returned
// by RemoveFromWhitelistResponse.Validate if the designated constraints
// aren't met.
type RemoveFromWhitelistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveFromWhitelistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveFromWhitelistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveFromWhitelistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveFromWhitelistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveFromWhitelistResponseValidationError) ErrorName() string {
	return "RemoveFromWhitelistResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveFromWhitelistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveFromWhitelistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveFromWhitelistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveFromWhitelistResponseValidationError{}
