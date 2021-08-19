package de.stroeer.article.v1;

import de.stroeer.page.stage.v1.Stage;
import de.stroeer.page.stage.v1.Stage.Item;
import org.junit.jupiter.api.Test;

public class PageTest {

  @Test
  void streamStage() {
    final Stage.Builder mainStage = Stage.newBuilder();
    final Item mainItem = Item.newBuilder().build();
    final Item companion = Item.newBuilder().build();
    mainStage.addCompanionItems(companion);
    mainStage.addStreamItems(mainItem);
  }
}
