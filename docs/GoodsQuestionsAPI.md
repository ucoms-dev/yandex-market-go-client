# \GoodsQuestionsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGoodsQuestionAnswers**](GoodsQuestionsAPI.md#GetGoodsQuestionAnswers) | **Post** /v1/businesses/{businessId}/goods-questions/answers | Получение ответов на вопрос
[**GetGoodsQuestions**](GoodsQuestionsAPI.md#GetGoodsQuestions) | **Post** /v1/businesses/{businessId}/goods-questions | Получение вопросов о товарах продавца
[**UpdateGoodsQuestionTextEntity**](GoodsQuestionsAPI.md#UpdateGoodsQuestionTextEntity) | **Post** /v1/businesses/{businessId}/goods-questions/update | Создание, изменение и удаление ответа или комментария



## GetGoodsQuestionAnswers

> GetAnswersResponse GetGoodsQuestionAnswers(ctx, businessId).GetAnswersRequest(getAnswersRequest).PageToken(pageToken).Limit(limit).Execute()

Получение ответов на вопрос



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getAnswersRequest := *openapiclient.NewGetAnswersRequest() // GetAnswersRequest | 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsQuestionsAPI.GetGoodsQuestionAnswers(context.Background(), businessId).GetAnswersRequest(getAnswersRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsQuestionsAPI.GetGoodsQuestionAnswers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsQuestionAnswers`: GetAnswersResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsQuestionsAPI.GetGoodsQuestionAnswers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsQuestionAnswersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getAnswersRequest** | [**GetAnswersRequest**](GetAnswersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 25]

### Return type

[**GetAnswersResponse**](GetAnswersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoodsQuestions

> GetQuestionsResponse GetGoodsQuestions(ctx, businessId).PageToken(pageToken).Limit(limit).GetQuestionsRequest(getQuestionsRequest).Execute()

Получение вопросов о товарах продавца



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 25)
	getQuestionsRequest := *openapiclient.NewGetQuestionsRequest() // GetQuestionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsQuestionsAPI.GetGoodsQuestions(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetQuestionsRequest(getQuestionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsQuestionsAPI.GetGoodsQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsQuestions`: GetQuestionsResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsQuestionsAPI.GetGoodsQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 25]
 **getQuestionsRequest** | [**GetQuestionsRequest**](GetQuestionsRequest.md) |  | 

### Return type

[**GetQuestionsResponse**](GetQuestionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGoodsQuestionTextEntity

> UpdateGoodsQuestionTextEntityResponse UpdateGoodsQuestionTextEntity(ctx, businessId).UpdateGoodsQuestionTextEntityRequest(updateGoodsQuestionTextEntityRequest).Execute()

Создание, изменение и удаление ответа или комментария



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	updateGoodsQuestionTextEntityRequest := *openapiclient.NewUpdateGoodsQuestionTextEntityRequest(openapiclient.QuestionsTextEntityOperationType("UPDATE")) // UpdateGoodsQuestionTextEntityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsQuestionsAPI.UpdateGoodsQuestionTextEntity(context.Background(), businessId).UpdateGoodsQuestionTextEntityRequest(updateGoodsQuestionTextEntityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsQuestionsAPI.UpdateGoodsQuestionTextEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoodsQuestionTextEntity`: UpdateGoodsQuestionTextEntityResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsQuestionsAPI.UpdateGoodsQuestionTextEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoodsQuestionTextEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoodsQuestionTextEntityRequest** | [**UpdateGoodsQuestionTextEntityRequest**](UpdateGoodsQuestionTextEntityRequest.md) |  | 

### Return type

[**UpdateGoodsQuestionTextEntityResponse**](UpdateGoodsQuestionTextEntityResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

