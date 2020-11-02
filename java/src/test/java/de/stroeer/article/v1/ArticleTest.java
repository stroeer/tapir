package de.stroeer.article.v1;

import static org.junit.jupiter.api.Assertions.assertEquals;

import de.stroeer.core.v1.Article;
import de.stroeer.core.v1.Article.Builder;
import org.junit.jupiter.api.Test;

class ArticleTest {

  @Test
  void ModelCompiles() {
    Builder articleBuilder = Article.newBuilder();
    articleBuilder.setId(42L);

    Article article = articleBuilder.build();

    assertEquals(42L, article.getId());
  }
}
