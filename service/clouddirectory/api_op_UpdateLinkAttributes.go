// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a given typed link’s attributes. Attributes to be updated must not
// contribute to the typed link’s identity, as defined by its
// IdentityAttributeOrder.
func (c *Client) UpdateLinkAttributes(ctx context.Context, params *UpdateLinkAttributesInput, optFns ...func(*Options)) (*UpdateLinkAttributesOutput, error) {
	if params == nil {
		params = &UpdateLinkAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLinkAttributes", params, optFns, addOperationUpdateLinkAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLinkAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLinkAttributesInput struct {

	// The attributes update structure.
	//
	// This member is required.
	AttributeUpdates []types.LinkAttributeUpdate

	// The Amazon Resource Name (ARN) that is associated with the Directory where the
	// updated typed link resides. For more information, see arns or Typed Links
	// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	//
	// This member is required.
	DirectoryArn *string

	// Allows a typed link specifier to be accepted as input.
	//
	// This member is required.
	TypedLinkSpecifier *types.TypedLinkSpecifier
}

type UpdateLinkAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateLinkAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLinkAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLinkAttributes{}, middleware.After)
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
	if err = addOpUpdateLinkAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLinkAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLinkAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpdateLinkAttributes",
	}
}
