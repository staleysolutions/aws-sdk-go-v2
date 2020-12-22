// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified VPC with the specified local gateway route table.
func (c *Client) CreateLocalGatewayRouteTableVpcAssociation(ctx context.Context, params *CreateLocalGatewayRouteTableVpcAssociationInput, optFns ...func(*Options)) (*CreateLocalGatewayRouteTableVpcAssociationOutput, error) {
	if params == nil {
		params = &CreateLocalGatewayRouteTableVpcAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocalGatewayRouteTableVpcAssociation", params, optFns, addOperationCreateLocalGatewayRouteTableVpcAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocalGatewayRouteTableVpcAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocalGatewayRouteTableVpcAssociationInput struct {

	// The ID of the local gateway route table.
	//
	// This member is required.
	LocalGatewayRouteTableId *string

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The tags to assign to the local gateway route table VPC association.
	TagSpecifications []types.TagSpecification
}

type CreateLocalGatewayRouteTableVpcAssociationOutput struct {

	// Information about the association.
	LocalGatewayRouteTableVpcAssociation *types.LocalGatewayRouteTableVpcAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLocalGatewayRouteTableVpcAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateLocalGatewayRouteTableVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateLocalGatewayRouteTableVpcAssociation{}, middleware.After)
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
	if err = addOpCreateLocalGatewayRouteTableVpcAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocalGatewayRouteTableVpcAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocalGatewayRouteTableVpcAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateLocalGatewayRouteTableVpcAssociation",
	}
}
