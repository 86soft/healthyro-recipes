package postgresql

type RecipeModel struct {
	Base
	title        string
	description  string
	externalLink string
}
