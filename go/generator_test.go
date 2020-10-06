package test

import (
	"testing"

	"github.com/stroeer/tapir/go/stroeer/core/v1"
	"github.com/stroeer/tapir/go/stroeer/pages/article/v1"
	"github.com/stroeer/tapir/go/stroeer/pages/section/v1"
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
		{name: "v1.GetArticlePageRequest", input: &article.GetArticlePageRequest{}, exp: "stroeer.pages.article.v1.GetArticlePageRequest"},
		{name: "v1.GetArticlePageResponse", input: &article.GetArticlePageResponse{}, exp: "stroeer.pages.article.v1.GetArticlePageResponse"},
		{name: "v1.GetSectionPageRequest", input: &section.GetSectionPageRequest{}, exp: "stroeer.pages.section.v1.GetSectionPageRequest"},
		{name: "v1.GetSectionPageResponse", input: &section.GetSectionPageResponse{}, exp: "stroeer.pages.section.v1.GetSectionPageResponse"},
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
