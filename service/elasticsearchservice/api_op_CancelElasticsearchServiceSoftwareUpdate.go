// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a scheduled service software update for an Amazon ES domain. You can
// only perform this operation before the AutomatedUpdateDate and when the
// UpdateStatus is in the PENDING_UPDATE state.
func (c *Client) CancelElasticsearchServiceSoftwareUpdate(ctx context.Context, params *CancelElasticsearchServiceSoftwareUpdateInput, optFns ...func(*Options)) (*CancelElasticsearchServiceSoftwareUpdateOutput, error) {
	if params == nil {
		params = &CancelElasticsearchServiceSoftwareUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelElasticsearchServiceSoftwareUpdate", params, optFns, addOperationCancelElasticsearchServiceSoftwareUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelElasticsearchServiceSoftwareUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the CancelElasticsearchServiceSoftwareUpdate
// operation. Specifies the name of the Elasticsearch domain that you wish to
// cancel a service software update on.
type CancelElasticsearchServiceSoftwareUpdateInput struct {

	// The name of the domain that you want to stop the latest service software update
	// on.
	//
	// This member is required.
	DomainName *string
}

// The result of a CancelElasticsearchServiceSoftwareUpdate operation. Contains the
// status of the update.
type CancelElasticsearchServiceSoftwareUpdateOutput struct {

	// The current status of the Elasticsearch service software update.
	ServiceSoftwareOptions *types.ServiceSoftwareOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelElasticsearchServiceSoftwareUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelElasticsearchServiceSoftwareUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelElasticsearchServiceSoftwareUpdate{}, middleware.After)
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
	if err = addOpCancelElasticsearchServiceSoftwareUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelElasticsearchServiceSoftwareUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelElasticsearchServiceSoftwareUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CancelElasticsearchServiceSoftwareUpdate",
	}
}
