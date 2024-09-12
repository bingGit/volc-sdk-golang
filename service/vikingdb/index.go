package vikingdb

import (
	"context"
	"fmt"
)

type Index struct {
	CollectionName  string
	IndexName       string
	Description     string
	VectorIndex     *VectorIndexParams
	ScalarIndex     interface{}
	Stat            string
	VikingDBService *VikingDBService
	CpuQuota        int64
	PartitionBy     string
	CreateTime      string
	UpdateTime      string
	UpdatePerson    string
	IndexCost       interface{}
	ShardCount      int64
	ShardPolicy     string
	primaryKey      string
}

type SearchOptions struct {
	filter         map[string]interface{}
	limit          int64
	outputFields   []string
	partition      string
	dense_weight   *float64
	sparse_vectors []map[string]interface{}
}

func NewSearchOptions() *SearchOptions {
	searchOptions := &SearchOptions{
		filter:         nil,
		limit:          10,
		outputFields:   nil,
		partition:      "default",
		dense_weight:   nil,
		sparse_vectors: nil,
	}
	return searchOptions
}
func (searchOptions *SearchOptions) SetFilter(filter map[string]interface{}) *SearchOptions {
	searchOptions.filter = filter
	return searchOptions
}
func (searchOptions *SearchOptions) SetLimit(limit int64) *SearchOptions {
	searchOptions.limit = limit
	return searchOptions
}
func (searchOptions *SearchOptions) SetOutputFields(outputFields []string) *SearchOptions {
	searchOptions.outputFields = outputFields
	return searchOptions
}
func (searchOptions *SearchOptions) SetPartition(partition string) *SearchOptions {
	searchOptions.partition = partition
	return searchOptions
}
func (searchOptions *SearchOptions) SetDenseWeight(dense_weight float64) *SearchOptions {
	searchOptions.dense_weight = &dense_weight
	return searchOptions
}
func (searchOptions *SearchOptions) SetSparseVectors(sparse_vectors map[string]interface{}) *SearchOptions {
	searchOptions.sparse_vectors = []map[string]interface{}{sparse_vectors}
	return searchOptions
}

func (index *Index) Search(order interface{}, searchOptions *SearchOptions) ([]*Data, error) {
	if _, ok := order.(VectorOrder); ok {
		order := order.(VectorOrder)
		if order.Vector != nil {
			return index.SearchByVector(order.Vector.([]float64), searchOptions)

		} else if order.Id != nil {
			return index.SearchById(order.Id, searchOptions)
		}
	} else if _, ok := order.(ScalarOrder); ok {
		order := order.(ScalarOrder)
		orderByScalar := map[string]interface{}{
			"order":      order.Order,
			"field_name": order.FieldName,
		}
		search := map[string]interface{}{
			"order_by_scalar": orderByScalar,
			"limit":           searchOptions.limit,
			"partition":       searchOptions.partition,
		}
		if searchOptions.filter != nil {
			search["filter"] = searchOptions.filter
		}
		var outputFields interface{} = nil
		if searchOptions.outputFields != nil {
			search["output_fields"] = searchOptions.outputFields
			outputFields = searchOptions.outputFields
		}
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"search":          search,
		}
		res, err := index.VikingDBService.DoRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params))
		if err != nil {
			return nil, err
		}
		return index.getData(res, outputFields)
	} else if order == nil {
		search := map[string]interface{}{
			"limit":     searchOptions.limit,
			"partition": searchOptions.partition,
		}
		if searchOptions.filter != nil {
			search["filter"] = searchOptions.filter
		}
		var outputFields interface{} = nil
		if searchOptions.outputFields != nil {
			search["output_fields"] = searchOptions.outputFields
			outputFields = searchOptions.outputFields
		}
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"search":          search,
		}
		res, err := index.VikingDBService.DoRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params))
		if err != nil {
			return nil, err
		}
		return index.getData(res, outputFields)
	}
	return nil, nil
}
func (index *Index) SearchById(id interface{}, searchOptions *SearchOptions) ([]*Data, error) {
	orderById := map[string]interface{}{"primary_keys": id}
	search := map[string]interface{}{
		"order_by_vector": orderById,
		"limit":           searchOptions.limit,
		"partition":       searchOptions.partition,
	}
	if searchOptions.filter != nil {
		search["filter"] = searchOptions.filter
	}
	if searchOptions.dense_weight != nil {
		search["dense_weight"] = *searchOptions.dense_weight
	}
	var outputFields interface{} = nil
	if searchOptions.outputFields != nil {
		search["output_fields"] = searchOptions.outputFields
		outputFields = searchOptions.outputFields
	}
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"search":          search,
	}
	res, err := index.VikingDBService.DoRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	return index.getData(res, outputFields)
}
func (index *Index) SearchByVector(vector []float64, searchOptions *SearchOptions) ([]*Data, error) {
	orderByVector := map[string]interface{}{"vectors": []interface{}{vector}}
	if searchOptions.sparse_vectors != nil {
		orderByVector["sparse_vectors"] = searchOptions.sparse_vectors
	}
	search := map[string]interface{}{
		"order_by_vector": orderByVector,
		"limit":           searchOptions.limit,
		"partition":       searchOptions.partition,
	}
	if searchOptions.filter != nil {
		search["filter"] = searchOptions.filter
	}
	if searchOptions.dense_weight != nil {
		search["dense_weight"] = *searchOptions.dense_weight
	}
	var outputFields interface{} = nil
	if searchOptions.outputFields != nil {
		search["output_fields"] = searchOptions.outputFields
		outputFields = searchOptions.outputFields
	}
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"search":          search,
	}
	//fmt.Println(params)
	res, err := index.VikingDBService.DoRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%#v\n", res)
	return index.getData(res, outputFields)
}
func (index *Index) SearchByText(text TextObject, searchOptions *SearchOptions) ([]*Data, error) {
	orderByRaw := map[string]interface{}{"text": text.Text}
	search := map[string]interface{}{
		"order_by_raw": orderByRaw,
		"limit":        searchOptions.limit,
		"partition":    searchOptions.partition,
	}
	if searchOptions.filter != nil {
		search["filter"] = searchOptions.filter
	}
	if searchOptions.dense_weight != nil {
		search["dense_weight"] = *searchOptions.dense_weight
	}
	var outputFields interface{} = nil
	if searchOptions.outputFields != nil {
		search["output_fields"] = searchOptions.outputFields
		outputFields = searchOptions.outputFields
	}
	params := map[string]interface{}{
		"collection_name": index.CollectionName,
		"index_name":      index.IndexName,
		"search":          search,
	}
	res, err := index.VikingDBService.DoRequest(context.Background(), "SearchIndex", nil, index.VikingDBService.convertMapToJson(params))
	if err != nil {
		return nil, err
	}
	//fmt.Println(res)
	return index.getData(res, outputFields)
}
func (index *Index) FetchData(id interface{}, searchOptions *SearchOptions) ([]*Data, error) {
	_, intType := id.(int)
	_, stringType := id.(string)
	datas := []*Data{}
	if intType || stringType {
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"primary_keys":    id,
			"partition":       searchOptions.partition,
		}
		if searchOptions.outputFields != nil {
			params["output_fields"] = searchOptions.outputFields
		}
		res, err := index.VikingDBService.DoRequest(context.Background(), "FetchIndexData", nil, index.VikingDBService.convertMapToJson(params))
		if err != nil {
			return nil, err
		}
		//fmt.Println(res)
		var resData []interface{}
		var fields map[string]interface{}
		if d, ok := res["data"]; !ok {
			return nil, fmt.Errorf("invalid response, data does not exist: %v", res)
		} else if resData, ok = d.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not a list: %v", res)
		} else if fields, ok = resData[0].(map[string]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not list[map]: %v", res)
		}
		fmt.Println(res)
		data := &Data{
			Id: id,
			Fields: map[string]interface{}{
				"message": "no data found",
			},
			Timestamp: nil,
		}
		if _, exists := fields["fields"]; exists {
			var dataFields map[string]interface{}
			var ok bool
			if dataFields, ok = fields["fields"].(map[string]interface{}); !ok {
				return nil, fmt.Errorf("invalid response, fields is not a map: %v", res)
			}
			data = &Data{
				Id:        id,
				Fields:    dataFields,
				Timestamp: nil,
			}
		}
		datas = append(datas, data)
	} else {
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
			"index_name":      index.IndexName,
			"primary_keys":    id,
			"partition":       searchOptions.partition,
		}
		if searchOptions.outputFields != nil {
			params["output_fields"] = searchOptions.outputFields
		}
		//fmt.Println(params)
		res, err := index.VikingDBService.DoRequest(context.Background(), "FetchIndexData", nil, index.VikingDBService.convertMapToJson(params))
		if err != nil {
			return nil, err
		}
		//fmt.Println(res)
		var resData []interface{}
		if d, ok := res["data"]; !ok {
			return nil, fmt.Errorf("invalid response, data does not exist: %v", res)
		} else if resData, ok = d.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not a list: %v", res)
		}
		for _, item := range resData {
			var itemMap map[string]interface{}
			var ok bool
			if itemMap, ok = item.(map[string]interface{}); !ok {
				return nil, fmt.Errorf("invalid response, data is not list[map]: %v", res)
			}
			primaryKey, err := index.getPrimaryKey()
			if err != nil {
				return nil, err
			}
			data := &Data{
				Id: itemMap[primaryKey],
				Fields: map[string]interface{}{
					"message": "no data found",
				},
				Timestamp: nil,
			}
			if _, exists := itemMap["fields"]; exists {
				var dataFields map[string]interface{}
				var ok bool
				if dataFields, ok = itemMap["fields"].(map[string]interface{}); !ok {
					return nil, fmt.Errorf("invalid response, fields is not a map: %v", res)
				}

				primaryKey, err := index.getPrimaryKey()
				if err != nil {
					return nil, err
				}
				data = &Data{
					Id:        itemMap[primaryKey],
					Fields:    dataFields,
					Timestamp: nil,
				}
			}

			datas = append(datas, data)
		}
	}
	return datas, nil
}
func (index *Index) getData(resData map[string]interface{}, outputField interface{}) ([]*Data, error) {
	var res []interface{}
	if d, ok := resData["data"]; !ok {
		return nil, fmt.Errorf("invalid response, data does not exist: %v", resData)
	} else if res, ok = d.([]interface{}); !ok {
		return nil, fmt.Errorf("invalid response, data is not a list: %v", resData)
	}

	datas := []*Data{}
	for _, items := range res {
		var itemsList []interface{}
		var ok bool
		if itemsList, ok = items.([]interface{}); !ok {
			return nil, fmt.Errorf("invalid response, data is not list[list]: %v", resData)
		}
		for _, l := range itemsList {
			var item map[string]interface{}
			var ok bool
			if item, ok = l.(map[string]interface{}); !ok {
				return nil, fmt.Errorf("invalid response, data is not list[list[map]]: %v", resData)
			}
			//fmt.Println(item)
			fields := map[string]interface{}{}
			if outputField == nil || len(outputField.([]string)) != 0 {
				if f, ok := item["fields"]; !ok {
					return nil, fmt.Errorf("invalid response, fields does not exist: %v", resData)
				} else if fields, ok = f.(map[string]interface{}); !ok {
					return nil, fmt.Errorf("invalid response, fields is not a map: %v", resData)
				}
			}
			text := ""
			if value, exists := item["text"]; exists {
				if t, ok := value.(string); !ok {
					return nil, fmt.Errorf("invalid response, text is not string: %v", resData)
				} else {
					text = t
				}
			}
			primaryKey, err := index.getPrimaryKey()
			if err != nil {
				return nil, err
			}
			var scoreFloat float64
			if s, ok := item["score"]; !ok {
				return nil, fmt.Errorf("invalid response, score not exists: %v", resData)
			} else if scoreFloat, ok = s.(float64); !ok {
				return nil, fmt.Errorf("invalid response, score is not float64: %v", resData)
			}
			data := &Data{
				Fields:    fields,
				Id:        item[primaryKey],
				Score:     scoreFloat,
				Timestamp: nil,
				Text:      text,
			}
			datas = append(datas, data)
		}
	}
	return datas, nil
}
func (index *Index) getPrimaryKey() (string, error) {
	if index.primaryKey == "" {
		// collection, err := index.VikingDBService.GetCollection(index.CollectionName)
		// if err != nil {
		// 	return "", err
		// }
		// index.primaryKey = collection.PrimaryKey
		params := map[string]interface{}{
			"collection_name": index.CollectionName,
		}
		resData, err := index.VikingDBService.DoRequest(context.Background(), "GetCollection", nil, index.VikingDBService.convertMapToJson(params))
		if err != nil {
			return "", err
		}
		var resDataItem map[string]interface{}
		if d, ok := resData["data"]; !ok {
			return "", fmt.Errorf("invalid response, data does not exist: %v", resData)
		} else if resDataItem, ok = d.(map[string]interface{}); !ok {
			return "", fmt.Errorf("invalid response, data is not a map: %v", resData)
		}
		if value, exist := resDataItem["primary_key"]; exist {
			if v, ok := value.(string); !ok {
				return "", fmt.Errorf("invalid response, primary_key is invalid:%v", resDataItem)
			} else {
				index.primaryKey = v
			}
		}
		return index.primaryKey, err
	} else {
		return index.primaryKey, nil
	}
}
