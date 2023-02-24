//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package appcamera

// Note: The code in this file was created from actual JSON payloads, using 1 or more of the
//       many JSON -> Go struct converters available.

type StreamUriRequest struct {
	StreamSetup  StreamSetup `json:"StreamSetup"`
	ProfileToken string      `json:"ProfileToken"`
}
type USBStartStreamingRequest struct {
	InputFps           string `json:"InputFps,omitempty"`
	InputImageSize     string `json:"InputImageSize,omitempty"`
	InputPixelFormat   string `json:"InputPixelFormat,omitempty"`
	OutputFrames       string `json:"OutputFrames,omitempty"`
	OutputFps          string `json:"OutputFps,omitempty"`
	OutputImageSize    string `json:"OutputImageSize,omitempty"`
	OutputAspect       string `json:"OutputAspect,omitempty"`
	OutputVideoCodec   string `json:"OutputVideoCodec,omitempty"`
	OutputVideoQuality string `json:"OutputVideoQuality,omitempty"`
}
type Transport struct {
	Protocol string `json:"Protocol"`
}
type StreamSetup struct {
	Stream    string    `json:"Stream"`
	Transport Transport `json:"Transport"`
}

type StreamingStatusResponse struct {
	Error              string `json:"Error"`
	InputFps           string `json:"InputFps"`
	InputImageSize     string `json:"InputImageSize"`
	IsStreaming        bool   `json:"IsStreaming"`
	OutputAspect       string `json:"OutputAspect"`
	OutputFps          string `json:"OutputFps"`
	OutputFrames       string `json:"OutputFrames"`
	OutputImageSize    string `json:"OutputImageSize"`
	OutputVideoQuality string `json:"OutputVideoQuality"`
}

type CameraType string

const (
	USB     CameraType = "USB"
	Onvif   CameraType = "Onvif"
	Unknown CameraType = "Unknown"
)

type CameraFeatures struct {
	PTZ        bool       `json:"PTZ"`
	Zoom       bool       `json:"Zoom"`
	CameraType CameraType `json:"CameraType"`
}

type PipelineRequest struct {
	Source      Source      `json:"source"`
	Destination Destination `json:"destination"`
}
type Source struct {
	URI  string `json:"uri"`
	Type string `json:"type"`
}
type Metadata struct {
	Type  string `json:"type"`
	Host  string `json:"host"`
	Topic string `json:"topic"`
}
type Frame struct {
	Type string `json:"type"`
	Path string `json:"path"`
}
type Destination struct {
	Metadata Metadata `json:"metadata"`
	Frame    Frame    `json:"frame"`
}

type OnvifPipelineConfig struct {
	ProfileToken string `json:"profile_token"`
}

type StartPipelineRequest struct {
	Onvif           *OnvifPipelineConfig      `json:"onvif,omitempty"`
	USB             *USBStartStreamingRequest `json:"usb,omitempty"`
	PipelineName    string                    `json:"pipeline_name"`
	PipelineVersion string                    `json:"pipeline_version"`
}

type PTZRange struct {
	XRange float64 `json:"XRange"`
	YRange float64 `json:"YRange"`
	ZRange float64 `json:"ZRange"`
}
