package contentmoderatorapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/test-repo-arcturus/azure-sdk-for-go/services/cognitiveservices/v1.0/contentmoderator"
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// ImageModerationClientAPI contains the set of methods on the ImageModerationClient type.
type ImageModerationClientAPI interface {
	EvaluateFileInput(ctx context.Context, imageStream io.ReadCloser, cacheImage *bool) (result contentmoderator.Evaluate, err error)
	EvaluateMethod(ctx context.Context, cacheImage *bool) (result contentmoderator.Evaluate, err error)
	EvaluateURLInput(ctx context.Context, contentType string, imageURL contentmoderator.ImageURL, cacheImage *bool) (result contentmoderator.Evaluate, err error)
	FindFaces(ctx context.Context, cacheImage *bool) (result contentmoderator.FoundFaces, err error)
	FindFacesFileInput(ctx context.Context, imageStream io.ReadCloser, cacheImage *bool) (result contentmoderator.FoundFaces, err error)
	FindFacesURLInput(ctx context.Context, contentType string, imageURL contentmoderator.ImageURL, cacheImage *bool) (result contentmoderator.FoundFaces, err error)
	MatchFileInput(ctx context.Context, imageStream io.ReadCloser, listID string, cacheImage *bool) (result contentmoderator.MatchResponse, err error)
	MatchMethod(ctx context.Context, listID string, cacheImage *bool) (result contentmoderator.MatchResponse, err error)
	MatchURLInput(ctx context.Context, contentType string, imageURL contentmoderator.ImageURL, listID string, cacheImage *bool) (result contentmoderator.MatchResponse, err error)
	OCRFileInput(ctx context.Context, language string, imageStream io.ReadCloser, cacheImage *bool, enhanced *bool) (result contentmoderator.OCR, err error)
	OCRMethod(ctx context.Context, language string, cacheImage *bool, enhanced *bool) (result contentmoderator.OCR, err error)
	OCRURLInput(ctx context.Context, language string, contentType string, imageURL contentmoderator.ImageURL, cacheImage *bool, enhanced *bool) (result contentmoderator.OCR, err error)
}

var _ ImageModerationClientAPI = (*contentmoderator.ImageModerationClient)(nil)

// TextModerationClientAPI contains the set of methods on the TextModerationClient type.
type TextModerationClientAPI interface {
	DetectLanguage(ctx context.Context, textContentType string, textContent io.ReadCloser) (result contentmoderator.DetectedLanguage, err error)
	ScreenText(ctx context.Context, textContentType string, textContent io.ReadCloser, language string, autocorrect *bool, pii *bool, listID string, classify *bool) (result contentmoderator.Screen, err error)
}

var _ TextModerationClientAPI = (*contentmoderator.TextModerationClient)(nil)

// ListManagementImageListsClientAPI contains the set of methods on the ListManagementImageListsClient type.
type ListManagementImageListsClientAPI interface {
	Create(ctx context.Context, contentType string, body contentmoderator.Body) (result contentmoderator.ImageList, err error)
	Delete(ctx context.Context, listID string) (result contentmoderator.String, err error)
	GetAllImageLists(ctx context.Context) (result contentmoderator.ListImageList, err error)
	GetDetails(ctx context.Context, listID string) (result contentmoderator.ImageList, err error)
	RefreshIndexMethod(ctx context.Context, listID string) (result contentmoderator.RefreshIndex, err error)
	Update(ctx context.Context, listID string, contentType string, body contentmoderator.Body) (result contentmoderator.ImageList, err error)
}

var _ ListManagementImageListsClientAPI = (*contentmoderator.ListManagementImageListsClient)(nil)

// ListManagementTermListsClientAPI contains the set of methods on the ListManagementTermListsClient type.
type ListManagementTermListsClientAPI interface {
	Create(ctx context.Context, contentType string, body contentmoderator.Body) (result contentmoderator.TermList, err error)
	Delete(ctx context.Context, listID string) (result contentmoderator.String, err error)
	GetAllTermLists(ctx context.Context) (result contentmoderator.ListTermList, err error)
	GetDetails(ctx context.Context, listID string) (result contentmoderator.TermList, err error)
	RefreshIndexMethod(ctx context.Context, listID string, language string) (result contentmoderator.RefreshIndex, err error)
	Update(ctx context.Context, listID string, contentType string, body contentmoderator.Body) (result contentmoderator.TermList, err error)
}

var _ ListManagementTermListsClientAPI = (*contentmoderator.ListManagementTermListsClient)(nil)

// ListManagementImageClientAPI contains the set of methods on the ListManagementImageClient type.
type ListManagementImageClientAPI interface {
	AddImage(ctx context.Context, listID string, tag *int32, label string) (result contentmoderator.Image, err error)
	AddImageFileInput(ctx context.Context, listID string, imageStream io.ReadCloser, tag *int32, label string) (result contentmoderator.Image, err error)
	AddImageURLInput(ctx context.Context, listID string, contentType string, imageURL contentmoderator.ImageURL, tag *int32, label string) (result contentmoderator.Image, err error)
	DeleteAllImages(ctx context.Context, listID string) (result contentmoderator.String, err error)
	DeleteImage(ctx context.Context, listID string, imageID string) (result contentmoderator.String, err error)
	GetAllImageIds(ctx context.Context, listID string) (result contentmoderator.ImageIds, err error)
}

var _ ListManagementImageClientAPI = (*contentmoderator.ListManagementImageClient)(nil)

// ListManagementTermClientAPI contains the set of methods on the ListManagementTermClient type.
type ListManagementTermClientAPI interface {
	AddTerm(ctx context.Context, listID string, term string, language string) (result autorest.Response, err error)
	DeleteAllTerms(ctx context.Context, listID string, language string) (result contentmoderator.String, err error)
	DeleteTerm(ctx context.Context, listID string, term string, language string) (result contentmoderator.String, err error)
	GetAllTerms(ctx context.Context, listID string, language string, offset *int32, limit *int32) (result contentmoderator.Terms, err error)
}

var _ ListManagementTermClientAPI = (*contentmoderator.ListManagementTermClient)(nil)

// ReviewsClientAPI contains the set of methods on the ReviewsClient type.
type ReviewsClientAPI interface {
	AddVideoFrame(ctx context.Context, teamName string, reviewID string, timescale *int32) (result autorest.Response, err error)
	AddVideoFrameStream(ctx context.Context, contentType string, teamName string, reviewID string, frameImageZip io.ReadCloser, frameMetadata string, timescale *int32) (result autorest.Response, err error)
	AddVideoFrameURL(ctx context.Context, contentType string, teamName string, reviewID string, videoFrameBody []contentmoderator.VideoFrameBodyItem, timescale *int32) (result autorest.Response, err error)
	AddVideoTranscript(ctx context.Context, teamName string, reviewID string, vttfile io.ReadCloser) (result autorest.Response, err error)
	AddVideoTranscriptModerationResult(ctx context.Context, contentType string, teamName string, reviewID string, transcriptModerationBody []contentmoderator.TranscriptModerationBodyItem) (result autorest.Response, err error)
	CreateJob(ctx context.Context, teamName string, contentType string, contentID string, workflowName string, jobContentType string, content contentmoderator.Content, callBackEndpoint string) (result contentmoderator.JobID, err error)
	CreateReviews(ctx context.Context, URLContentType string, teamName string, createReviewBody []contentmoderator.CreateReviewBodyItem, subTeam string) (result contentmoderator.ListString, err error)
	CreateVideoReviews(ctx context.Context, contentType string, teamName string, createVideoReviewsBody []contentmoderator.CreateVideoReviewsBodyItem, subTeam string) (result contentmoderator.ListString, err error)
	GetJobDetails(ctx context.Context, teamName string, jobID string) (result contentmoderator.Job, err error)
	GetReview(ctx context.Context, teamName string, reviewID string) (result contentmoderator.Review, err error)
	GetVideoFrames(ctx context.Context, teamName string, reviewID string, startSeed *int32, noOfRecords *int32, filter string) (result contentmoderator.Frames, err error)
	PublishVideoReview(ctx context.Context, teamName string, reviewID string) (result autorest.Response, err error)
}

var _ ReviewsClientAPI = (*contentmoderator.ReviewsClient)(nil)
