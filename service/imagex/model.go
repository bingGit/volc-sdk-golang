package imagex

import (
	"encoding/json"
	"fmt"

	"github.com/volcengine/volc-sdk-golang/base"
)

const (
	SEGMENT_CLASS_GENERAL    = "general"
	SEGMENT_CLASS_HUMAN      = "human"
	SEGMENT_CLASS_PRODUCT    = "product"
	SEGMENT_CLASS_HUMAN_V2   = "humanv2"
	SEGMENT_CLASS_PRODUCT_V2 = "productv2"

	SMARTCROP_POLICY_DEMOTION_CENTER = "center"
	SMARTCROP_POLICY_DEMOTION_FGLASS = "fglass"
	SMARTCROP_POLICY_DEMOTION_TOP    = "top"

	SMARTCROP_SCENE_NORMAL  = "normal"
	SMARTCROP_SCENE_CARTOON = "cartoon"

	FunctionEncryption = "Encryption"
)

// GetAllImageServices
type GetServicesResult struct {
	Services []Service `json:"Services"`
}

// GetImageService
type Service struct {
	ServiceName    string   `json:"ServiceName"`
	ServiceId      string   `json:"ServiceId"`
	ServiceRegion  string   `json:"ServiceRegion"`
	CreateAt       string   `json:"CreateAt"`
	ServiceStatus  string   `json:"ServiceStatus"`
	HasSigkey      bool     `json:"HasSigkey"`
	TemplatePrefix string   `json:"TemplatePrefix"`
	DomainInfos    []Domain `json:"DomainInfos"`
	PrimaryKey     string   `json:"PrimaryKey"`
	SecondaryKey   string   `json:"SecondaryKey"`
	ObjectAccess   bool     `json:"ObjectAccess"`
	CompactURL     bool     `json:"CompactURL"`
	Mirror         Mirror   `json:"Mirror"`
	Storage        Storage  `json:"Storage"`
	AllowBkts      []string `json:"AllowBkts"`
}

type Mirror struct {
	Schema  string            `json:"Schema"`
	Host    string            `json:"Host"`
	Source  string            `json:"Source"`
	Headers map[string]string `json:"Headers"`
}

type Domain struct {
	DomainName string `json:"DomainName"`
	CNAME      string `json:"CNAME"`
	Status     string `json:"Status"`
	IsDefault  bool   `json:"IsDefault"`
}

type Storage struct {
	BktName  string `json:"BktName"`
	TTL      int    `json:"TTL"`
	AllTypes bool   `json:"AllTypes"`
}

// GetServiceDomains
type DomainResult struct {
	Domain      string      `json:"domain"`
	CNAME       string      `json:"cname"`
	Status      string      `json:"status"`
	HttpsConfig HttpsConfig `json:"https_config"`
	IsDefault   bool        `json:"is_default"`
	Resolved    bool        `json:"resolved"`
}

type HttpsConfig struct {
	EnableHttps bool   `json:"enable_https"`
	ForceHttps  bool   `json:"force_https"`
	CertId      string `json:"cert_id"`
}

type AccessControlConfig struct {
	ReferLink ReferLink `json:"refer_link"`
}

type ReferLink struct {
	Enabled         bool     `json:"enabled"`
	IsWhiteMode     bool     `json:"is_white_mode"`
	Values          []string `json:"values"`
	AllowEmptyRefer bool     `json:"allow_empty_refer"`
}

type RespHdr struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

// DeleteImageUploadFiles
type DeleteImageParam struct {
	StoreUris []string `json:"StoreUris"`
}

type DeleteImageResult struct {
	ServiceId    string   `json:"ServiceId"`
	DeletedFiles []string `json:"DeletedFiles"`
}

// ApplyImageUpload
type ApplyUploadImageParam struct {
	ServiceId     string
	SessionKey    string
	UploadNum     int
	StoreKeys     []string
	ContentTypes  []string
	Prefix        string
	FileExtension string
	CommitParam   *CommitUploadImageParam
	SkipMeta      bool
	SkipCommit    bool
	Overwrite     bool
}

type ApplyUploadImageResult struct {
	UploadAddress UploadAddress `json:"UploadAddress"`
	RequestId     string        `json:"RequestId"`
}

type UploadAddress struct {
	SessionKey  string      `json:"SessionKey"`
	UploadHosts []string    `json:"UploadHosts"`
	StoreInfos  []StoreInfo `json:"StoreInfos"`
}

type StoreInfo struct {
	StoreUri string `json:"StoreUri"`
	Auth     string `json:"Auth"`
}

// CommitImageUpload
type CommitUploadImageParam struct {
	ServiceId   string     `json:"-"`
	SessionKey  string     `json:"SessionKey"`
	SuccessOids []string   `json:"SuccessOids"`
	Functions   []Function `json:"Functions"`
	SkipMeta    bool       `json:"-"`
}

type Function struct {
	Name  string      `json:"Name"`
	Input interface{} `json:"Input"`
}

type EncryptionInput struct {
	Config       map[string]string `json:"Config"`
	PolicyParams map[string]string `json:"PolicyParams"`
}

type CommitUploadImageResult struct {
	Results    []Result    `json:"Results"`
	RequestId  string      `json:"RequestId"`
	ImageInfos []ImageInfo `json:"PluginResult"`
}

type Result struct {
	Uri        string     `json:"Uri"`
	UriStatus  int        `json:"UriStatus"` // 图片上传结果（2000:成功，2001:失败）需要传 SuccessOids 才会返回
	Encryption Encryption `json:"Encryption"`
	PutError   *PutError  `json:"-"` // 上传阶段错误
}

type PutError struct {
	ErrorCode int
	Error     string
	Message   string
}

type Encryption struct {
	Uri       string            `json:"Uri"`
	SecretKey string            `json:"SecretKey"`
	Algorithm string            `json:"Algorithm"`
	Version   string            `json:"Version"`
	SourceMd5 string            `json:"SourceMd5"`
	Extra     map[string]string `json:"Extra"`
}

type ImageInfo struct {
	FileName    string `json:"FileName"`
	ImageUri    string `json:"ImageUri"`
	ImageWidth  int    `json:"ImageWidth"`
	ImageHeight int    `json:"ImageHeight"`
	ImageMd5    string `json:"ImageMd5"`
	ImageFormat string `json:"ImageFormat"`
	ImageSize   int    `json:"ImageSize"`
	FrameCnt    int    `json:"FrameCnt"`
	Duration    int    `json:"Duration"`
}

type UploadPolicy struct {
	ContentTypeBlackList []string `json:"ContentTypeBlackList,omitempty"`
	ContentTypeWhiteList []string `json:"ContentTypeWhiteList,omitempty"`
	FileSizeUpLimit      string   `json:"FileSizeUpLimit,omitempty"`
	FileSizeBottomLimit  string   `json:"FileSizeBottomLimit,omitempty"`
}

// CreateImageContentTask
type CreateImageContentTaskReq struct {
	ServiceId string   `query:"ServiceId" json:"-"`
	TaskType  string   `json:"TaskType"`
	Urls      []string `json:"Urls"`
}

type CreateImageContentTaskResp struct {
	Msg    string `json:"Msg"`
	TaskId string `json:"TaskId"`
}

// GetImageContentTaskDetail
type GetImageContentTaskDetailReq struct {
	ServiceId string `query:"ServiceId" json:"-"`
	TaskType  string `json:"TaskType"`
	TaskId    string `json:"TaskId"`
	State     string `json:"State"`
	Order     string `json:"Order"`
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Url       string `json:"Url"`
	PageNum   int    `json:"PageNum"`
	PageSize  int    `json:"PageSize"`
}

type GetImageContentTaskDetailResp struct {
	PageNum  int                                 `json:"PageNum"`
	PageSize int                                 `json:"PageSize"`
	Total    int                                 `json:"Total"`
	Data     []GetImageContentTaskDetailRespData `json:"Data"`
}

type GetImageContentTaskDetailRespData struct {
	Url        string `json:"Url"`
	State      string `json:"State"`
	Process    string `json:"Process"`
	TaskType   string `json:"TaskType"`
	CreateTime int64  `json:"CreateTime"`
	UpdateTime int64  `json:"UpdateTime"`
	Msg        string `json:"Msg"`
	TaskId     string `json:"TaskId"`
}

// GetImageContentBlockList
type GetImageContentBlockListReq struct {
	ServiceId string `query:"ServiceId" json:"-"`
	State     string `json:"State"`
	Order     string `json:"Order"`
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	Url       string `json:"Url"`
	PageNum   int    `json:"PageNum"`
	PageSize  int    `json:"PageSize"`
}

type GetImageContentBlockListResp struct {
	PageNum  int                                `json:"PageNum"`
	PageSize int                                `json:"PageSize"`
	Total    int                                `json:"Total"`
	Data     []GetImageContentBlockListRespData `json:"Data"`
}

type GetImageContentBlockListRespData struct {
	Url        string `json:"Url"`
	State      string `json:"State"`
	Process    string `json:"Process"`
	TaskType   string `json:"TaskType"`
	CreateTime int64  `json:"CreateTime"`
	UpdateTime int64  `json:"UpdateTime"`
	Msg        string `json:"Msg"`
}

type CreateImageCompressTaskReq struct {
	ServiceId string
	FileList  []UncompressFileInfo
	IndexFile string
	Output    string
	Callback  string
	ZipMode   int
}

type UncompressFileInfo struct {
	Url    string
	Alias  string
	Folder string
}

type CreateImageCompressTaskResp struct {
	TaskId string
}

type GetImageCompressTaskInfoReq struct {
	ServiceId string
	TaskId    string
}

type GetImageCompressTaskInfoResp struct {
	Status       string
	TaskId       string
	ProcessCount int
	OutputUri    string
	ErrMsg       string
	ErrCode      int
}

// FetchImageUrl
type FetchUrlReq struct {
	Url              string              `json:"Url"`
	ServiceId        string              `json:"ServiceId"`
	StoreKey         string              `json:"StoreKey"`
	RequestHeader    map[string][]string `json:"RequestHeader"`
	TimeOut          int                 `json:"TimeOut"`
	Async            bool                `json:"Async"`
	Callback         string              `json:"Callback"`
	Host             string              `json:"Host"`
	MD5              string              `json:"MD5"`
	CallbackBodyType string              `json:"CallbackBodyType"`
	CallbackBody     string              `json:"CallbackBody"`
	CallbackHost     string              `json:"CallbackHost"`
	IgnoreSameKey    bool                `json:"IgnoreSameKey"`
}

type FetchUrlResp struct {
	Url         string `json:"Url"`
	StoreUri    string `json:"StoreUri"`
	FSize       int    `json:"FSize"`
	ImageWidth  int    `json:"ImageWidth,omitempty"`
	ImageHeight int    `json:"ImageHeight,omitempty"`
	Duration    int    `json:"Duration,omitempty"`
	FrameCnt    int    `json:"FrameCnt,omitempty"`
	ImageFormat string `json:"ImageFormat,omitempty"`
	TaskId      string `json:"TaskId"`
}

// GetUrlFetchTask
type GetUrlFetchTaskReq struct {
	Id        string `query:"Id"`
	ServiceId string `query:"ServiceId"`
}

type GetUrlFetchTaskResp struct {
	Id               string `json:"Id"`
	Status           string `json:"Status"`
	Url              string `json:"Url"`
	StoreUri         string `json:"StoreUri"`
	Callback         string `json:"Callback"`
	CallbackBody     string `json:"CallbackBody"`
	CallbackBodyType string `json:"CallbackBodyType"`
	Err              string `json:"Err"`
	Code             string `json:"Code"`
}

// UpdateImageStorageTTL
type UpdateImageStorageTTLReq struct {
	ServiceId string `json:"ServiceId"`
	TTL       int    `json:"TTL"`
}

type UpdateImageStorageTTLResp struct {
}

// GetImageStyleResult
type GetImageStyleResultReq struct {
	ServiceId     string
	StyleId       string            `json:"StyleId"`
	Params        map[string]string `json:"Params"`
	OutputFormat  string            `json:"OutputFormat"`
	OutputQuality int               `json:"OutputQuality"`
}

type GetImageStyleResultResp struct {
	ResUri       string `json:"ResUri"`
	RenderDetail []struct {
		Element string `json:"Element"`
		ErrMsg  string `json:"ErrMsg"`
	} `json:"RenderDetail"`
}

// GetImageOCR
type GetImageOCRResult struct {
	Scene         string                                  `json:"Scene"`
	GeneralResult *[]*GetImageOCRGeneralTextInfo          `json:"GeneralResult"`
	LicenseResult *map[string]*GetImageOCRLicenseTextInfo `json:"LicenseResult"`
}

type GetImageOCRGeneralTextInfo struct {
	Content    string  `json:"Content"`
	Confidence string  `json:"Confidence"`
	Location   [][]int `json:"Location"`
}

type GetImageOCRLicenseTextInfo struct {
	Content  string `json:"Content"`
	Location []int  `json:"Location"`
}

type GetImageOCRParam struct {
	ServiceId string `json:"-"`
	StoreUri  string `json:"StoreUri"`
	ImageUrl  string `json:"ImageUrl"`
	Scene     string `json:"Scene"`
}

// GetImageBgFillResult
type GetImageBgFillParam struct {
	ServiceId string  `json:"ServiceId"`
	StoreUri  string  `json:"StoreUri"`
	Model     int     `json:"Model"`
	Top       float64 `json:"Top"`
	Bottom    float64 `json:"Bottom"`
	Left      float64 `json:"Left"`
	Right     float64 `json:"Right"`
}

type GetImageBgFillResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageEnhanceResult
type GetImageEnhanceParam struct {
	ServiceId    string   `json:"ServiceId"`
	StoreUri     string   `json:"StoreUri"`
	Model        int      `json:"Model"`
	DisableAr    bool     `json:"DisableAr"`
	DisableSharp bool     `json:"DisableSharp"`
	Resolution   string   `json:"Resolution"`
	Actions      []string `json:"Actions"`
	OutFormat    string   `json:"OutFormat"`
	Quality      int      `json:"Quality"`
}

type GetImageEnhanceResult struct {
	ResUri string `json:"ResUri"`
	Method string `json:"Method"`
}

// GetImageEraseResult
type GetImageEraseParam struct {
	ServiceId string     `json:"ServiceId"`
	StoreUri  string     `json:"StoreUri"`
	Model     string     `json:"Model"`
	BBox      []EraseBox `json:"BBox"`
}

type EraseBox struct {
	X1 float64 `json:"X1"`
	Y1 float64 `json:"Y1"`
	X2 float64 `json:"X2"`
	Y2 float64 `json:"Y2"`
}

type GetImageEraseResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageQuality
type GetImageQualityParam struct {
	ServiceId string
	ImageUrl  string `json:"ImageUrl"`
	VqType    string `json:"VqType"`
}

type GetImageQualityResult struct {
	VqType   string                 `json:"VqType"`
	NrScores map[string]interface{} `json:"NrScores"`
}

// GetImageSegment
type GetImageSegmentParam struct {
	ServiceId string
	StoreUri  string   `json:"StoreUri"`
	Class     string   `json:"Class"`
	Refine    bool     `json:"Refine"`
	OutFormat string   `json:"OutFormat"`
	TransBg   bool     `json:"TransBg"`
	Contour   *Contour `json:"Contour"`
}

type Contour struct {
	Color string `json:"Color"`
	Size  int    `json:"Size"`
}

type GetImageSegmentResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageSuperResolutionResult
type GetImageSuperResolutionParam struct {
	ServiceId string  `json:"ServiceId"`
	StoreUri  string  `json:"StoreUri"`
	Multiple  float64 `json:"Multiple"`
}

type GetImageSuperResolutionResp struct {
	ResUri string `json:"ResUri"`
}

// GetImageUploadFile
type GetImageUploadFileParam struct {
	ServiceId string `json:"ServiceId"`
	StoreUri  string `json:"StoreUri"`
}

type GetImageUploadFileResult struct {
	ServiceId    string `json:"ServiceId"`
	StoreUri     string `json:"StoreUri"`
	LastModified string `json:"LastModified"`
	FileSize     int    `json:"FileSize"`
	Marker       int64  `json:"Marker"`
}

// GetImageUploadFiles
type GetImageUploadFilesParam struct {
	ServiceId string `json:"ServiceId"`
	Limit     int    `json:"Limit"`
	Marker    int64  `json:"Marker"`
}

type GetImageUploadFilesResult struct {
	ServiceId   string       `json:"ServiceId"`
	FileObjects []FileObject `json:"FileObjects"`
	Status      string       `json:"Status"`
	HasMore     bool         `json:"hasMore"`
}

type FileObject struct {
	StoreUri     string `json:"StoreUri"`
	Key          string `json:"Key"`
	LastModified string `json:"LastModified"`
	FileSize     int    `json:"FileSize"`
	Marker       int64  `json:"Marker"`
}

// GetImageDuplicateDetection
type GetImageDuplicateDetectionParam struct {
	Urls      []string `json:"Urls"`
	ServiceId string   `json:"-"`
}

type GetImageDuplicateDetectionAsyncParam struct {
	GetImageDuplicateDetectionParam
	Callback string `json:"Callback"`
}

type getImageDuplicateDetectionParam struct {
	Urls     []string `json:"Urls"`
	Async    bool     `json:"Async"`
	Callback string   `json:"Callback"`
}

type GetImageDuplicateDetectionResult struct {
	Score  string              `json:"Score"`
	Groups map[string][]string `json:"Groups"`
}

type GetImageDuplicateDetectionAsyncResult struct {
	TaskId   string `json:"TaskId"`
	Callback string `json:"Callback"`
}

type GetImageDuplicateDetectionCallbackBody struct {
	Status int                              `json:"Status"`
	TaskId string                           `json:"TaskId"`
	Result GetImageDuplicateDetectionResult `json:"Result"`
}

type GetDedupTaskStatusResult struct {
	Status int                              `json:"Status"`
	TaskId string                           `json:"TaskId"`
	Result GetImageDuplicateDetectionResult `json:"Result"`
}

// GetDenoisingImage
type GetDenoisingImageParam struct {
	ServiceId   string  `json:"-"`
	StoreUri    string  `json:"StoreUri"`
	ImageUrl    string  `json:"ImageUrl"`
	OutFormat   string  `json:"OutFormat"`
	Intensity   float64 `json:"Intensity"`
	CanDemotion bool    `json:"CanDemotion"`
}

type GetDenoisingImageResult struct {
	ResUri   string `json:"ResUri"`
	Demotion bool   `json:"Demotion"`
}

// GetImageComicResult
type GetImageComicParam struct {
	ServiceId string `json:"ServiceId"`
	StoreUri  string `json:"StoreUri"`
}

type GetImageComicResult struct {
	ResUri string `json:"ResUri"`
}

// GetImageSmartCropResult
type GetImageSmartCropParam struct {
	ServiceId string  `json:"ServiceId"`
	StoreUri  string  `json:"StoreUri"`
	Policy    string  `json:"Policy"`
	Scene     string  `json:"Scene"`
	Sigma     float64 `json:"Sigma"`
	Width     int     `json:"Width"`
	Height    int     `json:"Height"`
}

type GetImageSmartCropResult struct {
	ResUri   string `json:"ResUri"`
	Demotion bool   `json:"Demotion"`
}

// GetLicensePlateDetection
type GetLicensePlateDetectionParam struct {
	ServiceId string `json:"-"`
	ImageUri  string `json:"ImageUri"`
}

type LocationsLower struct {
	X1 int `json:"x1"`
	X2 int `json:"x2"`
	Y1 int `json:"y1"`
	Y2 int `json:"y2"`
}

type GetLicensePlateDetectionResult struct {
	Locations []LocationsLower `json:"Locations"`
}

// GetImagePSDetection
type GetImagePSDetectionParam struct {
	ServiceId string `json:"-"`
	ImageUri  string `json:"ImageUri"`
	Method    string `json:"Method"`
}

type GetImagePSDetectionResult struct {
	ElaImage []byte  `json:"ElaImage"`
	NaImage  []byte  `json:"NaImage"`
	HeImage  []byte  `json:"HeImage"`
	Score    float64 `json:"Score"`
	Label    int     `json:"Label"`
}

// GetPrivateImageType
type GetPrivateImageTypeParam struct {
	ServiceId  string  `json:"-"`
	ImageUri   string  `json:"ImageUri"`
	Method     string  `json:"Method"`
	ThresFace  float64 `json:"ThresFace"`
	ThresCloth float64 `json:"ThresCloth"`
}

type GetPrivateImageTypeResult struct {
	Face  int `json:"Face"`
	Cloth int `json:"Cloth"`
}

type CreateHiddenWatermarkImageParam struct {
	ServiceId string
	Algorithm string `json:"Algorithm"`
	Strength  string `json:"Strength"`
	Info      string `json:"Info"`
}

type CreateHiddenWatermarkImageResult struct {
	StoreUri string `json:"StoreUri"`
}

// CreateImageHmEmbed
type CreateImageHmEmbedParam struct {
	ServiceId     string `json:"ServiceId"`
	StoreUri      string `json:"StoreUri"`
	Algorithm     string `json:"Algorithm"`
	Info          string `json:"Info"`
	OutFormat     string `json:"OutFormat"`
	OutQuality    int    `json:"OutQuality"`
	StrengthLevel string `json:"StrengthLevel"`
	ImageUrl      string `json:"ImageUrl"`
	Strength      int    `json:"Strength"`
}

type CreateImageHmEmbedResult struct {
	StoreUri string `json:"StoreUri"`
}

// CreateImageHmExtract
type CreateImageHmExtractParam struct {
	ServiceId string `json:"ServiceId"`
	StoreUri  string `json:"StoreUri"`
	Algorithm string `json:"Algorithm"`
	ImageUrl  string
}

type CreateImageHmExtractResult struct {
	Info       string `json:"Info"`
	StatusCode int    `json:"StatusCode"`
}

func UnmarshalResultInto(data []byte, result interface{}) error {
	resp := new(base.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	data, err := json.Marshal(resp.Result)
	if err != nil {
		return fmt.Errorf("fail to marshal result, %v", err)
	}
	if err = json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

// DescribeImageVolcCdnAccessLog
type DescribeImageVolcCdnAccessLogReq struct {
	ServiceId string `query:"ServiceId"`
	Domain    string `json:"Domain"`
	Region    string `json:"Region"`
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	PageNum   int    `json:"PageNum"`
	PageSize  int    `json:"PageSize"`
}

type DescribeImageVolcCdnAccessLogResp struct {
	Domain     string                                 `json:"Domain"`
	PageSize   int                                    `json:"PageSize"`
	PageNum    int                                    `json:"PageNum"`
	TotalCount int                                    `json:"TotalCount"`
	Logs       []DescribeImageVolcCdnAccessLogRespLog `json:"Logs"`
}

type DescribeImageVolcCdnAccessLogRespLog struct {
	StartTime int64  `json:"StartTime"`
	EndTime   int64  `json:"EndTime"`
	LogName   string `json:"LogName"`
	LogPath   string `json:"LogPath"`
	LogSize   int64  `json:"LogSize"`
}

// DeleteImageAuditResult
type DeleteImageAuditResultReq struct {
	TaskId  string `json:"TaskId"`
	EntryId string `json:"EntryId"`
}

type DeleteImageAuditResultResp struct {
}

// CreateImageAuditTask
type CreateImageAuditTaskReq struct {
	Type                   string            `json:"Type"`
	ServiceID              string            `json:"ServiceId"`
	Region                 string            `json:"Region"`
	TaskType               string            `json:"TaskType"`
	AuditAbility           int               `json:"AuditAbility"`
	AuditDimensions        []string          `json:"AuditDimensions"`
	AuditTextDimensions    []string          `json:"AuditTextDimensions"`
	FreezeStrategy         int               `json:"FreezeStrategy"`
	ResURI                 []string          `json:"ResUri"`
	CallbackDimensions     []string          `json:"CallbackDimensions"`
	FreezeType             []string          `json:"FreezeType"`
	CallbackImageTypes     []string          `json:"CallbackImageTypes"`
	CallbackURL            string            `json:"CallbackUrl"`
	EnableFreeze           bool              `json:"EnableFreeze"`
	EnableCallback         bool              `json:"EnableCallback"`
	AuditPrefix            []string          `json:"AuditPrefix"`
	FreezeDimensions       []string          `json:"FreezeDimensions"`
	EnableAuditRange       int               `json:"EnableAuditRange"`
	NoAuditPrefix          []string          `json:"NoAuditPrefix"`
	ImageInfos             []*AuditImageInfo `json:"ImageInfos"`
	EnableLargeImageDetect bool              `json:"EnableLargeImageDetect"`
}

type AuditImageInfo struct {
	ImageUri string `json:"ImageUri"`
	DataId   string `json:"DataId"`
}

type CreateImageAuditTaskResp struct {
	TaskID string `json:"TaskId"`
}

// UpdateAuditImageStatus
type UpdateAuditImageStatusReq struct {
	TaskID  string `json:"TaskId"`
	EntryID string `json:"EntryId"`
	Status  string `json:"Status"`
}

type UpdateAuditImageStatusResp struct {
}

type UpdateImageAuditTaskStatusReq struct {
	TaskID string `json:"TaskId"`
	Status string `json:"Status"`
}

type UpdateImageAuditTaskStatusResp struct {
}

// GetAuditEntrysCount
type GetAuditEntrysCountReq struct {
	TaskId string `query:"TaskId"`
}

type GetAuditEntrysCountResp struct {
	Total               int `json:"Total"`
	NopassCount         int `json:"NopassCount"`
	RecheckCount        int `json:"RecheckCount"`
	SuccessCount        int `json:"SuccessCount"`
	FailedCount         int `json:"FailedCount"`
	ExceptionCount      int `json:"ExceptionCount"`
	AuditFailCount      int `json:"AuditFailCount"`
	AuditSuccessCount   int `json:"AuditSuccessCount"`
	AuditTotal          int `json:"AuditTotal"`
	AuditExceptionCount int `json:"AuditExceptionCount"`
	AuditNopassCount    int `json:"AuditNopassCount"`
	AuditRecheckCount   int `json:"AuditRecheckCount"`
}

// GetImageAuditResult
type GetImageAuditResultReq struct {
	TaskId          string `query:"TaskId"`
	AuditSuggestion string `query:"AuditSuggestion"`
	Problem         string `query:"Problem"`
	ImageType       string `query:"ImageType"`
	Marker          string `query:"Marker"`
	Limit           int    `query:"Limit"`
	Type            string `query:"Type"`
}

type GetImageAuditResultResp struct {
	Results  []GetImageAuditResult `json:"Results"`
	HaveMore bool                  `json:"HaveMore"`
}

type GetImageAuditResult struct {
	Marker          string   `json:"Marker"`
	EndAt           string   `json:"EndAt"`
	Type            string   `json:"Type"`
	ErrMsg          string   `json:"ErrMsg"`
	TaskID          string   `json:"TaskId"`
	Problems        []string `json:"Problems"`
	ImageType       string   `json:"ImageType"`
	EntryID         string   `json:"EntryId"`
	AuditSuggestion string   `json:"AuditSuggestion"`
	ImageURI        string   `json:"ImageUri"`
	Ability         int      `json:"Ability"`
}

// CreateImageRetryAuditTask
type CreateImageRetryAuditTaskReq struct {
	TaskID string `json:"TaskId"`
	Entry  string `json:"Entry"`
}

type CreateImageRetryAuditTaskResp struct {
	TaskId string `json:"TaskId"`
}

// UpdateImageAuditTask
type UpdateImageAuditTaskReq struct {
	ServiceID              string   `json:"ServiceId"`
	AuditDimensions        []string `json:"AuditDimensions"`
	AuditTextDimensions    []string `json:"AuditTextDimensions"`
	TaskID                 string   `json:"TaskId"`
	EnableFreeze           bool     `json:"EnableFreeze"`
	CallbackURL            string   `json:"CallbackUrl"`
	AuditPrefix            []string `json:"AuditPrefix"`
	NoAuditPrefix          []string `json:"NoAuditPrefix"`
	ResURI                 []string `json:"ResUri"`
	FreezeStrategy         int      `json:"FreezeStrategy"`
	EnableAuditRange       string   `json:"EnableAuditRange"`
	CallbackDimensions     []string `json:"CallbackDimensions"`
	FreezeDimensions       []string `json:"FreezeDimensions"`
	Region                 string   `json:"Region"`
	CallbackImageTypes     []string `json:"CallbackImageTypes"`
	EnableCallback         bool     `json:"EnableCallback"`
	FreezeType             []string `json:"FreezeType"`
	EnableLargeImageDetect bool     `json:"EnableLargeImageDetect"`
}

type UpdateImageAuditTaskResp struct {
}

// GetImageAuditTasks
type GetImageAuditTasksReq struct {
	Region       string `query:"Region"`
	TaskType     string `query:"TaskType"`
	AuditAbility string `query:"AuditAbility"`
	Status       string `query:"Status"`
	Offset       int    `query:"Offset"`
	Limit        int    `query:"Limit"`
	Type         string `query:"Type"`
}

type GetImageAuditTasksResp struct {
	Tasks []GetImageAuditTask `json:"Tasks"`
	Total int                 `json:"Total"`
}
type GetImageAuditTaskInfo struct {
	AuditAbility           int               `json:"AuditAbility"`
	FreezeType             []string          `json:"FreezeType"`
	FreezeDimensions       []string          `json:"FreezeDimensions"`
	EnableFreeze           bool              `json:"EnableFreeze"`
	AuditDimensions        []string          `json:"AuditDimensions"`
	CallbackDimensions     []string          `json:"CallbackDimensions"`
	EnableCallback         bool              `json:"EnableCallback"`
	FreezeStrategy         int               `json:"FreezeStrategy"`
	CallbackURL            string            `json:"CallbackUrl"`
	CallbackImageTypes     []string          `json:"CallbackImageTypes"`
	EnableAuditRange       int               `json:"EnableAuditRange"`
	NoAuditPrefix          []string          `json:"NoAuditPrefix"`
	AuditPrefix            []string          `json:"AuditPrefix"`
	ProcessedNumber        int               `json:"ProcessedNumber"`
	FailedCount            int               `json:"FailedCount"`
	SuccessCount           int               `json:"SuccessCount"`
	ImageInfos             []*AuditImageInfo `json:"ImageInfos"`
	AuditTextDimensions    []string          `json:"AuditTextDimensions"`
	EnableLargeImageDetect bool              `json:"EnableLargeImageDetect"`
}
type GetImageAuditTask struct {
	Status    string                 `json:"Status"`
	ID        string                 `json:"Id"`
	TaskType  string                 `json:"TaskType"`
	AuditTask *GetImageAuditTaskInfo `json:"AuditTask"`
	Type      string                 `json:"Type"`
	ResURI    []string               `json:"ResUri"`
	CreateAt  string                 `json:"CreateAt"`
	UpdateAt  string                 `json:"UpdateAt"`
	ServiceID string                 `json:"ServiceId"`
}

type GetImageAiGenerateTaskReq struct {
	TaskId string
}

type GetImageAiGenerateTaskResp struct {
	Uris            string    `json:"Uris"`
	Status          string    `json:"Status"`
	AestheticScores []float64 `json:"AestheticScores"`
	SDScores        []float64 `json:"SDScores"`
}

type GetAiGenerateImageReq struct {
	ServiceId      string  `json:"-"`
	URL            string  `json:"Url"`
	NeedSeg        bool    `json:"NeedSeg"`
	PositivePrompt string  `json:"PositivePrompt"`
	NegativePrompt string  `json:"NegativePrompt"`
	ProductRatio   float64 `json:"ProductRatio"`
	OutputSize     int     `json:"OutputSize"`
	Callback       string  `json:"Callback"`
	Scene          string  `json:"Scene"`
	UseCaption     bool    `json:"UseCaption"`
	ReturnTop1     bool    `json:"ReturnTop1"`
	CX             int     `json:"CX"`
	CY             int     `json:"CY"`
	SafeH          int     `json:"SafeH"`
	SafeW          int     `json:"SafeW"`
	Extra          string  `json:"Extra"`
}

type GetAiGenerateImageResp struct {
	TaskId string `json:"TaskId"`
}
