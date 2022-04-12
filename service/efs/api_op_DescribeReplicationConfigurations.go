// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the replication configuration for a specific file system. If a file
// system is not specified, all of the replication configurations for the Amazon
// Web Services account in an Amazon Web Services Region are retrieved.
func (c *Client) DescribeReplicationConfigurations(ctx context.Context, params *DescribeReplicationConfigurationsInput, optFns ...func(*Options)) (*DescribeReplicationConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeReplicationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationConfigurations", params, optFns, c.addOperationDescribeReplicationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationConfigurationsInput struct {

	// You can retrieve the replication configuration for a specific file system by
	// providing its file system ID.
	FileSystemId *string

	// (Optional) To limit the number of objects returned in a response, you can
	// specify the MaxItems parameter. The default value is 100.
	MaxResults *int32

	// NextToken is present if the response is paginated. You can use NextToken in a
	// subsequent request to fetch the next page of output.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeReplicationConfigurationsOutput struct {

	// You can use the NextToken from the previous response in a subsequent request to
	// fetch the additional descriptions.
	NextToken *string

	// The collection of replication configurations that is returned.
	Replications []types.ReplicationConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReplicationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReplicationConfigurations{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReplicationConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeReplicationConfigurations",
	}
}
