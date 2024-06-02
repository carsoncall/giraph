const jspb = require("google-protobuf");

const Edge = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
jspb.inherits(Edge, jspb.Message);

Edge.prototype.getName = function() {
  return jspb.Message.getFieldWithDefault(this, 1, "");
};

Edge.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};

Edge.prototype.getParentHash = function() {
  return jspb.Message.getFieldWithDefault(this, 2, "");
};

Edge.prototype.setParentHash = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};

Edge.prototype.getChildHash = function() {
  return jspb.Message.getFieldWithDefault(this, 3, "");
};

Edge.prototype.setChildHash = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};

Edge.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = reader.readString();
        msg.setName(value);
        break;
      case 2:
        var value = reader.readString();
        msg.setParentHash(value);
        break;
      case 3:
        var value = reader.readString();
        msg.setChildHash(value);
        break;
      default:
        reader.skipField();
        break;
    }
  }
  return msg;
};

exports.Edge = Edge;