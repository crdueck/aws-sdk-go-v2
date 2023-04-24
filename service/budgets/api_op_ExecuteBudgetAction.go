// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Executes a budget action.
func (c *Client) ExecuteBudgetAction(ctx context.Context, params *ExecuteBudgetActionInput, optFns ...func(*Options)) (*ExecuteBudgetActionOutput, error) {
	if params == nil {
		params = &ExecuteBudgetActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteBudgetAction", params, optFns, c.addOperationExecuteBudgetActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteBudgetActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteBudgetActionInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	// A string that represents the budget name. The ":" and "\" characters aren't
	// allowed.
	//
	// This member is required.
	BudgetName *string

	// The type of execution.
	//
	// This member is required.
	ExecutionType types.ExecutionType

	noSmithyDocumentSerde
}

type ExecuteBudgetActionOutput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	// A string that represents the budget name. The ":" and "\" characters aren't
	// allowed.
	//
	// This member is required.
	BudgetName *string

	// The type of execution.
	//
	// This member is required.
	ExecutionType types.ExecutionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteBudgetActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExecuteBudgetAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExecuteBudgetAction{}, middleware.After)
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
	if err = addOpExecuteBudgetActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteBudgetAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExecuteBudgetAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "ExecuteBudgetAction",
	}
}
