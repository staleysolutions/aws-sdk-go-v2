// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Creates a RegexPatternSet, which you reference in a
// RegexPatternSetReferenceStatement, to have AWS WAF inspect a web request
// component for the specified patterns.
func (c *Client) CreateRegexPatternSet(ctx context.Context, params *CreateRegexPatternSetInput, optFns ...func(*Options)) (*CreateRegexPatternSetOutput, error) {
	if params == nil {
		params = &CreateRegexPatternSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegexPatternSet", params, optFns, addOperationCreateRegexPatternSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegexPatternSetInput struct {

	// The name of the set. You cannot change the name after you create the set.
	//
	// This member is required.
	Name *string

	// Array of regular expression strings.
	//
	// This member is required.
	RegularExpressionList []types.Regex

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB),
	// an API Gateway REST API, or an AppSync GraphQL API. To work with CloudFront, you
	// must also specify the Region US East (N. Virginia) as follows:
	//
	// * CLI - Specify
	// the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	// --region=us-east-1.
	//
	// * API and SDKs - For all calls, use the Region endpoint
	// us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// A description of the set that helps with identification. You cannot change the
	// description of a set after you create it.
	Description *string

	// An array of key:value pairs to associate with the resource.
	Tags []types.Tag
}

type CreateRegexPatternSetOutput struct {

	// High-level information about a RegexPatternSet, returned by operations like
	// create and list. This provides information like the ID, that you can use to
	// retrieve and manage a RegexPatternSet, and the ARN, that you provide to the
	// RegexPatternSetReferenceStatement to use the pattern set in a Rule.
	Summary *types.RegexPatternSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRegexPatternSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRegexPatternSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRegexPatternSet{}, middleware.After)
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
	if err = addOpCreateRegexPatternSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegexPatternSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRegexPatternSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "CreateRegexPatternSet",
	}
}
