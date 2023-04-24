// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The following example serializes a payload that uses an XML namespace.
func (c *Client) HttpPayloadWithXmlNamespaceAndPrefix(ctx context.Context, params *HttpPayloadWithXmlNamespaceAndPrefixInput, optFns ...func(*Options)) (*HttpPayloadWithXmlNamespaceAndPrefixOutput, error) {
	if params == nil {
		params = &HttpPayloadWithXmlNamespaceAndPrefixInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPayloadWithXmlNamespaceAndPrefix", params, optFns, c.addOperationHttpPayloadWithXmlNamespaceAndPrefixMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPayloadWithXmlNamespaceAndPrefixOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadWithXmlNamespaceAndPrefixInput struct {
	Nested *types.PayloadWithXmlNamespaceAndPrefix

	noSmithyDocumentSerde
}

type HttpPayloadWithXmlNamespaceAndPrefixOutput struct {
	Nested *types.PayloadWithXmlNamespaceAndPrefix

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHttpPayloadWithXmlNamespaceAndPrefixMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHttpPayloadWithXmlNamespaceAndPrefix{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHttpPayloadWithXmlNamespaceAndPrefix{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespaceAndPrefix(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opHttpPayloadWithXmlNamespaceAndPrefix(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadWithXmlNamespaceAndPrefix",
	}
}
