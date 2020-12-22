// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a named query in the specified workgroup. Requires that you have access
// to the workgroup. For code samples using the AWS SDK for Java, see Examples and
// Code Samples (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in
// the Amazon Athena User Guide.
func (c *Client) CreateNamedQuery(ctx context.Context, params *CreateNamedQueryInput, optFns ...func(*Options)) (*CreateNamedQueryOutput, error) {
	if params == nil {
		params = &CreateNamedQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNamedQuery", params, optFns, addOperationCreateNamedQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNamedQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNamedQueryInput struct {

	// The database to which the query belongs.
	//
	// This member is required.
	Database *string

	// The query name.
	//
	// This member is required.
	Name *string

	// The contents of the query with all query statements.
	//
	// This member is required.
	QueryString *string

	// A unique case-sensitive string used to ensure the request to create the query is
	// idempotent (executes only once). If another CreateNamedQuery request is
	// received, the same response is returned and another query is not created. If a
	// parameter has changed, for example, the QueryString, an error is returned. This
	// token is listed as not required because AWS SDKs (for example the AWS SDK for
	// Java) auto-generate the token for users. If you are not using the AWS SDK or the
	// AWS CLI, you must provide this token or the action will fail.
	ClientRequestToken *string

	// The query description.
	Description *string

	// The name of the workgroup in which the named query is being created.
	WorkGroup *string
}

type CreateNamedQueryOutput struct {

	// The unique ID of the query.
	NamedQueryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNamedQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNamedQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNamedQuery{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateNamedQueryMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNamedQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNamedQuery(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateNamedQuery struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNamedQuery) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNamedQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNamedQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNamedQueryInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNamedQueryMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNamedQuery{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNamedQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "CreateNamedQuery",
	}
}
