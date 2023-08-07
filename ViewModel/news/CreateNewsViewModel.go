package news

type CreateNewsViewModel struct {
	Title            string `validate:"required"`
	ShortDescription string `validate:"required"`
	Description      string `validate:"required"`
	ImageName        string
	CreatorUserId    string
}
