package main

import (
	u "user-auth/protofiles"
)

type UserAuthService struct {
	*u.UnimplementedUserAuthServer
}
