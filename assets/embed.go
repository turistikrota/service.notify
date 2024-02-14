package assets

import "embed"

//go:embed mail/*
var emailTemplate embed.FS

func EmbedMailTemplate() embed.FS {
	return emailTemplate
}
