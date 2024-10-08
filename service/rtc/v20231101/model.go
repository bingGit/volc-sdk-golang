package rtc_v20231101

type BanRoomUserBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// 封禁时长，取值范围为 {0,[10,604800]}，单位为秒。 若传入值为空或 0 表示允许用户重新进房。 若传入值小于 10，自动修改为 10。 若传入值大于 604800，自动修改为 604800。
	ForbiddenInterval *int32 `json:"ForbiddenInterval,omitempty"`

	// 指定房间 ID
	RoomID *string `json:"RoomId,omitempty"`

	// 希望封禁用户的 ID
	UserID *string `json:"UserId,omitempty"`
}

type BanRoomUserRes struct {

	// REQUIRED
	ResponseMetadata BanRoomUserResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result BanRoomUserResResult `json:"Result"`
}

type BanRoomUserResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *BanRoomUserResResponseMetadataError `json:"Error,omitempty"`
}

// BanRoomUserResResponseMetadataError - 仅在请求失败时返回。
type BanRoomUserResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type BanRoomUserResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 "success", 失败时为空。
	Message string `json:"Message"`
}

type BanUserStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 你需要封禁音/视频流的用户的 ID
	UserID string `json:"UserId"`

	// 是否解封音频流。
	// * true：解封音频流。
	// * false：封禁音频流。
	// 默认值为 true。
	Audio *bool `json:"Audio,omitempty"`

	// 封禁时长，取值范围为{0,[60,259200]}，单位为秒。 若传入值为空或 0 表示封禁音视频流至主动调用unbanroomuser 接口解封。 若传入值大于 0，且小于10，自动修改为10。 若传入值大于259200，自动修改为259200。
	ForbiddenInterval *int64 `json:"ForbiddenInterval,omitempty"`

	// 是否解封视频流。
	// * true：解封视频流。
	// * false：封禁视频流。
	// 默认值为 true。
	Video *bool `json:"Video,omitempty"`
}

type BanUserStreamRes struct {

	// REQUIRED
	ResponseMetadata BanUserStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *BanUserStreamResResult          `json:"Result,omitempty"`
}

type BanUserStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *BanUserStreamResResponseMetadataError `json:"Error,omitempty"`
}

// BanUserStreamResResponseMetadataError - 仅在请求失败时返回。
type BanUserStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type BanUserStreamResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 "success", 失败时为空。
	Message string `json:"Message"`
}

type CreateAppBody struct {

	// REQUIRED; * 应用名称
	// * 命名规则：字符串中只能包含中文字符、英文大小写字符、数字和下划线；长度不能超过24个字符。
	AppName string `json:"AppName"`
}

type CreateAppRes struct {

	// REQUIRED
	ResponseMetadata CreateAppResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateAppResResult `json:"Result"`
}

type CreateAppResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *CreateAppResResponseMetadataError `json:"Error,omitempty"`
}

// CreateAppResResponseMetadataError - 仅在请求失败时返回。
type CreateAppResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type CreateAppResResult struct {

	// REQUIRED; 应用 ID
	AppID string `json:"AppId"`

	// REQUIRED; 主 AppKey，用于生成 Token
	AppKey string `json:"AppKey"`

	// REQUIRED; 创建时间
	CreateDate string `json:"CreateDate"`

	// REQUIRED; 实例 ID
	InstanceID string `json:"InstanceId"`

	// REQUIRED; * 服务状态。取值及含义如下：
	// * 0: 创建中——指 AppID 在初始化
	// * 1: 运行中——指当前 AppID 为正常服务状态
	// * 6: 欠费关停
	// * 98: 已停用——指调用服务端停用接口，当前 AppID 被设置为不可用状态
	InstanceStatus int32 `json:"InstanceStatus"`

	// REQUIRED; 应用名称
	Name string `json:"Name"`

	// REQUIRED; 应用所属的账号 ID
	Owner string `json:"Owner"`

	// REQUIRED; 副 AppKey，启用后可用于生成 Token
	SecondaryAppKey string `json:"SecondaryAppKey"`
}

type Event20231101 struct {

	// 原因，只有error事件有
	Cause *string `json:"Cause,omitempty"`

	// 事件描述
	Description *string `json:"Description,omitempty"`

	// 原始日志信息，这一期不做在前端展示
	Detail *string `json:"Detail,omitempty"`

	// 展示名
	DisplayName *string `json:"DisplayName,omitempty"`

	// 事件级别 Info、Warning、Error
	Level *int64 `json:"Level,omitempty"`

	// 事件名，英文
	Name *string `json:"Name,omitempty"`

	// 子事件，构成该聚合事件的事件，这一期不做在前端展示
	SubEvents []*Event20231101SubEventsItem `json:"SubEvents,omitempty"`

	// 标签
	Tags []*string `json:"Tags,omitempty"`

	// 事件时间，若为聚合类事件，取最大时间
	Time *int64 `json:"Time,omitempty"`
}

type Event20231101SubEventsItem struct {

	// 原因，只有error事件有
	Cause *string `json:"Cause,omitempty"`

	// 事件描述
	Description *string `json:"Description,omitempty"`

	// 原始日志信息，这一期不做在前端展示
	Detail *string `json:"Detail,omitempty"`

	// 展示名
	DisplayName *string `json:"DisplayName,omitempty"`

	// 事件级别 Info、Warning、Error
	Level *int64 `json:"Level,omitempty"`

	// 事件名，英文
	Name *string `json:"Name,omitempty"`

	// 子事件，构成该聚合事件的事件，这一期不做在前端展示
	SubEvents []*Event20231101 `json:"SubEvents,omitempty"`

	// 标签
	Tags []*string `json:"Tags,omitempty"`

	// 事件时间，若为聚合类事件，取最大时间
	Time *int64 `json:"Time,omitempty"`
}

type GetPushMixedStreamToCDNTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 要查询的转推直播任务 ID。通过服务端发起时，该值为调用 OpenAPI 时传入的 TaskId。通过客户端 SDK 发起时，TaskId 是按照 userId@@taskId 格式拼接而成的字符串；当传入的 taskId 为空时，这里的
	// TaskId 为 userId。
	// TaskId和UserId均为非必填参数，但是你需要至少填一个参数以保证可以正常发起请求。
	TaskID *string `json:"TaskId,omitempty" query:"TaskId"`

	// 客户端发起转推任务的用户 ID。
	// 你在客户端发起多个任务，当使用该接口进行查询时：
	// * 查询第一个任务时，UserId 可以传入发起转推任务的用户 UserId，TaskId 传入空字符串；或直接将该用户的 UserId 传入 TaskId。
	// * 查询第二个以上任务时，UserId 和 TaskId 为发起转推任务的用户 UserId 和 TaskId。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type GetPushMixedStreamToCDNTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPushMixedStreamToCDNTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetPushMixedStreamToCDNTaskResResult          `json:"Result,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetPushMixedStreamToCDNTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResponseMetadataError - 仅在请求失败时返回。
type GetPushMixedStreamToCDNTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResult struct {

	// 合流转推任务信息
	PushMixedStreamToCDNTask *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask `json:"PushMixedStreamToCDNTask,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask - 合流转推任务信息
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTask struct {
	Control *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControl `json:"Control,omitempty"`

	// 音视频编码参数。
	// * 单流录制时，你仅可以设置VideoFps、VideoBitrate和所有音频相关参数。
	// * 合流录制时，你仅可以设置VideoWidth、VideoHeight、VideoFps、VideoBitrate和所有音频相关参数。
	Encode *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskEncode `json:"Encode,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime        *int64                                                                      `json:"EndTime,omitempty"`
	ExcludeStreams *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreams `json:"ExcludeStreams,omitempty"`
	Layout         *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayout         `json:"Layout,omitempty"`

	// 推流状态。
	// * 0：运行中，未获取到任务状态，建议稍后重新查询
	// * 1：未开始推流
	// * 2：首次连接 CDN 服务
	// * 3：正在重连 CDN 服务
	// * 4：连接 CDN 服务成功，正在尝试推流。
	// * 5：连接 CDN 服务成功，推流成功
	// * 6：已停止推流。
	// 仅当Status=2 时，PushStreamState 有实际意义；当Status=3 时，PushStreamState=6; status 为其他值时，PushStreamState 均为0。
	PushStreamState *int32 `json:"PushStreamState,omitempty"`

	// 推流 CDN 地址。仅支持 RTMP 协议。
	PushURL *string `json:"PushURL,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0：未知异常状态
	// * 1：未开始
	// * 2： 运行中
	// * 3： 已结束
	// * 4： 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因。
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：客户端用户主动退房/调用停止转推接口
	StopReason    *string                                                                    `json:"StopReason,omitempty"`
	TargetStreams *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreams `json:"TargetStreams,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControl struct {

	// 选择补帧模式。支持取值及含义如下：
	// * 0：补最后一帧。
	// * 1：补黑帧
	// 默认值为0。 自动布局模式下，没有补帧的逻辑。 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系: 你可以在 Region 中传入
	// Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为 [10, 86400],单位为秒，默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。支持取值及含义如下：
	// * 0 ：音视频；
	// * 1 ：纯音频。
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 转推直播推流模式，用于控制触发推流的时机。支持取值及含义如下：
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。
	// 默认值为 0。 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                                                            `json:"PushStreamMode,omitempty"`
	SEIConfig      *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSEIConfig     `json:"SEIConfig,omitempty"`
	SpatialConfig  *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfig `json:"SpatialConfig,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSEIConfig struct {

	// SEI 中是否包含用户说话音量值。
	// * false：否。
	// * true：是。
	// 默认值为 false。
	AddVolumeValue *bool `json:"AddVolumeValue,omitempty"`

	// 是否透传 RTC 流里的 SEI 信息。
	// * true：是。
	// * false：否。
	// 默认值为 true。
	PassthroughRTCSEIMessage *bool `json:"PassthroughRTCSEIMessage,omitempty"`

	// 开启音量指示模式后，非关键帧携带 SEI 包含的信息类型。支持取值及含义如下：
	// * 0：全量信息。
	// * 1：只有音量信息和时间戳。
	// 默认值为 0。关于 SEI 结构，参看服务端合流转推 SEI 结构 [1163740#sei]
	SEIContentMode *int32 `json:"SEIContentMode,omitempty"`

	// 设置 SEI 的 Payload Type，对 服务端合流转推 SEI 和 RTC 透传SEI 均生效。支持取值为 5 或 100，默认为 100。
	SEIPayLoadType *int32 `json:"SEIPayLoadType,omitempty"`

	// 该字段为长度为 32 位的 16 进制字符，每个字符的范围为 a-f，A-F，0-9，不可包含 -。如果校验失败，会返回错误码 InvalidParameter。
	// 仅当 SEIPayLoadType=5时，该字段需要填写，SEIPayLoadType=100时，该字段内容会被忽略。
	// 此参数仅对 RTC 透传SEI 生效。当用户设置 SEIPayloadType = 5 时，服务端合流转推SEI 的 SEIPayloadUUID 为固定值：566f6c63656e67696e65525443534549，对应16位字符串
	// VolcengineRTCSEI。
	SEIPayloadUUID *string `json:"SEIPayloadUUID,omitempty"`

	// 设置SEI数据结构 [1163740#sei]中 app_data 参数的值，最大长度为 4096。此参数支持动态更新。
	UserConfigExtraInfo *string `json:"UserConfigExtraInfo,omitempty"`

	// 用户说话音量的回调间隔。取值范围为 [0.3,∞]，单位为秒，默认值为 2。仅当用户说话音量发生变化时此回调才会被触发。
	VolumeIndicationInterval *float32 `json:"VolumeIndicationInterval,omitempty"`

	// 是否开启音量指示模式。
	// * true：是。此时非关键帧中也可能携带 SEI 信息。
	// * false：否。此时只会在关键帧中携带 SEI 信息。
	// 默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。 关于音量指示模式的用法，参看 音量指示模式 [1163740#volume]。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfig struct {
	AudienceSpatialOrientation *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。
	// * false：关闭。
	// * true：开启
	// 默认值为 false。
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskControlSpatialConfigAudienceSpatialOrientation struct {

	// 前方朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 头顶朝向
	Up []*float32 `json:"Up,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskEncode - 音视频编码参数。
// * 单流录制时，你仅可以设置VideoFps、VideoBitrate和所有音频相关参数。
// * 合流录制时，你仅可以设置VideoWidth、VideoHeight、VideoFps、VideoBitrate和所有音频相关参数。
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps。
	// * 该参数若不填，RTC 会根据音频配置文件类型、采样率、声道数的配置生成音频码率取值。
	// * 该参数若填写，实际生效的码率也会随着音频配置文件类型、采样率、声道数的配置而变化。若你发现生效的码率值没有达到你设定的值，可能是该值已经超过该场景下的极限码率。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。支持取值及含义如下：
	// * 1：单声道。
	// * 2：双声道。
	// 默认值为 2。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。仅支持取 0，表示 AAC 编码协议。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型。支持取值及含义如下：
	// * 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。此时，只支持输出流音频声道数为双声道。
	// 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。可取值为：{32000,44100，48000}，单位为 Hz，默认值为 48000。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 视频码率。取值范围为 [0,10000]，单位为 Kbps，默认值为 0。0 表示自适应码率，会自动根据 VideoFps , VideoWidth 以及VideoHeight 计算出合理的码率。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。支持取值及含义如下：
	// * 0：H.264
	// * 1：ByteVC1
	// 默认值为 0。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 视频帧率。取值范围为 [1,60]，单位为 FPS，默认值为 15。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。取值范围为 [1,5]，单位为秒，默认值为 4。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 画面高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Height 设置画面宽度。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 画面宽度。取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Width 设置画面宽度。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayout struct {
	CustomLayout *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。支持取值及含义如下：
	// * 0：自适应布局 [1167930#adapt]。
	// * 1：垂直布局 [1167930#vertical]
	// * 2 ：自定义布局。
	// * 3 ：并排布局 [1167930#horizontal]
	// 默认值为0
	LayoutMode      *int32                                                                             `json:"LayoutMode,omitempty"`
	MainVideoStream *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayout struct {
	Canvas *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。
	Regions []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色, 范围为 #000000 ~ #ffffff (大小写均可)，格式为 #RGB(16进制)，默认值为 #000000（黑色）。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。 如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。
	Width *int32 `json:"Width,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]，默认值为 1.0。0.0 表示完全透明，1.0 表示完全不透明。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 [0,1]。默认值为 0，表示没有圆角效果。 圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。 例如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 ：按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 ：按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形。
	// 默认值为 0。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]，默认值为 0。0 表示该区域图像位于最下层，100 表示该区域图像位于最上层。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskLayoutMainVideoStream struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetPushMixedStreamToCDNTaskResResultPushMixedStreamToCDNTaskTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetPushSingleStreamToCDNTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 要查询的转推直播任务 ID。通过服务端发起时，该值为调用 OpenAPI 时传入的 TaskId。通过客户端 SDK 发起时，TaskId 是按照 userId@@taskId 格式拼接而成的字符串；当传入的 taskId 为空时，这里的
	// TaskId 为 userId。
	// TaskId和UserId均为非必填参数，但是你需要至少填一个参数以保证可以正常发起请求。
	TaskID *string `json:"TaskId,omitempty" query:"TaskId"`

	// 客户端发起转推任务的用户 ID。
	// 你在客户端发起多个任务，当使用该接口进行查询时：
	// * 查询第一个任务时，UserId 可以传入发起转推任务的用户 UserId，TaskId 传入空字符串；或直接将该用户的 UserId 传入 TaskId。
	// * 查询第二个以上任务时，UserId 和 TaskId 为发起转推任务的用户 UserId 和 TaskId。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type GetPushSingleStreamToCDNTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPushSingleStreamToCDNTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetPushSingleStreamToCDNTaskResResult          `json:"Result,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetPushSingleStreamToCDNTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetPushSingleStreamToCDNTaskResResponseMetadataError - 仅在请求失败时返回。
type GetPushSingleStreamToCDNTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResult struct {

	// 单流转推任务信息
	PushSingleStreamToCDNTask *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask `json:"PushSingleStreamToCDNTask,omitempty"`
}

// GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask - 单流转推任务信息
type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTask struct {
	Control *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskControl `json:"Control,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 推流状态
	// * 0：未获取到任务状态。建议稍后重新查询
	// * 1：未开始推流
	// * 2：首次连接 CDN 服务
	// * 3：正在重连 CDN 服务
	// * 4：连接 CDN 服务成功，正在尝试推流。
	// * 5：连接 CDN 服务成功，推流成功
	// * 6：已停止推流。
	// 仅当Status=2 时，PushStreamState 有实际意义；当Status=3 时，PushStreamState=6; status 为其他值时，PushStreamState 均为0。
	PushStreamState *int32 `json:"PushStreamState,omitempty"`

	// 推流地址。目前仅支持 rtmp 协议
	PushURL *string `json:"PushURL,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：客户端用户主动退房/调用停止转推接口
	StopReason *string                                                               `json:"StopReason,omitempty"`
	Stream     *GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskStream `json:"Stream,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskControl struct {

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为 [10, 86400], 单位为秒，默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为 0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type GetPushSingleStreamToCDNTaskResResultPushSingleStreamToCDNTaskStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetRecordTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的云端录制任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata GetRecordTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetRecordTaskResResult          `json:"Result,omitempty"`
}

type GetRecordTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetRecordTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetRecordTaskResResponseMetadataError - 仅在请求失败时返回。
type GetRecordTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetRecordTaskResResult struct {

	// -
	RecordTask *GetRecordTaskResResultRecordTask `json:"RecordTask,omitempty"`
}

// GetRecordTaskResResultRecordTask - -
type GetRecordTaskResResultRecordTask struct {

	// 任务结束的时间，为 Unix 时间戳，单位为毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 录制生成的文件列表。
	RecordFileList []*GetRecordTaskResResultRecordTaskRecordFileListItem `json:"RecordFileList,omitempty"`

	// 任务开始的时间，为 Unix 时间戳，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。取值及含义如下：
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因。取值及含义如下：
	// * 空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动通过 API 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：超过了最大空闲时间
	StopReason *string `json:"StopReason,omitempty"`
}

type GetRecordTaskResResultRecordTaskRecordFileListItem struct {

	// 音频录制编码器
	AudioCodec *string `json:"AudioCodec,omitempty"`

	// 文件时长，单位为毫秒。
	Duration *int64 `json:"Duration,omitempty"`

	// 文件在对象存储平台中的完整路径，如abc/efg/123.mp4。仅在你选择配置存储到对象存储平台时，此参数有效。
	ObjectKey *string `json:"ObjectKey,omitempty"`

	// 文件大小，单位为字节。
	Size *int64 `json:"Size,omitempty"`

	// 当前录制文件创建的时间，为 Unix 时间戳，单位为毫秒。
	StartTime *int64 `json:"StartTime,omitempty"`

	// 录制文件中包含流的列表。
	StreamList []*GetRecordTaskResResultRecordTaskRecordFileListPropertiesItemsItem `json:"StreamList,omitempty"`

	// 文件在火山引擎视频点播 VOD [https://www.volcengine.com/product/vod] 平台的唯一标识。你可以根据 vid 在点播平台上找到对应的文件。仅在你选择配置存储到 Vod 平台时，此参数有效。
	Vid *string `json:"Vid,omitempty"`

	// 视频录制编码协议
	VideoCodec *string `json:"VideoCodec,omitempty"`

	// 录制视频高度，单位为像素。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 录制视频宽度，单位为像素。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type GetRecordTaskResResultRecordTaskRecordFileListPropertiesItemsItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetRoomOnlineUsersQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 指定房间 ID
	RoomID string `json:"RoomId" query:"RoomId"`
}

type GetRoomOnlineUsersRes struct {

	// REQUIRED
	ResponseMetadata GetRoomOnlineUsersResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetRoomOnlineUsersResResult          `json:"Result,omitempty"`
}

type GetRoomOnlineUsersResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetRoomOnlineUsersResResponseMetadataError `json:"Error,omitempty"`
}

// GetRoomOnlineUsersResResponseMetadataError - 仅在请求失败时返回。
type GetRoomOnlineUsersResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetRoomOnlineUsersResResult struct {

	// 不可见用户列表
	InvisibleUserList []*string `json:"InvisibleUserList,omitempty"`

	// 查询的房间是否存在。
	// * true：存在。
	// * false：不存在。
	// 当 RoomExists 的值为 false 时，不会返回其他字段。
	RoomExists *bool `json:"RoomExists,omitempty"`

	// 查询到的不可见用户总数。不可见用户最多返回 10000 名。
	TotalInvisibleUser *int32 `json:"TotalInvisibleUser,omitempty"`

	// 查询到的用户总数
	TotalUser *int32 `json:"TotalUser,omitempty"`

	// 可见用户列表
	VisibleUserList []*string `json:"VisibleUserList,omitempty"`
}

type GetSegmentTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的音频切片任务 ID。自动切片任务下，该字段可传入用户 UserId。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetSegmentTaskRes struct {

	// REQUIRED
	ResponseMetadata GetSegmentTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetSegmentTaskResResult          `json:"Result,omitempty"`
}

type GetSegmentTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetSegmentTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetSegmentTaskResResponseMetadataError - 仅在请求失败时返回。
type GetSegmentTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetSegmentTaskResResult struct {

	// 音频切片任务信息
	SegmentTask *GetSegmentTaskResResultSegmentTask `json:"SegmentTask,omitempty"`
}

// GetSegmentTaskResResultSegmentTask - 音频切片任务信息
type GetSegmentTaskResResultSegmentTask struct {
	Control *GetSegmentTaskResResultSegmentTaskControl `json:"Control,omitempty"`

	// 每个音频切片的时长。
	Duration *int32 `json:"Duration,omitempty"`

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime *int64 `json:"EndTime,omitempty"`

	// 是否在开启音视频切片时，立刻开始切片。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。
	Identifier *string `json:"Identifier,omitempty"`

	// 任务最大的空闲超时时间。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因。
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：自动切片任务中，切片任务对应的客户端用户主动退房。
	StopReason *string `json:"StopReason,omitempty"`

	// 录制文件的存储平台配置。
	// 支持：
	// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * Amazon S3
	// * 阿里云对象存储 OSS
	// * 华为云 OBS
	// * 腾讯云 COS
	// * 七牛云 Kodo。
	StorageConfig *GetSegmentTaskResResultSegmentTaskStorageConfig `json:"StorageConfig,omitempty"`
	TargetStreams *GetSegmentTaskResResultSegmentTaskTargetStreams `json:"TargetStreams,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskControl struct {

	// 是否开启切片对齐功能，你可以使用该功能，对齐各个用户音频切片的开始和结束时刻。
	// * false：关闭音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。如果用户在切片的周期中，有部分时间未发布音频，返回的音频切片时长会小于切片周期。各个用户音频切片开始时间不一定一致。
	// * true：开启音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。切片长度和切片周期相等，且各个用户音频切片开始的时间一致。如果用户在切片的周期中，有部分时间未发布音频，切片长度不变，这段时间呈现静音帧。如果用户在某个切片周期中始终没有发布音频，则不生成音频切片。
	// 默认值为false。
	Align *bool `json:"Align,omitempty"`

	// 是否忽略静音切片。
	// * true：忽略。
	// * false：不忽略。
	// 默认值为 false
	IgnoreSilence *bool `json:"IgnoreSilence,omitempty"`

	// 是否开启合流切片功能。
	// * False：只会对 TargetStreams 中指定的音频流分别切片。
	// * True：除了会对 TargetStreams 中指定的音频流分别切片，还会对指定的音频流进行混音，生成合流切片，合流切片对应的用户名为 mixed。此时，任务创建后，不管是否有人上麦，会持续回调混音切片。
	// 默认值为 false。 不同平台的回调参看：
	// 操作 ANDROID API IOS API WINDOWS API
	// 本地麦克风录制和远端所有用户混音后的音频数据回调 onMixedAudioFrame [70081#onmixedaudioframe] onMixedAudioFrame: [70087#onmixedaudioframe] onMixedAudioFrame
	// [70096#onmixedaudioframe]
	Mixed *bool `json:"Mixed,omitempty"`

	// 冗余切片时长，单位为毫秒。
	// 当前 RTC 按照传入的Duration值进行固定时长切片,可能存在敏感词被截断，未被识别的情况。此时你可以添加冗余切片，即上一段切片的末尾指定时长，来避免此情况，此时切片的时长变为RedundantDuration + Duration。
	// 例如：当 Duration = 20，redundantDuration = 3 时，最终输出的前三个文件时长为：[0-20], [17-40],
	// [37-60]。
	RedundantDuration *int32 `json:"RedundantDuration,omitempty"`
}

// GetSegmentTaskResResultSegmentTaskStorageConfig - 录制文件的存储平台配置。
// 支持：
// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
// * Amazon S3
// * 阿里云对象存储 OSS
// * 华为云 OBS
// * 腾讯云 COS
// * 七牛云 Kodo。
type GetSegmentTaskResResultSegmentTaskStorageConfig struct {
	CustomConfig *GetSegmentTaskResResultSegmentTaskStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *GetSegmentTaskResResultSegmentTaskStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型。支持取值及含义如下：
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 支持 S3 协议的第三方存储平台
	// * 3: VeImageX (当前仅支持抽帧截图功能)
	// 默认值为 0。
	Type           *int32                                                         `json:"Type,omitempty"`
	VeImageXConfig *GetSegmentTaskResResultSegmentTaskStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *GetSegmentTaskResResultSegmentTaskStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。支持取值及含义如下：
	// * 0：Amazon S3
	// * 1：阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。
	// 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	// * 此账号 ID 为火山引擎主账号 ID。
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。 你也可以通过
	// OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为 0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 上传到视频点播平台时, 是否需要根据文件后缀自动设置 FileExtension。关于 FileExtension 的详细说明，参看 FileExtension [https://www.volcengine.com/docs/4/164295#fileextension]。
	// * true：需要；
	// * false：不需要。
	// 默认值为 false。
	AutoSetFileExtension *bool `json:"AutoSetFileExtension,omitempty"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 上传到视频点播平台时, 文件的存储类型。支持取值及含义如下：：
	// * 1：标准存储。
	// * 2：归档存储。
	// * 3：低频存储。
	// 默认值为 1。 关于存储类型的详细说明，参看媒资存储存储类型 [https://www.volcengine.com/docs/4/73629#%E5%AA%92%E8%B5%84%E5%AD%98%E5%82%A8]
	StorageClass *int32 `json:"StorageClass,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetSegmentTaskResResultSegmentTaskTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetSegmentTaskResResultSegmentTaskTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetSnapshotTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的抽帧截图任务 ID。自动抽帧任务下，该字段可传入用户 UserId。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata GetSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetSnapshotTaskResResult          `json:"Result,omitempty"`
}

type GetSnapshotTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetSnapshotTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetSnapshotTaskResResponseMetadataError - 仅在请求失败时返回。
type GetSnapshotTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetSnapshotTaskResResult struct {

	// 抽帧截图任务信息
	SnapshotTask *GetSnapshotTaskResResultSnapshotTask `json:"SnapshotTask,omitempty"`
}

// GetSnapshotTaskResResultSnapshotTask - 抽帧截图任务信息
type GetSnapshotTaskResResultSnapshotTask struct {

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime     *int64                                           `json:"EndTime,omitempty"`
	ImageConfig *GetSnapshotTaskResResultSnapshotTaskImageConfig `json:"ImageConfig,omitempty"`

	// 任务最大的空闲超时时间。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * IdleTimeOut：任务超过最大空闲时间
	// * UserDisconnect：自动截图任务中，截图任务对应的客户端用户主动退房。
	StopReason *string `json:"StopReason,omitempty"`

	// 录制文件的存储平台配置。
	// 支持：
	// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * Amazon S3
	// * 阿里云对象存储 OSS
	// * 华为云 OBS
	// * 腾讯云 COS
	// * 七牛云 Kodo。
	StorageConfig *GetSnapshotTaskResResultSnapshotTaskStorageConfig `json:"StorageConfig,omitempty"`
	TargetStreams *GetSnapshotTaskResResultSnapshotTaskTargetStreams `json:"TargetStreams,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskImageConfig struct {

	// 图片的格式。支持取值及含义如下：
	// * 0：JEPG
	// * 1：PNG
	// 默认值为0。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，取值范围为 [0, 1920]，单位为像素，默认值为 0，此时，和视频流的实际高度相同。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，取值范围为[1, 600]，单位为秒，默认值为2。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，取值范围为[0, 1920]，单位为像素。默认值为0，表示和视频流的实际宽度相同。
	Width *int32 `json:"Width,omitempty"`
}

// GetSnapshotTaskResResultSnapshotTaskStorageConfig - 录制文件的存储平台配置。
// 支持：
// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
// * Amazon S3
// * 阿里云对象存储 OSS
// * 华为云 OBS
// * 腾讯云 COS
// * 七牛云 Kodo。
type GetSnapshotTaskResResultSnapshotTaskStorageConfig struct {
	CustomConfig *GetSnapshotTaskResResultSnapshotTaskStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *GetSnapshotTaskResResultSnapshotTaskStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型。支持取值及含义如下：
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 支持 S3 协议的第三方存储平台
	// * 3: VeImageX (当前仅支持抽帧截图功能)
	// 默认值为 0。
	Type           *int32                                                           `json:"Type,omitempty"`
	VeImageXConfig *GetSnapshotTaskResResultSnapshotTaskStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *GetSnapshotTaskResResultSnapshotTaskStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。支持取值及含义如下：
	// * 0：Amazon S3
	// * 1：阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。
	// 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	// * 此账号 ID 为火山引擎主账号 ID。
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。 你也可以通过
	// OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为 0。
	Region *int32 `json:"Region,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 上传到视频点播平台时, 是否需要根据文件后缀自动设置 FileExtension。关于 FileExtension 的详细说明，参看 FileExtension [https://www.volcengine.com/docs/4/164295#fileextension]。
	// * true：需要；
	// * false：不需要。
	// 默认值为 false。
	AutoSetFileExtension *bool `json:"AutoSetFileExtension,omitempty"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 上传到视频点播平台时, 文件的存储类型。支持取值及含义如下：：
	// * 1：标准存储。
	// * 2：归档存储。
	// * 3：低频存储。
	// 默认值为 1。 关于存储类型的详细说明，参看媒资存储存储类型 [https://www.volcengine.com/docs/4/73629#%E5%AA%92%E8%B5%84%E5%AD%98%E5%82%A8]
	StorageClass *int32 `json:"StorageClass,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*GetSnapshotTaskResResultSnapshotTaskTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type GetSnapshotTaskResResultSnapshotTaskTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type GetWebCastTaskQuery struct {

	// REQUIRED; 你的音视频应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 要查询的云录屏任务 ID。
	TaskID string `json:"TaskId" query:"TaskId"`
}

type GetWebCastTaskRes struct {

	// REQUIRED
	ResponseMetadata GetWebCastTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *GetWebCastTaskResResult          `json:"Result,omitempty"`
}

type GetWebCastTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *GetWebCastTaskResResponseMetadataError `json:"Error,omitempty"`
}

// GetWebCastTaskResResponseMetadataError - 仅在请求失败时返回。
type GetWebCastTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type GetWebCastTaskResResult struct {

	// 云录屏任务信息
	WebCastTask *GetWebCastTaskResResultWebCastTask `json:"WebCastTask,omitempty"`
}

// GetWebCastTaskResResultWebCastTask - 云录屏任务信息
type GetWebCastTaskResResultWebCastTask struct {

	// 任务结束时间戳，Unix 时间，单位为毫秒。0 表示任务未结束
	EndTime           *int64                                               `json:"EndTime,omitempty"`
	EventNotifyConfig *GetWebCastTaskResResultWebCastTaskEventNotifyConfig `json:"EventNotifyConfig,omitempty"`

	// 最大运行时间，超过此时间后，任务自动终止。单位为秒。取值范围为 [10,86400]，默认值为 86400。不填时自动调整为默认值。
	MaxRunningTime *int32                                           `json:"MaxRunningTime,omitempty"`
	MonitorConfig  *GetWebCastTaskResResultWebCastTaskMonitorConfig `json:"MonitorConfig,omitempty"`

	// 需要转推的网页地址，可以携带自定义的 queryParams 来鉴权等，总体长度不超过 1024。
	SourceURL *string `json:"SourceURL,omitempty"`

	// 任务开始时间戳，Unix 时间，单位为毫秒
	StartTime *int64 `json:"StartTime,omitempty"`

	// 任务状态。支持取值及含义如下：
	// * 0: 未知异常状态
	// * 1: 未开始
	// * 2: 运行中
	// * 3: 已结束
	// * 4: 任务运行失败
	Status *int64 `json:"Status,omitempty"`

	// 任务停止的原因。支持取值及含义如下：
	// * 返回为空：表示任务未结束
	// * UnknownStopReason：未知停止原因
	// * StopByAPI：用户主动调用 服务端 OpenAPI 停止
	// * StartTaskFailed：任务启动失败
	// * ExceedMaxRunningTime：任务超过最大运行时间
	StopReason *string `json:"StopReason,omitempty"`

	// 推送网页音视频内容的用户对应的 UserId
	UserID *string `json:"UserId,omitempty"`

	// 输出的视频参数，最多支持 2 路，以大小流的方式支持接收端按需订阅，将以最大的视频流分辨率作为网页渲染分辨率，为空时按默认值填充一路
	VideoSolutions []*GetWebCastTaskResResultWebCastTaskVideoSolutionsItem `json:"VideoSolutions,omitempty"`
}

type GetWebCastTaskResResultWebCastTaskEventNotifyConfig struct {

	// 是否启用页面主动事件通知,。
	// * false：页面在打开后就会开始采集，在收到 StopWebCast openAPI 请求后结束采集。
	// * true：在页面中注入两个 JS 函数：onWebcastStart()和 onWebcastEnd()。 默认值为false。
	// 当页面判断资源加载完成之后调用onWebcastStart()，控制程序才会开始进行页面内容的采集。当页面判断本次任务内容已完成时调用onWebcastEnd() 通知控制程序结束本次任务。StopWebCast openAPI 效果不变，业务可提前结束任务。其他页面内容、JS
	// 线程的检测（若启用），将在收到 onWebcastStart()事件后才开始。
	// 当启用页面主动事件通知后，你可以参考以下示例代码来通知采集开始。
	// <script>
	// if (ready() && typeof onWebcastStart === 'function')
	// onWebcastStart();
	// </script>
	EnableEventNotify *bool `json:"EnableEventNotify,omitempty"`

	// 启用页面主动事件通知后，等待开始事件的超时时间。取值范围为 [0,60]，单位为秒。默认值为 0，表示不启用。仅当 EnableEventNotify 为 true 时，此参数有效。
	// * 当在超时时间内收到开始事件，采集功能正常运行，用户将收到 Status=1的回调。
	// * 当超时时间内未收到开始事件，将进行刷新，等待时间被重置，再次发生超时后将进行任务重调度。刷新时将回调 Status=4，Reason=" StartEventTimeout"。重调度时将回调 Status=5，Reason="StartEventTimeout"。
	StartTimeout *int32 `json:"StartTimeout,omitempty"`
}

type GetWebCastTaskResResultWebCastTaskMonitorConfig struct {

	// 对页面是否白屏的检测间隔。取值范围为 [2,30]，单位为秒。默认值为0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageBlank" 。
	// * 再次出现连续两次检测命中时将进行任务重调度，并回调Status=5，Reason="PageBlank"。
	// 注意：页面全白可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	BlankCheckInterval *int32 `json:"BlankCheckInterval,omitempty"`

	// 对页面 JS 线程是否崩溃/卡死的检测间隔。 取值范围为 [0,30]，单位为秒。默认值为 0，表示不启用。
	// 当出现检测命中时将进行任务重调度，并回调 Status=5，Reason="PageCrash"。
	CrashCheckInterval *int32 `json:"CrashCheckInterval,omitempty"`

	// 对页面内容是否无变化的检测间隔。取值范围为 [2,30]，单位为秒。默认值为 0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageFreeze"。
	// * 再次出现连续两次检测命中时，将进行任务重调度，并回调Status=5，Reason="PageFreeze"。
	// 注意：页面无变化可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	FreezeCheckInterval *int32 `json:"FreezeCheckInterval,omitempty"`
}

type GetWebCastTaskResResultWebCastTaskVideoSolutionsItem struct {

	// 最大发送码率，取值范围为 [0,10000]，单位为 Kbps，默认值 0，为 0 时表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，单位为 fps，范围为[1,60]，默认值为 15。帧率和码率设置建议参照视频发布参数对照表 [70122#param]以获取最佳体验。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 视频高度，单位为像素，范围为 [50,1080]，默认值为 720。必须是偶数，值为奇数时自动调整为偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频宽度，单位为像素，范围为 [50,1920]，默认值为 1280。必须是偶数，值为奇数时自动调整为偶数。
	Width *int32 `json:"Width,omitempty"`
}

type IndicatorTag struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                      `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTagSubTagsItem `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTag202201 struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                            `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag202201SubTagsItem `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTag202201SubTagsItem struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool                 `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag202201 `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type IndicatorTagSubTagsItem struct {
	Alias *string `json:"Alias,omitempty"`

	// 是否隐藏掉地域分布信息，注意这个参数只有在质量概览模块有效
	HidenDistribution *bool   `json:"HidenDistribution,omitempty"`
	IsPositive        *bool   `json:"IsPositive,omitempty"`
	Name              *string `json:"Name,omitempty"`

	// 可以使用采样归因功能
	SampleAvailable *bool           `json:"SampleAvailable,omitempty"`
	SubTags         []*IndicatorTag `json:"SubTags,omitempty"`

	// 是否支持用户分析
	SupportDetailAnalysis *bool `json:"SupportDetailAnalysis,omitempty"`

	// 用量统计是否支持切换新旧展示类型
	SupportSwitchDisplayType *bool `json:"SupportSwitchDisplayType,omitempty"`

	// 是否支持用户占比统计
	SupportUserProportion *bool `json:"SupportUserProportion,omitempty"`

	// 分布统计中取topN
	TopN *int64 `json:"TopN,omitempty"`

	// 分布统计中只取在这里面的值
	ValueDict []*string `json:"ValueDict,omitempty"`

	// 是否在控制台可见，若是，则在后面加🌟
	VisibleOnConsole *bool `json:"VisibleOnConsole,omitempty"`
}

type LimitTokenPrivilegeBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 限制 Token 发布权限时长，取值范围为 [60,259200]，单位为秒。 若传入值大于 0，且小于 60，自动修改为 60。 若传入值大于 259200，自动修改为 259200。
	ForbiddenInterval int32 `json:"ForbiddenInterval"`

	// REQUIRED; 指定房间 ID
	RoomID string `json:"RoomId"`

	// REQUIRED; 需要限制发布权限的 Token
	Token string `json:"Token"`

	// REQUIRED; 指定用户 ID
	UserID string `json:"UserId"`
}

type LimitTokenPrivilegeRes struct {

	// REQUIRED
	ResponseMetadata LimitTokenPrivilegeResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result LimitTokenPrivilegeResResult `json:"Result"`
}

type LimitTokenPrivilegeResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *LimitTokenPrivilegeResResponseMetadataError `json:"Error,omitempty"`
}

// LimitTokenPrivilegeResResponseMetadataError - 仅在请求失败时返回。
type LimitTokenPrivilegeResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type LimitTokenPrivilegeResResult struct {

	// REQUIRED; Token 权限限制结束时间，Unix 时间，单位为秒
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 请求成功时返回 “Success"，失败时为空
	Message string `json:"Message"`
}

type ListAppsBody struct {

	// 已创建的应用的 ID
	AppID *string `json:"AppId,omitempty"`

	// 返回条目数的上限，默认为 50，最大为 100
	Limit *string `json:"Limit,omitempty"`

	// * 偏移量，单位为条；
	// * 因单次查询返回结果数量有限，如果需要大量查询，需要借助偏移量实现分页查询；
	Offset *string `json:"Offset,omitempty"`

	// 传入项目名称，接口只返回传入项目包含的应用
	ProjectName *string `json:"ProjectName,omitempty"`

	// * 默认按照 App 创建时间升序返回；
	// * 设置为 1 时，降序返回；
	Reverse *int32 `json:"Reverse,omitempty"`
}

type ListAppsRes struct {

	// REQUIRED
	ResponseMetadata ListAppsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 查询的指定 AppId 的信息在 Result 中显示。
	Result ListAppsResResult `json:"Result"`
}

type ListAppsResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListAppsResResponseMetadataError `json:"Error,omitempty"`
}

// ListAppsResResponseMetadataError - 仅在请求失败时返回。
type ListAppsResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// ListAppsResResult - 查询的指定 AppId 的信息在 Result 中显示。
type ListAppsResResult struct {

	// REQUIRED; 应用详细信息
	AppList []ListAppsResResultAppListItem `json:"AppList"`

	// REQUIRED; 返回条目数的限制
	Limit int32 `json:"Limit"`

	// REQUIRED; 偏移量，单位为条
	Offset int32 `json:"Offset"`

	// REQUIRED; 当前条件查询到的 App 数量
	Total int32 `json:"Total"`
}

type ListAppsResResultAppListItem struct {

	// 应用 ID
	AppID *string `json:"AppId,omitempty"`

	// 主 AppKey，用于生成 Token
	AppKey *string `json:"AppKey,omitempty"`

	// 应用名称
	AppName *string `json:"AppName,omitempty"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitempty"`

	// 副 AppKey，启用后可用于生成 Token
	SecondaryAppKey *string `json:"SecondaryAppKey,omitempty"`

	// 服务状态。取值及含义如下：0: 创建中——指 AppID 在初始化1: 运行中——指当前 AppID 为正常服务状态6: 欠费关停98: 已停用——指调用服务端停用接口，当前 AppID 被设置为不可用状态
	Status *int32 `json:"Status,omitempty"`
}

type ListCallDetailQuery struct {

	// REQUIRED; 应用的唯一标志应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒
	StartTime string `json:"StartTime" query:"StartTime"`

	// REQUIRED; 查询的用户ID列表，多个用户ID用逗号隔开，最多可以指定 20 个。值不合法时默认剔除。
	UserID string `json:"UserId" query:"UserId"`

	// 获取数据的传输方向。枚举值为 0，1，2。
	// * 0:上下行
	// * 1:上行
	// * 2: 下行 默认值为0。不传时自动调整为默认值。
	Direction *string `json:"Direction,omitempty" query:"Direction"`
}

type ListCallDetailRes struct {
	ResponseMetadata *ListCallDetailResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *ListCallDetailResResult           `json:"Result,omitempty"`
}

type ListCallDetailResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *ListCallDetailResResponseMetadataError `json:"Error,omitempty"`
}

type ListCallDetailResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message     string  `json:"Message"`
	CodeN       *int32  `json:"CodeN,omitempty"`
	Description *string `json:"Description,omitempty"`
}

type ListCallDetailResResult struct {

	// 用户通话质量指标信息
	UserCallDetail []*ListCallDetailResResultUserCallDetailItem `json:"UserCallDetail,omitempty"`
}

type ListCallDetailResResultUserCallDetailItem struct {

	// 通话质量指标数据详情。
	CallDetail []*ListCallDetailResResultUserCallDetailPropertiesItemsItem `json:"CallDetail,omitempty"`

	// 获取数据的传输方向。
	Direction *string `json:"Direction,omitempty"`

	// 对端用户 Id，只有UserId用户处于下行时存在
	PeerID *string `json:"PeerId,omitempty"`

	// 查询房间 Id
	RoomID *string `json:"RoomId,omitempty"`

	// 查询用户 Id
	UserID *string `json:"UserId,omitempty"`
}

type ListCallDetailResResultUserCallDetailPropertiesItemsDataItem struct {

	// 该指标数据点的 Unix 时间戳，单位为毫秒。
	Time *int64 `json:"Time,omitempty"`

	// 该指标数据点的数据值
	Value *float64 `json:"Value,omitempty"`
}

type ListCallDetailResResultUserCallDetailPropertiesItemsItem struct {

	// 指标数据详情
	Data []*ListCallDetailResResultUserCallDetailPropertiesItemsDataItem `json:"Data,omitempty"`

	// 该指标中文名称
	Desc *string `json:"Desc,omitempty"`

	// 该项指标字段名称
	Name *string `json:"Name,omitempty"`

	// 指标单位
	Unit *string `json:"Unit,omitempty"`
}

type ListDetectionTaskQuery struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId" query:"RoomId"`

	// 用户 ID
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type ListDetectionTaskRes struct {

	// REQUIRED
	ResponseMetadata ListDetectionTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListDetectionTaskResResult `json:"Result,omitempty"`
}

type ListDetectionTaskResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListDetectionTaskResResponseMetadataError `json:"Error,omitempty"`
}

// ListDetectionTaskResResponseMetadataError - 仅在请求失败时返回。
type ListDetectionTaskResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// ListDetectionTaskResResult - 视请求的接口而定
type ListDetectionTaskResResult struct {

	// REQUIRED; 任务状态信息。如果从未对指定的用户发起审核，查询其审核状态时，返回的 EventData 值为空列表。
	EventData []ListDetectionTaskResResultEventDataItem `json:"EventData"`

	// REQUIRED; 仅在请求成功时返回success。
	Message string `json:"Message"`
}

// ListDetectionTaskResResultEventDataItem - 任务状态信息。如果从未对指定的用户发起审核，查询其审核状态时，返回的 EventData 值为空列表。
type ListDetectionTaskResResultEventDataItem struct {

	// REQUIRED; 用户审核发起时间戳，unix 时间，单位为秒。
	CreateTime int32 `json:"CreateTime"`

	// REQUIRED; 进行审核的内容类型，支持取值及含义如下：1：视频截图；2：音频切片；3：视频截图+音频切片（默认值）
	MediaType int32 `json:"MediaType"`

	// REQUIRED; 审核房间 ID
	RoomID string `json:"RoomId"`

	// REQUIRED; 审核状态，支持取值及含义如下：1：运行中：2：已完成；
	Status int32 `json:"Status"`

	// REQUIRED; 若开始审核接口填入 UserId，此处返回填入 UserId。 若开始审核接口未填入 UserId，此处返回房间内触发过审核任务的用户的 UserId。
	UserID string `json:"UserId"`
}

type ListOperationDataBody struct {

	// REQUIRED; 返回聚合时间的粒度，支持取值及含义如下：
	// * 1d：粒度为 1 天
	// * 1h：粒度为 1 小时
	AggregateGranularity string `json:"AggregateGranularity"`

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。Indicator 可选值，参看离线运营数据相关 indicator [1167931#indicator]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`
}

type ListOperationDataRes struct {

	// REQUIRED
	ResponseMetadata ListOperationDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListOperationDataResResult          `json:"Result,omitempty"`
}

type ListOperationDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListOperationDataResResponseMetadataError `json:"Error,omitempty"`
}

// ListOperationDataResResponseMetadataError - 仅在请求失败时返回。
type ListOperationDataResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListOperationDataResResult struct {

	// 返回聚合时间的粒度
	AggregateGranularity *string `json:"AggregateGranularity,omitempty"`

	// 具体指标数据
	Indicators []*ListOperationDataResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListOperationDataResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListOperationDataResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                           `json:"Name,omitempty"`
	Overview *ListOperationDataResResultIndicatorsItemOverview `json:"overview,omitempty"`

	// 指标单位
	Unit *string `json:"Unit,omitempty"`
}

type ListOperationDataResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListOperationDataResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListOperationDistributionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询的维度，一次仅支持查询一个维度。支持取值及含义如下：
	// * Province：一级行政区（包括港澳台地区）
	// * Country：国家
	Dimension string `json:"Dimension"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。
	// 目前仅支持查询 user_count，即此 AppId 在所选日期的通话总人数。通话人数按用户 id 去重。
	Indicator string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`
}

type ListOperationDistributionRes struct {

	// REQUIRED
	ResponseMetadata ListOperationDistributionResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListOperationDistributionResResult          `json:"Result,omitempty"`
}

type ListOperationDistributionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListOperationDistributionResResponseMetadataError `json:"Error,omitempty"`
}

// ListOperationDistributionResResponseMetadataError - 仅在请求失败时返回。
type ListOperationDistributionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListOperationDistributionResResult struct {

	// 具体的指标值以及对应时间
	Data []*ListOperationDistributionResResultDataItem `json:"Data,omitempty"`

	// 查询维度
	Dimension *string `json:"Dimension,omitempty"`

	// 指标名称
	Name *string `json:"Name,omitempty"`
}

type ListOperationDistributionResResultDataItem struct {

	// 维度名
	DistributionName *string `json:"DistributionName,omitempty"`

	// 占比
	Proportion *float64 `json:"Proportion,omitempty"`

	// 指标值
	Value *float64 `json:"Value,omitempty"`
}

type ListQualityBody struct {

	// REQUIRED; 返回聚合时间的粒度，支持取值及含义如下：
	// * 1d：粒度为 1 天
	// * 5min：粒度为 5 分钟
	AggregateGranularity string `json:"AggregateGranularity"`

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。Indicator 可选值，参看质量数据相关 indicator [1167931#listquality]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`

	// 网络类型，枚举值为：2g、3g、4g、5g、wifi
	// ProductType = web 时该参数不生效
	Access []*string `json:"Access,omitempty"`

	// 枚举值为：android、iOS、linux、mac、windows
	// ProductType = web 时该参数不生效
	OS []*string `json:"OS,omitempty"`

	// 要查询的产品类型，支持取值及含义如下：
	// * native：指 android、iOS、linux、mac或 windows 平台的 veRTC SDK。
	// * web：指 Web 平台的veRTC SDK。
	// 默认值为native
	ProductType *string `json:"ProductType,omitempty"`

	// 房间 ID。如果不填，代表查询该 AppId 的整体离线指标。
	RoomID *string `json:"RoomId,omitempty"`

	// 查询的用户 ID 列表，最多可以指定 20 个。值不合法时默认剔除。此字段仅在 RoomId 不为空时生效。
	UserID []*string `json:"UserId,omitempty"`
}

type ListQualityDistributionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询的维度，一次仅支持查询一个维度。支持取值及含义如下：
	// * Province：一级行政区（包括港澳台地区）
	// * Country：国家
	// * OS：用户设备平台。支持取值：android、ios、linux、mac、windows
	// * Access：用户网络类型。支持取值：2g、3g、4g、5g、wifi
	Dimension string `json:"Dimension"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称，一次仅支持查询一个指标。Indicator 支持取值参看质量数据相关 indicator [1167931#listquality]。
	Indicator string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`

	// 设备平台，目前支持 Android、iOS 和 Web。
	Platform *string `json:"Platform,omitempty"`

	// 要查询的数据所属设备端，支持取值及含义如下：
	// * native：指 Android、iOS、Linux、Mac 或 Windows 端。
	// * web：指 Web 端。 默认值为 native。
	ProductType *string `json:"ProductType,omitempty"`
}

type ListQualityDistributionRes struct {

	// REQUIRED
	ResponseMetadata ListQualityDistributionResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListQualityDistributionResResult          `json:"Result,omitempty"`
}

type ListQualityDistributionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListQualityDistributionResResponseMetadataError `json:"Error,omitempty"`
}

// ListQualityDistributionResResponseMetadataError - 仅在请求失败时返回。
type ListQualityDistributionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListQualityDistributionResResult struct {

	// 具体指标数据。
	Indicators []*ListQualityDistributionResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListQualityDistributionResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListQualityDistributionResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name *string `json:"Name,omitempty"`
}

type ListQualityDistributionResResultIndicatorsPropertiesItemsItem struct {

	// 具体的维度信息。
	Dimension *string `json:"Dimension,omitempty"`

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标数值，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListQualityRes struct {

	// REQUIRED
	ResponseMetadata ListQualityResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListQualityResResult          `json:"Result,omitempty"`
}

type ListQualityResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListQualityResResponseMetadataError `json:"Error,omitempty"`
}

// ListQualityResResponseMetadataError - 仅在请求失败时返回。
type ListQualityResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListQualityResResult struct {

	// 返回数据的粒度
	AggregateGranularity *string `json:"AggregateGranularity,omitempty"`

	// 具体指标数据
	Indicators []*ListQualityResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListQualityResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListQualityResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                     `json:"Name,omitempty"`
	Overview *ListQualityResResultIndicatorsItemOverview `json:"overview,omitempty"`

	// 指标单位
	Unit *string `json:"Unit,omitempty"`
}

type ListQualityResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListQualityResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRealTimeOperationDataBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。 Indicator 支持取值参看实时运营数据相关 indicator [1167931#realtime]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`
}

type ListRealTimeOperationDataRes struct {

	// REQUIRED
	ResponseMetadata ListRealTimeOperationDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRealTimeOperationDataResResult          `json:"Result,omitempty"`
}

type ListRealTimeOperationDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRealTimeOperationDataResResponseMetadataError `json:"Error,omitempty"`
}

// ListRealTimeOperationDataResResponseMetadataError - 仅在请求失败时返回。
type ListRealTimeOperationDataResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRealTimeOperationDataResResult struct {

	// 具体指标数据
	Indicators []*ListRealTimeOperationDataResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListRealTimeOperationDataResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListRealTimeOperationDataResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                                   `json:"Name,omitempty"`
	Overview *ListRealTimeOperationDataResResultIndicatorsItemOverview `json:"overview,omitempty"`

	// 指标单位
	Unit *string `json:"Unit,omitempty"`
}

type ListRealTimeOperationDataResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListRealTimeOperationDataResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRealTimeQualityBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。可同时查询多个指标。Indicator 可选值，参看质量数据相关 indicator [1167931#listquality]。
	Indicator []string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`

	// 要查询的数据所属设备端，支持取值及含义如下：
	// * native：Android、iOS、Linux、Mac 或 Windows 端。
	// * web：指 Web 端。
	// 默认值为native。
	ProductType *string `json:"ProductType,omitempty"`

	// 房间 ID，是房间的唯一标志。如果不填，代表查询该 AppId 的整体实时指标。
	RoomID *string `json:"RoomId,omitempty"`

	// 查询的用户 ID 列表，最多可以指定 20 个。值不合法时默认剔除。此字段仅在 RoomId 不为空时生效。
	UserID []*string `json:"UserId,omitempty"`
}

type ListRealTimeQualityDistributionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询的维度，一次仅支持查询一个维度。支持取值及含义如下：
	// * Province：一级行政区（包括港澳台地区）
	// * Country：国家
	// * OS：用户设备平台。支持取值：android、ios、linux、mac、windows
	// * Access：用户网络类型。支持取值：2g、3g、4g、5g、wifi
	Dimension string `json:"Dimension"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的指标名称。一次仅支持查询一个指标。Indicator 支持取值参看质量数据相关 indicator [1167931#listquality]。
	Indicator string `json:"Indicator"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒。
	StartTime string `json:"StartTime"`

	// 设备平台，目前支持 Android、iOS 和 Web。
	Platform *string `json:"Platform,omitempty"`

	// 要查询的产品类型，支持取值及含义如下：
	// * native：Android、iOS、linux、Nac或 Windows 平台的 veRTC SDK。
	// * web：Web 平台的veRTC SDK。
	// 默认值为 native。
	ProductType *string `json:"ProductType,omitempty"`
}

type ListRealTimeQualityDistributionRes struct {

	// REQUIRED
	ResponseMetadata ListRealTimeQualityDistributionResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRealTimeQualityDistributionResResult          `json:"Result,omitempty"`
}

type ListRealTimeQualityDistributionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRealTimeQualityDistributionResResponseMetadataError `json:"Error,omitempty"`
}

// ListRealTimeQualityDistributionResResponseMetadataError - 仅在请求失败时返回。
type ListRealTimeQualityDistributionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRealTimeQualityDistributionResResult struct {

	// 具体指标数据
	Indicators []*ListRealTimeQualityDistributionResResultIndicatorsItem `json:"Indicators,omitempty"`
}

type ListRealTimeQualityDistributionResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListRealTimeQualityDistributionResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name *string `json:"Name,omitempty"`
}

type ListRealTimeQualityDistributionResResultIndicatorsPropertiesItemsItem struct {

	// 具体的维度信息。
	Dimension *string `json:"Dimension,omitempty"`

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标数值，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRealTimeQualityRes struct {

	// REQUIRED
	ResponseMetadata ListRealTimeQualityResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRealTimeQualityResResult          `json:"Result,omitempty"`
}

type ListRealTimeQualityResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRealTimeQualityResResponseMetadataError `json:"Error,omitempty"`
}

// ListRealTimeQualityResResponseMetadataError - 仅在请求失败时返回。
type ListRealTimeQualityResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRealTimeQualityResResult struct {

	// 具体指标数据
	Indicators []*ListRealTimeQualityResResultIndicatorsItem `json:"Indicators,omitempty"`

	// 房间 ID
	RoomID *string `json:"RoomId,omitempty"`

	// 用户 ID
	UserID []*string `json:"UserId,omitempty"`
}

type ListRealTimeQualityResResultIndicatorsItem struct {

	// 具体的指标值以及对应时间
	Data []*ListRealTimeQualityResResultIndicatorsPropertiesItemsItem `json:"Data,omitempty"`

	// 指标名称
	Name     *string                                             `json:"Name,omitempty"`
	Overview *ListRealTimeQualityResResultIndicatorsItemOverview `json:"overview,omitempty"`

	// 指标单位
	Unit *string `json:"Unit,omitempty"`
}

type ListRealTimeQualityResResultIndicatorsItemOverview struct {

	// 指标中文名
	Alias     *string `json:"Alias,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty"`

	// hover时提醒信息
	HoverInfo *string `json:"HoverInfo,omitempty"`
	MaxValue  *string `json:"MaxValue,omitempty"`
	MinValue  *string `json:"MinValue,omitempty"`

	// 指标名
	Name *string `json:"Name,omitempty"`

	// 单位
	Unit *string `json:"Unit,omitempty"`

	// 值
	Value *string `json:"Value,omitempty"`
}

type ListRealTimeQualityResResultIndicatorsPropertiesItemsItem struct {

	// 指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。指标聚合时间，即指标值对应时间段的开始时刻（每 30s 一聚合）。格式为 RFC3339 规范。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 指标值，浮点数，保留两位小数。指标值，浮点数，保留两位小数。
	Value *float64 `json:"Value,omitempty"`
}

type ListRelayStreamQuery struct {

	// REQUIRED; 应用的唯一标志。
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 要查询房间的 ID
	RoomID string `json:"RoomId" query:"RoomId"`

	// 页大小，取值范围为[1,100]，默认值为 10。
	Limit *int32 `json:"Limit,omitempty" query:"Limit"`

	// 起始位置，取值范围为[0,9999]，默认值为 0。
	Offset *int32 `json:"Offset,omitempty" query:"Offset"`

	// 接口调用时间顺序。
	// * 0：倒序。
	// * 1：正序。
	// 默认值为0。
	Order *int32 `json:"Order,omitempty" query:"Order"`
}

type ListRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata ListRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRelayStreamResResult          `json:"Result,omitempty"`
}

type ListRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// ListRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type ListRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListRelayStreamResResult struct {

	// REQUIRED; 页大小
	Limit int32 `json:"Limit"`

	// REQUIRED; 下一页索引
	Offset int32 `json:"Offset"`

	// REQUIRED; 任务列表
	Task []ListRelayStreamResResultTaskItem `json:"Task"`

	// REQUIRED; 当前页大小
	Total int32 `json:"Total"`
}

// ListRelayStreamResResultTaskItem - 任务列表
type ListRelayStreamResResultTaskItem struct {

	// 你的音视频应用的唯一标志。
	AppID *string `json:"AppId,omitempty"`

	// 任务开始时间，Unix 时间戳，单位为秒
	CreateTime *int32 `json:"CreateTime,omitempty"`

	// 发送帧率，取值范围为[1，60]，默认值为 15，转码时生效。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 媒体类型。0：音视频 1：音频2：视频
	MediaType *int32 `json:"MediaType,omitempty"`

	// 房间 ID，是房间的唯一标志
	RoomID *string `json:"RoomId,omitempty"`

	// 任务当前状态。 1：任务启动中。2：任务运行中。3：任务已结束。
	Status *int32 `json:"Status,omitempty"`

	// 流处理模式。0：转码1：转封装
	StreamMode *int32 `json:"StreamMode,omitempty"`

	// 在线媒体流地址
	StreamURL *string `json:"StreamUrl,omitempty"`

	// 任务 ID
	TaskID *string `json:"TaskId,omitempty"`

	// 在线媒体流对应的 UserId
	UserID *string `json:"UserId,omitempty"`

	// 视频高度，取值范围为 [16, 1920]，单位为像素。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 视频宽度，取值范围为 [16, 1920]，单位为像素。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type ListRoomInfoQuery struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒
	StartTime string `json:"StartTime" query:"StartTime"`

	// 分页序号，默认值为1
	PageNum *int64 `json:"PageNum,omitempty" query:"PageNum"`

	// 每页房间数，最大不能超过 100。默认为 20。如果指定值超过 100，每页的房间数为 100。
	PageSize *int64 `json:"PageSize,omitempty" query:"PageSize"`

	// 房间 ID，是房间的唯一标志
	RoomID *string `json:"RoomId,omitempty" query:"RoomId"`
}

type ListRoomInfoRes struct {
	ResponseMetadata *ListRoomInfoResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *ListRoomInfoResResult           `json:"Result,omitempty"`
}

type ListRoomInfoResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *ListRoomInfoResResponseMetadataError `json:"Error,omitempty"`
}

type ListRoomInfoResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message     string  `json:"Message"`
	CodeN       *int32  `json:"CodeN,omitempty"`
	Description *string `json:"Description,omitempty"`
}

type ListRoomInfoResResult struct {

	// 是否还有房间未列出。若为true，表示符合条件房间过多。若要查询房间未列出，请修改查询条件。
	HasMore *bool `json:"HasMore,omitempty"`

	// 分页序号
	PageNum *int64 `json:"PageNum,omitempty"`

	// 每页房间数
	PageSize *int64 `json:"PageSize,omitempty"`

	// 房间信息
	RoomList []*ListRoomInfoResResultRoomListItem `json:"RoomList,omitempty"`

	// 查询到的房间总数。若同一 RoomId 下有多次通话，记为多个房间。
	Total *int64 `json:"Total,omitempty"`
}

type ListRoomInfoResResultRoomListItem struct {

	// 前仍在房间用户数
	ActiveUserNum *int64 `json:"ActiveUserNum,omitempty"`

	// 通话 Id，是通话的唯一标识
	CallID *string `json:"CallId,omitempty"`

	// 通话创建时间，格式为 RFC3339，单位秒
	CreatedTime *string `json:"CreatedTime,omitempty"`

	// 通话结束时间，格式为 RFC3339，单位秒，若查询时还未结束，则返回空值。
	DestroyTime *string `json:"DestroyTime,omitempty"`

	// 通话是否结束
	IsFinished *bool `json:"IsFinished,omitempty"`

	// 返回房间 Id
	RoomID *string `json:"RoomId,omitempty"`
}

type ListUsagesBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 查询数据的结束时间，格式要求符合 RFC3339 规范。起止时间的跨度最多为 30 天
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询数据的开始时间，格式要求符合 RFC3339 规范
	StartTime string `json:"StartTime"`
}

type ListUsagesRes struct {

	// REQUIRED
	ResponseMetadata ListUsagesResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListUsagesResResult          `json:"Result,omitempty"`
}

type ListUsagesResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ListUsagesResResponseMetadataError `json:"Error,omitempty"`
}

// ListUsagesResResponseMetadataError - 仅在请求失败时返回。
type ListUsagesResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type ListUsagesResResult struct {

	// 具体的通话时长数据
	Usages []*ListUsagesResResultUsagesItem `json:"Usages,omitempty"`
}

type ListUsagesResResultUsagesItem struct {

	// 计价档位为纯音频的通话时长，单位为分钟
	Audio *int64 `json:"Audio,omitempty"`

	// 指标聚合的时间，格式为 RFC3339 规范。这个参数的值，指通话时长对应时间段的开始时刻。以示例值为例，意味着对应的时间段为 北京时间 2021 年 7 月 24 日 0 点起的 1 天。
	TimeStamp *string `json:"TimeStamp,omitempty"`

	// 计价档位为 1080P 视频的通话时长，单位为分钟
	VideoHD *int64 `json:"VideoHD,omitempty"`

	// 计价档位为 720P 视频的通话时长，单位为分钟
	VideoSD *int64 `json:"VideoSD,omitempty"`

	// 计价档位为 360P 视频的通话时长，单位为分钟
	VideoSDM *int64 `json:"VideoSDM,omitempty"`
}

type ListUserInfoQuery struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 查询结束时间戳，格式为 RFC3339，单位为秒
	EndTime string `json:"EndTime" query:"EndTime"`

	// REQUIRED; 查询起始时间戳，格式为 RFC3339，单位为秒
	StartTime string `json:"StartTime" query:"StartTime"`

	// 是否排除 Linux 用户。默认为 true，表示排除 Linux 用户
	ExcludeServerUser *bool `json:"ExcludeServerUser,omitempty" query:"ExcludeServerUser"`

	// 分页序号，默认值为1
	PageNum *int64 `json:"PageNum,omitempty" query:"PageNum"`

	// 每页用户数，最大不能超过 100。默认为 20。如果指定值超过 100，每页的用户数为 100。
	PageSize *int64 `json:"PageSize,omitempty" query:"PageSize"`

	// 房间 ID，是房间的唯一标志
	RoomID *string `json:"RoomId,omitempty" query:"RoomId"`

	// 查询的用户 ID 列表，多个用户 ID 用逗号隔开，最多可以指定 10 个。值不合法时默认剔除。为空时，查询房间内全部用户信息。
	UserID *string `json:"UserId,omitempty" query:"UserId"`
}

type ListUserInfoRes struct {
	ResponseMetadata *ListUserInfoResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *ListUserInfoResResult           `json:"Result,omitempty"`
}

type ListUserInfoResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *ListUserInfoResResponseMetadataError `json:"Error,omitempty"`
}

type ListUserInfoResResponseMetadataError struct {

	// REQUIRED
	Code string `json:"Code"`

	// REQUIRED
	Message     string  `json:"Message"`
	CodeN       *int32  `json:"CodeN,omitempty"`
	Description *string `json:"Description,omitempty"`
}

type ListUserInfoResResult struct {

	// 分页序号
	PageNum *int64 `json:"PageNum,omitempty"`

	// 每页用户数，若同一用户有多次通话记为多个用户
	PageSize *int64 `json:"PageSize,omitempty"`

	// 查询到的用户总数，若同一用户有多次通话记为多个用户
	Total *int64 `json:"Total,omitempty"`

	// 用户信息
	UserList []*ListUserInfoResResultUserListItem `json:"UserList,omitempty"`
}

type ListUserInfoResResultUserListItem struct {

	// CallId下用户最后一次进房对应的网络类型。枚举值：2g、3g、4g、5g、wifi、unknown。
	Access *string `json:"Access,omitempty"`

	// 通话 Id，是通话的唯一标识
	CallID *string `json:"CallId,omitempty"`

	// CallId下用户第一次进入通话时间，格式为 RFC3339，单位秒
	CreatedAt *string `json:"CreatedAt,omitempty"`

	// CallId下用户最后一次进房对应的设备型号
	DeviceType *string `json:"DeviceType,omitempty"`

	// CallId下用户从第一次进房到最后一次离开房间的时间范围内真实在线时长（多次进出房时间间隔累加），单位为秒
	Duration *int64 `json:"Duration,omitempty"`

	// CallId下用户是否离开房间
	Finished *bool `json:"Finished,omitempty"`

	// CallId下用户最后一次退出通话时间，格式为 RFC3339，单位秒。如果此时用户在线，返回为空。
	LeaveAt *string `json:"LeaveAt,omitempty"`

	// CallId下用户最后一次进房对应的设备平台。枚举值：android、ios、linux、mac、windows、web、unknown
	OS *string `json:"OS,omitempty"`

	// CallId下用户是否发布过流
	Pub *bool `json:"Pub,omitempty"`

	// CallId下用户在通话内全部进退房记录
	Record []*ListUserInfoResResultUserListPropertiesItemsItem `json:"Record,omitempty"`

	// 查询房间 Id
	RoomID *string `json:"RoomId,omitempty"`

	// CallId下用户最后一次进房使用 sdk 版本号
	SdkVersion *string `json:"SdkVersion,omitempty"`

	// 查询用户 Id
	UserID *string `json:"UserId,omitempty"`
}

type ListUserInfoResResultUserListPropertiesItemsItem struct {

	// 用户进房/退房时间，格式为 RFC3339，单位秒。
	Time *string `json:"Time,omitempty"`

	// 用户进房/退房类型，取值为 joinroom或 leaveroom。
	Type *string `json:"Type,omitempty"`
}

type ModifyAppStatusBody struct {

	// REQUIRED; 需要停用/启用的 AppId
	AppID string `json:"AppId"`

	// REQUIRED; * 1：将 App 状态设置为启用；
	// * 2：将 App 状态设置为停用；
	Status int32 `json:"Status"`
}

type ModifyAppStatusRes struct {

	// REQUIRED
	ResponseMetadata ModifyAppStatusResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 返回值Result仅在请求成功时返回 "success",失败时为空。
	Result string `json:"Result"`
}

type ModifyAppStatusResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *ModifyAppStatusResResponseMetadataError `json:"Error,omitempty"`
}

// ModifyAppStatusResResponseMetadataError - 仅在请求失败时返回。
type ModifyAppStatusResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartDetectionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// 回调地址 开启审核后，如果可能存在违规信息，此地址会收到违规信息回调。如果地址无效或为空，审核会继续，但你不会收到违规信息的回调结果。
	Callback *string `json:"Callback,omitempty"`

	// 回调种类。支持取值及含义如下：
	// * 0：违规回调
	// * 1：全部回调。
	// 默认值为 0。
	CallbackType *int32 `json:"CallbackType,omitempty"`

	// 每段音频切片的时长，取值范围为 [1000，600000]，单位为毫秒，默认值为 20000。该参数不建议设置过大，如果设置过大, 会出现审核延迟的问题，且造成最后一段切片计费误差向上取整偏大。
	Duration *int32 `json:"Duration,omitempty"`

	// 任务最大空闲超时时间。如果指定的用户停止推流或素材间隔过长，导致素材接收不到，那么审核任务会在空闲时间超过设定值后自动停止。取值范围为 [1，10800]，单位为秒，默认值为 180。
	IdleSec *int32 `json:"IdleSec,omitempty"`

	// 相邻截图之间的间隔时间，取值范围为 [100，600000]，单位为毫秒，默认值为 2000。
	Interval *int32 `json:"Interval,omitempty"`

	// 进行审核的内容类型。支持取值及含义如下：
	// * 1：视频截图；
	// * 2：音频切片；
	// * 3：视频截图+音频切片
	// 默认值为 3。 视频截图：审核过程中，RTC 会按照设定的时间间隔，进行周期性截图，并对截图进行审核。若出现违规信息，会返回审核结果。
	// > 音频切片：审核过程中，RTC 会按设定的音频切片时长，保存每段音频切片，并对切片进行审核。若出现违规信息，会返回审核结果。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 要审核的用户ID。若不填，则审核房间内所有推流用户。最多可以审核 17 路流。
	// * 如果先发起单流审核 再发起房间级审核，会发起房间级审核并停止单流审核；
	// * 如果先发起房间级审核，再发起单流审核，会引发错误，提示：已有审核进行中；
	// * 如果先发起单流音频审核和单流视频审核，再发起房间音频审核，会合并单流音频审核到房间音频审核，单流视频审核无变化；
	// * 如果先发起单流音频审核和单流视频审核，再发起房间音频审核，之后又发起房间音视频审核，会合并所有审核到房间音视频审核。
	UserID *string `json:"UserId,omitempty"`
}

type StartDetectionRes struct {

	// REQUIRED
	ResponseMetadata StartDetectionResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result StartDetectionResResult `json:"Result"`
}

type StartDetectionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartDetectionResResponseMetadataError `json:"Error,omitempty"`
}

// StartDetectionResResponseMetadataError - 仅在请求失败时返回。
type StartDetectionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartDetectionResResult struct {

	// REQUIRED; 仅在请求成功时返回"success",失败时为空
	Message string `json:"Message"`
}

type StartPushMixedStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 推流 CDN 地址。仅支持 RTMP 协议。
	PushURL string `json:"PushURL"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 合流转推任务 ID。你必须对每个合流转推任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	// 若任务运行中，使用相同的 TaskId 重复调用开始接口不会导致请求失败，BaseResponse.Result 会提示 The task has been started. Please do not call the startup
	// task interface repeatedly。
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string                               `json:"BusinessId,omitempty"`
	Control    *StartPushMixedStreamToCDNBodyControl `json:"Control,omitempty"`

	// 音视频编码参数。
	// * 单流录制时，你仅可以设置VideoFps、VideoBitrate和所有音频相关参数。
	// * 合流录制时，你仅可以设置VideoWidth、VideoHeight、VideoFps、VideoBitrate和所有音频相关参数。
	Encode         *StartPushMixedStreamToCDNBodyEncode         `json:"Encode,omitempty"`
	ExcludeStreams *StartPushMixedStreamToCDNBodyExcludeStreams `json:"ExcludeStreams,omitempty"`
	Layout         *StartPushMixedStreamToCDNBodyLayout         `json:"Layout,omitempty"`
	TargetStreams  *StartPushMixedStreamToCDNBodyTargetStreams  `json:"TargetStreams,omitempty"`
}

type StartPushMixedStreamToCDNBodyControl struct {

	// 选择补帧模式。支持取值及含义如下：
	// * 0：补最后一帧。
	// * 1：补黑帧
	// 默认值为0。 自动布局模式下，没有补帧的逻辑。 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系: 你可以在 Region 中传入
	// Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为 [10, 86400],单位为秒，默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。支持取值及含义如下：
	// * 0 ：音视频；
	// * 1 ：纯音频。
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 转推直播推流模式，用于控制触发推流的时机。支持取值及含义如下：
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。
	// 默认值为 0。 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                             `json:"PushStreamMode,omitempty"`
	SEIConfig      *StartPushMixedStreamToCDNBodyControlSEIConfig     `json:"SEIConfig,omitempty"`
	SpatialConfig  *StartPushMixedStreamToCDNBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`
}

type StartPushMixedStreamToCDNBodyControlSEIConfig struct {

	// SEI 中是否包含用户说话音量值。
	// * false：否。
	// * true：是。
	// 默认值为 false。
	AddVolumeValue *bool `json:"AddVolumeValue,omitempty"`

	// 是否透传 RTC 流里的 SEI 信息。
	// * true：是。
	// * false：否。
	// 默认值为 true。
	PassthroughRTCSEIMessage *bool `json:"PassthroughRTCSEIMessage,omitempty"`

	// 开启音量指示模式后，非关键帧携带 SEI 包含的信息类型。支持取值及含义如下：
	// * 0：全量信息。
	// * 1：只有音量信息和时间戳。
	// 默认值为 0。关于 SEI 结构，参看服务端合流转推 SEI 结构 [1163740#sei]
	SEIContentMode *int32 `json:"SEIContentMode,omitempty"`

	// 设置 SEI 的 Payload Type，对 服务端合流转推 SEI 和 RTC 透传SEI 均生效。支持取值为 5 或 100，默认为 100。
	SEIPayLoadType *int32 `json:"SEIPayLoadType,omitempty"`

	// 该字段为长度为 32 位的 16 进制字符，每个字符的范围为 a-f，A-F，0-9，不可包含 -。如果校验失败，会返回错误码 InvalidParameter。
	// 仅当 SEIPayLoadType=5时，该字段需要填写，SEIPayLoadType=100时，该字段内容会被忽略。
	// 此参数仅对 RTC 透传SEI 生效。当用户设置 SEIPayloadType = 5 时，服务端合流转推SEI 的 SEIPayloadUUID 为固定值：566f6c63656e67696e65525443534549，对应16位字符串
	// VolcengineRTCSEI。
	SEIPayloadUUID *string `json:"SEIPayloadUUID,omitempty"`

	// 设置SEI数据结构 [1163740#sei]中 app_data 参数的值，最大长度为 4096。此参数支持动态更新。
	UserConfigExtraInfo *string `json:"UserConfigExtraInfo,omitempty"`

	// 用户说话音量的回调间隔。取值范围为 [0.3,∞]，单位为秒，默认值为 2。仅当用户说话音量发生变化时此回调才会被触发。
	VolumeIndicationInterval *float32 `json:"VolumeIndicationInterval,omitempty"`

	// 是否开启音量指示模式。
	// * true：是。此时非关键帧中也可能携带 SEI 信息。
	// * false：否。此时只会在关键帧中携带 SEI 信息。
	// 默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。 关于音量指示模式的用法，参看 音量指示模式 [1163740#volume]。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type StartPushMixedStreamToCDNBodyControlSpatialConfig struct {
	AudienceSpatialOrientation *StartPushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。
	// * false：关闭。
	// * true：开启
	// 默认值为 false。
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

type StartPushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 前方朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 头顶朝向
	Up []*float32 `json:"Up,omitempty"`
}

// StartPushMixedStreamToCDNBodyEncode - 音视频编码参数。
// * 单流录制时，你仅可以设置VideoFps、VideoBitrate和所有音频相关参数。
// * 合流录制时，你仅可以设置VideoWidth、VideoHeight、VideoFps、VideoBitrate和所有音频相关参数。
type StartPushMixedStreamToCDNBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps。
	// * 该参数若不填，RTC 会根据音频配置文件类型、采样率、声道数的配置生成音频码率取值。
	// * 该参数若填写，实际生效的码率也会随着音频配置文件类型、采样率、声道数的配置而变化。若你发现生效的码率值没有达到你设定的值，可能是该值已经超过该场景下的极限码率。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。支持取值及含义如下：
	// * 1：单声道。
	// * 2：双声道。
	// 默认值为 2。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。仅支持取 0，表示 AAC 编码协议。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型。支持取值及含义如下：
	// * 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。此时，只支持输出流音频声道数为双声道。
	// 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。可取值为：{32000,44100，48000}，单位为 Hz，默认值为 48000。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 视频码率。取值范围为 [0,10000]，单位为 Kbps，默认值为 0。0 表示自适应码率，会自动根据 VideoFps , VideoWidth 以及VideoHeight 计算出合理的码率。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。支持取值及含义如下：
	// * 0：H.264
	// * 1：ByteVC1
	// 默认值为 0。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 视频帧率。取值范围为 [1,60]，单位为 FPS，默认值为 15。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。取值范围为 [1,5]，单位为秒，默认值为 4。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 画面高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Height 设置画面宽度。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 画面宽度。取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Width 设置画面宽度。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type StartPushMixedStreamToCDNBodyExcludeStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartPushMixedStreamToCDNBodyExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartPushMixedStreamToCDNBodyExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayout struct {
	CustomLayout *StartPushMixedStreamToCDNBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。支持取值及含义如下：
	// * 0：自适应布局 [1167930#adapt]。
	// * 1：垂直布局 [1167930#vertical]
	// * 2 ：自定义布局。
	// * 3 ：并排布局 [1167930#horizontal]
	// 默认值为0
	LayoutMode      *int32                                              `json:"LayoutMode,omitempty"`
	MainVideoStream *StartPushMixedStreamToCDNBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayout struct {
	Canvas *StartPushMixedStreamToCDNBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。
	Regions []*StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色, 范围为 #000000 ~ #ffffff (大小写均可)，格式为 #RGB(16进制)，默认值为 #000000（黑色）。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。 如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。
	Width *int32 `json:"Width,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]，默认值为 1.0。0.0 表示完全透明，1.0 表示完全不透明。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 [0,1]。默认值为 0，表示没有圆角效果。 圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。 例如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 ：按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 ：按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形。
	// 默认值为 0。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]，默认值为 0。0 表示该区域图像位于最下层，100 表示该区域图像位于最上层。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartPushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartPushMixedStreamToCDNBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNBodyTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartPushMixedStreamToCDNBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartPushMixedStreamToCDNBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushMixedStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StartPushMixedStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                      `json:"Result,omitempty"`
}

type StartPushMixedStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartPushMixedStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushMixedStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StartPushMixedStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartPushPublicStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 公共流 ID。你必须对每路公共流，设定 PublicStreamId 时命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	PublicStreamID string `json:"PublicStreamId"`

	// REQUIRED; 为公共流指定单路或多路媒体流及对应参数，Stream 数组。最高支持 17 路。
	TargetStreams []StartPushPublicStreamBodyTargetStreamsItem `json:"TargetStreams"`

	// 你创建的业务标识
	BusinessID *string                           `json:"BusinessId,omitempty"`
	Control    *StartPushPublicStreamBodyControl `json:"Control,omitempty"`
	Encode     *StartPushPublicStreamBodyEncode  `json:"Encode,omitempty"`

	// 你可以通过本参数排除掉不需要包含在公共流中的音视频流，即黑名单。参数默认为空。黑名单中的流不得超过 17 路。此参数中的 stream 不应和 TargetStreams 中重复。
	ExcludeStreams []*StartPushPublicStreamBodyExcludeStreamsItem `json:"ExcludeStreams,omitempty"`
	Layout         *StartPushPublicStreamBodyLayout               `json:"Layout,omitempty"`

	// 公共流处理模式。支持取值及含义如下：0：转码。1：转封装。
	// 当 TranscodeMode=1 时，
	// * TargetStreams 只能指定一路流，且该路流的 UserId不能为空，需为对应房间用户的 UserId。
	// * ExcludeStreams 必须为空。
	// * Encode.VideoConfig 设置不生效。
	// * Layout 设置不生效。
	TranscodeMode *int32 `json:"TranscodeMode,omitempty"`
}

type StartPushPublicStreamBodyControl struct {

	// 插入公共流的自定义信息，可用于随流信息同步，长度不超过 4 kB。 数据会添加到当前视频帧开始的连续 30 个视频帧中。 只在调用UpdatePublicStreamParam时有效。
	DataMsg *string `json:"DataMsg,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为[10, 86400]，单位为秒，默认值为180。只在调用StartPushPublicStream时有效。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 房间用户发布状态回调间隔，仅在纯音频时生效。取值范围为 [1000,2147483647]，单位为毫秒，默认值为 2000。
	StreamPublishStatsInterval *int32 `json:"StreamPublishStatsInterval,omitempty"`

	// 是否开启房间用户发布状态回调。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户发布状态回调。
	// * false：不开启房间用户发布状态回调。
	// 默认值为false。
	StreamPublishStatsMode *bool `json:"StreamPublishStatsMode,omitempty"`

	// 房间用户采集状态回调间隔，仅在纯音频时生效。取值范围为[1000,2147483647]，单位为毫秒，默认值为2000。
	UserCaptureStatsInterval *int32 `json:"UserCaptureStatsInterval,omitempty"`

	// 是否开启房间用户采集状态回调。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户采集状态回调。
	// * false：不开启房间用户采集状态回调。
	// 默认值为false。
	UserCaptureStatsMode *bool `json:"UserCaptureStatsMode,omitempty"`

	// 音量指示的回调间隔。最小值为100，单位为毫秒，默认值为2000。
	// VideoConfig.FrameRate大于 10 fps 时，回调间隔才能达到 100ms。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// 是否开启音量指示模式。
	// * true：开启音量提示。
	// * false：不开启音量提示。
	// 默认值为false。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type StartPushPublicStreamBodyEncode struct {

	// 视频编码配置
	VideoConfig *StartPushPublicStreamBodyEncodeVideoConfig `json:"VideoConfig,omitempty"`
}

// StartPushPublicStreamBodyEncodeVideoConfig - 视频编码配置
type StartPushPublicStreamBodyEncodeVideoConfig struct {

	// 最高输出视频码率。取值范围为[0,10000]，单位为 Kbps，默认值为0，0表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 输出视频帧率。取值范围为 [1,60]，单位为 fps，默认为15。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 输出画面的高度，默认值为 480。范围为 [16, 1920]，必须是偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频编码协议。支持取值及含义如下：
	// * 0：H.264。
	// * 5：VP8。 如果选择 VP8 格式，请先联系火山技术支持配置。
	// 默认值为0。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出画面的宽度。默认值为 640，范围为 [16, 1920]，必须是偶数。
	Width *int32 `json:"Width,omitempty"`
}

type StartPushPublicStreamBodyExcludeStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	//
	//
	// * 1：纯音频
	//
	//
	// * 2：纯视频
	//
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。支持取值及含义如下：
	// * 0：音视频流
	//
	//
	// * 1：屏幕流
	//
	// 默认值为 0。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。UserId 为空时，表示订阅房间内所有流。UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type StartPushPublicStreamBodyLayout struct {

	// REQUIRED; 布局模式。支持取值及含义如下：
	// * 0：自适应布局 [1167930#adapt]
	// * 1：：垂直布局 [1167930#vertical]
	// * 2 ：自定义布局
	// * 3 ：并排布局 [1167930#horizontal]
	// 默认值为0
	LayoutMode     int32                                          `json:"LayoutMode"`
	CustomLayout   *StartPushPublicStreamBodyLayoutCustomLayout   `json:"CustomLayout,omitempty"`
	VerticalLayout *StartPushPublicStreamBodyLayoutVerticalLayout `json:"VerticalLayout,omitempty"`
}

type StartPushPublicStreamBodyLayoutCustomLayout struct {

	// REQUIRED; 自定义布局下，多路视频配置
	Regions []StartPushPublicStreamBodyLayoutCustomLayoutRegionsItem `json:"Regions"`

	// 背整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundColor *string `json:"BackgroundColor,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 补帧模式。支持取值及含义如下：
	// * 0：补最后一帧，
	// * 1：补黑帧。
	// 默认值为0。自动布局模式下，该参数不生效。 补帧是指在音视频录制时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 如果同时配置Alternateimage 和FrameInterpolationMode ，优先使用 Alternateimage参数。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`
}

type StartPushPublicStreamBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流标识。
	// StreamIndex 即 Stream.Index，用来指定布局设置将被应用到哪路流。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]，默认值为 1.0。0.0 表示完全透明，1.0 表示完全不透明。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 占位图片的 url
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的纵坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 ：按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 ：按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形。
	// 默认值为 0。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartPushPublicStreamBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 当画面有重叠时，使用此参数设置指定画面的图层顺序，取值范围为 [0, 100]，默认值为 0。0 表示该区域图像位于最下层，100 表示该区域图像位于最上层。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartPushPublicStreamBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartPushPublicStreamBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartPushPublicStreamBodyLayoutVerticalLayout struct {
	MainStream *StartPushPublicStreamBodyLayoutVerticalLayoutMainStream `json:"MainStream,omitempty"`
}

type StartPushPublicStreamBodyLayoutVerticalLayoutMainStream struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	//
	//
	// * 1：纯音频
	//
	//
	// * 2：纯视频
	//
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。支持取值及含义如下：
	// * 0：音视频流
	//
	//
	// * 1：屏幕流
	//
	// 默认值为 0。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。UserId 为空时，表示订阅房间内所有流。UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type StartPushPublicStreamBodyTargetStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	//
	//
	// * 1：纯音频
	//
	//
	// * 2：纯视频
	//
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。支持取值及含义如下：
	// * 0：音视频流
	//
	//
	// * 1：屏幕流
	//
	// 默认值为 0。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。UserId 为空时，表示订阅房间内所有流。UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type StartPushPublicStreamRes struct {

	// REQUIRED
	ResponseMetadata StartPushPublicStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                  `json:"Result,omitempty"`
}

type StartPushPublicStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartPushPublicStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushPublicStreamResResponseMetadataError - 仅在请求失败时返回。
type StartPushPublicStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartPushSingleStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED
	Control StartPushSingleStreamToCDNBodyControl `json:"Control"`

	// REQUIRED; 推流地址。目前仅支持 rtmp 协议
	PushURL string `json:"PushURL"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED
	Stream StartPushSingleStreamToCDNBodyStream `json:"Stream"`

	// REQUIRED; 转推任务 ID。你必须对每个转推任务设定 TaskId，且在后续进行任务结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	// 若任务运行中，使用相同的 TaskId 重复调用开始接口不会导致请求失败，BaseResponse.Result 会提示 The task has been started. Please do not call the startup
	// task interface repeatedly。
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StartPushSingleStreamToCDNBodyControl struct {

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为 [10, 86400], 单位为秒，默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。默认值为 0，可以取0和1。0表示音视频，1表示纯音频，暂不支持纯视频。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type StartPushSingleStreamToCDNBodyStream struct {

	// REQUIRED; 用户Id，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartPushSingleStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StartPushSingleStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                       `json:"Result,omitempty"`
}

type StartPushSingleStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartPushSingleStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StartPushSingleStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StartPushSingleStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制文件的存储平台配置。
	// 支持：
	// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * Amazon S3
	// * 阿里云对象存储 OSS
	// * 华为云 OBS
	// * 腾讯云 COS
	// * 七牛云 Kodo。
	StorageConfig StartRecordBodyStorageConfig `json:"StorageConfig"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	// 若任务运行中，使用相同的 TaskId 重复调用开始接口不会导致请求失败，BaseResponse.Result 会提示 The task has been started. Please do not call the startup
	// task interface repeatedly。
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 高级配置选项
	Control          *StartRecordBodyControl          `json:"Control,omitempty"`
	Encode           *StartRecordBodyEncode           `json:"Encode,omitempty"`
	ExcludeStreams   *StartRecordBodyExcludeStreams   `json:"ExcludeStreams,omitempty"`
	FileFormatConfig *StartRecordBodyFileFormatConfig `json:"FileFormatConfig,omitempty"`

	// 录制文件的命名设置。
	// 录制文件的名称由 [StorageConfig.Type] 和 [FileNameConfig] 共同决定。
	// * 当StorageConfig.Type配置为 0 ，即存储在 TOS 平台时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern +随机数组成
	// * 当StorageConfig.Type配置为 1 ，即存储在 火山引擎视频点播平台 [https://www.volcengine.com/product/vod] 平台时，录制文件名称由FilenameConfig.Pattern
	// + 随机数组成
	// * 当StorageConfig.Type配置为 2 ，即存储在指定第三方S3 对象存储平台 [155125#storage]时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern
	// +随机数组成。 例如：当 StorageConfig.Type 配置为 0 ，FilenameConfig.Prefix配置为
	// directory1/directory2/，FilenameConfig.Pattern 配置为 {TaskId}_{RoomId}_{StartTime}_{Duration}， 若此时该文件的 TaskId = mytask123456789,
	// RoomId = myroom99991234，StartTime =1645769481126，Duration = 30s
	// 且生成的随机八位字符为 c4694fa1，则生成录制文件的文件名为：directory1/directory2/mytask123456789_myroom99991234_1645769481126_000030_c4694fa1.mp4。
	// > 文件名在 视频点播平台 [https://www.volcengine.com/product/vod] 上可以重复，但在 TOS [https://www.volcengine.com/product/tos] 或第三方存储平台 [155125#storage]上默认不可以重复。
	FileNameConfig *StartRecordBodyFileNameConfig `json:"FileNameConfig,omitempty"`
	Layout         *StartRecordBodyLayout         `json:"Layout,omitempty"`

	// 录制模式。支持取值及含义如下：
	// * 0 ：合流录制，
	// * 1 ：单流录制。
	// 默认值为 0。
	RecordMode    *int32                        `json:"RecordMode,omitempty"`
	TargetStreams *StartRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

// StartRecordBodyControl - 高级配置选项
type StartRecordBodyControl struct {

	// 是否开启边录制边上传功能。
	// * false：关闭
	// * true：开启
	// 默认值为 false。 :::warning 该功能仅对HLS格式存储文件生效，即：FileFormatConfig.FileFormat取值必须包含 HLS。 若 HLS 格式文件名中包含Duration填充变量，开通该功能Duration的值始终为
	// 0。 :::
	EnableSyncUpload *bool `json:"EnableSyncUpload,omitempty"`

	// 补帧模式。支持取值及含义如下：
	// * 0：补最后一帧，
	// * 1：补黑帧。
	// 默认值为0。自动布局模式下，该参数不生效。 补帧是指在音视频录制时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 如果同时配置Alternateimage 和FrameInterpolationMode ，优先使用 Alternateimage参数。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为 [10, 86400],单位为秒，默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 最大录制时长，取值为正整数，单位为秒。默认值为 0。0 表示不限制录制时长。
	MaxRecordTime *int32 `json:"MaxRecordTime,omitempty"`

	// 任务最大中断时间，仅对单流录制生效。取值范围为 [60, 3600],单位为秒，默认值为 3600。
	// * 若任务中断时间小于该参数值，则根据设置的补帧模式进行补帧。
	// * 若任务中断时间大于该参数值但小于空闲超时时间，任务恢复时会生成一个新文件。 建议该参数取值小于 MaxIdleTime 取值。
	MaxSilenceSeconds *int32 `json:"MaxSilenceSeconds,omitempty"`

	// 订阅流类型。支持取值及含义如下：
	// * 0：音视频；
	// * 1：纯音频。
	// 默认值为0。
	MediaType *int32 `json:"MediaType,omitempty"`
}

type StartRecordBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps，默认值为 64。
	// 当AudioProfile=0时： 若输入参数取值范围为 [32,192]，编码码率等于输入码率。
	// 当AudioProfile=1时：
	// * 若输入参数取值范围为 [32,128]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [128,192]，编码码率固定为128。
	// 当AudioProfile=2时：
	// * 若输入参数取值范围为 [32,64]，编码码率等于输入码率。
	// * 若输入参数取值范围为 [64,192]，编码码率固定为64。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。支持取值及含义如下：
	// * 1：单声道。
	// * 2：双声道。
	// 默认值为 2。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。仅支持取 0，表示 AAC 编码协议。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型。支持取值及含义如下：
	// * 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。此时，只支持输出流音频声道数为双声道。
	// 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。可取值为：32000,44100，48000，单位为 Hz，默认值为 48000。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 视频码率。取值范围为 [0,10000]，单位为 Kbps，默认值为 0。0 表示自适应码率，会自动根据 VideoFps , VideoWidth 以及VideoHeight 计算出合理的码率。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。支持取值及含义如下：
	// * 0：H.264
	// * 1：ByteVC1
	// 默认值为 0。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 视频帧率。取值范围为 [1,60]，单位为 FPS，默认值为 15。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。取值范围为 [1,5]，单位为秒，默认值为 4。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 画面高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Height 设置画面宽度。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 画面宽度。取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Width 设置画面宽度。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type StartRecordBodyExcludeStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartRecordBodyExcludeStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyExcludeStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordBodyFileFormatConfig struct {

	// 存储文件格式。支持取值包括：MP4、HLS、FLV、MP3、 AAC、M4A。可同时选择多个存储文件格式，最终生成所选存储格式对应的多个文件。
	// MP3、AAC和M4A仅在编码纯音频时有效。
	FileFormat []*string `json:"FileFormat,omitempty"`
}

// StartRecordBodyFileNameConfig - 录制文件的命名设置。
// 录制文件的名称由 [StorageConfig.Type] 和 [FileNameConfig] 共同决定。
// * 当StorageConfig.Type配置为 0 ，即存储在 TOS 平台时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern +随机数组成
// * 当StorageConfig.Type配置为 1 ，即存储在 火山引擎视频点播平台 [https://www.volcengine.com/product/vod] 平台时，录制文件名称由FilenameConfig.Pattern
// + 随机数组成
// * 当StorageConfig.Type配置为 2 ，即存储在指定第三方S3 对象存储平台 [155125#storage]时，录制文件名称由FilenameConfig.Prefix + FilenameConfig.Pattern
// +随机数组成。 例如：当 StorageConfig.Type 配置为 0 ，FilenameConfig.Prefix配置为
// directory1/directory2/，FilenameConfig.Pattern 配置为 {TaskId}_{RoomId}_{StartTime}_{Duration}， 若此时该文件的 TaskId = mytask123456789,
// RoomId = myroom99991234，StartTime =1645769481126，Duration = 30s
// 且生成的随机八位字符为 c4694fa1，则生成录制文件的文件名为：directory1/directory2/mytask123456789_myroom99991234_1645769481126_000030_c4694fa1.mp4。
// > 文件名在 视频点播平台 [https://www.volcengine.com/product/vod] 上可以重复，但在 TOS [https://www.volcengine.com/product/tos] 或第三方存储平台 [155125#storage]上默认不可以重复。
type StartRecordBodyFileNameConfig struct {

	// 自定义录制文件名模式，具体参看自定义录制文件名 [106873]。 如果你设置了 Pattern，需自行保证最终文件名的唯一性，否则在 TOS 或第三方存储平台上同名文件将被覆盖; 你也可以给对应 bucket 开启版本控制，允许文件名重复，防止被覆盖的情况发生。
	Pattern *string `json:"Pattern,omitempty"`

	// 指定录制文件名的前缀，对TosConfig和CustomConfig生效。 在 TOS 以及支持 S3 的第三方存储平台上，可以实现指定文件夹的功能。如：Prefix=["directory1","directory2"]，将在录制文件名前加上前缀
	// "directory1/directory2/"。 前缀长度最大值为 128 个字符。仅支持大小写字母、数字。
	Prefix []*string `json:"Prefix,omitempty"`
}

type StartRecordBodyLayout struct {
	CustomLayout *StartRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。支持取值及含义如下：
	// * 0：自适应布局 [1167930#adapt]。
	// * 1：垂直布局 [1167930#vertical]
	// * 2 ：自定义布局。
	// * 3 ：并排布局 [1167930#horizontal]
	// 默认值为0
	LayoutMode      *int32                                `json:"LayoutMode,omitempty"`
	MainVideoStream *StartRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type StartRecordBodyLayoutCustomLayout struct {
	Canvas *StartRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。
	Regions []*StartRecordBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type StartRecordBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色, 范围为 #000000 ~ #ffffff (大小写均可)，格式为 #RGB(16进制)，默认值为 #000000（黑色）。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。 如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。
	Width *int32 `json:"Width,omitempty"`
}

type StartRecordBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]，默认值为 1.0。0.0 表示完全透明，1.0 表示完全不透明。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 [0,1]。默认值为 0，表示没有圆角效果。 圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。 例如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 ：按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 ：按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形。
	// 默认值为 0。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]，默认值为 0。0 表示该区域图像位于最下层，100 表示该区域图像位于最上层。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type StartRecordBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type StartRecordBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

// StartRecordBodyStorageConfig - 录制文件的存储平台配置。
// 支持：
// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
// * Amazon S3
// * 阿里云对象存储 OSS
// * 华为云 OBS
// * 腾讯云 COS
// * 七牛云 Kodo。
type StartRecordBodyStorageConfig struct {
	CustomConfig *StartRecordBodyStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *StartRecordBodyStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型。支持取值及含义如下：
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 支持 S3 协议的第三方存储平台
	// * 3: VeImageX (当前仅支持抽帧截图功能)
	// 默认值为 0。
	Type           *int32                                      `json:"Type,omitempty"`
	VeImageXConfig *StartRecordBodyStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *StartRecordBodyStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type StartRecordBodyStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。支持取值及含义如下：
	// * 0：Amazon S3
	// * 1：阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。
	// 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type StartRecordBodyStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartRecordBodyStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	// * 此账号 ID 为火山引擎主账号 ID。
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。 你也可以通过
	// OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为 0。
	Region *int32 `json:"Region,omitempty"`
}

type StartRecordBodyStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 上传到视频点播平台时, 是否需要根据文件后缀自动设置 FileExtension。关于 FileExtension 的详细说明，参看 FileExtension [https://www.volcengine.com/docs/4/164295#fileextension]。
	// * true：需要；
	// * false：不需要。
	// 默认值为 false。
	AutoSetFileExtension *bool `json:"AutoSetFileExtension,omitempty"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 上传到视频点播平台时, 文件的存储类型。支持取值及含义如下：：
	// * 1：标准存储。
	// * 2：归档存储。
	// * 3：低频存储。
	// 默认值为 1。 关于存储类型的详细说明，参看媒资存储存储类型 [https://www.volcengine.com/docs/4/73629#%E5%AA%92%E8%B5%84%E5%AD%98%E5%82%A8]
	StorageClass *int32 `json:"StorageClass,omitempty"`
}

type StartRecordBodyTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartRecordRes struct {

	// REQUIRED
	ResponseMetadata StartRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
}

type StartRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StartRecordResResponseMetadataError - 仅在请求失败时返回。
type StartRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartRelayStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED
	Control StartRelayStreamBodyControl `json:"Control"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个任务设定 TaskId，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。
	// 若任务运行中，使用相同的 TaskId 重复调用开始接口不会导致请求失败，BaseResponse.Result 会提示 The task has been started. Please do not call the startup
	// task interface repeatedly。
	TaskID string `json:"TaskId"`

	// REQUIRED; 客户端与业务服务端进行通讯时用于身份认证的 token 值
	Token string `json:"Token"`

	// REQUIRED; 在线媒体流对应的的 UserId
	UserID string `json:"UserId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 任务的空闲超时时间。超过此时间后，任务自动终止。取值范围为 [5, 600] ，单位为秒，默认值为 300。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`
}

type StartRelayStreamBodyControl struct {

	// REQUIRED; 在线流媒体地址。媒体格式应为：hls、rtmp、mp4、flv、dash、或 ts。如源流为海外，建议联系技术支持，以保障最佳体验。
	StreamURL string `json:"StreamUrl"`

	// 最大发送码率，单位为 Kbps,不填则不限制，转码时生效。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，值的范围为 [1，30]，默认值为15，转码时生效。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 是否循环播放，仅对源流为点播流时生效。
	Loop *bool `json:"Loop,omitempty"`

	// 媒体类型。
	// * 0：音视频
	// * 1：音频。采用此选项时，必须是 AAC 或 Opus 编码。
	// * 2：视频
	// 默认值为0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 任务起始时间戳，用于定时播放，Unix 时间，单位为秒。默认为 0，表示立即启动。此参数仅对 StartRelayStream接口生效。
	StartTimeStamp *int32 `json:"StartTimeStamp,omitempty"`

	// 流处理模式。
	// * 0：转码。采用此选项时，原视频编码方式必须是 H.264 或 ByteVC1。
	// * 1：转封装。采用此选项时，原视频编码方式必须是 H.264。转封装时，源流的视频关键帧间隔若过大，会影响 RTC 体验，建议 1s，但最大不超过 5s。
	// 默认值为0。
	StreamMode *int32 `json:"StreamMode,omitempty"`

	// 视频高度，转码时必填。单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 视频宽度。转码时必填，单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type StartRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata StartRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                             `json:"Result,omitempty"`
}

type StartRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StartRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type StartRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartSegmentBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制文件的存储平台配置。
	// 支持：
	// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * Amazon S3
	// * 阿里云对象存储 OSS
	// * 华为云 OBS
	// * 腾讯云 COS
	// * 七牛云 Kodo。
	StorageConfig StartSegmentBodyStorageConfig `json:"StorageConfig"`

	// REQUIRED; 截图任务 ID。你可以自行设定TaskId以区分不同的切片任务，且在后续更新和结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	// 若任务运行中，使用相同的 TaskId 重复调用开始接口不会导致请求失败，BaseResponse.Result 会提示 The task has been started. Please do not call the startup
	// task interface repeatedly。
	TaskID string `json:"TaskId"`

	// 业务标识。添加 BusinessId 后，你可以在控制台上查看根据 BusinessId 划分的质量和用量数据。
	BusinessID *string                  `json:"BusinessId,omitempty"`
	Control    *StartSegmentBodyControl `json:"Control,omitempty"`

	// 每个音频切片的时长。单位为秒。范围为 [1, 600]，默认值为 20。
	Duration *int32 `json:"Duration,omitempty"`

	// 是否在开启音视频切片时，立刻开始切片。 True 为在开启音视频切片时，立刻开启切片；False时，在开启音视频切片时，不开始切片，热启动任务。 默认值为 true。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。 切片文件名由 Identifier + UserId + 时间戳 + 序号组成。默认情况下 Identifier 为 随机生成的 UUID。在自定义文件名时，Identifier 的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。
	// 在自定义文件名时，你需确保 Identifier 命名全局唯一，否则在 TOS 平台会因文件名重复被覆盖。
	Identifier *string `json:"Identifier,omitempty"`

	// 任务最大的空闲超时时间。 如果切片任务订阅的所有流都已停止发布，那么任务会在空闲时间超过设定值后自动停止。取值范围为 [1, 86400]，单位为秒，默认值为 180。
	MaxIdleTime   *int32                         `json:"MaxIdleTime,omitempty"`
	TargetStreams *StartSegmentBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type StartSegmentBodyControl struct {

	// 是否开启切片对齐功能，你可以使用该功能，对齐各个用户音频切片的开始和结束时刻。
	// * false：关闭音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。如果用户在切片的周期中，有部分时间未发布音频，返回的音频切片时长会小于切片周期。各个用户音频切片开始时间不一定一致。
	// * true：开启音频切片对齐。在某个切片周期中，如果用户有发送音频流的行为，即生成音频切片。切片长度和切片周期相等，且各个用户音频切片开始的时间一致。如果用户在切片的周期中，有部分时间未发布音频，切片长度不变，这段时间呈现静音帧。如果用户在某个切片周期中始终没有发布音频，则不生成音频切片。
	// 默认值为false。
	Align *bool `json:"Align,omitempty"`

	// 是否忽略静音切片。
	// * true：忽略。
	// * false：不忽略。
	// 默认值为 false
	IgnoreSilence *bool `json:"IgnoreSilence,omitempty"`

	// 是否开启合流切片功能。
	// * False：只会对 TargetStreams 中指定的音频流分别切片。
	// * True：除了会对 TargetStreams 中指定的音频流分别切片，还会对指定的音频流进行混音，生成合流切片，合流切片对应的用户名为 mixed。此时，任务创建后，不管是否有人上麦，会持续回调混音切片。
	// 默认值为 false。 不同平台的回调参看：
	// 操作 ANDROID API IOS API WINDOWS API
	// 本地麦克风录制和远端所有用户混音后的音频数据回调 onMixedAudioFrame [70081#onmixedaudioframe] onMixedAudioFrame: [70087#onmixedaudioframe] onMixedAudioFrame
	// [70096#onmixedaudioframe]
	Mixed *bool `json:"Mixed,omitempty"`

	// 冗余切片时长，单位为毫秒。
	// 当前 RTC 按照传入的Duration值进行固定时长切片,可能存在敏感词被截断，未被识别的情况。此时你可以添加冗余切片，即上一段切片的末尾指定时长，来避免此情况，此时切片的时长变为RedundantDuration + Duration。
	// 例如：当 Duration = 20，redundantDuration = 3 时，最终输出的前三个文件时长为：[0-20], [17-40],
	// [37-60]。
	RedundantDuration *int32 `json:"RedundantDuration,omitempty"`
}

// StartSegmentBodyStorageConfig - 录制文件的存储平台配置。
// 支持：
// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
// * Amazon S3
// * 阿里云对象存储 OSS
// * 华为云 OBS
// * 腾讯云 COS
// * 七牛云 Kodo。
type StartSegmentBodyStorageConfig struct {
	CustomConfig *StartSegmentBodyStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *StartSegmentBodyStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型。支持取值及含义如下：
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 支持 S3 协议的第三方存储平台
	// * 3: VeImageX (当前仅支持抽帧截图功能)
	// 默认值为 0。
	Type           *int32                                       `json:"Type,omitempty"`
	VeImageXConfig *StartSegmentBodyStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *StartSegmentBodyStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type StartSegmentBodyStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。支持取值及含义如下：
	// * 0：Amazon S3
	// * 1：阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。
	// 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type StartSegmentBodyStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSegmentBodyStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	// * 此账号 ID 为火山引擎主账号 ID。
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。 你也可以通过
	// OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为 0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSegmentBodyStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 上传到视频点播平台时, 是否需要根据文件后缀自动设置 FileExtension。关于 FileExtension 的详细说明，参看 FileExtension [https://www.volcengine.com/docs/4/164295#fileextension]。
	// * true：需要；
	// * false：不需要。
	// 默认值为 false。
	AutoSetFileExtension *bool `json:"AutoSetFileExtension,omitempty"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 上传到视频点播平台时, 文件的存储类型。支持取值及含义如下：：
	// * 1：标准存储。
	// * 2：归档存储。
	// * 3：低频存储。
	// 默认值为 1。 关于存储类型的详细说明，参看媒资存储存储类型 [https://www.volcengine.com/docs/4/73629#%E5%AA%92%E8%B5%84%E5%AD%98%E5%82%A8]
	StorageClass *int32 `json:"StorageClass,omitempty"`
}

type StartSegmentBodyTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartSegmentBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartSegmentBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartSegmentRes struct {

	// REQUIRED
	ResponseMetadata StartSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type StartSegmentResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// StartSegmentResResponseMetadataError - 仅在请求失败时返回。
type StartSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartSnapshotBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制文件的存储平台配置。
	// 支持：
	// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * Amazon S3
	// * 阿里云对象存储 OSS
	// * 华为云 OBS
	// * 腾讯云 COS
	// * 七牛云 Kodo。
	StorageConfig StartSnapshotBodyStorageConfig `json:"StorageConfig"`

	// REQUIRED; 截图任务 ID。你可以自行设定TaskId以区分不同的切片任务，且在后续进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	// 若任务运行中，使用相同的 TaskId 重复调用开始接口不会导致请求失败，BaseResponse.Result 会提示 The task has been started. Please do not call the startup
	// task interface repeatedly。
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID  *string                       `json:"BusinessId,omitempty"`
	ImageConfig *StartSnapshotBodyImageConfig `json:"ImageConfig,omitempty"`

	// 任务最大的空闲超时时间。如果抽帧截图任务订阅的所有流都已停止发布，那么任务会在空闲时间超过设定值后自动停止。值的范围为 [1, 86400]，单位为秒，默认值为 180。
	MaxIdleTime   *int32                          `json:"MaxIdleTime,omitempty"`
	TargetStreams *StartSnapshotBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type StartSnapshotBodyImageConfig struct {

	// 图片的格式。支持取值及含义如下：
	// * 0：JEPG
	// * 1：PNG
	// 默认值为0。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，取值范围为 [0, 1920]，单位为像素，默认值为 0，此时，和视频流的实际高度相同。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，取值范围为[1, 600]，单位为秒，默认值为2。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，取值范围为[0, 1920]，单位为像素。默认值为0，表示和视频流的实际宽度相同。
	Width *int32 `json:"Width,omitempty"`
}

// StartSnapshotBodyStorageConfig - 录制文件的存储平台配置。
// 支持：
// * 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
// * 火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
// * Amazon S3
// * 阿里云对象存储 OSS
// * 华为云 OBS
// * 腾讯云 COS
// * 七牛云 Kodo。
type StartSnapshotBodyStorageConfig struct {
	CustomConfig *StartSnapshotBodyStorageConfigCustomConfig `json:"CustomConfig,omitempty"`
	TosConfig    *StartSnapshotBodyStorageConfigTosConfig    `json:"TosConfig,omitempty"`

	// 存储平台类型。支持取值及含义如下：
	// * 0：火山引擎对象存储 TOS [https://www.volcengine.com/product/tos]
	// * 1: 火山引擎视频点播 VOD [https://www.volcengine.com/product/vod]
	// * 2: 支持 S3 协议的第三方存储平台
	// * 3: VeImageX (当前仅支持抽帧截图功能)
	// 默认值为 0。
	Type           *int32                                        `json:"Type,omitempty"`
	VeImageXConfig *StartSnapshotBodyStorageConfigVeImageXConfig `json:"VeImageXConfig,omitempty"`
	VodConfig      *StartSnapshotBodyStorageConfigVodConfig      `json:"VodConfig,omitempty"`
}

type StartSnapshotBodyStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限。不建议开启读权限
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 第三方云存储平台。支持取值及含义如下：
	// * 0：Amazon S3
	// * 1：阿里云 OSS
	// * 2：华为云 OBS
	// * 3：腾讯云 COS
	// * 4：七牛云 Kodo。
	// 默认值为 0。
	Vendor *int32 `json:"Vendor,omitempty"`
}

type StartSnapshotBodyStorageConfigTosConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 存储桶的名称。
	Bucket string `json:"Bucket"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSnapshotBodyStorageConfigVeImageXConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	// * 此账号 ID 为火山引擎主账号 ID。
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 服务 ID。 你可以在 veImageX 控制台 服务管理 [https://console.volcengine.com/imagex/service_manage/]页面，通过创建好的图片服务中获取服务 ID。 你也可以通过
	// OpenAPI 的方式获取服务 ID，具体请参考获取所有服务信息 [https://www.volcengine.com/docs/508/9360]。
	ServiceID string `json:"ServiceId"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为 0。
	Region *int32 `json:"Region,omitempty"`
}

type StartSnapshotBodyStorageConfigVodConfig struct {

	// REQUIRED; 火山引擎平台账号 ID，例如：200000000。
	// * 火山引擎平台账号 ID 查看路径参看查看和管理账号信息 [https://www.volcengine.com/docs/6261/64929]。
	//
	//
	// * 此账号 ID 为火山引擎主账号 ID。
	//
	//
	// * 若你调用 OpenAPI 鉴权过程中使用的 AK、SK 为子用户 AK、SK，账号 ID 也必须为火山引擎主账号 ID，不能使用子用户账号 ID。
	AccountID string `json:"AccountId"`

	// REQUIRED; 点播空间名称。
	Space string `json:"Space"`

	// 上传到视频点播平台时, 是否需要根据文件后缀自动设置 FileExtension。关于 FileExtension 的详细说明，参看 FileExtension [https://www.volcengine.com/docs/4/164295#fileextension]。
	// * true：需要；
	// * false：不需要。
	// 默认值为 false。
	AutoSetFileExtension *bool `json:"AutoSetFileExtension,omitempty"`

	// 不同存储平台支持的 Region 不同，具体参看 Region对照表 [1167931#region]
	// 默认值为0。
	Region *int32 `json:"Region,omitempty"`

	// 上传到视频点播平台时, 文件的存储类型。支持取值及含义如下：：
	// * 1：标准存储。
	// * 2：归档存储。
	// * 3：低频存储。
	// 默认值为 1。 关于存储类型的详细说明，参看媒资存储存储类型 [https://www.volcengine.com/docs/4/73629#%E5%AA%92%E8%B5%84%E5%AD%98%E5%82%A8]
	StorageClass *int32 `json:"StorageClass,omitempty"`
}

type StartSnapshotBodyTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*StartSnapshotBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type StartSnapshotBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type StartSnapshotRes struct {

	// REQUIRED
	ResponseMetadata StartSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                          `json:"Result,omitempty"`
}

type StartSnapshotResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// StartSnapshotResResponseMetadataError - 仅在请求失败时返回。
type StartSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartWBRecordBody struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 AppId。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"AppId"`

	// REQUIRED; 需要录制的白板房间 ID。在同一个 AppId 中，该 ID为每个房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制任务 ID。你可以自行设定 TaskId 以区分不同的白板录制任务。 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// REQUIRED; 任务发起方的用户 ID。不能与房间中其他用户的 ID 重复，否则先进房的用户会被移出房间。
	UserID string `json:"UserId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 自定义 UI 的 Web 页面地址。 不填表示使用默认白板页面。
	SourceURL *string `json:"SourceURL,omitempty"`
}

type StartWBRecordRes struct {

	// REQUIRED
	ResponseMetadata StartWBRecordResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StartWBRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartWBRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StartWBRecordResResponseMetadataError - 仅在请求失败时返回。
type StartWBRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StartWebcastBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 向指定 RTC 房间推送网页音视频内容，房间 ID 是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 需要转推的网页地址，可以携带自定义的 queryParams 来鉴权等，总体长度不超过 1024。
	SourceURL string `json:"SourceURL"`

	// REQUIRED; 任务 ID。你必须对每个云录屏任务设定 TaskId，且在后续结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。
	// 若任务运行中，使用相同的 TaskId 重复调用开始接口不会导致请求失败，BaseResponse.Result 会提示 The task has been started. Please do not call the startup
	// task interface repeatedly。
	TaskID string `json:"TaskId"`

	// REQUIRED; 推送网页音视频内容的用户对应的 UserId。不能与房间中其他用户的 ID 重复，否则先进房的用户会被移出房间。 建议添加有规律的前缀，避免重复。例如， webcast_。
	UserID string `json:"UserId"`

	// 业务标识
	BusinessID        *string                            `json:"BusinessId,omitempty"`
	EventNotifyConfig *StartWebcastBodyEventNotifyConfig `json:"EventNotifyConfig,omitempty"`

	// 最大运行时间，超过此时间后，任务自动终止。取值范围为 [10,86400]，单位为秒，默认值为 86400。不填时自动调整为默认值。
	MaxRunningTime *int32                         `json:"MaxRunningTime,omitempty"`
	MonitorConfig  *StartWebcastBodyMonitorConfig `json:"MonitorConfig,omitempty"`

	// 输出流的视频参数。网页渲染分辨率为输入最大视频流的分辨率。该值为空时按默认值填充一路。
	// * 当设置一路流时，不会输出多路流，输出流参数与设置值相同。
	// * 当设置两路流时， RTC 会启动推送多路流 [70139]功能，设置的两路分辨率相当于分辨率等级划分中的大流和中流的上限。
	VideoSolutions []*StartWebcastBodyVideoSolutionsItem `json:"VideoSolutions,omitempty"`
}

type StartWebcastBodyEventNotifyConfig struct {

	// 是否启用页面主动事件通知,。
	// * false：页面在打开后就会开始采集，在收到 StopWebCast openAPI 请求后结束采集。
	// * true：在页面中注入两个 JS 函数：onWebcastStart()和 onWebcastEnd()。 默认值为false。
	// 当页面判断资源加载完成之后调用onWebcastStart()，控制程序才会开始进行页面内容的采集。当页面判断本次任务内容已完成时调用onWebcastEnd() 通知控制程序结束本次任务。StopWebCast openAPI 效果不变，业务可提前结束任务。其他页面内容、JS
	// 线程的检测（若启用），将在收到 onWebcastStart()事件后才开始。
	// 当启用页面主动事件通知后，你可以参考以下示例代码来通知采集开始。
	// <script>
	// if (ready() && typeof onWebcastStart === 'function')
	// onWebcastStart();
	// </script>
	EnableEventNotify *bool `json:"EnableEventNotify,omitempty"`

	// 启用页面主动事件通知后，等待开始事件的超时时间。取值范围为 [0,60]，单位为秒。默认值为 0，表示不启用。仅当 EnableEventNotify 为 true 时，此参数有效。
	// * 当在超时时间内收到开始事件，采集功能正常运行，用户将收到 Status=1的回调。
	// * 当超时时间内未收到开始事件，将进行刷新，等待时间被重置，再次发生超时后将进行任务重调度。刷新时将回调 Status=4，Reason=" StartEventTimeout"。重调度时将回调 Status=5，Reason="StartEventTimeout"。
	StartTimeout *int32 `json:"StartTimeout,omitempty"`
}

type StartWebcastBodyMonitorConfig struct {

	// 对页面是否白屏的检测间隔。取值范围为 [2,30]，单位为秒。默认值为0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageBlank" 。
	// * 再次出现连续两次检测命中时将进行任务重调度，并回调Status=5，Reason="PageBlank"。
	// 注意：页面全白可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	BlankCheckInterval *int32 `json:"BlankCheckInterval,omitempty"`

	// 对页面 JS 线程是否崩溃/卡死的检测间隔。 取值范围为 [0,30]，单位为秒。默认值为 0，表示不启用。
	// 当出现检测命中时将进行任务重调度，并回调 Status=5，Reason="PageCrash"。
	CrashCheckInterval *int32 `json:"CrashCheckInterval,omitempty"`

	// 对页面内容是否无变化的检测间隔。取值范围为 [2,30]，单位为秒。默认值为 0，表示不启用。
	// * 当连续两次出现检测命中时，将对页面进行刷新，并回调Status=4，Reason="PageFreeze"。
	// * 再次出现连续两次检测命中时，将进行任务重调度，并回调Status=5，Reason="PageFreeze"。
	// 注意：页面无变化可能是您业务的正常场景，请谨慎评估页面实际内容情况后再开启此功能，以免任务提前退出。
	FreezeCheckInterval *int32 `json:"FreezeCheckInterval,omitempty"`
}

type StartWebcastBodyVideoSolutionsItem struct {

	// 最大发送码率，取值范围为 [0,10000]，单位为 Kbps，默认值 0，为 0 时表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，单位为 fps，范围为[1,60]，默认值为 15。帧率和码率设置建议参照视频发布参数对照表 [70122#param]以获取最佳体验。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 视频高度，单位为像素，范围为 [50,1080]，默认值为 720。必须是偶数，值为奇数时自动调整为偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频宽度，单位为像素，范围为 [50,1920]，默认值为 1280。必须是偶数，值为奇数时自动调整为偶数。
	Width *int32 `json:"Width,omitempty"`
}

type StartWebcastRes struct {

	// REQUIRED
	ResponseMetadata StartWebcastResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StartWebcastResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StartWebcastResResponseMetadataError `json:"Error,omitempty"`
}

// StartWebcastResResponseMetadataError - 仅在请求失败时返回。
type StartWebcastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopDetectionBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// 用户 ID
	UserID *string `json:"UserId,omitempty"`
}

type StopDetectionRes struct {

	// REQUIRED
	ResponseMetadata StopDetectionResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result StopDetectionResResult `json:"Result"`
}

type StopDetectionResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopDetectionResResponseMetadataError `json:"Error,omitempty"`
}

// StopDetectionResResponseMetadataError - 仅在请求失败时返回。
type StopDetectionResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopDetectionResResult struct {

	// REQUIRED; 仅在请求成功时返回"success",失败时为空。
	Message string `json:"Message"`
}

type StopPushPublicStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 公共流 ID。
	PublicStreamID string `json:"PublicStreamId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopPushPublicStreamRes struct {

	// REQUIRED
	ResponseMetadata StopPushPublicStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                 `json:"Result,omitempty"`
}

type StopPushPublicStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopPushPublicStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StopPushPublicStreamResResponseMetadataError - 仅在请求失败时返回。
type StopPushPublicStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopPushStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 转推任务 ID。你必须对每个转推任务设定 TaskId，且在进行任务结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type StopPushStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata StopPushStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                `json:"Result,omitempty"`
}

type StopPushStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopPushStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// StopPushStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type StopPushStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopRecordRes struct {

	// REQUIRED
	ResponseMetadata StopRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                       `json:"Result,omitempty"`
}

type StopRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StopRecordResResponseMetadataError - 仅在请求失败时返回。
type StopRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopRelayStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个任务设定 TaskId，且在进行结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type StopRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata StopRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                            `json:"Result,omitempty"`
}

type StopRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// StopRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type StopRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopSegmentBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 你需要关闭的音频切片任务的 ID。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopSegmentRes struct {

	// REQUIRED
	ResponseMetadata StopSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
}

type StopSegmentResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// StopSegmentResResponseMetadataError - 仅在请求失败时返回。
type StopSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopSnapshotBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 你需要关闭的抽帧截图任务的 ID。TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId +
	// RoomId + TaskId 是任务的唯一标识，可以用来标识指定 AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`
}

type StopSnapshotRes struct {

	// REQUIRED
	ResponseMetadata StopSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type StopSnapshotResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// StopSnapshotResResponseMetadataError - 仅在请求失败时返回。
type StopSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopWBRecordBody struct {

	// REQUIRED; 应用的唯一标志
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID。在同一个 appId 中，该参数为每个房间的唯一标志。
	RoomID string `json:"RoomId"`

	// REQUIRED; 录制任务 ID。调用 StartWBRecord 时使用的任务 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 调用接口的用户 ID
	UserID string `json:"UserId"`
}

type StopWBRecordRes struct {

	// REQUIRED
	ResponseMetadata StopWBRecordResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopWBRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopWBRecordResResponseMetadataError `json:"Error,omitempty"`
}

// StopWBRecordResResponseMetadataError - 仅在请求失败时返回。
type StopWBRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type StopWebcastBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个云录屏任务设定 TaskId，且在结束任务时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type StopWebcastRes struct {

	// REQUIRED
	ResponseMetadata StopWebcastResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                        `json:"Result,omitempty"`
}

type StopWebcastResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *StopWebcastResResponseMetadataError `json:"Error,omitempty"`
}

// StopWebcastResResponseMetadataError - 仅在请求失败时返回。
type StopWebcastResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UnbanUserStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 需要被解封音/视频流的用户的 ID
	UserID string `json:"UserId"`

	// 是否解封音频流。
	// * true：解封音频流。
	// * false：封禁音频流。
	// 默认值为 true。
	Audio *bool `json:"Audio,omitempty"`

	// 是否解封视频流。
	// * true：解封视频流。
	// * false：封禁视频流。
	// 默认值为 true。
	Video *bool `json:"Video,omitempty"`
}

type UnbanUserStreamRes struct {

	// REQUIRED
	ResponseMetadata UnbanUserStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *UnbanUserStreamResResult          `json:"Result,omitempty"`
}

type UnbanUserStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UnbanUserStreamResResponseMetadataError `json:"Error,omitempty"`
}

// UnbanUserStreamResResponseMetadataError - 仅在请求失败时返回。
type UnbanUserStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UnbanUserStreamResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 "success", 失败时为空。
	Message string `json:"Message"`
}

type UpdateBanRoomUserRuleBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// 封禁时长，取值范围为 {0,[10,604800]}，单位为秒。 若传入值为空或 0 表示允许用户重新进房。 若传入值小于 10，自动修改为 10。 若传入值大于 604800，自动修改为 604800。
	ForbiddenInterval *int32 `json:"ForbiddenInterval,omitempty"`

	// 指定房间 ID
	RoomID *string `json:"RoomId,omitempty"`

	// 希望封禁用户的 ID
	UserID *string `json:"UserId,omitempty"`
}

type UpdateBanRoomUserRuleRes struct {

	// REQUIRED
	ResponseMetadata UpdateBanRoomUserRuleResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result UpdateBanRoomUserRuleResResult `json:"Result"`
}

type UpdateBanRoomUserRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateBanRoomUserRuleResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateBanRoomUserRuleResResponseMetadataError - 仅在请求失败时返回。
type UpdateBanRoomUserRuleResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateBanRoomUserRuleResResult struct {

	// REQUIRED; 仅在请求成功时返回消息 "success", 失败时为空。
	Message string `json:"Message"`
}

type UpdatePublicStreamParamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 公共流 ID。
	PublicStreamID string `json:"PublicStreamId"`

	// REQUIRED; 为公共流指定单路或多路媒体流及对应参数，Stream 数组。最高支持 17 路。
	TargetStreams []UpdatePublicStreamParamBodyTargetStreamsItem `json:"TargetStreams"`

	// 业务标识
	BusinessID *string                             `json:"BusinessId,omitempty"`
	Control    *UpdatePublicStreamParamBodyControl `json:"Control,omitempty"`
	Encode     *UpdatePublicStreamParamBodyEncode  `json:"Encode,omitempty"`

	// 你可以通过本参数排除掉不需要包含在公共流中的音视频流，即黑名单。参数默认为空。黑名单中的流不得超过 17 路。此参数中的 stream 不应和 TargetStreams 中重复。
	ExcludeStreams []*UpdatePublicStreamParamBodyExcludeStreamsItem `json:"ExcludeStreams,omitempty"`
	Layout         *UpdatePublicStreamParamBodyLayout               `json:"Layout,omitempty"`

	// 公共流处理模式。支持取值及含义如下：0：转码。1：转封装。
	// 当 TranscodeMode=1 时，
	// * TargetStreams 只能指定一路流，且该路流的 UserId不能为空，需为对应房间用户的 UserId。
	// * ExcludeStreams 必须为空。
	// * Encode.VideoConfig 设置不生效。
	// * Layout 设置不生效。
	TranscodeMode *int32 `json:"TranscodeMode,omitempty"`
}

type UpdatePublicStreamParamBodyControl struct {

	// 插入公共流的自定义信息，可用于随流信息同步，长度不超过 4 kB。 数据会添加到当前视频帧开始的连续 30 个视频帧中。 只在调用UpdatePublicStreamParam时有效。
	DataMsg *string `json:"DataMsg,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为[10, 86400]，单位为秒，默认值为180。只在调用StartPushPublicStream时有效。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 房间用户发布状态回调间隔，仅在纯音频时生效。取值范围为 [1000,2147483647]，单位为毫秒，默认值为 2000。
	StreamPublishStatsInterval *int32 `json:"StreamPublishStatsInterval,omitempty"`

	// 是否开启房间用户发布状态回调。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户发布状态回调。
	// * false：不开启房间用户发布状态回调。
	// 默认值为false。
	StreamPublishStatsMode *bool `json:"StreamPublishStatsMode,omitempty"`

	// 房间用户采集状态回调间隔，仅在纯音频时生效。取值范围为[1000,2147483647]，单位为毫秒，默认值为2000。
	UserCaptureStatsInterval *int32 `json:"UserCaptureStatsInterval,omitempty"`

	// 是否开启房间用户采集状态回调。开启后会通过onPublicStreamDataMessageReceived回调。
	// * true：开启房间用户采集状态回调。
	// * false：不开启房间用户采集状态回调。
	// 默认值为false。
	UserCaptureStatsMode *bool `json:"UserCaptureStatsMode,omitempty"`

	// 音量指示的回调间隔。最小值为100，单位为毫秒，默认值为2000。
	// VideoConfig.FrameRate大于 10 fps 时，回调间隔才能达到 100ms。
	VolumeIndicationInterval *int32 `json:"VolumeIndicationInterval,omitempty"`

	// 是否开启音量指示模式。
	// * true：开启音量提示。
	// * false：不开启音量提示。
	// 默认值为false。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type UpdatePublicStreamParamBodyEncode struct {

	// 视频编码配置
	VideoConfig *UpdatePublicStreamParamBodyEncodeVideoConfig `json:"VideoConfig,omitempty"`
}

// UpdatePublicStreamParamBodyEncodeVideoConfig - 视频编码配置
type UpdatePublicStreamParamBodyEncodeVideoConfig struct {

	// 最高输出视频码率。取值范围为[0,10000]，单位为 Kbps，默认值为0，0表示自适应码率。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 输出视频帧率。取值范围为 [1,60]，单位为 fps，默认为15。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 输出画面的高度，默认值为 480。范围为 [16, 1920]，必须是偶数。
	Height *int32 `json:"Height,omitempty"`

	// 视频编码协议。支持取值及含义如下：
	// * 0：H.264。
	// * 5：VP8。 如果选择 VP8 格式，请先联系火山技术支持配置。
	// 默认值为0。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 输出画面的宽度。默认值为 640，范围为 [16, 1920]，必须是偶数。
	Width *int32 `json:"Width,omitempty"`
}

type UpdatePublicStreamParamBodyExcludeStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	//
	//
	// * 1：纯音频
	//
	//
	// * 2：纯视频
	//
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。支持取值及含义如下：
	// * 0：音视频流
	//
	//
	// * 1：屏幕流
	//
	// 默认值为 0。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。UserId 为空时，表示订阅房间内所有流。UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type UpdatePublicStreamParamBodyLayout struct {

	// REQUIRED; 布局模式。支持取值及含义如下：
	// * 0：自适应布局 [1167930#adapt]
	// * 1：：垂直布局 [1167930#vertical]
	// * 2 ：自定义布局
	// * 3 ：并排布局 [1167930#horizontal]
	// 默认值为0
	LayoutMode     int32                                            `json:"LayoutMode"`
	CustomLayout   *UpdatePublicStreamParamBodyLayoutCustomLayout   `json:"CustomLayout,omitempty"`
	VerticalLayout *UpdatePublicStreamParamBodyLayoutVerticalLayout `json:"VerticalLayout,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutCustomLayout struct {

	// REQUIRED; 自定义布局下，多路视频配置
	Regions []UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItem `json:"Regions"`

	// 背整体屏幕（画布）的背景色，格式为 #RGB(16进制)，默认值为 #000000（黑色）, 范围为 #000000 ~ #ffffff (大小写均可)。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundColor *string `json:"BackgroundColor,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。
	// 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 补帧模式。支持取值及含义如下：
	// * 0：补最后一帧，
	// * 1：补黑帧。
	// 默认值为0。自动布局模式下，该参数不生效。 补帧是指在音视频录制时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 如果同时配置Alternateimage 和FrameInterpolationMode ，优先使用 Alternateimage参数。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	HeightProportion float32 `json:"HeightProportion"`

	// REQUIRED; 流标识。
	// StreamIndex 即 Stream.Index，用来指定布局设置将被应用到哪路流。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度相对整体画面的比例，取值的范围为 (0.0, 1.0]。
	WidthProportion float32 `json:"WidthProportion"`

	// 画面的透明度，取值范围为 (0.0, 1.0]，默认值为 1.0。0.0 表示完全透明，1.0 表示完全不透明。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 占位图片的 url
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的纵坐标相对整体画面的比例，取值的范围为 [0.0, 1.0)，默认值为 0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 ：按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 ：按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形。
	// 默认值为 0。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 当画面有重叠时，使用此参数设置指定画面的图层顺序，取值范围为 [0, 100]，默认值为 0。0 表示该区域图像位于最下层，100 表示该区域图像位于最上层。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdatePublicStreamParamBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutVerticalLayout struct {
	MainStream *UpdatePublicStreamParamBodyLayoutVerticalLayoutMainStream `json:"MainStream,omitempty"`
}

type UpdatePublicStreamParamBodyLayoutVerticalLayoutMainStream struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	//
	//
	// * 1：纯音频
	//
	//
	// * 2：纯视频
	//
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。支持取值及含义如下：
	// * 0：音视频流
	//
	//
	// * 1：屏幕流
	//
	// 默认值为 0。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。UserId 为空时，表示订阅房间内所有流。UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type UpdatePublicStreamParamBodyTargetStreamsItem struct {

	// REQUIRED; 发布公共流的用户所在的房间 ID
	RoomID string `json:"RoomId"`

	// 当选择自定义布局模式时，此字段必填。标记同一路公共流中不同的媒体流。 在同一个 TargetStreams 中，Stream.Index 是唯一的。
	Index *int32 `json:"Index,omitempty"`

	// 流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	//
	//
	// * 1：纯音频
	//
	//
	// * 2：纯视频
	//
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 流类型。支持取值及含义如下：
	// * 0：音视频流
	//
	//
	// * 1：屏幕流
	//
	// 默认值为 0。
	StreamType *int32 `json:"StreamType,omitempty"`

	// 媒体流的发布方的用户 ID。UserId 为空时，表示订阅房间内所有流。UserId 需全局唯一。不同房间内的 UserId 不能重复。
	UserID *string `json:"UserId,omitempty"`
}

type UpdatePublicStreamParamRes struct {

	// REQUIRED
	ResponseMetadata UpdatePublicStreamParamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                    `json:"Result,omitempty"`
}

type UpdatePublicStreamParamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdatePublicStreamParamResResponseMetadataError `json:"Error,omitempty"`
}

// UpdatePublicStreamParamResResponseMetadataError - 仅在请求失败时返回。
type UpdatePublicStreamParamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdatePushMixedStreamToCDNBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 合流转推任务 ID。你必须对每个合流转推任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string                                `json:"BusinessId,omitempty"`
	Control    *UpdatePushMixedStreamToCDNBodyControl `json:"Control,omitempty"`

	// 音视频编码参数。
	// * 单流录制时，你仅可以设置VideoFps、VideoBitrate和所有音频相关参数。
	// * 合流录制时，你仅可以设置VideoWidth、VideoHeight、VideoFps、VideoBitrate和所有音频相关参数。
	Encode *UpdatePushMixedStreamToCDNBodyEncode `json:"Encode,omitempty"`

	// 是否更新部分参数。
	// * false：否。
	// * true：是。
	// 默认值为 false。 开启部分更新后，必须按照参数层级传入，且数组类参数需要传入该数组中所有参数。
	IsUpdatePartialParam *bool                                 `json:"IsUpdatePartialParam,omitempty"`
	Layout               *UpdatePushMixedStreamToCDNBodyLayout `json:"Layout,omitempty"`

	// 更新请求序列号。填写该参数后，服务端会对请求进行校验，请确保最后一次更新请求的序列号大于前一次请求的序列号。
	// 建议更新部分参数场景下传入此参数，以确保服务端按照最新请求进行更新。
	SequenceNumber *int32                                       `json:"SequenceNumber,omitempty"`
	TargetStreams  *UpdatePushMixedStreamToCDNBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControl struct {

	// 选择补帧模式。支持取值及含义如下：
	// * 0：补最后一帧。
	// * 1：补黑帧
	// 默认值为0。 自动布局模式下，没有补帧的逻辑。 补帧是指在音视频录制或合流转推时，视频的帧率通常是固定的。但是，因为网络波动或其他原因，实际帧率可能无法达到预设的帧率。此时，需要补帧以保持视频流畅。补最后一帧意味着补充的视频帧和中断前的最后一帧相同，此时看到的画面可能会有短暂静止；补黑帧意味着补充的视频帧是全黑的。
	// 使用占位图、补帧和上一帧的关系: 你可以在 Region 中传入
	// Alternateimage 字段设置占位图,在 Control 中传入FrameInterpolationMode 字段设置补帧模式,但占位图优先级高于补帧模式。
	// * 在 Region.StreamIndex 对应的视频流停止发布时, Region 对应的画布空间会根据设置填充占位图或补帧。但当视频流为屏幕流时，补帧模式不生效。
	// * 当 Region.StreamIndex 对应的视频流发布后停止采集或推送时, Region 对应的画布空间会填充上一帧。
	// * 当 Region.StreamIndex 对应的视频流发布时,设置的占位图或补顿模式不造成任何影响。
	FrameInterpolationMode *int32 `json:"FrameInterpolationMode,omitempty"`

	// 任务的空闲超时时间，超过此时间后，任务自动终止。取值范围为 [10, 86400],单位为秒，默认值为 180。
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty"`

	// 流的类型，用于全局控制订阅的流的类型。支持取值及含义如下：
	// * 0 ：音视频；
	// * 1 ：纯音频。
	// 默认值为 0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 转推直播推流模式，用于控制触发推流的时机。支持取值及含义如下：
	// * 0：房间内有用户推 RTC 流时即触发 CDN 推流。
	// * 1：调用接口发起推流任务后，无论房间内是否有用户推 RTC 流,均可触发 CDN 推流。
	// 默认值为 0。 任务超时逻辑不变，依然是无用户推流即判定为超时。
	PushStreamMode *int32                                              `json:"PushStreamMode,omitempty"`
	SEIConfig      *UpdatePushMixedStreamToCDNBodyControlSEIConfig     `json:"SEIConfig,omitempty"`
	SpatialConfig  *UpdatePushMixedStreamToCDNBodyControlSpatialConfig `json:"SpatialConfig,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControlSEIConfig struct {

	// SEI 中是否包含用户说话音量值。
	// * false：否。
	// * true：是。
	// 默认值为 false。
	AddVolumeValue *bool `json:"AddVolumeValue,omitempty"`

	// 是否透传 RTC 流里的 SEI 信息。
	// * true：是。
	// * false：否。
	// 默认值为 true。
	PassthroughRTCSEIMessage *bool `json:"PassthroughRTCSEIMessage,omitempty"`

	// 开启音量指示模式后，非关键帧携带 SEI 包含的信息类型。支持取值及含义如下：
	// * 0：全量信息。
	// * 1：只有音量信息和时间戳。
	// 默认值为 0。关于 SEI 结构，参看服务端合流转推 SEI 结构 [1163740#sei]
	SEIContentMode *int32 `json:"SEIContentMode,omitempty"`

	// 设置 SEI 的 Payload Type，对 服务端合流转推 SEI 和 RTC 透传SEI 均生效。支持取值为 5 或 100，默认为 100。
	SEIPayLoadType *int32 `json:"SEIPayLoadType,omitempty"`

	// 该字段为长度为 32 位的 16 进制字符，每个字符的范围为 a-f，A-F，0-9，不可包含 -。如果校验失败，会返回错误码 InvalidParameter。
	// 仅当 SEIPayLoadType=5时，该字段需要填写，SEIPayLoadType=100时，该字段内容会被忽略。
	// 此参数仅对 RTC 透传SEI 生效。当用户设置 SEIPayloadType = 5 时，服务端合流转推SEI 的 SEIPayloadUUID 为固定值：566f6c63656e67696e65525443534549，对应16位字符串
	// VolcengineRTCSEI。
	SEIPayloadUUID *string `json:"SEIPayloadUUID,omitempty"`

	// 设置SEI数据结构 [1163740#sei]中 app_data 参数的值，最大长度为 4096。此参数支持动态更新。
	UserConfigExtraInfo *string `json:"UserConfigExtraInfo,omitempty"`

	// 用户说话音量的回调间隔。取值范围为 [0.3,∞]，单位为秒，默认值为 2。仅当用户说话音量发生变化时此回调才会被触发。
	VolumeIndicationInterval *float32 `json:"VolumeIndicationInterval,omitempty"`

	// 是否开启音量指示模式。
	// * true：是。此时非关键帧中也可能携带 SEI 信息。
	// * false：否。此时只会在关键帧中携带 SEI 信息。
	// 默认值为 false。
	// 若 VolumeIndicationMode = true 的同时设置 MediaType = 1，该流推向 CDN 地址时，服务端会补黑帧。 关于音量指示模式的用法，参看 音量指示模式 [1163740#volume]。
	VolumeIndicationMode *bool `json:"VolumeIndicationMode,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControlSpatialConfig struct {
	AudienceSpatialOrientation *UpdatePushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation `json:"AudienceSpatialOrientation,omitempty"`

	// 观众所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	AudienceSpatialPosition []*int32 `json:"AudienceSpatialPosition,omitempty"`

	// 是否开启空间音频处理功能。
	// * false：关闭。
	// * true：开启
	// 默认值为 false。
	EnableSpatialRender *bool `json:"EnableSpatialRender,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyControlSpatialConfigAudienceSpatialOrientation struct {

	// 前方朝向
	Forward []*float32 `json:"Forward,omitempty"`

	// 右边朝向
	Right []*float32 `json:"Right,omitempty"`

	// 头顶朝向
	Up []*float32 `json:"Up,omitempty"`
}

// UpdatePushMixedStreamToCDNBodyEncode - 音视频编码参数。
// * 单流录制时，你仅可以设置VideoFps、VideoBitrate和所有音频相关参数。
// * 合流录制时，你仅可以设置VideoWidth、VideoHeight、VideoFps、VideoBitrate和所有音频相关参数。
type UpdatePushMixedStreamToCDNBodyEncode struct {

	// 音频码率。取值范围为 [32,192],单位为 Kbps。
	// * 该参数若不填，RTC 会根据音频配置文件类型、采样率、声道数的配置生成音频码率取值。
	// * 该参数若填写，实际生效的码率也会随着音频配置文件类型、采样率、声道数的配置而变化。若你发现生效的码率值没有达到你设定的值，可能是该值已经超过该场景下的极限码率。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 音频声道数。支持取值及含义如下：
	// * 1：单声道。
	// * 2：双声道。
	// 默认值为 2。
	AudioChannels *int32 `json:"AudioChannels,omitempty"`

	// 音频编码协议。仅支持取 0，表示 AAC 编码协议。
	AudioCodec *int32 `json:"AudioCodec,omitempty"`

	// 音频配置文件类型。支持取值及含义如下：
	// * 0 :采用 LC 规格；
	// * 1: 采用 HE-AAC v1 规格；
	// * 2: 采用 HE-AAC v2 规格。此时，只支持输出流音频声道数为双声道。
	// 默认值为 0。
	AudioProfile *int32 `json:"AudioProfile,omitempty"`

	// 音频采样率。可取值为：{32000,44100，48000}，单位为 Hz，默认值为 48000。
	AudioSampleRate *int32 `json:"AudioSampleRate,omitempty"`

	// 视频码率。取值范围为 [0,10000]，单位为 Kbps，默认值为 0。0 表示自适应码率，会自动根据 VideoFps , VideoWidth 以及VideoHeight 计算出合理的码率。 自适应码率模式下，RTC 默认不会设置超高码率。如果订阅屏幕流，建议自行设置高码率。不同场景下设置码率等视频发布参数,请参考设置视频发布参数
	// [70122#videoprofiles]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频编码协议。支持取值及含义如下：
	// * 0：H.264
	// * 1：ByteVC1
	// 默认值为 0。
	VideoCodec *int32 `json:"VideoCodec,omitempty"`

	// 视频帧率。取值范围为 [1,60]，单位为 FPS，默认值为 15。
	VideoFps *int32 `json:"VideoFps,omitempty"`

	// 输出视频 GOP。取值范围为 [1,5]，单位为秒，默认值为 4。
	VideoGop *int32 `json:"VideoGop,omitempty"`

	// 画面高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Height 设置画面宽度。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 画面宽度。取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。该参数在垂直布局和并排布局下生效，自定义布局下请使用 canvas.Width 设置画面宽度。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayout struct {
	CustomLayout *UpdatePushMixedStreamToCDNBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。支持取值及含义如下：
	// * 0：自适应布局 [1167930#adapt]。
	// * 1：垂直布局 [1167930#vertical]
	// * 2 ：自定义布局。
	// * 3 ：并排布局 [1167930#horizontal]
	// 默认值为0
	LayoutMode      *int32                                               `json:"LayoutMode,omitempty"`
	MainVideoStream *UpdatePushMixedStreamToCDNBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayout struct {
	Canvas *UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。
	Regions []*UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色, 范围为 #000000 ~ #ffffff (大小写均可)，格式为 #RGB(16进制)，默认值为 #000000（黑色）。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。 如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。
	Width *int32 `json:"Width,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]，默认值为 1.0。0.0 表示完全透明，1.0 表示完全不透明。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 [0,1]。默认值为 0，表示没有圆角效果。 圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。 例如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 ：按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 ：按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形。
	// 默认值为 0。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]，默认值为 0。0 表示该区域图像位于最下层，100 表示该区域图像位于最上层。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdatePushMixedStreamToCDNBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*UpdatePushMixedStreamToCDNBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type UpdatePushMixedStreamToCDNBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdatePushMixedStreamToCDNRes struct {

	// REQUIRED
	ResponseMetadata UpdatePushMixedStreamToCDNResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                                       `json:"Result,omitempty"`
}

type UpdatePushMixedStreamToCDNResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdatePushMixedStreamToCDNResResponseMetadataError `json:"Error,omitempty"`
}

// UpdatePushMixedStreamToCDNResResponseMetadataError - 仅在请求失败时返回。
type UpdatePushMixedStreamToCDNResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateRecordBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 云端录制任务 ID。你必须对每个云端录制任务设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID    *string                        `json:"BusinessId,omitempty"`
	Layout        *UpdateRecordBodyLayout        `json:"Layout,omitempty"`
	TargetStreams *UpdateRecordBodyTargetStreams `json:"TargetStreams,omitempty"`
}

type UpdateRecordBodyLayout struct {
	CustomLayout *UpdateRecordBodyLayoutCustomLayout `json:"CustomLayout,omitempty"`

	// 布局模式。支持取值及含义如下：
	// * 0：自适应布局 [1167930#adapt]。
	// * 1：垂直布局 [1167930#vertical]
	// * 2 ：自定义布局。
	// * 3 ：并排布局 [1167930#horizontal]
	// 默认值为0
	LayoutMode      *int32                                 `json:"LayoutMode,omitempty"`
	MainVideoStream *UpdateRecordBodyLayoutMainVideoStream `json:"MainVideoStream,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayout struct {
	Canvas *UpdateRecordBodyLayoutCustomLayoutCanvas `json:"Canvas,omitempty"`

	// 在自定义布局模式下，你可以使用 Regions 对每一路视频流进行画面布局设置。其中，每个 Region 对一路视频流进行画面布局设置。
	// 自定义布局模式下，对于 StreamList 中的每个 Stream，Regions 中都需要给出对应的布局信息，否则会返回请求不合法的错误。即 Regions.Region.StreamIndex 要与 TargetStreams.StreamList.Stream.Index
	// 的值一一对应，否则自定义布局设置失败，返回 InvalidParameter 错误码。
	// > 当传入的必填参数值不合法时，返回错误码 InvalidParameter 。
	Regions []*UpdateRecordBodyLayoutCustomLayoutRegionsItem `json:"Regions,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayoutCanvas struct {

	// 整体屏幕（画布）的背景色, 范围为 #000000 ~ #ffffff (大小写均可)，格式为 #RGB(16进制)，默认值为 #000000（黑色）。
	Background *string `json:"Background,omitempty"`

	// 背景图片的 URL。长度最大为 1024 byte。可以传入的图片的格式包括：JPG, JPEG, PNG。 如果背景图片的宽高和整体屏幕的宽高不一致，背景图片会缩放到铺满屏幕。 如果你设置了背景图片，背景图片会覆盖背景色。
	BackgroundImage *string `json:"BackgroundImage,omitempty"`

	// 整体屏幕（画布）的高度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 480。
	Height *int32 `json:"Height,omitempty"`

	// 整体屏幕（画布）的宽度，取值范围为 [2, 1920]，必须是偶数，单位为像素，默认值为 640。
	Width *int32 `json:"Width,omitempty"`
}

type UpdateRecordBodyLayoutCustomLayoutRegionsItem struct {

	// REQUIRED; 视频流对应区域高度的像素绝对值，取值的范围为 (0.0, Canvas.Height]。
	Height int32 `json:"Height"`

	// REQUIRED; 流的标识。这个标志应和 TargetStreams.StreamList.Stream.Index 对应。
	StreamIndex int32 `json:"StreamIndex"`

	// REQUIRED; 视频流对应区域宽度的像素绝对值，取值的范围为 (0.0, Canvas.Width]。
	Width int32 `json:"Width"`

	// 画面的透明度，取值范围为 (0.0, 1.0]，默认值为 1.0。0.0 表示完全透明，1.0 表示完全不透明。
	Alpha *float32 `json:"Alpha,omitempty"`

	// 补位图片的 url。长度不超过 1024 个字符串。
	// * 在 Region.StreamIndex 对应的视频流没有发布，或被暂停采集时，AlternateImage 对应的图片会填充 Region 对应的画布空间。当视频流被采集并发布时，AlternateImage 不造成任何影响。
	// * 可以传入的图片的格式包括：JPG, JPEG, PNG。
	// * 当图片和画布尺寸不一致时，图片根据 RenderMode 的设置，在画布中进行渲染。
	AlternateImage *string `json:"AlternateImage,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0：按照用户原始视频帧比例进行缩放
	// * 1：保持图片原有比例
	// 默认值为 0。
	AlternateImageFillMode *int32 `json:"AlternateImageFillMode,omitempty"`

	// 该路流对应的用户是否开启空间音频效果。
	// * true：开启空间音频效果。
	// * false：关闭空间音频效果。
	// 默认值为 true
	ApplySpatialAudio *bool `json:"ApplySpatialAudio,omitempty"`

	// 转推直播下边框圆角半径与画布宽度的比例值，取值范围为 [0,1]。默认值为 0，表示没有圆角效果。 圆角半径的像素位不能超过 Region 宽高最小值的一半，否则不会出现圆角效果。
	CornerRadius *float32 `json:"CornerRadius,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的横向位移，取值的范围为 [0.0, Canvas.Width)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	// 视频流对应区域左上角的实际坐标是通过画面尺寸和相对位置比例相乘，并四舍五入取整得到的。假如，画布尺寸为1920, 视频流对应区域左上角的横坐标相对整体画面的比例为 0.33，那么该画布左上角的横坐标为 634（1920*0.33=633.6，四舍五入取整）。
	LocationX *int32 `json:"LocationX,omitempty"`

	// 视频流对应区域左上角的横坐标相对整体画面左上角原点的纵向位移，取值的范围为 [0.0, Canvas.Height)。默认值为 0。若传入该参数，服务端会对该参数进行校验，若不合法会返回错误码 InvalidParameter。
	LocationY *int32 `json:"LocationY,omitempty"`

	// 该路流参与混流的媒体类型。支持取值及含义如下：
	// * 0：音视频
	// * 1：纯音频
	// * 2：纯视频
	// 默认值为 0。 例如该路流为音视频流，MediaType设为1，则只混入音频内容。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 画面的渲染模式。支持取值及含义如下：
	// * 0 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形
	// * 1 ：按照显示区域的长宽比裁减视频，然后等比拉伸或缩小视频，占满显示区域。
	// * 2 ：按照原始画面的宽高比缩放视频，在显示区域居中显示。如果原始画面宽高比与指定的宽高比不同，就会导致画面有空缺，空缺区域为填充的背景色值。
	// * 3 ：按照指定的宽高直接缩放。如果原始画面宽高比与指定的宽高比不同，就会导致画面变形。
	// 默认值为 0。 目前 0 和 3 均为按照指定的宽高直接缩放，但我们推荐你使用 3 以便与客户端实现相同逻辑。
	RenderMode *int32 `json:"RenderMode,omitempty"`

	// 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
	SourceCrop *UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop `json:"SourceCrop,omitempty"`

	// 空间音频下，房间内指定用户所在位置的三维坐标，默认值为[0,0,0]。数组长度为3，三个值依次对应X,Y,Z，每个值的取值范围为 [-100,100]。
	SpatialPosition []*int32 `json:"SpatialPosition,omitempty"`

	// 当多个流的画面有重叠时，使用此参数设置指定画面的图层顺序。取值范围为 [0, 100]，默认值为 0。0 表示该区域图像位于最下层，100 表示该区域图像位于最上层。
	ZOrder *int32 `json:"ZOrder,omitempty"`
}

// UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop - 如果裁剪后计算得到的实际分辨率的宽或高不是偶数，会被自动调整为偶数
type UpdateRecordBodyLayoutCustomLayoutRegionsItemSourceCrop struct {

	// 裁剪后得到的视频帧高度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	HeightProportion *float32 `json:"HeightProportion,omitempty"`

	// 裁剪后得到的视频帧左上角的横坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationX *float32 `json:"LocationX,omitempty"`

	// 裁剪后得到的视频帧左上角的纵坐标相对裁剪前整体画面的比例，取值的范围为 [0.0, 1.0)。默认值为 0.0。
	LocationY *float32 `json:"LocationY,omitempty"`

	// 裁剪后得到的视频帧宽度相对裁剪前整体画面宽度的比例，取值范围为 (0.0, 1.0]。默认值为 1.0。
	WidthProportion *float32 `json:"WidthProportion,omitempty"`
}

type UpdateRecordBodyLayoutMainVideoStream struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdateRecordBodyTargetStreams struct {

	// 音视频流列表，由Stream组成，可以为空。为空时，表示订阅房间内所有流。在一个 StreamList 中，Stream.Index 不能重复。
	StreamList []*UpdateRecordBodyTargetStreamsStreamListItem `json:"StreamList,omitempty"`
}

type UpdateRecordBodyTargetStreamsStreamListItem struct {

	// REQUIRED; 用户 ID，表示这个流所属的用户。
	UserID string `json:"UserId"`

	// 在自定义布局中，使用 Index 对流进行标志。后续在 Layout.regions.StreamIndex 中，你需要使用 Index 指定对应流的布局设置。
	Index *int32 `json:"Index,omitempty"`

	// 流的类型。支持取值及含义如下：
	// * 0：普通音视频流，
	// * 1：屏幕流。
	// 默认值为0。
	StreamType *int32 `json:"StreamType,omitempty"`
}

type UpdateRecordRes struct {

	// REQUIRED
	ResponseMetadata UpdateRecordResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                         `json:"Result,omitempty"`
}

type UpdateRecordResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateRecordResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateRecordResResponseMetadataError - 仅在请求失败时返回。
type UpdateRecordResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateRelayStreamBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED
	Control UpdateRelayStreamBodyControl `json:"Control"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 任务 ID。你必须对每个任务设定 TaskId，且在进行任务更新和结束时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`
}

type UpdateRelayStreamBodyControl struct {

	// REQUIRED; 在线流媒体地址。媒体格式应为：hls、rtmp、mp4、flv、dash、或 ts。如源流为海外，建议联系技术支持，以保障最佳体验。
	StreamURL string `json:"StreamUrl"`

	// 最大发送码率，单位为 Kbps,不填则不限制，转码时生效。
	Bitrate *int32 `json:"Bitrate,omitempty"`

	// 发送帧率，值的范围为 [1，30]，默认值为15，转码时生效。
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 是否循环播放，仅对源流为点播流时生效。
	Loop *bool `json:"Loop,omitempty"`

	// 媒体类型。
	// * 0：音视频
	// * 1：音频。采用此选项时，必须是 AAC 或 Opus 编码。
	// * 2：视频
	// 默认值为0。
	MediaType *int32 `json:"MediaType,omitempty"`

	// 任务起始时间戳，用于定时播放，Unix 时间，单位为秒。默认为 0，表示立即启动。此参数仅对 StartRelayStream接口生效。
	StartTimeStamp *int32 `json:"StartTimeStamp,omitempty"`

	// 流处理模式。
	// * 0：转码。采用此选项时，原视频编码方式必须是 H.264 或 ByteVC1。
	// * 1：转封装。采用此选项时，原视频编码方式必须是 H.264。转封装时，源流的视频关键帧间隔若过大，会影响 RTC 体验，建议 1s，但最大不超过 5s。
	// 默认值为0。
	StreamMode *int32 `json:"StreamMode,omitempty"`

	// 视频高度，转码时必填。单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoHeight *int32 `json:"VideoHeight,omitempty"`

	// 视频宽度。转码时必填，单位为像素，范围为 [16, 1920]，必须是偶数，值为奇数时自动调整为偶数。
	VideoWidth *int32 `json:"VideoWidth,omitempty"`
}

type UpdateRelayStreamRes struct {

	// REQUIRED
	ResponseMetadata UpdateRelayStreamResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                              `json:"Result,omitempty"`
}

type UpdateRelayStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateRelayStreamResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateRelayStreamResResponseMetadataError - 仅在请求失败时返回。
type UpdateRelayStreamResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateSegmentBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 切片任务 ID。你必须对每个切片任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID *string `json:"BusinessId,omitempty"`

	// 音频切片的时长。单位为秒。取值范围为 [1, 600]，默认值为 20。 更新该字段后，计时器会重新启动，当前切片立即停止，生成一个新切片。
	Duration *int32 `json:"Duration,omitempty"`

	// 是否在开启音视切片时，立刻开始切片。
	// * true：立刻开启切片。
	// * false：不切片。
	// 默认值 true。
	Handle *bool `json:"Handle,omitempty"`

	// 自定义文件前缀。 切片文件名由 Identifier + UserId + 时间戳 + 序号组成。默认情况下 Identifier 为 随机生成的 UUID。在自定义文件名时，Identifier 的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}。
	// 在自定义文件名时，你需确保 identifier 命名全局唯一，否则在 TOS 平台会因文件名重复被覆盖。
	Identifier *string `json:"Identifier,omitempty"`
}

type UpdateSegmentRes struct {

	// REQUIRED
	ResponseMetadata UpdateSegmentResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                          `json:"Result,omitempty"`
}

type UpdateSegmentResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateSegmentResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateSegmentResResponseMetadataError - 仅在请求失败时返回。
type UpdateSegmentResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type UpdateSnapshotBody struct {

	// REQUIRED; 你的音视频应用的唯一标志，参看获取 AppId [https://www.volcengine.com/docs/6348/69865#%E6%AD%A5%E9%AA%A44%EF%BC%9A%E5%88%9B%E5%BB%BA-rtc-%E5%BA%94%E7%94%A8%EF%BC%8C%E8%8E%B7%E5%8F%96-appid]。
	AppID string `json:"AppId"`

	// REQUIRED; 房间的 ID，是房间的唯一标志
	RoomID string `json:"RoomId"`

	// REQUIRED; 截图任务 ID。你必须对每个截图任务，设定 TaskId，且在进行任务更新时也须使用该 TaskId。
	// TaskId 是任务的标识，在一个 AppId 的 RoomId 下 taskId 是唯一的，不同 AppId 或者不同 RoomId 下 TaskId 可以重复，因此 AppId + RoomId + TaskId 是任务的唯一标识，可以用来标识指定
	// AppId 下某个房间内正在运行的任务，从而能在此任务运行中进行更新或者停止此任务。
	// 关于 TaskId 及以上 Id 字段的命名规则符合正则表达式：[a-zA-Z0-9_@\-\.]{1,128}
	TaskID string `json:"TaskId"`

	// 业务标识
	BusinessID  *string                        `json:"BusinessId,omitempty"`
	ImageConfig *UpdateSnapshotBodyImageConfig `json:"ImageConfig,omitempty"`
}

type UpdateSnapshotBodyImageConfig struct {

	// 图片的格式。支持取值及含义如下：
	// * 0：JEPG
	// * 1：PNG
	// 默认值为0。
	Format *int32 `json:"Format,omitempty"`

	// 实际使用视频帧的高度，取值范围为 [0, 1920]，单位为像素，默认值为 0，此时，和视频流的实际高度相同。
	Height *int32 `json:"Height,omitempty"`

	// 相邻截图之间的间隔时间，取值范围为[1, 600]，单位为秒，默认值为2。
	Interval *int32 `json:"Interval,omitempty"`

	// 实际使用视频帧的宽度，取值范围为[0, 1920]，单位为像素。默认值为0，表示和视频流的实际宽度相同。
	Width *int32 `json:"Width,omitempty"`
}

type UpdateSnapshotRes struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                           `json:"Result,omitempty"`
}

type UpdateSnapshotResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *UpdateSnapshotResResponseMetadataError `json:"Error,omitempty"`
}

// UpdateSnapshotResResponseMetadataError - 仅在请求失败时返回。
type UpdateSnapshotResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type WbTranscodeCreateBody struct {

	// REQUIRED; RTC AppId
	AppID string `json:"AppId"`

	// REQUIRED; 发起转码的Operator
	Operator string `json:"Operator"`

	// REQUIRED; 转码优先级
	Priority int32 `json:"Priority"`

	// REQUIRED; 源文件URL
	Resource string `json:"Resource"`

	// REQUIRED; 转码模式
	TranscodeMode int32 `json:"TranscodeMode"`

	// TOS 桶中存储对象的预签名 URL 有限期。单位为秒，取值范围为[0,604800]。默认值为 0，表示无限期。
	PreSignDuration *int32 `json:"PreSignDuration,omitempty"`

	// 源文件属性，动态转码需要
	ResourceAttr *WbTranscodeCreateBodyResourceAttr `json:"ResourceAttr,omitempty"`

	// 转码结果存储配置，参数优先级高于控制台配置
	StorageConfig *WbTranscodeCreateBodyStorageConfig `json:"StorageConfig,omitempty"`

	// 转码配置，静态转码需要，动态转码需要生成缩略图时需要
	TranscodeConfig *WbTranscodeCreateBodyTranscodeConfig `json:"TranscodeConfig,omitempty"`
}

// WbTranscodeCreateBodyResourceAttr - 源文件属性，动态转码需要
type WbTranscodeCreateBodyResourceAttr struct {

	// REQUIRED; 文件名
	FileName string `json:"FileName"`

	// REQUIRED; 文件大小，单位字节
	Size int32 `json:"Size"`
}

// WbTranscodeCreateBodyStorageConfig - 转码结果存储配置，参数优先级高于控制台配置
type WbTranscodeCreateBodyStorageConfig struct {

	// 第三方对象存储服务参数设置
	CustomConfig *WbTranscodeCreateBodyStorageConfigCustomConfig `json:"CustomConfig,omitempty"`

	// 火山引擎的对象存储服务参数设置
	TosConfig *WbTranscodeCreateBodyStorageConfigTosConfig `json:"TosConfig,omitempty"`

	// 存储平台。0: TOS1: 第三方对象存储接口。 默认值为 0。
	Type *int32 `json:"Type,omitempty"`
}

// WbTranscodeCreateBodyStorageConfigCustomConfig - 第三方对象存储服务参数设置
type WbTranscodeCreateBodyStorageConfigCustomConfig struct {

	// REQUIRED; 第三方存储平台账号的密钥。需确保此账号对存储桶有写权限，不建议开启读权限。
	AccessKey string `json:"AccessKey"`

	// REQUIRED; 桶名称
	Bucket string `json:"Bucket"`

	// REQUIRED; 参看 Region 对照表 [https://www.volcengine.com/docs/6348/1167931]当 vendor =0 时，7、9、15、16、17、23 均不支持。
	Region int32 `json:"Region"`

	// REQUIRED; 第三方存储平台账号的密钥
	SecretKey string `json:"SecretKey"`

	// REQUIRED; 第三方存储供应商0: Amazon 亚马逊1: Alicloud 阿里云
	Vendor int32 `json:"Vendor"`
}

// WbTranscodeCreateBodyStorageConfigTosConfig - 火山引擎的对象存储服务参数设置
type WbTranscodeCreateBodyStorageConfigTosConfig struct {

	// REQUIRED; Bucket 所属的火山引擎账号 ID。在登录火山引擎后，可在头像的悬浮菜单中找到账号 ID。[https://portal.volccdn.com/obj/volcfe/cloud-universal-doc/upload_0819c44c6aadff358a7dfc52c5daab57.png]
	AccountID string `json:"AccountId"`

	// REQUIRED; 桶名称。登录TOS 控制台 [https://console.volcengine.com/tos/bucket]开通和获取。
	Bucket string `json:"Bucket"`

	// REQUIRED; 参看 Region 对照表 [1217553]。
	Region int32 `json:"Region"`
}

// WbTranscodeCreateBodyTranscodeConfig - 转码配置，静态转码需要，动态转码需要生成缩略图时需要
type WbTranscodeCreateBodyTranscodeConfig struct {

	// REQUIRED; 是否强制使用指定分辨率
	ForceUseResolution bool `json:"ForceUseResolution"`

	// REQUIRED; 输入格式
	InputFormat int32 `json:"InputFormat"`

	// REQUIRED; 输出格式
	OutputFormat int32 `json:"OutputFormat"`

	// REQUIRED; 输出分辨率高
	OutputHeight int32 `json:"OutputHeight"`

	// REQUIRED; 输出分辨率宽
	OutputWidth int32 `json:"OutputWidth"`

	// REQUIRED; 是否需要缩略图
	Thumbnail bool `json:"Thumbnail"`

	// REQUIRED; 输出缩略图分辨率高
	ThumbnailHeight int32 `json:"ThumbnailHeight"`

	// REQUIRED; 输出缩略图分辨率宽
	ThumbnailWidth int32 `json:"ThumbnailWidth"`
}

type WbTranscodeCreateRes struct {

	// REQUIRED
	ResponseMetadata WbTranscodeCreateResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *WbTranscodeCreateResResult `json:"Result,omitempty"`
}

type WbTranscodeCreateResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *WbTranscodeCreateResResponseMetadataError `json:"Error,omitempty"`
}

// WbTranscodeCreateResResponseMetadataError - 仅在请求失败时返回。
type WbTranscodeCreateResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

// WbTranscodeCreateResResult - 视请求的接口而定
type WbTranscodeCreateResResult struct {

	// REQUIRED; 任务ID
	TaskID string `json:"TaskId"`
}

type WbTranscodeGetQuery struct {

	// REQUIRED; 应用的唯一标志。你可以通过控制台 [https://console.volcengine.com/rtc/listRTC]查看和复制你的 app_id。或通过调用ListApps [https://www.volcengine.com/docs/6348/74489]接口获取。
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 成功调用WbTranscodeQuery后返回的任务ID
	TaskID string `json:"TaskId" query:"TaskId"`
}

type WbTranscodeGetRes struct {

	// REQUIRED
	ResponseMetadata WbTranscodeGetResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result WbTranscodeGetResResult `json:"Result"`
}

type WbTranscodeGetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *WbTranscodeGetResResponseMetadataError `json:"Error,omitempty"`
}

// WbTranscodeGetResResponseMetadataError - 仅在请求失败时返回。
type WbTranscodeGetResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type WbTranscodeGetResResult struct {

	// REQUIRED; 文件名
	FileName string `json:"FileName"`

	// REQUIRED; 分辨率高
	Height int32 `json:"Height"`

	// REQUIRED; 转码结果
	Images []WbTranscodeGetResResultImagesItem `json:"Images"`

	// REQUIRED; 转码模式
	TranscodeMode int32 `json:"TranscodeMode"`

	// REQUIRED; 分辨率宽
	Width int32 `json:"Width"`

	// 动态转码结果FileId，使用动态转码结果时在SDK接口传入
	FileID *string `json:"FileId,omitempty"`

	// 缩略图分辨率高
	ThumbnailHeight *int32 `json:"ThumbnailHeight,omitempty"`

	// 缩略图分辨率宽
	ThumbnailWidth *int32 `json:"ThumbnailWidth,omitempty"`
}

type WbTranscodeGetResResultImagesItem struct {

	// REQUIRED; 转码图片URL
	ImgURL string `json:"ImgUrl"`

	// REQUIRED; 页码
	PageID float32 `json:"PageId"`

	// REQUIRED; 缩略图URL
	ThumbnailURL string `json:"ThumbnailUrl"`

	// REQUIRED; 如果该页异常，会转为空白图片，并在此给出提示信息
	Warning string `json:"Warning"`
}

type WbTranscodeQueryQuery struct {

	// REQUIRED; yourappid
	AppID string `json:"AppId" query:"AppId"`

	// REQUIRED; 成功调用WbTranscodeQuery后返回的任务ID
	TaskID string `json:"TaskId" query:"TaskId"`
}

type WbTranscodeQueryRes struct {

	// REQUIRED
	ResponseMetadata WbTranscodeQueryResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result WbTranscodeQueryResResult `json:"Result"`
}

type WbTranscodeQueryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`

	// 仅在请求失败时返回。
	Error *WbTranscodeQueryResResponseMetadataError `json:"Error,omitempty"`
}

// WbTranscodeQueryResResponseMetadataError - 仅在请求失败时返回。
type WbTranscodeQueryResResponseMetadataError struct {

	// REQUIRED; API 的错误码
	Code string `json:"Code"`

	// REQUIRED; 具体的错误信息
	Message string `json:"Message"`

	// 网关的错误码。（请求失败时返回）
	CodeN *int32 `json:"CodeN,omitempty"`
}

type WbTranscodeQueryResResult struct {

	// REQUIRED; 转码进度
	Progress int32 `json:"Progress"`

	// REQUIRED; 转码任务状态
	Status int32 `json:"Status"`

	// 错误码描述
	ErrCode *string `json:"ErrCode,omitempty"`

	// 错误码
	ErrCodeN *int32 `json:"ErrCodeN,omitempty"`

	// 错误信息
	ErrMsg *string `json:"ErrMsg,omitempty"`
}
type BanRoomUser struct{}
type BanRoomUserQuery struct{}
type BanUserStream struct{}
type BanUserStreamQuery struct{}
type CreateApp struct{}
type CreateAppQuery struct{}
type GetPushMixedStreamToCDNTask struct{}
type GetPushMixedStreamToCDNTaskBody struct{}
type GetPushSingleStreamToCDNTask struct{}
type GetPushSingleStreamToCDNTaskBody struct{}
type GetRecordTask struct{}
type GetRecordTaskBody struct{}
type GetRoomOnlineUsers struct{}
type GetRoomOnlineUsersBody struct{}
type GetSegmentTask struct{}
type GetSegmentTaskBody struct{}
type GetSnapshotTask struct{}
type GetSnapshotTaskBody struct{}
type GetWebCastTask struct{}
type GetWebCastTaskBody struct{}
type LimitTokenPrivilege struct{}
type LimitTokenPrivilegeQuery struct{}
type ListApps struct{}
type ListAppsQuery struct{}
type ListCallDetail struct{}
type ListCallDetailBody struct{}
type ListDetectionTask struct{}
type ListDetectionTaskBody struct{}
type ListOperationData struct{}
type ListOperationDataQuery struct{}
type ListOperationDistribution struct{}
type ListOperationDistributionQuery struct{}
type ListQuality struct{}
type ListQualityDistribution struct{}
type ListQualityDistributionQuery struct{}
type ListQualityQuery struct{}
type ListRealTimeOperationData struct{}
type ListRealTimeOperationDataQuery struct{}
type ListRealTimeQuality struct{}
type ListRealTimeQualityDistribution struct{}
type ListRealTimeQualityDistributionQuery struct{}
type ListRealTimeQualityQuery struct{}
type ListRelayStream struct{}
type ListRelayStreamBody struct{}
type ListRoomInfo struct{}
type ListRoomInfoBody struct{}
type ListUsages struct{}
type ListUsagesQuery struct{}
type ListUserInfo struct{}
type ListUserInfoBody struct{}
type ModifyAppStatus struct{}
type ModifyAppStatusQuery struct{}
type StartDetection struct{}
type StartDetectionQuery struct{}
type StartPushMixedStreamToCDN struct{}
type StartPushMixedStreamToCDNQuery struct{}
type StartPushPublicStream struct{}
type StartPushPublicStreamQuery struct{}
type StartPushSingleStreamToCDN struct{}
type StartPushSingleStreamToCDNQuery struct{}
type StartRecord struct{}
type StartRecordQuery struct{}
type StartRelayStream struct{}
type StartRelayStreamQuery struct{}
type StartSegment struct{}
type StartSegmentQuery struct{}
type StartSnapshot struct{}
type StartSnapshotQuery struct{}
type StartWBRecord struct{}
type StartWBRecordQuery struct{}
type StartWebcast struct{}
type StartWebcastQuery struct{}
type StopDetection struct{}
type StopDetectionQuery struct{}
type StopPushPublicStream struct{}
type StopPushPublicStreamQuery struct{}
type StopPushStreamToCDN struct{}
type StopPushStreamToCDNQuery struct{}
type StopRecord struct{}
type StopRecordQuery struct{}
type StopRelayStream struct{}
type StopRelayStreamQuery struct{}
type StopSegment struct{}
type StopSegmentQuery struct{}
type StopSnapshot struct{}
type StopSnapshotQuery struct{}
type StopWBRecord struct{}
type StopWBRecordQuery struct{}
type StopWebcast struct{}
type StopWebcastQuery struct{}
type UnbanUserStream struct{}
type UnbanUserStreamQuery struct{}
type UpdateBanRoomUserRule struct{}
type UpdateBanRoomUserRuleQuery struct{}
type UpdatePublicStreamParam struct{}
type UpdatePublicStreamParamQuery struct{}
type UpdatePushMixedStreamToCDN struct{}
type UpdatePushMixedStreamToCDNQuery struct{}
type UpdateRecord struct{}
type UpdateRecordQuery struct{}
type UpdateRelayStream struct{}
type UpdateRelayStreamQuery struct{}
type UpdateSegment struct{}
type UpdateSegmentQuery struct{}
type UpdateSnapshot struct{}
type UpdateSnapshotQuery struct{}
type WbTranscodeCreate struct{}
type WbTranscodeCreateQuery struct{}
type WbTranscodeGet struct{}
type WbTranscodeGetBody struct{}
type WbTranscodeQueryBody struct{}
type BanRoomUserReq struct {
	*BanRoomUserQuery
	*BanRoomUserBody
}
type BanUserStreamReq struct {
	*BanUserStreamQuery
	*BanUserStreamBody
}
type CreateAppReq struct {
	*CreateAppQuery
	*CreateAppBody
}
type GetPushMixedStreamToCDNTaskReq struct {
	*GetPushMixedStreamToCDNTaskQuery
	*GetPushMixedStreamToCDNTaskBody
}
type GetPushSingleStreamToCDNTaskReq struct {
	*GetPushSingleStreamToCDNTaskQuery
	*GetPushSingleStreamToCDNTaskBody
}
type GetRecordTaskReq struct {
	*GetRecordTaskQuery
	*GetRecordTaskBody
}
type GetRoomOnlineUsersReq struct {
	*GetRoomOnlineUsersQuery
	*GetRoomOnlineUsersBody
}
type GetSegmentTaskReq struct {
	*GetSegmentTaskQuery
	*GetSegmentTaskBody
}
type GetSnapshotTaskReq struct {
	*GetSnapshotTaskQuery
	*GetSnapshotTaskBody
}
type GetWebCastTaskReq struct {
	*GetWebCastTaskQuery
	*GetWebCastTaskBody
}
type LimitTokenPrivilegeReq struct {
	*LimitTokenPrivilegeQuery
	*LimitTokenPrivilegeBody
}
type ListAppsReq struct {
	*ListAppsQuery
	*ListAppsBody
}
type ListCallDetailReq struct {
	*ListCallDetailQuery
	*ListCallDetailBody
}
type ListDetectionTaskReq struct {
	*ListDetectionTaskQuery
	*ListDetectionTaskBody
}
type ListOperationDataReq struct {
	*ListOperationDataQuery
	*ListOperationDataBody
}
type ListOperationDistributionReq struct {
	*ListOperationDistributionQuery
	*ListOperationDistributionBody
}
type ListQualityReq struct {
	*ListQualityQuery
	*ListQualityBody
}
type ListQualityDistributionReq struct {
	*ListQualityDistributionQuery
	*ListQualityDistributionBody
}
type ListRealTimeOperationDataReq struct {
	*ListRealTimeOperationDataQuery
	*ListRealTimeOperationDataBody
}
type ListRealTimeQualityReq struct {
	*ListRealTimeQualityQuery
	*ListRealTimeQualityBody
}
type ListRealTimeQualityDistributionReq struct {
	*ListRealTimeQualityDistributionQuery
	*ListRealTimeQualityDistributionBody
}
type ListRelayStreamReq struct {
	*ListRelayStreamQuery
	*ListRelayStreamBody
}
type ListRoomInfoReq struct {
	*ListRoomInfoQuery
	*ListRoomInfoBody
}
type ListUsagesReq struct {
	*ListUsagesQuery
	*ListUsagesBody
}
type ListUserInfoReq struct {
	*ListUserInfoQuery
	*ListUserInfoBody
}
type ModifyAppStatusReq struct {
	*ModifyAppStatusQuery
	*ModifyAppStatusBody
}
type StartDetectionReq struct {
	*StartDetectionQuery
	*StartDetectionBody
}
type StartPushMixedStreamToCDNReq struct {
	*StartPushMixedStreamToCDNQuery
	*StartPushMixedStreamToCDNBody
}
type StartPushPublicStreamReq struct {
	*StartPushPublicStreamQuery
	*StartPushPublicStreamBody
}
type StartPushSingleStreamToCDNReq struct {
	*StartPushSingleStreamToCDNQuery
	*StartPushSingleStreamToCDNBody
}
type StartRecordReq struct {
	*StartRecordQuery
	*StartRecordBody
}
type StartRelayStreamReq struct {
	*StartRelayStreamQuery
	*StartRelayStreamBody
}
type StartSegmentReq struct {
	*StartSegmentQuery
	*StartSegmentBody
}
type StartSnapshotReq struct {
	*StartSnapshotQuery
	*StartSnapshotBody
}
type StartWBRecordReq struct {
	*StartWBRecordQuery
	*StartWBRecordBody
}
type StartWebcastReq struct {
	*StartWebcastQuery
	*StartWebcastBody
}
type StopDetectionReq struct {
	*StopDetectionQuery
	*StopDetectionBody
}
type StopPushPublicStreamReq struct {
	*StopPushPublicStreamQuery
	*StopPushPublicStreamBody
}
type StopPushStreamToCDNReq struct {
	*StopPushStreamToCDNQuery
	*StopPushStreamToCDNBody
}
type StopRecordReq struct {
	*StopRecordQuery
	*StopRecordBody
}
type StopRelayStreamReq struct {
	*StopRelayStreamQuery
	*StopRelayStreamBody
}
type StopSegmentReq struct {
	*StopSegmentQuery
	*StopSegmentBody
}
type StopSnapshotReq struct {
	*StopSnapshotQuery
	*StopSnapshotBody
}
type StopWBRecordReq struct {
	*StopWBRecordQuery
	*StopWBRecordBody
}
type StopWebcastReq struct {
	*StopWebcastQuery
	*StopWebcastBody
}
type UnbanUserStreamReq struct {
	*UnbanUserStreamQuery
	*UnbanUserStreamBody
}
type UpdateBanRoomUserRuleReq struct {
	*UpdateBanRoomUserRuleQuery
	*UpdateBanRoomUserRuleBody
}
type UpdatePublicStreamParamReq struct {
	*UpdatePublicStreamParamQuery
	*UpdatePublicStreamParamBody
}
type UpdatePushMixedStreamToCDNReq struct {
	*UpdatePushMixedStreamToCDNQuery
	*UpdatePushMixedStreamToCDNBody
}
type UpdateRecordReq struct {
	*UpdateRecordQuery
	*UpdateRecordBody
}
type UpdateRelayStreamReq struct {
	*UpdateRelayStreamQuery
	*UpdateRelayStreamBody
}
type UpdateSegmentReq struct {
	*UpdateSegmentQuery
	*UpdateSegmentBody
}
type UpdateSnapshotReq struct {
	*UpdateSnapshotQuery
	*UpdateSnapshotBody
}
type WbTranscodeCreateReq struct {
	*WbTranscodeCreateQuery
	*WbTranscodeCreateBody
}
type WbTranscodeGetReq struct {
	*WbTranscodeGetQuery
	*WbTranscodeGetBody
}
type WbTranscodeQueryReq struct {
	*WbTranscodeQueryQuery
	*WbTranscodeQueryBody
}
