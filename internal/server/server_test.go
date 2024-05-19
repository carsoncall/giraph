// internal/server/server_test.go

package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carsoncall/giraph/internal/server"

	"github.com/carsoncall/giraph/internal/server/protobuf"
	"github.com/golang/protobuf/proto"
)

func TestGraphHandler(t *testing.T) {
	req := &protobuf.GraphRequest{
		NodeNames: []string{"node1", "node2"},
	}
	data, err := proto.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	r, err := http.NewRequest("POST", "/graph", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	server.GraphHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}

	expectedContentType := "application/protobuf"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("expected content type %q; got %q", expectedContentType, contentType)
	}

	var resp protobuf.GraphResponse
	if err := proto.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	expectedNodes := []*protobuf.Node{
		{Name: "node1"},
		{Name: "node2"},
	}
	if len(resp.Nodes) != len(expectedNodes) {
		t.Errorf("expected %d nodes; got %d", len(expectedNodes), len(resp.Nodes))
	}
	for i, node := range resp.Nodes {
		if node.Name != expectedNodes[i].Name {
			t.Errorf("expected node name %q; got %q", expectedNodes[i].Name, node.Name)
		}
	}

}
