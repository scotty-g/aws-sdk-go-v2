// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables a user in the user pool. After being enabled, users can sign in to
// AppStream 2.0 and open applications from the stacks to which they are assigned.
func (c *Client) EnableUser(ctx context.Context, params *EnableUserInput, optFns ...func(*Options)) (*EnableUserOutput, error) {
	stack := middleware.NewStack("EnableUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableUserMiddlewares(stack)
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
	addOpEnableUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableUser(options.Region), middleware.Before)

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
			OperationName: "EnableUser",
			Err:           err,
		}
	}
	out := result.(*EnableUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableUserInput struct {
	// The authentication type for the user. You must specify USERPOOL.
	AuthenticationType types.AuthenticationType
	// The email address of the user.  <note> <p>Users' email addresses are
	// case-sensitive. During login, if they specify an email address that doesn't use
	// the same capitalization as the email address specified when their user pool
	// account was created, a "user does not exist" error message displays. </p>
	// </note>
	UserName *string
}

type EnableUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "EnableUser",
	}
}