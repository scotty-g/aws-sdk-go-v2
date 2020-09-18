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

// Describes the target networks associated with the specified Client VPN endpoint.
func (c *Client) DescribeClientVpnTargetNetworks(ctx context.Context, params *DescribeClientVpnTargetNetworksInput, optFns ...func(*Options)) (*DescribeClientVpnTargetNetworksOutput, error) {
	stack := middleware.NewStack("DescribeClientVpnTargetNetworks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeClientVpnTargetNetworksMiddlewares(stack)
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
	addOpDescribeClientVpnTargetNetworksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClientVpnTargetNetworks(options.Region), middleware.Before)

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
			OperationName: "DescribeClientVpnTargetNetworks",
			Err:           err,
		}
	}
	out := result.(*DescribeClientVpnTargetNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClientVpnTargetNetworksInput struct {
	// One or more filters. Filter names and values are case-sensitive.
	//
	//     *
	// association-id - The ID of the association.
	//
	//     * target-network-id - The ID of
	// the subnet specified as the target network.
	//
	//     * vpc-id - The ID of the VPC in
	// which the target network is located.
	Filters []*types.Filter
	// The IDs of the target network associations.
	AssociationIds []*string
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string
	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int32
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeClientVpnTargetNetworksOutput struct {
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// Information about the associated target networks.
	ClientVpnTargetNetworks []*types.TargetNetwork

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeClientVpnTargetNetworksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeClientVpnTargetNetworks{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClientVpnTargetNetworks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClientVpnTargetNetworks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeClientVpnTargetNetworks",
	}
}