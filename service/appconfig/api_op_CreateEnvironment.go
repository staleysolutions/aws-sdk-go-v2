// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For each application, you define one or more environments. An environment is a
// logical deployment group of AppConfig targets, such as applications in a Beta or
// Production environment. You can also define environments for application
// subcomponents such as the Web, Mobile and Back-end components for your
// application. You can configure Amazon CloudWatch alarms for each environment.
// The system monitors alarms during a configuration deployment. If an alarm is
// triggered, the system rolls back the configuration.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// A name for the environment.
	//
	// This member is required.
	Name *string

	// A description of the environment.
	Description *string

	// Amazon CloudWatch alarms to monitor during the deployment process.
	Monitors []types.Monitor

	// Metadata to assign to the environment. Tags help organize and categorize your
	// AppConfig resources. Each tag consists of a key and an optional value, both of
	// which you define.
	Tags map[string]string
}

type CreateEnvironmentOutput struct {

	// The application ID.
	ApplicationId *string

	// The description of the environment.
	Description *string

	// The environment ID.
	Id *string

	// Amazon CloudWatch alarms monitored during the deployment.
	Monitors []types.Monitor

	// The name of the environment.
	Name *string

	// The state of the environment. An environment can be in one of the following
	// states: READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK, or ROLLED_BACK
	State types.EnvironmentState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironment{}, middleware.After)
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
	if err = addOpCreateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateEnvironment",
	}
}
