package main;

import de.stroeer.article.v1.Article;
import de.stroeer.article.v1.Article.Builder;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class ArticleTest {

  @Test
  public void ModelCompiles() {
    Builder articleBuilder = Article.newBuilder();
    articleBuilder.setId("id");
    Article article = articleBuilder.build();
    assertEquals(article.getId(), "id");
  }
}