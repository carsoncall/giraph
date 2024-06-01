// package: protobuf
// file: graph_response.proto

import * as jspb from "google-protobuf";
import * as node_pb from "./node_pb";
import * as edge_pb from "./edge_pb";

export class GraphResponse extends jspb.Message {
  clearNodesList(): void;
  getNodesList(): Array<node_pb.Node>;
  setNodesList(value: Array<node_pb.Node>): void;
  addNodes(value?: node_pb.Node, index?: number): node_pb.Node;

  clearEdgesList(): void;
  getEdgesList(): Array<edge_pb.Edge>;
  setEdgesList(value: Array<edge_pb.Edge>): void;
  addEdges(value?: edge_pb.Edge, index?: number): edge_pb.Edge;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GraphResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GraphResponse): GraphResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GraphResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GraphResponse;
  static deserializeBinaryFromReader(message: GraphResponse, reader: jspb.BinaryReader): GraphResponse;
}

export namespace GraphResponse {
  export type AsObject = {
    nodesList: Array<node_pb.Node.AsObject>,
    edgesList: Array<edge_pb.Edge.AsObject>,
  }
}

