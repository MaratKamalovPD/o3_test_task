package resolvers

import "fmt"

const maxCommentLength = 2000

var (
	errCommentIsTooLong      = fmt.Errorf("comment is too long")
	errIdIsLessThanOne       = fmt.Errorf("id should not be less than 1")
	errInvalidPaginationArgs = fmt.Errorf("invalid pagination args")
)

func ValidateCommentLength(comment string) error {

	if len(comment) >= maxCommentLength {

		return errCommentIsTooLong
	}

	return nil
}

func ValidateID(IDs ...int) error {

	for _, id := range IDs {

		if id < 1 {

			return errIdIsLessThanOne
		}
	}

	return nil
}

func ValidatePaginationArgs(limit, offset int) error {

	if limit < 1 || offset < 0 {

		return errInvalidPaginationArgs
	}
	return nil
}
