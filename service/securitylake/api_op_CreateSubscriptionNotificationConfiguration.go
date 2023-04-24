// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Notifies the subscriber when new data is written to the data lake for the
// sources that the subscriber consumes in Security Lake. You can create only one
// subscriber notification per subscriber.
func (c *Client) CreateSubscriptionNotificationConfiguration(ctx context.Context, params *CreateSubscriptionNotificationConfigurationInput, optFns ...func(*Options)) (*CreateSubscriptionNotificationConfigurationOutput, error) {
	if params == nil {
		params = &CreateSubscriptionNotificationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscriptionNotificationConfiguration", params, optFns, c.addOperationCreateSubscriptionNotificationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriptionNotificationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubscriptionNotificationConfigurationInput struct {

	// The subscription ID for the notification subscription.
	//
	// This member is required.
	SubscriptionId *string

	// Create an Amazon Simple Queue Service queue.
	CreateSqs *bool

	// The key name for the notification subscription.
	HttpsApiKeyName *string

	// The key value for the notification subscription.
	HttpsApiKeyValue *string

	// The HTTPS method used for the notification subscription.
	HttpsMethod types.HttpsMethod

	// The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role
	// that you created. For more information about ARNs and how to use them in
	// policies, see Managing data access (https://docs.aws.amazon.com//security-lake/latest/userguide/subscriber-data-access.html)
	// and Amazon Web Services Managed Policies (https://docs.aws.amazon.com/security-lake/latest/userguide/security-iam-awsmanpol.html)
	// in the Amazon Security Lake User Guide.
	RoleArn *string

	// The subscription endpoint in Security Lake. If you prefer notification with an
	// HTTPs endpoint, populate this field.
	SubscriptionEndpoint *string

	noSmithyDocumentSerde
}

type CreateSubscriptionNotificationConfigurationOutput struct {

	// Returns the Amazon Resource Name (ARN) of the queue.
	QueueArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriptionNotificationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSubscriptionNotificationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSubscriptionNotificationConfiguration{}, middleware.After)
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
	if err = addOpCreateSubscriptionNotificationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscriptionNotificationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubscriptionNotificationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateSubscriptionNotificationConfiguration",
	}
}
