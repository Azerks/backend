package errs

// slug represents a unique error type that can be detected by the client
type slug string

// available slugs
const (
	SlugNotFound       slug = "not-found"
	SlugRequestMissing slug = "request-missing"
	SlugRequestInvalid slug = "request-invalid"
	SlugUnauthorized   slug = "unauthorized"
	SlugForbidden      slug = "forbidden"
	SlugDuplicate      slug = "already-exists"
	SlugNotImplemented slug = "not-implemented"
	SlugConstraint     slug = "constraint-violation"
	SlugInternal       slug = "internal-error"
	SlugCreate         slug = "create-error"
	SlugUpdate         slug = "update-error"
	SlugDelete         slug = "delete-error"

	SlugUnknown slug = "unknown"
)
