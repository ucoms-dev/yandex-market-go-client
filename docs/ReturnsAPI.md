# \ReturnsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReturn**](ReturnsAPI.md#GetReturn) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId} | Информация о невыкупе или возврате
[**GetReturnApplication**](ReturnsAPI.md#GetReturnApplication) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/application | Получение заявления на возврат
[**GetReturnPhoto**](ReturnsAPI.md#GetReturnPhoto) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/{itemId}/image/{imageHash} | Получение фотографий товаров в возврате
[**GetReturns**](ReturnsAPI.md#GetReturns) | **Get** /campaigns/{campaignId}/returns | Список невыкупов и возвратов
[**SetReturnDecision**](ReturnsAPI.md#SetReturnDecision) | **Post** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision | Принятие или изменение решения по возврату
[**SubmitReturnDecision**](ReturnsAPI.md#SubmitReturnDecision) | **Post** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/submit | Подтверждение решения по возврату



## GetReturn

> GetReturnResponse GetReturn(ctx, campaignId, orderId, returnId).Execute()

Информация о невыкупе или возврате



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.GetReturn(context.Background(), campaignId, orderId, returnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturn`: GetReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор невыкупа или возврата. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetReturnResponse**](GetReturnResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnApplication

> *os.File GetReturnApplication(ctx, campaignId, orderId, returnId).Execute()

Получение заявления на возврат



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.GetReturnApplication(context.Background(), campaignId, orderId, returnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturnApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturnApplication`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturnApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор невыкупа или возврата. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnPhoto

> *os.File GetReturnPhoto(ctx, campaignId, orderId, returnId, itemId, imageHash).Execute()

Получение фотографий товаров в возврате



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.
	itemId := int64(789) // int64 | Идентификатор товара в возврате.
	imageHash := "imageHash_example" // string | Хеш ссылки изображения для загрузки.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.GetReturnPhoto(context.Background(), campaignId, orderId, returnId, itemId, imageHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturnPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturnPhoto`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturnPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор невыкупа или возврата. | 
**itemId** | **int64** | Идентификатор товара в возврате. | 
**imageHash** | **string** | Хеш ссылки изображения для загрузки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturns

> GetReturnsResponse GetReturns(ctx, campaignId).PageToken(pageToken).Limit(limit).OrderIds(orderIds).Statuses(statuses).Type_(type_).FromDate(fromDate).ToDate(toDate).FromDate2(fromDate2).ToDate2(toDate2).Execute()

Список невыкупов и возвратов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)
	orderIds := []int64{int64(123543)} // []int64 | Идентификаторы заказов — для фильтрации результатов.  Несколько идентификаторов перечисляются через запятую без пробела.  (optional)
	statuses := []openapiclient.RefundStatusType{openapiclient.RefundStatusType("STARTED_BY_USER")} // []RefundStatusType | Статусы невыкупов или возвратов — для фильтрации результатов.  Несколько статусов перечисляются через запятую.  (optional)
	type_ := openapiclient.ReturnType("UNREDEEMED") // ReturnType | Тип заказа для фильтрации:  * `UNREDEEMED` — невыкуп.  * `RETURN` — возврат.  Если не указать, в ответе будут и невыкупы, и возвраты.  (optional)
	fromDate := time.Now() // string | Начальная дата для фильтрации невыкупов или возвратов по дате обновления.  Формат: `ГГГГ-ММ-ДД`.  (optional)
	toDate := time.Now() // string | Конечная дата для фильтрации невыкупов или возвратов по дате обновления.  Формат: `ГГГГ-ММ-ДД`.  (optional)
	fromDate2 := time.Now() // string | {% note warning \"Вместо него используйте `fromDate`.\" %}     {% endnote %}  Начальная дата для фильтрации невыкупов или возвратов по дате обновления.  (optional)
	toDate2 := time.Now() // string | {% note warning \"Вместо него используйте `toDate`.\" %}     {% endnote %}  Конечная дата для фильтрации невыкупов или возвратов по дате обновления.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.GetReturns(context.Background(), campaignId).PageToken(pageToken).Limit(limit).OrderIds(orderIds).Statuses(statuses).Type_(type_).FromDate(fromDate).ToDate(toDate).FromDate2(fromDate2).ToDate2(toDate2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.GetReturns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturns`: GetReturnsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.GetReturns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
 **orderIds** | **[]int64** | Идентификаторы заказов — для фильтрации результатов.  Несколько идентификаторов перечисляются через запятую без пробела.  | 
 **statuses** | [**[]RefundStatusType**](RefundStatusType.md) | Статусы невыкупов или возвратов — для фильтрации результатов.  Несколько статусов перечисляются через запятую.  | 
 **type_** | [**ReturnType**](ReturnType.md) | Тип заказа для фильтрации:  * &#x60;UNREDEEMED&#x60; — невыкуп.  * &#x60;RETURN&#x60; — возврат.  Если не указать, в ответе будут и невыкупы, и возвраты.  | 
 **fromDate** | **string** | Начальная дата для фильтрации невыкупов или возвратов по дате обновления.  Формат: &#x60;ГГГГ-ММ-ДД&#x60;.  | 
 **toDate** | **string** | Конечная дата для фильтрации невыкупов или возвратов по дате обновления.  Формат: &#x60;ГГГГ-ММ-ДД&#x60;.  | 
 **fromDate2** | **string** | {% note warning \&quot;Вместо него используйте &#x60;fromDate&#x60;.\&quot; %}     {% endnote %}  Начальная дата для фильтрации невыкупов или возвратов по дате обновления.  | 
 **toDate2** | **string** | {% note warning \&quot;Вместо него используйте &#x60;toDate&#x60;.\&quot; %}     {% endnote %}  Конечная дата для фильтрации невыкупов или возвратов по дате обновления.  | 

### Return type

[**GetReturnsResponse**](GetReturnsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetReturnDecision

> EmptyApiResponse SetReturnDecision(ctx, campaignId, orderId, returnId).SetReturnDecisionRequest(setReturnDecisionRequest).Execute()

Принятие или изменение решения по возврату



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.
	setReturnDecisionRequest := *openapiclient.NewSetReturnDecisionRequest(int64(123), openapiclient.ReturnRequestDecisionType("REFUND_MONEY")) // SetReturnDecisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.SetReturnDecision(context.Background(), campaignId, orderId, returnId).SetReturnDecisionRequest(setReturnDecisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.SetReturnDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetReturnDecision`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.SetReturnDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор невыкупа или возврата. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetReturnDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setReturnDecisionRequest** | [**SetReturnDecisionRequest**](SetReturnDecisionRequest.md) |  | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitReturnDecision

> EmptyApiResponse SubmitReturnDecision(ctx, campaignId, orderId, returnId).Body(body).Execute()

Подтверждение решения по возврату



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.
	body := map[string]interface{}{ ... } // map[string]interface{} | description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAPI.SubmitReturnDecision(context.Background(), campaignId, orderId, returnId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAPI.SubmitReturnDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitReturnDecision`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAPI.SubmitReturnDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 
**returnId** | **int64** | Идентификатор невыкупа или возврата. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitReturnDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **map[string]interface{}** | description | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

