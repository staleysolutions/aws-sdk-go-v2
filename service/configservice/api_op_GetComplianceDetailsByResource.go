// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the evaluation results for the specified AWS resource. The results
// indicate which AWS Config rules were used to evaluate the resource, when each
// rule was last used, and whether the resource complies with each rule.
func (c *Client) GetComplianceDetailsByResource(ctx context.Context, params *GetComplianceDetailsByResourceInput, optFns ...func(*Options)) (*GetComplianceDetailsByResourceOutput, error) {
	if params == nil {
		params = &GetComplianceDetailsByResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceDetailsByResource", params, optFns, addOperationGetComplianceDetailsByResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceDetailsByResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type GetComplianceDetailsByResourceInput struct {

	// The ID of the AWS resource for which you want compliance information.
	//
	// This member is required.
	ResourceId *string

	// The type of the AWS resource for which you want compliance information.
	//
	// This member is required.
	ResourceType *string

	// Filters the results by compliance. The allowed values are COMPLIANT,
	// NON_COMPLIANT, and NOT_APPLICABLE.
	ComplianceTypes []types.ComplianceType

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

//
type GetComplianceDetailsByResourceOutput struct {

	// Indicates whether the specified AWS resource complies each AWS Config rule.
	EvaluationResults []types.EvaluationResult

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetComplianceDetailsByResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceDetailsByResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceDetailsByResource{}, middleware.After)
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
	if err = addOpGetComplianceDetailsByResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceDetailsByResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetComplianceDetailsByResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetComplianceDetailsByResource",
	}
}
