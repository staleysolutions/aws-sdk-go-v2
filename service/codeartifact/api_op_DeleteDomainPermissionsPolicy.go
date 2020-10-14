// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the resource policy set on a domain.
func (c *Client) DeleteDomainPermissionsPolicy(ctx context.Context, params *DeleteDomainPermissionsPolicyInput, optFns ...func(*Options)) (*DeleteDomainPermissionsPolicyOutput, error) {
	if params == nil {
		params = &DeleteDomainPermissionsPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDomainPermissionsPolicy", params, optFns, addOperationDeleteDomainPermissionsPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDomainPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDomainPermissionsPolicyInput struct {

	// The name of the domain associated with the resource policy to be deleted.
	//
	// This member is required.
	Domain *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The current revision of the resource policy to be deleted. This revision is used
	// for optimistic locking, which prevents others from overwriting your changes to
	// the domain's resource policy.
	PolicyRevision *string
}

type DeleteDomainPermissionsPolicyOutput struct {

	// Information about the deleted resource policy after processing the request.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDomainPermissionsPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDomainPermissionsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDomainPermissionsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteDomainPermissionsPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDomainPermissionsPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDomainPermissionsPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DeleteDomainPermissionsPolicy",
	}
}