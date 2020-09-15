// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an analytics configuration for the bucket (specified by the analytics
// configuration ID). To use this operation, you must have permissions to perform
// the s3:PutAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
// <p>For information about the Amazon S3 analytics feature, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html">Amazon
// S3 Analytics – Storage Class Analysis</a>. </p> <p>The following operations are
// related to <code>DeleteBucketAnalyticsConfiguration</code>:</p> <ul> <li> <p>
// </p> </li> <li> <p> </p> </li> <li> <p> </p> </li> </ul>
func (c *Client) DeleteBucketAnalyticsConfiguration(ctx context.Context, params *DeleteBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*DeleteBucketAnalyticsConfigurationOutput, error) {
	stack := middleware.NewStack("DeleteBucketAnalyticsConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteBucketAnalyticsConfigurationMiddlewares(stack)
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
	addOpDeleteBucketAnalyticsConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketAnalyticsConfiguration(options.Region), middleware.Before)
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
			OperationName: "DeleteBucketAnalyticsConfiguration",
			Err:           err,
		}
	}
	out := result.(*DeleteBucketAnalyticsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketAnalyticsConfigurationInput struct {
	// The name of the bucket from which an analytics configuration is deleted.
	Bucket *string
	// The ID that identifies the analytics configuration.
	Id *string
}

type DeleteBucketAnalyticsConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteBucketAnalyticsConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketAnalyticsConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketAnalyticsConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucketAnalyticsConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketAnalyticsConfiguration",
	}
}
