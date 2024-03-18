// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ActorService views
//
// Command:
// $ goa gen film-lib/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ActorResultCollection is the viewed result type that is projected based on a
// view.
type ActorResultCollection struct {
	// Type to project
	Projected ActorResultCollectionView
	// View to render
	View string
}

// ActorResult is the viewed result type that is projected based on a view.
type ActorResult struct {
	// Type to project
	Projected *ActorResultView
	// View to render
	View string
}

// ActorResultCollectionView is a type that runs validations on a projected
// type.
type ActorResultCollectionView []*ActorResultView

// ActorResultView is a type that runs validations on a projected type.
type ActorResultView struct {
	// Unique ID of an Actor
	ActorID        *uint64
	ActorName      *string
	ActorSex       *string
	ActorBirthdate *string
}

var (
	// ActorResultCollectionMap is a map indexing the attribute names of
	// ActorResultCollection by view name.
	ActorResultCollectionMap = map[string][]string{
		"default": {
			"ActorID",
			"ActorName",
			"ActorSex",
			"ActorBirthdate",
		},
	}
	// ActorResultMap is a map indexing the attribute names of ActorResult by view
	// name.
	ActorResultMap = map[string][]string{
		"default": {
			"ActorID",
			"ActorName",
			"ActorSex",
			"ActorBirthdate",
		},
	}
)

// ValidateActorResultCollection runs the validations defined on the viewed
// result type ActorResultCollection.
func ValidateActorResultCollection(result ActorResultCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateActorResultCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateActorResult runs the validations defined on the viewed result type
// ActorResult.
func ValidateActorResult(result *ActorResult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateActorResultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateActorResultCollectionView runs the validations defined on
// ActorResultCollectionView using the "default" view.
func ValidateActorResultCollectionView(result ActorResultCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateActorResultView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActorResultView runs the validations defined on ActorResultView
// using the "default" view.
func ValidateActorResultView(result *ActorResultView) (err error) {
	if result.ActorID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ActorID", "result"))
	}
	return
}
