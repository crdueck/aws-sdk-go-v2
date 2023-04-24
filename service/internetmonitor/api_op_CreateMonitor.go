// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a monitor in Amazon CloudWatch Internet Monitor. A monitor is built
// based on information from the application resources that you add: Amazon Virtual
// Private Clouds (VPCs), Amazon CloudFront distributions, and WorkSpaces
// directories. Internet Monitor then publishes internet measurements from Amazon
// Web Services that are specific to the city-networks, that is, the locations and
// ASNs (typically internet service providers or ISPs), where clients access your
// application. For more information, see Using Amazon CloudWatch Internet Monitor (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-InternetMonitor.html)
// in the Amazon CloudWatch User Guide. When you create a monitor, you set a
// maximum limit for the number of city-networks where client traffic is monitored.
// The city-network maximum that you choose is the limit, but you only pay for the
// number of city-networks that are actually monitored. You can change the maximum
// at any time by updating your monitor. For more information, see Choosing a
// city-network maximum value (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html)
// in the Amazon CloudWatch User Guide.
func (c *Client) CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) {
	if params == nil {
		params = &CreateMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMonitor", params, optFns, c.addOperationCreateMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMonitorInput struct {

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// A unique, case-sensitive string of up to 64 ASCII characters that you specify
	// to make an idempotent API request. Don't reuse the same client token for other
	// API requests.
	ClientToken *string

	// Publish internet measurements for Internet Monitor to an Amazon S3 bucket in
	// addition to CloudWatch Logs.
	InternetMeasurementsLogDelivery *types.InternetMeasurementsLogDelivery

	// The maximum number of city-networks to monitor for your resources. A
	// city-network is the location (city) where clients access your application
	// resources from and the network or ASN, such as an internet service provider
	// (ISP), that clients access the resources through. This limit helps control
	// billing costs. To learn more, see Choosing a city-network maximum value  (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	MaxCityNetworksToMonitor int32

	// The resources to include in a monitor, which you provide as a set of Amazon
	// Resource Names (ARNs). You can add a combination of Amazon Virtual Private
	// Clouds (VPCs) and Amazon CloudFront distributions, or you can add Amazon
	// WorkSpaces directories. You can't add all three types of resources. If you add
	// only VPC resources, at least one VPC must have an Internet Gateway attached to
	// it, to make sure that it has internet connectivity.
	Resources []string

	// The tags for a monitor. You can add a maximum of 50 tags in Internet Monitor.
	Tags map[string]string

	// The percentage of the internet-facing traffic for your application that you
	// want to monitor with this monitor.
	TrafficPercentageToMonitor int32

	noSmithyDocumentSerde
}

type CreateMonitorOutput struct {

	// The Amazon Resource Name (ARN) of the monitor.
	//
	// This member is required.
	Arn *string

	// The status of a monitor.
	//
	// This member is required.
	Status types.MonitorConfigState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMonitor{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateMonitorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMonitor(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMonitor struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMonitor) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMonitorInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMonitorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMonitor{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "internetmonitor",
		OperationName: "CreateMonitor",
	}
}
