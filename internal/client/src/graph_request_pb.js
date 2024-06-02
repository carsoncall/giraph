const jspb = require("google-protobuf");

const GraphRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
jspb.inherits(GraphRequest, jspb.Message);

GraphRequest.prototype.getRequest = function() {
  return jspb.Message.getFieldWithDefault(this, 1, "");
};

GraphRequest.prototype.setRequest = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};

GraphRequest.prototype.getProjectRoot = function() {
  return jspb.Message.getFieldWithDefault(this, 2, "");
};

GraphRequest.prototype.setProjectRoot = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};

GraphRequest.prototype.getNumSteps = function() {
  return jspb.Message.getFieldWithDefault(this, 3, 0);
};

GraphRequest.prototype.setNumSteps = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};

GraphRequest.prototype.getStartNodeHash = function() {
  return jspb.Message.getFieldWithDefault(this, 4, "");
};

GraphRequest.prototype.setStartNodeHash = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};

GraphRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  GraphRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};

GraphRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequest();
  if (f.length > 0) {
    writer.writeString(1, f);
  }
  f = message.getProjectRoot();
  if (f.length > 0) {
    writer.writeString(2, f);
  }
  f = message.getNumSteps();
  if (f !== 0) {
    writer.writeInt32(3, f);
  }
  f = message.getStartNodeHash();
  if (f.length > 0) {
    writer.writeString(4, f);
  }
};

exports.GraphRequest = GraphRequest;