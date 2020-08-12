import { expect, test } from '@jest/globals';

import WebArticleModule, {
  WebArticle,
  RelatedWebArticle,
} from '../stroeer/web/article/v1/web_article_pb';

test('protobuf generated module of `WebArticle` matches snapshot', () => {
  expect(WebArticleModule).toMatchInlineSnapshot(`
    Object {
      "RelatedWebArticle": [Function],
      "RelatedWebArticleRole": Object {
        "RELATED_WEB_ARTICLE_ROLE_EDITORIAL": 1,
        "RELATED_WEB_ARTICLE_ROLE_UNSPECIFIED": 0,
      },
      "WebArticle": [Function],
    }
  `);
});

test('protobuf generated model of a empty `WebArticle` matches snapshot', () => {
  expect(new WebArticle().toObject()).toMatchInlineSnapshot(`
    Object {
      "authorsList": Array [],
      "bodyList": Array [],
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
  expect(new RelatedWebArticle().toObject()).toMatchInlineSnapshot(`
    Object {
      "article": undefined,
      "role": 0,
    }
  `);
});
