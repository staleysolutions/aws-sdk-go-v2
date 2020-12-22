// Code generated by smithy-go-codegen DO NOT EDIT.

package pricing

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGetAttributeValues struct {
}

func (*validateOpGetAttributeValues) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAttributeValues) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAttributeValuesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAttributeValuesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProducts struct {
}

func (*validateOpGetProducts) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProducts) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProductsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProductsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetAttributeValuesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAttributeValues{}, middleware.After)
}

func addOpGetProductsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProducts{}, middleware.After)
}

func validateFilter(v *types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filter"}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if v.Field == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Field"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilters(v []types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filters"}
	for i := range v {
		if err := validateFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAttributeValuesInput(v *GetAttributeValuesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAttributeValuesInput"}
	if v.AttributeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeName"))
	}
	if v.ServiceCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProductsInput(v *GetProductsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProductsInput"}
	if v.Filters != nil {
		if err := validateFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
