// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the user import jobs.
func (c *Client) ListUserImportJobs(ctx context.Context, params *ListUserImportJobsInput, optFns ...func(*Options)) (*ListUserImportJobsOutput, error) {
	if params == nil {
		params = &ListUserImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserImportJobs", params, optFns, c.addOperationListUserImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list the user import jobs.
type ListUserImportJobsInput struct {

	// The maximum number of import jobs you want the request to return.
	//
	// This member is required.
	MaxResults int32

	// The user pool ID for the user pool that the users are being imported into.
	//
	// This member is required.
	UserPoolId *string

	// An identifier that was returned from the previous call to ListUserImportJobs ,
	// which can be used to return the next set of import jobs in the list.
	PaginationToken *string

	noSmithyDocumentSerde
}

// Represents the response from the server to the request to list the user import
// jobs.
type ListUserImportJobsOutput struct {

	// An identifier that can be used to return the next set of user import jobs in
	// the list.
	PaginationToken *string

	// The user import jobs.
	UserImportJobs []types.UserImportJobType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUserImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUserImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUserImportJobs{}, middleware.After)
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
	if err = addOpListUserImportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserImportJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListUserImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUserImportJobs",
	}
}
