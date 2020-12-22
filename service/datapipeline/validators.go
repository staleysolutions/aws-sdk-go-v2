// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpActivatePipeline struct {
}

func (*validateOpActivatePipeline) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpActivatePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ActivatePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpActivatePipelineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAddTags struct {
}

func (*validateOpAddTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePipeline struct {
}

func (*validateOpCreatePipeline) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePipelineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeactivatePipeline struct {
}

func (*validateOpDeactivatePipeline) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeactivatePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeactivatePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeactivatePipelineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePipeline struct {
}

func (*validateOpDeletePipeline) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePipeline) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePipelineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePipelineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeObjects struct {
}

func (*validateOpDescribeObjects) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeObjects) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeObjectsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeObjectsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribePipelines struct {
}

func (*validateOpDescribePipelines) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribePipelines) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribePipelinesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribePipelinesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEvaluateExpression struct {
}

func (*validateOpEvaluateExpression) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEvaluateExpression) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EvaluateExpressionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEvaluateExpressionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPipelineDefinition struct {
}

func (*validateOpGetPipelineDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPipelineDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPipelineDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPipelineDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPollForTask struct {
}

func (*validateOpPollForTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPollForTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PollForTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPollForTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutPipelineDefinition struct {
}

func (*validateOpPutPipelineDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutPipelineDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutPipelineDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutPipelineDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpQueryObjects struct {
}

func (*validateOpQueryObjects) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpQueryObjects) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*QueryObjectsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpQueryObjectsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveTags struct {
}

func (*validateOpRemoveTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpReportTaskProgress struct {
}

func (*validateOpReportTaskProgress) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpReportTaskProgress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ReportTaskProgressInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpReportTaskProgressInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpReportTaskRunnerHeartbeat struct {
}

func (*validateOpReportTaskRunnerHeartbeat) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpReportTaskRunnerHeartbeat) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ReportTaskRunnerHeartbeatInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpReportTaskRunnerHeartbeatInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetStatus struct {
}

func (*validateOpSetStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetTaskStatus struct {
}

func (*validateOpSetTaskStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetTaskStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetTaskStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetTaskStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpValidatePipelineDefinition struct {
}

func (*validateOpValidatePipelineDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpValidatePipelineDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ValidatePipelineDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpValidatePipelineDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpActivatePipelineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpActivatePipeline{}, middleware.After)
}

func addOpAddTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddTags{}, middleware.After)
}

func addOpCreatePipelineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePipeline{}, middleware.After)
}

func addOpDeactivatePipelineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeactivatePipeline{}, middleware.After)
}

func addOpDeletePipelineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePipeline{}, middleware.After)
}

func addOpDescribeObjectsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeObjects{}, middleware.After)
}

func addOpDescribePipelinesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribePipelines{}, middleware.After)
}

func addOpEvaluateExpressionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEvaluateExpression{}, middleware.After)
}

func addOpGetPipelineDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPipelineDefinition{}, middleware.After)
}

func addOpPollForTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPollForTask{}, middleware.After)
}

func addOpPutPipelineDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutPipelineDefinition{}, middleware.After)
}

func addOpQueryObjectsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpQueryObjects{}, middleware.After)
}

func addOpRemoveTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveTags{}, middleware.After)
}

func addOpReportTaskProgressValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpReportTaskProgress{}, middleware.After)
}

func addOpReportTaskRunnerHeartbeatValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpReportTaskRunnerHeartbeat{}, middleware.After)
}

func addOpSetStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetStatus{}, middleware.After)
}

func addOpSetTaskStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetTaskStatus{}, middleware.After)
}

func addOpValidatePipelineDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpValidatePipelineDefinition{}, middleware.After)
}

func validateField(v *types.Field) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Field"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFieldList(v []types.Field) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FieldList"}
	for i := range v {
		if err := validateField(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateParameterAttribute(v *types.ParameterAttribute) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ParameterAttribute"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.StringValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StringValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateParameterAttributeList(v []types.ParameterAttribute) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ParameterAttributeList"}
	for i := range v {
		if err := validateParameterAttribute(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateParameterObject(v *types.ParameterObject) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ParameterObject"}
	if v.Attributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attributes"))
	} else if v.Attributes != nil {
		if err := validateParameterAttributeList(v.Attributes); err != nil {
			invalidParams.AddNested("Attributes", err.(smithy.InvalidParamsError))
		}
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateParameterObjectList(v []types.ParameterObject) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ParameterObjectList"}
	for i := range v {
		if err := validateParameterObject(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateParameterValue(v *types.ParameterValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ParameterValue"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.StringValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StringValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateParameterValueList(v []types.ParameterValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ParameterValueList"}
	for i := range v {
		if err := validateParameterValue(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipelineObject(v *types.PipelineObject) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipelineObject"}
	if v.Fields == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Fields"))
	} else if v.Fields != nil {
		if err := validateFieldList(v.Fields); err != nil {
			invalidParams.AddNested("Fields", err.(smithy.InvalidParamsError))
		}
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePipelineObjectList(v []types.PipelineObject) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PipelineObjectList"}
	for i := range v {
		if err := validatePipelineObject(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpActivatePipelineInput(v *ActivatePipelineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActivatePipelineInput"}
	if v.ParameterValues != nil {
		if err := validateParameterValueList(v.ParameterValues); err != nil {
			invalidParams.AddNested("ParameterValues", err.(smithy.InvalidParamsError))
		}
	}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddTagsInput(v *AddTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddTagsInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePipelineInput(v *CreatePipelineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePipelineInput"}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.UniqueId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UniqueId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeactivatePipelineInput(v *DeactivatePipelineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeactivatePipelineInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePipelineInput(v *DeletePipelineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePipelineInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeObjectsInput(v *DescribeObjectsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeObjectsInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if v.ObjectIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObjectIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribePipelinesInput(v *DescribePipelinesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribePipelinesInput"}
	if v.PipelineIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEvaluateExpressionInput(v *EvaluateExpressionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EvaluateExpressionInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if v.Expression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Expression"))
	}
	if v.ObjectId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObjectId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPipelineDefinitionInput(v *GetPipelineDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPipelineDefinitionInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPollForTaskInput(v *PollForTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PollForTaskInput"}
	if v.WorkerGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkerGroup"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutPipelineDefinitionInput(v *PutPipelineDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutPipelineDefinitionInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if v.ParameterObjects != nil {
		if err := validateParameterObjectList(v.ParameterObjects); err != nil {
			invalidParams.AddNested("ParameterObjects", err.(smithy.InvalidParamsError))
		}
	}
	if v.ParameterValues != nil {
		if err := validateParameterValueList(v.ParameterValues); err != nil {
			invalidParams.AddNested("ParameterValues", err.(smithy.InvalidParamsError))
		}
	}
	if v.PipelineObjects == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineObjects"))
	} else if v.PipelineObjects != nil {
		if err := validatePipelineObjectList(v.PipelineObjects); err != nil {
			invalidParams.AddNested("PipelineObjects", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpQueryObjectsInput(v *QueryObjectsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueryObjectsInput"}
	if v.Sphere == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sphere"))
	}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveTagsInput(v *RemoveTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveTagsInput"}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpReportTaskProgressInput(v *ReportTaskProgressInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReportTaskProgressInput"}
	if v.TaskId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskId"))
	}
	if v.Fields != nil {
		if err := validateFieldList(v.Fields); err != nil {
			invalidParams.AddNested("Fields", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpReportTaskRunnerHeartbeatInput(v *ReportTaskRunnerHeartbeatInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReportTaskRunnerHeartbeatInput"}
	if v.TaskrunnerId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskrunnerId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetStatusInput(v *SetStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetStatusInput"}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if v.Status == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Status"))
	}
	if v.ObjectIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObjectIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetTaskStatusInput(v *SetTaskStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetTaskStatusInput"}
	if v.TaskId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskId"))
	}
	if len(v.TaskStatus) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("TaskStatus"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpValidatePipelineDefinitionInput(v *ValidatePipelineDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ValidatePipelineDefinitionInput"}
	if v.ParameterValues != nil {
		if err := validateParameterValueList(v.ParameterValues); err != nil {
			invalidParams.AddNested("ParameterValues", err.(smithy.InvalidParamsError))
		}
	}
	if v.ParameterObjects != nil {
		if err := validateParameterObjectList(v.ParameterObjects); err != nil {
			invalidParams.AddNested("ParameterObjects", err.(smithy.InvalidParamsError))
		}
	}
	if v.PipelineId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineId"))
	}
	if v.PipelineObjects == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PipelineObjects"))
	} else if v.PipelineObjects != nil {
		if err := validatePipelineObjectList(v.PipelineObjects); err != nil {
			invalidParams.AddNested("PipelineObjects", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
