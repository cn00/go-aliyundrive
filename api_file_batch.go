/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package aliyundrive

import (
	"context"
	"net/http"
)

type BatchReqItem struct {
	Id, // FileID
	Method, // POST
	Url string // "/recyclebin/trash"
	Body    GetFileReq //struct{ DriveID, FileID string }
	Headers struct {
		ContentType string `json:"Content-Type"`
	}
}

const (
	BatchTrash  = "/recyclebin/trash"
	BatchDelete = "/file/delete"
)

// batch
func (r *FileService) Batch(ctx context.Context, url, driveID string, fileIds []string) (resp *batchResp, err error) {
	request := BatchReq{Resource: "file"}
	for _, i := range fileIds {
		request.Requests = append(
			request.Requests,
			BatchReqItem{
				Id:     i,
				Method: http.MethodPost,
				Url:    url,
				Body:   GetFileReq{DriveID: driveID, FileID: i},
			},
		)
	}
	req := &RawRequestReq{
		Scope:  "File",
		API:    "DeleteFile",
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/v3/batch",
		Body:   request,
	}
	resp = new(batchResp)

	_, err = r.cli.RawRequest(ctx, req, resp)
	return
}

type BatchReq struct {
	Requests []BatchReqItem
	Resource string // file
}

type batchResp struct {
	Responses []struct {
		Id     string
		Status int64
	}
}
