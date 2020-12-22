// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the current status of an existing association or disassociation request.
// A ResourceNotFoundException is thrown when no recent association or
// disassociation request with the specified token is found, or when the server
// does not exist. A ValidationException is raised when parameters of the request
// are not valid.
func (c *Client) DescribeNodeAssociationStatus(ctx context.Context, params *DescribeNodeAssociationStatusInput, optFns ...func(*Options)) (*DescribeNodeAssociationStatusOutput, error) {
	if params == nil {
		params = &DescribeNodeAssociationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNodeAssociationStatus", params, optFns, addOperationDescribeNodeAssociationStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNodeAssociationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNodeAssociationStatusInput struct {

	// The token returned in either the AssociateNodeResponse or the
	// DisassociateNodeResponse.
	//
	// This member is required.
	NodeAssociationStatusToken *string

	// The name of the server from which to disassociate the node.
	//
	// This member is required.
	ServerName *string
}

type DescribeNodeAssociationStatusOutput struct {

	// Attributes specific to the node association. In Puppet, the attibute
	// PUPPET_NODE_CERT contains the signed certificate (the result of the CSR).
	EngineAttributes []types.EngineAttribute

	// The status of the association or disassociation request. Possible values:
	//
	// *
	// SUCCESS: The association or disassociation succeeded.
	//
	// * FAILED: The association
	// or disassociation failed.
	//
	// * IN_PROGRESS: The association or disassociation is
	// still in progress.
	NodeAssociationStatus types.NodeAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNodeAssociationStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNodeAssociationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNodeAssociationStatus{}, middleware.After)
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
	if err = addOpDescribeNodeAssociationStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNodeAssociationStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeNodeAssociationStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "DescribeNodeAssociationStatus",
	}
}
