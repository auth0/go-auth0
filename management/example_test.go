package management_test

import (
	"fmt"
	"os"
	"time"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/management"
)

var (
	domain = os.Getenv("AUTH0_DOMAIN")
	id     = os.Getenv("AUTH0_CLIENT_ID")
	secret = os.Getenv("AUTH0_CLIENT_SECRET")

	api *management.Management
)

func init() {
	var err error
	api, err = management.New(domain, management.WithClientCredentials(id, secret))
	if err != nil {
		panic(err)
	}
}

func ExampleNew() {
	api, err := management.New(domain, management.WithClientCredentials(id, secret))
	if err != nil {
		// handle err
		return
	}
	_, _ = api.Stat.ActiveUsers()
}

func ExampleClientManager_Create() {
	c := &management.Client{
		Name:        auth0.String("Example Client"),
		Description: auth0.String("This client was created from the Auth0 SDK examples"),
	}

	err := api.Client.Create(c)
	if err != nil {
		// handle err
		return
	}
	defer api.Client.Delete(c.GetClientID())

	_ = c.GetClientID()
	_ = c.GetClientSecret() // Generated values are available after creation
}

func ExampleResourceServerManager_List() {
	l, err := api.ResourceServer.List()
	if err != nil {
		// handle err
		return
	}
	_ = l.ResourceServers
}

func ExampleUserManager_Create() {
	u := &management.User{
		Connection: auth0.String("Username-Password-Authentication"),
		Email:      auth0.String("smith@example.com"),
		Username:   auth0.String("smith"),
		Password:   auth0.String("F4e3DA1a6cDD"),
	}

	err := api.User.Create(u)
	if err != nil {
		// handle err
		return
	}
	defer api.User.Delete(u.GetID())

	_ = u.GetID()
	_ = u.GetCreatedAt()
}

func ExampleRoleManager_Create() {
	r := &management.Role{
		Name:        auth0.String("admin"),
		Description: auth0.String("Administrator"),
	}
	err := api.Role.Create(r)
	if err != nil {
		// handle err
		return
	}
	defer api.Role.Delete(r.GetID())
}

var (
	user  = &management.User{}
	admin = &management.Role{}
)

func ExampleUserManager_AssignRoles() {
	err := api.User.AssignRoles(user.GetID(), []*management.Role{admin})
	if err != nil {
		// handle err
		return
	}
	defer api.User.RemoveRoles(user.GetID(), []*management.Role{admin})
}

func ExampleUserManager_List() {
	q := management.Query(`name:"jane smith"`)
	l, err := api.User.List(q)
	if err != nil {
		// handle err
		return
	}
	_ = l.Users // users matching name "jane smith"
}

func ExampleUserManager_List_pagination() {
	var page int
	for {
		l, err := api.User.List(
			management.Query(`logins_count:{100 TO *]`),
			management.Page(page))
		if err != nil {
			// handle err
			return
		}
		for _, u := range l.Users {
			u.GetID() // do something with each user
		}
		if !l.HasNext() {
			break
		}
		page++
	}
}

func ExampleConnectionManager_List() {
	connectionList, err := api.Connection.List(
		management.Parameter("strategy", "auth0"),
	)
	if err != nil {
		// handle err
		return
	}
	for _, connection := range connectionList.Connections {
		fmt.Println(connection.GetName())

		if options, ok := connection.Options.(*management.ConnectionOptions); ok {
			fmt.Printf("\tPassword Policy: %s\n", options.GetPasswordPolicy())
			fmt.Printf("\tMulti-Factor Auth Enabled: %t\n", options.MFA["active"])
		}
	}
}

func ExampleConnectionManager_Create() {
	connection := &management.Connection{
		Name:     auth0.Stringf("Test-Google-OAuth2-%d", time.Now().Unix()),
		Strategy: auth0.String("google-oauth2"),
		Options: &management.ConnectionOptionsGoogleOAuth2{
			ClientID:     auth0.String(""), // replace with your client id
			ClientSecret: auth0.String(""),
			AllowedAudiences: &[]string{
				"example.com",
				"api.example.com",
			},
			Profile:  auth0.Bool(true),
			Calendar: auth0.Bool(true),
			Youtube:  auth0.Bool(false),
		},
	}

	defer api.Connection.Delete(connection.GetID())

	err := api.Connection.Create(connection)
	if err != nil {
		// handle err
		return
	}
}
