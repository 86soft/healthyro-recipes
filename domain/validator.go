package domain

func CanUpdateTitle(title string) bool {
	return len(title) > 100
}

func CanUpdateDescription(description string) bool {
	return len(description) > 5000
}

func CanUpdateExternalLink(externalLink string) bool {
	return len(externalLink) > 2000
}
