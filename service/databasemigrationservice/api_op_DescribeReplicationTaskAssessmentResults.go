// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the task assessment results from Amazon S3. This action always returns
// the latest results.
func (c *Client) DescribeReplicationTaskAssessmentResults(ctx context.Context, params *DescribeReplicationTaskAssessmentResultsInput, optFns ...func(*Options)) (*DescribeReplicationTaskAssessmentResultsOutput, error) {
	stack := middleware.NewStack("DescribeReplicationTaskAssessmentResults", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeReplicationTaskAssessmentResultsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentResults(options.Region), middleware.Before)

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
			OperationName: "DescribeReplicationTaskAssessmentResults",
			Err:           err,
		}
	}
	out := result.(*DescribeReplicationTaskAssessmentResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeReplicationTaskAssessmentResultsInput struct {
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// The Amazon Resource Name (ARN) string that uniquely identifies the task. When
	// this input parameter is specified, the API returns only one result and ignore
	// the values of the MaxRecords and Marker parameters.
	ReplicationTaskArn *string
}

//
type DescribeReplicationTaskAssessmentResultsOutput struct {
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string
	// - The Amazon S3 bucket where the task assessment report is located.
	BucketName *string
	// The task assessment report.
	ReplicationTaskAssessmentResults []*types.ReplicationTaskAssessmentResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeReplicationTaskAssessmentResultsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationTaskAssessmentResults{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationTaskAssessmentResults{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReplicationTaskAssessmentResults(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationTaskAssessmentResults",
	}
}