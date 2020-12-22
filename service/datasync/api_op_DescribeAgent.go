// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata such as the name, the network interfaces, and the status (that
// is, whether the agent is running or not) for an agent. To specify which agent to
// describe, use the Amazon Resource Name (ARN) of the agent in your request.
func (c *Client) DescribeAgent(ctx context.Context, params *DescribeAgentInput, optFns ...func(*Options)) (*DescribeAgentOutput, error) {
	if params == nil {
		params = &DescribeAgentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAgent", params, optFns, addOperationDescribeAgentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAgentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeAgent
type DescribeAgentInput struct {

	// The Amazon Resource Name (ARN) of the agent to describe.
	//
	// This member is required.
	AgentArn *string
}

// DescribeAgentResponse
type DescribeAgentOutput struct {

	// The Amazon Resource Name (ARN) of the agent.
	AgentArn *string

	// The time that the agent was activated (that is, created in your account).
	CreationTime *time.Time

	// The type of endpoint that your agent is connected to. If the endpoint is a VPC
	// endpoint, the agent is not accessible over the public internet.
	EndpointType types.EndpointType

	// The time that the agent last connected to DataSyc.
	LastConnectionTime *time.Time

	// The name of the agent.
	Name *string

	// The subnet and the security group that DataSync used to access a VPC endpoint.
	PrivateLinkConfig *types.PrivateLinkConfig

	// The status of the agent. If the status is ONLINE, then the agent is configured
	// properly and is available to use. The Running status is the normal running
	// status for an agent. If the status is OFFLINE, the agent's VM is turned off or
	// the agent is in an unhealthy state. When the issue that caused the unhealthy
	// state is resolved, the agent returns to ONLINE status.
	Status types.AgentStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAgentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAgent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAgent{}, middleware.After)
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
	if err = addOpDescribeAgentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAgent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAgent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeAgent",
	}
}
