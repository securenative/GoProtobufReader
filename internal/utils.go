package internal

import (
	"github.com/emicklei/proto"
	"strings"
)

func trimComment(comment *proto.Comment) string {
	if comment == nil {
		return ""
	}

	return strings.Trim(comment.Message(), " ")
}
