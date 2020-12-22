// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a PackageVersionDescription
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html)
// object that contains information about the requested package version.
func (c *Client) DescribePackageVersion(ctx context.Context, params *DescribePackageVersionInput, optFns ...func(*Options)) (*DescribePackageVersionOutput, error) {
	if params == nil {
		params = &DescribePackageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackageVersion", params, optFns, addOperationDescribePackageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackageVersionInput struct {

	// The name of the domain that contains the repository that contains the package
	// version.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of the requested package version. The valid
	// values are:
	//
	// * npm
	//
	// * pypi
	//
	// * maven
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the requested package version.
	//
	// This member is required.
	Package *string

	// A string that contains the package version (for example, 3.5.2).
	//
	// This member is required.
	PackageVersion *string

	// The name of the repository that contains the package version.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string
}

type DescribePackageVersionOutput struct {

	// A PackageVersionDescription
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html)
	// object that contains information about the requested package version.
	//
	// This member is required.
	PackageVersion *types.PackageVersionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePackageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackageVersion{}, middleware.After)
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
	if err = addOpDescribePackageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackageVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePackageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DescribePackageVersion",
	}
}
