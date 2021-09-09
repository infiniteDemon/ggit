package models

type HeartbeatEvent struct {
	// Toggles between every JSON message as an "I am alive" indicator.
	// Required: Yes.
	Pulse bool `json:"pulse"`
	// Current active profile.
	// Required: No.
	CurrentProfile string `json:"current-profile"`
	// Current active scene.
	// Required: No.
	CurrentScene string `json:"current-scene"`
	// Current streaming state.
	// Required: No.
	Streaming bool `json:"streaming"`
	// Total time (in seconds) since the stream started.
	// Required: No.
	TotalStreamTime int `json:"total-stream-time"`
	// Total bytes sent since the stream started.
	// Required: No.
	TotalStreamBytes int `json:"total-stream-bytes"`
	// Total frames streamed since the stream started.
	// Required: No.
	TotalStreamFrames int `json:"total-stream-frames"`
	// Current recording state.
	// Required: No.
	Recording bool `json:"recording"`
	// Total time (in seconds) since recording started.
	// Required: No.
	TotalRecordTime int `json:"total-record-time"`
	// Total bytes recorded since the recording started.
	// Required: No.
	TotalRecordBytes int `json:"total-record-bytes"`
	// Total frames recorded since the recording started.
	// Required: No.
	TotalRecordFrames int `json:"total-record-frames"`
	// OBS Stats.
	// Required: Yes.
	Stats  *OBSStats `json:"stats"`
	_event `json:",squash"`
}

type _event struct {
	Type_           string `json:"update-type"`
	StreamTimecode_ string `json:"stream-timecode"`
	RecTimecode_    string `json:"rec-timecode"`
}

type OBSStats struct {
	FPS                float64 `json:"fps"`
	RenderTotalFrames  int     `json:"render-total-frames"`
	RenderMissedFrames int     `json:"render-missed-frames"`
	OutputTotalFrames  int     `json:"output-total-frames"`
	OutputMissedFrames int     `json:"output-missed-frames"`
	AverageFrameTime   float64 `json:"average-frame-time"`
	CPUUsage           float64 `json:"cpu-usage"`
	MemoryUsage        float64 `json:"memory-usage"`
	FreeDiskSpace      float64 `json:"free-disk-space`
}

type Scene struct {
	Name    string       `json:"name"`
	Sources []*SceneItem `json:"sources"`
}

type SceneItem struct {
	CY              int          `json:"cy"`
	CX              int          `json:"cx"`
	Name            string       `json:"name"`
	ID              int          `json:"id"`
	Render          bool         `json:"render"` // Visible or not
	Locked          bool         `json:"locked"`
	SourceCX        int          `json:"source_cx"`
	SourceCY        int          `json:"source_cy"`
	Type            string       `json:"type"` // One of: "input", "filter", "transition", "scene" or "unknown"
	Volume          int          `json:"volume"`
	X               int          `json:"x"`
	Y               int          `json:"y"`
	ParentGroupName string       `json:"parentGroupName,omitempty"` // Name of the item's parent (if this item belongs to a group)
	GroupChildren   []*SceneItem `json:"groupChildren"`             // List of children (if this item is a group)
}

type RevKey struct {
	KeyId   string `json:"keyId"`
	Shift   bool   `json:"shift"`
	Alt     bool   `json:"alt"`
	Control bool   `json:"control"`
	Command bool   `json:"command"`
}

type RevAction struct {
	Code int         `json:"code"`
	Role string      `json:"role"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var OBSKEYS = map[int]string{
	1:  "OBS_KEY_A",
	2:  "OBS_KEY_B",
	3:  "OBS_KEY_C",
	4:  "OBS_KEY_D",
	5:  "OBS_KEY_E",
	6:  "OBS_KEY_F",
	7:  "OBS_KEY_G",
	8:  "OBS_KEY_H",
	9:  "OBS_KEY_I",
	10: "OBS_KEY_J",
	11: "OBS_KEY_K",
	12: "OBS_KEY_L",
	13: "OBS_KEY_M",
	14: "OBS_KEY_N",
	15: "OBS_KEY_O",
}
