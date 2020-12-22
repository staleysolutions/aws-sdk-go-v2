// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Purchases a Reserved Instance for use with your account. With Reserved
// Instances, you pay a lower hourly rate compared to On-Demand instance pricing.
// Use DescribeReservedInstancesOfferings to get a list of Reserved Instance
// offerings that match your specifications. After you've purchased a Reserved
// Instance, you can check for your new Reserved Instance with
// DescribeReservedInstances. To queue a purchase for a future date and time,
// specify a purchase time. If you do not specify a purchase time, the default is
// the current time. For more information, see Reserved Instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts-on-demand-reserved-instances.html)
// and Reserved Instance Marketplace
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) PurchaseReservedInstancesOffering(ctx context.Context, params *PurchaseReservedInstancesOfferingInput, optFns ...func(*Options)) (*PurchaseReservedInstancesOfferingOutput, error) {
	if params == nil {
		params = &PurchaseReservedInstancesOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseReservedInstancesOffering", params, optFns, addOperationPurchaseReservedInstancesOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseReservedInstancesOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for PurchaseReservedInstancesOffering.
type PurchaseReservedInstancesOfferingInput struct {

	// The number of Reserved Instances to purchase.
	//
	// This member is required.
	InstanceCount int32

	// The ID of the Reserved Instance offering to purchase.
	//
	// This member is required.
	ReservedInstancesOfferingId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// Specified for Reserved Instance Marketplace offerings to limit the total order
	// and ensure that the Reserved Instances are not purchased at unexpected prices.
	LimitPrice *types.ReservedInstanceLimitPrice

	// The time at which to purchase the Reserved Instance, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	PurchaseTime *time.Time
}

// Contains the output of PurchaseReservedInstancesOffering.
type PurchaseReservedInstancesOfferingOutput struct {

	// The IDs of the purchased Reserved Instances.
	ReservedInstancesId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPurchaseReservedInstancesOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpPurchaseReservedInstancesOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpPurchaseReservedInstancesOffering{}, middleware.After)
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
	if err = addOpPurchaseReservedInstancesOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedInstancesOffering(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPurchaseReservedInstancesOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "PurchaseReservedInstancesOffering",
	}
}
