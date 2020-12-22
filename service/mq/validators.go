// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateBroker struct {
}

func (*validateOpCreateBroker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBrokerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTags struct {
}

func (*validateOpCreateTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateUser struct {
}

func (*validateOpCreateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteBroker struct {
}

func (*validateOpDeleteBroker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteBrokerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTags struct {
}

func (*validateOpDeleteTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteUser struct {
}

func (*validateOpDeleteUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeBroker struct {
}

func (*validateOpDescribeBroker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeBrokerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeConfiguration struct {
}

func (*validateOpDescribeConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeConfigurationRevision struct {
}

func (*validateOpDescribeConfigurationRevision) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeConfigurationRevision) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeConfigurationRevisionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeConfigurationRevisionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeUser struct {
}

func (*validateOpDescribeUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListConfigurationRevisions struct {
}

func (*validateOpListConfigurationRevisions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListConfigurationRevisions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListConfigurationRevisionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListConfigurationRevisionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTags struct {
}

func (*validateOpListTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListUsers struct {
}

func (*validateOpListUsers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListUsers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListUsersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListUsersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRebootBroker struct {
}

func (*validateOpRebootBroker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRebootBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RebootBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRebootBrokerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateBroker struct {
}

func (*validateOpUpdateBroker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateBrokerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateConfiguration struct {
}

func (*validateOpUpdateConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateUser struct {
}

func (*validateOpUpdateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateBrokerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBroker{}, middleware.After)
}

func addOpCreateTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTags{}, middleware.After)
}

func addOpCreateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateUser{}, middleware.After)
}

func addOpDeleteBrokerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteBroker{}, middleware.After)
}

func addOpDeleteTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTags{}, middleware.After)
}

func addOpDeleteUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteUser{}, middleware.After)
}

func addOpDescribeBrokerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeBroker{}, middleware.After)
}

func addOpDescribeConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeConfiguration{}, middleware.After)
}

func addOpDescribeConfigurationRevisionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeConfigurationRevision{}, middleware.After)
}

func addOpDescribeUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeUser{}, middleware.After)
}

func addOpListConfigurationRevisionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListConfigurationRevisions{}, middleware.After)
}

func addOpListTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTags{}, middleware.After)
}

func addOpListUsersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListUsers{}, middleware.After)
}

func addOpRebootBrokerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRebootBroker{}, middleware.After)
}

func addOpUpdateBrokerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateBroker{}, middleware.After)
}

func addOpUpdateConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateConfiguration{}, middleware.After)
}

func addOpUpdateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateUser{}, middleware.After)
}

func validateEncryptionOptions(v *types.EncryptionOptions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EncryptionOptions"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateBrokerInput(v *CreateBrokerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBrokerInput"}
	if v.EncryptionOptions != nil {
		if err := validateEncryptionOptions(v.EncryptionOptions); err != nil {
			invalidParams.AddNested("EncryptionOptions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTagsInput(v *CreateTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTagsInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateUserInput(v *CreateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateUserInput"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteBrokerInput(v *DeleteBrokerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteBrokerInput"}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTagsInput(v *DeleteTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTagsInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteUserInput(v *DeleteUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteUserInput"}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeBrokerInput(v *DescribeBrokerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeBrokerInput"}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeConfigurationInput(v *DescribeConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeConfigurationInput"}
	if v.ConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeConfigurationRevisionInput(v *DescribeConfigurationRevisionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeConfigurationRevisionInput"}
	if v.ConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationId"))
	}
	if v.ConfigurationRevision == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationRevision"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeUserInput(v *DescribeUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeUserInput"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListConfigurationRevisionsInput(v *ListConfigurationRevisionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListConfigurationRevisionsInput"}
	if v.ConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsInput(v *ListTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListUsersInput(v *ListUsersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListUsersInput"}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRebootBrokerInput(v *RebootBrokerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RebootBrokerInput"}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateBrokerInput(v *UpdateBrokerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBrokerInput"}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateConfigurationInput(v *UpdateConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateConfigurationInput"}
	if v.ConfigurationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateUserInput(v *UpdateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateUserInput"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.BrokerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BrokerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
