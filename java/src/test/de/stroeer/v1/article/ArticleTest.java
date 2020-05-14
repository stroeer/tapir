package de.stroeer.v1.article;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.module.SimpleModule;
import de.stroeer.api.v1.article.Article;
import de.stroeer.api.v1.article.Article.Builder;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;
import de.stroeer.api.config.*;


class ArticleTest {

  @Test
  public void ModelCompiles() {
    Builder articleBuilder = de.stroeer.api.v1.article.Article.newBuilder();
    articleBuilder.setId("id");
    de.stroeer.api.v1.article.Article article = articleBuilder.build();
    assertEquals(article.getId(), "id");
  }

  @Test
  void ProtoSerializerIsBundled() throws JsonProcessingException {
    ObjectMapper mapper = new ObjectMapper();
    SimpleModule module = new SimpleModule();
    module.addSerializer(Article.class, new ProtobufSerializer());
    mapper.registerModule(module);
    Builder articleBuilder = de.stroeer.api.v1.article.Article.newBuilder();
    articleBuilder.setId("id");
    de.stroeer.api.v1.article.Article article = articleBuilder.build();
    String s = mapper.writeValueAsString(article);
    assertNotNull(s);
  }
}