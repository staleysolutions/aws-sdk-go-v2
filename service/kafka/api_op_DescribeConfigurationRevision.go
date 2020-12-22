// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of this revision of the configuration.
func (c *Client) DescribeConfigurationRevision(ctx context.Context, params *DescribeConfigurationRevisionInput, optFns ...func(*Options)) (*DescribeConfigurationRevisionOutput, error) {
	if params == nil {
		params = &DescribeConfigurationRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationRevision", params, optFns, addOperationDescribeConfigurationRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationRevisionInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies an MSK configuration and
	// all of its revisions.
	//
	// This member is required.
	Arn *string

	// A string that uniquely identifies a revision of an MSK configuration.
	//
	// This member is required.
	Revision int64
}

type DescribeConfigurationRevisionOutput struct {

	// The Amazon Resource Name (ARN) of the configuration.
	Arn *string

	// The time when the configuration was created.
	CreationTime *time.Time

	// The description of the configuration.
	Description *string

	// The revision number.
	Revision int64

	// Contents of the server.properties file. When using the API, you must ensure that
	// the contents of the file are base64 encoded. When using the AWS Management
	// Console, the SDK, or the AWS CLI, the contents of server.properties can be in
	// plaintext.
	ServerProperties []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeConfigurationRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConfigurationRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConfigurationRevision{}, middleware.After)
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
	if err = addOpDescribeConfigurationRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationRevision(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfigurationRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "DescribeConfigurationRevision",
	}
}
