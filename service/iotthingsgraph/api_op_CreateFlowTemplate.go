// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a workflow template. Workflows can be created only in the user's
// namespace. (The public namespace contains only entities.) The workflow can
// contain only entities in the specified namespace. The workflow is validated
// against the entities in the latest version of the user's namespace unless
// another namespace version is specified in the request.
func (c *Client) CreateFlowTemplate(ctx context.Context, params *CreateFlowTemplateInput, optFns ...func(*Options)) (*CreateFlowTemplateOutput, error) {
	if params == nil {
		params = &CreateFlowTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlowTemplate", params, optFns, addOperationCreateFlowTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowTemplateInput struct {

	// The workflow DefinitionDocument.
	//
	// This member is required.
	Definition *types.DefinitionDocument

	// The namespace version in which the workflow is to be created. If no value is
	// specified, the latest version is used by default.
	CompatibleNamespaceVersion *int64
}

type CreateFlowTemplateOutput struct {

	// The summary object that describes the created workflow.
	Summary *types.FlowTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFlowTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFlowTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFlowTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateFlowTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlowTemplate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateFlowTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "CreateFlowTemplate",
	}
}
