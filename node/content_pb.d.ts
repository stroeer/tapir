// package: 
// file: content.proto

import * as jspb from "google-protobuf";

export class Article extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getWebUrl(): string;
  setWebUrl(value: string): void;

  getCanonicalUrl(): string;
  setCanonicalUrl(value: string): void;

  getType(): ContentTypeMap[keyof ContentTypeMap];
  setType(value: ContentTypeMap[keyof ContentTypeMap]): void;

  getSectionPath(): string;
  setSectionPath(value: string): void;

  getHeadline(): string;
  setHeadline(value: string): void;

  getFieldsMap(): jspb.Map<string, string>;
  clearFieldsMap(): void;
  clearKeywordsList(): void;
  getKeywordsList(): Array<string>;
  setKeywordsList(value: Array<string>): void;
  addKeywords(value: string, index?: number): string;

  hasMetadata(): boolean;
  clearMetadata(): void;
  getMetadata(): Metadata | undefined;
  setMetadata(value?: Metadata): void;

  clearElementsList(): void;
  getElementsList(): Array<Element>;
  setElementsList(value: Array<Element>): void;
  addElements(value?: Element, index?: number): Element;

  clearAuthorsList(): void;
  getAuthorsList(): Array<Author>;
  setAuthorsList(value: Array<Author>): void;
  addAuthors(value?: Author, index?: number): Author;

  clearOnwardList(): void;
  getOnwardList(): Array<Onward>;
  setOnwardList(value: Array<Onward>): void;
  addOnward(value?: Onward, index?: number): Onward;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Article.AsObject;
  static toObject(includeInstance: boolean, msg: Article): Article.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Article, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Article;
  static deserializeBinaryFromReader(message: Article, reader: jspb.BinaryReader): Article;
}

export namespace Article {
  export type AsObject = {
    id: string,
    webUrl: string,
    canonicalUrl: string,
    type: ContentTypeMap[keyof ContentTypeMap],
    sectionPath: string,
    headline: string,
    fieldsMap: Array<[string, string]>,
    keywordsList: Array<string>,
    metadata?: Metadata.AsObject,
    elementsList: Array<Element.AsObject>,
    authorsList: Array<Author.AsObject>,
    onwardList: Array<Onward.AsObject>,
  }
}

export class Metadata extends jspb.Message {
  getValidFromDate(): string;
  setValidFromDate(value: string): void;

  getValidToDate(): string;
  setValidToDate(value: string): void;

  getState(): string;
  setState(value: string): void;

  getTransformationDate(): string;
  setTransformationDate(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Metadata.AsObject;
  static toObject(includeInstance: boolean, msg: Metadata): Metadata.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Metadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Metadata;
  static deserializeBinaryFromReader(message: Metadata, reader: jspb.BinaryReader): Metadata;
}

export namespace Metadata {
  export type AsObject = {
    validFromDate: string,
    validToDate: string,
    state: string,
    transformationDate: string,
  }
}

export class Element extends jspb.Message {
  getType(): ContentTypeMap[keyof ContentTypeMap];
  setType(value: ContentTypeMap[keyof ContentTypeMap]): void;

  getGuid(): string;
  setGuid(value: string): void;

  clearRelationsList(): void;
  getRelationsList(): Array<string>;
  setRelationsList(value: Array<string>): void;
  addRelations(value: string, index?: number): string;

  clearAssetsList(): void;
  getAssetsList(): Array<Asset>;
  setAssetsList(value: Array<Asset>): void;
  addAssets(value?: Asset, index?: number): Asset;

  getContentId(): string;
  setContentId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Element.AsObject;
  static toObject(includeInstance: boolean, msg: Element): Element.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Element, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Element;
  static deserializeBinaryFromReader(message: Element, reader: jspb.BinaryReader): Element;
}

export namespace Element {
  export type AsObject = {
    type: ContentTypeMap[keyof ContentTypeMap],
    guid: string,
    relationsList: Array<string>,
    assetsList: Array<Asset.AsObject>,
    contentId: string,
  }
}

export class Asset extends jspb.Message {
  getType(): AssetTypeMap[keyof AssetTypeMap];
  setType(value: AssetTypeMap[keyof AssetTypeMap]): void;

  getIndex(): number;
  setIndex(value: number): void;

  clearRelationsList(): void;
  getRelationsList(): Array<string>;
  setRelationsList(value: Array<string>): void;
  addRelations(value: string, index?: number): string;

  getFieldsMap(): jspb.Map<string, string>;
  clearFieldsMap(): void;
  hasMetadata(): boolean;
  clearMetadata(): void;
  getMetadata(): Metadata | undefined;
  setMetadata(value?: Metadata): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Asset.AsObject;
  static toObject(includeInstance: boolean, msg: Asset): Asset.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Asset, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Asset;
  static deserializeBinaryFromReader(message: Asset, reader: jspb.BinaryReader): Asset;
}

export namespace Asset {
  export type AsObject = {
    type: AssetTypeMap[keyof AssetTypeMap],
    index: number,
    relationsList: Array<string>,
    fieldsMap: Array<[string, string]>,
    metadata?: Metadata.AsObject,
  }
}

export class Author extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getAlias(): string;
  setAlias(value: string): void;

  getId(): string;
  setId(value: string): void;

  getImageUrl(): string;
  setImageUrl(value: string): void;

  getLandingPage(): string;
  setLandingPage(value: string): void;

  getFirstname(): string;
  setFirstname(value: string): void;

  getSurname(): string;
  setSurname(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Author.AsObject;
  static toObject(includeInstance: boolean, msg: Author): Author.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Author, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Author;
  static deserializeBinaryFromReader(message: Author, reader: jspb.BinaryReader): Author;
}

export namespace Author {
  export type AsObject = {
    name: string,
    alias: string,
    id: string,
    imageUrl: string,
    landingPage: string,
    firstname: string,
    surname: string,
  }
}

export class Onward extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  clearRolesList(): void;
  getRolesList(): Array<string>;
  setRolesList(value: Array<string>): void;
  addRoles(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Onward.AsObject;
  static toObject(includeInstance: boolean, msg: Onward): Onward.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Onward, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Onward;
  static deserializeBinaryFromReader(message: Onward, reader: jspb.BinaryReader): Onward;
}

export namespace Onward {
  export type AsObject = {
    id: string,
    rolesList: Array<string>,
  }
}

export interface ContentTypeMap {
  ARTICLE_CONTENT: 0;
  CODE_CONTENT: 1;
  EXTERNAL_CONTENT: 2;
  GALLERY_CONTENT: 3;
  IMAGE_CONTENT: 4;
  OEMBED_CONTENT: 5;
  PAGE_CONTENT: 6;
  TEASER_CONTENT: 7;
  UNDEFINED_CONTENT: 8;
  VIDEO_CONTENT: 9;
}

export const ContentType: ContentTypeMap;

export interface AssetTypeMap {
  IMAGE_ASSET: 0;
  VIDEO_ASSET: 1;
  IMAGES_ASSET: 2;
  URL_ASSET: 3;
  METADATA_ASSET: 4;
  UNDEFINED_ASSET: 5;
}

export const AssetType: AssetTypeMap;

