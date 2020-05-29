import { expect, test } from '@jest/globals';

import WebArticleModule, {
  WebArticle,
  RelatedArticle,
} from '../stroeer/web/article/v1/web_article_pb';

test('protobuf generated module of `WebArticle` matches snapshot', () => {
  expect(WebArticleModule).toMatchInlineSnapshot(`
    Object {
      "RelatedArticle": [Function],
      "RelatedArticleRole": Object {
        "RELATED_ARTICLE_ROLE_EDITORIAL": 1,
        "RELATED_ARTICLE_ROLE_UNSPECIFIED": 0,
      },
      "WebArticle": [Function],
    }
  `);
});

test('protobuf generated model of a empty `WebArticle` matches snapshot', () => {
  expect(new WebArticle().toObject()).toMatchInlineSnapshot(`
    Object {
      "authorsList": Array [],
      "canonicalUrl": "",
      "elementsList": Array [],
      "fieldsMap": Array [],
      "headline": "",
      "id": "",
      "sectionPath": "",
      "type": 0,
      "webUrl": "",
    }
  `);
});

test('protobuf generated model of a empty `RelatedArticle` matches snapshot', () => {
  expect(new RelatedArticle().toObject()).toMatchInlineSnapshot(`
    Object {
      "article": undefined,
      "role": 0,
    }
  `);
});
