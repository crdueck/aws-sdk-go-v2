// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows a subscription owner to set an attribute of the subscription to a new
// value.
func (c *Client) SetSubscriptionAttributes(ctx context.Context, params *SetSubscriptionAttributesInput, optFns ...func(*Options)) (*SetSubscriptionAttributesOutput, error) {
	if params == nil {
		params = &SetSubscriptionAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetSubscriptionAttributes", params, optFns, c.addOperationSetSubscriptionAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetSubscriptionAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetSubscriptionAttributes action.
type SetSubscriptionAttributesInput struct {

	// A map of attributes with their corresponding values. The following lists the
	// names, descriptions, and values of the special request parameters that this
	// action uses:
	//   - DeliveryPolicy – The policy that defines how Amazon SNS retries failed
	//   deliveries to HTTP/S endpoints.
	//   - FilterPolicy – The simple JSON object that lets your subscriber receive only
	//   a subset of messages, rather than receiving every message published to the
	//   topic.
	//   - FilterPolicyScope – This attribute lets you choose the filtering scope by
	//   using one of the following string value types:
	//   - MessageAttributes (default) – The filter is applied on the message
	//   attributes.
	//   - MessageBody – The filter is applied on the message body.
	//   - RawMessageDelivery – When set to true , enables raw message delivery to
	//   Amazon SQS or HTTP/S endpoints. This eliminates the need for the endpoints to
	//   process JSON formatting, which is otherwise created for Amazon SNS metadata.
	//   - RedrivePolicy – When specified, sends undeliverable messages to the
	//   specified Amazon SQS dead-letter queue. Messages that can't be delivered due to
	//   client errors (for example, when the subscribed endpoint is unreachable) or
	//   server errors (for example, when the service that powers the subscribed endpoint
	//   becomes unavailable) are held in the dead-letter queue for further analysis or
	//   reprocessing.
	// The following attribute applies only to Amazon Kinesis Data Firehose delivery
	// stream subscriptions:
	//   - SubscriptionRoleArn – The ARN of the IAM role that has the following:
	//   - Permission to write to the Kinesis Data Firehose delivery stream
	//   - Amazon SNS listed as a trusted entity Specifying a valid ARN for this
	//   attribute is required for Kinesis Data Firehose delivery stream subscriptions.
	//   For more information, see Fanout to Kinesis Data Firehose delivery streams (https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html)
	//   in the Amazon SNS Developer Guide.
	//
	// This member is required.
	AttributeName *string

	// The ARN of the subscription to modify.
	//
	// This member is required.
	SubscriptionArn *string

	// The new value for the attribute in JSON format.
	AttributeValue *string

	noSmithyDocumentSerde
}

type SetSubscriptionAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetSubscriptionAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetSubscriptionAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetSubscriptionAttributes{}, middleware.After)
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
	if err = addOpSetSubscriptionAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetSubscriptionAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetSubscriptionAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "SetSubscriptionAttributes",
	}
}
