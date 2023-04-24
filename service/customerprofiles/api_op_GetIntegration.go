// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns an integration for a domain.
func (c *Client) GetIntegration(ctx context.Context, params *GetIntegrationInput, optFns ...func(*Options)) (*GetIntegrationOutput, error) {
	if params == nil {
		params = &GetIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntegration", params, optFns, c.addOperationGetIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntegrationInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The URI of the S3 bucket or any other type of data source.
	//
	// This member is required.
	Uri *string

	noSmithyDocumentSerde
}

type GetIntegrationOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The URI of the S3 bucket or any other type of data source.
	//
	// This member is required.
	Uri *string

	// Boolean that shows if the Flow that's associated with the Integration is
	// created in Amazon Appflow, or with ObjectTypeName equals _unstructured via
	// API/CLI in flowDefinition.
	IsUnstructured *bool

	// The name of the profile object type.
	ObjectTypeName *string

	// A map in which each key is an event type from an external application such as
	// Segment or Shopify, and each value is an ObjectTypeName (template) used to
	// ingest the event. It supports the following event types: SegmentIdentify ,
	// ShopifyCreateCustomers , ShopifyUpdateCustomers , ShopifyCreateDraftOrders ,
	// ShopifyUpdateDraftOrders , ShopifyCreateOrders , and ShopifyUpdatedOrders .
	ObjectTypeNames map[string]string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Unique identifier for the workflow.
	WorkflowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntegration{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntegration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "GetIntegration",
	}
}
