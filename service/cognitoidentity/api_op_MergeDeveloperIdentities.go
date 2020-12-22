// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Merges two users having different IdentityIds, existing in the same identity
// pool, and identified by the same developer provider. You can use this action to
// request that discrete users be merged and identified as a single user in the
// Cognito environment. Cognito associates the given source user
// (SourceUserIdentifier) with the IdentityId of the DestinationUserIdentifier.
// Only developer-authenticated users can be merged. If the users to be merged are
// associated with the same public provider, but as two different users, an
// exception will be thrown. The number of linked logins is limited to 20. So, the
// number of linked logins for the source user, SourceUserIdentifier, and the
// destination user, DestinationUserIdentifier, together should not be larger than
// 20. Otherwise, an exception will be thrown. You must use AWS Developer
// credentials to call this API.
func (c *Client) MergeDeveloperIdentities(ctx context.Context, params *MergeDeveloperIdentitiesInput, optFns ...func(*Options)) (*MergeDeveloperIdentitiesOutput, error) {
	if params == nil {
		params = &MergeDeveloperIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MergeDeveloperIdentities", params, optFns, addOperationMergeDeveloperIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MergeDeveloperIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the MergeDeveloperIdentities action.
type MergeDeveloperIdentitiesInput struct {

	// User identifier for the destination user. The value should be a
	// DeveloperUserIdentifier.
	//
	// This member is required.
	DestinationUserIdentifier *string

	// The "domain" by which Cognito will refer to your users. This is a (pseudo)
	// domain name that you provide while creating an identity pool. This name acts as
	// a placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName, you can use letters
	// as well as period (.), underscore (_), and dash (-).
	//
	// This member is required.
	DeveloperProviderName *string

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// User identifier for the source user. The value should be a
	// DeveloperUserIdentifier.
	//
	// This member is required.
	SourceUserIdentifier *string
}

// Returned in response to a successful MergeDeveloperIdentities action.
type MergeDeveloperIdentitiesOutput struct {

	// A unique identifier in the format REGION:GUID.
	IdentityId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationMergeDeveloperIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMergeDeveloperIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergeDeveloperIdentities{}, middleware.After)
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
	if err = addOpMergeDeveloperIdentitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMergeDeveloperIdentities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMergeDeveloperIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "MergeDeveloperIdentities",
	}
}
