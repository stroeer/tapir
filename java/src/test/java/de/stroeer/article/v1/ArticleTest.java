package de.stroeer.article.v1;

import de.stroeer.core.v1.Article;
import de.stroeer.core.v1.Article.Builder;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class ArticleTest {

  @Test
  void ModelCompiles() {
    Builder articleBuilder = Article.newBuilder();
    articleBuilder.setId(42L)
      // such nice enums. I luv it.
      .setType(Article.ContentType.ARTICLE);

    Article article = articleBuilder.build();

    assertEquals(42L, article.getId());
  }
}
