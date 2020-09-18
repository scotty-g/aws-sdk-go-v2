// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns summary information about the versions of a type.
func (c *Client) ListTypeVersions(ctx context.Context, params *ListTypeVersionsInput, optFns ...func(*Options)) (*ListTypeVersionsOutput, error) {
	stack := middleware.NewStack("ListTypeVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListTypeVersionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTypeVersions(options.Region), middleware.Before)

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
			OperationName: "ListTypeVersions",
			Err:           err,
		}
	}
	out := result.(*ListTypeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypeVersionsInput struct {
	// The Amazon Resource Name (ARN) of the type for which you want version summary
	// information. Conditional: You must specify either TypeName and Type, or Arn.
	Arn *string
	// The name of the type for which you want version summary information.
	// Conditional: You must specify either TypeName and Type, or Arn.
	TypeName *string
	// The deprecation status of the type versions that you want to get summary
	// information about. Valid values include:
	//
	//     * LIVE: The type version is
	// registered and can be used in CloudFormation operations, dependent on its
	// provisioning behavior and visibility scope.
	//
	//     * DEPRECATED: The type version
	// has been deregistered and can no longer be used in CloudFormation
	// operations.
	//
	// The default is LIVE.
	DeprecatedStatus types.DeprecatedStatus
	// The kind of the type. Currently the only valid value is RESOURCE. Conditional:
	// You must specify either TypeName and Type, or Arn.
	Type types.RegistryType
	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results.
	MaxResults *int32
}

type ListTypeVersionsOutput struct {
	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call this action again and assign
	// that token to the request object's NextToken parameter. If the request returns
	// all results, NextToken is set to null.
	NextToken *string
	// A list of TypeVersionSummary structures that contain information about the
	// specified type's versions.
	TypeVersionSummaries []*types.TypeVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListTypeVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListTypeVersions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListTypeVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTypeVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListTypeVersions",
	}
}