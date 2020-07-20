package test

import (
	"testing"

	"github.com/stroeer/tapir/go/stroeer/core/v1"
	"github.com/stroeer/tapir/go/stroeer/web/article/v1"
	"google.golang.org/protobuf/proto"
)

func TestGenerator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input proto.Message
		exp   string
	}{
		{name: "v1.Article", input: &core.Article{}, exp: "stroeer.core.v1.Article"},
		{name: "v1.GetArticleRequest", input: &core.GetArticleRequest{}, exp: "stroeer.core.v1.GetArticleRequest"},
		{name: "v1.WebArticle", input: &article.WebArticle{}, exp: "stroeer.web.article.v1.WebArticle"},
		{name: "v1.GetWebArticlePageRequest", input: &article.GetWebArticlePageRequest{}, exp: "stroeer.web.article.v1.GetWebArticlePageRequest"},
		{name: "v1.GetWebArticlePageResponse", input: &article.GetWebArticlePageResponse{}, exp: "stroeer.web.article.v1.GetWebArticlePageResponse"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			name := string(tc.input.ProtoReflect().Descriptor().FullName())
			if name != tc.exp {
				t.Fatalf("expected proto message with name %s, got: %s", tc.exp, name)
			}
		})
	}
}
