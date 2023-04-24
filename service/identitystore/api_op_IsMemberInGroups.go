// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks the user's membership in all requested groups and returns if the member
// exists in all queried groups.
func (c *Client) IsMemberInGroups(ctx context.Context, params *IsMemberInGroupsInput, optFns ...func(*Options)) (*IsMemberInGroupsOutput, error) {
	if params == nil {
		params = &IsMemberInGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IsMemberInGroups", params, optFns, c.addOperationIsMemberInGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IsMemberInGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IsMemberInGroupsInput struct {

	// A list of identifiers for groups in the identity store.
	//
	// This member is required.
	GroupIds []string

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// An object containing the identifier of a group member.
	//
	// This member is required.
	MemberId types.MemberId

	noSmithyDocumentSerde
}

type IsMemberInGroupsOutput struct {

	// A list containing the results of membership existence checks.
	//
	// This member is required.
	Results []types.GroupMembershipExistenceResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationIsMemberInGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpIsMemberInGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpIsMemberInGroups{}, middleware.After)
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
	if err = addOpIsMemberInGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIsMemberInGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIsMemberInGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "identitystore",
		OperationName: "IsMemberInGroups",
	}
}
