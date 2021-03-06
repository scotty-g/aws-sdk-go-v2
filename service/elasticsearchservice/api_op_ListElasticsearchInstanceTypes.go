// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all Elasticsearch instance types that are supported for given
// ElasticsearchVersion
func (c *Client) ListElasticsearchInstanceTypes(ctx context.Context, params *ListElasticsearchInstanceTypesInput, optFns ...func(*Options)) (*ListElasticsearchInstanceTypesOutput, error) {
	if params == nil {
		params = &ListElasticsearchInstanceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListElasticsearchInstanceTypes", params, optFns, addOperationListElasticsearchInstanceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListElasticsearchInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the ListElasticsearchInstanceTypes operation.
type ListElasticsearchInstanceTypesInput struct {

	// Version of Elasticsearch for which list of supported elasticsearch instance
	// types are needed.
	//
	// This member is required.
	ElasticsearchVersion *string

	// DomainName represents the name of the Domain that we are trying to modify. This
	// should be present only if we are querying for list of available Elasticsearch
	// instance types when modifying existing domain.
	DomainName *string

	// Set this value to limit the number of results returned. Value provided must be
	// greater than 30 else it wont be honored.
	MaxResults *int32

	// NextToken should be sent in case if earlier API call produced result containing
	// NextToken. It is used for pagination.
	NextToken *string
}

// Container for the parameters returned by ListElasticsearchInstanceTypes
// operation.
type ListElasticsearchInstanceTypesOutput struct {

	// List of instance types supported by Amazon Elasticsearch service for given
	// ElasticsearchVersion
	ElasticsearchInstanceTypes []types.ESPartitionInstanceType

	// In case if there are more results available NextToken would be present, make
	// further request to the same API with received NextToken to paginate remaining
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListElasticsearchInstanceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListElasticsearchInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListElasticsearchInstanceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListElasticsearchInstanceTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListElasticsearchInstanceTypes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListElasticsearchInstanceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListElasticsearchInstanceTypes",
	}
}
