// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about DB snapshots. This API action supports pagination.
func (c *Client) DescribeDBSnapshots(ctx context.Context, params *DescribeDBSnapshotsInput, optFns ...func(*Options)) (*DescribeDBSnapshotsOutput, error) {
	stack := middleware.NewStack("DescribeDBSnapshots", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeDBSnapshotsMiddlewares(stack)
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
	addOpDescribeDBSnapshotsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBSnapshots(options.Region), middleware.Before)

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
			OperationName: "DescribeDBSnapshots",
			Err:           err,
		}
	}
	out := result.(*DescribeDBSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBSnapshotsInput struct {
	// The type of snapshots to be returned. You can specify one of the following
	// values:
	//
	//     * automated - Return all DB snapshots that have been automatically
	// taken by Amazon RDS for my AWS account.
	//
	//     * manual - Return all DB snapshots
	// that have been taken by my AWS account.
	//
	//     * shared - Return all manual DB
	// snapshots that have been shared to my AWS account.
	//
	//     * public - Return all DB
	// snapshots that have been marked as public.
	//
	//     * awsbackup - Return the DB
	// snapshots managed by the AWS Backup service. For information about AWS Backup,
	// see the  AWS Backup Developer Guide.
	// (https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html) The
	// awsbackup type does not apply to Aurora.
	//
	// If you don't specify a SnapshotType
	// value, then both automated and manual snapshots are returned. Shared and public
	// DB snapshots are not included in the returned results by default. You can
	// include shared snapshots with these results by enabling the IncludeShared
	// parameter. You can include public snapshots with these results by enabling the
	// IncludePublic parameter. The IncludeShared and IncludePublic parameters don't
	// apply for SnapshotType values of manual or automated. The IncludePublic
	// parameter doesn't apply when SnapshotType is set to shared. The IncludeShared
	// parameter doesn't apply when SnapshotType is set to public.
	SnapshotType *string
	// An optional pagination token provided by a previous DescribeDBSnapshots request.
	// If this parameter is specified, the response includes only records beyond the
	// marker, up to the value specified by MaxRecords.
	Marker *string
	// A value that indicates whether to include manual DB cluster snapshots that are
	// public and can be copied or restored by any AWS account. By default, the public
	// snapshots are not included. You can share a manual DB snapshot as public by
	// using the ModifyDBSnapshotAttribute () API.
	IncludePublic *bool
	// A value that indicates whether to include shared manual DB cluster snapshots
	// from other AWS accounts that this AWS account has been given permission to copy
	// or restore. By default, these snapshots are not included. You can give an AWS
	// account permission to restore a manual DB snapshot from another AWS account by
	// using the ModifyDBSnapshotAttribute API action.
	IncludeShared *bool
	// A filter that specifies one or more DB snapshots to describe. Supported
	// filters:
	//
	//     * db-instance-id - Accepts DB instance identifiers and DB instance
	// Amazon Resource Names (ARNs).
	//
	//     * db-snapshot-id - Accepts DB snapshot
	// identifiers.
	//
	//     * dbi-resource-id - Accepts identifiers of source DB
	// instances.
	//
	//     * snapshot-type - Accepts types of DB snapshots.
	//
	//     * engine -
	// Accepts names of database engines.
	Filters []*types.Filter
	// A specific DB snapshot identifier to describe. This parameter can't be used in
	// conjunction with DBInstanceIdentifier. This value is stored as a lowercase
	// string. Constraints:
	//
	//     * If supplied, must match the identifier of an
	// existing DBSnapshot.
	//
	//     * If this identifier is for an automated snapshot, the
	// SnapshotType parameter must also be specified.
	DBSnapshotIdentifier *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// A specific DB resource ID to describe.
	DbiResourceId *string
	// The ID of the DB instance to retrieve the list of DB snapshots for. This
	// parameter can't be used in conjunction with DBSnapshotIdentifier. This parameter
	// isn't case-sensitive. Constraints:
	//
	//     * If supplied, must match the identifier
	// of an existing DBInstance.
	DBInstanceIdentifier *string
}

// Contains the result of a successful invocation of the DescribeDBSnapshots
// action.
type DescribeDBSnapshotsOutput struct {
	// A list of DBSnapshot instances.
	DBSnapshots []*types.DBSnapshot
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeDBSnapshotsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBSnapshots{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBSnapshots{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDBSnapshots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBSnapshots",
	}
}