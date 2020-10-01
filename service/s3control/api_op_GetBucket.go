// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets an Amazon S3 on Outposts bucket. For more information, see  Using Amazon S3
// on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html)
// in the Amazon Simple Storage Service Developer Guide. The following actions are
// related to GetBucket for Amazon S3 on Outposts:
//
//     * PutObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//
//     *
// CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_CreateBucket.html)
//
//
// * DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucket.html)
func (c *Client) GetBucket(ctx context.Context, params *GetBucketInput, optFns ...func(*Options)) (*GetBucketOutput, error) {
	stack := middleware.NewStack("GetBucket", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketMiddlewares(stack)
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
	addOpGetBucketValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucket(options.Region), middleware.Before)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)

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
			OperationName: "GetBucket",
			Err:           err,
		}
	}
	out := result.(*GetBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The ARN of the bucket. For Amazon S3 on Outposts specify the ARN of the bucket
	// accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For example, to
	// access the bucket reports through outpost my-outpost owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type GetBucketOutput struct {

	// The creation date of the Outposts bucket.
	CreationDate *time.Time

	//
	PublicAccessBlockEnabled *bool

	// The Outposts bucket requested.
	Bucket *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucket{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucket{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucket(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucket",
	}
}