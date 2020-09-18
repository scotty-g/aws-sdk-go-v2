// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Obtains information about the directory snapshots that belong to this account.
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// DescribeSnapshots.NextToken member contains a token that you pass in the next
// call to DescribeSnapshots () to retrieve the next set of items. You can also
// specify a maximum number of return results with the Limit parameter.
func (c *Client) DescribeSnapshots(ctx context.Context, params *DescribeSnapshotsInput, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	stack := middleware.NewStack("DescribeSnapshots", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeSnapshotsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshots(options.Region), middleware.Before)

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
			OperationName: "DescribeSnapshots",
			Err:           err,
		}
	}
	out := result.(*DescribeSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DescribeSnapshots () operation.
type DescribeSnapshotsInput struct {
	// The identifier of the directory for which to retrieve snapshot information.
	DirectoryId *string
	// The maximum number of objects to return.
	Limit *int32
	// A list of identifiers of the snapshots to obtain the information for. If this
	// member is null or empty, all snapshots are returned using the Limit and
	// NextToken members.
	SnapshotIds []*string
	// The DescribeSnapshotsResult.NextToken value from a previous call to
	// DescribeSnapshots (). Pass null if this is the first call.
	NextToken *string
}

// Contains the results of the DescribeSnapshots () operation.
type DescribeSnapshotsOutput struct {
	// If not null, more results are available. Pass this value in the NextToken member
	// of a subsequent call to DescribeSnapshots ().
	NextToken *string
	// The list of Snapshot () objects that were retrieved. It is possible that this
	// list contains less than the number of items specified in the Limit member of the
	// request. This occurs if there are less than the requested number of items left
	// to retrieve, or if the limitations of the operation have been exceeded.
	Snapshots []*types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeSnapshotsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSnapshots{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSnapshots{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeSnapshots",
	}
}