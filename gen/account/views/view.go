// Code generated by goa v3.4.3, DO NOT EDIT.
//
// account views
//
// Command:
// $ goa gen algodexidx/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// TrackedAccountCollection is the viewed result type that is projected based
// on a view.
type TrackedAccountCollection struct {
	// Type to project
	Projected TrackedAccountCollectionView
	// View to render
	View string
}

// TrackedAccountCollectionView is a type that runs validations on a projected
// type.
type TrackedAccountCollectionView []*TrackedAccountView

// TrackedAccountView is a type that runs validations on a projected type.
type TrackedAccountView struct {
	// Public Account address
	Address *string
	// Opted-in ASA IDs
	Holdings map[string]uint64
}

var (
	// TrackedAccountCollectionMap is a map of attribute names in result type
	// TrackedAccountCollection indexed by view name.
	TrackedAccountCollectionMap = map[string][]string{
		"default": []string{
			"address",
		},
		"full": []string{
			"address",
			"holdings",
		},
	}
	// TrackedAccountMap is a map of attribute names in result type TrackedAccount
	// indexed by view name.
	TrackedAccountMap = map[string][]string{
		"default": []string{
			"address",
		},
		"full": []string{
			"address",
			"holdings",
		},
	}
)

// ValidateTrackedAccountCollection runs the validations defined on the viewed
// result type TrackedAccountCollection.
func ValidateTrackedAccountCollection(result TrackedAccountCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateTrackedAccountCollectionView(result.Projected)
	case "full":
		err = ValidateTrackedAccountCollectionViewFull(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "full"})
	}
	return
}

// ValidateTrackedAccountCollectionView runs the validations defined on
// TrackedAccountCollectionView using the "default" view.
func ValidateTrackedAccountCollectionView(result TrackedAccountCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateTrackedAccountView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTrackedAccountCollectionViewFull runs the validations defined on
// TrackedAccountCollectionView using the "full" view.
func ValidateTrackedAccountCollectionViewFull(result TrackedAccountCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateTrackedAccountViewFull(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTrackedAccountView runs the validations defined on
// TrackedAccountView using the "default" view.
func ValidateTrackedAccountView(result *TrackedAccountView) (err error) {
	if result.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "result"))
	}
	if result.Address != nil {
		if utf8.RuneCountInString(*result.Address) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.address", *result.Address, utf8.RuneCountInString(*result.Address), 58, false))
		}
	}
	return
}

// ValidateTrackedAccountViewFull runs the validations defined on
// TrackedAccountView using the "full" view.
func ValidateTrackedAccountViewFull(result *TrackedAccountView) (err error) {
	if result.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "result"))
	}
	if result.Holdings == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("holdings", "result"))
	}
	if result.Address != nil {
		if utf8.RuneCountInString(*result.Address) > 58 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.address", *result.Address, utf8.RuneCountInString(*result.Address), 58, false))
		}
	}
	return
}
