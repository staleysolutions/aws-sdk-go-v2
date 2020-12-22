// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about an instance as part of a deployment.
//
// Deprecated: This operation is deprecated, use GetDeploymentTarget instead.
func (c *Client) GetDeploymentInstance(ctx context.Context, params *GetDeploymentInstanceInput, optFns ...func(*Options)) (*GetDeploymentInstanceOutput, error) {
	if params == nil {
		params = &GetDeploymentInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeploymentInstance", params, optFns, addOperationGetDeploymentInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetDeploymentInstance operation.
type GetDeploymentInstanceInput struct {

	// The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	// The unique ID of an instance in the deployment group.
	//
	// This member is required.
	InstanceId *string
}

// Represents the output of a GetDeploymentInstance operation.
type GetDeploymentInstanceOutput struct {

	// Information about the instance.
	//
	// Deprecated: InstanceSummary is deprecated, use DeploymentTarget instead.
	InstanceSummary *types.InstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDeploymentInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeploymentInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeploymentInstance{}, middleware.After)
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
	if err = addOpGetDeploymentInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeploymentInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeploymentInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetDeploymentInstance",
	}
}
