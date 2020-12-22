// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// ServiceSetting is an account-level setting for an AWS service. This setting
// defines how a user interacts with or uses a service or a feature of a service.
// For example, if an AWS service charges money to the account based on feature or
// service usage, then the AWS service team might create a default setting of
// "false". This means the user can't use this feature unless they change the
// setting to "true" and intentionally opt in for a paid feature. Services map a
// SettingId object to a setting value. AWS services teams define the default value
// for a SettingId. You can't create a new SettingId, but you can overwrite the
// default value if you have the ssm:UpdateServiceSetting permission for the
// setting. Use the UpdateServiceSetting API action to change the default setting.
// Or use the ResetServiceSetting to change the value back to the original value
// defined by the AWS service team. Query the current service setting for the
// account.
func (c *Client) GetServiceSetting(ctx context.Context, params *GetServiceSettingInput, optFns ...func(*Options)) (*GetServiceSettingOutput, error) {
	if params == nil {
		params = &GetServiceSettingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceSetting", params, optFns, addOperationGetServiceSettingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body of the GetServiceSetting API action.
type GetServiceSettingInput struct {

	// The ID of the service setting to get. The setting ID can be
	// /ssm/parameter-store/default-parameter-tier,
	// /ssm/parameter-store/high-throughput-enabled, or
	// /ssm/managed-instance/activation-tier.
	//
	// This member is required.
	SettingId *string
}

// The query result body of the GetServiceSetting API action.
type GetServiceSettingOutput struct {

	// The query result of the current service setting.
	ServiceSetting *types.ServiceSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetServiceSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceSetting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceSetting{}, middleware.After)
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
	if err = addOpGetServiceSettingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceSetting(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetServiceSetting",
	}
}
