// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the configured contact methods. Specify a protocol in
// your request to return information about a specific contact method. A contact
// method is used to send you notifications about your Amazon Lightsail resources.
// You can add one email address and one mobile phone number contact method in each
// AWS Region. However, SMS text messaging is not supported in some AWS Regions,
// and SMS text messages cannot be sent to some countries/regions. For more
// information, see Notifications in Amazon Lightsail
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-notifications).
func (c *Client) GetContactMethods(ctx context.Context, params *GetContactMethodsInput, optFns ...func(*Options)) (*GetContactMethodsOutput, error) {
	if params == nil {
		params = &GetContactMethodsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContactMethods", params, optFns, addOperationGetContactMethodsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContactMethodsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContactMethodsInput struct {

	// The protocols used to send notifications, such as Email, or SMS (text
	// messaging). Specify a protocol in your request to return information about a
	// specific contact method protocol.
	Protocols []types.ContactProtocol
}

type GetContactMethodsOutput struct {

	// An array of objects that describe the contact methods.
	ContactMethods []types.ContactMethod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetContactMethodsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContactMethods{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContactMethods{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContactMethods(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContactMethods(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetContactMethods",
	}
}
