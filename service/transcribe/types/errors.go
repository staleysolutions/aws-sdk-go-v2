// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Your request didn't pass one or more validation tests. For example, if the
// entity that you're trying to delete doesn't exist or if it is in a non-terminal
// state (for example, it's "in progress"). See the exception Message field for
// more information.
type BadRequestException struct {
	Message *string
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is already a resource with that name.
type ConflictException struct {
	Message *string
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was an internal error. Check the error message and try your request again.
type InternalFailureException struct {
	Message *string
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string             { return "InternalFailureException" }
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Either you have sent too many requests or your input file is too long. Wait
// before you resend your request, or use a smaller file and resend the request.
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

// We can't find the requested resource. Check the name and try your request again.
type NotFoundException struct {
	Message *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
