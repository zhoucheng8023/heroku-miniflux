// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package oauth2 // import "miniflux.app/oauth2"

import (
	"context"

	"miniflux.app/model"
)

// Provider is an interface for OAuth2 providers.
type Provider interface {
	GetUserExtraKey() string
	GetRedirectURL(state string) string
	GetProfile(ctx context.Context, code string) (*Profile, error)
	PopulateUserCreationWithProfileID(user *model.UserCreationRequest, profile *Profile)
	PopulateUserWithProfileID(user *model.User, profile *Profile)
	UnsetUserProfileID(user *model.User)
}
