/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package registry

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/docker/distribution/manifest/schema2"
	registry_error "github.com/vmware/harbor/utils/registry/error"
	"github.com/vmware/harbor/utils/test"
)

var (
	repository = "library/hello-world"
	tag        = "latest"

	mediaType = schema2.MediaTypeManifest
	manifest  = []byte("manifest")

	blob = []byte("blob")

	uuid = "0663ff44-63bb-11e6-8b77-86f30ca893d3"

	digest = "sha256:6c3c624b58dbbcd3c0dd82b4c53f04194d1247c6eebdaab7c610cf7d66709b3b"
)

func TestNewRepositoryWithModifiers(t *testing.T) {
	_, err := NewRepositoryWithModifiers("library/ubuntu",
		"http://registry.org", true, nil)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}
}

func TestBlobExist(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		dgt := path[strings.LastIndex(path, "/")+1 : len(path)]
		if dgt == digest {
			w.Header().Add(http.CanonicalHeaderKey("Content-Length"), strconv.Itoa(len(blob)))
			w.Header().Add(http.CanonicalHeaderKey("Docker-Content-Digest"), digest)
			w.Header().Add(http.CanonicalHeaderKey("Content-Type"), "application/octet-stream")
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "HEAD",
			Pattern: fmt.Sprintf("/v2/%s/blobs/", repository),
			Handler: handler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		err = parseError(err)
		t.Fatalf("failed to create client for repository: %v", err)
	}

	exist, err := client.BlobExist(digest)
	if err != nil {
		t.Fatalf("failed to check the existence of blob: %v", err)
	}

	if !exist {
		t.Errorf("blob should exist on registry, but it does not exist")
	}

	exist, err = client.BlobExist("invalid_digest")
	if err != nil {
		t.Fatalf("failed to check the existence of blob: %v", err)
	}

	if exist {
		t.Errorf("blob should not exist on registry, but it exists")
	}
}

func TestPullBlob(t *testing.T) {
	handler := test.Handler(&test.Response{
		Headers: map[string]string{
			"Content-Length":        strconv.Itoa(len(blob)),
			"Docker-Content-Digest": digest,
			"Content-Type":          "application/octet-stream",
		},
		Body: blob,
	})

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "GET",
			Pattern: fmt.Sprintf("/v2/%s/blobs/%s", repository, digest),
			Handler: handler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	size, reader, err := client.PullBlob(digest)
	if err != nil {
		t.Fatalf("failed to pull blob: %v", err)
	}

	if size != int64(len(blob)) {
		t.Errorf("unexpected size of blob: %d != %d", size, len(blob))
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read from reader: %v", err)
	}

	if bytes.Compare(b, blob) != 0 {
		t.Errorf("unexpected blob: %s != %s", string(b), string(blob))
	}
}

func TestPushBlob(t *testing.T) {
	location := ""
	initUploadHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(http.CanonicalHeaderKey("Content-Length"), "0")
		w.Header().Add(http.CanonicalHeaderKey("Location"), location)
		w.Header().Add(http.CanonicalHeaderKey("Range"), "0-0")
		w.Header().Add(http.CanonicalHeaderKey("Docker-Upload-UUID"), uuid)
		w.WriteHeader(http.StatusAccepted)
	}

	monolithicUploadHandler := test.Handler(&test.Response{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Length":        "0",
			"Location":              fmt.Sprintf("/v2/%s/blobs/%s", repository, digest),
			"Docker-Content-Digest": digest,
		},
	})

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "POST",
			Pattern: fmt.Sprintf("/v2/%s/blobs/uploads/", repository),
			Handler: initUploadHandler,
		},
		&test.RequestHandlerMapping{
			Method:  "PUT",
			Pattern: fmt.Sprintf("/v2/%s/blobs/uploads/%s", repository, uuid),
			Handler: monolithicUploadHandler,
		})
	defer server.Close()
	location = fmt.Sprintf("%s/v2/%s/blobs/uploads/%s", server.URL, repository, uuid)

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	if err = client.PushBlob(digest, int64(len(blob)), bytes.NewReader(blob)); err != nil {
		t.Fatalf("failed to push blob: %v", err)
	}
}

func TestDeleteBlob(t *testing.T) {
	handler := test.Handler(&test.Response{
		StatusCode: http.StatusAccepted,
	})

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "DELETE",
			Pattern: fmt.Sprintf("/v2/%s/blobs/%s", repository, digest),
			Handler: handler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	if err = client.DeleteBlob(digest); err != nil {
		t.Fatalf("failed to delete blob: %v", err)
	}
}

func TestManifestExist(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		tg := path[strings.LastIndex(path, "/")+1 : len(path)]
		if tg == tag {
			w.Header().Add(http.CanonicalHeaderKey("Docker-Content-Digest"), digest)
			w.Header().Add(http.CanonicalHeaderKey("Content-Type"), mediaType)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "HEAD",
			Pattern: fmt.Sprintf("/v2/%s/manifests/%s", repository, tag),
			Handler: handler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	d, exist, err := client.ManifestExist(tag)
	if err != nil {
		t.Fatalf("failed to check the existence of manifest: %v", err)
	}

	if !exist || d != digest {
		t.Errorf("manifest should exist on registry, but it does not exist")
	}

	_, exist, err = client.ManifestExist("invalid_tag")
	if err != nil {
		t.Fatalf("failed to check the existence of manifest: %v", err)
	}

	if exist {
		t.Errorf("manifest should not exist on registry, but it exists")
	}
}

func TestPullManifest(t *testing.T) {
	handler := test.Handler(&test.Response{
		Headers: map[string]string{
			"Docker-Content-Digest": digest,
			"Content-Type":          mediaType,
		},
		Body: manifest,
	})

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "GET",
			Pattern: fmt.Sprintf("/v2/%s/manifests/%s", repository, tag),
			Handler: handler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	d, md, payload, err := client.PullManifest(tag, []string{mediaType})
	if err != nil {
		t.Fatalf("failed to pull manifest: %v", err)
	}

	if d != digest {
		t.Errorf("unexpected digest of manifest: %s != %s", d, digest)
	}

	if md != mediaType {
		t.Errorf("unexpected media type of manifest: %s != %s", md, mediaType)
	}

	if bytes.Compare(payload, manifest) != 0 {
		t.Errorf("unexpected manifest: %s != %s", string(payload), string(manifest))
	}
}

func TestPushManifest(t *testing.T) {
	handler := test.Handler(&test.Response{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Length":        "0",
			"Docker-Content-Digest": digest,
			"Location":              "",
		},
	})

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "PUT",
			Pattern: fmt.Sprintf("/v2/%s/manifests/%s", repository, tag),
			Handler: handler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	d, err := client.PushManifest(tag, mediaType, manifest)
	if err != nil {
		t.Fatalf("failed to pull manifest: %v", err)
	}

	if d != digest {
		t.Errorf("unexpected digest of manifest: %s != %s", d, digest)
	}
}

func TestDeleteTag(t *testing.T) {
	manifestExistHandler := test.Handler(&test.Response{
		Headers: map[string]string{
			"Docker-Content-Digest": digest,
			"Content-Type":          mediaType,
		},
	})

	deleteManifestandler := test.Handler(&test.Response{
		StatusCode: http.StatusAccepted,
	})

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "HEAD",
			Pattern: fmt.Sprintf("/v2/%s/manifests/", repository),
			Handler: manifestExistHandler,
		},
		&test.RequestHandlerMapping{
			Method:  "DELETE",
			Pattern: fmt.Sprintf("/v2/%s/manifests/%s", repository, digest),
			Handler: deleteManifestandler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	if err = client.DeleteTag(tag); err != nil {
		t.Fatalf("failed to delete tag: %v", err)
	}
}

func TestListTag(t *testing.T) {
	handler := test.Handler(&test.Response{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(fmt.Sprintf("{\"name\": \"%s\",\"tags\": [\"%s\"]}", repository, tag)),
	})

	server := test.NewServer(
		&test.RequestHandlerMapping{
			Method:  "GET",
			Pattern: fmt.Sprintf("/v2/%s/tags/list", repository),
			Handler: handler,
		})
	defer server.Close()

	client, err := newRepository(server.URL)
	if err != nil {
		t.Fatalf("failed to create client for repository: %v", err)
	}

	tags, err := client.ListTag()
	if err != nil {
		t.Fatalf("failed to list tags: %v", err)
	}

	if len(tags) != 1 {
		t.Fatalf("unexpected length of tags: %d != %d", len(tags), 1)
	}

	if tags[0] != tag {
		t.Errorf("unexpected tag: %s != %s", tags[0], tag)
	}
}

func TestParseError(t *testing.T) {
	err := &url.Error{
		Err: &registry_error.Error{},
	}
	e := parseError(err)
	if _, ok := e.(*registry_error.Error); !ok {
		t.Errorf("error type does not match registry error")
	}
}

func newRepository(endpoint string) (*Repository, error) {
	return NewRepository(repository, endpoint, &http.Client{})
}
