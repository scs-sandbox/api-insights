// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/cisco-developer/api-insights/api/pkg/utils"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/emicklei/go-restful/v3"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

const (
	SpecDiffTableName = "spec_diffs"
)

// SpecDiff represents a specDiff
type SpecDiff struct {
	ID string `json:"id,omitempty" gorm:"column:id;primaryKey"`
	*SpecDiffRequest

	SpecDiffResult

	ServiceID string    `json:"service_id,omitempty" gorm:"column:service_id;index"`
	Status    string    `json:"status,omitempty" gorm:"column:status;index"` // Submitted, Invalid, Diffed
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// BeforeSave is a hook called before creation by GORM (https://gorm.io/docs/hooks.html).
// For handling large SpecDiffResult.RawResult(s), compressData conditionally compresses SpecDiffResult.RawResult into SpecDiffResult.RawResultCompressed.
func (m *SpecDiff) BeforeSave(tx *gorm.DB) (err error) {
	if m.Config != nil {
		if m.RawConfig, err = json.Marshal(m.Config); err != nil {
			return err
		}
	}
	if m.Result != nil {
		if m.RawResult, err = json.Marshal(m.Result); err != nil {
			return err
		}
	}

	// Handle compression.
	m.RawResultCompressed, err = compressData(m.RawResult)
	if err != nil {
		return err
	} else if m.RawResultCompressed != nil {
		m.internalRawResult = m.RawResult
		m.RawResult = nil
	}

	return
}

// AfterSave is a hook called after creation by GORM (https://gorm.io/docs/hooks.html).
// For handling large SpecDiff.SpecDiffResult(s), resets the temporary staging of SpecDiff.SpecDiffResult.
func (m *SpecDiff) AfterSave(tx *gorm.DB) (err error) {
	if m.RawResultCompressed != nil {
		m.RawResult = m.internalRawResult
		m.internalRawResult = nil
		m.RawResultCompressed = nil
	}
	return
}

// AfterFind is a hook called after querying by GORM (https://gorm.io/docs/hooks.html).
// For handling large SpecDiff.SpecDiffResult(s), if SpecDiffResult.RawResultCompressed contains the compression, decompresses it back into SpecDiffResult.RawResult.
func (m *SpecDiff) AfterFind(tx *gorm.DB) (err error) {
	if m.RawConfig != nil {
		m.Config = &diff.Config{}
		if err = json.Unmarshal(m.RawConfig, m.Config); err != nil {
			return err
		}
	}
	if m.RawResultCompressed != nil {
		m.RawResult, _, err = utils.GUNZIP(m.RawResultCompressed)
		if err != nil {
			return err
		}
	}
	if m.RawResult != nil {
		m.Result = &diff.Result{}
		if err = json.Unmarshal(m.RawResult, m.Result); err != nil {
			return err
		}
	}
	return
}

// TableName implements gorm Tabler interface
func (m *SpecDiff) TableName() string {
	return SpecDiffTableName
}

func (m *SpecDiff) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = shared.TimeUUID()
	return
}

// GetID returns the ID of specDiff object
func (m *SpecDiff) GetID() string {
	return fmt.Sprintf("%v", m.ID)
}

// GetTags returns all the tags
func (m *SpecDiff) GetTags() []string {
	tags := make([]string, 0, 10)
	tags = append(tags, m.ServiceID)
	tags = append(tags, m.OldSpecID)
	tags = append(tags, m.NewSpecID)
	tags = append(tags, m.Status)
	return tags
}

// String returns the text representation of specDiff object
func (m *SpecDiff) String() string {
	return fmt.Sprintf("%v", *m)
}

// GetIndex returns an index for specific field
func (m *SpecDiff) GetIndex(field string) string {
	return m.GetIndexes()[field]
}

// GetIndexes returns all the field indexes
func (m *SpecDiff) GetIndexes() map[string]string {
	return map[string]string{
		"service_id":  "idx_service_id",
		"old_spec_id": "idx_old_spec_id",
		"new_spec_id": "idx_new_spec_id",
		"status":      "idx_status",
	}
}

// GetIndexValue return index value for specified field
func (m *SpecDiff) GetIndexValue(field string) string {
	return m.GetIndexValues()[field]
}

// GetIndexValues return all field index values
func (m *SpecDiff) GetIndexValues() map[string]string {
	return map[string]string{
		"service_id":  m.ServiceID,
		"old_spec_id": m.OldSpecID,
		"new_spec_id": m.NewSpecID,
		"status":      m.Status,
	}
}

// Sortable checks if field is sortable.
func (m *SpecDiff) Sortable(field string) bool {
	_, found := m.SortableFields()[field]
	return found
}

// SortableFields returns all sortable fields
func (m *SpecDiff) SortableFields() map[string]struct{} {
	return map[string]struct{}{
		"created_at": {},
		"updated_at": {},
	}
}

func (m *SpecDiff) SetResult(result *diff.Result, status string) error {
	if result == nil {
		return fmt.Errorf("spec_diff: cannot set nil result")
	}
	if status == "" {
		return fmt.Errorf("spec_diff: cannot set empty status")
	}
	m.Result = result
	m.Status = status
	return nil
}

// SpecDiffResponse wrappers specDiff response
type SpecDiffResponse struct {
	Pagination
	Data []SpecDiff `json:"data"`
}

type SpecDiffResult struct {
	Result              *diff.Result   `json:"result" gorm:"-"`
	RawResult           datatypes.JSON `json:"-" gorm:"column:result"`
	RawResultCompressed []byte         `json:"-" gorm:"column:result_compressed"`
	// internalRawResult is an internal state variable for temporarily storing SpecDiffResult.RawResult between SpecDiff.BeforeSave & SpecDiff.AfterSave for data compression.
	internalRawResult datatypes.JSON
}

type SpecDiffRequest struct {
	NewSpecID string `json:"new_spec_id,omitempty" gorm:"column:new_spec_id;index"`
	OldSpecID string `json:"old_spec_id,omitempty" gorm:"column:old_spec_id;index"`

	OldSpecDoc SpecDoc `json:"old_spec_doc,omitempty" gorm:"-"`
	NewSpecDoc SpecDoc `json:"new_spec_doc,omitempty" gorm:"-"`

	SpecDiffConfig
}

func (r *SpecDiffRequest) Compare(with *SpecDiffRequest) bool {
	if r.NewSpecID != with.NewSpecID ||
		r.OldSpecID != with.OldSpecID ||
		((r.Config == nil) != (with.Config == nil)) {
		return false
	}
	if r.Config != nil && with.Config != nil {
		if r.Config.OutputFormat != with.Config.OutputFormat {
			return false
		}
	}
	return true
}

func (r *SpecDiffRequest) tryAsMultipartForm(req *restful.Request) (isMultipart bool, err error) {
	const (
		oldSpecFilename = "old_spec_file"
		newSpecFilename = "new_spec_file"
	)

	of, _, err := req.Request.FormFile(oldSpecFilename)
	if err != nil {
		return err != http.ErrNotMultipart, err
	}
	oldSpecData, err := io.ReadAll(of)
	if err != nil {
		return true, err
	}
	r.OldSpecDoc = NewSpecDocFromBytes(oldSpecData)

	nf, _, err := req.Request.FormFile(newSpecFilename)
	if err != nil {
		return err != http.ErrNotMultipart, err
	}
	newSpecData, err := io.ReadAll(nf)
	if err != nil {
		return true, err
	}
	r.NewSpecDoc = NewSpecDocFromBytes(newSpecData)

	defer func() {
		_ = of.Close()
		_ = nf.Close()
		_ = req.Request.MultipartForm.RemoveAll()
	}()

	return true, nil
}

func (r *SpecDiffRequest) From(req *restful.Request, specsGetter func(oldSpecID, newSpecID string) (oldSpec *Spec, newSpec *Spec, err error)) error {
	// Try handling request as multipart/form-data.
	if isMultipart, err := r.tryAsMultipartForm(req); isMultipart {
		if err != nil && err != http.ErrNotMultipart {
			return err
		}
		return nil
	}

	// Try handling request as normal entity request body.
	if err := req.ReadEntity(r); err != nil {
		return err
	}

	// Spec docs already provided.
	if hasOld, hasNew := r.HasSpecDocs(); hasOld && hasNew {
		return nil
	}

	// Last resort. Fetch specs by ID.
	oldSpec, newSpec, err := specsGetter(r.OldSpecID, r.NewSpecID)
	if err != nil {
		return err
	}

	r.OldSpecDoc = oldSpec.Doc
	r.NewSpecDoc = newSpec.Doc

	return nil
}

func (r *SpecDiffRequest) HasSpecDocs() (hasOldSpecDoc, hasNewSpecDoc bool) {
	var hasSpecDoc = func(specDoc SpecDoc) bool {
		return specDoc != nil && *specDoc != ""
	}
	return hasSpecDoc(r.OldSpecDoc), hasSpecDoc(r.NewSpecDoc)
}

type SpecDiffConfig struct {
	Config    *diff.Config   `json:"config,omitempty" gorm:"-"`
	RawConfig datatypes.JSON `json:"-" gorm:"column:config"`
}
