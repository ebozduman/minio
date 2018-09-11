/*
 * Minio Cloud Storage, (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
)

// Data types used for returning dummy tagging XML.
// These variables shouldn't be used elsewhere.
// They are only defined to be used in this file alone.

type tagging struct {
	XMLName xml.Name `xml:"Tagging"`
	TagSet  tagSet   `xml:"TagSet"`
}

type tagSet struct {
	Tag []tagElem `xml:"Tag"`
}

type tagElem struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

// GetBucketWebsite  - GET bucket website, a dummy api
func (api objectAPIHandlers) GetBucketWebsiteHandler(w http.ResponseWriter, r *http.Request) {
}

// GetBucketVersioning - GET bucket versioning, a dummy api
func (api objectAPIHandlers) GetBucketVersioningHandler(w http.ResponseWriter, r *http.Request) {
}

// GetBucketAccelerate  - GET bucket accelerate, a dummy api
func (api objectAPIHandlers) GetBucketAccelerateHandler(w http.ResponseWriter, r *http.Request) {
}

// GetBucketRequestPaymentHandler - GET bucket requestPayment, a dummy api
func (api objectAPIHandlers) GetBucketRequestPaymentHandler(w http.ResponseWriter, r *http.Request) {
}

// GetBucketLoggingHandler - GET bucket logging, a dummy api
func (api objectAPIHandlers) GetBucketLoggingHandler(w http.ResponseWriter, r *http.Request) {
}

// GetBucketLifecycleHandler - GET bucket lifecycle, a dummy api
func (api objectAPIHandlers) GetBucketLifecycleHandler(w http.ResponseWriter, r *http.Request) {
}

// GetBucketReplicationHandler - GET bucket replication, a dummy api
func (api objectAPIHandlers) GetBucketReplicationHandler(w http.ResponseWriter, r *http.Request) {
}

// DeleteBucketTaggingHandler - DELETE bucket tagging, a dummy api
func (api objectAPIHandlers) DeleteBucketTaggingHandler(w http.ResponseWriter, r *http.Request) {
	writeSuccessNoContent(w)
	w.(http.Flusher).Flush()
}

// DeleteBucketWebsiteHandler - DELETE bucket website, a dummy api
func (api objectAPIHandlers) DeleteBucketWebsiteHandler(w http.ResponseWriter, r *http.Request) {
}

// GetBucketCorsHandler - GET bucket cors, a dummy api
func (api objectAPIHandlers) GetBucketCorsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "GetBucketTagging")

	type allowedMethod int
	const (
		GET allowedMethod = iota
		PUT
		HEAD
		POST
		DELETE
	)
	type corsRule struct {
		AllowedMethod allowedMethod `xml:"AllowedMethod"`
		AllowedOrigin string        `xml:"AllowedOrigin"`
		ExposeHeader  string        `xml:"ExposeHeader"`
		ID            string        `xml:"ID"`
		MaxAgeSeconds int64         `xml:"MaxAgeSeconds"`
	}
	type corsConfiguration struct {
		XMLName  xml.Name   `xml:"CORSConfiguration"`
		CorsRule []corsRule `xml:"CORSRule"`
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(w, ErrServerNotInitialized, r.URL)
		return
	}

	// Validate if bucket exists, before proceeding further...
	_, err := objAPI.GetBucketInfo(ctx, bucket)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	tags := &tagging{}
	tags.TagSet.Tag = append(tags.TagSet.Tag, tagElem{
		Key:   "",
		Value: "",
	})

	if err := xml.NewEncoder(w).Encode(tags); err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	w.(http.Flusher).Flush()

}

// GetBucketTaggingHandler - GET bucket tagging, a dummy api
func (api objectAPIHandlers) GetBucketTaggingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "GetBucketTagging")

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(w, ErrServerNotInitialized, r.URL)
		return
	}

	// Validate if bucket exists, before proceeding further...
	_, err := objAPI.GetBucketInfo(ctx, bucket)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	tags := &tagging{}
	tags.TagSet.Tag = append(tags.TagSet.Tag, tagElem{
		Key:   "",
		Value: "",
	})

	if err := xml.NewEncoder(w).Encode(tags); err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	w.(http.Flusher).Flush()
}

// GetObjectTaggingHandler - GET object tagging, a dummy api
func (api objectAPIHandlers) GetObjectTaggingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "GetObjectTagging")

	vars := mux.Vars(r)
	bucket := vars["bucket"]
	object := vars["object"]

	objAPI := api.ObjectAPI()
	if objAPI == nil {
		writeErrorResponse(w, ErrServerNotInitialized, r.URL)
		return
	}

	// Validate if object exists, before proceeding further...
	_, err := objAPI.GetObjectInfo(ctx, bucket, object)
	if err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	tags := &tagging{}
	tags.TagSet.Tag = append(tags.TagSet.Tag, tagElem{
		Key:   "",
		Value: "",
	})

	if err := xml.NewEncoder(w).Encode(tags); err != nil {
		writeErrorResponse(w, toAPIErrorCode(err), r.URL)
		return
	}

	w.(http.Flusher).Flush()
}
