// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Uploads a part in a multipart upload. In this operation, you provide part data
// in your request. However, you have an option to specify your existing Amazon S3
// object as a data source for the part you are uploading. To upload a part from an
// existing object, you use the UploadPartCopy () operation.  <p>You must initiate
// a multipart upload (see <a>CreateMultipartUpload</a>) before you can upload any
// part. In response to your initiate request, Amazon S3 returns an upload ID, a
// unique identifier, that you must include in your upload part request.</p>
// <p>Part numbers can be any number from 1 to 10,000, inclusive. A part number
// uniquely identifies a part and also defines its position within the object being
// created. If you upload a new part using the same part number that was used with
// a previous part, the previously uploaded part is overwritten. Each part must be
// at least 5 MB in size, except the last part. There is no size limit on the last
// part of your multipart upload.</p> <p>To ensure that data is not corrupted when
// traversing the network, specify the <code>Content-MD5</code> header in the
// upload part request. Amazon S3 checks the part data against the provided MD5
// value. If they do not match, Amazon S3 returns an error. </p> <p> <b>Note:</b>
// After you initiate multipart upload and upload one or more parts, you must
// either complete or abort multipart upload in order to stop getting charged for
// storage of the uploaded parts. Only after you either complete or abort multipart
// upload, Amazon S3 frees up the parts storage and stops charging you for the
// parts storage.</p> <p>For more information on multipart uploads, go to <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html">Multipart
// Upload Overview</a> in the <i>Amazon Simple Storage Service Developer Guide
// </i>.</p> <p>For information on the permissions required to use the multipart
// upload API, go to <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html">Multipart
// Upload API and Permissions</a> in the <i>Amazon Simple Storage Service Developer
// Guide</i>.</p> <p>You can optionally request server-side encryption where Amazon
// S3 encrypts your data as it writes it to disks in its data centers and decrypts
// it for you when you access it. You have the option of providing your own
// encryption key, or you can use the AWS managed encryption keys. If you choose to
// provide your own encryption key, the request headers you provide in the request
// must match the headers you used in the request to initiate the upload by using
// <a>CreateMultipartUpload</a>. For more information, go to <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html">Using
// Server-Side Encryption</a> in the <i>Amazon Simple Storage Service Developer
// Guide</i>.</p> <p>Server-side encryption is supported by the S3 Multipart Upload
// actions. Unless you are using a customer-provided encryption key, you don't need
// to specify the encryption parameters in each UploadPart request. Instead, you
// only need to specify the server-side encryption parameters in the initial
// Initiate Multipart request. For more information, see
// <a>CreateMultipartUpload</a>.</p> <p>If you requested server-side encryption
// using a customer-provided encryption key in your initiate multipart upload
// request, you must provide identical encryption information in each part upload
// using the following headers.</p> <ul> <li>
// <p>x-amz-server-side-encryption-customer-algorithm</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key-MD5</p> </li> </ul> <p
// class="title"> <b>Special Errors</b> </p> <ul> <li> <p class="title"> <b></b>
// </p> <ul> <li> <p> <i>Code: NoSuchUpload</i> </p> </li> <li> <p> <i>Cause: The
// specified multipart upload does not exist. The upload ID might be invalid, or
// the multipart upload might have been aborted or completed.</i> </p> </li> <li>
// <p> <i> HTTP Status Code: 404 Not Found </i> </p> </li> <li> <p> <i>SOAP Fault
// Code Prefix: Client</i> </p> </li> </ul> </li> </ul> <p class="title">
// <b>Related Resources</b> </p> <ul> <li> <p> <a>CreateMultipartUpload</a> </p>
// </li> <li> <p> <a>CompleteMultipartUpload</a> </p> </li> <li> <p>
// <a>AbortMultipartUpload</a> </p> </li> <li> <p> <a>ListParts</a> </p> </li> <li>
// <p> <a>ListMultipartUploads</a> </p> </li> </ul>
func (c *Client) UploadPart(ctx context.Context, params *UploadPartInput, optFns ...func(*Options)) (*UploadPartOutput, error) {
	stack := middleware.NewStack("UploadPart", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUploadPartMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUploadPartValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUploadPart(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "UploadPart",
			Err:           err,
		}
	}
	out := result.(*UploadPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadPartInput struct {
	// Size of the body in bytes. This parameter is useful when the size of the body
	// cannot be determined automatically.
	ContentLength *int64
	// Name of the bucket to which the multipart upload was initiated.
	Bucket *string
	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header. This must be the same
	// encryption key specified in the initiate multipart upload request.
	SSECustomerKey *string
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
	// The base64-encoded 128-bit MD5 digest of the part data. This parameter is
	// auto-populated when using the command from the CLI. This parameter is required
	// if object lock parameters are specified.
	ContentMD5 *string
	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string
	// Object data.
	Body io.Reader
	// Part number of part being uploaded. This is a positive integer between 1 and
	// 10,000.
	PartNumber *int32
	// Upload ID identifying the multipart upload whose part is being uploaded.
	UploadId *string
	// Object key for which the multipart upload was initiated.
	Key *string
}

type UploadPartOutput struct {
	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) was used for the object.
	SSEKMSKeyId *string
	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged
	// Entity tag for the uploaded object.
	ETag *string
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUploadPartMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUploadPart{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUploadPart{}, middleware.After)
}

func newServiceMetadataMiddleware_opUploadPart(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UploadPart",
	}
}