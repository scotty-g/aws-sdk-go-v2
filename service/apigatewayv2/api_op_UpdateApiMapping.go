// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The API mapping.
func (c *Client) UpdateApiMapping(ctx context.Context, params *UpdateApiMappingInput, optFns ...func(*Options)) (*UpdateApiMappingOutput, error) {
	stack := middleware.NewStack("UpdateApiMapping", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateApiMappingMiddlewares(stack)
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
	addOpUpdateApiMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApiMapping(options.Region), middleware.Before)

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
			OperationName: "UpdateApiMapping",
			Err:           err,
		}
	}
	out := result.(*UpdateApiMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates an ApiMapping.
type UpdateApiMappingInput struct {
	// The API identifier.
	ApiId *string
	// The domain name.
	DomainName *string
	// The API stage.
	Stage *string
	// The API mapping key.
	ApiMappingKey *string
	// The API mapping identifier.
	ApiMappingId *string
}

type UpdateApiMappingOutput struct {
	// The API stage.
	Stage *string
	// The API identifier.
	ApiId *string
	// The API mapping key.
	ApiMappingKey *string
	// The API mapping identifier.
	ApiMappingId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateApiMappingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApiMapping{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApiMapping{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateApiMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateApiMapping",
	}
}