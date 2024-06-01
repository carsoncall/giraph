# what the protobuf needs to

## graph_request
request string = "import" to start
path of project root =
two optional fields
- how many steps
- hash of a specific

else whole tree

## graph_response

list of nodes
list of edges

## nodes

filepath
start byte
end byte
contents
hash

## edges

name for type of relation
two hashes parent and child
