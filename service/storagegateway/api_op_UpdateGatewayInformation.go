// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a gateway's metadata, which includes the gateway's name and time zone.
// To specify which gateway to update, use the Amazon Resource Name (ARN) of the
// gateway in your request. For gateways activated after September 2, 2015, the
// gateway's ARN contains the gateway ID rather than the gateway name. However,
// changing the name of the gateway has no effect on the gateway's ARN.
func (c *Client) UpdateGatewayInformation(ctx context.Context, params *UpdateGatewayInformationInput, optFns ...func(*Options)) (*UpdateGatewayInformationOutput, error) {
	if params == nil {
		params = &UpdateGatewayInformationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGatewayInformation", params, optFns, c.addOperationUpdateGatewayInformationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGatewayInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGatewayInformationInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that you want
	// to use to monitor and log events in the gateway. For more information, see What
	// is Amazon CloudWatch Logs? (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/WhatIsCloudWatchLogs.html)
	CloudWatchLogGroupARN *string

	// Specifies the size of the gateway's metadata cache.
	GatewayCapacity types.GatewayCapacity

	// The name you configured for your gateway.
	GatewayName *string

	// A value that indicates the time zone of the gateway.
	GatewayTimezone *string

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway that was
// updated.
type UpdateGatewayInformationOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// The name you configured for your gateway.
	GatewayName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGatewayInformationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGatewayInformation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGatewayInformation{}, middleware.After)
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
	if err = addOpUpdateGatewayInformationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGatewayInformation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGatewayInformation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateGatewayInformation",
	}
}
