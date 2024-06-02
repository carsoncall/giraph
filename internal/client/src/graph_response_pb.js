const jspb = require("google-protobuf");
const Node = require("./node_pb").Node;
const Edge = require("./edge_pb").Edge;

const GraphResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, GraphResponse.repeatedFields_, null);
};
jspb.inherits(GraphResponse, jspb.Message);
GraphResponse.repeatedFields_ = [1, 2];

GraphResponse.prototype.getNodesList = function() {
  return jspb.Message.getRepeatedWrapperField(this, Node, 1);
};

GraphResponse.prototype.setNodesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};

GraphResponse.prototype.addNodes = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, Node, opt_index);
};

GraphResponse.prototype.getEdgesList = function() {
  return jspb.Message.getRepeatedWrapperField(this, Edge, 2);
};

GraphResponse.prototype.setEdgesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};

GraphResponse.prototype.addEdges = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, Edge, opt_index);
};

GraphResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new GraphResponse();
  return GraphResponse.deserializeBinaryFromReader(msg, reader);
};

GraphResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = new Node();
        reader.readMessage(value, Node.deserializeBinaryFromReader);
        msg.addNodes(value);
        break;
      case 2:
        var value = new Edge();
        reader.readMessage(value, Edge.deserializeBinaryFromReader);
        msg.addEdges(value);
        break;
      default:
        reader.skipField();
        break;
    }
  }
  return msg;
};

exports.GraphResponse = GraphResponse;