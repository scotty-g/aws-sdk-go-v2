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

// Describes available services to which you can create a VPC endpoint.
func (c *Client) DescribeVpcEndpointServices(ctx context.Context, params *DescribeVpcEndpointServicesInput, optFns ...func(*Options)) (*DescribeVpcEndpointServicesOutput, error) {
	stack := middleware.NewStack("DescribeVpcEndpointServices", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVpcEndpointServicesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointServices(options.Region), middleware.Before)

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
			OperationName: "DescribeVpcEndpointServices",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcEndpointServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeVpcEndpointServices.
type DescribeVpcEndpointServicesInput struct {
	// The token for the next set of items to return. (You received this token from a
	// prior call.)
	NextToken *string
	// One or more filters.
	//
	//     * service-name - The name of the service.
	//
	//     * tag:
	// - The key/value combination of a tag assigned to the resource. Use the tag key
	// in the filter name and the tag value as the filter value. For example, to find
	// all resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	//     * tag-key -
	// The key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	Filters []*types.Filter
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// Constraint: If the value is greater than 1,000, we return only 1,000 items.
	MaxResults *int32
	// One or more service names.
	ServiceNames []*string
}

// Contains the output of DescribeVpcEndpointServices.
type DescribeVpcEndpointServicesOutput struct {
	// A list of supported services.
	ServiceNames []*string
	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string
	// Information about the service.
	ServiceDetails []*types.ServiceDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVpcEndpointServicesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointServices{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointServices{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcEndpointServices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointServices",
	}
}