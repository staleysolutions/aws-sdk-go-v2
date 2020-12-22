// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists tags for an Amazon FSx file systems and backups in the case of Amazon FSx
// for Windows File Server. When retrieving all tags, you can optionally specify
// the MaxResults parameter to limit the number of tags in a response. If more tags
// remain, Amazon FSx returns a NextToken value in the response. In this case, send
// a later request with the NextToken request parameter set to the value of
// NextToken from the last response. This action is used in an iterative process to
// retrieve a list of your tags. ListTagsForResource is called first without a
// NextTokenvalue. Then the action continues to be called with the NextToken
// parameter set to the value of the last NextToken value until a response has no
// NextToken. When using this action, keep the following in mind:
//
// * The
// implementation might return fewer than MaxResults file system descriptions while
// still including a NextToken value.
//
// * The order of tags returned in the response
// of one ListTagsForResource call and the order of tags returned across the
// responses of a multi-call iteration is unspecified.
func (c *Client) ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResource", params, optFns, addOperationListTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for ListTagsForResource operation.
type ListTagsForResourceInput struct {

	// The ARN of the Amazon FSx resource that will have its tags listed.
	//
	// This member is required.
	ResourceARN *string

	// Maximum number of tags to return in the response (integer). This parameter value
	// must be greater than 0. The number of items that Amazon FSx returns is the
	// minimum of the MaxResults parameter specified in the request and the service's
	// internal maximum number of items per page.
	MaxResults *int32

	// Opaque pagination token returned from a previous ListTagsForResource operation
	// (String). If a token present, the action continues the list from where the
	// returning call left off.
	NextToken *string
}

// The response object for ListTagsForResource operation.
type ListTagsForResourceOutput struct {

	// This is present if there are more tags than returned in the response (String).
	// You can use the NextToken value in the later request to fetch the tags.
	NextToken *string

	// A list of tags on the resource.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsForResource{}, middleware.After)
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
	if err = addOpListTagsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "ListTagsForResource",
	}
}
