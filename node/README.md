# @stroeer/tapir-v1

## Usage

```bash
npm i @stroeer/tapir-v1
# or
yarn add @stroeer/tapir-v1
```

## Development

Make sure you generated the gRPC models and services for node.

### node

```bash
# https://github.com/nvm-sh/nvm#installing-and-updating
brew install nvm

# add this to your shell/bash profile:
# source /usr/local/opt/nvm/nvm.sh

nvm install 14
nvm use 14
# optional
nvm alias system 14
```

### install

```bash
npm i
```

### Test/Lint

```bash
npm run test
npm run lint
```

## How to update jest snapshots

```bash
npm run test -- --u
```

## Helpful resources

Examples for generating TypeScript types of generated services / models.

https://github.com/badsyntax/grpc-js-typescript

## Example Script

A very basic example script to call the article page via the given node module and log the plain text to console. The `GRPC_SERVICE` variable needs to be set via environment variables or replaced in code to get this snippet to work.

```js
import ArticlePageService from '@stroeer/tapir-v1/stroeer/page/article/v1/article_page_service_pb.js';
import { ArticlePageServiceClient } from '@stroeer/tapir-v1/stroeer/page/article/v1/article_page_service_grpc_pb.js';
import { ChannelCredentials } from '@grpc/grpc-js';

// Client
const credentials = ChannelCredentials.createSsl();
const client = new ArticlePageServiceClient(process.env.GRPC_SERVICE, credentials);

// Request
const request = new ArticlePageService.GetArticlePageRequest();
// Set Article Id
request.setId(87971076);

// Content
const page = client.getArticlePage(request, (err, response) => {
	const page = response.getArticlePage();
	const article = page.getArticle();

	let articleText = '';

	function handleChild(child) {
		const children = child.getChildrenList();
		if (children.length > 0) {
			children.forEach(child => handleChild(child));
		} else {
			articleText += child.getText();
		}
	}

	const bodyNodes = article.getBody().getChildrenList();

	bodyNodes.forEach(handleChild);

	console.log(articleText);
});
```
