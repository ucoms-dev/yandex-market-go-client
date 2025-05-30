# \SupplyRequestsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupplyRequestDocuments**](SupplyRequestsAPI.md#GetSupplyRequestDocuments) | **Post** /campaigns/{campaignId}/supply-requests/documents | Получение документов по заявке на поставку, вывоз или утилизацию
[**GetSupplyRequestItems**](SupplyRequestsAPI.md#GetSupplyRequestItems) | **Post** /campaigns/{campaignId}/supply-requests/items | Получение товаров в заявке на поставку, вывоз или утилизацию
[**GetSupplyRequests**](SupplyRequestsAPI.md#GetSupplyRequests) | **Post** /campaigns/{campaignId}/supply-requests | Получение информации о заявках на поставку, вывоз и утилизацию



## GetSupplyRequestDocuments

> GetSupplyRequestDocumentsResponse GetSupplyRequestDocuments(ctx, campaignId).GetSupplyRequestDocumentsRequest(getSupplyRequestDocumentsRequest).Execute()

Получение документов по заявке на поставку, вывоз или утилизацию



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	getSupplyRequestDocumentsRequest := *openapiclient.NewGetSupplyRequestDocumentsRequest(int64(123)) // GetSupplyRequestDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupplyRequestsAPI.GetSupplyRequestDocuments(context.Background(), campaignId).GetSupplyRequestDocumentsRequest(getSupplyRequestDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupplyRequestsAPI.GetSupplyRequestDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupplyRequestDocuments`: GetSupplyRequestDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `SupplyRequestsAPI.GetSupplyRequestDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyRequestDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSupplyRequestDocumentsRequest** | [**GetSupplyRequestDocumentsRequest**](GetSupplyRequestDocumentsRequest.md) |  | 

### Return type

[**GetSupplyRequestDocumentsResponse**](GetSupplyRequestDocumentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupplyRequestItems

> GetSupplyRequestItemsResponse GetSupplyRequestItems(ctx, campaignId).GetSupplyRequestItemsRequest(getSupplyRequestItemsRequest).PageToken(pageToken).Limit(limit).Execute()

Получение товаров в заявке на поставку, вывоз или утилизацию



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	getSupplyRequestItemsRequest := *openapiclient.NewGetSupplyRequestItemsRequest(int64(123)) // GetSupplyRequestItemsRequest | 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupplyRequestsAPI.GetSupplyRequestItems(context.Background(), campaignId).GetSupplyRequestItemsRequest(getSupplyRequestItemsRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupplyRequestsAPI.GetSupplyRequestItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupplyRequestItems`: GetSupplyRequestItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `SupplyRequestsAPI.GetSupplyRequestItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyRequestItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSupplyRequestItemsRequest** | [**GetSupplyRequestItemsRequest**](GetSupplyRequestItemsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetSupplyRequestItemsResponse**](GetSupplyRequestItemsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupplyRequests

> GetSupplyRequestsResponse GetSupplyRequests(ctx, campaignId).PageToken(pageToken).Limit(limit).GetSupplyRequestsRequest(getSupplyRequestsRequest).Execute()

Получение информации о заявках на поставку, вывоз и утилизацию



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)
	getSupplyRequestsRequest := *openapiclient.NewGetSupplyRequestsRequest() // GetSupplyRequestsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupplyRequestsAPI.GetSupplyRequests(context.Background(), campaignId).PageToken(pageToken).Limit(limit).GetSupplyRequestsRequest(getSupplyRequestsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupplyRequestsAPI.GetSupplyRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupplyRequests`: GetSupplyRequestsResponse
	fmt.Fprintf(os.Stdout, "Response from `SupplyRequestsAPI.GetSupplyRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
 **getSupplyRequestsRequest** | [**GetSupplyRequestsRequest**](GetSupplyRequestsRequest.md) |  | 

### Return type

[**GetSupplyRequestsResponse**](GetSupplyRequestsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

