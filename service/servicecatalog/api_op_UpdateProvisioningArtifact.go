// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified provisioning artifact (also known as a version) for the
// specified product. You cannot update a provisioning artifact for a product that
// was shared with you.
func (c *Client) UpdateProvisioningArtifact(ctx context.Context, params *UpdateProvisioningArtifactInput, optFns ...func(*Options)) (*UpdateProvisioningArtifactOutput, error) {
	stack := middleware.NewStack("UpdateProvisioningArtifact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateProvisioningArtifactMiddlewares(stack)
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
	addOpUpdateProvisioningArtifactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisioningArtifact(options.Region), middleware.Before)

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
			OperationName: "UpdateProvisioningArtifact",
			Err:           err,
		}
	}
	out := result.(*UpdateProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisioningArtifactInput struct {
	// The product identifier.
	ProductId *string
	// The updated description of the provisioning artifact.
	Description *string
	// Information set by the administrator to provide guidance to end users about
	// which provisioning artifacts to use. The DEFAULT value indicates that the
	// product version is active. The administrator can set the guidance to DEPRECATED
	// to inform users that the product version is deprecated. Users are able to make
	// updates to a provisioned product of a deprecated version but cannot launch new
	// provisioned products using a deprecated version.
	Guidance types.ProvisioningArtifactGuidance
	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string
	// The updated name of the provisioning artifact.
	Name *string
	// Indicates whether the product version is active. Inactive provisioning artifacts
	// are invisible to end users. End users cannot launch or update a provisioned
	// product from an inactive provisioning artifact.
	Active *bool
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type UpdateProvisioningArtifactOutput struct {
	// The status of the current request.
	Status types.Status
	// Information about the provisioning artifact.
	ProvisioningArtifactDetail *types.ProvisioningArtifactDetail
	// The URL of the CloudFormation template in Amazon S3.
	Info map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateProvisioningArtifactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProvisioningArtifact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProvisioningArtifact{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "UpdateProvisioningArtifact",
	}
}