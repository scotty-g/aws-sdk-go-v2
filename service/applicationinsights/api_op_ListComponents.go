// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the auto-grouped, standalone, and custom components of the application.
func (c *Client) ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) {
	stack := middleware.NewStack("ListComponents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListComponentsMiddlewares(stack)
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
	addOpListComponentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListComponents(options.Region), middleware.Before)

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
			OperationName: "ListComponents",
			Err:           err,
		}
	}
	out := result.(*ListComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentsInput struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32
	// The name of the resource group.
	ResourceGroupName *string
	// The token to request the next page of results.
	NextToken *string
}

type ListComponentsOutput struct {
	// The token to request the next page of results.
	NextToken *string
	// The list of application components.
	ApplicationComponentList []*types.ApplicationComponent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListComponentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListComponents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComponents{}, middleware.After)
}

func newServiceMetadataMiddleware_opListComponents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListComponents",
	}
}