// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified placement groups or all of your placement groups. For
// more information, see Placement groups
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribePlacementGroups(ctx context.Context, params *DescribePlacementGroupsInput, optFns ...func(*Options)) (*DescribePlacementGroupsOutput, error) {
	stack := middleware.NewStack("DescribePlacementGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribePlacementGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePlacementGroups(options.Region), middleware.Before)

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
			OperationName: "DescribePlacementGroups",
			Err:           err,
		}
	}
	out := result.(*DescribePlacementGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePlacementGroupsInput struct {
	// The IDs of the placement groups.
	GroupIds []*string
	// The names of the placement groups. Default: Describes all your placement groups,
	// or only those otherwise specified.
	GroupNames []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The filters.
	//
	//     * group-name - The name of the placement group.
	//
	//     * state -
	// The state of the placement group (pending | available | deleting | deleted).
	//
	//
	// * strategy - The strategy of the placement group (cluster | spread |
	// partition).
	//
	//     * tag: - The key/value combination of a tag assigned to the
	// resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	//     * tag-key - The key of a tag assigned to the resource. Use this
	// filter to find all resources that have a tag with a specific key, regardless of
	// the tag value.
	Filters []*types.Filter
}

type DescribePlacementGroupsOutput struct {
	// Information about the placement groups.
	PlacementGroups []*types.PlacementGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribePlacementGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribePlacementGroups{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribePlacementGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePlacementGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribePlacementGroups",
	}
}