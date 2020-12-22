// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// There is concurrent modification on a rule, target, archive, or replay.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred because a replay can be canceled only when the state is
// Running or Starting.
type IllegalStatusException struct {
	Message *string
}

func (e *IllegalStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IllegalStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IllegalStatusException) ErrorCode() string             { return "IllegalStatusException" }
func (e *IllegalStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception occurs due to unexpected causes.
type InternalException struct {
	Message *string
}

func (e *InternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalException) ErrorCode() string             { return "InternalException" }
func (e *InternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The event pattern is not valid.
type InvalidEventPatternException struct {
	Message *string
}

func (e *InvalidEventPatternException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEventPatternException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEventPatternException) ErrorCode() string             { return "InvalidEventPatternException" }
func (e *InvalidEventPatternException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified state is not a valid state for an event source.
type InvalidStateException struct {
	Message *string
}

func (e *InvalidStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidStateException) ErrorCode() string             { return "InvalidStateException" }
func (e *InvalidStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because it attempted to create resource beyond the allowed
// service quota.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This rule was created by an AWS service on behalf of your account. It is managed
// by that service. If you see this error in response to DeleteRule or
// RemoveTargets, you can use the Force parameter in those calls to delete the rule
// or remove targets from the rule. You cannot modify these managed rules by using
// DisableRule, EnableRule, PutTargets, PutRule, TagResource, or UntagResource.
type ManagedRuleException struct {
	Message *string
}

func (e *ManagedRuleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ManagedRuleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ManagedRuleException) ErrorCode() string             { return "ManagedRuleException" }
func (e *ManagedRuleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation you are attempting is not available in this region.
type OperationDisabledException struct {
	Message *string
}

func (e *OperationDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationDisabledException) ErrorCode() string             { return "OperationDisabledException" }
func (e *OperationDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The event bus policy is too long. For more information, see the limits.
type PolicyLengthExceededException struct {
	Message *string
}

func (e *PolicyLengthExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyLengthExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyLengthExceededException) ErrorCode() string             { return "PolicyLengthExceededException" }
func (e *PolicyLengthExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource you are trying to create already exists.
type ResourceAlreadyExistsException struct {
	Message *string
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An entity that you specified does not exist.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
