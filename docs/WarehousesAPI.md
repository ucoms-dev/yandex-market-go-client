# \WarehousesAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFulfillmentWarehouses**](WarehousesAPI.md#GetFulfillmentWarehouses) | **Get** /v2/warehouses | Идентификаторы фулфилмент-складов Маркета
[**GetPagedWarehouses**](WarehousesAPI.md#GetPagedWarehouses) | **Post** /v2/businesses/{businessId}/warehouses | Список складов
[**GetWarehouses**](WarehousesAPI.md#GetWarehouses) | **Get** /v2/businesses/{businessId}/warehouses | Список складов и групп складов
[**UpdateWarehouseStatus**](WarehousesAPI.md#UpdateWarehouseStatus) | **Post** /v2/campaigns/{campaignId}/warehouse/status | Изменение статуса склада



## GetFulfillmentWarehouses

> GetFulfillmentWarehousesResponse GetFulfillmentWarehouses(ctx).CampaignId(campaignId).Execute()

Идентификаторы фулфилмент-складов Маркета



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
	campaignId := int64(789) // int64 | Идентификатор кампании магазина.  Указывается, если нужно вернуть все склады Маркета, которые привязаны к определенной кампании магазина.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.GetFulfillmentWarehouses(context.Background()).CampaignId(campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.GetFulfillmentWarehouses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFulfillmentWarehouses`: GetFulfillmentWarehousesResponse
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.GetFulfillmentWarehouses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFulfillmentWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignId** | **int64** | Идентификатор кампании магазина.  Указывается, если нужно вернуть все склады Маркета, которые привязаны к определенной кампании магазина.  | 

### Return type

[**GetFulfillmentWarehousesResponse**](GetFulfillmentWarehousesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagedWarehouses

> GetPagedWarehousesResponse GetPagedWarehouses(ctx, businessId).PageToken(pageToken).Limit(limit).GetPagedWarehousesRequest(getPagedWarehousesRequest).Execute()

Список складов



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
	businessId := int64(789) // int64 | Идентификатор кабинета.  {% if audience == \"partner\" %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %} 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 15)
	getPagedWarehousesRequest := *openapiclient.NewGetPagedWarehousesRequest() // GetPagedWarehousesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.GetPagedWarehouses(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetPagedWarehousesRequest(getPagedWarehousesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.GetPagedWarehouses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagedWarehouses`: GetPagedWarehousesResponse
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.GetPagedWarehouses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagedWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 15]
 **getPagedWarehousesRequest** | [**GetPagedWarehousesRequest**](GetPagedWarehousesRequest.md) |  | 

### Return type

[**GetPagedWarehousesResponse**](GetPagedWarehousesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarehouses

> GetWarehousesResponse GetWarehouses(ctx, businessId).Execute()

Список складов и групп складов



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
	businessId := int64(789) // int64 | Идентификатор кабинета.  {% if audience == \"partner\" %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %} 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.GetWarehouses(context.Background(), businessId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.GetWarehouses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarehouses`: GetWarehousesResponse
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.GetWarehouses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWarehousesResponse**](GetWarehousesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWarehouseStatus

> UpdateWarehouseStatusResponse UpdateWarehouseStatus(ctx, campaignId).UpdateWarehouseStatusRequest(updateWarehouseStatusRequest).Execute()

Изменение статуса склада



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
	updateWarehouseStatusRequest := *openapiclient.NewUpdateWarehouseStatusRequest(false) // UpdateWarehouseStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.UpdateWarehouseStatus(context.Background(), campaignId).UpdateWarehouseStatusRequest(updateWarehouseStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.UpdateWarehouseStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWarehouseStatus`: UpdateWarehouseStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.UpdateWarehouseStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWarehouseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWarehouseStatusRequest** | [**UpdateWarehouseStatusRequest**](UpdateWarehouseStatusRequest.md) |  | 

### Return type

[**UpdateWarehouseStatusResponse**](UpdateWarehouseStatusResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

