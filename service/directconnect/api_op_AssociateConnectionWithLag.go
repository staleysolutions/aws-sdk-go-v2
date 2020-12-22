// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Associates an existing connection with a link aggregation group (LAG). The
// connection is interrupted and re-established as a member of the LAG
// (connectivity to AWS is interrupted). The connection must be hosted on the same
// AWS Direct Connect endpoint as the LAG, and its bandwidth must match the
// bandwidth for the LAG. You can re-associate a connection that's currently
// associated with a different LAG; however, if removing the connection would cause
// the original LAG to fall below its setting for minimum number of operational
// connections, the request fails. Any virtual interfaces that are directly
// associated with the connection are automatically re-associated with the LAG. If
// the connection was originally associated with a different LAG, the virtual
// interfaces remain associated with the original LAG. For interconnects, any
// hosted connections are automatically re-associated with the LAG. If the
// interconnect was originally associated with a different LAG, the hosted
// connections remain associated with the original LAG.
func (c *Client) AssociateConnectionWithLag(ctx context.Context, params *AssociateConnectionWithLagInput, optFns ...func(*Options)) (*AssociateConnectionWithLagOutput, error) {
	if params == nil {
		params = &AssociateConnectionWithLagInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateConnectionWithLag", params, optFns, addOperationAssociateConnectionWithLagMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateConnectionWithLagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateConnectionWithLagInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	// The ID of the LAG with which to associate the connection.
	//
	// This member is required.
	LagId *string
}

// Information about an AWS Direct Connect connection.
type AssociateConnectionWithLagOutput struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//
	// * ordering:
	// The initial state of a hosted connection provisioned on an interconnect. The
	// connection stays in the ordering state until the owner of the hosted connection
	// confirms or declines the connection order.
	//
	// * requested: The initial state of a
	// standard connection. The connection stays in the requested state until the
	// Letter of Authorization (LOA) is sent to the customer.
	//
	// * pending: The
	// connection has been approved and is being initialized.
	//
	// * available: The network
	// link is up and the connection is ready for use.
	//
	// * down: The network link is
	// down.
	//
	// * deleting: The connection is being deleted.
	//
	// * deleted: The connection
	// has been deleted.
	//
	// * rejected: A hosted connection in the ordering state enters
	// the rejected state if it is deleted by the customer.
	//
	// * unknown: The state of
	// the connection is not available.
	ConnectionState types.ConnectionState

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// The ID of the AWS account that owns the connection.
	OwnerAccount *string

	// The name of the AWS Direct Connect service provider associated with the
	// connection.
	PartnerName *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The AWS Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []types.Tag

	// The ID of the VLAN.
	Vlan int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateConnectionWithLagMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateConnectionWithLag{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateConnectionWithLag{}, middleware.After)
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
	if err = addOpAssociateConnectionWithLagValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateConnectionWithLag(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateConnectionWithLag(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "AssociateConnectionWithLag",
	}
}
