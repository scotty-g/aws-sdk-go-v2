// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a website authorization provider from a specified fleet. After the
// disassociation, users can't load any associated websites that require this
// authorization provider.
func (c *Client) DisassociateWebsiteAuthorizationProvider(ctx context.Context, params *DisassociateWebsiteAuthorizationProviderInput, optFns ...func(*Options)) (*DisassociateWebsiteAuthorizationProviderOutput, error) {
	stack := middleware.NewStack("DisassociateWebsiteAuthorizationProvider", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateWebsiteAuthorizationProviderMiddlewares(stack)
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
	addOpDisassociateWebsiteAuthorizationProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateWebsiteAuthorizationProvider(options.Region), middleware.Before)

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
			OperationName: "DisassociateWebsiteAuthorizationProvider",
			Err:           err,
		}
	}
	out := result.(*DisassociateWebsiteAuthorizationProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateWebsiteAuthorizationProviderInput struct {
	// The ARN of the fleet.
	FleetArn *string
	// A unique identifier for the authorization provider.
	AuthorizationProviderId *string
}

type DisassociateWebsiteAuthorizationProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateWebsiteAuthorizationProviderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateWebsiteAuthorizationProvider{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateWebsiteAuthorizationProvider{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateWebsiteAuthorizationProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DisassociateWebsiteAuthorizationProvider",
	}
}