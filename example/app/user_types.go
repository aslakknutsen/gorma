//************************************************************************//
// API "congo": Application User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// ProposalPayload user type.
type proposalPayload struct {
	Abstract  *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail    *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	Withdrawn *bool   `json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// Validate validates the proposalPayload type instance.
func (ut *proposalPayload) Validate() (err error) {
	if ut.Abstract != nil {
		if len(*ut.Abstract) < 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 50, true))
		}
	}
	if ut.Abstract != nil {
		if len(*ut.Abstract) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 500, false))
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) < 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 100, true))
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 2000, false))
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 10, true))
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 200, false))
		}
	}
	return
}

// Publicize creates ProposalPayload from proposalPayload
func (ut *proposalPayload) Publicize() *ProposalPayload {
	var pub ProposalPayload
	if ut.Abstract != nil {
		pub.Abstract = ut.Abstract
	}
	if ut.Detail != nil {
		pub.Detail = ut.Detail
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	if ut.Withdrawn != nil {
		pub.Withdrawn = ut.Withdrawn
	}
	return &pub
}

// ProposalPayload user type.
type ProposalPayload struct {
	Abstract  *string `json:"abstract,omitempty" xml:"abstract,omitempty"`
	Detail    *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	Withdrawn *bool   `json:"withdrawn,omitempty" xml:"withdrawn,omitempty"`
}

// Validate validates the ProposalPayload type instance.
func (ut *ProposalPayload) Validate() (err error) {
	if ut.Abstract != nil {
		if len(*ut.Abstract) < 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 50, true))
		}
	}
	if ut.Abstract != nil {
		if len(*ut.Abstract) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.abstract`, *ut.Abstract, len(*ut.Abstract), 500, false))
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) < 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 100, true))
		}
	}
	if ut.Detail != nil {
		if len(*ut.Detail) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.detail`, *ut.Detail, len(*ut.Detail), 2000, false))
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 10, true))
		}
	}
	if ut.Title != nil {
		if len(*ut.Title) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, len(*ut.Title), 200, false))
		}
	}
	return err
}

// ReviewPayload user type.
type reviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  *int    `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the reviewPayload type instance.
func (ut *reviewPayload) Validate() (err error) {
	if ut.Comment != nil {
		if len(*ut.Comment) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 10, true))
		}
	}
	if ut.Comment != nil {
		if len(*ut.Comment) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 200, false))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false))
		}
	}
	return
}

// Publicize creates ReviewPayload from reviewPayload
func (ut *reviewPayload) Publicize() *ReviewPayload {
	var pub ReviewPayload
	if ut.Comment != nil {
		pub.Comment = ut.Comment
	}
	if ut.Rating != nil {
		pub.Rating = ut.Rating
	}
	return &pub
}

// ReviewPayload user type.
type ReviewPayload struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Rating  *int    `json:"rating,omitempty" xml:"rating,omitempty"`
}

// Validate validates the ReviewPayload type instance.
func (ut *ReviewPayload) Validate() (err error) {
	if ut.Comment != nil {
		if len(*ut.Comment) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 10, true))
		}
	}
	if ut.Comment != nil {
		if len(*ut.Comment) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.comment`, *ut.Comment, len(*ut.Comment), 200, false))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false))
		}
	}
	return err
}

// UserPayload user type.
type userPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	Firstname *string `json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty" xml:"lastname,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Bio != nil {
		if len(*ut.Bio) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.bio`, *ut.Bio, len(*ut.Bio), 500, false))
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Bio != nil {
		pub.Bio = ut.Bio
	}
	if ut.City != nil {
		pub.City = ut.City
	}
	if ut.Country != nil {
		pub.Country = ut.Country
	}
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Firstname != nil {
		pub.Firstname = ut.Firstname
	}
	if ut.Lastname != nil {
		pub.Lastname = ut.Lastname
	}
	if ut.State != nil {
		pub.State = ut.State
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	Bio       *string `json:"bio,omitempty" xml:"bio,omitempty"`
	City      *string `json:"city,omitempty" xml:"city,omitempty"`
	Country   *string `json:"country,omitempty" xml:"country,omitempty"`
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	Firstname *string `json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty" xml:"lastname,omitempty"`
	State     *string `json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Bio != nil {
		if len(*ut.Bio) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.bio`, *ut.Bio, len(*ut.Bio), 500, false))
		}
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	return err
}
