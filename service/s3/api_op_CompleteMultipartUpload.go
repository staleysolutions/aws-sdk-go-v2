// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Completes a multipart upload by assembling previously uploaded parts. You first
// initiate the multipart upload and then upload all parts using the UploadPart
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html) operation.
// After successfully uploading all relevant parts of an upload, you call this
// operation to complete the upload. Upon receiving this request, Amazon S3
// concatenates all the parts in ascending order by part number to create a new
// object. In the Complete Multipart Upload request, you must provide the parts
// list. You must ensure that the parts list is complete. This operation
// concatenates the parts that you provide in the list. For each part in the list,
// you must provide the part number and the ETag value, returned after that part
// was uploaded. Processing of a Complete Multipart Upload request could take
// several minutes to complete. After Amazon S3 begins processing the request, it
// sends an HTTP response header that specifies a 200 OK response. While processing
// is in progress, Amazon S3 periodically sends white space characters to keep the
// connection from timing out. Because a request could fail after the initial 200
// OK response has been sent, it is important that you check the response body to
// determine whether the request succeeded. Note that if CompleteMultipartUpload
// fails, applications should be prepared to retry the failed requests. For more
// information, see Amazon S3 Error Best Practices
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ErrorBestPractices.html). For
// more information about multipart uploads, see Uploading Objects Using Multipart
// Upload (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html).
// For information about permissions required to use the multipart upload API, see
// Multipart Upload API and Permissions
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html).
// CompleteMultipartUpload has the following special errors:
//
// * Error code:
// EntityTooSmall
//
// * Description: Your proposed upload is smaller than the minimum
// allowed object size. Each part must be at least 5 MB in size, except the last
// part.
//
// * 400 Bad Request
//
// * Error code: InvalidPart
//
// * Description: One or more
// of the specified parts could not be found. The part might not have been
// uploaded, or the specified entity tag might not have matched the part's entity
// tag.
//
// * 400 Bad Request
//
// * Error code: InvalidPartOrder
//
// * Description: The list
// of parts was not in ascending order. The parts list must be specified in order
// by part number.
//
// * 400 Bad Request
//
// * Error code: NoSuchUpload
//
// * Description:
// The specified multipart upload does not exist. The upload ID might be invalid,
// or the multipart upload might have been aborted or completed.
//
// * 404 Not
// Found
//
// The following operations are related to CompleteMultipartUpload:
//
// *
// CreateMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
//
// *
// UploadPart
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
//
// *
// AbortMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)
//
// *
// ListParts
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)
//
// *
// ListMultipartUploads
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html)
func (c *Client) CompleteMultipartUpload(ctx context.Context, params *CompleteMultipartUploadInput, optFns ...func(*Options)) (*CompleteMultipartUploadOutput, error) {
	if params == nil {
		params = &CompleteMultipartUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteMultipartUpload", params, optFns, addOperationCompleteMultipartUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteMultipartUploadInput struct {

	// Name of the bucket to which the multipart upload was initiated.
	//
	// This member is required.
	Bucket *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// ID for the initiated multipart upload.
	//
	// This member is required.
	UploadId *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// The container for the multipart upload request information.
	MultipartUpload *types.CompletedMultipartUpload

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
}

type CompleteMultipartUploadOutput struct {

	// The name of the bucket that contains the newly created object. When using this
	// API with an access point, you must direct requests to the access point hostname.
	// The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation with an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide. When using this API with
	// Amazon S3 on Outposts, you must direct requests to the S3 on Outposts hostname.
	// The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When using
	// this operation using S3 on Outposts through the AWS SDKs, you provide the
	// Outposts bucket ARN in place of the bucket name. For more information about S3
	// on Outposts ARNs, see Using S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
	// Amazon Simple Storage Service Developer Guide.
	Bucket *string

	// Entity tag that identifies the newly created object's data. Objects with
	// different object data will have different entity tags. The entity tag is an
	// opaque string. The entity tag may or may not be an MD5 digest of the object
	// data. If the entity tag is not an MD5 digest of the object data, it will contain
	// one or more nonhexadecimal characters and/or will consist of less than 32 or
	// more than 32 hexadecimal digits.
	ETag *string

	// If the object expiration is configured, this will contain the expiration date
	// (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration *string

	// The object key of the newly created object.
	Key *string

	// The URI that identifies the newly created object.
	Location *string

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// If you specified server-side encryption either with an Amazon S3-managed
	// encryption key or an AWS KMS customer master key (CMK) in your initiate
	// multipart upload request, the response includes this header. It confirms the
	// encryption algorithm that Amazon S3 used to encrypt the object.
	ServerSideEncryption types.ServerSideEncryption

	// Version ID of the newly created object, in case the bucket has versioning turned
	// on.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCompleteMultipartUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCompleteMultipartUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCompleteMultipartUpload{}, middleware.After)
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
	if err = addOpCompleteMultipartUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteMultipartUpload(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addCompleteMultipartUploadUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = s3cust.HandleResponseErrorWith200Status(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCompleteMultipartUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CompleteMultipartUpload",
	}
}

// getCompleteMultipartUploadBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getCompleteMultipartUploadBucketMember(input interface{}) (*string, bool) {
	in := input.(*CompleteMultipartUploadInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addCompleteMultipartUploadUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getCompleteMultipartUploadBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
