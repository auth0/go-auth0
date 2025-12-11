# Reference
## Actions
<details><summary><code>client.Actions.List() -> *management.ListActionsPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all actions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListActionsRequestParameters{
        TriggerID: management.String(
            "triggerId",
        ),
        ActionName: management.String(
            "actionName",
        ),
        Deployed: management.Bool(
            true,
        ),
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        Installed: management.Bool(
            true,
        ),
    }
client.Actions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**triggerID:** `*management.ActionTriggerTypeEnum` — An actions extensibility point.
    
</dd>
</dl>

<dl>
<dd>

**actionName:** `*string` — The name of the action to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**deployed:** `*bool` — Optional filter to only retrieve actions that are deployed.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Use this field to request a specific page of the list results.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — The maximum number of results to be returned by the server in single response. 20 by default
    
</dd>
</dl>

<dl>
<dd>

**installed:** `*bool` — Optional. When true, return only installed actions. When false, return only custom actions. Returns all actions by default.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Create(request) -> *management.CreateActionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an action. Once an action is created, it must be deployed, and then bound to a trigger before it will be executed as part of a flow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateActionRequestContent{
        Name: "name",
        SupportedTriggers: []*management.ActionTrigger{
            &management.ActionTrigger{
                ID: "id",
            },
        },
    }
client.Actions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of an action.
    
</dd>
</dl>

<dl>
<dd>

**supportedTriggers:** `[]*management.ActionTrigger` — The list of triggers that this action supports. At this time, an action can only target a single trigger at a time.
    
</dd>
</dl>

<dl>
<dd>

**code:** `*string` — The source code of the action.
    
</dd>
</dl>

<dl>
<dd>

**dependencies:** `[]*management.ActionVersionDependency` — The list of third party npm modules, and their versions, that this action depends on.
    
</dd>
</dl>

<dl>
<dd>

**runtime:** `*string` — The Node runtime. For example: `node22`, defaults to `node22`
    
</dd>
</dl>

<dl>
<dd>

**secrets:** `[]*management.ActionSecretRequest` — The list of secrets that are included in an action or a version of an action.
    
</dd>
</dl>

<dl>
<dd>

**deploy:** `*bool` — True if the action should be deployed after creation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Get(ID) -> *management.GetActionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve an action by its ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Actions.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the action to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes an action and all of its associated versions. An action must be unbound from all triggers before it can be deleted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteActionRequestParameters{
        Force: management.Bool(
            true,
        ),
    }
client.Actions.Delete(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the action to delete.
    
</dd>
</dl>

<dl>
<dd>

**force:** `*bool` — Force action deletion detaching bindings
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Update(ID, request) -> *management.UpdateActionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing action. If this action is currently bound to a trigger, updating it will <strong>not</strong> affect any user flows until the action is deployed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateActionRequestContent{}
client.Actions.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the action to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of an action.
    
</dd>
</dl>

<dl>
<dd>

**supportedTriggers:** `[]*management.ActionTrigger` — The list of triggers that this action supports. At this time, an action can only target a single trigger at a time.
    
</dd>
</dl>

<dl>
<dd>

**code:** `*string` — The source code of the action.
    
</dd>
</dl>

<dl>
<dd>

**dependencies:** `[]*management.ActionVersionDependency` — The list of third party npm modules, and their versions, that this action depends on.
    
</dd>
</dl>

<dl>
<dd>

**runtime:** `*string` — The Node runtime. For example: `node22`, defaults to `node22`
    
</dd>
</dl>

<dl>
<dd>

**secrets:** `[]*management.ActionSecretRequest` — The list of secrets that are included in an action or a version of an action.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Deploy(ID) -> *management.DeployActionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deploy an action. Deploying an action will create a new immutable version of the action. If the action is currently bound to a trigger, then the system will begin executing the newly deployed version of the action immediately. Otherwise, the action will only be executed as a part of a flow once it is bound to that flow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Actions.Deploy(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of an action.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Test(ID, request) -> *management.TestActionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Test an action. After updating an action, it can be tested prior to being deployed to ensure it behaves as expected.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.TestActionRequestContent{
        Payload: map[string]any{
            "key": "value",
        },
    }
client.Actions.Test(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the action to test.
    
</dd>
</dl>

<dl>
<dd>

**payload:** `management.TestActionPayload` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Branding
<details><summary><code>client.Branding.Get() -> *management.GetBrandingResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve branding settings.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Update(request) -> *management.UpdateBrandingResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update branding settings.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateBrandingRequestContent{}
client.Branding.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**colors:** `*management.UpdateBrandingColors` 
    
</dd>
</dl>

<dl>
<dd>

**faviconURL:** `*string` — URL for the favicon. Must use HTTPS.
    
</dd>
</dl>

<dl>
<dd>

**logoURL:** `*string` — URL for the logo. Must use HTTPS.
    
</dd>
</dl>

<dl>
<dd>

**font:** `*management.UpdateBrandingFont` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ClientGrants
<details><summary><code>client.ClientGrants.List() -> *management.ListClientGrantPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of <a href="https://auth0.com/docs/get-started/applications/application-access-to-apis-client-grants">client grants</a>, including the scopes associated with the application/API pair.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListClientGrantsRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
        Audience: management.String(
            "audience",
        ),
        ClientID: management.String(
            "client_id",
        ),
        AllowAnyOrganization: management.Bool(
            true,
        ),
        SubjectType: management.ClientGrantSubjectTypeEnumClient.Ptr(),
    }
client.ClientGrants.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**audience:** `*string` — Optional filter on audience.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — Optional filter on client_id.
    
</dd>
</dl>

<dl>
<dd>

**allowAnyOrganization:** `*management.ClientGrantAllowAnyOrganizationEnum` — Optional filter on allow_any_organization.
    
</dd>
</dl>

<dl>
<dd>

**subjectType:** `*management.ClientGrantSubjectTypeEnum` — The type of application access the client grant allows. Use of this field is subject to the applicable Free Trial terms in Okta’s <a href="https://www.okta.com/legal/"> Master Subscription Agreement.</a>
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ClientGrants.Create(request) -> *management.CreateClientGrantResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a client grant for a machine-to-machine login flow. To learn more, read <a href="https://www.auth0.com/docs/get-started/authentication-and-authorization-flow/client-credentials-flow">Client Credential Flow</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateClientGrantRequestContent{
        ClientID: "client_id",
        Audience: "audience",
    }
client.ClientGrants.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — ID of the client.
    
</dd>
</dl>

<dl>
<dd>

**audience:** `string` — The audience (API identifier) of this client grant
    
</dd>
</dl>

<dl>
<dd>

**organizationUsage:** `*management.ClientGrantOrganizationUsageEnum` 
    
</dd>
</dl>

<dl>
<dd>

**allowAnyOrganization:** `*bool` — If enabled, any organization can be used with this grant. If disabled (default), the grant must be explicitly assigned to the desired organizations.
    
</dd>
</dl>

<dl>
<dd>

**scope:** `[]string` — Scopes allowed for this client grant.
    
</dd>
</dl>

<dl>
<dd>

**subjectType:** `*management.ClientGrantSubjectTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**authorizationDetailsTypes:** `[]string` — Types of authorization_details allowed for this client grant. Use of this field is subject to the applicable Free Trial terms in Okta’s <a href= "https://www.okta.com/legal/"> Master Subscription Agreement.</a>
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ClientGrants.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete the <a href="https://www.auth0.com/docs/get-started/authentication-and-authorization-flow/client-credentials-flow">Client Credential Flow</a> from your machine-to-machine application.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ClientGrants.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client grant to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ClientGrants.Update(ID, request) -> *management.UpdateClientGrantResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a client grant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateClientGrantRequestContent{}
client.ClientGrants.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client grant to update.
    
</dd>
</dl>

<dl>
<dd>

**scope:** `[]string` — Scopes allowed for this client grant.
    
</dd>
</dl>

<dl>
<dd>

**organizationUsage:** `*management.ClientGrantOrganizationNullableUsageEnum` 
    
</dd>
</dl>

<dl>
<dd>

**allowAnyOrganization:** `*bool` — Controls allowing any organization to be used with this grant
    
</dd>
</dl>

<dl>
<dd>

**authorizationDetailsTypes:** `[]string` — Types of authorization_details allowed for this client grant. Use of this field is subject to the applicable Free Trial terms in Okta’s <a href= "https://www.okta.com/legal/"> Master Subscription Agreement.</a>
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Clients
<details><summary><code>client.Clients.List() -> *management.ListClientsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve clients (applications and SSO integrations) matching provided filters. A list of fields to include or exclude may also be specified.
For more information, read <a href="https://www.auth0.com/docs/get-started/applications"> Applications in Auth0</a> and <a href="https://www.auth0.com/docs/authenticate/single-sign-on"> Single Sign-On</a>.

<ul>
  <li>
    The following can be retrieved with any scope:
    <code>client_id</code>, <code>app_type</code>, <code>name</code>, and <code>description</code>.
  </li>
  <li>
    The following properties can only be retrieved with the <code>read:clients</code> or
    <code>read:client_keys</code> scope:
    <code>callbacks</code>, <code>oidc_logout</code>, <code>allowed_origins</code>,
    <code>web_origins</code>, <code>tenant</code>, <code>global</code>, <code>config_route</code>,
    <code>callback_url_template</code>, <code>jwt_configuration</code>,
    <code>jwt_configuration.lifetime_in_seconds</code>, <code>jwt_configuration.secret_encoded</code>,
    <code>jwt_configuration.scopes</code>, <code>jwt_configuration.alg</code>, <code>api_type</code>,
    <code>logo_uri</code>, <code>allowed_clients</code>, <code>owners</code>, <code>custom_login_page</code>,
    <code>custom_login_page_off</code>, <code>sso</code>, <code>addons</code>, <code>form_template</code>,
    <code>custom_login_page_codeview</code>, <code>resource_servers</code>, <code>client_metadata</code>,
    <code>mobile</code>, <code>mobile.android</code>, <code>mobile.ios</code>, <code>allowed_logout_urls</code>,
    <code>token_endpoint_auth_method</code>, <code>is_first_party</code>, <code>oidc_conformant</code>,
    <code>is_token_endpoint_ip_header_trusted</code>, <code>initiate_login_uri</code>, <code>grant_types</code>,
    <code>refresh_token</code>, <code>refresh_token.rotation_type</code>, <code>refresh_token.expiration_type</code>,
    <code>refresh_token.leeway</code>, <code>refresh_token.token_lifetime</code>, <code>refresh_token.policies</code>, <code>organization_usage</code>,
    <code>organization_require_behavior</code>.
  </li>
  <li>
    The following properties can only be retrieved with the
    <code>read:client_keys</code> or <code>read:client_credentials</code> scope:
    <code>encryption_key</code>, <code>encryption_key.pub</code>, <code>encryption_key.cert</code>,
    <code>client_secret</code>, <code>client_authentication_methods</code> and <code>signing_key</code>.
  </li>
</ul>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListClientsRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        IsGlobal: management.Bool(
            true,
        ),
        IsFirstParty: management.Bool(
            true,
        ),
        AppType: management.String(
            "app_type",
        ),
        Q: management.String(
            "q",
        ),
    }
client.Clients.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Default value is 50, maximum value is 100
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**isGlobal:** `*bool` — Optional filter on the global client parameter.
    
</dd>
</dl>

<dl>
<dd>

**isFirstParty:** `*bool` — Optional filter on whether or not a client is a first-party client.
    
</dd>
</dl>

<dl>
<dd>

**appType:** `*string` — Optional filter by a comma-separated list of application types.
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Advanced Query in <a href="http://www.lucenetutorial.com/lucene-query-syntax.html">Lucene</a> syntax.<br /><b>Permitted Queries</b>:<br /><ul><li><i>client_grant.organization_id:{organization_id}</i></li><li><i>client_grant.allow_any_organization:true</i></li></ul><b>Additional Restrictions</b>:<br /><ul><li>Cannot be used in combination with other filters</li><li>Requires use of the <i>from</i> and <i>take</i> paging parameters (checkpoint paginatinon)</li><li>Reduced rate limits apply. See <a href="https://auth0.com/docs/troubleshoot/customer-support/operational-policies/rate-limit-policy/rate-limit-configurations/enterprise-public">Rate Limit Configurations</a></li></ul><i><b>Note</b>: Recent updates may not be immediately reflected in query results</i>
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Create(request) -> *management.CreateClientResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new client (application or SSO integration). For more information, read <a href="https://www.auth0.com/docs/get-started/auth0-overview/create-applications">Create Applications</a>
<a href="https://www.auth0.com/docs/authenticate/single-sign-on/api-endpoints-for-single-sign-on>">API Endpoints for Single Sign-On</a>. 

Notes: 
- We recommend leaving the `client_secret` parameter unspecified to allow the generation of a safe secret.
- The <code>client_authentication_methods</code> and <code>token_endpoint_auth_method</code> properties are mutually exclusive. Use 
<code>client_authentication_methods</code> to configure the client with Private Key JWT authentication method. Otherwise, use <code>token_endpoint_auth_method</code>
to configure the client with client secret (basic or post) or with no authentication method (none).
- When using <code>client_authentication_methods</code> to configure the client with Private Key JWT authentication method, specify fully defined credentials. 
These credentials will be automatically enabled for Private Key JWT authentication on the client. 
- To configure <code>client_authentication_methods</code>, the <code>create:client_credentials</code> scope is required.
- To configure <code>client_authentication_methods</code>, the property <code>jwt_configuration.alg</code> must be set to RS256.

<div class="alert alert-warning">SSO Integrations created via this endpoint will accept login requests and share user profile information.</div>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateClientRequestContent{
        Name: "name",
    }
client.Clients.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Name of this client (min length: 1 character, does not allow `<` or `>`).
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Free text description of this client (max length: 140 characters).
    
</dd>
</dl>

<dl>
<dd>

**logoURI:** `*string` — URL of the logo to display for this client. Recommended size is 150x150 pixels.
    
</dd>
</dl>

<dl>
<dd>

**callbacks:** `[]string` — Comma-separated list of URLs whitelisted for Auth0 to use as a callback to the client after authentication.
    
</dd>
</dl>

<dl>
<dd>

**oidcLogout:** `*management.ClientOidcBackchannelLogoutSettings` 
    
</dd>
</dl>

<dl>
<dd>

**oidcBackchannelLogout:** `*management.ClientOidcBackchannelLogoutSettings` 
    
</dd>
</dl>

<dl>
<dd>

**sessionTransfer:** `*management.ClientSessionTransferConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**allowedOrigins:** `[]string` — Comma-separated list of URLs allowed to make requests from JavaScript to Auth0 API (typically used with CORS). By default, all your callback URLs will be allowed. This field allows you to enter other origins if necessary. You can also use wildcards at the subdomain level (e.g., https://*.contoso.com). Query strings and hash information are not taken into account when validating these URLs.
    
</dd>
</dl>

<dl>
<dd>

**webOrigins:** `[]string` — Comma-separated list of allowed origins for use with <a href='https://auth0.com/docs/cross-origin-authentication'>Cross-Origin Authentication</a>, <a href='https://auth0.com/docs/flows/concepts/device-auth'>Device Flow</a>, and <a href='https://auth0.com/docs/protocols/oauth2#how-response-mode-works'>web message response mode</a>.
    
</dd>
</dl>

<dl>
<dd>

**clientAliases:** `[]string` — List of audiences/realms for SAML protocol. Used by the wsfed addon.
    
</dd>
</dl>

<dl>
<dd>

**allowedClients:** `[]string` — List of allow clients and API ids that are allowed to make delegation requests. Empty means all all your clients are allowed.
    
</dd>
</dl>

<dl>
<dd>

**allowedLogoutURLs:** `[]string` — Comma-separated list of URLs that are valid to redirect to after logout from Auth0. Wildcards are allowed for subdomains.
    
</dd>
</dl>

<dl>
<dd>

**grantTypes:** `[]string` — List of grant types supported for this application. Can include `authorization_code`, `implicit`, `refresh_token`, `client_credentials`, `password`, `http://auth0.com/oauth/grant-type/password-realm`, `http://auth0.com/oauth/grant-type/mfa-oob`, `http://auth0.com/oauth/grant-type/mfa-otp`, `http://auth0.com/oauth/grant-type/mfa-recovery-code`, `urn:openid:params:grant-type:ciba`, `urn:ietf:params:oauth:grant-type:device_code`, and `urn:auth0:params:oauth:grant-type:token-exchange:federated-connection-access-token`.
    
</dd>
</dl>

<dl>
<dd>

**tokenEndpointAuthMethod:** `*management.ClientTokenEndpointAuthMethodEnum` 
    
</dd>
</dl>

<dl>
<dd>

**isTokenEndpointIPHeaderTrusted:** `*bool` — If true, trust that the IP specified in the `auth0-forwarded-for` header is the end-user's IP for brute-force-protection on token endpoint.
    
</dd>
</dl>

<dl>
<dd>

**appType:** `*management.ClientAppTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**isFirstParty:** `*bool` — Whether this client a first party client or not
    
</dd>
</dl>

<dl>
<dd>

**oidcConformant:** `*bool` — Whether this client conforms to <a href='https://auth0.com/docs/api-auth/tutorials/adoption'>strict OIDC specifications</a> (true) or uses legacy features (false).
    
</dd>
</dl>

<dl>
<dd>

**jwtConfiguration:** `*management.ClientJwtConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**encryptionKey:** `*management.ClientEncryptionKey` 
    
</dd>
</dl>

<dl>
<dd>

**sso:** `*bool` — Applies only to SSO clients and determines whether Auth0 will handle Single Sign On (true) or whether the Identity Provider will (false).
    
</dd>
</dl>

<dl>
<dd>

**crossOriginAuthentication:** `*bool` — Whether this client can be used to make cross-origin authentication requests (true) or it is not allowed to make such requests (false).
    
</dd>
</dl>

<dl>
<dd>

**crossOriginLoc:** `*string` — URL of the location in your site where the cross origin verification takes place for the cross-origin auth flow when performing Auth in your own domain instead of Auth0 hosted login page.
    
</dd>
</dl>

<dl>
<dd>

**ssoDisabled:** `*bool` — <code>true</code> to disable Single Sign On, <code>false</code> otherwise (default: <code>false</code>)
    
</dd>
</dl>

<dl>
<dd>

**customLoginPageOn:** `*bool` — <code>true</code> if the custom login page is to be used, <code>false</code> otherwise. Defaults to <code>true</code>
    
</dd>
</dl>

<dl>
<dd>

**customLoginPage:** `*string` — The content (HTML, CSS, JS) of the custom login page.
    
</dd>
</dl>

<dl>
<dd>

**customLoginPagePreview:** `*string` — The content (HTML, CSS, JS) of the custom login page. (Used on Previews)
    
</dd>
</dl>

<dl>
<dd>

**formTemplate:** `*string` — HTML form template to be used for WS-Federation.
    
</dd>
</dl>

<dl>
<dd>

**addons:** `*management.ClientAddons` 
    
</dd>
</dl>

<dl>
<dd>

**clientMetadata:** `*management.ClientMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**mobile:** `*management.ClientMobile` 
    
</dd>
</dl>

<dl>
<dd>

**initiateLoginURI:** `*string` — Initiate login uri, must be https
    
</dd>
</dl>

<dl>
<dd>

**nativeSocialLogin:** `*management.NativeSocialLogin` 
    
</dd>
</dl>

<dl>
<dd>

**refreshToken:** `*management.ClientRefreshTokenConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**defaultOrganization:** `*management.ClientDefaultOrganization` 
    
</dd>
</dl>

<dl>
<dd>

**organizationUsage:** `*management.ClientOrganizationUsageEnum` 
    
</dd>
</dl>

<dl>
<dd>

**organizationRequireBehavior:** `*management.ClientOrganizationRequireBehaviorEnum` 
    
</dd>
</dl>

<dl>
<dd>

**organizationDiscoveryMethods:** `[]*management.ClientOrganizationDiscoveryEnum` — Defines the available methods for organization discovery during the `pre_login_prompt`. Users can discover their organization either by `email`, `organization_name` or both.
    
</dd>
</dl>

<dl>
<dd>

**clientAuthenticationMethods:** `*management.ClientCreateAuthenticationMethod` 
    
</dd>
</dl>

<dl>
<dd>

**requirePushedAuthorizationRequests:** `*bool` — Makes the use of Pushed Authorization Requests mandatory for this client
    
</dd>
</dl>

<dl>
<dd>

**requireProofOfPossession:** `*bool` — Makes the use of Proof-of-Possession mandatory for this client
    
</dd>
</dl>

<dl>
<dd>

**signedRequestObject:** `*management.ClientSignedRequestObjectWithPublicKey` 
    
</dd>
</dl>

<dl>
<dd>

**complianceLevel:** `*management.ClientComplianceLevelEnum` 
    
</dd>
</dl>

<dl>
<dd>

**skipNonVerifiableCallbackURIConfirmationPrompt:** `*bool` 

Controls whether a confirmation prompt is shown during login flows when the redirect URI uses non-verifiable callback URIs (for example, a custom URI schema such as `myapp://`, or `localhost`).
If set to true, a confirmation prompt will not be shown. We recommend that this is set to false for improved protection from malicious apps.
See https://auth0.com/docs/secure/security-guidance/measures-against-app-impersonation for more information.
    
</dd>
</dl>

<dl>
<dd>

**tokenExchange:** `*management.ClientTokenExchangeConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**parRequestExpiry:** `*int` — Specifies how long, in seconds, a Pushed Authorization Request URI remains valid
    
</dd>
</dl>

<dl>
<dd>

**tokenQuota:** `*management.CreateTokenQuota` 
    
</dd>
</dl>

<dl>
<dd>

**resourceServerIdentifier:** `*string` — The identifier of the resource server that this client is linked to.
    
</dd>
</dl>

<dl>
<dd>

**expressConfiguration:** `*management.ExpressConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**asyncApprovalNotificationChannels:** `*management.ClientAsyncApprovalNotificationsChannelsAPIPostConfiguration` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Get(ID) -> *management.GetClientResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve client details by ID. Clients are SSO connections or Applications linked with your Auth0 tenant. A list of fields to include or exclude may also be specified. 
For more information, read <a href="https://www.auth0.com/docs/get-started/applications"> Applications in Auth0</a> and <a href="https://www.auth0.com/docs/authenticate/single-sign-on"> Single Sign-On</a>.
<ul>
  <li>
    The following properties can be retrieved with any of the scopes:
    <code>client_id</code>, <code>app_type</code>, <code>name</code>, and <code>description</code>.
  </li>
  <li>
    The following properties can only be retrieved with the <code>read:clients</code> or
    <code>read:client_keys</code> scopes:
    <code>callbacks</code>, <code>oidc_logout</code>, <code>allowed_origins</code>,
    <code>web_origins</code>, <code>tenant</code>, <code>global</code>, <code>config_route</code>,
    <code>callback_url_template</code>, <code>jwt_configuration</code>,
    <code>jwt_configuration.lifetime_in_seconds</code>, <code>jwt_configuration.secret_encoded</code>,
    <code>jwt_configuration.scopes</code>, <code>jwt_configuration.alg</code>, <code>api_type</code>,
    <code>logo_uri</code>, <code>allowed_clients</code>, <code>owners</code>, <code>custom_login_page</code>,
    <code>custom_login_page_off</code>, <code>sso</code>, <code>addons</code>, <code>form_template</code>,
    <code>custom_login_page_codeview</code>, <code>resource_servers</code>, <code>client_metadata</code>,
    <code>mobile</code>, <code>mobile.android</code>, <code>mobile.ios</code>, <code>allowed_logout_urls</code>,
    <code>token_endpoint_auth_method</code>, <code>is_first_party</code>, <code>oidc_conformant</code>,
    <code>is_token_endpoint_ip_header_trusted</code>, <code>initiate_login_uri</code>, <code>grant_types</code>,
    <code>refresh_token</code>, <code>refresh_token.rotation_type</code>, <code>refresh_token.expiration_type</code>,
    <code>refresh_token.leeway</code>, <code>refresh_token.token_lifetime</code>, <code>refresh_token.policies</code>, <code>organization_usage</code>,
    <code>organization_require_behavior</code>.
  </li>
  <li>
    The following properties can only be retrieved with the <code>read:client_keys</code> or <code>read:client_credentials</code> scopes:
    <code>encryption_key</code>, <code>encryption_key.pub</code>, <code>encryption_key.cert</code>,
    <code>client_secret</code>, <code>client_authentication_methods</code> and <code>signing_key</code>.
  </li>
</ul>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetClientRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Clients.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a client and related configuration (rules, connections, etc).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Clients.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Update(ID, request) -> *management.UpdateClientResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a client's settings. For more information, read <a href="https://www.auth0.com/docs/get-started/applications"> Applications in Auth0</a> and <a href="https://www.auth0.com/docs/authenticate/single-sign-on"> Single Sign-On</a>.

Notes:
- The `client_secret` and `signing_key` attributes can only be updated with the `update:client_keys` scope.
- The <code>client_authentication_methods</code> and <code>token_endpoint_auth_method</code> properties are mutually exclusive. Use <code>client_authentication_methods</code> to configure the client with Private Key JWT authentication method. Otherwise, use <code>token_endpoint_auth_method</code> to configure the client with client secret (basic or post) or with no authentication method (none).
- When using <code>client_authentication_methods</code> to configure the client with Private Key JWT authentication method, only specify the credential IDs that were generated when creating the credentials on the client.
- To configure <code>client_authentication_methods</code>, the <code>update:client_credentials</code> scope is required.
- To configure <code>client_authentication_methods</code>, the property <code>jwt_configuration.alg</code> must be set to RS256.
- To change a client's <code>is_first_party</code> property to <code>false</code>, the <code>organization_usage</code> and <code>organization_require_behavior</code> properties must be unset.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateClientRequestContent{}
client.Clients.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the client. Must contain at least one character. Does not allow '<' or '>'.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Free text description of the purpose of the Client. (Max character length: <code>140</code>)
    
</dd>
</dl>

<dl>
<dd>

**clientSecret:** `*string` — The secret used to sign tokens for the client
    
</dd>
</dl>

<dl>
<dd>

**logoURI:** `*string` — The URL of the client logo (recommended size: 150x150)
    
</dd>
</dl>

<dl>
<dd>

**callbacks:** `[]string` — A set of URLs that are valid to call back from Auth0 when authenticating users
    
</dd>
</dl>

<dl>
<dd>

**oidcLogout:** `*management.ClientOidcBackchannelLogoutSettings` 
    
</dd>
</dl>

<dl>
<dd>

**oidcBackchannelLogout:** `*management.ClientOidcBackchannelLogoutSettings` 
    
</dd>
</dl>

<dl>
<dd>

**sessionTransfer:** `*management.ClientSessionTransferConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**allowedOrigins:** `[]string` — A set of URLs that represents valid origins for CORS
    
</dd>
</dl>

<dl>
<dd>

**webOrigins:** `[]string` — A set of URLs that represents valid web origins for use with web message response mode
    
</dd>
</dl>

<dl>
<dd>

**grantTypes:** `[]string` — A set of grant types that the client is authorized to use. Can include `authorization_code`, `implicit`, `refresh_token`, `client_credentials`, `password`, `http://auth0.com/oauth/grant-type/password-realm`, `http://auth0.com/oauth/grant-type/mfa-oob`, `http://auth0.com/oauth/grant-type/mfa-otp`, `http://auth0.com/oauth/grant-type/mfa-recovery-code`, `urn:openid:params:grant-type:ciba`, `urn:ietf:params:oauth:grant-type:device_code`, and `urn:auth0:params:oauth:grant-type:token-exchange:federated-connection-access-token`.
    
</dd>
</dl>

<dl>
<dd>

**clientAliases:** `[]string` — List of audiences for SAML protocol
    
</dd>
</dl>

<dl>
<dd>

**allowedClients:** `[]string` — Ids of clients that will be allowed to perform delegation requests. Clients that will be allowed to make delegation request. By default, all your clients will be allowed. This field allows you to specify specific clients
    
</dd>
</dl>

<dl>
<dd>

**allowedLogoutURLs:** `[]string` — URLs that are valid to redirect to after logout from Auth0.
    
</dd>
</dl>

<dl>
<dd>

**jwtConfiguration:** `*management.ClientJwtConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**encryptionKey:** `*management.ClientEncryptionKey` 
    
</dd>
</dl>

<dl>
<dd>

**sso:** `*bool` — <code>true</code> to use Auth0 instead of the IdP to do Single Sign On, <code>false</code> otherwise (default: <code>false</code>)
    
</dd>
</dl>

<dl>
<dd>

**crossOriginAuthentication:** `*bool` — <code>true</code> if this client can be used to make cross-origin authentication requests, <code>false</code> otherwise if cross origin is disabled
    
</dd>
</dl>

<dl>
<dd>

**crossOriginLoc:** `*string` — URL for the location in your site where the cross origin verification takes place for the cross-origin auth flow when performing Auth in your own domain instead of Auth0 hosted login page.
    
</dd>
</dl>

<dl>
<dd>

**ssoDisabled:** `*bool` — <code>true</code> to disable Single Sign On, <code>false</code> otherwise (default: <code>false</code>)
    
</dd>
</dl>

<dl>
<dd>

**customLoginPageOn:** `*bool` — <code>true</code> if the custom login page is to be used, <code>false</code> otherwise.
    
</dd>
</dl>

<dl>
<dd>

**tokenEndpointAuthMethod:** `*management.ClientTokenEndpointAuthMethodOrNullEnum` 
    
</dd>
</dl>

<dl>
<dd>

**isTokenEndpointIPHeaderTrusted:** `*bool` — If true, trust that the IP specified in the `auth0-forwarded-for` header is the end-user's IP for brute-force-protection on token endpoint.
    
</dd>
</dl>

<dl>
<dd>

**appType:** `*management.ClientAppTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**isFirstParty:** `*bool` — Whether this client a first party client or not
    
</dd>
</dl>

<dl>
<dd>

**oidcConformant:** `*bool` — Whether this client will conform to strict OIDC specifications
    
</dd>
</dl>

<dl>
<dd>

**customLoginPage:** `*string` — The content (HTML, CSS, JS) of the custom login page
    
</dd>
</dl>

<dl>
<dd>

**customLoginPagePreview:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**tokenQuota:** `*management.UpdateTokenQuota` 
    
</dd>
</dl>

<dl>
<dd>

**formTemplate:** `*string` — Form template for WS-Federation protocol
    
</dd>
</dl>

<dl>
<dd>

**addons:** `*management.ClientAddons` 
    
</dd>
</dl>

<dl>
<dd>

**clientMetadata:** `*management.ClientMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**mobile:** `*management.ClientMobile` 
    
</dd>
</dl>

<dl>
<dd>

**initiateLoginURI:** `*string` — Initiate login uri, must be https
    
</dd>
</dl>

<dl>
<dd>

**nativeSocialLogin:** `*management.NativeSocialLogin` 
    
</dd>
</dl>

<dl>
<dd>

**refreshToken:** `*management.ClientRefreshTokenConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**defaultOrganization:** `*management.ClientDefaultOrganization` 
    
</dd>
</dl>

<dl>
<dd>

**organizationUsage:** `*management.ClientOrganizationUsagePatchEnum` 
    
</dd>
</dl>

<dl>
<dd>

**organizationRequireBehavior:** `*management.ClientOrganizationRequireBehaviorPatchEnum` 
    
</dd>
</dl>

<dl>
<dd>

**organizationDiscoveryMethods:** `[]*management.ClientOrganizationDiscoveryEnum` — Defines the available methods for organization discovery during the `pre_login_prompt`. Users can discover their organization either by `email`, `organization_name` or both.
    
</dd>
</dl>

<dl>
<dd>

**clientAuthenticationMethods:** `*management.ClientAuthenticationMethod` 
    
</dd>
</dl>

<dl>
<dd>

**requirePushedAuthorizationRequests:** `*bool` — Makes the use of Pushed Authorization Requests mandatory for this client
    
</dd>
</dl>

<dl>
<dd>

**requireProofOfPossession:** `*bool` — Makes the use of Proof-of-Possession mandatory for this client
    
</dd>
</dl>

<dl>
<dd>

**signedRequestObject:** `*management.ClientSignedRequestObjectWithCredentialID` 
    
</dd>
</dl>

<dl>
<dd>

**complianceLevel:** `*management.ClientComplianceLevelEnum` 
    
</dd>
</dl>

<dl>
<dd>

**skipNonVerifiableCallbackURIConfirmationPrompt:** `*bool` 

Controls whether a confirmation prompt is shown during login flows when the redirect URI uses non-verifiable callback URIs (for example, a custom URI schema such as `myapp://`, or `localhost`).
If set to true, a confirmation prompt will not be shown. We recommend that this is set to false for improved protection from malicious apps.
See https://auth0.com/docs/secure/security-guidance/measures-against-app-impersonation for more information.
    
</dd>
</dl>

<dl>
<dd>

**tokenExchange:** `*management.ClientTokenExchangeConfigurationOrNull` 
    
</dd>
</dl>

<dl>
<dd>

**parRequestExpiry:** `*int` — Specifies how long, in seconds, a Pushed Authorization Request URI remains valid
    
</dd>
</dl>

<dl>
<dd>

**expressConfiguration:** `*management.ExpressConfigurationOrNull` 
    
</dd>
</dl>

<dl>
<dd>

**asyncApprovalNotificationChannels:** `*management.ClientAsyncApprovalNotificationsChannelsAPIPatchConfiguration` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.RotateSecret(ID) -> *management.RotateClientSecretResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Rotate a client secret.

This endpoint cannot be used with clients configured with Private Key JWT authentication method (client_authentication_methods configured with private_key_jwt). The generated secret is NOT base64 encoded.

For more information, read <a href="https://www.auth0.com/docs/get-started/applications/rotate-client-secret">Rotate Client Secrets</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Clients.RotateSecret(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client that will rotate secrets.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ConnectionProfiles
<details><summary><code>client.ConnectionProfiles.List() -> *management.ListConnectionProfilesPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of Connection Profiles. This endpoint supports Checkpoint pagination.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListConnectionProfileRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.ConnectionProfiles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 5.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectionProfiles.Create(request) -> *management.CreateConnectionProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a Connection Profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateConnectionProfileRequestContent{
        Name: "name",
    }
client.ConnectionProfiles.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `management.ConnectionProfileName` 
    
</dd>
</dl>

<dl>
<dd>

**organization:** `*management.ConnectionProfileOrganization` 
    
</dd>
</dl>

<dl>
<dd>

**connectionNamePrefixTemplate:** `*management.ConnectionNamePrefixTemplate` 
    
</dd>
</dl>

<dl>
<dd>

**enabledFeatures:** `*management.ConnectionProfileEnabledFeatures` 
    
</dd>
</dl>

<dl>
<dd>

**connectionConfig:** `*management.ConnectionProfileConfig` 
    
</dd>
</dl>

<dl>
<dd>

**strategyOverrides:** `*management.ConnectionProfileStrategyOverrides` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectionProfiles.ListTemplates() -> *management.ListConnectionProfileTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of Connection Profile Templates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ConnectionProfiles.ListTemplates(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectionProfiles.GetTemplate(ID) -> *management.GetConnectionProfileTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a Connection Profile Template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ConnectionProfiles.GetTemplate(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the connection-profile-template to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectionProfiles.Get(ID) -> *management.GetConnectionProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a single Connection Profile specified by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ConnectionProfiles.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the connection-profile to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectionProfiles.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a single Connection Profile specified by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ConnectionProfiles.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the connection-profile to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConnectionProfiles.Update(ID, request) -> *management.UpdateConnectionProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the details of a specific Connection Profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateConnectionProfileRequestContent{}
client.ConnectionProfiles.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the connection profile to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*management.ConnectionProfileName` 
    
</dd>
</dl>

<dl>
<dd>

**organization:** `*management.ConnectionProfileOrganization` 
    
</dd>
</dl>

<dl>
<dd>

**connectionNamePrefixTemplate:** `*management.ConnectionNamePrefixTemplate` 
    
</dd>
</dl>

<dl>
<dd>

**enabledFeatures:** `*management.ConnectionProfileEnabledFeatures` 
    
</dd>
</dl>

<dl>
<dd>

**connectionConfig:** `*management.ConnectionProfileConfig` 
    
</dd>
</dl>

<dl>
<dd>

**strategyOverrides:** `*management.ConnectionProfileStrategyOverrides` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections
<details><summary><code>client.Connections.List() -> *management.ListConnectionsCheckpointPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves detailed list of all <a href="https://auth0.com/docs/authenticate/identity-providers">connections</a> that match the specified strategy. If no strategy is provided, all connections within your tenant are retrieved. This action can accept a list of fields to include or exclude from the resulting list of connections. 

This endpoint supports two types of pagination:
<ul>
<li>Offset pagination</li>
<li>Checkpoint pagination</li>
</ul>

Checkpoint pagination must be used if you need to retrieve more than 1000 connections.

<h2>Checkpoint Pagination</h2>

To search by checkpoint, use the following parameters:
<ul>
<li><code>from</code>: Optional id from which to start selection.</li>
<li><code>take</code>: The total amount of entries to retrieve when using the from parameter. Defaults to 50.</li>
</ul>

<b>Note</b>: The first time you call this endpoint using checkpoint pagination, omit the <code>from</code> parameter. If there are more results, a <code>next</code> value is included in the response. You can use this for subsequent API calls. When <code>next</code> is no longer included in the response, no pages are remaining.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListConnectionsQueryParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
        Name: management.String(
            "name",
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Connections.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**strategy:** `*management.ConnectionStrategyEnum` — Provide strategies to only retrieve connections with such strategies
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Provide the name of the connection to retrieve
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise (defaults to <code>true</code>)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Create(request) -> *management.CreateConnectionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a new connection according to the JSON object received in <code>body</code>.<br/>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateConnectionRequestContent{
        Name: "name",
        Strategy: management.ConnectionIdentityProviderEnumAd,
    }
client.Connections.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the connection. Must start and end with an alphanumeric character and can only contain alphanumeric characters and '-'. Max length 128
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` — Connection name used in the new universal login experience
    
</dd>
</dl>

<dl>
<dd>

**strategy:** `*management.ConnectionIdentityProviderEnum` 
    
</dd>
</dl>

<dl>
<dd>

**options:** `*management.ConnectionPropertiesOptions` 
    
</dd>
</dl>

<dl>
<dd>

**enabledClients:** `[]string` — DEPRECATED property. Use the PATCH /v2/connections/{id}/clients endpoint to enable the connection for a set of clients.
    
</dd>
</dl>

<dl>
<dd>

**isDomainConnection:** `*bool` — <code>true</code> promotes to a domain-level connection so that third-party applications can use it. <code>false</code> does not promote the connection, so only first-party applications with the connection enabled can use it. (Defaults to <code>false</code>.)
    
</dd>
</dl>

<dl>
<dd>

**showAsButton:** `*bool` — Enables showing a button for the connection in the login page (new experience only). If false, it will be usable only by HRD. (Defaults to <code>false</code>.)
    
</dd>
</dl>

<dl>
<dd>

**realms:** `[]string` — Defines the realms for which the connection will be used (ie: email domains). If the array is empty or the property is not specified, the connection name will be added as realm.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `*management.ConnectionsMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**authentication:** `*management.ConnectionAuthenticationPurpose` 
    
</dd>
</dl>

<dl>
<dd>

**connectedAccounts:** `*management.ConnectionConnectedAccountsPurpose` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Get(ID) -> *management.GetConnectionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details for a specified <a href="https://auth0.com/docs/authenticate/identity-providers">connection</a> along with options that can be used for identity provider configuration.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetConnectionRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Connections.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to retrieve
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise (defaults to <code>true</code>)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Removes a specific <a href="https://auth0.com/docs/authenticate/identity-providers">connection</a> from your tenant. This action cannot be undone. Once removed, users can no longer use this connection to authenticate.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Update(ID, request) -> *management.UpdateConnectionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update details for a specific <a href="https://auth0.com/docs/authenticate/identity-providers">connection</a>, including option properties for identity provider configuration.

<b>Note</b>: If you use the <code>options</code> parameter, the entire <code>options</code> object is overriden. To avoid partial data or other issues, ensure all parameters are present when using this option.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateConnectionRequestContent{}
client.Connections.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to update
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` — The connection name used in the new universal login experience. If display_name is not included in the request, the field will be overwritten with the name value.
    
</dd>
</dl>

<dl>
<dd>

**options:** `*management.UpdateConnectionOptions` 
    
</dd>
</dl>

<dl>
<dd>

**enabledClients:** `[]string` — DEPRECATED property. Use the PATCH /v2/connections/{id}/clients endpoint to enable or disable the connection for any clients.
    
</dd>
</dl>

<dl>
<dd>

**isDomainConnection:** `*bool` — <code>true</code> promotes to a domain-level connection so that third-party applications can use it. <code>false</code> does not promote the connection, so only first-party applications with the connection enabled can use it. (Defaults to <code>false</code>.)
    
</dd>
</dl>

<dl>
<dd>

**showAsButton:** `*bool` — Enables showing a button for the connection in the login page (new experience only). If false, it will be usable only by HRD. (Defaults to <code>false</code>.)
    
</dd>
</dl>

<dl>
<dd>

**realms:** `[]string` — Defines the realms for which the connection will be used (ie: email domains). If the array is empty or the property is not specified, the connection name will be added as realm.
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `*management.ConnectionsMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**authentication:** `*management.ConnectionAuthenticationPurpose` 
    
</dd>
</dl>

<dl>
<dd>

**connectedAccounts:** `*management.ConnectionConnectedAccountsPurpose` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.CheckStatus(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the status of an ad/ldap connection referenced by its <code>ID</code>. <code>200 OK</code> http status code response is returned  when the connection is online, otherwise a <code>404</code> status code is returned along with an error message
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.CheckStatus(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the connection to check
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## CustomDomains
<details><summary><code>client.CustomDomains.List() -> management.ListCustomDomainsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details on <a href="https://auth0.com/docs/custom-domains">custom domains</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListCustomDomainsRequestParameters{
        Q: management.String(
            "q",
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        Sort: management.String(
            "sort",
        ),
    }
client.CustomDomains.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `*string` — Query in <a href ="http://www.lucenetutorial.com/lucene-query-syntax.html">Lucene query string syntax</a>.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Field to sort by. Only <code>domain:1</code> (ascending order by domain) is supported at this time.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CustomDomains.Create(request) -> *management.CreateCustomDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new custom domain.

Note: The custom domain will need to be verified before it will accept
requests.

Optional attributes that can be updated:

- custom_client_ip_header
- tls_policy


TLS Policies:

- recommended - for modern usage this includes TLS 1.2 only
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateCustomDomainRequestContent{
        Domain: "domain",
        Type: management.CustomDomainProvisioningTypeEnumAuth0ManagedCerts,
    }
client.CustomDomains.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**domain:** `string` — Domain name.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*management.CustomDomainProvisioningTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**verificationMethod:** `*management.CustomDomainVerificationMethodEnum` 
    
</dd>
</dl>

<dl>
<dd>

**tlsPolicy:** `*management.CustomDomainTLSPolicyEnum` 
    
</dd>
</dl>

<dl>
<dd>

**customClientIPHeader:** `*management.CustomDomainCustomClientIPHeader` 
    
</dd>
</dl>

<dl>
<dd>

**domainMetadata:** `*management.DomainMetadata` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CustomDomains.Get(ID) -> *management.GetCustomDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a custom domain configuration and status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.CustomDomains.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the custom domain to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CustomDomains.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a custom domain and stop serving requests for it.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.CustomDomains.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the custom domain to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CustomDomains.Update(ID, request) -> *management.UpdateCustomDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a custom domain.

These are the attributes that can be updated:

- custom_client_ip_header
- tls_policy

<h5>Updating CUSTOM_CLIENT_IP_HEADER for a custom domain</h5>To update the <code>custom_client_ip_header</code> for a domain, the body to
send should be:
<pre><code>{ "custom_client_ip_header": "cf-connecting-ip" }</code></pre>

<h5>Updating TLS_POLICY for a custom domain</h5>To update the <code>tls_policy</code> for a domain, the body to send should be:
<pre><code>{ "tls_policy": "recommended" }</code></pre>


TLS Policies:

- recommended - for modern usage this includes TLS 1.2 only


Some considerations:

- The TLS ciphers and protocols available in each TLS policy follow industry recommendations, and may be updated occasionally.
- The <code>compatible</code> TLS policy is no longer supported.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateCustomDomainRequestContent{}
client.CustomDomains.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the custom domain to update
    
</dd>
</dl>

<dl>
<dd>

**tlsPolicy:** `*management.CustomDomainTLSPolicyEnum` 
    
</dd>
</dl>

<dl>
<dd>

**customClientIPHeader:** `*management.CustomDomainCustomClientIPHeader` 
    
</dd>
</dl>

<dl>
<dd>

**domainMetadata:** `*management.DomainMetadata` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CustomDomains.Test(ID) -> *management.TestCustomDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Run the test process on a custom domain.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.CustomDomains.Test(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the custom domain to test.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.CustomDomains.Verify(ID) -> *management.VerifyCustomDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Run the verification process on a custom domain.

Note: Check the <code>status</code> field to see its verification status. Once verification is complete, it may take up to 10 minutes before the custom domain can start accepting requests.

For <code>self_managed_certs</code>, when the custom domain is verified for the first time, the response will also include the <code>cname_api_key</code> which you will need to configure your proxy. This key must be kept secret, and is used to validate the proxy requests.

<a href="https://auth0.com/docs/custom-domains#step-2-verify-ownership">Learn more</a> about verifying custom domains that use Auth0 Managed certificates.
<a href="https://auth0.com/docs/custom-domains/self-managed-certificates#step-2-verify-ownership">Learn more</a> about verifying custom domains that use Self Managed certificates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.CustomDomains.Verify(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the custom domain to verify.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## DeviceCredentials
<details><summary><code>client.DeviceCredentials.List() -> *management.ListDeviceCredentialsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve device credential information (<code>public_key</code>, <code>refresh_token</code>, or <code>rotating_refresh_token</code>) associated with a specific user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListDeviceCredentialsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        UserID: management.String(
            "user_id",
        ),
        ClientID: management.String(
            "client_id",
        ),
        Type: management.DeviceCredentialTypeEnumPublicKey.Ptr(),
    }
client.DeviceCredentials.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.  There is a maximum of 1000 results allowed from this endpoint.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — user_id of the devices to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — client_id of the devices to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*management.DeviceCredentialTypeEnum` — Type of credentials to retrieve. Must be `public_key`, `refresh_token` or `rotating_refresh_token`. The property will default to `refresh_token` when paging is requested
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DeviceCredentials.CreatePublicKey(request) -> *management.CreatePublicKeyDeviceCredentialResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a device credential public key to manage refresh token rotation for a given <code>user_id</code>. Device Credentials APIs are designed for ad-hoc administrative use only and paging is by default enabled for GET requests.

When refresh token rotation is enabled, the endpoint becomes consistent. For more information, read <a href="https://auth0.com/docs/get-started/tenant-settings/signing-keys"> Signing Keys</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreatePublicKeyDeviceCredentialRequestContent{
        DeviceName: "device_name",
        Type: management.DeviceCredentialPublicKeyTypeEnum(
            "public_key",
        ),
        Value: "value",
        DeviceID: "device_id",
    }
client.DeviceCredentials.CreatePublicKey(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**deviceName:** `string` — Name for this device easily recognized by owner.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `management.DeviceCredentialPublicKeyTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**value:** `string` — Base64 encoded string containing the credential.
    
</dd>
</dl>

<dl>
<dd>

**deviceID:** `string` — Unique identifier for the device. Recommend using <a href="http://developer.android.com/reference/android/provider/Settings.Secure.html#ANDROID_ID">Android_ID</a> on Android and <a href="https://developer.apple.com/library/ios/documentation/UIKit/Reference/UIDevice_Class/index.html#//apple_ref/occ/instp/UIDevice/identifierForVendor">identifierForVendor</a>.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — client_id of the client (application) this credential is for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.DeviceCredentials.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Permanently delete a device credential (such as a refresh token or public key) with the given ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.DeviceCredentials.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the credential to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## EmailTemplates
<details><summary><code>client.EmailTemplates.Create(request) -> *management.CreateEmailTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an email template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateEmailTemplateRequestContent{
        Template: management.EmailTemplateNameEnumVerifyEmail,
    }
client.EmailTemplates.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**template:** `*management.EmailTemplateNameEnum` 
    
</dd>
</dl>

<dl>
<dd>

**body:** `*string` — Body of the email template.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Senders `from` email address.
    
</dd>
</dl>

<dl>
<dd>

**resultURL:** `*string` — URL to redirect the user to after a successful action.
    
</dd>
</dl>

<dl>
<dd>

**subject:** `*string` — Subject line of the email.
    
</dd>
</dl>

<dl>
<dd>

**syntax:** `*string` — Syntax of the template body.
    
</dd>
</dl>

<dl>
<dd>

**urlLifetimeInSeconds:** `*float64` — Lifetime in seconds that the link within the email will be valid for.
    
</dd>
</dl>

<dl>
<dd>

**includeEmailInRedirect:** `*bool` — Whether the `reset_email` and `verify_email` templates should include the user's email address as the `email` parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the template is enabled (true) or disabled (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailTemplates.Get(TemplateName) -> *management.GetEmailTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve an email template by pre-defined name. These names are `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, and `async_approval`. The names `change_password`, and `password_reset` are also supported for legacy scenarios.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.EmailTemplates.Get(
        context.TODO(),
        management.EmailTemplateNameEnumVerifyEmail.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateName:** `*management.EmailTemplateNameEnum` — Template name. Can be `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, `async_approval`, `change_password` (legacy), or `password_reset` (legacy).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailTemplates.Set(TemplateName, request) -> *management.SetEmailTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an email template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetEmailTemplateRequestContent{
        Template: management.EmailTemplateNameEnumVerifyEmail,
    }
client.EmailTemplates.Set(
        context.TODO(),
        management.EmailTemplateNameEnumVerifyEmail.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateName:** `*management.EmailTemplateNameEnum` — Template name. Can be `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, `async_approval`, `change_password` (legacy), or `password_reset` (legacy).
    
</dd>
</dl>

<dl>
<dd>

**template:** `*management.EmailTemplateNameEnum` 
    
</dd>
</dl>

<dl>
<dd>

**body:** `*string` — Body of the email template.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Senders `from` email address.
    
</dd>
</dl>

<dl>
<dd>

**resultURL:** `*string` — URL to redirect the user to after a successful action.
    
</dd>
</dl>

<dl>
<dd>

**subject:** `*string` — Subject line of the email.
    
</dd>
</dl>

<dl>
<dd>

**syntax:** `*string` — Syntax of the template body.
    
</dd>
</dl>

<dl>
<dd>

**urlLifetimeInSeconds:** `*float64` — Lifetime in seconds that the link within the email will be valid for.
    
</dd>
</dl>

<dl>
<dd>

**includeEmailInRedirect:** `*bool` — Whether the `reset_email` and `verify_email` templates should include the user's email address as the `email` parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the template is enabled (true) or disabled (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EmailTemplates.Update(TemplateName, request) -> *management.UpdateEmailTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify an email template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateEmailTemplateRequestContent{}
client.EmailTemplates.Update(
        context.TODO(),
        management.EmailTemplateNameEnumVerifyEmail.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**templateName:** `*management.EmailTemplateNameEnum` — Template name. Can be `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, `async_approval`, `change_password` (legacy), or `password_reset` (legacy).
    
</dd>
</dl>

<dl>
<dd>

**template:** `*management.EmailTemplateNameEnum` 
    
</dd>
</dl>

<dl>
<dd>

**body:** `*string` — Body of the email template.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Senders `from` email address.
    
</dd>
</dl>

<dl>
<dd>

**resultURL:** `*string` — URL to redirect the user to after a successful action.
    
</dd>
</dl>

<dl>
<dd>

**subject:** `*string` — Subject line of the email.
    
</dd>
</dl>

<dl>
<dd>

**syntax:** `*string` — Syntax of the template body.
    
</dd>
</dl>

<dl>
<dd>

**urlLifetimeInSeconds:** `*float64` — Lifetime in seconds that the link within the email will be valid for.
    
</dd>
</dl>

<dl>
<dd>

**includeEmailInRedirect:** `*bool` — Whether the `reset_email` and `verify_email` templates should include the user's email address as the `email` parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the template is enabled (true) or disabled (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## EventStreams
<details><summary><code>client.EventStreams.List() -> []*management.EventStreamResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListEventStreamsRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.EventStreams.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EventStreams.Create(request) -> *management.CreateEventStreamResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.EventStreamsCreateRequest{
        CreateEventStreamWebHookRequestContent: &management.CreateEventStreamWebHookRequestContent{
            Destination: &management.EventStreamWebhookDestination{
                Type: management.EventStreamWebhookDestinationTypeEnum(
                    "webhook",
                ),
                Configuration: &management.EventStreamWebhookConfiguration{
                    WebhookEndpoint: "webhook_endpoint",
                    WebhookAuthorization: &management.EventStreamWebhookAuthorizationResponse{
                        EventStreamWebhookBasicAuth: &management.EventStreamWebhookBasicAuth{
                            Method: management.EventStreamWebhookBasicAuthMethodEnum(
                                "basic",
                            ),
                            Username: "username",
                        },
                    },
                },
            },
        },
    }
client.EventStreams.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.EventStreamsCreateRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EventStreams.Get(ID) -> *management.GetEventStreamResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.EventStreams.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EventStreams.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.EventStreams.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EventStreams.Update(ID, request) -> *management.UpdateEventStreamResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateEventStreamRequestContent{}
client.EventStreams.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name of the event stream.
    
</dd>
</dl>

<dl>
<dd>

**subscriptions:** `[]*management.EventStreamSubscription` — List of event types subscribed to in this stream.
    
</dd>
</dl>

<dl>
<dd>

**destination:** `*management.EventStreamDestinationPatch` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*management.EventStreamStatusEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EventStreams.Test(ID, request) -> *management.CreateEventStreamTestEventResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateEventStreamTestEventRequestContent{
        EventType: management.EventStreamTestEventTypeEnumUserCreated,
    }
client.EventStreams.Test(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*management.EventStreamTestEventTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `*management.TestEventDataContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Flows
<details><summary><code>client.Flows.List() -> *management.ListFlowsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.FlowsListRequest{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Synchronous: management.Bool(
            true,
        ),
    }
client.Flows.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**hydrate:** `*string` — hydration param
    
</dd>
</dl>

<dl>
<dd>

**synchronous:** `*bool` — flag to filter by sync/async flows
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Create(request) -> *management.CreateFlowResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateFlowRequestContent{
        Name: "name",
    }
client.Flows.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**actions:** `[]*management.FlowAction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Get(ID) -> *management.GetFlowResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetFlowRequestParameters{}
client.Flows.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Flow identifier
    
</dd>
</dl>

<dl>
<dd>

**hydrate:** `*management.GetFlowRequestParametersHydrateEnum` — hydration param
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Flows.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Flow id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Update(ID, request) -> *management.UpdateFlowResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateFlowRequestContent{}
client.Flows.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Flow identifier
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**actions:** `[]*management.FlowAction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Forms
<details><summary><code>client.Forms.List() -> *management.ListFormsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListFormsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Forms.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**hydrate:** `*management.FormsRequestParametersHydrateEnum` — Query parameter to hydrate the response with additional data
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Forms.Create(request) -> *management.CreateFormResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateFormRequestContent{
        Name: "name",
    }
client.Forms.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**messages:** `*management.FormMessages` 
    
</dd>
</dl>

<dl>
<dd>

**languages:** `*management.FormLanguages` 
    
</dd>
</dl>

<dl>
<dd>

**translations:** `*management.FormTranslations` 
    
</dd>
</dl>

<dl>
<dd>

**nodes:** `*management.FormNodeList` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*management.FormStartNode` 
    
</dd>
</dl>

<dl>
<dd>

**ending:** `*management.FormEndingNode` 
    
</dd>
</dl>

<dl>
<dd>

**style:** `*management.FormStyle` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Forms.Get(ID) -> *management.GetFormResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetFormRequestParameters{}
client.Forms.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the form to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**hydrate:** `*management.FormsRequestParametersHydrateEnum` — Query parameter to hydrate the response with additional data
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Forms.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Forms.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the form to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Forms.Update(ID, request) -> *management.UpdateFormResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateFormRequestContent{}
client.Forms.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the form to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**messages:** `*management.FormMessagesNullable` 
    
</dd>
</dl>

<dl>
<dd>

**languages:** `*management.FormLanguagesNullable` 
    
</dd>
</dl>

<dl>
<dd>

**translations:** `*management.FormTranslationsNullable` 
    
</dd>
</dl>

<dl>
<dd>

**nodes:** `*management.FormNodeListNullable` 
    
</dd>
</dl>

<dl>
<dd>

**start:** `*management.FormStartNodeNullable` 
    
</dd>
</dl>

<dl>
<dd>

**ending:** `*management.FormEndingNodeNullable` 
    
</dd>
</dl>

<dl>
<dd>

**style:** `*management.FormStyleNullable` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## UserGrants
<details><summary><code>client.UserGrants.List() -> *management.ListUserGrantsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the <a href="https://auth0.com/docs/api-auth/which-oauth-flow-to-use">grants</a> associated with your account. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserGrantsRequestParameters{
        PerPage: management.Int(
            1,
        ),
        Page: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        UserID: management.String(
            "user_id",
        ),
        ClientID: management.String(
            "client_id",
        ),
        Audience: management.String(
            "audience",
        ),
    }
client.UserGrants.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — user_id of the grants to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — client_id of the grants to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**audience:** `*string` — audience of the grants to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGrants.DeleteByUserID() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a grant associated with your account. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteUserGrantByUserIDRequestParameters{
        UserID: "user_id",
    }
client.UserGrants.DeleteByUserID(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — user_id of the grant to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserGrants.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a grant associated with your account. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserGrants.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the grant to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hooks
<details><summary><code>client.Hooks.List() -> *management.ListHooksOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all <a href="https://auth0.com/docs/hooks">hooks</a>. Accepts a list of fields to include or exclude in the result.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListHooksRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Enabled: management.Bool(
            true,
        ),
        Fields: management.String(
            "fields",
        ),
        TriggerID: management.HookTriggerIDEnumCredentialsExchange.Ptr(),
    }
client.Hooks.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Optional filter on whether a hook is enabled (true) or disabled (false).
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**triggerID:** `*management.HookTriggerIDEnum` — Retrieves hooks that match the trigger
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.Create(request) -> *management.CreateHookResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new hook.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateHookRequestContent{
        Name: "name",
        Script: "script",
        TriggerID: management.HookTriggerIDEnumCredentialsExchange,
    }
client.Hooks.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Name of this hook.
    
</dd>
</dl>

<dl>
<dd>

**script:** `string` — Code to be executed when this hook runs.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether this hook will be executed (true) or ignored (false).
    
</dd>
</dl>

<dl>
<dd>

**dependencies:** `*management.HookDependencies` 
    
</dd>
</dl>

<dl>
<dd>

**triggerID:** `*management.HookTriggerIDEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.Get(ID) -> *management.GetHookResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve <a href="https://auth0.com/docs/hooks">a hook</a> by its ID. Accepts a list of fields to include in the result.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetHookRequestParameters{
        Fields: management.String(
            "fields",
        ),
    }
client.Hooks.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the hook to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a hook.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hooks.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the hook to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.Update(ID, request) -> *management.UpdateHookResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing hook.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateHookRequestContent{}
client.Hooks.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the hook to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name of this hook.
    
</dd>
</dl>

<dl>
<dd>

**script:** `*string` — Code to be executed when this hook runs.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether this hook will be executed (true) or ignored (false).
    
</dd>
</dl>

<dl>
<dd>

**dependencies:** `*management.HookDependencies` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Jobs
<details><summary><code>client.Jobs.Get(ID) -> *management.GetJobResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a job. Useful to check its status.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jobs.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the job.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LogStreams
<details><summary><code>client.LogStreams.List() -> []*management.LogStreamResponseSchema</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details on <a href="https://auth0.com/docs/logs/streams">log streams</a>.
<h5>Sample Response</h5><pre><code>[{
	"id": "string",
	"name": "string",
	"type": "eventbridge",
	"status": "active|paused|suspended",
	"sink": {
		"awsAccountId": "string",
		"awsRegion": "string",
		"awsPartnerEventSource": "string"
	}
}, {
	"id": "string",
	"name": "string",
	"type": "http",
	"status": "active|paused|suspended",
	"sink": {
		"httpContentFormat": "JSONLINES|JSONARRAY",
		"httpContentType": "string",
		"httpEndpoint": "string",
		"httpAuthorization": "string"
	}
},
{
	"id": "string",
	"name": "string",
	"type": "eventgrid",
	"status": "active|paused|suspended",
	"sink": {
		"azureSubscriptionId": "string",
		"azureResourceGroup": "string",
		"azureRegion": "string",
		"azurePartnerTopic": "string"
	}
},
{
	"id": "string",
	"name": "string",
	"type": "splunk",
	"status": "active|paused|suspended",
	"sink": {
		"splunkDomain": "string",
		"splunkToken": "string",
		"splunkPort": "string",
		"splunkSecure": "boolean"
	}
},
{
	"id": "string",
	"name": "string",
	"type": "sumo",
	"status": "active|paused|suspended",
	"sink": {
		"sumoSourceAddress": "string",
	}
},
{
	"id": "string",
	"name": "string",
	"type": "datadog",
	"status": "active|paused|suspended",
	"sink": {
		"datadogRegion": "string",
		"datadogApiKey": "string"
	}
}]</code></pre>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LogStreams.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LogStreams.Create(request) -> *management.CreateLogStreamResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a log stream.
<h5>Log Stream Types</h5> The <code>type</code> of log stream being created determines the properties required in the <code>sink</code> payload.
<h5>HTTP Stream</h5> For an <code>http</code> Stream, the <code>sink</code> properties are listed in the payload below
Request: <pre><code>{
	"name": "string",
	"type": "http",
	"sink": {
		"httpEndpoint": "string",
		"httpContentType": "string",
		"httpContentFormat": "JSONLINES|JSONARRAY",
		"httpAuthorization": "string"
	}
}</code></pre>
Response: <pre><code>{
	"id": "string",
	"name": "string",
	"type": "http",
	"status": "active",
	"sink": {
		"httpEndpoint": "string",
		"httpContentType": "string",
		"httpContentFormat": "JSONLINES|JSONARRAY",
		"httpAuthorization": "string"
	}
}</code></pre>
<h5>Amazon EventBridge Stream</h5> For an <code>eventbridge</code> Stream, the <code>sink</code> properties are listed in the payload below
Request: <pre><code>{
	"name": "string",
	"type": "eventbridge",
	"sink": {
		"awsRegion": "string",
		"awsAccountId": "string"
	}
}</code></pre>
The response will include an additional field <code>awsPartnerEventSource</code> in the <code>sink</code>: <pre><code>{
	"id": "string",
	"name": "string",
	"type": "eventbridge",
	"status": "active",
	"sink": {
		"awsAccountId": "string",
		"awsRegion": "string",
		"awsPartnerEventSource": "string"
	}
}</code></pre>
<h5>Azure Event Grid Stream</h5> For an <code>Azure Event Grid</code> Stream, the <code>sink</code> properties are listed in the payload below
Request: <pre><code>{
	"name": "string",
	"type": "eventgrid",
	"sink": {
		"azureSubscriptionId": "string",
		"azureResourceGroup": "string",
		"azureRegion": "string"
	}
}</code></pre>
Response: <pre><code>{
	"id": "string",
	"name": "string",
	"type": "http",
	"status": "active",
	"sink": {
		"azureSubscriptionId": "string",
		"azureResourceGroup": "string",
		"azureRegion": "string",
		"azurePartnerTopic": "string"
	}
}</code></pre>
<h5>Datadog Stream</h5> For a <code>Datadog</code> Stream, the <code>sink</code> properties are listed in the payload below
Request: <pre><code>{
	"name": "string",
	"type": "datadog",
	"sink": {
		"datadogRegion": "string",
		"datadogApiKey": "string"
	}
}</code></pre>
Response: <pre><code>{
	"id": "string",
	"name": "string",
	"type": "datadog",
	"status": "active",
	"sink": {
		"datadogRegion": "string",
		"datadogApiKey": "string"
	}
}</code></pre>
<h5>Splunk Stream</h5> For a <code>Splunk</code> Stream, the <code>sink</code> properties are listed in the payload below
Request: <pre><code>{
	"name": "string",
	"type": "splunk",
	"sink": {
		"splunkDomain": "string",
		"splunkToken": "string",
		"splunkPort": "string",
		"splunkSecure": "boolean"
	}
}</code></pre>
Response: <pre><code>{
	"id": "string",
	"name": "string",
	"type": "splunk",
	"status": "active",
	"sink": {
		"splunkDomain": "string",
		"splunkToken": "string",
		"splunkPort": "string",
		"splunkSecure": "boolean"
	}
}</code></pre>
<h5>Sumo Logic Stream</h5> For a <code>Sumo Logic</code> Stream, the <code>sink</code> properties are listed in the payload below
Request: <pre><code>{
	"name": "string",
	"type": "sumo",
	"sink": {
		"sumoSourceAddress": "string",
	}
}</code></pre>
Response: <pre><code>{
	"id": "string",
	"name": "string",
	"type": "sumo",
	"status": "active",
	"sink": {
		"sumoSourceAddress": "string",
	}
}</code></pre>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateLogStreamRequestContent{
        CreateLogStreamHTTPRequestBody: &management.CreateLogStreamHTTPRequestBody{
            Type: management.LogStreamHTTPEnum(
                "http",
            ),
            Sink: &management.LogStreamHTTPSink{
                HTTPEndpoint: "httpEndpoint",
            },
        },
    }
client.LogStreams.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateLogStreamRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LogStreams.Get(ID) -> *management.GetLogStreamResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a log stream configuration and status.
<h5>Sample responses</h5><h5>Amazon EventBridge Log Stream</h5><pre><code>{
	"id": "string",
	"name": "string",
	"type": "eventbridge",
	"status": "active|paused|suspended",
	"sink": {
		"awsAccountId": "string",
		"awsRegion": "string",
		"awsPartnerEventSource": "string"
	}
}</code></pre> <h5>HTTP Log Stream</h5><pre><code>{
	"id": "string",
	"name": "string",
	"type": "http",
	"status": "active|paused|suspended",
	"sink": {
		"httpContentFormat": "JSONLINES|JSONARRAY",
		"httpContentType": "string",
		"httpEndpoint": "string",
		"httpAuthorization": "string"
	}
}</code></pre> <h5>Datadog Log Stream</h5><pre><code>{
	"id": "string",
	"name": "string",
	"type": "datadog",
	"status": "active|paused|suspended",
	"sink": {
		"datadogRegion": "string",
		"datadogApiKey": "string"
	}

}</code></pre><h5>Mixpanel</h5>
	
	Request: <pre><code>{
	  "name": "string",
	  "type": "mixpanel",
	  "sink": {
		"mixpanelRegion": "string", // "us" | "eu",
		"mixpanelProjectId": "string",
		"mixpanelServiceAccountUsername": "string",
		"mixpanelServiceAccountPassword": "string"
	  }
	} </code></pre>
	
	
	Response: <pre><code>{
		"id": "string",
		"name": "string",
		"type": "mixpanel",
		"status": "active",
		"sink": {
		  "mixpanelRegion": "string", // "us" | "eu",
		  "mixpanelProjectId": "string",
		  "mixpanelServiceAccountUsername": "string",
		  "mixpanelServiceAccountPassword": "string" // the following is redacted on return
		}
	  } </code></pre>

	<h5>Segment</h5>

	Request: <pre><code> {
	  "name": "string",
	  "type": "segment",
	  "sink": {
		"segmentWriteKey": "string"
	  }
	}</code></pre>
	
	Response: <pre><code>{
	  "id": "string",
	  "name": "string",
	  "type": "segment",
	  "status": "active",
	  "sink": {
		"segmentWriteKey": "string"
	  }
	} </code></pre>
	
<h5>Splunk Log Stream</h5><pre><code>{
	"id": "string",
	"name": "string",
	"type": "splunk",
	"status": "active|paused|suspended",
	"sink": {
		"splunkDomain": "string",
		"splunkToken": "string",
		"splunkPort": "string",
		"splunkSecure": "boolean"
	}
}</code></pre> <h5>Sumo Logic Log Stream</h5><pre><code>{
	"id": "string",
	"name": "string",
	"type": "sumo",
	"status": "active|paused|suspended",
	"sink": {
		"sumoSourceAddress": "string",
	}
}</code></pre> <h5>Status</h5> The <code>status</code> of a log stream maybe any of the following:
1. <code>active</code> - Stream is currently enabled.
2. <code>paused</code> - Stream is currently user disabled and will not attempt log delivery.
3. <code>suspended</code> - Stream is currently disabled because of errors and will not attempt log delivery.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LogStreams.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the log stream to get
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LogStreams.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a log stream.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.LogStreams.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the log stream to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LogStreams.Update(ID, request) -> *management.UpdateLogStreamResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a log stream.
<h4>Examples of how to use the PATCH endpoint.</h4> The following fields may be updated in a PATCH operation: <ul><li>name</li><li>status</li><li>sink</li></ul> Note: For log streams of type <code>eventbridge</code> and <code>eventgrid</code>, updating the <code>sink</code> is not permitted.
<h5>Update the status of a log stream</h5><pre><code>{
	"status": "active|paused"
}</code></pre>
<h5>Update the name of a log stream</h5><pre><code>{
	"name": "string"
}</code></pre>
<h5>Update the sink properties of a stream of type <code>http</code></h5><pre><code>{
  "sink": {
    "httpEndpoint": "string",
    "httpContentType": "string",
    "httpContentFormat": "JSONARRAY|JSONLINES",
    "httpAuthorization": "string"
  }
}</code></pre>
<h5>Update the sink properties of a stream of type <code>datadog</code></h5><pre><code>{
  "sink": {
		"datadogRegion": "string",
		"datadogApiKey": "string"
  }
}</code></pre>
<h5>Update the sink properties of a stream of type <code>splunk</code></h5><pre><code>{
  "sink": {
    "splunkDomain": "string",
    "splunkToken": "string",
    "splunkPort": "string",
    "splunkSecure": "boolean"
  }
}</code></pre>
<h5>Update the sink properties of a stream of type <code>sumo</code></h5><pre><code>{
  "sink": {
    "sumoSourceAddress": "string"
  }
}</code></pre> 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateLogStreamRequestContent{}
client.LogStreams.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the log stream to get
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — log stream name
    
</dd>
</dl>

<dl>
<dd>

**status:** `*management.LogStreamStatusEnum` 
    
</dd>
</dl>

<dl>
<dd>

**isPriority:** `*bool` — True for priority log streams, false for non-priority
    
</dd>
</dl>

<dl>
<dd>

**filters:** `[]*management.LogStreamFilter` — Only logs events matching these filters will be delivered by the stream. If omitted or empty, all events will be delivered.
    
</dd>
</dl>

<dl>
<dd>

**piiConfig:** `*management.LogStreamPiiConfig` 
    
</dd>
</dl>

<dl>
<dd>

**sink:** `*management.LogStreamSinkPatch` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Logs
<details><summary><code>client.Logs.List() -> *management.ListLogOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve log entries that match the specified search criteria (or all log entries if no criteria specified).

Set custom search criteria using the <code>q</code> parameter, or search from a specific log ID (<i>"search from checkpoint"</i>).

For more information on all possible event types, their respective acronyms, and descriptions, see <a href="https://auth0.com/docs/logs/log-event-type-codes">Log Event Type Codes</a>.

<h5>To set custom search criteria, use the following parameters:</h5>

<ul>
    <li><b>q:</b> Search Criteria using <a href="https://auth0.com/docs/logs/log-search-query-syntax">Query String Syntax</a></li>
    <li><b>page:</b> Page index of the results to return. First page is 0.</li>
    <li><b>per_page:</b> Number of results per page.</li>
    <li><b>sort:</b> Field to use for sorting appended with `:1` for ascending and `:-1` for descending. e.g. `date:-1`</li>
    <li><b>fields:</b> Comma-separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields.</li>
    <li><b>include_fields:</b> Whether specified fields are to be included (true) or excluded (false).</li>
    <li><b>include_totals:</b> Return results inside an object that contains the total result count (true) or as a direct array of results (false, default). <b>Deprecated:</b> this field is deprecated and should be removed from use. See <a href="https://auth0.com/docs/product-lifecycle/deprecations-and-migrations/migrate-to-tenant-log-search-v3#pagination">Search Engine V3 Breaking Changes</a></li>
</ul>

For more information on the list of fields that can be used in <code>fields</code> and <code>sort</code>, see <a href="https://auth0.com/docs/logs/log-search-query-syntax#searchable-fields">Searchable Fields</a>.

Auth0 <a href="https://auth0.com/docs/logs/retrieve-log-events-using-mgmt-api#limitations">limits the number of logs</a> you can return by search criteria to 100 logs per request. Furthermore, you may paginate only through 1,000 search results. If you exceed this threshold, please redefine your search or use the <a href="https://auth0.com/docs/logs/retrieve-log-events-using-mgmt-api#retrieve-logs-by-checkpoint">get logs by checkpoint method</a>.

<h5>To search from a checkpoint log ID, use the following parameters:</h5>
<ul>
    <li><b>from:</b> Log Event ID from which to start retrieving logs. You can limit the number of logs returned using the <code>take</code> parameter. If you use <code>from</code> at the same time as <code>q</code>, <code>from</code> takes precedence and <code>q</code> is ignored.</li>
    <li><b>take:</b> Number of entries to retrieve when using the <code>from</code> parameter.</li>
</ul>

<strong>Important:</strong> When fetching logs from a checkpoint log ID, any parameter other than <code>from</code> and <code>take</code> will be ignored, and date ordering is not guaranteed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListLogsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        Sort: management.String(
            "sort",
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Search: management.String(
            "search",
        ),
    }
client.Logs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` —  Number of results per page. Paging is disabled if parameter not sent. Default: <code>50</code>. Max value: <code>100</code>
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Field to use for sorting appended with <code>:1</code>  for ascending and <code>:-1</code> for descending. e.g. <code>date:-1</code>
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for <code>include_fields</code>) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (<code>true</code>) or excluded (<code>false</code>)
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results as an array when false (default). Return results inside an object that also contains a total result count when true.
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` 

Retrieves logs that match the specified search criteria. This parameter can be combined with all the others in the /api/logs endpoint but is specified separately for clarity.
If no fields are provided a case insensitive 'starts with' search is performed on all of the following fields: client_name, connection, user_name. Otherwise, you can specify multiple fields and specify the search using the %field%:%search%, for example: application:node user:"John@contoso.com".
Values specified without quotes are matched using a case insensitive 'starts with' search. If quotes are used a case insensitve exact search is used. If multiple fields are used, the AND operator is used to join the clauses.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Logs.Get(ID) -> *management.GetLogResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve an individual log event.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Logs.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — log_id of the log to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## NetworkACLs
<details><summary><code>client.NetworkACLs.List() -> *management.ListNetworkACLsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all access control list entries for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListNetworkACLsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.NetworkACLs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Use this field to request a specific page of the list results.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — The amount of results per page.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkACLs.Create(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new access control list for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateNetworkACLRequestContent{
        Description: "description",
        Active: true,
        Priority: 1.1,
        Rule: &management.NetworkACLRule{
            Action: &management.NetworkACLAction{},
            Scope: management.NetworkACLRuleScopeEnumManagement,
        },
    }
client.NetworkACLs.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**active:** `bool` — Indicates whether or not this access control list is actively being used
    
</dd>
</dl>

<dl>
<dd>

**priority:** `float64` — Indicates the order in which the ACL will be evaluated relative to other ACL rules.
    
</dd>
</dl>

<dl>
<dd>

**rule:** `*management.NetworkACLRule` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkACLs.Get(ID) -> *management.GetNetworkACLsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific access control list entry for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkACLs.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the access control list to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkACLs.Set(ID, request) -> *management.SetNetworkACLsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update existing access control list for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetNetworkACLRequestContent{
        Description: "description",
        Active: true,
        Priority: 1.1,
        Rule: &management.NetworkACLRule{
            Action: &management.NetworkACLAction{},
            Scope: management.NetworkACLRuleScopeEnumManagement,
        },
    }
client.NetworkACLs.Set(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the ACL to update.
    
</dd>
</dl>

<dl>
<dd>

**description:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**active:** `bool` — Indicates whether or not this access control list is actively being used
    
</dd>
</dl>

<dl>
<dd>

**priority:** `float64` — Indicates the order in which the ACL will be evaluated relative to other ACL rules.
    
</dd>
</dl>

<dl>
<dd>

**rule:** `*management.NetworkACLRule` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkACLs.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete existing access control list for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.NetworkACLs.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the ACL to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.NetworkACLs.Update(ID, request) -> *management.UpdateNetworkACLResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update existing access control list for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateNetworkACLRequestContent{}
client.NetworkACLs.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the ACL to update.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**active:** `*bool` — Indicates whether or not this access control list is actively being used
    
</dd>
</dl>

<dl>
<dd>

**priority:** `*float64` — Indicates the order in which the ACL will be evaluated relative to other ACL rules.
    
</dd>
</dl>

<dl>
<dd>

**rule:** `*management.NetworkACLRule` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations
<details><summary><code>client.Organizations.List() -> *management.ListOrganizationsPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed list of all Organizations available in your tenant. For more information, see Auth0 Organizations.

This endpoint supports two types of pagination:
<ul>
<li>Offset pagination</li>
<li>Checkpoint pagination</li>
</ul>

Checkpoint pagination must be used if you need to retrieve more than 1000 organizations.

<h2>Checkpoint Pagination</h2>

To search by checkpoint, use the following parameters:
<ul>
<li><code>from</code>: Optional id from which to start selection.</li>
<li><code>take</code>: The total number of entries to retrieve when using the <code>from</code> parameter. Defaults to 50.</li>
</ul>

<b>Note</b>: The first time you call this endpoint using checkpoint pagination, omit the <code>from</code> parameter. If there are more results, a <code>next</code> value is included in the response. You can use this for subsequent API calls. When <code>next</code> is no longer included in the response, no pages are remaining.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationsRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
        Sort: management.String(
            "sort",
        ),
    }
client.Organizations.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Field to sort by. Use <code>field:order</code> where order is <code>1</code> for ascending and <code>-1</code> for descending. e.g. <code>created_at:1</code>. We currently support sorting by the following fields: <code>name</code>, <code>display_name</code> and <code>created_at</code>.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Create(request) -> *management.CreateOrganizationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Organization within your tenant.  To learn more about Organization settings, behavior, and configuration options, review <a href="https://auth0.com/docs/manage-users/organizations/create-first-organization">Create Your First Organization</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateOrganizationRequestContent{
        Name: "name",
    }
client.Organizations.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of this organization.
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` — Friendly name of this organization.
    
</dd>
</dl>

<dl>
<dd>

**branding:** `*management.OrganizationBranding` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `*management.OrganizationMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**enabledConnections:** `[]*management.ConnectionForOrganization` — Connections that will be enabled for this organization. See POST enabled_connections endpoint for the object format. (Max of 10 connections allowed)
    
</dd>
</dl>

<dl>
<dd>

**tokenQuota:** `*management.CreateTokenQuota` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.GetByName(Name) -> *management.GetOrganizationByNameResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a single Organization specified by name.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.GetByName(
        context.TODO(),
        "name",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — name of the organization to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Get(ID) -> *management.GetOrganizationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a single Organization specified by ID. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the organization to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove an Organization from your tenant.  This action cannot be undone. 

<b>Note</b>: Members are automatically disassociated from an Organization when it is deleted. However, this action does <b>not</b> delete these users from your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Update(ID, request) -> *management.UpdateOrganizationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the details of a specific <a href="https://auth0.com/docs/manage-users/organizations/configure-organizations/create-organizations">Organization</a>, such as name and display name, branding options, and metadata.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateOrganizationRequestContent{}
client.Organizations.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the organization to update.
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` — Friendly name of this organization.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of this organization.
    
</dd>
</dl>

<dl>
<dd>

**branding:** `*management.OrganizationBranding` 
    
</dd>
</dl>

<dl>
<dd>

**metadata:** `*management.OrganizationMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**tokenQuota:** `*management.UpdateTokenQuota` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompts
<details><summary><code>client.Prompts.GetSettings() -> *management.GetSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the Universal Login configuration of your tenant. This includes the <a href="https://auth0.com/docs/authenticate/login/auth0-universal-login/identifier-first">Identifier First Authentication</a> and <a href="https://auth0.com/docs/secure/multi-factor-authentication/fido-authentication-with-webauthn/configure-webauthn-device-biometrics-for-mfa">WebAuthn with Device Biometrics for MFA</a> features.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Prompts.GetSettings(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.UpdateSettings(request) -> *management.UpdateSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the Universal Login configuration of your tenant. This includes the <a href="https://auth0.com/docs/authenticate/login/auth0-universal-login/identifier-first">Identifier First Authentication</a> and <a href="https://auth0.com/docs/secure/multi-factor-authentication/fido-authentication-with-webauthn/configure-webauthn-device-biometrics-for-mfa">WebAuthn with Device Biometrics for MFA</a> features.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateSettingsRequestContent{}
client.Prompts.UpdateSettings(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**universalLoginExperience:** `*management.UniversalLoginExperienceEnum` 
    
</dd>
</dl>

<dl>
<dd>

**identifierFirst:** `*bool` — Whether identifier first is enabled or not
    
</dd>
</dl>

<dl>
<dd>

**webauthnPlatformFirstFactor:** `*bool` — Use WebAuthn with Device Biometrics as the first authentication factor
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## RefreshTokens
<details><summary><code>client.RefreshTokens.Get(ID) -> *management.GetRefreshTokenResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve refresh token information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.RefreshTokens.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID refresh token to retrieve
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RefreshTokens.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a refresh token by its ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.RefreshTokens.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the refresh token to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ResourceServers
<details><summary><code>client.ResourceServers.List() -> *management.ListResourceServerOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of all APIs associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListResourceServerRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.ResourceServers.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifiers:** `*string` — An optional filter on the resource server identifier. Must be URL encoded and may be specified multiple times (max 10).<br /><b>e.g.</b> <i>../resource-servers?identifiers=id1&identifiers=id2</i>
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResourceServers.Create(request) -> *management.CreateResourceServerResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new API associated with your tenant. Note that all new APIs must be registered with Auth0. For more information, read <a href="https://www.auth0.com/docs/get-started/apis"> APIs</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateResourceServerRequestContent{
        Identifier: "identifier",
    }
client.ResourceServers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*string` — Friendly name for this resource server. Can not contain `<` or `>` characters.
    
</dd>
</dl>

<dl>
<dd>

**identifier:** `string` — Unique identifier for the API used as the audience parameter on authorization calls. Can not be changed once set.
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]*management.ResourceServerScope` — List of permissions (scopes) that this API uses.
    
</dd>
</dl>

<dl>
<dd>

**signingAlg:** `*management.SigningAlgorithmEnum` 
    
</dd>
</dl>

<dl>
<dd>

**signingSecret:** `*string` — Secret used to sign tokens when using symmetric algorithms (HS256).
    
</dd>
</dl>

<dl>
<dd>

**allowOfflineAccess:** `*bool` — Whether refresh tokens can be issued for this API (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**tokenLifetime:** `*int` — Expiration value (in seconds) for access tokens issued for this API from the token endpoint.
    
</dd>
</dl>

<dl>
<dd>

**tokenDialect:** `*management.ResourceServerTokenDialectSchemaEnum` 
    
</dd>
</dl>

<dl>
<dd>

**skipConsentForVerifiableFirstPartyClients:** `*bool` — Whether to skip user consent for applications flagged as first party (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**enforcePolicies:** `*bool` — Whether to enforce authorization policies (true) or to ignore them (false).
    
</dd>
</dl>

<dl>
<dd>

**tokenEncryption:** `*management.ResourceServerTokenEncryption` 
    
</dd>
</dl>

<dl>
<dd>

**consentPolicy:** `*management.ResourceServerConsentPolicyEnum` 
    
</dd>
</dl>

<dl>
<dd>

**authorizationDetails:** `[]any` 
    
</dd>
</dl>

<dl>
<dd>

**proofOfPossession:** `*management.ResourceServerProofOfPossession` 
    
</dd>
</dl>

<dl>
<dd>

**subjectTypeAuthorization:** `*management.ResourceServerSubjectTypeAuthorization` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResourceServers.Get(ID) -> *management.GetResourceServerResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve <a href="https://auth0.com/docs/apis">API</a> details with the given ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetResourceServerRequestParameters{
        IncludeFields: management.Bool(
            true,
        ),
    }
client.ResourceServers.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID or audience of the resource server to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResourceServers.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete an existing API by ID. For more information, read <a href="https://www.auth0.com/docs/get-started/apis/api-settings">API Settings</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ResourceServers.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID or the audience of the resource server to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ResourceServers.Update(ID, request) -> *management.UpdateResourceServerResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Change an existing API setting by resource server ID. For more information, read <a href="https://www.auth0.com/docs/get-started/apis/api-settings">API Settings</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateResourceServerRequestContent{}
client.ResourceServers.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID or audience of the resource server to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Friendly name for this resource server. Can not contain `<` or `>` characters.
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]*management.ResourceServerScope` — List of permissions (scopes) that this API uses.
    
</dd>
</dl>

<dl>
<dd>

**signingAlg:** `*management.SigningAlgorithmEnum` 
    
</dd>
</dl>

<dl>
<dd>

**signingSecret:** `*string` — Secret used to sign tokens when using symmetric algorithms (HS256).
    
</dd>
</dl>

<dl>
<dd>

**skipConsentForVerifiableFirstPartyClients:** `*bool` — Whether to skip user consent for applications flagged as first party (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**allowOfflineAccess:** `*bool` — Whether refresh tokens can be issued for this API (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**tokenLifetime:** `*int` — Expiration value (in seconds) for access tokens issued for this API from the token endpoint.
    
</dd>
</dl>

<dl>
<dd>

**tokenDialect:** `*management.ResourceServerTokenDialectSchemaEnum` 
    
</dd>
</dl>

<dl>
<dd>

**enforcePolicies:** `*bool` — Whether authorization policies are enforced (true) or not enforced (false).
    
</dd>
</dl>

<dl>
<dd>

**tokenEncryption:** `*management.ResourceServerTokenEncryption` 
    
</dd>
</dl>

<dl>
<dd>

**consentPolicy:** `*management.ResourceServerConsentPolicyEnum` 
    
</dd>
</dl>

<dl>
<dd>

**authorizationDetails:** `[]any` 
    
</dd>
</dl>

<dl>
<dd>

**proofOfPossession:** `*management.ResourceServerProofOfPossession` 
    
</dd>
</dl>

<dl>
<dd>

**subjectTypeAuthorization:** `*management.ResourceServerSubjectTypeAuthorization` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Roles
<details><summary><code>client.Roles.List() -> *management.ListRolesOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed list of user roles created in your tenant.

<b>Note</b>: The returned list does not include standard roles available for tenant members, such as Admin or Support Access.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListRolesRequestParameters{
        PerPage: management.Int(
            1,
        ),
        Page: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        NameFilter: management.String(
            "name_filter",
        ),
    }
client.Roles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**nameFilter:** `*string` — Optional filter on name (case-insensitive).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Create(request) -> *management.CreateRoleResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a user role for <a href="https://auth0.com/docs/manage-users/access-control/rbac">Role-Based Access Control</a>.

<b>Note</b>: New roles are not associated with any permissions by default. To assign existing permissions to your role, review Associate Permissions with a Role. To create new permissions, review Add API Permissions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateRoleRequestContent{
        Name: "name",
    }
client.Roles.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Name of the role.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Description of the role.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Get(ID) -> *management.GetRoleResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a specific <a href="https://auth0.com/docs/manage-users/access-control/rbac">user role</a> specified by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Roles.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a specific <a href="https://auth0.com/docs/manage-users/access-control/rbac">user role</a> from your tenant. Once deleted, it is removed from any user who was previously assigned that role. This action cannot be undone.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Roles.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Update(ID, request) -> *management.UpdateRoleResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify the details of a specific <a href="https://auth0.com/docs/manage-users/access-control/rbac">user role</a> specified by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateRoleRequestContent{}
client.Roles.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name of this role.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — Description of this role.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Rules
<details><summary><code>client.Rules.List() -> *management.ListRulesOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a filtered list of <a href="https://auth0.com/docs/rules">rules</a>. Accepts a list of fields to include or exclude.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListRulesRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Enabled: management.Bool(
            true,
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Rules.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Optional filter on whether a rule is enabled (true) or disabled (false).
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Rules.Create(request) -> *management.CreateRuleResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a <a href="https://auth0.com/docs/rules#create-a-new-rule-using-the-management-api">new rule</a>.

Note: Changing a rule's stage of execution from the default <code>login_success</code> can change the rule's function signature to have user omitted.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateRuleRequestContent{
        Name: "name",
        Script: "script",
    }
client.Rules.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Name of this rule.
    
</dd>
</dl>

<dl>
<dd>

**script:** `string` — Code to be executed when this rule runs.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*float64` — Order that this rule should execute in relative to other rules. Lower-valued rules execute first.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the rule is enabled (true), or disabled (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Rules.Get(ID) -> *management.GetRuleResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve <a href="https://auth0.com/docs/rules">rule</a> details. Accepts a list of fields to include or exclude in the result.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetRuleRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Rules.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the rule to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Rules.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a rule.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Rules.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the rule to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Rules.Update(ID, request) -> *management.UpdateRuleResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an existing rule.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateRuleRequestContent{}
client.Rules.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the rule to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**script:** `*string` — Code to be executed when this rule runs.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name of this rule.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*float64` — Order that this rule should execute in relative to other rules. Lower-valued rules execute first.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the rule is enabled (true), or disabled (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## RulesConfigs
<details><summary><code>client.RulesConfigs.List() -> []*management.RulesConfig</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve rules config variable keys.

    Note: For security, config variable values cannot be retrieved outside rule execution.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.RulesConfigs.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RulesConfigs.Set(Key, request) -> *management.SetRulesConfigResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Sets a rules config variable.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetRulesConfigRequestContent{
        Value: "value",
    }
client.RulesConfigs.Set(
        context.TODO(),
        "key",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Key of the rules config variable to set (max length: 127 characters).
    
</dd>
</dl>

<dl>
<dd>

**value:** `string` — Value for a rules config variable.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RulesConfigs.Delete(Key) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a rules config variable identified by its key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.RulesConfigs.Delete(
        context.TODO(),
        "key",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**key:** `string` — Key of the rules config variable to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SelfServiceProfiles
<details><summary><code>client.SelfServiceProfiles.List() -> *management.ListSelfServiceProfilesPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves self-service profiles.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListSelfServiceProfilesRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.SelfServiceProfiles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SelfServiceProfiles.Create(request) -> *management.CreateSelfServiceProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a self-service profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateSelfServiceProfileRequestContent{
        Name: "name",
    }
client.SelfServiceProfiles.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the self-service Profile.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*string` — The description of the self-service Profile.
    
</dd>
</dl>

<dl>
<dd>

**branding:** `*management.SelfServiceProfileBrandingProperties` 
    
</dd>
</dl>

<dl>
<dd>

**allowedStrategies:** `[]*management.SelfServiceProfileAllowedStrategyEnum` — List of IdP strategies that will be shown to users during the Self-Service SSO flow. Possible values: [`oidc`, `samlp`, `waad`, `google-apps`, `adfs`, `okta`, `keycloak-samlp`, `pingfederate`]
    
</dd>
</dl>

<dl>
<dd>

**userAttributes:** `[]*management.SelfServiceProfileUserAttribute` — List of attributes to be mapped that will be shown to the user during the SS-SSO flow.
    
</dd>
</dl>

<dl>
<dd>

**userAttributeProfileID:** `*string` — ID of the user-attribute-profile to associate with this self-service profile.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SelfServiceProfiles.Get(ID) -> *management.GetSelfServiceProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a self-service profile by Id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SelfServiceProfiles.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the self-service profile to retrieve
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SelfServiceProfiles.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a self-service profile by Id.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SelfServiceProfiles.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the self-service profile to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SelfServiceProfiles.Update(ID, request) -> *management.UpdateSelfServiceProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a self-service profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateSelfServiceProfileRequestContent{}
client.SelfServiceProfiles.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the self-service profile to update
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The name of the self-service Profile.
    
</dd>
</dl>

<dl>
<dd>

**description:** `*management.SelfServiceProfileDescription` 
    
</dd>
</dl>

<dl>
<dd>

**branding:** `*management.SelfServiceProfileBranding` 
    
</dd>
</dl>

<dl>
<dd>

**allowedStrategies:** `[]*management.SelfServiceProfileAllowedStrategyEnum` — List of IdP strategies that will be shown to users during the Self-Service SSO flow. Possible values: [`oidc`, `samlp`, `waad`, `google-apps`, `adfs`, `okta`, `keycloak-samlp`, `pingfederate`]
    
</dd>
</dl>

<dl>
<dd>

**userAttributes:** `*management.SelfServiceProfileUserAttributes` 
    
</dd>
</dl>

<dl>
<dd>

**userAttributeProfileID:** `*string` — ID of the user-attribute-profile to associate with this self-service profile.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sessions
<details><summary><code>client.Sessions.Get(ID) -> *management.GetSessionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve session information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sessions.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of session to retrieve
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sessions.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a session by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sessions.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the session to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sessions.Update(ID, request) -> *management.UpdateSessionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update session information.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateSessionRequestContent{}
client.Sessions.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the session to update.
    
</dd>
</dl>

<dl>
<dd>

**sessionMetadata:** `*management.SessionMetadata` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sessions.Revoke(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revokes a session by ID and all associated refresh tokens.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sessions.Revoke(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the session to revoke.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Stats
<details><summary><code>client.Stats.GetActiveUsersCount() -> management.GetActiveUsersCountStatsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of active users that logged in during the last 30 days.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Stats.GetActiveUsersCount(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.GetDaily() -> []*management.DailyStats</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of logins, signups and breached-password detections (subscription required) that occurred each day within a specified date range.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetDailyStatsRequestParameters{
        From: management.String(
            "from",
        ),
        To: management.String(
            "to",
        ),
    }
client.Stats.GetDaily(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional first day of the date range (inclusive) in YYYYMMDD format.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*string` — Optional last day of the date range (inclusive) in YYYYMMDD format.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SupplementalSignals
<details><summary><code>client.SupplementalSignals.Get() -> *management.GetSupplementalSignalsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the supplemental signals configuration for a tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SupplementalSignals.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SupplementalSignals.Patch(request) -> *management.PatchSupplementalSignalsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the supplemental signals configuration for a tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateSupplementalSignalsRequestContent{
        AkamaiEnabled: true,
    }
client.SupplementalSignals.Patch(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**akamaiEnabled:** `bool` — Indicates if incoming Akamai Headers should be processed
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tickets
<details><summary><code>client.Tickets.VerifyEmail(request) -> *management.VerifyEmailTicketResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an email verification ticket for a given user. An email verification ticket is a generated URL that the user can consume to verify their email address.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.VerifyEmailTicketRequestContent{
        UserID: "user_id",
    }
client.Tickets.VerifyEmail(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resultURL:** `*string` — URL the user will be redirected to in the classic Universal Login experience once the ticket is used. Cannot be specified when using client_id or organization_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — user_id of for whom the ticket should be created.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — ID of the client (application). If provided for tenants using the New Universal Login experience, the email template and UI displays application details, and the user is prompted to redirect to the application's <a target='' href='https://auth0.com/docs/authenticate/login/auth0-universal-login/configure-default-login-routes#completing-the-password-reset-flow'>default login route</a> after the ticket is used. client_id is required to use the <a target='' href='https://auth0.com/docs/customize/actions/flows-and-triggers/post-change-password-flow'>Password Reset Post Challenge</a> trigger.
    
</dd>
</dl>

<dl>
<dd>

**organizationID:** `*string` — (Optional) Organization ID – the ID of the Organization. If provided, organization parameters will be made available to the email template and organization branding will be applied to the prompt. In addition, the redirect link in the prompt will include organization_id and organization_name query string parameters.
    
</dd>
</dl>

<dl>
<dd>

**ttlSec:** `*int` — Number of seconds for which the ticket is valid before expiration. If unspecified or set to 0, this value defaults to 432000 seconds (5 days).
    
</dd>
</dl>

<dl>
<dd>

**includeEmailInRedirect:** `*bool` — Whether to include the email address as part of the returnUrl in the reset_email (true), or not (false).
    
</dd>
</dl>

<dl>
<dd>

**identity:** `*management.Identity` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tickets.ChangePassword(request) -> *management.ChangePasswordTicketResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a password change ticket for a given user. A password change ticket is a generated URL that the user can consume to start a reset password flow.

Note: This endpoint does not verify the given user’s identity. If you call this endpoint within your application, you must design your application to verify the user’s identity.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ChangePasswordTicketRequestContent{}
client.Tickets.ChangePassword(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**resultURL:** `*string` — URL the user will be redirected to in the classic Universal Login experience once the ticket is used. Cannot be specified when using client_id or organization_id.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — user_id of for whom the ticket should be created.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — ID of the client (application). If provided for tenants using the New Universal Login experience, the email template and UI displays application details, and the user is prompted to redirect to the application's <a target='' href='https://auth0.com/docs/authenticate/login/auth0-universal-login/configure-default-login-routes#completing-the-password-reset-flow'>default login route</a> after the ticket is used. client_id is required to use the <a target='' href='https://auth0.com/docs/customize/actions/flows-and-triggers/post-change-password-flow'>Password Reset Post Challenge</a> trigger.
    
</dd>
</dl>

<dl>
<dd>

**organizationID:** `*string` — (Optional) Organization ID – the ID of the Organization. If provided, organization parameters will be made available to the email template and organization branding will be applied to the prompt. In addition, the redirect link in the prompt will include organization_id and organization_name query string parameters.
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `*string` — ID of the connection. If provided, allows the user to be specified using email instead of user_id. If you set this value, you must also send the email parameter. You cannot send user_id when specifying a connection_id.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Email address of the user for whom the tickets should be created. Requires the connection_id parameter. Cannot be specified when using user_id.
    
</dd>
</dl>

<dl>
<dd>

**ttlSec:** `*int` — Number of seconds for which the ticket is valid before expiration. If unspecified or set to 0, this value defaults to 432000 seconds (5 days).
    
</dd>
</dl>

<dl>
<dd>

**markEmailAsVerified:** `*bool` — Whether to set the email_verified attribute to true (true) or whether it should not be updated (false).
    
</dd>
</dl>

<dl>
<dd>

**includeEmailInRedirect:** `*bool` — Whether to include the email address as part of the returnUrl in the reset_email (true), or not (false).
    
</dd>
</dl>

<dl>
<dd>

**identity:** `*management.ChangePasswordTicketIdentity` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TokenExchangeProfiles
<details><summary><code>client.TokenExchangeProfiles.List() -> *management.ListTokenExchangeProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of all Token Exchange Profiles available in your tenant.

By using this feature, you agree to the applicable Free Trial terms in <a href="https://www.okta.com/legal/">Okta’s Master Subscription Agreement</a>. It is your responsibility to securely validate the user’s subject_token. See <a href="https://auth0.com/docs/authenticate/custom-token-exchange">User Guide</a> for more details.

This endpoint supports Checkpoint pagination. To search by checkpoint, use the following parameters:
<ul>
<li><code>from</code>: Optional id from which to start selection.</li>
<li><code>take</code>: The total amount of entries to retrieve when using the from parameter. Defaults to 50.</li>
</ul>

<b>Note</b>: The first time you call this endpoint using checkpoint pagination, omit the <code>from</code> parameter. If there are more results, a <code>next</code> value is included in the response. You can use this for subsequent API calls. When <code>next</code> is no longer included in the response, no pages are remaining.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.TokenExchangeProfilesListRequest{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.TokenExchangeProfiles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TokenExchangeProfiles.Create(request) -> *management.CreateTokenExchangeProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new Token Exchange Profile within your tenant.

By using this feature, you agree to the applicable Free Trial terms in <a href="https://www.okta.com/legal/">Okta’s Master Subscription Agreement</a>. It is your responsibility to securely validate the user’s subject_token. See <a href="https://auth0.com/docs/authenticate/custom-token-exchange">User Guide</a> for more details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateTokenExchangeProfileRequestContent{
        Name: "name",
        SubjectTokenType: "subject_token_type",
        ActionID: "action_id",
        Type: management.TokenExchangeProfileTypeEnum(
            "custom_authentication",
        ),
    }
client.TokenExchangeProfiles.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Friendly name of this profile.
    
</dd>
</dl>

<dl>
<dd>

**subjectTokenType:** `string` — Subject token type for this profile. When receiving a token exchange request on the Authentication API, the corresponding token exchange profile with a matching subject_token_type will be executed. This must be a URI.
    
</dd>
</dl>

<dl>
<dd>

**actionID:** `string` — The ID of the Custom Token Exchange action to execute for this profile, in order to validate the subject_token. The action must use the custom-token-exchange trigger.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `management.TokenExchangeProfileTypeEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TokenExchangeProfiles.Get(ID) -> *management.GetTokenExchangeProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a single Token Exchange Profile specified by ID.

By using this feature, you agree to the applicable Free Trial terms in <a href="https://www.okta.com/legal/">Okta’s Master Subscription Agreement</a>. It is your responsibility to securely validate the user’s subject_token. See <a href="https://auth0.com/docs/authenticate/custom-token-exchange">User Guide</a> for more details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.TokenExchangeProfiles.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the Token Exchange Profile to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TokenExchangeProfiles.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a Token Exchange Profile within your tenant.

By using this feature, you agree to the applicable Free Trial terms in <a href="https://www.okta.com/legal/">Okta's Master Subscription Agreement</a>. It is your responsibility to securely validate the user's subject_token. See <a href="https://auth0.com/docs/authenticate/custom-token-exchange">User Guide</a> for more details.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.TokenExchangeProfiles.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the Token Exchange Profile to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TokenExchangeProfiles.Update(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a Token Exchange Profile within your tenant.

By using this feature, you agree to the applicable Free Trial terms in <a href="https://www.okta.com/legal/">Okta's Master Subscription Agreement</a>. It is your responsibility to securely validate the user's subject_token. See <a href="https://auth0.com/docs/authenticate/custom-token-exchange">User Guide</a> for more details.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateTokenExchangeProfileRequestContent{}
client.TokenExchangeProfiles.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the Token Exchange Profile to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Friendly name of this profile.
    
</dd>
</dl>

<dl>
<dd>

**subjectTokenType:** `*string` — Subject token type for this profile. When receiving a token exchange request on the Authentication API, the corresponding token exchange profile with a matching subject_token_type will be executed. This must be a URI.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## UserAttributeProfiles
<details><summary><code>client.UserAttributeProfiles.List() -> *management.ListUserAttributeProfilesPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of User Attribute Profiles. This endpoint supports Checkpoint pagination.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserAttributeProfileRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.UserAttributeProfiles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 5.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserAttributeProfiles.Create(request) -> *management.CreateUserAttributeProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a single User Attribute Profile specified by ID. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateUserAttributeProfileRequestContent{
        Name: "name",
        UserAttributes: map[string]*management.UserAttributeProfileUserAttributeAdditionalProperties{
            "key": &management.UserAttributeProfileUserAttributeAdditionalProperties{
                Description: "description",
                Label: "label",
                ProfileRequired: true,
                Auth0Mapping: "auth0_mapping",
            },
        },
    }
client.UserAttributeProfiles.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `management.UserAttributeProfileName` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*management.UserAttributeProfileUserID` 
    
</dd>
</dl>

<dl>
<dd>

**userAttributes:** `management.UserAttributeProfileUserAttributes` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserAttributeProfiles.ListTemplates() -> *management.ListUserAttributeProfileTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of User Attribute Profile Templates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserAttributeProfiles.ListTemplates(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserAttributeProfiles.GetTemplate(ID) -> *management.GetUserAttributeProfileTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a User Attribute Profile Template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserAttributeProfiles.GetTemplate(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user-attribute-profile-template to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserAttributeProfiles.Get(ID) -> *management.GetUserAttributeProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a single User Attribute Profile specified by ID. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserAttributeProfiles.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user-attribute-profile to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserAttributeProfiles.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a single User Attribute Profile specified by ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserAttributeProfiles.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user-attribute-profile to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserAttributeProfiles.Update(ID, request) -> *management.UpdateUserAttributeProfileResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the details of a specific User attribute profile, such as name, user_id and user_attributes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateUserAttributeProfileRequestContent{}
client.UserAttributeProfiles.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user attribute profile to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*management.UserAttributeProfileName` 
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*management.UserAttributeProfilePatchUserID` 
    
</dd>
</dl>

<dl>
<dd>

**userAttributes:** `*management.UserAttributeProfileUserAttributes` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## UserBlocks
<details><summary><code>client.UserBlocks.ListByIdentifier() -> *management.ListUserBlocksByIdentifierResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of all <a href="https://auth0.com/docs/secure/attack-protection/brute-force-protection">Brute-force Protection</a> blocks for a user with the given identifier (username, phone number, or email).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserBlocksByIdentifierRequestParameters{
        Identifier: "identifier",
        ConsiderBruteForceEnablement: management.Bool(
            true,
        ),
    }
client.UserBlocks.ListByIdentifier(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Should be any of a username, phone number, or email.
    
</dd>
</dl>

<dl>
<dd>

**considerBruteForceEnablement:** `*bool` 


          If true and Brute Force Protection is enabled and configured to block logins, will return a list of blocked IP addresses.
          If true and Brute Force Protection is disabled, will return an empty list.
        
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserBlocks.DeleteByIdentifier() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove all <a href="https://auth0.com/docs/secure/attack-protection/brute-force-protection">Brute-force Protection</a> blocks for the user with the given identifier (username, phone number, or email).

Note: This endpoint does not unblock users that were <a href="https://auth0.com/docs/user-profile#block-and-unblock-a-user">blocked by a tenant administrator</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteUserBlocksByIdentifierRequestParameters{
        Identifier: "identifier",
    }
client.UserBlocks.DeleteByIdentifier(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — Should be any of a username, phone number, or email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserBlocks.List(ID) -> *management.ListUserBlocksResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of all <a href="https://auth0.com/docs/secure/attack-protection/brute-force-protection">Brute-force Protection</a> blocks for the user with the given ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserBlocksRequestParameters{
        ConsiderBruteForceEnablement: management.Bool(
            true,
        ),
    }
client.UserBlocks.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — user_id of the user blocks to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**considerBruteForceEnablement:** `*bool` 


          If true and Brute Force Protection is enabled and configured to block logins, will return a list of blocked IP addresses.
          If true and Brute Force Protection is disabled, will return an empty list.
        
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.UserBlocks.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove all <a href="https://auth0.com/docs/secure/attack-protection/brute-force-protection">Brute-force Protection</a> blocks for the user with the given ID.

Note: This endpoint does not unblock users that were <a href="https://auth0.com/docs/user-profile#block-and-unblock-a-user">blocked by a tenant administrator</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.UserBlocks.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The user_id of the user to update.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users
<details><summary><code>client.Users.List() -> *management.ListUsersOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of users. It is possible to:

- Specify a search criteria for users
- Sort the users to be returned
- Select the fields to be returned
- Specify the number of users to retrieve per page and the page index
 <!-- only v3 is available -->
The <code>q</code> query parameter can be used to get users that match the specified criteria <a href="https://auth0.com/docs/users/search/v3/query-syntax">using query string syntax.</a>

<a href="https://auth0.com/docs/users/search/v3">Learn more about searching for users.</a>

Read about <a href="https://auth0.com/docs/users/search/best-practices">best practices</a> when working with the API endpoints for retrieving users.

Auth0 limits the number of users you can return. If you exceed this threshold, please redefine your search, use the <a href="https://auth0.com/docs/api/management/v2#!/Jobs/post_users_exports">export job</a>, or the <a href="https://auth0.com/docs/extensions/user-import-export">User Import / Export</a> extension.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUsersRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Sort: management.String(
            "sort",
        ),
        Connection: management.String(
            "connection",
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        Q: management.String(
            "q",
        ),
        SearchEngine: management.SearchEngineVersionsEnumV1.Ptr(),
        PrimaryOrder: management.Bool(
            true,
        ),
    }
client.Users.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Field to sort by. Use <code>field:order</code> where order is <code>1</code> for ascending and <code>-1</code> for descending. e.g. <code>created_at:1</code>
    
</dd>
</dl>

<dl>
<dd>

**connection:** `*string` — Connection filter. Only applies when using <code>search_engine=v1</code>. To filter by connection with <code>search_engine=v2|v3</code>, use <code>q=identities.connection:"connection_name"</code>
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — Query in <a target='_new' href ='http://www.lucenetutorial.com/lucene-query-syntax.html'>Lucene query string syntax</a>. Some query types cannot be used on metadata fields, for details see <a href='https://auth0.com/docs/users/search/v3/query-syntax#searchable-fields'>Searchable Fields</a>.
    
</dd>
</dl>

<dl>
<dd>

**searchEngine:** `*management.SearchEngineVersionsEnum` — The version of the search engine
    
</dd>
</dl>

<dl>
<dd>

**primaryOrder:** `*bool` — If true (default), results are returned in a deterministic order. If false, results may be returned in a non-deterministic order, which can enhance performance for complex queries targeting a small number of users. Set to false only when consistent ordering and pagination is not required.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Create(request) -> *management.CreateUserResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new user for a given <a href="https://auth0.com/docs/connections/database">database</a> or <a href="https://auth0.com/docs/connections/passwordless">passwordless</a> connection.

Note: <code>connection</code> is required but other parameters such as <code>email</code> and <code>password</code> are dependent upon the type of connection.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateUserRequestContent{
        Connection: "connection",
    }
client.Users.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**email:** `*string` — The user's email.
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `*string` — The user's phone number (following the E.164 recommendation).
    
</dd>
</dl>

<dl>
<dd>

**userMetadata:** `*management.UserMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**blocked:** `*bool` — Whether this user was blocked by an administrator (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**emailVerified:** `*bool` — Whether this email address is verified (true) or unverified (false). User will receive a verification email after creation if `email_verified` is false or not specified
    
</dd>
</dl>

<dl>
<dd>

**phoneVerified:** `*bool` — Whether this phone number has been verified (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**appMetadata:** `*management.AppMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**givenName:** `*string` — The user's given name(s).
    
</dd>
</dl>

<dl>
<dd>

**familyName:** `*string` — The user's family name(s).
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — The user's full name.
    
</dd>
</dl>

<dl>
<dd>

**nickname:** `*string` — The user's nickname.
    
</dd>
</dl>

<dl>
<dd>

**picture:** `*string` — A URI pointing to the user's picture.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*string` — The external user's id provided by the identity provider.
    
</dd>
</dl>

<dl>
<dd>

**connection:** `string` — Name of the connection this user should be created in.
    
</dd>
</dl>

<dl>
<dd>

**password:** `*string` — Initial password for this user. Only valid for auth0 connection strategy.
    
</dd>
</dl>

<dl>
<dd>

**verifyEmail:** `*bool` — Whether the user will receive a verification email after creation (true) or no email (false). Overrides behavior of `email_verified` parameter.
    
</dd>
</dl>

<dl>
<dd>

**username:** `*string` — The user's username. Only valid if the connection requires a username.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.ListUsersByEmail() -> []*management.UserResponseSchema</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Find users by email. If Auth0 is the identity provider (idP), the email address associated with a user is saved in lower case, regardless of how you initially provided it. 

For example, if you register a user as JohnSmith@example.com, Auth0 saves the user's email as johnsmith@example.com. 

Therefore, when using this endpoint, make sure that you are searching for users via email addresses using the correct case.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUsersByEmailRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        Email: "email",
    }
client.Users.ListUsersByEmail(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false). Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` — Email address to search for (case-sensitive).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Get(ID) -> *management.GetUserResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve user details. A list of fields to include or exclude may also be specified. For more information, see <a href="https://auth0.com/docs/manage-users/user-search/retrieve-users-with-get-users-endpoint">Retrieve Users with the Get Users Endpoint</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetUserRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Users.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a user by user ID. This action cannot be undone. For Auth0 Dashboard instructions, see <a href="https://auth0.com/docs/manage-users/user-accounts/delete-users">Delete Users</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Update(ID, request) -> *management.UpdateUserResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a user.

These are the attributes that can be updated at the root level:

<ul>
    <li>app_metadata</li>
    <li>blocked</li>
    <li>email</li>
    <li>email_verified</li>
    <li>family_name</li>
    <li>given_name</li>
    <li>name</li>
    <li>nickname</li>
    <li>password</li>
    <li>phone_number</li>
    <li>phone_verified</li>
    <li>picture</li>
    <li>username</li>
    <li>user_metadata</li>
    <li>verify_email</li>
</ul>

Some considerations:
<ul>
    <li>The properties of the new object will replace the old ones.</li>
    <li>The metadata fields are an exception to this rule (<code>user_metadata</code> and <code>app_metadata</code>). These properties are merged instead of being replaced but be careful, the merge only occurs on the first level.</li>
    <li>If you are updating <code>email</code>, <code>email_verified</code>, <code>phone_number</code>, <code>phone_verified</code>, <code>username</code> or <code>password</code> of a secondary identity, you need to specify the <code>connection</code> property too.</li>
    <li>If you are updating <code>email</code> or <code>phone_number</code> you can specify, optionally, the <code>client_id</code> property.</li>
    <li>Updating <code>email_verified</code> is not supported for enterprise and passwordless sms connections.</li>
    <li>Updating the <code>blocked</code> to <code>false</code> does not affect the user's blocked state from an excessive amount of incorrectly provided credentials. Use the "Unblock a user" endpoint from the "User Blocks" API to change the user's state.</li>
    <li>Supported attributes can be unset by supplying <code>null</code> as the value.</li>
</ul>

<h5>Updating a field (non-metadata property)</h5>
To mark the email address of a user as verified, the body to send should be:
<pre><code>{ "email_verified": true }</code></pre>

<h5>Updating a user metadata root property</h5>Let's assume that our test user has the following <code>user_metadata</code>:
<pre><code>{ "user_metadata" : { "profileCode": 1479 } }</code></pre>

To add the field <code>addresses</code> the body to send should be:
<pre><code>{ "user_metadata" : { "addresses": {"work_address": "100 Industrial Way"} }}</code></pre>

The modified object ends up with the following <code>user_metadata</code> property:<pre><code>{
  "user_metadata": {
    "profileCode": 1479,
    "addresses": { "work_address": "100 Industrial Way" }
  }
}</code></pre>

<h5>Updating an inner user metadata property</h5>If there's existing user metadata to which we want to add  <code>"home_address": "742 Evergreen Terrace"</code> (using the <code>addresses</code> property) we should send the whole <code>addresses</code> object. Since this is a first-level object, the object will be merged in, but its own properties will not be. The body to send should be:
<pre><code>{
  "user_metadata": {
    "addresses": {
      "work_address": "100 Industrial Way",
      "home_address": "742 Evergreen Terrace"
    }
  }
}</code></pre>

The modified object ends up with the following <code>user_metadata</code> property:
<pre><code>{
  "user_metadata": {
    "profileCode": 1479,
    "addresses": {
      "work_address": "100 Industrial Way",
      "home_address": "742 Evergreen Terrace"
    }
  }
}</code></pre>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateUserRequestContent{}
client.Users.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to update.
    
</dd>
</dl>

<dl>
<dd>

**blocked:** `*bool` — Whether this user was blocked by an administrator (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**emailVerified:** `*bool` — Whether this email address is verified (true) or unverified (false). If set to false the user will not receive a verification email unless `verify_email` is set to true.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Email address of this user.
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `*string` — The user's phone number (following the E.164 recommendation).
    
</dd>
</dl>

<dl>
<dd>

**phoneVerified:** `*bool` — Whether this phone number has been verified (true) or not (false).
    
</dd>
</dl>

<dl>
<dd>

**userMetadata:** `*management.UserMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**appMetadata:** `*management.AppMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**givenName:** `*string` — Given name/first name/forename of this user.
    
</dd>
</dl>

<dl>
<dd>

**familyName:** `*string` — Family name/last name/surname of this user.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Name of this user.
    
</dd>
</dl>

<dl>
<dd>

**nickname:** `*string` — Preferred nickname or alias of this user.
    
</dd>
</dl>

<dl>
<dd>

**picture:** `*string` — URL to picture, photo, or avatar of this user.
    
</dd>
</dl>

<dl>
<dd>

**verifyEmail:** `*bool` — Whether this user will receive a verification email after creation (true) or no email (false). Overrides behavior of `email_verified` parameter.
    
</dd>
</dl>

<dl>
<dd>

**verifyPhoneNumber:** `*bool` — Whether this user will receive a text after changing the phone number (true) or no text (false). Only valid when changing phone number for SMS connections.
    
</dd>
</dl>

<dl>
<dd>

**password:** `*string` — New password for this user. Only valid for database connections.
    
</dd>
</dl>

<dl>
<dd>

**connection:** `*string` — Name of the connection to target for this user update.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — Auth0 client ID. Only valid when updating email address.
    
</dd>
</dl>

<dl>
<dd>

**username:** `*string` — The user's username. Only valid if the connection requires a username.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.RegenerateRecoveryCode(ID) -> *management.RegenerateUsersRecoveryCodeResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove an existing multi-factor authentication (MFA) <a href="https://auth0.com/docs/secure/multi-factor-authentication/reset-user-mfa">recovery code</a> and generate a new one. If a user cannot access the original device or account used for MFA enrollment, they can use a recovery code to authenticate. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.RegenerateRecoveryCode(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to regenerate a multi-factor authentication recovery code for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.RevokeAccess(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revokes selected resources related to a user (sessions, refresh tokens, ...).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.RevokeUserAccessRequestContent{}
client.Users.RevokeAccess(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user.
    
</dd>
</dl>

<dl>
<dd>

**sessionID:** `*string` — ID of the session to revoke.
    
</dd>
</dl>

<dl>
<dd>

**preserveRefreshTokens:** `*bool` — Whether to preserve the refresh tokens associated with the session.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Actions Versions
<details><summary><code>client.Actions.Versions.List(ActionID) -> *management.ListActionVersionsPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all of an action's versions. An action version is created whenever an action is deployed. An action version is immutable, once created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListActionVersionsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
    }
client.Actions.Versions.List(
        context.TODO(),
        "actionId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**actionID:** `string` — The ID of the action.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Use this field to request a specific page of the list results.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — This field specify the maximum number of results to be returned by the server. 20 by default
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Versions.Get(ActionID, ID) -> *management.GetActionVersionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a specific version of an action. An action version is created whenever an action is deployed. An action version is immutable, once created.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Actions.Versions.Get(
        context.TODO(),
        "actionId",
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**actionID:** `string` — The ID of the action.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The ID of the action version.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Versions.Deploy(ActionID, ID, request) -> *management.DeployActionVersionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Performs the equivalent of a roll-back of an action to an earlier, specified version. Creates a new, deployed action version that is identical to the specified version. If this action is currently bound to a trigger, the system will begin executing the newly-created version immediately.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeployActionVersionRequestBodyParams{}
client.Actions.Versions.Deploy(
        context.TODO(),
        "actionId",
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**actionID:** `string` — The ID of an action.
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The ID of an action version.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.DeployActionVersionRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Actions Executions
<details><summary><code>client.Actions.Executions.Get(ID) -> *management.GetActionExecutionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve information about a specific execution of a trigger. Relevant execution IDs will be included in tenant logs generated as part of that authentication flow. Executions will only be stored for 10 days after their creation.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Actions.Executions.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the execution to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Actions Triggers
<details><summary><code>client.Actions.Triggers.List() -> *management.ListActionTriggersResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the set of triggers currently available within actions. A trigger is an extensibility point to which actions can be bound.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Actions.Triggers.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Actions Triggers Bindings
<details><summary><code>client.Actions.Triggers.Bindings.List(TriggerID) -> *management.ListActionBindingsPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the actions that are bound to a trigger. Once an action is created and deployed, it must be attached (i.e. bound) to a trigger so that it will be executed as part of a flow. The list of actions returned reflects the order in which they will be executed during the appropriate flow.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListActionTriggerBindingsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
    }
client.Actions.Triggers.Bindings.List(
        context.TODO(),
        "triggerId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**triggerID:** `management.ActionTriggerTypeEnum` — An actions extensibility point.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Use this field to request a specific page of the list results.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — The maximum number of results to be returned in a single request. 20 by default
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Actions.Triggers.Bindings.UpdateMany(TriggerID, request) -> *management.UpdateActionBindingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the actions that are bound (i.e. attached) to a trigger. Once an action is created and deployed, it must be attached (i.e. bound) to a trigger so that it will be executed as part of a flow. The order in which the actions are provided will determine the order in which they are executed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateActionBindingsRequestContent{}
client.Actions.Triggers.Bindings.UpdateMany(
        context.TODO(),
        "triggerId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**triggerID:** `management.ActionTriggerTypeEnum` — An actions extensibility point.
    
</dd>
</dl>

<dl>
<dd>

**bindings:** `[]*management.ActionBindingWithRef` — The actions that will be bound to this trigger. The order in which they are included will be the order in which they are executed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Anomaly Blocks
<details><summary><code>client.Anomaly.Blocks.CheckIP(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Check if the given IP address is blocked via the <a href="https://auth0.com/docs/configure/attack-protection/suspicious-ip-throttling">Suspicious IP Throttling</a> due to multiple suspicious attempts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Anomaly.Blocks.CheckIP(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `management.AnomalyIPFormat` — IP address to check.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Anomaly.Blocks.UnblockIP(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a block imposed by <a href="https://auth0.com/docs/configure/attack-protection/suspicious-ip-throttling">Suspicious IP Throttling</a> for the given IP address.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Anomaly.Blocks.UnblockIP(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `management.AnomalyIPFormat` — IP address to unblock.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AttackProtection BotDetection
<details><summary><code>client.AttackProtection.BotDetection.Get() -> *management.GetBotDetectionSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the Bot Detection configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AttackProtection.BotDetection.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AttackProtection.BotDetection.Update(request) -> *management.UpdateBotDetectionSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the Bot Detection configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateBotDetectionSettingsRequestContent{}
client.AttackProtection.BotDetection.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**botDetectionLevel:** `*management.BotDetectionLevelEnum` 
    
</dd>
</dl>

<dl>
<dd>

**challengePasswordPolicy:** `*management.BotDetectionChallengePolicyPasswordFlowEnum` 
    
</dd>
</dl>

<dl>
<dd>

**challengePasswordlessPolicy:** `*management.BotDetectionChallengePolicyPasswordlessFlowEnum` 
    
</dd>
</dl>

<dl>
<dd>

**challengePasswordResetPolicy:** `*management.BotDetectionChallengePolicyPasswordResetFlowEnum` 
    
</dd>
</dl>

<dl>
<dd>

**allowlist:** `*management.BotDetectionAllowlist` 
    
</dd>
</dl>

<dl>
<dd>

**monitoringModeEnabled:** `*management.BotDetectionMonitoringModeEnabled` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AttackProtection BreachedPasswordDetection
<details><summary><code>client.AttackProtection.BreachedPasswordDetection.Get() -> *management.GetBreachedPasswordDetectionSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the Breached Password Detection configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AttackProtection.BreachedPasswordDetection.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AttackProtection.BreachedPasswordDetection.Update(request) -> *management.UpdateBreachedPasswordDetectionSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update details of the Breached Password Detection configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateBreachedPasswordDetectionSettingsRequestContent{}
client.AttackProtection.BreachedPasswordDetection.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enabled:** `*bool` — Whether or not breached password detection is active.
    
</dd>
</dl>

<dl>
<dd>

**shields:** `[]*management.BreachedPasswordDetectionShieldsEnum` 

Action to take when a breached password is detected during a login.
      Possible values: <code>block</code>, <code>user_notification</code>, <code>admin_notification</code>.
    
</dd>
</dl>

<dl>
<dd>

**adminNotificationFrequency:** `[]*management.BreachedPasswordDetectionAdminNotificationFrequencyEnum` 

When "admin_notification" is enabled, determines how often email notifications are sent.
        Possible values: <code>immediately</code>, <code>daily</code>, <code>weekly</code>, <code>monthly</code>.
    
</dd>
</dl>

<dl>
<dd>

**method:** `*management.BreachedPasswordDetectionMethodEnum` 
    
</dd>
</dl>

<dl>
<dd>

**stage:** `*management.BreachedPasswordDetectionStage` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AttackProtection BruteForceProtection
<details><summary><code>client.AttackProtection.BruteForceProtection.Get() -> *management.GetBruteForceSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the Brute-force Protection configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AttackProtection.BruteForceProtection.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AttackProtection.BruteForceProtection.Update(request) -> *management.UpdateBruteForceSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the Brute-force Protection configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateBruteForceSettingsRequestContent{}
client.AttackProtection.BruteForceProtection.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enabled:** `*bool` — Whether or not brute force attack protections are active.
    
</dd>
</dl>

<dl>
<dd>

**shields:** `[]*attackprotection.UpdateBruteForceSettingsRequestContentShieldsItem` 

Action to take when a brute force protection threshold is violated.
        Possible values: <code>block</code>, <code>user_notification</code>.
    
</dd>
</dl>

<dl>
<dd>

**allowlist:** `[]string` — List of trusted IP addresses that will not have attack protection enforced against them.
    
</dd>
</dl>

<dl>
<dd>

**mode:** `*attackprotection.UpdateBruteForceSettingsRequestContentMode` 

Account Lockout: Determines whether or not IP address is used when counting failed attempts.
          Possible values: <code>count_per_identifier_and_ip</code>, <code>count_per_identifier</code>.
    
</dd>
</dl>

<dl>
<dd>

**maxAttempts:** `*int` — Maximum number of unsuccessful attempts.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AttackProtection Captcha
<details><summary><code>client.AttackProtection.Captcha.Get() -> *management.GetAttackProtectionCaptchaResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the CAPTCHA configuration for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AttackProtection.Captcha.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AttackProtection.Captcha.Update(request) -> *management.UpdateAttackProtectionCaptchaResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update existing CAPTCHA configuration for your client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateAttackProtectionCaptchaRequestContent{}
client.AttackProtection.Captcha.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**activeProviderID:** `*management.AttackProtectionCaptchaProviderID` 
    
</dd>
</dl>

<dl>
<dd>

**arkose:** `*management.AttackProtectionUpdateCaptchaArkose` 
    
</dd>
</dl>

<dl>
<dd>

**authChallenge:** `*management.AttackProtectionCaptchaAuthChallengeRequest` 
    
</dd>
</dl>

<dl>
<dd>

**hcaptcha:** `*management.AttackProtectionUpdateCaptchaHcaptcha` 
    
</dd>
</dl>

<dl>
<dd>

**friendlyCaptcha:** `*management.AttackProtectionUpdateCaptchaFriendlyCaptcha` 
    
</dd>
</dl>

<dl>
<dd>

**recaptchaEnterprise:** `*management.AttackProtectionUpdateCaptchaRecaptchaEnterprise` 
    
</dd>
</dl>

<dl>
<dd>

**recaptchaV2:** `*management.AttackProtectionUpdateCaptchaRecaptchaV2` 
    
</dd>
</dl>

<dl>
<dd>

**simpleCaptcha:** `*management.AttackProtectionCaptchaSimpleCaptchaResponseContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AttackProtection SuspiciousIPThrottling
<details><summary><code>client.AttackProtection.SuspiciousIPThrottling.Get() -> *management.GetSuspiciousIPThrottlingSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the Suspicious IP Throttling configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.AttackProtection.SuspiciousIPThrottling.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AttackProtection.SuspiciousIPThrottling.Update(request) -> *management.UpdateSuspiciousIPThrottlingSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the details of the Suspicious IP Throttling configuration of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateSuspiciousIPThrottlingSettingsRequestContent{}
client.AttackProtection.SuspiciousIPThrottling.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enabled:** `*bool` — Whether or not suspicious IP throttling attack protections are active.
    
</dd>
</dl>

<dl>
<dd>

**shields:** `[]*management.SuspiciousIPThrottlingShieldsEnum` 

Action to take when a suspicious IP throttling threshold is violated.
          Possible values: <code>block</code>, <code>admin_notification</code>.
    
</dd>
</dl>

<dl>
<dd>

**allowlist:** `*management.SuspiciousIPThrottlingAllowlist` 
    
</dd>
</dl>

<dl>
<dd>

**stage:** `*management.SuspiciousIPThrottlingStage` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Branding Templates
<details><summary><code>client.Branding.Templates.GetUniversalLogin() -> *management.GetUniversalLoginTemplateResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Templates.GetUniversalLogin(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Templates.UpdateUniversalLogin(request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the Universal Login branding template.

<p>When <code>content-type</code> header is set to <code>application/json</code>:</p>
<pre>
{
  "template": "&lt;!DOCTYPE html&gt;{% assign resolved_dir = dir | default: "auto" %}&lt;html lang="{{locale}}" dir="{{resolved_dir}}"&gt;&lt;head&gt;{%- auth0:head -%}&lt;/head&gt;&lt;body class="_widget-auto-layout"&gt;{%- auth0:widget -%}&lt;/body&gt;&lt;/html&gt;"
}
</pre>

<p>
  When <code>content-type</code> header is set to <code>text/html</code>:
</p>
<pre>
&lt!DOCTYPE html&gt;
{% assign resolved_dir = dir | default: "auto" %}
&lt;html lang="{{locale}}" dir="{{resolved_dir}}"&gt;
  &lt;head&gt;
    {%- auth0:head -%}
  &lt;/head&gt;
  &lt;body class="_widget-auto-layout"&gt;
    {%- auth0:widget -%}
  &lt;/body&gt;
&lt;/html&gt;
</pre>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateUniversalLoginTemplateRequestContent{
        String: "string",
    }
client.Branding.Templates.UpdateUniversalLogin(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.UpdateUniversalLoginTemplateRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Templates.DeleteUniversalLogin() -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Templates.DeleteUniversalLogin(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Branding Themes
<details><summary><code>client.Branding.Themes.Create(request) -> *management.CreateBrandingThemeResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create branding theme.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateBrandingThemeRequestContent{
        Borders: &management.BrandingThemeBorders{
            ButtonBorderRadius: 1.1,
            ButtonBorderWeight: 1.1,
            ButtonsStyle: management.BrandingThemeBordersButtonsStyleEnumPill,
            InputBorderRadius: 1.1,
            InputBorderWeight: 1.1,
            InputsStyle: management.BrandingThemeBordersInputsStyleEnumPill,
            ShowWidgetShadow: true,
            WidgetBorderWeight: 1.1,
            WidgetCornerRadius: 1.1,
        },
        Colors: &management.BrandingThemeColors{
            BodyText: "body_text",
            Error: "error",
            Header: "header",
            Icons: "icons",
            InputBackground: "input_background",
            InputBorder: "input_border",
            InputFilledText: "input_filled_text",
            InputLabelsPlaceholders: "input_labels_placeholders",
            LinksFocusedComponents: "links_focused_components",
            PrimaryButton: "primary_button",
            PrimaryButtonLabel: "primary_button_label",
            SecondaryButtonBorder: "secondary_button_border",
            SecondaryButtonLabel: "secondary_button_label",
            Success: "success",
            WidgetBackground: "widget_background",
            WidgetBorder: "widget_border",
        },
        Fonts: &management.BrandingThemeFonts{
            BodyText: &management.BrandingThemeFontBodyText{
                Bold: true,
                Size: 1.1,
            },
            ButtonsText: &management.BrandingThemeFontButtonsText{
                Bold: true,
                Size: 1.1,
            },
            FontURL: "font_url",
            InputLabels: &management.BrandingThemeFontInputLabels{
                Bold: true,
                Size: 1.1,
            },
            Links: &management.BrandingThemeFontLinks{
                Bold: true,
                Size: 1.1,
            },
            LinksStyle: management.BrandingThemeFontLinksStyleEnumNormal,
            ReferenceTextSize: 1.1,
            Subtitle: &management.BrandingThemeFontSubtitle{
                Bold: true,
                Size: 1.1,
            },
            Title: &management.BrandingThemeFontTitle{
                Bold: true,
                Size: 1.1,
            },
        },
        PageBackground: &management.BrandingThemePageBackground{
            BackgroundColor: "background_color",
            BackgroundImageURL: "background_image_url",
            PageLayout: management.BrandingThemePageBackgroundPageLayoutEnumCenter,
        },
        Widget: &management.BrandingThemeWidget{
            HeaderTextAlignment: management.BrandingThemeWidgetHeaderTextAlignmentEnumCenter,
            LogoHeight: 1.1,
            LogoPosition: management.BrandingThemeWidgetLogoPositionEnumCenter,
            LogoURL: "logo_url",
            SocialButtonsLayout: management.BrandingThemeWidgetSocialButtonsLayoutEnumBottom,
        },
    }
client.Branding.Themes.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**borders:** `*management.BrandingThemeBorders` 
    
</dd>
</dl>

<dl>
<dd>

**colors:** `*management.BrandingThemeColors` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` — Display Name
    
</dd>
</dl>

<dl>
<dd>

**fonts:** `*management.BrandingThemeFonts` 
    
</dd>
</dl>

<dl>
<dd>

**pageBackground:** `*management.BrandingThemePageBackground` 
    
</dd>
</dl>

<dl>
<dd>

**widget:** `*management.BrandingThemeWidget` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Themes.GetDefault() -> *management.GetBrandingDefaultThemeResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve default branding theme.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Themes.GetDefault(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Themes.Get(ThemeID) -> *management.GetBrandingThemeResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve branding theme.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Themes.Get(
        context.TODO(),
        "themeId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**themeID:** `string` — The ID of the theme
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Themes.Delete(ThemeID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete branding theme.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Themes.Delete(
        context.TODO(),
        "themeId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**themeID:** `string` — The ID of the theme
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Themes.Update(ThemeID, request) -> *management.UpdateBrandingThemeResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update branding theme.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateBrandingThemeRequestContent{
        Borders: &management.BrandingThemeBorders{
            ButtonBorderRadius: 1.1,
            ButtonBorderWeight: 1.1,
            ButtonsStyle: management.BrandingThemeBordersButtonsStyleEnumPill,
            InputBorderRadius: 1.1,
            InputBorderWeight: 1.1,
            InputsStyle: management.BrandingThemeBordersInputsStyleEnumPill,
            ShowWidgetShadow: true,
            WidgetBorderWeight: 1.1,
            WidgetCornerRadius: 1.1,
        },
        Colors: &management.BrandingThemeColors{
            BodyText: "body_text",
            Error: "error",
            Header: "header",
            Icons: "icons",
            InputBackground: "input_background",
            InputBorder: "input_border",
            InputFilledText: "input_filled_text",
            InputLabelsPlaceholders: "input_labels_placeholders",
            LinksFocusedComponents: "links_focused_components",
            PrimaryButton: "primary_button",
            PrimaryButtonLabel: "primary_button_label",
            SecondaryButtonBorder: "secondary_button_border",
            SecondaryButtonLabel: "secondary_button_label",
            Success: "success",
            WidgetBackground: "widget_background",
            WidgetBorder: "widget_border",
        },
        Fonts: &management.BrandingThemeFonts{
            BodyText: &management.BrandingThemeFontBodyText{
                Bold: true,
                Size: 1.1,
            },
            ButtonsText: &management.BrandingThemeFontButtonsText{
                Bold: true,
                Size: 1.1,
            },
            FontURL: "font_url",
            InputLabels: &management.BrandingThemeFontInputLabels{
                Bold: true,
                Size: 1.1,
            },
            Links: &management.BrandingThemeFontLinks{
                Bold: true,
                Size: 1.1,
            },
            LinksStyle: management.BrandingThemeFontLinksStyleEnumNormal,
            ReferenceTextSize: 1.1,
            Subtitle: &management.BrandingThemeFontSubtitle{
                Bold: true,
                Size: 1.1,
            },
            Title: &management.BrandingThemeFontTitle{
                Bold: true,
                Size: 1.1,
            },
        },
        PageBackground: &management.BrandingThemePageBackground{
            BackgroundColor: "background_color",
            BackgroundImageURL: "background_image_url",
            PageLayout: management.BrandingThemePageBackgroundPageLayoutEnumCenter,
        },
        Widget: &management.BrandingThemeWidget{
            HeaderTextAlignment: management.BrandingThemeWidgetHeaderTextAlignmentEnumCenter,
            LogoHeight: 1.1,
            LogoPosition: management.BrandingThemeWidgetLogoPositionEnumCenter,
            LogoURL: "logo_url",
            SocialButtonsLayout: management.BrandingThemeWidgetSocialButtonsLayoutEnumBottom,
        },
    }
client.Branding.Themes.Update(
        context.TODO(),
        "themeId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**themeID:** `string` — The ID of the theme
    
</dd>
</dl>

<dl>
<dd>

**borders:** `*management.BrandingThemeBorders` 
    
</dd>
</dl>

<dl>
<dd>

**colors:** `*management.BrandingThemeColors` 
    
</dd>
</dl>

<dl>
<dd>

**displayName:** `*string` — Display Name
    
</dd>
</dl>

<dl>
<dd>

**fonts:** `*management.BrandingThemeFonts` 
    
</dd>
</dl>

<dl>
<dd>

**pageBackground:** `*management.BrandingThemePageBackground` 
    
</dd>
</dl>

<dl>
<dd>

**widget:** `*management.BrandingThemeWidget` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Branding Phone Providers
<details><summary><code>client.Branding.Phone.Providers.List() -> *management.ListBrandingPhoneProvidersResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a list of <a href="https://auth0.com/docs/customize/phone-messages/configure-phone-messaging-providers">phone providers</a> details set for a Tenant. A list of fields to include or exclude may also be specified.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListBrandingPhoneProvidersRequestParameters{
        Disabled: management.Bool(
            true,
        ),
    }
client.Branding.Phone.Providers.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disabled:** `*bool` — Whether the provider is enabled (false) or disabled (true).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Providers.Create(request) -> *management.CreateBrandingPhoneProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a <a href="https://auth0.com/docs/customize/phone-messages/configure-phone-messaging-providers">phone provider</a>.
The <code>credentials</code> object requires different properties depending on the phone provider (which is specified using the <code>name</code> property).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateBrandingPhoneProviderRequestContent{
        Name: management.PhoneProviderNameEnumTwilio,
        Credentials: &management.PhoneProviderCredentials{
            TwilioProviderCredentials: &management.TwilioProviderCredentials{
                AuthToken: "auth_token",
            },
        },
    }
client.Branding.Phone.Providers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*management.PhoneProviderNameEnum` 
    
</dd>
</dl>

<dl>
<dd>

**disabled:** `*bool` — Whether the provider is enabled (false) or disabled (true).
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `*management.PhoneProviderConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**credentials:** `*management.PhoneProviderCredentials` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Providers.Get(ID) -> *management.GetBrandingPhoneProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve <a href="https://auth0.com/docs/customize/phone-messages/configure-phone-messaging-providers">phone provider</a> details. A list of fields to include or exclude may also be specified.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Phone.Providers.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Providers.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete the configured phone provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Phone.Providers.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Providers.Update(ID, request) -> *management.UpdateBrandingPhoneProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a <a href="https://auth0.com/docs/customize/phone-messages/configure-phone-messaging-providers">phone provider</a>.
The <code>credentials</code> object requires different properties depending on the phone provider (which is specified using the <code>name</code> property).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateBrandingPhoneProviderRequestContent{}
client.Branding.Phone.Providers.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*management.PhoneProviderNameEnum` 
    
</dd>
</dl>

<dl>
<dd>

**disabled:** `*bool` — Whether the provider is enabled (false) or disabled (true).
    
</dd>
</dl>

<dl>
<dd>

**credentials:** `*management.PhoneProviderCredentials` 
    
</dd>
</dl>

<dl>
<dd>

**configuration:** `*management.PhoneProviderConfiguration` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Providers.Test(ID, request) -> *management.CreatePhoneProviderSendTestResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreatePhoneProviderSendTestRequestContent{
        To: "to",
    }
client.Branding.Phone.Providers.Test(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — The recipient phone number to receive a given notification.
    
</dd>
</dl>

<dl>
<dd>

**deliveryMethod:** `*management.PhoneProviderDeliveryMethodEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Branding Phone Templates
<details><summary><code>client.Branding.Phone.Templates.List() -> *management.ListPhoneTemplatesResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListPhoneTemplatesRequestParameters{
        Disabled: management.Bool(
            true,
        ),
    }
client.Branding.Phone.Templates.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**disabled:** `*bool` — Whether the template is enabled (false) or disabled (true).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Templates.Create(request) -> *management.CreatePhoneTemplateResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreatePhoneTemplateRequestContent{}
client.Branding.Phone.Templates.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `*management.PhoneTemplateNotificationTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**disabled:** `*bool` — Whether the template is enabled (false) or disabled (true).
    
</dd>
</dl>

<dl>
<dd>

**content:** `*management.PhoneTemplateContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Templates.Get(ID) -> *management.GetPhoneTemplateResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Phone.Templates.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Templates.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Branding.Phone.Templates.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Templates.Update(ID, request) -> *management.UpdatePhoneTemplateResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdatePhoneTemplateRequestContent{}
client.Branding.Phone.Templates.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**content:** `*management.PartialPhoneTemplateContent` 
    
</dd>
</dl>

<dl>
<dd>

**disabled:** `*bool` — Whether the template is enabled (false) or disabled (true).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Templates.Reset(ID, request) -> *management.ResetPhoneTemplateResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Branding.Phone.Templates.Reset(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.ResetPhoneTemplateRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Branding.Phone.Templates.Test(ID, request) -> *management.CreatePhoneTemplateTestNotificationResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreatePhoneTemplateTestNotificationRequestContent{
        To: "to",
    }
client.Branding.Phone.Templates.Test(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `string` — Destination of the testing phone notification
    
</dd>
</dl>

<dl>
<dd>

**deliveryMethod:** `*management.PhoneProviderDeliveryMethodEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ClientGrants Organizations
<details><summary><code>client.ClientGrants.Organizations.List(ID) -> *management.ListClientGrantOrganizationsPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListClientGrantOrganizationsRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.ClientGrants.Organizations.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client grant
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Clients Credentials
<details><summary><code>client.Clients.Credentials.List(ClientID) -> []*management.ClientCredential</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the details of a client credential.

<b>Important</b>: To enable credentials to be used for a client authentication method, set the <code>client_authentication_methods</code> property on the client. To enable credentials to be used for JWT-Secured Authorization requests set the <code>signed_request_object</code> property on the client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Clients.Credentials.List(
        context.TODO(),
        "client_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — ID of the client.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Credentials.Create(ClientID, request) -> *management.PostClientCredentialResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a client credential associated to your application. Credentials can be used to configure Private Key JWT and mTLS authentication methods, as well as for JWT-secured Authorization requests.

<h5>Public Key</h5>Public Key credentials can be used to set up Private Key JWT client authentication and JWT-secured Authorization requests.

Sample: <pre><code>{
  "credential_type": "public_key",
  "name": "string",
  "pem": "string",
  "alg": "RS256",
  "parse_expiry_from_cert": false,
  "expires_at": "2022-12-31T23:59:59Z"
}</code></pre>
<h5>Certificate (CA-signed & self-signed)</h5>Certificate credentials can be used to set up mTLS client authentication. CA-signed certificates can be configured either with a signed certificate or with just the certificate Subject DN.

CA-signed Certificate Sample (pem): <pre><code>{
  "credential_type": "x509_cert",
  "name": "string",
  "pem": "string"
}</code></pre>CA-signed Certificate Sample (subject_dn): <pre><code>{
  "credential_type": "cert_subject_dn",
  "name": "string",
  "subject_dn": "string"
}</code></pre>Self-signed Certificate Sample: <pre><code>{
  "credential_type": "cert_subject_dn",
  "name": "string",
  "pem": "string"
}</code></pre>

The credential will be created but not yet enabled for use until you set the corresponding properties in the client:
<ul>
  <li>To enable the credential for Private Key JWT or mTLS authentication methods, set the <code>client_authentication_methods</code> property on the client. For more information, read <a href="https://auth0.com/docs/get-started/applications/configure-private-key-jwt">Configure Private Key JWT Authentication</a> and <a href="https://auth0.com/docs/get-started/applications/configure-mtls">Configure mTLS Authentication</a></li>
  <li>To enable the credential for JWT-secured Authorization requests, set the <code>signed_request_object</code>property on the client. For more information, read <a href="https://auth0.com/docs/get-started/applications/configure-jar">Configure JWT-secured Authorization Requests (JAR)</a></li>
</ul>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.PostClientCredentialRequestContent{
        CredentialType: management.ClientCredentialTypeEnumPublicKey,
    }
client.Clients.Credentials.Create(
        context.TODO(),
        "client_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — ID of the client.
    
</dd>
</dl>

<dl>
<dd>

**credentialType:** `*management.ClientCredentialTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Friendly name for a credential.
    
</dd>
</dl>

<dl>
<dd>

**subjectDn:** `*string` — Subject Distinguished Name. Mutually exclusive with `pem` property. Applies to `cert_subject_dn` credential type.
    
</dd>
</dl>

<dl>
<dd>

**pem:** `*string` — PEM-formatted public key (SPKI and PKCS1) or X509 certificate. Must be JSON escaped.
    
</dd>
</dl>

<dl>
<dd>

**alg:** `*management.PublicKeyCredentialAlgorithmEnum` 
    
</dd>
</dl>

<dl>
<dd>

**parseExpiryFromCert:** `*bool` — Parse expiry from x509 certificate. If true, attempts to parse the expiry date from the provided PEM. Applies to `public_key` credential type.
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` — The ISO 8601 formatted date representing the expiration of the credential. If not specified (not recommended), the credential never expires. Applies to `public_key` credential type.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Credentials.Get(ClientID, CredentialID) -> *management.GetClientCredentialResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the details of a client credential.

<b>Important</b>: To enable credentials to be used for a client authentication method, set the <code>client_authentication_methods</code> property on the client. To enable credentials to be used for JWT-Secured Authorization requests set the <code>signed_request_object</code> property on the client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Clients.Credentials.Get(
        context.TODO(),
        "client_id",
        "credential_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — ID of the client.
    
</dd>
</dl>

<dl>
<dd>

**credentialID:** `string` — ID of the credential.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Credentials.Delete(ClientID, CredentialID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a client credential you previously created. May be enabled or disabled. For more information, read <a href="https://www.auth0.com/docs/get-started/authentication-and-authorization-flow/client-credentials-flow">Client Credential Flow</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Clients.Credentials.Delete(
        context.TODO(),
        "client_id",
        "credential_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — ID of the client.
    
</dd>
</dl>

<dl>
<dd>

**credentialID:** `string` — ID of the credential to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Clients.Credentials.Update(ClientID, CredentialID, request) -> *management.PatchClientCredentialResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Change a client credential you previously created. May be enabled or disabled. For more information, read <a href="https://www.auth0.com/docs/get-started/authentication-and-authorization-flow/client-credentials-flow">Client Credential Flow</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.PatchClientCredentialRequestContent{}
client.Clients.Credentials.Update(
        context.TODO(),
        "client_id",
        "credential_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**clientID:** `string` — ID of the client.
    
</dd>
</dl>

<dl>
<dd>

**credentialID:** `string` — ID of the credential.
    
</dd>
</dl>

<dl>
<dd>

**expiresAt:** `*time.Time` — The ISO 8601 formatted date representing the expiration of the credential.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Clients Connections
<details><summary><code>client.Clients.Connections.Get(ID) -> *management.ListClientConnectionsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all connections that are enabled for the specified <a href="https://www.auth0.com/docs/get-started/applications"> Application</a>, using checkpoint pagination. A list of fields to include or exclude for each connection may also be specified.
<ul>
  <li>
    This endpoint requires the <code>read:connections</code> scope and any one of <code>read:clients</code> or <code>read:client_summary</code>.
  </li>
  <li>
    <b>Note</b>: The first time you call this endpoint, omit the <code>from</code> parameter. If there are more results, a <code>next</code> value is included in the response. You can use this for subsequent API calls. When <code>next</code> is no longer included in the response, no further results are remaining.
  </li>
</ul>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ConnectionsGetRequest{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Clients.Connections.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the client for which to retrieve enabled connections.
    
</dd>
</dl>

<dl>
<dd>

**strategy:** `*management.ConnectionStrategyEnum` — Provide strategies to only retrieve connections with such strategies
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — A comma separated list of fields to include or exclude (depending on include_fields) from the result, empty to retrieve all fields
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — <code>true</code> if the fields specified are to be included in the result, <code>false</code> otherwise (defaults to <code>true</code>)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections Clients
<details><summary><code>client.Connections.Clients.Get(ID) -> *management.GetConnectionEnabledClientsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all clients that have the specified <a href="https://auth0.com/docs/authenticate/identity-providers">connection</a> enabled.

<b>Note</b>: The first time you call this endpoint, omit the <code>from</code> parameter. If there are more results, a <code>next</code> value is included in the response. You can use this for subsequent API calls. When <code>next</code> is no longer included in the response, no further results are remaining.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetConnectionEnabledClientsRequestParameters{
        Take: management.Int(
            1,
        ),
        From: management.String(
            "from",
        ),
    }
client.Connections.Clients.Get(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection for which enabled clients are to be retrieved
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Clients.Update(ID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.UpdateEnabledClientConnectionsRequestContentItem{
        &management.UpdateEnabledClientConnectionsRequestContentItem{
            ClientID: "client_id",
            Status: true,
        },
    }
client.Connections.Clients.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to modify
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateEnabledClientConnectionsRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections DirectoryProvisioning
<details><summary><code>client.Connections.DirectoryProvisioning.Get(ID) -> *management.GetDirectoryProvisioningResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the directory provisioning configuration of a connection.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.DirectoryProvisioning.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to retrieve its directory provisioning configuration
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.DirectoryProvisioning.Create(ID, request) -> *management.CreateDirectoryProvisioningResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a directory provisioning configuration for a connection.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateDirectoryProvisioningRequestContent{}
client.Connections.DirectoryProvisioning.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to create its directory provisioning configuration
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.CreateDirectoryProvisioningRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.DirectoryProvisioning.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete the directory provisioning configuration of a connection.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.DirectoryProvisioning.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to delete its directory provisioning configuration
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.DirectoryProvisioning.Update(ID, request) -> *management.UpdateDirectoryProvisioningResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the directory provisioning configuration of a connection.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateDirectoryProvisioningRequestContent{}
client.Connections.DirectoryProvisioning.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to create its directory provisioning configuration
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.UpdateDirectoryProvisioningRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.DirectoryProvisioning.GetDefaultMapping(ID) -> *management.GetDirectoryProvisioningDefaultMappingResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the directory provisioning default attribute mapping of a connection.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.DirectoryProvisioning.GetDefaultMapping(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to retrieve its directory provisioning configuration
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections Keys
<details><summary><code>client.Connections.Keys.Get(ID) -> []*management.ConnectionKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets the connection keys for the Okta or OIDC connection strategy.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.Keys.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the connection
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.Keys.Rotate(ID, request) -> *management.RotateConnectionsKeysResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Rotates the connection keys for the Okta or OIDC connection strategies.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.RotateConnectionKeysRequestContent{}
client.Connections.Keys.Rotate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the connection
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.RotateConnectionKeysRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections SCIMConfiguration
<details><summary><code>client.Connections.SCIMConfiguration.Get(ID) -> *management.GetSCIMConfigurationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a scim configuration by its <code>connectionId</code>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.SCIMConfiguration.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to retrieve its SCIM configuration
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.SCIMConfiguration.Create(ID, request) -> *management.CreateSCIMConfigurationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a scim configuration for a connection.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateSCIMConfigurationRequestContent{}
client.Connections.SCIMConfiguration.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to create its SCIM configuration
    
</dd>
</dl>

<dl>
<dd>

**request:** `*management.CreateSCIMConfigurationRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.SCIMConfiguration.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a scim configuration by its <code>connectionId</code>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.SCIMConfiguration.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to delete its SCIM configuration
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.SCIMConfiguration.Update(ID, request) -> *management.UpdateSCIMConfigurationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a scim configuration by its <code>connectionId</code>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateSCIMConfigurationRequestContent{
        UserIDAttribute: "user_id_attribute",
        Mapping: []*management.SCIMMappingItem{
            &management.SCIMMappingItem{},
        },
    }
client.Connections.SCIMConfiguration.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to update its SCIM configuration
    
</dd>
</dl>

<dl>
<dd>

**userIDAttribute:** `string` — User ID attribute for generating unique user ids
    
</dd>
</dl>

<dl>
<dd>

**mapping:** `[]*management.SCIMMappingItem` — The mapping between auth0 and SCIM
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.SCIMConfiguration.GetDefaultMapping(ID) -> *management.GetSCIMConfigurationDefaultMappingResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a scim configuration's default mapping by its <code>connectionId</code>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.SCIMConfiguration.GetDefaultMapping(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to retrieve its default SCIM mapping
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections Users
<details><summary><code>client.Connections.Users.DeleteByEmail(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a specified connection user by its email (you cannot delete all users from specific connection). Currently, only Database Connections are supported.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteConnectionUsersByEmailQueryParameters{
        Email: "email",
    }
client.Connections.Users.DeleteByEmail(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection (currently only database connections are supported)
    
</dd>
</dl>

<dl>
<dd>

**email:** `string` — The email of the user to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections DirectoryProvisioning Synchronizations
<details><summary><code>client.Connections.DirectoryProvisioning.Synchronizations.Create(ID) -> *management.CreateDirectorySynchronizationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Request an on-demand synchronization of the directory.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.DirectoryProvisioning.Synchronizations.Create(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to trigger synchronization for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Connections SCIMConfiguration Tokens
<details><summary><code>client.Connections.SCIMConfiguration.Tokens.Get(ID) -> management.GetSCIMTokensResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves all scim tokens by its connection <code>id</code>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.SCIMConfiguration.Tokens.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to retrieve its SCIM configuration
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.SCIMConfiguration.Tokens.Create(ID, request) -> *management.CreateSCIMTokenResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a scim token for a scim client.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateSCIMTokenRequestContent{}
client.Connections.SCIMConfiguration.Tokens.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the connection to create its SCIM token
    
</dd>
</dl>

<dl>
<dd>

**scopes:** `[]string` — The scopes of the scim token
    
</dd>
</dl>

<dl>
<dd>

**tokenLifetime:** `*int` — Lifetime of the token in seconds. Must be greater than 900
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Connections.SCIMConfiguration.Tokens.Delete(ID, TokenID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes a scim token by its connection <code>id</code> and <code>tokenId</code>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Connections.SCIMConfiguration.Tokens.Delete(
        context.TODO(),
        "id",
        "tokenId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The connection id that owns the SCIM token to delete
    
</dd>
</dl>

<dl>
<dd>

**tokenID:** `string` — The id of the scim token to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Emails Provider
<details><summary><code>client.Emails.Provider.Get() -> *management.GetEmailProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the <a href="https://auth0.com/docs/customize/email/smtp-email-providers">email provider configuration</a> in your tenant. A list of fields to include or exclude may also be specified.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetEmailProviderRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Emails.Provider.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (dependent upon include_fields) from the result. Leave empty to retrieve `name` and `enabled`. Additional fields available include `credentials`, `default_from_address`, and `settings`.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Emails.Provider.Create(request) -> *management.CreateEmailProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an <a href="https://auth0.com/docs/email/providers">email provider</a>. The <code>credentials</code> object
requires different properties depending on the email provider (which is specified using the <code>name</code> property):
<ul>
  <li><code>mandrill</code> requires <code>api_key</code></li>
  <li><code>sendgrid</code> requires <code>api_key</code></li>
  <li>
    <code>sparkpost</code> requires <code>api_key</code>. Optionally, set <code>region</code> to <code>eu</code> to use
    the SparkPost service hosted in Western Europe; set to <code>null</code> to use the SparkPost service hosted in
    North America. <code>eu</code> or <code>null</code> are the only valid values for <code>region</code>.
  </li>
  <li>
    <code>mailgun</code> requires <code>api_key</code> and <code>domain</code>. Optionally, set <code>region</code> to
    <code>eu</code> to use the Mailgun service hosted in Europe; set to <code>null</code> otherwise. <code>eu</code> or
    <code>null</code> are the only valid values for <code>region</code>.
  </li>
  <li><code>ses</code> requires <code>accessKeyId</code>, <code>secretAccessKey</code>, and <code>region</code></li>
  <li>
    <code>smtp</code> requires <code>smtp_host</code>, <code>smtp_port</code>, <code>smtp_user</code>, and
    <code>smtp_pass</code>
  </li>
</ul>
Depending on the type of provider it is possible to specify <code>settings</code> object with different configuration
options, which will be used when sending an email:
<ul>
  <li>
    <code>smtp</code> provider, <code>settings</code> may contain <code>headers</code> object.
    <ul>
      <li>
        When using AWS SES SMTP host, you may provide a name of configuration set in
        <code>X-SES-Configuration-Set</code> header. Value must be a string.
      </li>
      <li>
        When using Sparkpost host, you may provide value for
        <code>X-MSYS_API</code> header. Value must be an object.
      </li>
    </ul>
  </li>
  <li>
    for <code>ses</code> provider, <code>settings</code> may contain <code>message</code> object, where you can provide
    a name of configuration set in <code>configuration_set_name</code> property. Value must be a string.
  </li>
</ul>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateEmailProviderRequestContent{
        Name: management.EmailProviderNameEnumMailgun,
        Credentials: &management.EmailProviderCredentialsSchema{
            EmailProviderCredentialsSchemaZero: &management.EmailProviderCredentialsSchemaZero{
                APIKey: "api_key",
            },
        },
    }
client.Emails.Provider.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*management.EmailProviderNameEnum` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the provider is enabled (true) or disabled (false).
    
</dd>
</dl>

<dl>
<dd>

**defaultFromAddress:** `*string` — Email address to use as "from" when no other address specified.
    
</dd>
</dl>

<dl>
<dd>

**credentials:** `*management.EmailProviderCredentialsSchema` 
    
</dd>
</dl>

<dl>
<dd>

**settings:** `*management.EmailSpecificProviderSettingsWithAdditionalProperties` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Emails.Provider.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete the email provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Emails.Provider.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Emails.Provider.Update(request) -> *management.UpdateEmailProviderResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update an <a href="https://auth0.com/docs/email/providers">email provider</a>. The <code>credentials</code> object
requires different properties depending on the email provider (which is specified using the <code>name</code> property):
<ul>
  <li><code>mandrill</code> requires <code>api_key</code></li>
  <li><code>sendgrid</code> requires <code>api_key</code></li>
  <li>
    <code>sparkpost</code> requires <code>api_key</code>. Optionally, set <code>region</code> to <code>eu</code> to use
    the SparkPost service hosted in Western Europe; set to <code>null</code> to use the SparkPost service hosted in
    North America. <code>eu</code> or <code>null</code> are the only valid values for <code>region</code>.
  </li>
  <li>
    <code>mailgun</code> requires <code>api_key</code> and <code>domain</code>. Optionally, set <code>region</code> to
    <code>eu</code> to use the Mailgun service hosted in Europe; set to <code>null</code> otherwise. <code>eu</code> or
    <code>null</code> are the only valid values for <code>region</code>.
  </li>
  <li><code>ses</code> requires <code>accessKeyId</code>, <code>secretAccessKey</code>, and <code>region</code></li>
  <li>
    <code>smtp</code> requires <code>smtp_host</code>, <code>smtp_port</code>, <code>smtp_user</code>, and
    <code>smtp_pass</code>
  </li>
</ul>
Depending on the type of provider it is possible to specify <code>settings</code> object with different configuration
options, which will be used when sending an email:
<ul>
  <li>
    <code>smtp</code> provider, <code>settings</code> may contain <code>headers</code> object.
    <ul>
      <li>
        When using AWS SES SMTP host, you may provide a name of configuration set in
        <code>X-SES-Configuration-Set</code> header. Value must be a string.
      </li>
      <li>
        When using Sparkpost host, you may provide value for
        <code>X-MSYS_API</code> header. Value must be an object.
      </li>
    </ul>
    for <code>ses</code> provider, <code>settings</code> may contain <code>message</code> object, where you can provide
    a name of configuration set in <code>configuration_set_name</code> property. Value must be a string.
  </li>
</ul>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateEmailProviderRequestContent{}
client.Emails.Provider.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*management.EmailProviderNameEnum` 
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `*bool` — Whether the provider is enabled (true) or disabled (false).
    
</dd>
</dl>

<dl>
<dd>

**defaultFromAddress:** `*string` — Email address to use as "from" when no other address specified.
    
</dd>
</dl>

<dl>
<dd>

**credentials:** `*management.EmailProviderCredentialsSchema` 
    
</dd>
</dl>

<dl>
<dd>

**settings:** `*management.EmailSpecificProviderSettingsWithAdditionalProperties` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## EventStreams Deliveries
<details><summary><code>client.EventStreams.Deliveries.List(ID) -> []*management.EventStreamDelivery</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListEventStreamDeliveriesRequestParameters{
        Statuses: management.String(
            "statuses",
        ),
        EventTypes: management.String(
            "event_types",
        ),
        DateFrom: management.String(
            "date_from",
        ),
        DateTo: management.String(
            "date_to",
        ),
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.EventStreams.Deliveries.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>

<dl>
<dd>

**statuses:** `*string` — Comma-separated list of statuses by which to filter
    
</dd>
</dl>

<dl>
<dd>

**eventTypes:** `*string` — Comma-separated list of event types by which to filter
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*string` — An RFC-3339 date-time for redelivery start, inclusive. Does not allow sub-second precision.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*string` — An RFC-3339 date-time for redelivery end, exclusive. Does not allow sub-second precision.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EventStreams.Deliveries.GetHistory(ID, EventID) -> *management.GetEventStreamDeliveryHistoryResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.EventStreams.Deliveries.GetHistory(
        context.TODO(),
        "id",
        "event_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>

<dl>
<dd>

**eventID:** `string` — Unique identifier for the event
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## EventStreams Redeliveries
<details><summary><code>client.EventStreams.Redeliveries.Create(ID, request) -> *management.CreateEventStreamRedeliveryResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateEventStreamRedeliveryRequestContent{}
client.EventStreams.Redeliveries.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>

<dl>
<dd>

**dateFrom:** `*time.Time` — An RFC-3339 date-time for redelivery start, inclusive. Does not allow sub-second precision.
    
</dd>
</dl>

<dl>
<dd>

**dateTo:** `*time.Time` — An RFC-3339 date-time for redelivery end, exclusive. Does not allow sub-second precision.
    
</dd>
</dl>

<dl>
<dd>

**statuses:** `[]management.EventStreamDeliveryStatusEnum` — Filter by status
    
</dd>
</dl>

<dl>
<dd>

**eventTypes:** `[]*management.EventStreamEventTypeEnum` — Filter by event type
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EventStreams.Redeliveries.CreateByID(ID, EventID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.EventStreams.Redeliveries.CreateByID(
        context.TODO(),
        "id",
        "event_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Unique identifier for the event stream.
    
</dd>
</dl>

<dl>
<dd>

**eventID:** `string` — Unique identifier for the event
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Flows Executions
<details><summary><code>client.Flows.Executions.List(FlowID) -> *management.ListFlowExecutionsPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ExecutionsListRequest{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.Flows.Executions.List(
        context.TODO(),
        "flow_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**flowID:** `string` — Flow id
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Executions.Get(FlowID, ExecutionID) -> *management.GetFlowExecutionResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ExecutionsGetRequest{}
client.Flows.Executions.Get(
        context.TODO(),
        "flow_id",
        "execution_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**flowID:** `string` — Flow id
    
</dd>
</dl>

<dl>
<dd>

**executionID:** `string` — Flow execution id
    
</dd>
</dl>

<dl>
<dd>

**hydrate:** `*string` — Hydration param
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Executions.Delete(FlowID, ExecutionID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Flows.Executions.Delete(
        context.TODO(),
        "flow_id",
        "execution_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**flowID:** `string` — Flows id
    
</dd>
</dl>

<dl>
<dd>

**executionID:** `string` — Flow execution identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Flows Vault Connections
<details><summary><code>client.Flows.Vault.Connections.List() -> *management.ListFlowsVaultConnectionsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListFlowsVaultConnectionsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Flows.Vault.Connections.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Vault.Connections.Create(request) -> *management.CreateFlowsVaultConnectionResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateFlowsVaultConnectionRequestContent{
        CreateFlowsVaultConnectionActivecampaign: &management.CreateFlowsVaultConnectionActivecampaign{
            CreateFlowsVaultConnectionActivecampaignAPIKey: &management.CreateFlowsVaultConnectionActivecampaignAPIKey{
                Name: "name",
                AppID: management.FlowsVaultConnectionAppIDActivecampaignEnum(
                    "ACTIVECAMPAIGN",
                ),
                Setup: &management.FlowsVaultConnectioSetupAPIKeyWithBaseURL{
                    Type: management.FlowsVaultConnectioSetupTypeAPIKeyEnum(
                        "API_KEY",
                    ),
                    APIKey: "api_key",
                    BaseURL: "base_url",
                },
            },
        },
    }
client.Flows.Vault.Connections.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*management.CreateFlowsVaultConnectionRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Vault.Connections.Get(ID) -> *management.GetFlowsVaultConnectionResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Flows.Vault.Connections.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Flows Vault connection ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Vault.Connections.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Flows.Vault.Connections.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Vault connection id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Flows.Vault.Connections.Update(ID, request) -> *management.UpdateFlowsVaultConnectionResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateFlowsVaultConnectionRequestContent{}
client.Flows.Vault.Connections.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Flows Vault connection ID
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — Flows Vault Connection name.
    
</dd>
</dl>

<dl>
<dd>

**setup:** `*management.UpdateFlowsVaultConnectionSetup` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Guardian Enrollments
<details><summary><code>client.Guardian.Enrollments.CreateTicket(request) -> *management.CreateGuardianEnrollmentTicketResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a <a href="https://auth0.com/docs/secure/multi-factor-authentication/auth0-guardian/create-custom-enrollment-tickets">multi-factor authentication (MFA) enrollment ticket</a>, and optionally send an email with the created ticket, to a given user.
Create a <a href="https://auth0.com/docs/secure/multi-factor-authentication/auth0-guardian/create-custom-enrollment-tickets">multi-factor authentication (MFA) enrollment ticket</a>, and optionally send an email with the created ticket to a given user. Enrollment tickets can specify which factor users must enroll with or allow existing MFA users to enroll in additional factors.<br/> 

Note: Users cannot enroll in Email as a factor through custom enrollment tickets. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateGuardianEnrollmentTicketRequestContent{
        UserID: "user_id",
    }
client.Guardian.Enrollments.CreateTicket(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — user_id for the enrollment ticket
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — alternate email to which the enrollment email will be sent. Optional - by default, the email will be sent to the user's default address
    
</dd>
</dl>

<dl>
<dd>

**sendMail:** `*bool` — Send an email to the user to start the enrollment
    
</dd>
</dl>

<dl>
<dd>

**emailLocale:** `*string` — Optional. Specify the locale of the enrollment email. Used with send_email.
    
</dd>
</dl>

<dl>
<dd>

**factor:** `*management.GuardianEnrollmentFactorEnum` 
    
</dd>
</dl>

<dl>
<dd>

**allowMultipleEnrollments:** `*bool` — Optional. Allows a user who has previously enrolled in MFA to enroll with additional factors.<br />Note: Parameter can only be used with Universal Login; it cannot be used with Classic Login or custom MFA pages.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Enrollments.Get(ID) -> *management.GetGuardianEnrollmentResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details, such as status and type, for a specific multi-factor authentication enrollment registered to a user account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Enrollments.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the enrollment to be retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Enrollments.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a specific multi-factor authentication (MFA) enrollment from a user's account. This allows the user to re-enroll with MFA. For more information, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/reset-user-mfa">Reset User Multi-Factor Authentication and Recovery Codes</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Enrollments.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the enrollment to be deleted.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Guardian Factors
<details><summary><code>client.Guardian.Factors.List() -> []*management.GuardianFactor</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of all <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors">multi-factor authentication factors</a> associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Set(Name, request) -> *management.SetGuardianFactorResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the status (i.e., enabled or disabled) of a specific multi-factor authentication factor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorRequestContent{
        Enabled: true,
    }
client.Guardian.Factors.Set(
        context.TODO(),
        management.GuardianFactorNameEnumPushNotification.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `*management.GuardianFactorNameEnum` — Factor name. Can be `sms`, `push-notification`, `email`, `duo` `otp` `webauthn-roaming`, `webauthn-platform`, or `recovery-code`.
    
</dd>
</dl>

<dl>
<dd>

**enabled:** `bool` — Whether this factor is enabled (true) or disabled (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Guardian Policies
<details><summary><code>client.Guardian.Policies.List() -> management.ListGuardianPoliciesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the <a href="https://auth0.com/docs/secure/multi-factor-authentication/enable-mfa">multi-factor authentication (MFA) policies</a> configured for your tenant.

The following policies are supported:
<ul>
<li><code>all-applications</code> policy prompts with MFA for all logins.</li>
<li><code>confidence-score</code> policy prompts with MFA only for low confidence logins.</li>
</ul>

<b>Note</b>: The <code>confidence-score</code> policy is part of the <a href="https://auth0.com/docs/secure/multi-factor-authentication/adaptive-mfa">Adaptive MFA feature</a>. Adaptive MFA requires an add-on for the Enterprise plan; review <a href="https://auth0.com/pricing">Auth0 Pricing</a> for more details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Policies.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Policies.Set(request) -> management.SetGuardianPoliciesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set <a href="https://auth0.com/docs/secure/multi-factor-authentication/enable-mfa">multi-factor authentication (MFA) policies</a> for your tenant.

The following policies are supported:
<ul>
<li><code>all-applications</code> policy prompts with MFA for all logins.</li>
<li><code>confidence-score</code> policy prompts with MFA only for low confidence logins.</li>
</ul>

<b>Note</b>: The <code>confidence-score</code> policy is part of the <a href="https://auth0.com/docs/secure/multi-factor-authentication/adaptive-mfa">Adaptive MFA feature</a>. Adaptive MFA requires an add-on for the Enterprise plan; review <a href="https://auth0.com/pricing">Auth0 Pricing</a> for more details.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []management.MfaPolicyEnum{
        management.MfaPolicyEnumAllApplications,
    }
client.Guardian.Policies.Set(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `management.SetGuardianPoliciesRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Guardian Factors Phone
<details><summary><code>client.Guardian.Factors.Phone.GetMessageTypes() -> *management.GetGuardianFactorPhoneMessageTypesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve list of <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors/configure-sms-voice-notifications-mfa">phone-type MFA factors</a> (i.e., sms and voice) that are enabled for your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Phone.GetMessageTypes(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Phone.SetMessageTypes(request) -> *management.SetGuardianFactorPhoneMessageTypesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace the list of <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors/configure-sms-voice-notifications-mfa">phone-type MFA factors</a> (i.e., sms and voice) that are enabled for your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorPhoneMessageTypesRequestContent{
        MessageTypes: []management.GuardianFactorPhoneFactorMessageTypeEnum{
            management.GuardianFactorPhoneFactorMessageTypeEnumSms,
        },
    }
client.Guardian.Factors.Phone.SetMessageTypes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**messageTypes:** `[]*management.GuardianFactorPhoneFactorMessageTypeEnum` — The list of phone factors to enable on the tenant. Can include `sms` and `voice`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Phone.GetTwilioProvider() -> *management.GetGuardianFactorsProviderPhoneTwilioResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve configuration details for a Twilio phone provider that has been set up in your tenant. To learn more, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors/configure-sms-voice-notifications-mfa">Configure SMS and Voice Notifications for MFA</a>. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Phone.GetTwilioProvider(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Phone.SetTwilioProvider(request) -> *management.SetGuardianFactorsProviderPhoneTwilioResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the configuration of a Twilio phone provider that has been set up in your tenant. To learn more, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors/configure-sms-voice-notifications-mfa">Configure SMS and Voice Notifications for MFA</a>. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderPhoneTwilioRequestContent{}
client.Guardian.Factors.Phone.SetTwilioProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — From number
    
</dd>
</dl>

<dl>
<dd>

**messagingServiceSid:** `*string` — Copilot SID
    
</dd>
</dl>

<dl>
<dd>

**authToken:** `*string` — Twilio Authentication token
    
</dd>
</dl>

<dl>
<dd>

**sid:** `*string` — Twilio SID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Phone.GetSelectedProvider() -> *management.GetGuardianFactorsProviderPhoneResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the multi-factor authentication phone provider configured for your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Phone.GetSelectedProvider(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Phone.SetProvider(request) -> *management.SetGuardianFactorsProviderPhoneResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderPhoneRequestContent{
        Provider: management.GuardianFactorsProviderSmsProviderEnumAuth0,
    }
client.Guardian.Factors.Phone.SetProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*management.GuardianFactorsProviderSmsProviderEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Phone.GetTemplates() -> *management.GetGuardianFactorPhoneTemplatesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the multi-factor authentication enrollment and verification templates for phone-type factors available in your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Phone.GetTemplates(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Phone.SetTemplates(request) -> *management.SetGuardianFactorPhoneTemplatesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Customize the messages sent to complete phone enrollment and verification (subscription required).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorPhoneTemplatesRequestContent{
        EnrollmentMessage: "enrollment_message",
        VerificationMessage: "verification_message",
    }
client.Guardian.Factors.Phone.SetTemplates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentMessage:** `string` — Message sent to the user when they are invited to enroll with a phone number.
    
</dd>
</dl>

<dl>
<dd>

**verificationMessage:** `string` — Message sent to the user when they are prompted to verify their account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Guardian Factors PushNotification
<details><summary><code>client.Guardian.Factors.PushNotification.GetApnsProvider() -> *management.GetGuardianFactorsProviderApnsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve configuration details for the multi-factor authentication APNS provider associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.PushNotification.GetApnsProvider(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.SetApnsProvider(request) -> *management.SetGuardianFactorsProviderPushNotificationApnsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Overwrite all configuration details of the multi-factor authentication APNS provider associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderPushNotificationApnsRequestContent{}
client.Guardian.Factors.PushNotification.SetApnsProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandbox:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**bundleID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**p12:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.UpdateApnsProvider(request) -> *management.UpdateGuardianFactorsProviderPushNotificationApnsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify configuration details of the multi-factor authentication APNS provider associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateGuardianFactorsProviderPushNotificationApnsRequestContent{}
client.Guardian.Factors.PushNotification.UpdateApnsProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sandbox:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**bundleID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**p12:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.SetFcmProvider(request) -> management.SetGuardianFactorsProviderPushNotificationFcmResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Overwrite all configuration details of the multi-factor authentication FCM provider associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderPushNotificationFcmRequestContent{}
client.Guardian.Factors.PushNotification.SetFcmProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serverKey:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.UpdateFcmProvider(request) -> management.UpdateGuardianFactorsProviderPushNotificationFcmResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify configuration details of the multi-factor authentication FCM provider associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateGuardianFactorsProviderPushNotificationFcmRequestContent{}
client.Guardian.Factors.PushNotification.UpdateFcmProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serverKey:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.SetFcmv1Provider(request) -> management.SetGuardianFactorsProviderPushNotificationFcmv1ResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Overwrite all configuration details of the multi-factor authentication FCMV1 provider associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderPushNotificationFcmv1RequestContent{}
client.Guardian.Factors.PushNotification.SetFcmv1Provider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serverCredentials:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.UpdateFcmv1Provider(request) -> management.UpdateGuardianFactorsProviderPushNotificationFcmv1ResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify configuration details of the multi-factor authentication FCMV1 provider associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateGuardianFactorsProviderPushNotificationFcmv1RequestContent{}
client.Guardian.Factors.PushNotification.UpdateFcmv1Provider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**serverCredentials:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.GetSnsProvider() -> *management.GetGuardianFactorsProviderSnsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve configuration details for an AWS SNS push notification provider that has been enabled for MFA. To learn more, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors/configure-push-notifications-for-mfa">Configure Push Notifications for MFA</a>. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.PushNotification.GetSnsProvider(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.SetSnsProvider(request) -> *management.SetGuardianFactorsProviderPushNotificationSnsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Configure the <a href="https://auth0.com/docs/multifactor-authentication/developer/sns-configuration">AWS SNS push notification provider configuration</a> (subscription required).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderPushNotificationSnsRequestContent{}
client.Guardian.Factors.PushNotification.SetSnsProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**awsAccessKeyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**awsSecretAccessKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**awsRegion:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**snsApnsPlatformApplicationArn:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**snsGcmPlatformApplicationArn:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.UpdateSnsProvider(request) -> *management.UpdateGuardianFactorsProviderPushNotificationSnsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Configure the <a href="https://auth0.com/docs/multifactor-authentication/developer/sns-configuration">AWS SNS push notification provider configuration</a> (subscription required).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateGuardianFactorsProviderPushNotificationSnsRequestContent{}
client.Guardian.Factors.PushNotification.UpdateSnsProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**awsAccessKeyID:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**awsSecretAccessKey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**awsRegion:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**snsApnsPlatformApplicationArn:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**snsGcmPlatformApplicationArn:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.GetSelectedProvider() -> *management.GetGuardianFactorsProviderPushNotificationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify the push notification provider configured for your tenant. For more information, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors/configure-push-notifications-for-mfa">Configure Push Notifications for MFA</a>. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.PushNotification.GetSelectedProvider(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.PushNotification.SetProvider(request) -> *management.SetGuardianFactorsProviderPushNotificationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify the push notification provider configured for your tenant. For more information, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors/configure-push-notifications-for-mfa">Configure Push Notifications for MFA</a>. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderPushNotificationRequestContent{
        Provider: management.GuardianFactorsProviderPushNotificationProviderDataEnumGuardian,
    }
client.Guardian.Factors.PushNotification.SetProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*management.GuardianFactorsProviderPushNotificationProviderDataEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Guardian Factors Sms
<details><summary><code>client.Guardian.Factors.Sms.GetTwilioProvider() -> *management.GetGuardianFactorsProviderSmsTwilioResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the <a href="https://auth0.com/docs/multifactor-authentication/twilio-configuration">Twilio SMS provider configuration</a> (subscription required).

    A new endpoint is available to retrieve the Twilio configuration related to phone factors (<a href='https://auth0.com/docs/api/management/v2/#!/Guardian/get_twilio'>phone Twilio configuration</a>). It has the same payload as this one. Please use it instead.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Sms.GetTwilioProvider(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Sms.SetTwilioProvider(request) -> *management.SetGuardianFactorsProviderSmsTwilioResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint has been deprecated. To complete this action, use the <a href="https://auth0.com/docs/api/management/v2/guardian/put-twilio">Update Twilio phone configuration</a> endpoint.

    <b>Previous functionality</b>: Update the Twilio SMS provider configuration.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderSmsTwilioRequestContent{}
client.Guardian.Factors.Sms.SetTwilioProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — From number
    
</dd>
</dl>

<dl>
<dd>

**messagingServiceSid:** `*string` — Copilot SID
    
</dd>
</dl>

<dl>
<dd>

**authToken:** `*string` — Twilio Authentication token
    
</dd>
</dl>

<dl>
<dd>

**sid:** `*string` — Twilio SID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Sms.GetSelectedProvider() -> *management.GetGuardianFactorsProviderSmsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint has been deprecated. To complete this action, use the <a href="https://auth0.com/docs/api/management/v2/guardian/get-phone-providers">Retrieve phone configuration</a> endpoint instead.

    <b>Previous functionality</b>: Retrieve details for the multi-factor authentication SMS provider configured for your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Sms.GetSelectedProvider(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Sms.SetProvider(request) -> *management.SetGuardianFactorsProviderSmsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint has been deprecated. To complete this action, use the <a href="https://auth0.com/docs/api/management/v2/guardian/put-phone-providers">Update phone configuration</a> endpoint instead.

    <b>Previous functionality</b>: Update the multi-factor authentication SMS provider configuration in your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorsProviderSmsRequestContent{
        Provider: management.GuardianFactorsProviderSmsProviderEnumAuth0,
    }
client.Guardian.Factors.Sms.SetProvider(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**provider:** `*management.GuardianFactorsProviderSmsProviderEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Sms.GetTemplates() -> *management.GetGuardianFactorSmsTemplatesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint has been deprecated. To complete this action, use the <a href="https://auth0.com/docs/api/management/v2/guardian/get-factor-phone-templates">Retrieve enrollment and verification phone templates</a> endpoint instead.

    <b>Previous function</b>: Retrieve details of SMS enrollment and verification templates configured for your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Sms.GetTemplates(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Sms.SetTemplates(request) -> *management.SetGuardianFactorSmsTemplatesResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

This endpoint has been deprecated. To complete this action, use the <a href="https://auth0.com/docs/api/management/v2/guardian/put-factor-phone-templates">Update enrollment and verification phone templates</a> endpoint instead.

    <b>Previous functionality</b>: Customize the messages sent to complete SMS enrollment and verification.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorSmsTemplatesRequestContent{
        EnrollmentMessage: "enrollment_message",
        VerificationMessage: "verification_message",
    }
client.Guardian.Factors.Sms.SetTemplates(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enrollmentMessage:** `string` — Message sent to the user when they are invited to enroll with a phone number.
    
</dd>
</dl>

<dl>
<dd>

**verificationMessage:** `string` — Message sent to the user when they are prompted to verify their account.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Guardian Factors Duo Settings
<details><summary><code>client.Guardian.Factors.Duo.Settings.Get() -> *management.GetGuardianFactorDuoSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the DUO account and factor configuration.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Guardian.Factors.Duo.Settings.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Duo.Settings.Set(request) -> *management.SetGuardianFactorDuoSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set the DUO account configuration and other properties specific to this factor.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetGuardianFactorDuoSettingsRequestContent{}
client.Guardian.Factors.Duo.Settings.Set(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ikey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**skey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**host:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Guardian.Factors.Duo.Settings.Update(request) -> *management.UpdateGuardianFactorDuoSettingsResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateGuardianFactorDuoSettingsRequestContent{}
client.Guardian.Factors.Duo.Settings.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ikey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**skey:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**host:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hooks Secrets
<details><summary><code>client.Hooks.Secrets.Get(ID) -> management.GetHookSecretResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a hook's secrets by the ID of the hook. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hooks.Secrets.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the hook to retrieve secrets from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.Secrets.Create(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add one or more secrets to an existing hook. Accepts an object of key-value pairs, where the key is the name of the secret. A hook can have a maximum of 20 secrets. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]string{
        "key": "value",
    }
client.Hooks.Secrets.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the hook to retrieve
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.CreateHookSecretRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.Secrets.Delete(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete one or more existing secrets for a given hook. Accepts an array of secret names to delete. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []string{
        "string",
    }
client.Hooks.Secrets.Delete(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the hook whose secrets to delete.
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.DeleteHookSecretRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hooks.Secrets.Update(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update one or more existing secrets for an existing hook. Accepts an object of key-value pairs, where the key is the name of the existing secret. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]string{
        "key": "value",
    }
client.Hooks.Secrets.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the hook whose secrets to update.
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.UpdateHookSecretRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Jobs UsersExports
<details><summary><code>client.Jobs.UsersExports.Create(request) -> *management.CreateExportUsersResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Export all users to a file via a long-running job.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateExportUsersRequestContent{}
client.Jobs.UsersExports.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**connectionID:** `*string` — connection_id of the connection from which users will be exported.
    
</dd>
</dl>

<dl>
<dd>

**format:** `*management.JobFileFormatEnum` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Limit the number of records.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `[]*management.CreateExportUsersFields` — List of fields to be included in the CSV. Defaults to a predefined set of fields.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Jobs UsersImports
<details><summary><code>client.Jobs.UsersImports.Create(request) -> *management.CreateImportUsersResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Import users from a <a href="https://auth0.com/docs/users/references/bulk-import-database-schema-examples">formatted file</a> into a connection via a long-running job. When importing users, with or without upsert, the `email_verified` is set to `false` when the email address is added or updated. Users must verify their email address. To avoid this behavior, set `email_verified` to `true` in the imported data.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateImportUsersRequestContent{
        Users: strings.NewReader(
            "",
        ),
        ConnectionID: "connection_id",
    }
client.Jobs.UsersImports.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Jobs VerificationEmail
<details><summary><code>client.Jobs.VerificationEmail.Create(request) -> *management.CreateVerificationEmailResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Send an email to the specified user that asks them to click a link to <a href="https://auth0.com/docs/email/custom#verification-email">verify their email address</a>.

Note: You must have the `Status` toggle enabled for the verification email template for the email to be sent.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateVerificationEmailRequestContent{
        UserID: "user_id",
    }
client.Jobs.VerificationEmail.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — user_id of the user to send the verification email to.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — client_id of the client (application). If no value provided, the global Client ID will be used.
    
</dd>
</dl>

<dl>
<dd>

**identity:** `*management.Identity` 
    
</dd>
</dl>

<dl>
<dd>

**organizationID:** `*string` — (Optional) Organization ID – the ID of the Organization. If provided, organization parameters will be made available to the email template and organization branding will be applied to the prompt. In addition, the redirect link in the prompt will include organization_id and organization_name query string parameters.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Jobs Errors
<details><summary><code>client.Jobs.Errors.Get(ID) -> *jobs.ErrorsGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve error details of a failed job.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Jobs.Errors.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the job.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Keys CustomSigning
<details><summary><code>client.Keys.CustomSigning.Get() -> *management.GetCustomSigningKeysResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get entire jwks representation of custom signing keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.CustomSigning.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.CustomSigning.Set(request) -> *management.SetCustomSigningKeysResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or replace entire jwks representation of custom signing keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.SetCustomSigningKeysRequestContent{
        Keys: []*management.CustomSigningKeyJwk{
            &management.CustomSigningKeyJwk{
                Kty: management.CustomSigningKeyTypeEnumEc,
            },
        },
    }
client.Keys.CustomSigning.Set(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**keys:** `[]*management.CustomSigningKeyJwk` — An array of custom public signing keys.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.CustomSigning.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete entire jwks representation of custom signing keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.CustomSigning.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Keys Encryption
<details><summary><code>client.Keys.Encryption.List() -> *management.ListEncryptionKeyOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of all the encryption keys associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListEncryptionKeysRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Keys.Encryption.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Default value is 50, maximum value is 100.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Encryption.Create(request) -> *management.CreateEncryptionKeyResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create the new, pre-activated encryption key, without the key material.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateEncryptionKeyRequestContent{
        Type: management.CreateEncryptionKeyTypeCustomerProvidedRootKey,
    }
client.Keys.Encryption.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**type_:** `*management.CreateEncryptionKeyType` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Encryption.Rekey() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Perform rekeying operation on the key hierarchy.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Encryption.Rekey(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Encryption.Get(Kid) -> *management.GetEncryptionKeyResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the encryption key with the given ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Encryption.Get(
        context.TODO(),
        "kid",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**kid:** `string` — Encryption key ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Encryption.Import(Kid, request) -> *management.ImportEncryptionKeyResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Import wrapped key material and activate encryption key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ImportEncryptionKeyRequestContent{
        WrappedKey: "wrapped_key",
    }
client.Keys.Encryption.Import(
        context.TODO(),
        "kid",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**kid:** `string` — Encryption key ID
    
</dd>
</dl>

<dl>
<dd>

**wrappedKey:** `string` — Base64 encoded ciphertext of key material wrapped by public wrapping key.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Encryption.Delete(Kid) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete the custom provided encryption key with the given ID and move back to using native encryption key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Encryption.Delete(
        context.TODO(),
        "kid",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**kid:** `string` — Encryption key ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Encryption.CreatePublicWrappingKey(Kid) -> *management.CreateEncryptionKeyPublicWrappingResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create the public wrapping key to wrap your own encryption key material.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Encryption.CreatePublicWrappingKey(
        context.TODO(),
        "kid",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**kid:** `string` — Encryption key ID
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Keys Signing
<details><summary><code>client.Keys.Signing.List() -> []*management.SigningKeys</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of all the application signing keys associated with your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Signing.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Signing.Rotate() -> *management.RotateSigningKeysResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Rotate the application signing key of your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Signing.Rotate(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Signing.Get(Kid) -> *management.GetSigningKeysResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details of the application signing key with the given ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Signing.Get(
        context.TODO(),
        "kid",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**kid:** `string` — Key id of the key to retrieve
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Keys.Signing.Revoke(Kid) -> *management.RevokedSigningKeysResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revoke the application signing key with the given ID.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Keys.Signing.Revoke(
        context.TODO(),
        "kid",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**kid:** `string` — Key id of the key to revoke
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations ClientGrants
<details><summary><code>client.Organizations.ClientGrants.List(ID) -> *management.ListOrganizationClientGrantsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationClientGrantsRequestParameters{
        Audience: management.String(
            "audience",
        ),
        ClientID: management.String(
            "client_id",
        ),
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Organizations.ClientGrants.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**audience:** `*string` — Optional filter on audience of the client grant.
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `*string` — Optional filter on client_id of the client grant.
    
</dd>
</dl>

<dl>
<dd>

**grantIDs:** `*string` — Optional filter on the ID of the client grant. Must be URL encoded and may be specified multiple times (max 10).<br /><b>e.g.</b> <i>../client-grants?grant_ids=id1&grant_ids=id2</i>
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ClientGrants.Create(ID, request) -> *management.AssociateOrganizationClientGrantResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.AssociateOrganizationClientGrantRequestContent{
        GrantID: "grant_id",
    }
client.Organizations.ClientGrants.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**grantID:** `string` — A Client Grant ID to add to the organization.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.ClientGrants.Delete(ID, GrantID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.ClientGrants.Delete(
        context.TODO(),
        "id",
        "grant_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**grantID:** `string` — The Client Grant ID to remove from the organization
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations DiscoveryDomains
<details><summary><code>client.Organizations.DiscoveryDomains.List(ID) -> *management.ListOrganizationDiscoveryDomainsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve list of all organization discovery domains associated with the specified organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationDiscoveryDomainsRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.Organizations.DiscoveryDomains.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the organization.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.DiscoveryDomains.Create(ID, request) -> *management.CreateOrganizationDiscoveryDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the verification status and/or use_for_organization_discovery for an organization discovery domain. The <code>status</code> field must be either <code>pending</code> or <code>verified</code>. The <code>use_for_organization_discovery</code> field can be <code>true</code> or <code>false</code> (default: <code>true</code>).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateOrganizationDiscoveryDomainRequestContent{
        Domain: "domain",
    }
client.Organizations.DiscoveryDomains.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the organization.
    
</dd>
</dl>

<dl>
<dd>

**domain:** `string` — The domain name to associate with the organization e.g. acme.com.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*management.OrganizationDiscoveryDomainStatus` 
    
</dd>
</dl>

<dl>
<dd>

**useForOrganizationDiscovery:** `*bool` — Indicates whether this discovery domain should be used for organization discovery.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.DiscoveryDomains.Get(ID, DiscoveryDomainID) -> *management.GetOrganizationDiscoveryDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a single organization discovery domain specified by ID. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.DiscoveryDomains.Get(
        context.TODO(),
        "id",
        "discovery_domain_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the organization.
    
</dd>
</dl>

<dl>
<dd>

**discoveryDomainID:** `string` — ID of the discovery domain.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.DiscoveryDomains.Delete(ID, DiscoveryDomainID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a discovery domain from an organization. This action cannot be undone. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.DiscoveryDomains.Delete(
        context.TODO(),
        "id",
        "discovery_domain_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the organization.
    
</dd>
</dl>

<dl>
<dd>

**discoveryDomainID:** `string` — ID of the discovery domain.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.DiscoveryDomains.Update(ID, DiscoveryDomainID, request) -> *management.UpdateOrganizationDiscoveryDomainResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update the verification status and/or use_for_organization_discovery for an organization discovery domain. The <code>status</code> field must be either <code>pending</code> or <code>verified</code>. The <code>use_for_organization_discovery</code> field can be <code>true</code> or <code>false</code> (default: <code>true</code>).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateOrganizationDiscoveryDomainRequestContent{}
client.Organizations.DiscoveryDomains.Update(
        context.TODO(),
        "id",
        "discovery_domain_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the organization.
    
</dd>
</dl>

<dl>
<dd>

**discoveryDomainID:** `string` — ID of the discovery domain to update.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*management.OrganizationDiscoveryDomainStatus` 
    
</dd>
</dl>

<dl>
<dd>

**useForOrganizationDiscovery:** `*bool` — Indicates whether this discovery domain should be used for organization discovery.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations EnabledConnections
<details><summary><code>client.Organizations.EnabledConnections.List(ID) -> *management.ListOrganizationConnectionsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a specific connection currently enabled for an Organization. Information returned includes details such as connection ID, name, strategy, and whether the connection automatically grants membership upon login.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationConnectionsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Organizations.EnabledConnections.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.EnabledConnections.Add(ID, request) -> *management.AddOrganizationConnectionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Enable a specific connection for a given Organization. To enable a connection, it must already exist within your tenant; connections cannot be created through this action.

<a href="https://auth0.com/docs/authenticate/identity-providers">Connections</a> represent the relationship between Auth0 and a source of users. Available types of connections include database, enterprise, and social.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.AddOrganizationConnectionRequestContent{
        ConnectionID: "connection_id",
    }
client.Organizations.EnabledConnections.Add(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `string` — Single connection ID to add to the organization.
    
</dd>
</dl>

<dl>
<dd>

**assignMembershipOnLogin:** `*bool` — When true, all users that log in with this connection will be automatically granted membership in the organization. When false, users must be granted membership in the organization before logging in with this connection.
    
</dd>
</dl>

<dl>
<dd>

**isSignupEnabled:** `*bool` — Determines whether organization signup should be enabled for this organization connection. Only applicable for database connections. Default: false.
    
</dd>
</dl>

<dl>
<dd>

**showAsButton:** `*bool` — Determines whether a connection should be displayed on this organization’s login prompt. Only applicable for enterprise connections. Default: true.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.EnabledConnections.Get(ID, ConnectionID) -> *management.GetOrganizationConnectionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details about a specific connection currently enabled for an Organization. Information returned includes details such as connection ID, name, strategy, and whether the connection automatically grants membership upon login.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.EnabledConnections.Get(
        context.TODO(),
        "id",
        "connectionId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `string` — Connection identifier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.EnabledConnections.Delete(ID, ConnectionID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Disable a specific connection for an Organization. Once disabled, Organization members can no longer use that connection to authenticate. 

<b>Note</b>: This action does not remove the connection from your tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.EnabledConnections.Delete(
        context.TODO(),
        "id",
        "connectionId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `string` — Connection identifier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.EnabledConnections.Update(ID, ConnectionID, request) -> *management.UpdateOrganizationConnectionResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify the details of a specific connection currently enabled for an Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateOrganizationConnectionRequestContent{}
client.Organizations.EnabledConnections.Update(
        context.TODO(),
        "id",
        "connectionId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `string` — Connection identifier.
    
</dd>
</dl>

<dl>
<dd>

**assignMembershipOnLogin:** `*bool` — When true, all users that log in with this connection will be automatically granted membership in the organization. When false, users must be granted membership in the organization before logging in with this connection.
    
</dd>
</dl>

<dl>
<dd>

**isSignupEnabled:** `*bool` — Determines whether organization signup should be enabled for this organization connection. Only applicable for database connections. Default: false.
    
</dd>
</dl>

<dl>
<dd>

**showAsButton:** `*bool` — Determines whether a connection should be displayed on this organization’s login prompt. Only applicable for enterprise connections. Default: true.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations Invitations
<details><summary><code>client.Organizations.Invitations.List(ID) -> *management.ListOrganizationInvitationsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a detailed list of invitations sent to users for a specific Organization. The list includes details such as inviter and invitee information, invitation URLs, and dates of creation and expiration. To learn more about Organization invitations, review <a href="https://auth0.com/docs/manage-users/organizations/configure-organizations/invite-members">Invite Organization Members</a>. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationInvitationsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        Sort: management.String(
            "sort",
        ),
    }
client.Organizations.Invitations.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — When true, return results inside an object that also contains the start and limit.  When false (default), a direct array of results is returned.  We do not yet support returning the total invitations count.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false). Defaults to true.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Field to sort by. Use field:order where order is 1 for ascending and -1 for descending Defaults to created_at:-1.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Invitations.Create(ID, request) -> *management.CreateOrganizationInvitationResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a user invitation for a specific Organization. Upon creation, the listed user receives an email inviting them to join the Organization. To learn more about Organization invitations, review <a href="https://auth0.com/docs/manage-users/organizations/configure-organizations/invite-members">Invite Organization Members</a>. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateOrganizationInvitationRequestContent{
        Inviter: &management.OrganizationInvitationInviter{
            Name: "name",
        },
        Invitee: &management.OrganizationInvitationInvitee{
            Email: "email",
        },
        ClientID: "client_id",
    }
client.Organizations.Invitations.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**inviter:** `*management.OrganizationInvitationInviter` 
    
</dd>
</dl>

<dl>
<dd>

**invitee:** `*management.OrganizationInvitationInvitee` 
    
</dd>
</dl>

<dl>
<dd>

**clientID:** `string` — Auth0 client ID. Used to resolve the application's login initiation endpoint.
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `*string` — The id of the connection to force invitee to authenticate with.
    
</dd>
</dl>

<dl>
<dd>

**appMetadata:** `*management.AppMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**userMetadata:** `*management.UserMetadata` 
    
</dd>
</dl>

<dl>
<dd>

**ttlSec:** `*int` — Number of seconds for which the invitation is valid before expiration. If unspecified or set to 0, this value defaults to 604800 seconds (7 days). Max value: 2592000 seconds (30 days).
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]string` — List of roles IDs to associated with the user.
    
</dd>
</dl>

<dl>
<dd>

**sendInvitationEmail:** `*bool` — Whether the user will receive an invitation email (true) or no email (false), true by default
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Invitations.Get(ID, InvitationID) -> *management.GetOrganizationInvitationResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetOrganizationInvitationRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Organizations.Invitations.Get(
        context.TODO(),
        "id",
        "invitation_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` — The id of the user invitation.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false). Defaults to true.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Invitations.Delete(ID, InvitationID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Organizations.Invitations.Delete(
        context.TODO(),
        "id",
        "invitation_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**invitationID:** `string` — The id of the user invitation.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations Members
<details><summary><code>client.Organizations.Members.List(ID) -> *management.ListOrganizationMembersPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List organization members.
This endpoint is subject to eventual consistency. New users may not be immediately included in the response and deleted users may not be immediately removed from it.

<ul>
  <li>
    Use the <code>fields</code> parameter to optionally define the specific member details retrieved. If <code>fields</code> is left blank, all fields (except roles) are returned.
  </li>
  <li>
    Member roles are not sent by default. Use <code>fields=roles</code> to retrieve the roles assigned to each listed member. To use this parameter, you must include the <code>read:organization_member_roles</code> scope in the token.
  </li>
</ul>

This endpoint supports two types of pagination:

- Offset pagination
- Checkpoint pagination

Checkpoint pagination must be used if you need to retrieve more than 1000 organization members.

<h2>Checkpoint Pagination</h2>

To search by checkpoint, use the following parameters: - from: Optional id from which to start selection. - take: The total amount of entries to retrieve when using the from parameter. Defaults to 50. Note: The first time you call this endpoint using Checkpoint Pagination, you should omit the <code>from</code> parameter. If there are more results, a <code>next</code> value will be included in the response. You can use this for subsequent API calls. When <code>next</code> is no longer included in the response, this indicates there are no more pages remaining.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationMembersRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Organizations.Members.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Members.Create(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set one or more existing users as members of a specific <a href="https://auth0.com/docs/manage-users/organizations">Organization</a>.

To add a user to an Organization through this action, the user must already exist in your tenant. If a user does not yet exist, you can <a href="https://auth0.com/docs/manage-users/organizations/configure-organizations/invite-members">invite them to create an account</a>, manually create them through the Auth0 Dashboard, or use the Management API.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateOrganizationMemberRequestContent{
        Members: []string{
            "members",
        },
    }
client.Organizations.Members.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**members:** `[]string` — List of user IDs to add to the organization as members.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Members.Delete(ID, request) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteOrganizationMembersRequestContent{
        Members: []string{
            "members",
        },
    }
client.Organizations.Members.Delete(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**members:** `[]string` — List of user IDs to remove from the organization.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Organizations Members Roles
<details><summary><code>client.Organizations.Members.Roles.List(ID, UserID) -> *management.ListOrganizationMemberRolesOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed list of roles assigned to a given user within the context of a specific Organization. 

Users can be members of multiple Organizations with unique roles assigned for each membership. This action only returns the roles associated with the specified Organization; any roles assigned to the user within other Organizations are not included.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListOrganizationMemberRolesRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Organizations.Members.Roles.List(
        context.TODO(),
        "id",
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — ID of the user to associate roles with.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Members.Roles.Assign(ID, UserID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assign one or more <a href="https://auth0.com/docs/manage-users/access-control/rbac">roles</a> to a user to determine their access for a specific Organization.

Users can be members of multiple Organizations with unique roles assigned for each membership. This action assigns roles to a user only for the specified Organization. Roles cannot be assigned to a user across multiple Organizations in the same call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.AssignOrganizationMemberRolesRequestContent{
        Roles: []string{
            "roles",
        },
    }
client.Organizations.Members.Roles.Assign(
        context.TODO(),
        "id",
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — ID of the user to associate roles with.
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]string` — List of roles IDs to associated with the user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Organizations.Members.Roles.Delete(ID, UserID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove one or more Organization-specific <a href="https://auth0.com/docs/manage-users/access-control/rbac">roles</a> from a given user.

Users can be members of multiple Organizations with unique roles assigned for each membership. This action removes roles from a user in relation to the specified Organization. Roles assigned to the user within a different Organization cannot be managed in the same call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteOrganizationMemberRolesRequestContent{
        Roles: []string{
            "roles",
        },
    }
client.Organizations.Members.Roles.Delete(
        context.TODO(),
        "id",
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Organization identifier.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — User ID of the organization member to remove roles from.
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]string` — List of roles IDs associated with the organization member to remove.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompts Rendering
<details><summary><code>client.Prompts.Rendering.List() -> *management.ListAculsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get render setting configurations for all screens.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListAculsRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
        Prompt: management.String(
            "prompt",
        ),
        Screen: management.String(
            "screen",
        ),
        RenderingMode: management.AculRenderingModeEnumAdvanced.Ptr(),
    }
client.Prompts.Rendering.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (default: true) or excluded (false).
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Maximum value is 100, default value is 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total configuration count (true) or as a direct array of results (false, default).
    
</dd>
</dl>

<dl>
<dd>

**prompt:** `*string` — Name of the prompt to filter by
    
</dd>
</dl>

<dl>
<dd>

**screen:** `*string` — Name of the screen to filter by
    
</dd>
</dl>

<dl>
<dd>

**renderingMode:** `*management.AculRenderingModeEnum` — Rendering mode to filter by
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.Rendering.BulkUpdate(request) -> *management.BulkUpdateAculResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Learn more about <a href='https://auth0.com/docs/customize/login-pages/advanced-customizations/getting-started/configure-acul-screens'>configuring render settings</a> for advanced customization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.BulkUpdateAculRequestContent{
        Configs: []*management.AculConfigsItem{
            &management.AculConfigsItem{
                Prompt: management.PromptGroupNameEnumLogin,
                Screen: management.ScreenGroupNameEnumLogin,
            },
        },
    }
client.Prompts.Rendering.BulkUpdate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**configs:** `management.AculConfigs` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.Rendering.Get(Prompt, Screen) -> *management.GetAculResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get render settings for a screen.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Prompts.Rendering.Get(
        context.TODO(),
        management.PromptGroupNameEnumLogin.Ptr(),
        management.ScreenGroupNameEnumLogin.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**prompt:** `*management.PromptGroupNameEnum` — Name of the prompt
    
</dd>
</dl>

<dl>
<dd>

**screen:** `*management.ScreenGroupNameEnum` — Name of the screen
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.Rendering.Update(Prompt, Screen, request) -> *management.UpdateAculResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Learn more about <a href='https://auth0.com/docs/customize/login-pages/advanced-customizations/getting-started/configure-acul-screens'>configuring render settings</a> for advanced customization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateAculRequestContent{}
client.Prompts.Rendering.Update(
        context.TODO(),
        management.PromptGroupNameEnumLogin.Ptr(),
        management.ScreenGroupNameEnumLogin.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**prompt:** `*management.PromptGroupNameEnum` — Name of the prompt
    
</dd>
</dl>

<dl>
<dd>

**screen:** `*management.ScreenGroupNameEnum` — Name of the screen
    
</dd>
</dl>

<dl>
<dd>

**renderingMode:** `*management.AculRenderingModeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**contextConfiguration:** `*management.AculContextConfiguration` 
    
</dd>
</dl>

<dl>
<dd>

**defaultHeadTagsDisabled:** `*bool` — Override Universal Login default head tags
    
</dd>
</dl>

<dl>
<dd>

**usePageTemplate:** `*bool` — Use page template with ACUL
    
</dd>
</dl>

<dl>
<dd>

**headTags:** `[]*management.AculHeadTag` — An array of head tags
    
</dd>
</dl>

<dl>
<dd>

**filters:** `*management.AculFilters` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompts CustomText
<details><summary><code>client.Prompts.CustomText.Get(Prompt, Language) -> management.GetCustomTextsByLanguageResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve custom text for a specific prompt and language.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Prompts.CustomText.Get(
        context.TODO(),
        management.PromptGroupNameEnumLogin.Ptr(),
        management.PromptLanguageEnumAm.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**prompt:** `*management.PromptGroupNameEnum` — Name of the prompt.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*management.PromptLanguageEnum` — Language to update.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.CustomText.Set(Prompt, Language, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set custom text for a specific prompt. Existing texts will be overwritten.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Prompts.CustomText.Set(
        context.TODO(),
        management.PromptGroupNameEnumLogin.Ptr(),
        management.PromptLanguageEnumAm.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**prompt:** `*management.PromptGroupNameEnum` — Name of the prompt.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*management.PromptLanguageEnum` — Language to update.
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.SetsCustomTextsByLanguageRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Prompts Partials
<details><summary><code>client.Prompts.Partials.Get(Prompt) -> management.GetPartialsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get template partials for a prompt
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Prompts.Partials.Get(
        context.TODO(),
        management.PartialGroupsEnumLogin.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**prompt:** `*management.PartialGroupsEnum` — Name of the prompt.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Prompts.Partials.Set(Prompt, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set template partials for a prompt
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]any{
        "key": "value",
    }
client.Prompts.Partials.Set(
        context.TODO(),
        management.PartialGroupsEnumLogin.Ptr(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**prompt:** `*management.PartialGroupsEnum` — Name of the prompt.
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.SetPartialsRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## RiskAssessments Settings
<details><summary><code>client.RiskAssessments.Settings.Get() -> *management.GetRiskAssessmentsSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets the tenant settings for risk assessments
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.RiskAssessments.Settings.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RiskAssessments.Settings.Update(request) -> *management.UpdateRiskAssessmentsSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the tenant settings for risk assessments
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateRiskAssessmentsSettingsRequestContent{
        Enabled: true,
    }
client.RiskAssessments.Settings.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**enabled:** `bool` — Whether or not risk assessment is enabled.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## RiskAssessments Settings NewDevice
<details><summary><code>client.RiskAssessments.Settings.NewDevice.Get() -> *management.GetRiskAssessmentsSettingsNewDeviceResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets the risk assessment settings for the new device assessor
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.RiskAssessments.Settings.NewDevice.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.RiskAssessments.Settings.NewDevice.Update(request) -> *management.UpdateRiskAssessmentsSettingsNewDeviceResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the risk assessment settings for the new device assessor
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateRiskAssessmentsSettingsNewDeviceRequestContent{
        RememberFor: 1,
    }
client.RiskAssessments.Settings.NewDevice.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**rememberFor:** `int` — Length of time to remember devices for, in days.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Roles Permissions
<details><summary><code>client.Roles.Permissions.List(ID) -> *management.ListRolePermissionsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed list (name, description, resource server) of permissions granted by a specified user role.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListRolePermissionsRequestParameters{
        PerPage: management.Int(
            1,
        ),
        Page: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Roles.Permissions.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to list granted permissions.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Permissions.Add(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Add one or more <a href="https://auth0.com/docs/manage-users/access-control/configure-core-rbac/manage-permissions">permissions</a> to a specified user role.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.AddRolePermissionsRequestContent{
        Permissions: []*management.PermissionRequestPayload{
            &management.PermissionRequestPayload{
                ResourceServerIdentifier: "resource_server_identifier",
                PermissionName: "permission_name",
            },
        },
    }
client.Roles.Permissions.Add(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to add permissions to.
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]*management.PermissionRequestPayload` — array of resource_server_identifier, permission_name pairs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Permissions.Delete(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove one or more <a href="https://auth0.com/docs/manage-users/access-control/configure-core-rbac/manage-permissions">permissions</a> from a specified user role.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteRolePermissionsRequestContent{
        Permissions: []*management.PermissionRequestPayload{
            &management.PermissionRequestPayload{
                ResourceServerIdentifier: "resource_server_identifier",
                PermissionName: "permission_name",
            },
        },
    }
client.Roles.Permissions.Delete(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to remove permissions from.
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]*management.PermissionRequestPayload` — array of resource_server_identifier, permission_name pairs.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Roles Users
<details><summary><code>client.Roles.Users.List(ID) -> *management.ListRoleUsersPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve list of users associated with a specific role. For Dashboard instructions, review <a href="https://auth0.com/docs/manage-users/access-control/configure-core-rbac/roles/view-users-assigned-to-roles">View Users Assigned to Roles</a>.

This endpoint supports two types of pagination:
<ul>
<li>Offset pagination</li>
<li>Checkpoint pagination</li>
</ul>

Checkpoint pagination must be used if you need to retrieve more than 1000 organization members.

<h2>Checkpoint Pagination</h2>

To search by checkpoint, use the following parameters:
<ul>
<li><code>from</code>: Optional id from which to start selection.</li>
<li><code>take</code>: The total amount of entries to retrieve when using the from parameter. Defaults to 50.</li>
</ul>

<b>Note</b>: The first time you call this endpoint using checkpoint pagination, omit the <code>from</code> parameter. If there are more results, a <code>next</code> value is included in the response. You can use this for subsequent API calls. When <code>next</code> is no longer included in the response, no pages are remaining.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListRoleUsersRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.Roles.Users.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to retrieve a list of users associated with.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Roles.Users.Assign(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assign one or more users to an existing user role. To learn more, review <a href="https://auth0.com/docs/manage-users/access-control/rbac">Role-Based Access Control</a>.

<b>Note</b>: New roles cannot be created through this action.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.AssignRoleUsersRequestContent{
        Users: []string{
            "users",
        },
    }
client.Roles.Users.Assign(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the role to assign users to.
    
</dd>
</dl>

<dl>
<dd>

**users:** `[]string` — user_id's of the users to assign the role to.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SelfServiceProfiles CustomText
<details><summary><code>client.SelfServiceProfiles.CustomText.List(ID, Language, Page) -> management.ListSelfServiceProfileCustomTextResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves text customizations for a given self-service profile, language and Self Service SSO Flow page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SelfServiceProfiles.CustomText.List(
        context.TODO(),
        "id",
        management.SelfServiceProfileCustomTextLanguageEnum(
            "en",
        ),
        management.SelfServiceProfileCustomTextPageEnum(
            "get-started",
        ),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the self-service profile.
    
</dd>
</dl>

<dl>
<dd>

**language:** `management.SelfServiceProfileCustomTextLanguageEnum` — The language of the custom text.
    
</dd>
</dl>

<dl>
<dd>

**page:** `management.SelfServiceProfileCustomTextPageEnum` — The page where the custom text is shown.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SelfServiceProfiles.CustomText.Set(ID, Language, Page, request) -> management.SetSelfServiceProfileCustomTextResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates text customizations for a given self-service profile, language and Self Service SSO Flow page.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := map[string]string{
        "key": "value",
    }
client.SelfServiceProfiles.CustomText.Set(
        context.TODO(),
        "id",
        management.SelfServiceProfileCustomTextLanguageEnum(
            "en",
        ),
        management.SelfServiceProfileCustomTextPageEnum(
            "get-started",
        ),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the self-service profile.
    
</dd>
</dl>

<dl>
<dd>

**language:** `management.SelfServiceProfileCustomTextLanguageEnum` — The language of the custom text.
    
</dd>
</dl>

<dl>
<dd>

**page:** `management.SelfServiceProfileCustomTextPageEnum` — The page where the custom text is shown.
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.SetSelfServiceProfileCustomTextRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SelfServiceProfiles SSOTicket
<details><summary><code>client.SelfServiceProfiles.SSOTicket.Create(ID, request) -> *management.CreateSelfServiceProfileSSOTicketResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an SSO access ticket to initiate the Self Service SSO Flow using a self-service profile.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateSelfServiceProfileSSOTicketRequestContent{}
client.SelfServiceProfiles.SSOTicket.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The id of the self-service profile to retrieve
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `*string` — If provided, this will allow editing of the provided connection during the SSO Flow
    
</dd>
</dl>

<dl>
<dd>

**connectionConfig:** `*management.SelfServiceProfileSSOTicketConnectionConfig` 
    
</dd>
</dl>

<dl>
<dd>

**enabledClients:** `[]string` — List of client_ids that the connection will be enabled for.
    
</dd>
</dl>

<dl>
<dd>

**enabledOrganizations:** `[]*management.SelfServiceProfileSSOTicketEnabledOrganization` — List of organizations that the connection will be enabled for.
    
</dd>
</dl>

<dl>
<dd>

**ttlSec:** `*int` — Number of seconds for which the ticket is valid before expiration. If unspecified or set to 0, this value defaults to 432000 seconds (5 days).
    
</dd>
</dl>

<dl>
<dd>

**domainAliasesConfig:** `*management.SelfServiceProfileSSOTicketDomainAliasesConfig` 
    
</dd>
</dl>

<dl>
<dd>

**provisioningConfig:** `*management.SelfServiceProfileSSOTicketProvisioningConfig` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SelfServiceProfiles.SSOTicket.Revoke(ProfileID, ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Revokes an SSO access ticket and invalidates associated sessions. The ticket will no longer be accepted to initiate a Self-Service SSO session. If any users have already started a session through this ticket, their session will be terminated. Clients should expect a `202 Accepted` response upon successful processing, indicating that the request has been acknowledged and that the revocation is underway but may not be fully completed at the time of response. If the specified ticket does not exist, a `202 Accepted` response is also returned, signaling that no further action is required.
Clients should treat these `202` responses as an acknowledgment that the request has been accepted and is in progress, even if the ticket was not found.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.SelfServiceProfiles.SSOTicket.Revoke(
        context.TODO(),
        "profileId",
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**profileID:** `string` — The id of the self-service profile
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` — The id of the ticket to revoke
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tenants Settings
<details><summary><code>client.Tenants.Settings.Get() -> *management.GetTenantSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve tenant settings. A list of fields to include or exclude may also be specified.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetTenantSettingsRequestParameters{
        Fields: management.String(
            "fields",
        ),
        IncludeFields: management.Bool(
            true,
        ),
    }
client.Tenants.Settings.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fields:** `*string` — Comma-separated list of fields to include or exclude (based on value provided for include_fields) in the result. Leave empty to retrieve all fields.
    
</dd>
</dl>

<dl>
<dd>

**includeFields:** `*bool` — Whether specified fields are to be included (true) or excluded (false).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tenants.Settings.Update(request) -> *management.UpdateTenantSettingsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update settings for a tenant.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateTenantSettingsRequestContent{}
client.Tenants.Settings.Update(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**changePassword:** `*management.TenantSettingsPasswordPage` 
    
</dd>
</dl>

<dl>
<dd>

**deviceFlow:** `*management.TenantSettingsDeviceFlow` 
    
</dd>
</dl>

<dl>
<dd>

**guardianMfaPage:** `*management.TenantSettingsGuardianPage` 
    
</dd>
</dl>

<dl>
<dd>

**defaultAudience:** `*string` — Default audience for API Authorization.
    
</dd>
</dl>

<dl>
<dd>

**defaultDirectory:** `*string` — Name of connection used for password grants at the `/token` endpoint. The following connection types are supported: LDAP, AD, Database Connections, Passwordless, Windows Azure Active Directory, ADFS.
    
</dd>
</dl>

<dl>
<dd>

**errorPage:** `*management.TenantSettingsErrorPage` 
    
</dd>
</dl>

<dl>
<dd>

**defaultTokenQuota:** `*management.DefaultTokenQuota` 
    
</dd>
</dl>

<dl>
<dd>

**flags:** `*management.TenantSettingsFlags` 
    
</dd>
</dl>

<dl>
<dd>

**friendlyName:** `*string` — Friendly name for this tenant.
    
</dd>
</dl>

<dl>
<dd>

**pictureURL:** `*string` — URL of logo to be shown for this tenant (recommended size: 150x150)
    
</dd>
</dl>

<dl>
<dd>

**supportEmail:** `*string` — End-user support email.
    
</dd>
</dl>

<dl>
<dd>

**supportURL:** `*string` — End-user support url.
    
</dd>
</dl>

<dl>
<dd>

**allowedLogoutURLs:** `[]string` — URLs that are valid to redirect to after logout from Auth0.
    
</dd>
</dl>

<dl>
<dd>

**sessionLifetime:** `*int` — Number of hours a session will stay valid.
    
</dd>
</dl>

<dl>
<dd>

**idleSessionLifetime:** `*int` — Number of hours for which a session can be inactive before the user must log in again.
    
</dd>
</dl>

<dl>
<dd>

**ephemeralSessionLifetime:** `*int` — Number of hours an ephemeral (non-persistent) session will stay valid.
    
</dd>
</dl>

<dl>
<dd>

**idleEphemeralSessionLifetime:** `*int` — Number of hours for which an ephemeral (non-persistent) session can be inactive before the user must log in again.
    
</dd>
</dl>

<dl>
<dd>

**sandboxVersion:** `*string` — Selected sandbox version for the extensibility environment
    
</dd>
</dl>

<dl>
<dd>

**legacySandboxVersion:** `*string` — Selected legacy sandbox version for the extensibility environment
    
</dd>
</dl>

<dl>
<dd>

**defaultRedirectionURI:** `*string` — The default absolute redirection uri, must be https
    
</dd>
</dl>

<dl>
<dd>

**enabledLocales:** `[]*tenants.UpdateTenantSettingsRequestContentEnabledLocalesItem` — Supported locales for the user interface
    
</dd>
</dl>

<dl>
<dd>

**sessionCookie:** `*management.SessionCookieSchema` 
    
</dd>
</dl>

<dl>
<dd>

**sessions:** `*management.TenantSettingsSessions` 
    
</dd>
</dl>

<dl>
<dd>

**oidcLogout:** `*management.TenantOidcLogoutSettings` 
    
</dd>
</dl>

<dl>
<dd>

**customizeMfaInPostloginAction:** `*bool` — Whether to enable flexible factors for MFA in the PostLogin action
    
</dd>
</dl>

<dl>
<dd>

**allowOrganizationNameInAuthenticationAPI:** `*bool` — Whether to accept an organization name instead of an ID on auth endpoints
    
</dd>
</dl>

<dl>
<dd>

**acrValuesSupported:** `[]string` — Supported ACR values
    
</dd>
</dl>

<dl>
<dd>

**mtls:** `*management.TenantSettingsMtls` 
    
</dd>
</dl>

<dl>
<dd>

**pushedAuthorizationRequestsSupported:** `*bool` — Enables the use of Pushed Authorization Requests
    
</dd>
</dl>

<dl>
<dd>

**authorizationResponseIssParameterSupported:** `*bool` — Supports iss parameter in authorization responses
    
</dd>
</dl>

<dl>
<dd>

**skipNonVerifiableCallbackURIConfirmationPrompt:** `*bool` 

Controls whether a confirmation prompt is shown during login flows when the redirect URI uses non-verifiable callback URIs (for example, a custom URI schema such as `myapp://`, or `localhost`).
If set to true, a confirmation prompt will not be shown. We recommend that this is set to false for improved protection from malicious apps.
See https://auth0.com/docs/secure/security-guidance/measures-against-app-impersonation for more information.
    
</dd>
</dl>

<dl>
<dd>

**resourceParameterProfile:** `*management.TenantSettingsResourceParameterProfile` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users AuthenticationMethods
<details><summary><code>client.Users.AuthenticationMethods.List(ID) -> *management.ListUserAuthenticationMethodsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed list of authentication methods associated with a specified user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserAuthenticationMethodsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Users.AuthenticationMethods.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the user in question.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0. Default is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Default is 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.AuthenticationMethods.Create(ID, request) -> *management.CreateUserAuthenticationMethodResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create an authentication method. Authentication methods created via this endpoint will be auto confirmed and should already have verification completed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateUserAuthenticationMethodRequestContent{
        Type: management.CreatedUserAuthenticationMethodTypeEnumPhone,
    }
client.Users.AuthenticationMethods.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the user to whom the new authentication method will be assigned.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*management.CreatedUserAuthenticationMethodTypeEnum` 
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — A human-readable label to identify the authentication method.
    
</dd>
</dl>

<dl>
<dd>

**totpSecret:** `*string` — Base32 encoded secret for TOTP generation.
    
</dd>
</dl>

<dl>
<dd>

**phoneNumber:** `*string` — Applies to phone authentication methods only. The destination phone number used to send verification codes via text and voice.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — Applies to email authentication methods only. The email address used to send verification messages.
    
</dd>
</dl>

<dl>
<dd>

**preferredAuthenticationMethod:** `*management.PreferredAuthenticationMethodEnum` 
    
</dd>
</dl>

<dl>
<dd>

**keyID:** `*string` — Applies to webauthn authentication methods only. The id of the credential.
    
</dd>
</dl>

<dl>
<dd>

**publicKey:** `*string` — Applies to webauthn authentication methods only. The public key, which is encoded as base64.
    
</dd>
</dl>

<dl>
<dd>

**relyingPartyIdentifier:** `*string` — Applies to webauthn authentication methods only. The relying party identifier.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.AuthenticationMethods.Set(ID, request) -> []*management.SetUserAuthenticationMethodResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Replace the specified user <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors"> authentication methods</a> with supplied values.

    <b>Note</b>: Authentication methods supplied through this action do not iterate on existing methods. Instead, any methods passed will overwrite the user&#8217s existing settings.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := []*management.SetUserAuthenticationMethods{
        &management.SetUserAuthenticationMethods{
            Type: management.AuthenticationTypeEnumPhone,
        },
    }
client.Users.AuthenticationMethods.Set(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the user in question.
    
</dd>
</dl>

<dl>
<dd>

**request:** `management.SetUserAuthenticationMethodsRequestContent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.AuthenticationMethods.DeleteAll(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove all authentication methods (i.e., enrolled MFA factors) from the specified user account. This action cannot be undone. 
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.AuthenticationMethods.DeleteAll(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the user in question.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.AuthenticationMethods.Get(ID, AuthenticationMethodID) -> *management.GetUserAuthenticationMethodResponseContent</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.AuthenticationMethods.Get(
        context.TODO(),
        "id",
        "authentication_method_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the user in question.
    
</dd>
</dl>

<dl>
<dd>

**authenticationMethodID:** `string` — The ID of the authentication methods in question.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.AuthenticationMethods.Delete(ID, AuthenticationMethodID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove the authentication method with the given ID from the specified user. For more information, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/manage-mfa-auth0-apis/manage-authentication-methods-with-management-api">Manage Authentication Methods with Management API</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.AuthenticationMethods.Delete(
        context.TODO(),
        "id",
        "authentication_method_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the user in question.
    
</dd>
</dl>

<dl>
<dd>

**authenticationMethodID:** `string` — The ID of the authentication method to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.AuthenticationMethods.Update(ID, AuthenticationMethodID, request) -> *management.UpdateUserAuthenticationMethodResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Modify the authentication method with the given ID from the specified user. For more information, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/manage-mfa-auth0-apis/manage-authentication-methods-with-management-api">Manage Authentication Methods with Management API</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateUserAuthenticationMethodRequestContent{}
client.Users.AuthenticationMethods.Update(
        context.TODO(),
        "id",
        "authentication_method_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The ID of the user in question.
    
</dd>
</dl>

<dl>
<dd>

**authenticationMethodID:** `string` — The ID of the authentication method to update.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — A human-readable label to identify the authentication method.
    
</dd>
</dl>

<dl>
<dd>

**preferredAuthenticationMethod:** `*management.PreferredAuthenticationMethodEnum` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Authenticators
<details><summary><code>client.Users.Authenticators.DeleteAll(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove all authenticators registered to a given user ID, such as OTP, email, phone, and push-notification. This action cannot be undone. For more information, review <a href="https://auth0.com/docs/secure/multi-factor-authentication/manage-mfa-auth0-apis/manage-authentication-methods-with-management-api">Manage Authentication Methods with Management API</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Authenticators.DeleteAll(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to delete.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users ConnectedAccounts
<details><summary><code>client.Users.ConnectedAccounts.List(ID) -> *management.ListUserConnectedAccountsResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all connected accounts associated with the user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.GetUserConnectedAccountsRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.Users.ConnectedAccounts.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to list connected accounts for.
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results to return.  Defaults to 10 with a maximum of 20
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Enrollments
<details><summary><code>client.Users.Enrollments.Get(ID) -> []*management.UsersEnrollment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the first <a href="https://auth0.com/docs/secure/multi-factor-authentication/multi-factor-authentication-factors">multi-factor authentication</a> enrollment that a specific user has confirmed.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Enrollments.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to list enrollments for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users FederatedConnectionsTokensets
<details><summary><code>client.Users.FederatedConnectionsTokensets.List(ID) -> []*management.FederatedConnectionTokenSet</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List active federated connections tokensets for a provided user
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.FederatedConnectionsTokensets.List(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — User identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.FederatedConnectionsTokensets.Delete(ID, TokensetID) -> error</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.FederatedConnectionsTokensets.Delete(
        context.TODO(),
        "id",
        "tokenset_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — Id of the user that owns the tokenset
    
</dd>
</dl>

<dl>
<dd>

**tokensetID:** `string` — The tokenset id
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Identities
<details><summary><code>client.Users.Identities.Link(ID, request) -> []*management.UserIdentity</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Link two user accounts together forming a primary and secondary relationship. On successful linking, the endpoint returns the new array of the primary account identities.

Note: There are two ways of invoking the endpoint:

<ul>
  <li>With the authenticated primary account's JWT in the Authorization header, which has the <code>update:current_user_identities</code> scope:
    <pre>
      POST /api/v2/users/PRIMARY_ACCOUNT_USER_ID/identities
      Authorization: "Bearer PRIMARY_ACCOUNT_JWT"
      {
        "link_with": "SECONDARY_ACCOUNT_JWT"
      }
    </pre>
    In this case, only the <code>link_with</code> param is required in the body, which also contains the JWT obtained upon the secondary account's authentication.
  </li>
  <li>With a token generated by the API V2 containing the <code>update:users</code> scope:
    <pre>
    POST /api/v2/users/PRIMARY_ACCOUNT_USER_ID/identities
    Authorization: "Bearer YOUR_API_V2_TOKEN"
    {
      "provider": "SECONDARY_ACCOUNT_PROVIDER",
      "connection_id": "SECONDARY_ACCOUNT_CONNECTION_ID(OPTIONAL)",
      "user_id": "SECONDARY_ACCOUNT_USER_ID"
    }
    </pre>
    In this case you need to send <code>provider</code> and <code>user_id</code> in the body. Optionally you can also send the <code>connection_id</code> param which is suitable for identifying a particular database connection for the 'auth0' provider.
  </li>
</ul>
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.LinkUserIdentityRequestContent{}
client.Users.Identities.Link(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the primary user account to link a second user account to.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*management.UserIdentityProviderEnum` 
    
</dd>
</dl>

<dl>
<dd>

**connectionID:** `*string` — connection_id of the secondary user account being linked when more than one `auth0` database provider exists.
    
</dd>
</dl>

<dl>
<dd>

**userID:** `*management.UserID` 
    
</dd>
</dl>

<dl>
<dd>

**linkWith:** `*string` — JWT for the secondary account being linked. If sending this parameter, `provider`, `user_id`, and `connection_id` must not be sent.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Identities.Delete(ID, Provider, UserID) -> management.DeleteUserIdentityResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Unlink a specific secondary account from a target user. This action requires the ID of both the target user and the secondary account. 

Unlinking the secondary account removes it from the identities array of the target user and creates a new standalone profile for the secondary account. To learn more, review <a href="https://auth0.com/docs/manage-users/user-accounts/user-account-linking/unlink-user-accounts">Unlink User Accounts</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Identities.Delete(
        context.TODO(),
        "id",
        management.UserIdentityProviderEnumAd.Ptr(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the primary user account.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*management.UserIdentityProviderEnum` — Identity provider name of the secondary linked account (e.g. `google-oauth2`).
    
</dd>
</dl>

<dl>
<dd>

**userID:** `string` — ID of the secondary linked account (e.g. `123456789081523216417` part after the `|` in `google-oauth2|123456789081523216417`).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Logs
<details><summary><code>client.Users.Logs.List(ID) -> *management.UserListLogOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve log events for a specific user.

Note: For more information on all possible event types, their respective acronyms and descriptions, see <a href="https://auth0.com/docs/logs/log-event-type-codes">Log Event Type Codes</a>.

For more information on the list of fields that can be used in `sort`, see <a href="https://auth0.com/docs/logs/log-search-query-syntax#searchable-fields">Searchable Fields</a>.

Auth0 <a href="https://auth0.com/docs/logs/retrieve-log-events-using-mgmt-api#limitations">limits the number of logs</a> you can return by search criteria to 100 logs per request. Furthermore, you may only paginate through up to 1,000 search results. If you exceed this threshold, please redefine your search.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserLogsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        Sort: management.String(
            "sort",
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Users.Logs.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user of the logs to retrieve
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Paging is disabled if parameter not sent.
    
</dd>
</dl>

<dl>
<dd>

**sort:** `*string` — Field to sort by. Use `fieldname:1` for ascending order and `fieldname:-1` for descending.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Multifactor
<details><summary><code>client.Users.Multifactor.InvalidateRememberBrowser(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Invalidate all remembered browsers across all <a href="https://auth0.com/docs/multifactor-authentication">authentication factors</a> for a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Multifactor.InvalidateRememberBrowser(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to invalidate all remembered browsers and authentication factors for.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Multifactor.DeleteProvider(ID, Provider) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove a <a href="https://auth0.com/docs/multifactor-authentication">multifactor</a> authentication configuration from a user's account. This forces the user to manually reconfigure the multi-factor provider.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Multifactor.DeleteProvider(
        context.TODO(),
        "id",
        management.UserMultifactorProviderEnumDuo.Ptr(),
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to remove a multifactor configuration from.
    
</dd>
</dl>

<dl>
<dd>

**provider:** `*management.UserMultifactorProviderEnum` — The multi-factor provider. Supported values 'duo' or 'google-authenticator'
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Organizations
<details><summary><code>client.Users.Organizations.List(ID) -> *management.ListUserOrganizationsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve list of the specified user's current Organization memberships. User must be specified by user ID. For more information, review <a href="https://auth0.com/docs/manage-users/organizations">Auth0 Organizations</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserOrganizationsRequestParameters{
        Page: management.Int(
            1,
        ),
        PerPage: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Users.Organizations.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to retrieve the organizations for.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Permissions
<details><summary><code>client.Users.Permissions.List(ID) -> *management.ListUserPermissionsOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve all permissions associated with the user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserPermissionsRequestParameters{
        PerPage: management.Int(
            1,
        ),
        Page: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Users.Permissions.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to retrieve the permissions for.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Permissions.Create(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assign permissions to a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateUserPermissionsRequestContent{
        Permissions: []*management.PermissionRequestPayload{
            &management.PermissionRequestPayload{
                ResourceServerIdentifier: "resource_server_identifier",
                PermissionName: "permission_name",
            },
        },
    }
client.Users.Permissions.Create(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to assign permissions to.
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]*management.PermissionRequestPayload` — List of permissions to add to this user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Permissions.Delete(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove permissions from a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteUserPermissionsRequestContent{
        Permissions: []*management.PermissionRequestPayload{
            &management.PermissionRequestPayload{
                ResourceServerIdentifier: "resource_server_identifier",
                PermissionName: "permission_name",
            },
        },
    }
client.Users.Permissions.Delete(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to remove permissions from.
    
</dd>
</dl>

<dl>
<dd>

**permissions:** `[]*management.PermissionRequestPayload` — List of permissions to remove from this user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users RiskAssessments
<details><summary><code>client.Users.RiskAssessments.Clear(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Clear risk assessment assessors for a specific user
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ClearAssessorsRequestContent{
        Connection: "connection",
        Assessors: []management.AssessorsTypeEnum{
            management.AssessorsTypeEnum(
                "new-device",
            ),
        },
    }
client.Users.RiskAssessments.Clear(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to clear assessors for.
    
</dd>
</dl>

<dl>
<dd>

**connection:** `string` — The name of the connection containing the user whose assessors should be cleared.
    
</dd>
</dl>

<dl>
<dd>

**assessors:** `[]management.AssessorsTypeEnum` — List of assessors to clear.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Roles
<details><summary><code>client.Users.Roles.List(ID) -> *management.ListUserRolesOffsetPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve detailed list of all user roles currently assigned to a user.

<b>Note</b>: This action retrieves all roles assigned to a user in the context of your whole tenant. To retrieve Organization-specific roles, use the following endpoint: <a href="https://auth0.com/docs/api/management/v2/organizations/get-organization-member-roles">Get user roles assigned to an Organization member</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserRolesRequestParameters{
        PerPage: management.Int(
            1,
        ),
        Page: management.Int(
            1,
        ),
        IncludeTotals: management.Bool(
            true,
        ),
    }
client.Users.Roles.List(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to list roles for.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Number of results per page.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — Page index of the results to return. First page is 0.
    
</dd>
</dl>

<dl>
<dd>

**includeTotals:** `*bool` — Return results inside an object that contains the total result count (true) or as a direct array of results (false, default).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Roles.Assign(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Assign one or more existing user roles to a user. For more information, review <a href="https://auth0.com/docs/manage-users/access-control/rbac">Role-Based Access Control</a>.

<b>Note</b>: New roles cannot be created through this action. Additionally, this action is used to assign roles to a user in the context of your whole tenant. To assign roles in the context of a specific Organization, use the following endpoint: <a href="https://auth0.com/docs/api/management/v2/organizations/post-organization-member-roles">Assign user roles to an Organization member</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.AssignUserRolesRequestContent{
        Roles: []string{
            "roles",
        },
    }
client.Users.Roles.Assign(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to associate roles with.
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]string` — List of roles IDs to associated with the user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Roles.Delete(ID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Remove one or more specified user roles assigned to a user.

<b>Note</b>: This action removes a role from a user in the context of your whole tenant. If you want to unassign a role from a user in the context of a specific Organization, use the following endpoint: <a href="https://auth0.com/docs/api/management/v2/organizations/delete-organization-member-roles">Delete user roles from an Organization member</a>.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.DeleteUserRolesRequestContent{
        Roles: []string{
            "roles",
        },
    }
client.Users.Roles.Delete(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the user to remove roles from.
    
</dd>
</dl>

<dl>
<dd>

**roles:** `[]string` — List of roles IDs to remove from the user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users RefreshToken
<details><summary><code>client.Users.RefreshToken.List(UserID) -> *management.ListRefreshTokensPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details for a user's refresh tokens.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListRefreshTokensRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.Users.RefreshToken.List(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — ID of the user to get refresh tokens for
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — An optional cursor from which to start the selection (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.RefreshToken.Delete(UserID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete all refresh tokens for a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.RefreshToken.Delete(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — ID of the user to get remove refresh tokens for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Users Sessions
<details><summary><code>client.Users.Sessions.List(UserID) -> *management.ListUserSessionsPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve details for a user's sessions.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListUserSessionsRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.Users.Sessions.List(
        context.TODO(),
        "user_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — ID of the user to get sessions for
    
</dd>
</dl>

<dl>
<dd>

**from:** `*string` — An optional cursor from which to start the selection (exclusive).
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Users.Sessions.Delete(UserID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete all sessions for a user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Users.Sessions.Delete(
        context.TODO(),
        "user_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — ID of the user to get sessions for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## VerifiableCredentials Verification Templates
<details><summary><code>client.VerifiableCredentials.Verification.Templates.List() -> *management.ListVerifiableCredentialTemplatesPaginatedResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List a verifiable credential templates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.ListVerifiableCredentialTemplatesRequestParameters{
        From: management.String(
            "from",
        ),
        Take: management.Int(
            1,
        ),
    }
client.VerifiableCredentials.Verification.Templates.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**from:** `*string` — Optional Id from which to start selection.
    
</dd>
</dl>

<dl>
<dd>

**take:** `*int` — Number of results per page. Defaults to 50.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiableCredentials.Verification.Templates.Create(request) -> *management.CreateVerifiableCredentialTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a verifiable credential template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.CreateVerifiableCredentialTemplateRequestContent{
        Name: "name",
        Type: "type",
        Dialect: "dialect",
        Presentation: &management.MdlPresentationRequest{
            OrgIso1801351MDl: &management.MdlPresentationRequestProperties{
                OrgIso1801351: &management.MdlPresentationProperties{},
            },
        },
        WellKnownTrustedIssuers: "well_known_trusted_issuers",
    }
client.VerifiableCredentials.Verification.Templates.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**dialect:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**presentation:** `*management.MdlPresentationRequest` 
    
</dd>
</dl>

<dl>
<dd>

**customCertificateAuthority:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**wellKnownTrustedIssuers:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiableCredentials.Verification.Templates.Get(ID) -> *management.GetVerifiableCredentialTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a verifiable credential template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.VerifiableCredentials.Verification.Templates.Get(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the template to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiableCredentials.Verification.Templates.Delete(ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a verifiable credential template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.VerifiableCredentials.Verification.Templates.Delete(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the template to retrieve.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.VerifiableCredentials.Verification.Templates.Update(ID, request) -> *management.UpdateVerifiableCredentialTemplateResponseContent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a verifiable credential template.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &management.UpdateVerifiableCredentialTemplateRequestContent{}
client.VerifiableCredentials.Verification.Templates.Update(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — ID of the template to retrieve.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**dialect:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**presentation:** `*management.MdlPresentationRequest` 
    
</dd>
</dl>

<dl>
<dd>

**wellKnownTrustedIssuers:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `*float64` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
