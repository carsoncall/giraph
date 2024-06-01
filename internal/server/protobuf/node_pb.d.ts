// package: protobuf
// file: node.proto

import * as jspb from "google-protobuf";

export class Node extends jspb.Message {
  getFilepath(): string;
  setFilepath(value: string): void;

  getStartByte(): number;
  setStartByte(value: number): void;

  getEndByte(): number;
  setEndByte(value: number): void;

  getContents(): string;
  setContents(value: string): void;

  getHash(): string;
  setHash(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Node.AsObject;
  static toObject(includeInstance: boolean, msg: Node): Node.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Node, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Node;
  static deserializeBinaryFromReader(message: Node, reader: jspb.BinaryReader): Node;
}

export namespace Node {
  export type AsObject = {
    filepath: string,
    startByte: number,
    endByte: number,
    contents: string,
    hash: string,
  }
}

