// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/honeycode/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGetScreenData struct {
}

func (*validateOpGetScreenData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetScreenData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetScreenDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetScreenDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInvokeScreenAutomation struct {
}

func (*validateOpInvokeScreenAutomation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInvokeScreenAutomation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InvokeScreenAutomationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInvokeScreenAutomationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetScreenDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetScreenData{}, middleware.After)
}

func addOpInvokeScreenAutomationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInvokeScreenAutomation{}, middleware.After)
}

func validateVariableValue(v *types.VariableValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableValue"}
	if v.RawValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RawValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVariableValueMap(v map[string]types.VariableValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableValueMap"}
	for key := range v {
		value := v[key]
		if err := validateVariableValue(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetScreenDataInput(v *GetScreenDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetScreenDataInput"}
	if v.Variables != nil {
		if err := validateVariableValueMap(v.Variables); err != nil {
			invalidParams.AddNested("Variables", err.(smithy.InvalidParamsError))
		}
	}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.ScreenId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScreenId"))
	}
	if v.WorkbookId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInvokeScreenAutomationInput(v *InvokeScreenAutomationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvokeScreenAutomationInput"}
	if v.WorkbookId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.ScreenId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScreenId"))
	}
	if v.AppId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.ScreenAutomationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScreenAutomationId"))
	}
	if v.Variables != nil {
		if err := validateVariableValueMap(v.Variables); err != nil {
			invalidParams.AddNested("Variables", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
