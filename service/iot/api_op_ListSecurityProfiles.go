// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the Device Defender security profiles you have created. You can use
// filters to list only those security profiles associated with a thing group or
// only those associated with your account.
func (c *Client) ListSecurityProfiles(ctx context.Context, params *ListSecurityProfilesInput, optFns ...func(*Options)) (*ListSecurityProfilesOutput, error) {
	stack := middleware.NewStack("ListSecurityProfiles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSecurityProfilesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityProfiles(options.Region), middleware.Before)

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
			OperationName: "ListSecurityProfiles",
			Err:           err,
		}
	}
	out := result.(*ListSecurityProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityProfilesInput struct {
	// The token for the next set of results.
	NextToken *string
	// The maximum number of results to return at one time.
	MaxResults *int32
	// A filter to limit results to the security profiles that use the defined
	// dimension.
	DimensionName *string
}

type ListSecurityProfilesOutput struct {
	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string
	// A list of security profile identifiers (names and ARNs).
	SecurityProfileIdentifiers []*types.SecurityProfileIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSecurityProfilesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityProfiles{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityProfiles{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSecurityProfiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListSecurityProfiles",
	}
}