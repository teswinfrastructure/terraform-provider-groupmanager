package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/teswinfrastructure/terraform-provider-groupmanager/internal/common"

	"github.com/hashicorp/go-azure-sdk/sdk/claims"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"

	groups "github.com/teswinfrastructure/terraform-provider-groupmanager/internal/services/groups/client"
	serviceprincipals "github.com/teswinfrastructure/terraform-provider-groupmanager/internal/services/serviceprincipals/client"
	users "github.com/teswinfrastructure/terraform-provider-groupmanager/internal/services/users/client"
)

// Client contains the handles to all the specific Azure AD resource classes' respective clients
type Client struct {
	Environment environments.Environment
	TenantID    string
	ClientID    string
	ObjectID    string
	Claims      *claims.Claims

	TerraformVersion string

	StopContext context.Context

	Groups            *groups.Client
	ServicePrincipals *serviceprincipals.Client
	Users             *users.Client
}

func (client *Client) build(ctx context.Context, o *common.ClientOptions) error {
	client.StopContext = ctx
	client.Groups = groups.NewClient(o)
	client.ServicePrincipals = serviceprincipals.NewClient(o)
	client.Users = users.NewClient(o)

	// Acquire an access token upfront, so we can decode the JWT and populate the claims
	token, err := o.Authorizer.Token(ctx, &http.Request{})
	if err != nil {
		return fmt.Errorf("unable to obtain access token: %v", err)
	}

	client.Claims, err = claims.ParseClaims(token)
	if err != nil {
		return fmt.Errorf("unable to parse claims in access token: %v", err)
	}

	// Log the claims for debugging
	claimsJson, err := json.Marshal(client.Claims)
	if err != nil {
		log.Printf("[DEBUG] GroupManager Provider could not marshal access token claims for log outout")
	} else if claimsJson == nil {
		log.Printf("[DEBUG] GroupManager Provider marshaled access token claims was nil")
	} else {
		log.Printf("[DEBUG] GroupManager Provider access token claims: %s", claimsJson)
	}

	// Missing object ID of token holder will break many things
	client.ObjectID = client.Claims.ObjectId
	if client.ObjectID == "" {
		if strings.Contains(strings.ToLower(client.Claims.Scopes), "openid") {
			log.Printf("[DEBUG] Querying Microsoft Graph to discover authenticated user principal object ID because the `oid` claim was missing from the access token")
			result, _, err := client.Users.MeClient.Get(ctx, odata.Query{})
			if err != nil {
				return fmt.Errorf("attempting to discover object ID for authenticated user principal: %+v", err)
			}

			id := result.ID
			if id == nil {
				return fmt.Errorf("attempting to discover object ID for authenticated user principal: returned object ID was nil")
			}

			client.ObjectID = *id
		} else {
			log.Printf("[DEBUG] Querying Microsoft Graph to discover authenticated service principal object ID because the `oid` claim was missing from the access token")
			query := odata.Query{
				Filter: fmt.Sprintf("appId eq '%s'", client.ClientID),
			}
			result, _, err := client.ServicePrincipals.ServicePrincipalsClient.List(ctx, query)
			if err != nil {
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: %+v", err)
			}

			if len(*result) != 1 {
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: unexpected number of results, expected 1, received %d", len(*result))
			}

			id := (*result)[0].ID()
			if id == nil {
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: returned object ID was nil")
			}

			client.ObjectID = *id
		}
	}

	if client.ObjectID == "" {
		return fmt.Errorf("parsing claims in access token: oid claim is empty")
	}

	return nil
}
