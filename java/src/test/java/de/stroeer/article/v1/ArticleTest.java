package de.stroeer.article.v1;

import static org.junit.jupiter.api.Assertions.assertEquals;

import de.stroeer.core.v1.Article;
import de.stroeer.core.v1.Article.Builder;
import org.junit.jupiter.api.Test;

public class ArticleTest {

  @Test
  public void ModelCompiles() {
    Builder articleBuilder = Article.newBuilder();
    articleBuilder.setId("id");
    Article article = articleBuilder.build();
    assertEquals(article.getId(), "id");
  }
}
