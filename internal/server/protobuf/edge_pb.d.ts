// package: protobuf
// file: edge.proto

import * as jspb from "google-protobuf";

export class Edge extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getParentHash(): string;
  setParentHash(value: string): void;

  getChildHash(): string;
  setChildHash(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Edge.AsObject;
  static toObject(includeInstance: boolean, msg: Edge): Edge.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Edge, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Edge;
  static deserializeBinaryFromReader(message: Edge, reader: jspb.BinaryReader): Edge;
}

export namespace Edge {
  export type AsObject = {
    name: string,
    parentHash: string,
    childHash: string,
  }
}

