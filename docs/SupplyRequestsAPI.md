# \SupplyRequestsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupplyRequestDocuments**](SupplyRequestsAPI.md#GetSupplyRequestDocuments) | **Post** /v2/campaigns/{campaignId}/supply-requests/documents | Получение документов по заявке на поставку, вывоз или утилизацию
[**GetSupplyRequestItems**](SupplyRequestsAPI.md#GetSupplyRequestItems) | **Post** /v2/campaigns/{campaignId}/supply-requests/items | Получение товаров в заявке на поставку, вывоз или утилизацию
[**GetSupplyRequests**](SupplyRequestsAPI.md#GetSupplyRequests) | **Post** /v2/campaigns/{campaignId}/supply-requests | Получение информации о заявках на поставку, вывоз и утилизацию



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
	campaignId := int64(789) // int64 | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями. 
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
**campaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 

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
	campaignId := int64(789) // int64 | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями. 
	getSupplyRequestItemsRequest := *openapiclient.NewGetSupplyRequestItemsRequest(int64(123)) // GetSupplyRequestItemsRequest | 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 250)

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
**campaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyRequestItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSupplyRequestItemsRequest** | [**GetSupplyRequestItemsRequest**](GetSupplyRequestItemsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 250]

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
	campaignId := int64(789) // int64 | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями. 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 50)
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
**campaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplyRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 50]
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

