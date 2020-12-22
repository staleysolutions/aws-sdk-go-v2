// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a route that is associated with a virtual router. You can route several
// different protocols and define a retry policy for a route. Traffic can be routed
// to one or more virtual nodes. For more information about routes, see Routes
// (https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html).
func (c *Client) CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) {
	if params == nil {
		params = &CreateRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRoute", params, optFns, addOperationCreateRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateRouteInput struct {

	// The name of the service mesh to create the route in.
	//
	// This member is required.
	MeshName *string

	// The name to use for the route.
	//
	// This member is required.
	RouteName *string

	// The route specification to apply.
	//
	// This member is required.
	Spec *types.RouteSpec

	// The name of the virtual router in which to create the route. If the virtual
	// router is in a shared mesh, then you must be the owner of the virtual router
	// resource.
	//
	// This member is required.
	VirtualRouterName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then the account that you specify must share the mesh with your account
	// before you can create the resource in the service mesh. For more information
	// about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// Optional metadata that you can apply to the route to assist with categorization
	// and organization. Each tag consists of a key and an optional value, both of
	// which you define. Tag keys can have a maximum character length of 128
	// characters, and tag values can have a maximum length of 256 characters.
	Tags []types.TagRef
}

//
type CreateRouteOutput struct {

	// The full description of your mesh following the create call.
	//
	// This member is required.
	Route *types.RouteData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoute{}, middleware.After)
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
	if err = addOpCreateRouteValidationMiddleware(stack); err != nil {
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
