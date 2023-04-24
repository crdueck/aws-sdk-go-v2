// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables or disables integration with a service that can be integrated with
// DevOps Guru. The one service that can be integrated with DevOps Guru is Amazon
// CodeGuru Profiler, which can produce proactive recommendations which can be
// stored and viewed in DevOps Guru.
func (c *Client) UpdateEventSourcesConfig(ctx context.Context, params *UpdateEventSourcesConfigInput, optFns ...func(*Options)) (*UpdateEventSourcesConfigOutput, error) {
	if params == nil {
		params = &UpdateEventSourcesConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventSourcesConfig", params, optFns, c.addOperationUpdateEventSourcesConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventSourcesConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEventSourcesConfigInput struct {

	// Configuration information about the integration of DevOps Guru as the Consumer
	// via EventBridge with another AWS Service.
	EventSources *types.EventSourcesConfig

	noSmithyDocumentSerde
}

type UpdateEventSourcesConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEventSourcesConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEventSourcesConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEventSourcesConfig{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEventSourcesConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventSourcesConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "UpdateEventSourcesConfig",
	}
}
