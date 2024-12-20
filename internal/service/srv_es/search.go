package srv_es

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
	"gvb/internal/global"
)

func DateAggregateSearch(index string, aggField string, interval string) ([]types.DateHistogramBucket, error) {
	var agg = make(map[string]types.Aggregations)
	format := "yyyy-MM-dd"
	aggName := "date"
	agg[aggName] = types.Aggregations{
		DateHistogram: &types.DateHistogramAggregation{
			Field:            &aggField,                                          //根据date字段进行分组
			CalendarInterval: &calendarinterval.CalendarInterval{Name: interval}, //按天分组
			Format:           &format,
		}}

	resp, err := global.ES.Search().
		Index(index).
		Aggregations(
			agg,
		).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return nil, err
	}
	aggResp, _ := resp.Aggregations[aggName].(*types.DateHistogramAggregate)
	return aggResp.Buckets.([]types.DateHistogramBucket), nil

}
func TagAggregateSearch(index string, aggName string, subAggName string) ([]types.StringTermsBucket, error) {
	var agg = make(map[string]types.Aggregations)
	var subAgg = make(map[string]types.Aggregations)
	requiredField := []string{"title"}
	termField := "tags.keyword"
	subAgg[subAggName] = types.Aggregations{
		TopHits: &types.TopHitsAggregation{
			Source_: &types.SourceFilter{
				Includes: requiredField,
			},
		},
	}
	agg[aggName] = types.Aggregations{
		Terms: &types.TermsAggregation{
			Field: &termField,
		},
		Aggregations: subAgg,
	}
	resp, err := global.ES.Search().
		Index(index).
		Aggregations(
			agg,
		).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return nil, err
	}

	buckets := resp.Aggregations[aggName].(*types.StringTermsAggregate).Buckets.([]types.StringTermsBucket)
	return buckets, nil
}
