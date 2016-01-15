//************************************************************************//
// API "congo" version v1: Resource and Payload Helpers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/bketelsen/gorma/example
// --design=github.com/bketelsen/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package user

import (
	"github.com/gopheracademy/congo/app/v1"
	"github.com/raphael/goa"
)

func UserFromV1CreatePayload(ctx *v1.CreateUserContext) User {
	payload := ctx.Payload
	var err error
	target, _ := MarshalCreateUserPayload(payload, err)

	return target
}

// MarshalCreateUserPayload validates and renders an instance of CreateUserPayload into a interface{}
func MarshalCreateUserPayload(source *v1.CreateUserPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp13 := map[string]interface{}{
		"bio":        source.Bio,
		"city":       source.City,
		"country":    source.Country,
		"created_at": source.CreatedAt,
		"email":      source.Email,
		"first_name": source.FirstName,
		"id":         source.ID,
		"last_name":  source.LastName,
		"role":       source.Role,
		"state":      source.State,
		"updated_at": source.UpdatedAt,
	}
	target = tmp13
	return
}

func UserFromV1UpdatePayload(ctx *v1.UpdateUserContext) User {
	payload := ctx.Payload
	var err error
	target, _ := MarshalUpdateUserPayload(payload, err)

	return target
}

// MarshalUpdateUserPayload validates and renders an instance of UpdateUserPayload into a interface{}
func MarshalUpdateUserPayload(source *v1.UpdateUserPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
		return
	}
	tmp14 := map[string]interface{}{
		"bio":        source.Bio,
		"city":       source.City,
		"country":    source.Country,
		"created_at": source.CreatedAt,
		"email":      source.Email,
		"first_name": source.FirstName,
		"id":         source.ID,
		"last_name":  source.LastName,
		"role":       source.Role,
		"state":      source.State,
		"updated_at": source.UpdatedAt,
	}
	target = tmp14
	return
}
