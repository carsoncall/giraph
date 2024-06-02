const grpcWeb = require('grpc-web');
const { GraphRequest } = require('./src/graph_request_pb');
const { GraphResponse } = require('./src/graph_response_pb');
const { Node } = require('./src/node_pb');
const { Edge } = require('./src/edge_pb');

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var graph_request_pb_1 = require("./src/graph_request_pb");
var graph_response_pb_1 = require("./src/graph_response_pb");
var request = new graph_request_pb_1.GraphRequest();
request.setRequest('Test Request');
request.setProjectRoot('/path/to/project');
request.setNumSteps(5);
request.setStartNodeHash('start_hash');
var url = 'http://localhost:8080/graph';
var requestData = request.serializeBinary();
fetch(url, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: 'request=' + encodeURIComponent(new TextDecoder().decode(requestData)),
})
    .then(function (response) { return response.arrayBuffer(); })
    .then(function (arrayBuffer) {
    var response = graph_response_pb_1.GraphResponse.deserializeBinary(new Uint8Array(arrayBuffer));
    var nodes = response.getNodesList();
    var edges = response.getEdgesList();
    console.log('Nodes:');
    nodes.forEach(function (node) {
        console.log("  Filepath: ".concat(node.getFilepath()));
        console.log("  StartByte: ".concat(node.getStartByte()));
        console.log("  EndByte: ".concat(node.getEndByte()));
        console.log("  Contents: ".concat(node.getContents()));
        console.log("  Hash: ".concat(node.getHash()));
    });
    console.log('Edges:');
    edges.forEach(function (edge) {
        console.log("  Name: ".concat(edge.getName()));
        console.log("  ParentHash: ".concat(edge.getParentHash()));
        console.log("  ChildHash: ".concat(edge.getChildHash()));
    });
})
    .catch(function (error) {
    console.error('Error:', error);
});
