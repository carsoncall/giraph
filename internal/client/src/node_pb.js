const jspb = require("google-protobuf");

const Node = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
jspb.inherits(Node, jspb.Message);

Node.prototype.getFilepath = function() {
  return jspb.Message.getFieldWithDefault(this, 1, "");
};

Node.prototype.setFilepath = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};

Node.prototype.getStartByte = function() {
  return jspb.Message.getFieldWithDefault(this, 2, 0);
};

Node.prototype.setStartByte = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};

Node.prototype.getEndByte = function() {
  return jspb.Message.getFieldWithDefault(this, 3, 0);
};

Node.prototype.setEndByte = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};

Node.prototype.getContents = function() {
  return jspb.Message.getFieldWithDefault(this, 4, "");
};

Node.prototype.setContents = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};

Node.prototype.getHash = function() {
  return jspb.Message.getFieldWithDefault(this, 5, "");
};

Node.prototype.setHash = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};

Node.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = reader.readString();
        msg.setFilepath(value);
        break;
      case 2:
        var value = reader.readInt64();
        msg.setStartByte(value);
        break;
      case 3:
        var value = reader.readInt64();
        msg.setEndByte(value);
        break;
      case 4:
        var value = reader.readString();
        msg.setContents(value);
        break;
      case 5:
        var value = reader.readString();
        msg.setHash(value);
        break;
      default:
        reader.skipField();
        break;
    }
  }
  return msg;
};

exports.Node = Node;