// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// A resource to be created or added already exists.
type AlreadyExistsException struct {
	Message *string
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string             { return "AlreadyExistsException" }
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Two processes are trying to modify a resource simultaneously.
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

// A specified entity does not exist
type EntityNotFoundException struct {
	Message *string
}

func (e *EntityNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityNotFoundException) ErrorCode() string             { return "EntityNotFoundException" }
func (e *EntityNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal service error occurred.
type InternalServiceException struct {
	Message *string
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The input provided was not valid.
type InvalidInputException struct {
	Message *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation timed out.
type OperationTimeoutException struct {
	Message *string
}

func (e *OperationTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationTimeoutException) ErrorCode() string             { return "OperationTimeoutException" }
func (e *OperationTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
