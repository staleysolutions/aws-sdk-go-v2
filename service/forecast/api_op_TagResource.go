// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified tags to a resource with the specified resourceArn. If
// existing tags on a resource are not specified in the request parameters, they
// are not changed. When a resource is deleted, the tags associated with that
// resource are also deleted.
func (c *Client) TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) {
	if params == nil {
		params = &TagResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagResource", params, optFns, addOperationTagResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagResourceInput struct {

	// The Amazon Resource Name (ARN) that identifies the resource for which to list
	// the tags. Currently, the supported resources are Forecast dataset groups,
	// datasets, dataset import jobs, predictors, forecasts, and forecast export jobs.
	//
	// This member is required.
	ResourceArn *string

	// The tags to add to the resource. A tag is an array of key-value pairs. The
	// following basic restrictions apply to tags:
	//
	// * Maximum number of tags per
	// resource - 50.
	//
	// * For each resource, each tag key must be unique, and each tag
	// key can have only one value.
	//
	// * Maximum key length - 128 Unicode characters in
	// UTF-8.
	//
	// * Maximum value length - 256 Unicode characters in UTF-8.
	//
	// * If your
	// tagging schema is used across multiple services and resources, remember that
	// other services may have restrictions on allowed characters. Generally allowed
	// characters are: letters, numbers, and spaces representable in UTF-8, and the
	// following characters: + - = . _ : / @.
	//
	// * Tag keys and values are case
	// sensitive.
	//
	// * Do not use aws:, AWS:, or any upper or lowercase combination of
	// such as a prefix for keys as it is reserved for AWS use. You cannot edit or
	// delete tag keys with this prefix. Values can have this prefix. If a tag value
	// has aws as its prefix but the key does not, then Forecast considers it to be a
	// user tag and will count against the limit of 50 tags. Tags with only the key
	// prefix of aws do not count against your tags per resource limit.
	//
	// This member is required.
	Tags []types.Tag
}

type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagResource{}, middleware.After)
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
	if err = addOpTagResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "TagResource",
	}
}
