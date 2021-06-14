import os
import pprint

import grpc

from stroeer.page.article.v1.article_page_service_pb2 import \
  GetArticlePageRequest, GetArticlePageResponse
from stroeer.page.article.v1.article_page_pb2 import \
  ArticlePage, RelatedArticle
from stroeer.page.article.v1.article_page_service_pb2_grpc import \
  ArticlePageServiceStub
from stroeer.core.v1.article_pb2 import Article
from typing import List

if __name__ == '__main__':
  credentials = grpc.ssl_channel_credentials()

  if not os.getenv('GRPC_ENDPOINT'):
    print('\033[91mPlease provide a valid endpoint via env.GRPC_ENDPOINT.\033[0m')
    exit(1)

  with grpc.secure_channel(os.getenv('GRPC_ENDPOINT'), credentials) as channel:
    stub: ArticlePageServiceStub = ArticlePageServiceStub(channel)
    response: GetArticlePageResponse = stub.GetArticlePage(GetArticlePageRequest(id=87971076))

    article_page: ArticlePage = response.article_page

    main_article: Article = article_page.article
    related_articles: List[RelatedArticle] = article_page.related_articles

    print(f"""
main article
============
{pprint.pformat(main_article)}
    """)

    print(f"""
related articles
================
{pprint.pformat(related_articles)}
    """)
