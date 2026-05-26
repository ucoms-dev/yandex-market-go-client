# \DeliveryOptionsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeliveryOptions**](DeliveryOptionsAPI.md#GetDeliveryOptions) | **Post** /v1/campaigns/{campaignId}/delivery-options | Получение доступных вариантов доставки заказов
[**GetReturnDeliveryOptions**](DeliveryOptionsAPI.md#GetReturnDeliveryOptions) | **Post** /v1/campaigns/{campaignId}/return-delivery-options | Получение подходящих для возврата пунктов выдачи



## GetDeliveryOptions

> GetDeliveryOptionsResponse GetDeliveryOptions(ctx, campaignId).GetDeliveryOptionsRequest(getDeliveryOptionsRequest).Execute()

Получение доступных вариантов доставки заказов



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
	getDeliveryOptionsRequest := *openapiclient.NewGetDeliveryOptionsRequest([]openapiclient.GetDeliveryOptionsItemDTO{*openapiclient.NewGetDeliveryOptionsItemDTO("OfferId_example", int32(123))}) // GetDeliveryOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveryOptionsAPI.GetDeliveryOptions(context.Background(), campaignId).GetDeliveryOptionsRequest(getDeliveryOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOptionsAPI.GetDeliveryOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeliveryOptions`: GetDeliveryOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeliveryOptionsAPI.GetDeliveryOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getDeliveryOptionsRequest** | [**GetDeliveryOptionsRequest**](GetDeliveryOptionsRequest.md) |  | 

### Return type

[**GetDeliveryOptionsResponse**](GetDeliveryOptionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReturnDeliveryOptions

> GetReturnDeliveryOptionsResponse GetReturnDeliveryOptions(ctx, campaignId).GetReturnDeliveryOptionsRequest(getReturnDeliveryOptionsRequest).Execute()

Получение подходящих для возврата пунктов выдачи



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
	getReturnDeliveryOptionsRequest := *openapiclient.NewGetReturnDeliveryOptionsRequest([]openapiclient.BasicOrderItemDTO{*openapiclient.NewBasicOrderItemDTO("OfferId_example", int32(123))}, *openapiclient.NewPickupDeliveryParametersDTO([]int64{int64(123)})) // GetReturnDeliveryOptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveryOptionsAPI.GetReturnDeliveryOptions(context.Background(), campaignId).GetReturnDeliveryOptionsRequest(getReturnDeliveryOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryOptionsAPI.GetReturnDeliveryOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturnDeliveryOptions`: GetReturnDeliveryOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeliveryOptionsAPI.GetReturnDeliveryOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnDeliveryOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getReturnDeliveryOptionsRequest** | [**GetReturnDeliveryOptionsRequest**](GetReturnDeliveryOptionsRequest.md) |  | 

### Return type

[**GetReturnDeliveryOptionsResponse**](GetReturnDeliveryOptionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

