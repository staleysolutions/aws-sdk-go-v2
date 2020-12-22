// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or removes the specified license configurations for the specified AWS
// resource. You can update the license specifications of AMIs, instances, and
// hosts. You cannot update the license specifications for launch templates and AWS
// CloudFormation templates, as they send license configurations to the operation
// that creates the resource.
func (c *Client) UpdateLicenseSpecificationsForResource(ctx context.Context, params *UpdateLicenseSpecificationsForResourceInput, optFns ...func(*Options)) (*UpdateLicenseSpecificationsForResourceOutput, error) {
	if params == nil {
		params = &UpdateLicenseSpecificationsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLicenseSpecificationsForResource", params, optFns, addOperationUpdateLicenseSpecificationsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLicenseSpecificationsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLicenseSpecificationsForResourceInput struct {

	// Amazon Resource Name (ARN) of the AWS resource.
	//
	// This member is required.
	ResourceArn *string

	// ARNs of the license configurations to add.
	AddLicenseSpecifications []types.LicenseSpecification

	// ARNs of the license configurations to remove.
	RemoveLicenseSpecifications []types.LicenseSpecification
}

type UpdateLicenseSpecificationsForResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateLicenseSpecificationsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLicenseSpecificationsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLicenseSpecificationsForResource{}, middleware.After)
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
	if err = addOpUpdateLicenseSpecificationsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLicenseSpecificationsForResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLicenseSpecificationsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "UpdateLicenseSpecificationsForResource",
	}
}
