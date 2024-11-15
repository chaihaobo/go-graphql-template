package book

type (
	CreateBookArgs struct {
		Input CreateBookInput
	}
	CreateBookInput struct {
		Title  string
		Author string
		UserID string
	}
)
