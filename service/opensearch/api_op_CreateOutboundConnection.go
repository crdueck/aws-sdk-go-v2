// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cross-cluster search connection from a source Amazon OpenSearch
// Service domain to a destination domain. For more information, see Cross-cluster
// search for Amazon OpenSearch Service (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html)
// .
func (c *Client) CreateOutboundConnection(ctx context.Context, params *CreateOutboundConnectionInput, optFns ...func(*Options)) (*CreateOutboundConnectionOutput, error) {
	if params == nil {
		params = &CreateOutboundConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOutboundConnection", params, optFns, c.addOperationCreateOutboundConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOutboundConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the CreateOutboundConnection operation.
type CreateOutboundConnectionInput struct {

	// Name of the connection.
	//
	// This member is required.
	ConnectionAlias *string

	// Name and Region of the source (local) domain.
	//
	// This member is required.
	LocalDomainInfo *types.DomainInformationContainer

	// Name and Region of the destination (remote) domain.
	//
	// This member is required.
	RemoteDomainInfo *types.DomainInformationContainer

	// The connection mode.
	ConnectionMode types.ConnectionMode

	noSmithyDocumentSerde
}

// The result of a CreateOutboundConnection request. Contains details about the
// newly created cross-cluster connection.
type CreateOutboundConnectionOutput struct {

	// Name of the connection.
	ConnectionAlias *string

	// The unique identifier for the created outbound connection, which is used for
	// subsequent operations on the connection.
	ConnectionId *string

	// The connection mode.
	ConnectionMode types.ConnectionMode

	// The ConnectionProperties for the newly created connection.
	ConnectionProperties *types.ConnectionProperties

	// The status of the connection.
	ConnectionStatus *types.OutboundConnectionStatus

	// Information about the source (local) domain.
	LocalDomainInfo *types.DomainInformationContainer

	// Information about the destination (remote) domain.
	RemoteDomainInfo *types.DomainInformationContainer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOutboundConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateOutboundConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOutboundConnection{}, middleware.After)
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
	if err = addOpCreateOutboundConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOutboundConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOutboundConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreateOutboundConnection",
	}
}
