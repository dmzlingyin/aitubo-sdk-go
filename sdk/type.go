package sdk

import "time"

/********************************************************* Job ********************************************************/

type BaseRequest struct {
	Prompt          string  `json:"prompt,omitempty"`
	NegativePrompt  string  `json:"negativePrompt,omitempty"`
	Width           int32   `json:"width,omitempty"`
	Height          int32   `json:"height,omitempty"`
	Count           int32   `json:"count,omitempty"`
	Steps           int32   `json:"steps,omitempty"`
	GuidScale       float32 `json:"guidanceScale,omitempty"`
	Art             bool    `json:"art,omitempty"`
	ModelID         string  `json:"modelId,omitempty"`
	Loras           []*Lora `json:"loras,omitempty"`
	NsfwThreshold   float32 `json:"nsfwThreshold,omitempty"`
	SafetyChecker   bool    `json:"safetyChecker,omitempty"`
	Seed            int64   `json:"seed,omitempty"`
	Scheduler       string  `json:"scheduler,omitempty"`
	PromptOptimizer bool    `json:"promptOptimizer,omitempty"`
	HDEnhance       bool    `json:"hdEnhance,omitempty"`
	ImageFormat     string  `json:"imageFormat,omitempty"` // default: jpg
}

type Lora struct {
	ModelID   string   `json:"modelId,omitempty"`
	Alpha     float32  `json:"alpha,omitempty"`
	Triggers  []string `json:"triggers,omitempty"`
	ModelName string   `json:"modelName,omitempty"`
}

type CreateText2ImageJobRequest struct {
	BaseRequest
}

type CreateImage2ImageJobRequest struct {
	Image    string  `json:"imagePath"`
	Strength float32 `json:"strength,omitempty"` // range: (0-1]
	BaseRequest
}

type CreateControlNetJobRequest struct {
	Image           string  `json:"imagePath"`
	ControlModel    string  `json:"controlModel,omitempty"`
	ControlStrength float32 `json:"controlStrength,omitempty"` // range: (0-2]
	ControlFilter   string  `json:"controlFilter,omitempty"`
	BaseRequest
}

type CreateUpscaleJobRequest struct {
	Image         string `json:"imagePath"`
	Model         string `json:"model,omitempty"`
	UpscaleFactor int32  `json:"upscaleFactor,omitempty"`
	Beauty        bool   `json:"beauty"`
	ArtID         string `json:"artId,omitempty"`
	Type          string `json:"type"`
}

type CreateJobResponse struct {
	Code int32  `json:"code"`
	Info string `json:"info"`
	Data *struct {
		ID             string `json:"id"`
		Credit         int32  `json:"credit"`
		SubCredit      int32  `json:"subCredit"`
		ConsumedCredit int32  `json:"consumedCredit"`
	} `json:"data"`
}

type QueryJobResponse struct {
	Code int32  `json:"code"`
	Info string `json:"info,omitempty"`
	Data *struct {
		Status int32 `json:"status"`
		Result struct {
			Code int32  `json:"code,omitempty"`
			Info string `json:"info,omitempty"`
			Data *struct {
				Domain string   `json:"domain"`
				Images []string `json:"images"`
				Nsfws  []bool   `json:"nsfws"`
				ArtIDs []string `json:"artIds"`
			}
		} `json:"result"`
	} `json:"data"`
}

/********************************************************* Model ******************************************************/

const (
	ModelTypePlatform  = "platform"
	ModelTypeCommunity = "community"
	ModelTypeUsed      = "used"

	ModelKindCheckpoint = "checkpoint"
	ModelKindLora       = "lora"

	BaseModelSD15   = "sd_1.5"
	BaseModelSDXL10 = "sdxl_1.0"
)

type ListModelRequest struct {
	PageIndex int
	PageSize  int
	Type      string
	Kind      string
	Name      string
	BaseModel string
}

type AituboModel struct {
	ID            string      `json:"id,omitempty"`
	Name          string      `json:"name,omitempty"`
	UserName      string      `json:"username,omitempty"`
	Desc          string      `json:"desc,omitempty"`
	Cover         string      `json:"cover,omitempty"`
	CategoryNames []string    `json:"categoryNames,omitempty"`
	StyleNames    []string    `json:"styleNames,omitempty"`
	Arts          []*modelArt `json:"arts,omitempty"` // template images
	IsFavourite   bool        `json:"isFavourite,omitempty"`
	Kind          string      `json:"kind,omitempty"`
	BaseModel     string      `json:"baseModel,omitempty"`
	Options       options     `json:"options,omitempty"`
	Order         int32       `json:"order,omitempty"`
	Nsfw          bool        `json:"nsfw,omitempty"`
	Free          bool        `json:"free,omitempty"`
	CreatedAt     time.Time   `json:"createdAt,omitempty"`
}

type options struct {
	Triggers     []string `json:"triggers,omitempty"`
	DefaultModel string   `json:"defaultModel,omitempty"`
	Alpha        float32  `json:"alpha,omitempty"`
	Width        int32    `json:"width,omitempty"`
	Height       int32    `json:"height,omitempty"`
}

type modelArt struct {
	Path  string `json:"path"`
	ArtID string `json:"artId"`
}

type ListModelResponse struct {
	Code int32  `json:"code"`
	Info string `json:"info,omitempty"`
	Data *struct {
		Domain string        `json:"domain"`
		Total  int32         `json:"total"`
		Models []AituboModel `json:"models"`
	} `json:"data,omitempty"`
}
