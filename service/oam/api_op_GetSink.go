// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns complete information about one monitoring account sink. To use this
// operation, provide the sink ARN. To retrieve a list of sink ARNs, use ListSinks (https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html)
// .
func (c *Client) GetSink(ctx context.Context, params *GetSinkInput, optFns ...func(*Options)) (*GetSinkOutput, error) {
	if params == nil {
		params = &GetSinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSink", params, optFns, c.addOperationGetSinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSinkInput struct {

	// The ARN of the sink to retrieve information for.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetSinkOutput struct {

	// The ARN of the sink.
	Arn *string

	// The random ID string that Amazon Web Services generated as part of the sink ARN.
	Id *string

	// The name of the sink.
	Name *string

	// The tags assigned to the sink.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSink{}, middleware.After)
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
	if err = addOpGetSinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "oam",
		OperationName: "GetSink",
	}
}
