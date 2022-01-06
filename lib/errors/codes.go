package errors

// Keep track of exit code statuses

const (
	// Success is returned when rclone finished without error.
	Success = iota

	// This error is the user's fault
	UserError

	// UnknownError doesn't have a category
	UnknownError

	// When a path or file is not found.
	PathNotFound

	// TODO add errors types as needed
)
