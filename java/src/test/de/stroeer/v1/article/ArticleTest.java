package de.stroeer.v1.article;

import de.stroeer.api.v1.article.Article.Builder;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class ArticleTest {

  @Test
  public void ModelCompiles() {
    Builder articleBuilder = de.stroeer.api.v1.article.Article.newBuilder();
    articleBuilder.setId("id");
    de.stroeer.api.v1.article.Article article = articleBuilder.build();
    assertEquals(article.getId(), "id");
  }
}