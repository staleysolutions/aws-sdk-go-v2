// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the permissions on a dataset. The permissions resource is
// arn:aws:quicksight:region:aws-account-id:dataset/data-set-id.
func (c *Client) DescribeDataSetPermissions(ctx context.Context, params *DescribeDataSetPermissionsInput, optFns ...func(*Options)) (*DescribeDataSetPermissionsOutput, error) {
	if params == nil {
		params = &DescribeDataSetPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataSetPermissions", params, optFns, addOperationDescribeDataSetPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataSetPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataSetPermissionsInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dataset that you want to create. This ID is unique per AWS Region
	// for each AWS account.
	//
	// This member is required.
	DataSetId *string
}

type DescribeDataSetPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	DataSetArn *string

	// The ID for the dataset that you want to create. This ID is unique per AWS Region
	// for each AWS account.
	DataSetId *string

	// A list of resource permissions on the dataset.
	Permissions []types.ResourcePermission

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDataSetPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDataSetPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDataSetPermissions{}, middleware.After)
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
	if err = addOpDescribeDataSetPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataSetPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDataSetPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeDataSetPermissions",
	}
}
