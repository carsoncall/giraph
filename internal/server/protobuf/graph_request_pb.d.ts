// package: protobuf
// file: graph_request.proto

import * as jspb from "google-protobuf";

export class GraphRequest extends jspb.Message {
  getRequest(): string;
  setRequest(value: string): void;

  getProjectRoot(): string;
  setProjectRoot(value: string): void;

  hasNumSteps(): boolean;
  clearNumSteps(): void;
  getNumSteps(): number;
  setNumSteps(value: number): void;

  hasStartNodeHash(): boolean;
  clearStartNodeHash(): void;
  getStartNodeHash(): string;
  setStartNodeHash(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GraphRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GraphRequest): GraphRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GraphRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GraphRequest;
  static deserializeBinaryFromReader(message: GraphRequest, reader: jspb.BinaryReader): GraphRequest;
}

export namespace GraphRequest {
  export type AsObject = {
    request: string,
    projectRoot: string,
    numSteps: number,
    startNodeHash: string,
  }
}

