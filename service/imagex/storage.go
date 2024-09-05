package imagex

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/volcengine/volc-sdk-golang/base"
)

func (c *ImageX) ImageXGet(action string, query url.Values, result interface{}) error {
	respBody, _, err := c.Client.Query(action, query)
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

func (c *ImageX) ImageXPost(action string, query url.Values, req, result interface{}) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("%s: fail to marshal request, %v", action, err)
	}
	data, _, err := c.Client.Json(action, query, string(body))
	if err != nil {
		return fmt.Errorf("%s: fail to do request, %v", action, err)
	}
	if err := UnmarshalResultInto(data, result); err != nil {
		return fmt.Errorf("%s: fail to unmarshal response, %v", action, err)
	}
	return nil
}

// GetAllImageServices 获取所有服务信息
func (c *ImageX) GetImageServices(searchPtn string) (*GetServicesResult, error) {
	query := url.Values{}
	if searchPtn != "" {
		query.Add("SearchPtn", searchPtn)
	}

	respBody, _, err := c.Client.Query("GetAllImageServices", query)
	if err != nil {
		return nil, fmt.Errorf("GetAllImageServices: fail to do request, %v", err)
	}

	result := new(GetServicesResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, fmt.Errorf("GetAllImageServices: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// GetServiceDomains 获取服务下的所有域名
func (c *ImageX) GetImageDomains(serviceId string) ([]DomainResult, error) {
	query := url.Values{}
	query.Add("ServiceId", serviceId)

	respBody, _, err := c.Client.Query("GetServiceDomains", query)
	if err != nil {
		return nil, fmt.Errorf("GetServiceDomains: fail to do request, %v", err)
	}

	result := make([]DomainResult, 0)
	if err := UnmarshalResultInto(respBody, &result); err != nil {
		return nil, fmt.Errorf("GetServiceDomains: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// DeleteImageUploadFiles 删除图片
func (c *ImageX) DeleteImages(serviceId string, uris []string) (*DeleteImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", serviceId)
	param := new(DeleteImageParam)
	param.StoreUris = uris

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("DeleteImages: fail to marshal request, %v", err)
	}

	data, _, err := c.Client.Json("DeleteImageUploadFiles", query, string(body))
	if err != nil {
		return nil, fmt.Errorf("DeleteImages: fail to do request, %v", err)
	}

	result := new(DeleteImageResult)
	if err := UnmarshalResultInto(data, result); err != nil {
		return nil, fmt.Errorf("DeleteImages: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// ApplyImageUpload 获取图片上传地址
func (c *ImageX) ApplyUploadImage(params *ApplyUploadImageParam) (*ApplyUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	if params.SessionKey != "" {
		query.Add("SessionKey", params.SessionKey)
	}
	if params.UploadNum > 0 {
		query.Add("UploadNum", strconv.Itoa(params.UploadNum))
	}
	for _, key := range params.StoreKeys {
		query.Add("StoreKeys", key)
	}
	query.Add("Prefix", params.Prefix)
	query.Add("FileExtension", params.FileExtension)
	query.Add("Overwrite", strconv.FormatBool(params.Overwrite))

	respBody, _, err := c.Client.Query("ApplyImageUpload", query)
	if err != nil {
		return nil, fmt.Errorf("ApplyUploadImage: fail to do request, %v", err)
	}

	result := new(ApplyUploadImageResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, fmt.Errorf("ApplyUploadImage: fail to unmarshal response, %v", err)
	}
	return result, nil
}

// CommitImageUpload 图片上传完成上报
func (c *ImageX) CommitUploadImage(params *CommitUploadImageParam) (*CommitUploadImageResult, error) {
	query := url.Values{}
	query.Add("ServiceId", params.ServiceId)
	query.Add("SkipMeta", fmt.Sprintf("%v", params.SkipMeta))

	bts, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("CommitUploadImage: fail to marshal request, %v", err)
	}

	respBody, _, err := c.Client.Json("CommitImageUpload", query, string(bts))
	if err != nil {
		return nil, fmt.Errorf("CommitUploadImage: fail to do request, %v", err)
	}

	result := new(CommitUploadImageResult)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) segmentedUpload(set *uploadTaskSet, item *uploadTaskElement) error {
	if item.size <= MinChunkSize {
		//goland:noinspection GoDeprecation
		bts, err := ioutil.ReadAll(item.content)
		if err != nil {
			return err
		}
		err = c.directUpload(item.ctx, item.host, item.idx, set, item.info, bts, item.ct)
		if err != nil {
			return err
		}
	} else {
		arg := &segmentedUploadParam{
			host:        item.host,
			StoreInfo:   item.info,
			content:     item.content,
			size:        item.size,
			isLargeFile: item.size > LargeFileSize,
			idx:         item.idx,
			set:         set,
			ct:          item.ct,
		}
		err := arg.chunkUpload()
		if err != nil {
			return err
		}
	}
	set.result[item.idx].success = true
	return nil
}

// 上传图片
// 请确保 content 长度和 size 长度一致
func (c *ImageX) SegmentedUploadImages(ctx context.Context, params *ApplyUploadImageParam, content []io.Reader, size []int64) (*CommitUploadImageResult, error) {
	if len(content) != len(size) {
		return nil, fmt.Errorf("expect len(size) == len(content), but len(size) = %d, len(content) = %d", len(size), len(content))
	}

	params.UploadNum = len(content)

	for idx, item := range size {
		if item == 0 {
			return nil, fmt.Errorf("size[%d] is zero", idx)
		}
	}

	// 1. apply
	applyResp, err := c.ApplyUploadImage(params)
	if err != nil {
		return nil, err
	}

	uploadAddr := applyResp.UploadAddress
	if len(uploadAddr.UploadHosts) == 0 {
		return nil, fmt.Errorf("UploadImages: no upload host found, request id %s", applyResp.RequestId)
	}
	if len(uploadAddr.StoreInfos) != params.UploadNum {
		return nil, fmt.Errorf("UploadImages: store infos num %d != upload num %d, request id %s",
			len(uploadAddr.StoreInfos), params.UploadNum, applyResp.RequestId)
	}

	// 2. upload
	host := uploadAddr.UploadHosts[0]

	wg := &sync.WaitGroup{}
	uploadTaskSet := &uploadTaskSet{
		ctx:     ctx,
		host:    host,
		info:    uploadAddr.StoreInfos,
		content: content,
		size:    size,
		cts:     params.ContentTypes,
	}
	uploadTaskSet.init()

	for i := 0; i < UploadRoutines; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				recover()
				wg.Done()
			}()

			for item := range uploadTaskSet.taskChan {
				uploadTaskSet.result[item.idx] = uploadTaskResult{
					uri: item.info.StoreUri,
				}
				err := c.segmentedUpload(uploadTaskSet, item)
				if err == nil {
					uploadTaskSet.result[item.idx].success = true
					uploadTaskSet.addSuccess(item.info.StoreUri)
				} else {
					uploadTaskSet.result[item.idx].errMsg = err.Error()
				}
			}
		}()
	}
	wg.Wait()
	if len(uploadTaskSet.successOids) == 0 || params.SkipCommit {
		return &CommitUploadImageResult{
			Results: uploadTaskSet.GetResult(),
		}, nil
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:   params.ServiceId,
		SessionKey:  uploadAddr.SessionKey,
		SuccessOids: uploadTaskSet.successOids,
		SkipMeta:    params.SkipMeta,
	}
	if params.CommitParam != nil {
		commitParams.Functions = params.CommitParam.Functions
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	uploadTaskSet.fill(commitResp)
	return commitResp, nil
}

// 上传图片
func (c *ImageX) UploadImages(params *ApplyUploadImageParam, images [][]byte) (*CommitUploadImageResult, error) {
	params.UploadNum = len(images)

	// 1. apply
	applyResp, err := c.ApplyUploadImage(params)
	if err != nil {
		return nil, err
	}
	uploadAddr := applyResp.UploadAddress
	if len(uploadAddr.UploadHosts) == 0 {
		return nil, fmt.Errorf("UploadImages: no upload host found, request id %s", applyResp.RequestId)
	} else if len(uploadAddr.StoreInfos) != params.UploadNum {
		return nil, fmt.Errorf("UploadImages: store infos num %d != upload num %d, request id %s",
			len(uploadAddr.StoreInfos), params.UploadNum, applyResp.RequestId)
	}

	// 2. upload
	host := uploadAddr.UploadHosts[0]
	uploadTaskSet := &uploadTaskSet{
		ctx:     context.Background(),
		host:    host,
		info:    uploadAddr.StoreInfos,
		content: make([]io.Reader, 0),
		size:    make([]int64, 0),
	}
	uploadTaskSet.init()
	for i, image := range images {
		imageCopy := image
		info := uploadAddr.StoreInfos[i]
		idx := i
		uploadTaskSet.result[idx] = uploadTaskResult{
			uri: info.StoreUri,
		}
		ct := ""
		if idx < len(params.ContentTypes) {
			ct = params.ContentTypes[idx]
		}
		err = retry.Do(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), c.ServiceInfo.Timeout)
			defer cancel()
			return c.directUpload(ctx, host, idx, uploadTaskSet, info, imageCopy, ct)
		}, retry.Attempts(2))
		if err != nil {
			uploadTaskSet.result[idx].errMsg = err.Error()
		} else {
			uploadTaskSet.result[idx].success = true
			uploadTaskSet.result[idx].errMsg = ""
			uploadTaskSet.result[idx].putErr = nil
			uploadTaskSet.addSuccess(info.StoreUri)
		}
	}
	if len(uploadTaskSet.successOids) == 0 || params.SkipCommit {
		return &CommitUploadImageResult{
			Results: uploadTaskSet.GetResult(),
		}, nil
	}

	// 3. commit
	commitParams := &CommitUploadImageParam{
		ServiceId:   params.ServiceId,
		SessionKey:  uploadAddr.SessionKey,
		SuccessOids: uploadTaskSet.successOids,
		SkipMeta:    params.SkipMeta,
	}
	if params.CommitParam != nil {
		commitParams.Functions = params.CommitParam.Functions
	}
	commitResp, err := c.CommitUploadImage(commitParams)
	if err != nil {
		return nil, err
	}
	uploadTaskSet.fill(commitResp)
	return commitResp, nil
}

// 获取临时上传凭证
func (c *ImageX) GetUploadAuthToken(query url.Values) (string, error) {
	ret := map[string]string{
		"Version": "v1",
	}

	applyUploadToken, err := c.Client.GetSignUrl("ApplyImageUpload", query)
	if err != nil {
		return "", fmt.Errorf("GetUploadAuthToken: fail to get apply token, %v", err)
	}
	ret["ApplyUploadToken"] = applyUploadToken

	commitUploadToken, err := c.Client.GetSignUrl("CommitImageUpload", query)
	if err != nil {
		return "", fmt.Errorf("GetUploadAuthToken: fail to get commit token, %v", err)
	}
	ret["CommitUploadToken"] = commitUploadToken

	b, err := json.Marshal(ret)
	if err != nil {
		return "", fmt.Errorf("GetUploadAuthToken: fail to marshal token, %v", err)
	}
	token := strings.ReplaceAll(string(b), "\\u0026", "&")
	return base64.StdEncoding.EncodeToString([]byte(token)), nil
}

type UploadAuthOpt func(option *uploadAuthOption)

type uploadAuthOption struct {
	keyPtn     string
	conditions map[string]string
}

func WithUploadKeyPtn(ptn string) UploadAuthOpt {
	return func(o *uploadAuthOption) {
		o.keyPtn = ptn
	}
}

func WithUploadOverwrite(overwrite bool) UploadAuthOpt {
	return func(op *uploadAuthOption) {
		op.conditions["UploadOverwrite"] = strconv.FormatBool(overwrite)
	}
}

func WithUploadPolicy(policy *UploadPolicy) UploadAuthOpt {
	res, _ := json.Marshal(policy)
	return func(op *uploadAuthOption) {
		op.conditions["UploadPolicy"] = string(res)
	}
}

// 获取上传临时密钥
func (c *ImageX) GetUploadAuth(serviceIds []string, opt ...UploadAuthOpt) (*base.SecurityToken2, error) {
	return c.GetUploadAuthWithExpire(serviceIds, time.Hour, opt...)
}

// 获取上传临时密钥
func (c *ImageX) GetUploadAuthWithExpire(serviceIds []string, expire time.Duration, opt ...UploadAuthOpt) (*base.SecurityToken2, error) {
	serviceIdRes := make([]string, 0)
	if len(serviceIds) == 0 {
		serviceIdRes = append(serviceIdRes, fmt.Sprintf(ResourceServiceIdTRN, "*"))
	} else {
		for _, sid := range serviceIds {
			serviceIdRes = append(serviceIdRes, fmt.Sprintf(ResourceServiceIdTRN, sid))
		}
	}
	op := &uploadAuthOption{
		conditions: map[string]string{},
	}
	for _, o := range opt {
		o(op)
	}
	storeKeyRes := []string{
		fmt.Sprintf(ResourceStoreKeyTRN, op.keyPtn),
	}

	inlinePolicy := new(base.Policy)
	inlinePolicy.Statement = append(inlinePolicy.Statement, base.NewAllowStatement([]string{
		"ImageX:ApplyImageUpload",
	}, append(serviceIdRes, storeKeyRes...)))
	inlinePolicy.Statement = append(inlinePolicy.Statement, base.NewAllowStatement([]string{
		"ImageX:CommitImageUpload",
	}, serviceIdRes))
	for k, v := range op.conditions {
		inlinePolicy.Statement = append(inlinePolicy.Statement, base.NewAllowStatement([]string{k}, []string{v}))
	}

	return c.Client.SignSts2(inlinePolicy, expire)
}

func (c *ImageX) CreateImageContentTask(req *CreateImageContentTaskReq) (*CreateImageContentTaskResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &CreateImageContentTaskResp{}
	err = c.ImageXPost("CreateImageContentTask", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageContentTaskDetail(req *GetImageContentTaskDetailReq) (*GetImageContentTaskDetailResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageContentTaskDetailResp{}
	err = c.ImageXPost("GetImageContentTaskDetail", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) GetImageContentBlockList(req *GetImageContentBlockListReq) (*GetImageContentBlockListResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := &GetImageContentBlockListResp{}
	err = c.ImageXPost("GetImageContentBlockList", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *ImageX) FetchImageUrl(req *FetchUrlReq) (*FetchUrlResp, error) {
	resp := new(FetchUrlResp)
	err := c.ImageXPost("FetchImageUrl", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetUrlFetchTask(req *GetUrlFetchTaskReq) (*GetUrlFetchTaskResp, error) {
	query, err := MarshalToQuery(req, SkipEmptyValue())
	if err != nil {
		return nil, err
	}

	resp := new(GetUrlFetchTaskResp)
	err = c.ImageXGet("GetUrlFetchTask", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetImageUploadFile(param *GetImageUploadFileParam) (*GetImageUploadFileResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)
	query.Add("StoreUri", param.StoreUri)

	result := new(GetImageUploadFileResult)
	err := c.ImageXGet("GetImageUploadFile", query, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) GetImageUploadFiles(param *GetImageUploadFilesParam) (*GetImageUploadFilesResult, error) {
	query := make(url.Values)
	query.Add("ServiceId", param.ServiceId)
	query.Add("Limit", strconv.Itoa(param.Limit))
	query.Add("Marker", strconv.FormatInt(param.Marker, 10))

	result := new(GetImageUploadFilesResult)
	err := c.ImageXGet("GetImageUploadFiles", query, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ImageX) UpdateImageStorageTTL(req *UpdateImageStorageTTLReq) (*UpdateImageStorageTTLResp, error) {
	resp := new(UpdateImageStorageTTLResp)
	err := c.ImageXPost("UpdateImageStorageTTL", nil, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) CreateImageCompressTask(req *CreateImageCompressTaskReq) (*CreateImageCompressTaskResp, error) {
	query := make(url.Values)
	query.Add("ServiceId", req.ServiceId)
	resp := new(CreateImageCompressTaskResp)
	err := c.ImageXPost("CreateImageCompressTask", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetImageCompressTaskInfo(req *GetImageCompressTaskInfoReq) (*GetImageCompressTaskInfoResp, error) {
	query := make(url.Values)
	query.Add("ServiceId", req.ServiceId)
	query.Add("TaskId", req.TaskId)
	resp := new(GetImageCompressTaskInfoResp)
	err := c.ImageXGet("GetCompressTaskInfo", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetAiGenerateImage(req *GetAiGenerateImageReq) (*GetAiGenerateImageResp, error) {
	query := make(url.Values)
	query.Add("ServiceId", req.ServiceId)
	resp := new(GetAiGenerateImageResp)
	err := c.ImageXPost("GetAiGenerateImage", query, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ImageX) GetImageAiGenerateTask(req *GetImageAiGenerateTaskReq) (*GetImageAiGenerateTaskResp, error) {
	query := make(url.Values)
	query.Add("TaskId", req.TaskId)
	resp := new(GetImageAiGenerateTaskResp)
	err := c.ImageXGet("GetImageAiGenerateTask", query, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
