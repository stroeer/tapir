// package: 
// file: service.proto

import * as jspb from "google-protobuf";
import * as content_pb from "./content_pb";

export class ArticleRequest extends jspb.Message {
  getArticleid(): string;
  setArticleid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ArticleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ArticleRequest): ArticleRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ArticleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ArticleRequest;
  static deserializeBinaryFromReader(message: ArticleRequest, reader: jspb.BinaryReader): ArticleRequest;
}

export namespace ArticleRequest {
  export type AsObject = {
    articleid: string,
  }
}

export class ArticleResponse extends jspb.Message {
  clearResultsList(): void;
  getResultsList(): Array<content_pb.Article>;
  setResultsList(value: Array<content_pb.Article>): void;
  addResults(value?: content_pb.Article, index?: number): content_pb.Article;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ArticleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ArticleResponse): ArticleResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ArticleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ArticleResponse;
  static deserializeBinaryFromReader(message: ArticleResponse, reader: jspb.BinaryReader): ArticleResponse;
}

export namespace ArticleResponse {
  export type AsObject = {
    resultsList: Array<content_pb.Article.AsObject>,
  }
}

