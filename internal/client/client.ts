import * as grpcWeb from 'grpc-web';
import { GraphRequest } from './src/graph_request_pb';
import { GraphResponse } from './src/graph_response_pb';
import { Node } from './src/node_pb';
import { Edge } from './src/edge_pb';

const request = new GraphRequest();
request.setRequest('Test Request');
request.setProjectRoot('/path/to/project');
request.setNumSteps(5);
request.setStartNodeHash('start_hash');
const url = 'http://localhost:8080/graph';
const requestData = request.serializeBinary();

fetch(url, {
  method: 'POST',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
  },
  body: 'request=' + encodeURIComponent(new TextDecoder().decode(requestData)),
})
  .then((response) => response.arrayBuffer())
  .then((arrayBuffer) => {
    const response = GraphResponse.deserializeBinary(new Uint8Array(arrayBuffer));

    const nodes = response.getNodesList();
    const edges = response.getEdgesList();

    console.log('Nodes:');
    nodes.forEach((node: Node) => {
      console.log(`  Filepath: ${node.getFilepath()}`);
      console.log(`  StartByte: ${node.getStartByte()}`);
      console.log(`  EndByte: ${node.getEndByte()}`);
      console.log(`  Contents: ${node.getContents()}`);
      console.log(`  Hash: ${node.getHash()}`);
    });

    console.log('Edges:');
    edges.forEach((edge: Edge) => {
      console.log(`  Name: ${edge.getName()}`);
      console.log(`  ParentHash: ${edge.getParentHash()}`);
      console.log(`  ChildHash: ${edge.getChildHash()}`);
    });
  })
  .catch((error) => {
    console.error('Error:', error);
  });