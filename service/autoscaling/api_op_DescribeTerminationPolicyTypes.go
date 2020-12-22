// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the termination policies supported by Amazon EC2 Auto Scaling. For
// more information, see Controlling which Auto Scaling instances terminate during
// scale in
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeTerminationPolicyTypes(ctx context.Context, params *DescribeTerminationPolicyTypesInput, optFns ...func(*Options)) (*DescribeTerminationPolicyTypesOutput, error) {
	if params == nil {
		params = &DescribeTerminationPolicyTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTerminationPolicyTypes", params, optFns, addOperationDescribeTerminationPolicyTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTerminationPolicyTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTerminationPolicyTypesInput struct {
}

type DescribeTerminationPolicyTypesOutput struct {

	// The termination policies supported by Amazon EC2 Auto Scaling: OldestInstance,
	// OldestLaunchConfiguration, NewestInstance, ClosestToNextInstanceHour, Default,
	// OldestLaunchTemplate, and AllocationStrategy.
	TerminationPolicyTypes []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTerminationPolicyTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTerminationPolicyTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTerminationPolicyTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTerminationPolicyTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTerminationPolicyTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeTerminationPolicyTypes",
	}
}
