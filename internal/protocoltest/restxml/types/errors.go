// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This error is thrown when a request is invalid.
type ComplexError struct {
	Message *string

	Header   *string
	TopLevel *string
	Nested   *ComplexNestedErrorData
}

func (e *ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ComplexError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ComplexError) ErrorCode() string             { return "ComplexError" }
func (e *ComplexError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error is thrown when an invalid greeting value is provided.
type InvalidGreeting struct {
	Message *string
}

func (e *InvalidGreeting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGreeting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGreeting) ErrorCode() string             { return "InvalidGreeting" }
func (e *InvalidGreeting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
