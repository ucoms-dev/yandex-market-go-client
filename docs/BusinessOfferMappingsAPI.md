# \BusinessOfferMappingsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOffersToArchive**](BusinessOfferMappingsAPI.md#AddOffersToArchive) | **Post** /v2/businesses/{businessId}/offer-mappings/archive | Добавление товаров в архив
[**DeleteOffers**](BusinessOfferMappingsAPI.md#DeleteOffers) | **Post** /v2/businesses/{businessId}/offer-mappings/delete | Удаление товаров из каталога
[**DeleteOffersFromArchive**](BusinessOfferMappingsAPI.md#DeleteOffersFromArchive) | **Post** /v2/businesses/{businessId}/offer-mappings/unarchive | Удаление товаров из архива
[**GenerateOfferBarcodes**](BusinessOfferMappingsAPI.md#GenerateOfferBarcodes) | **Post** /v1/businesses/{businessId}/offer-mappings/barcodes/generate | Генерация штрихкодов
[**GetOfferMappings**](BusinessOfferMappingsAPI.md#GetOfferMappings) | **Post** /v2/businesses/{businessId}/offer-mappings | Информация о товарах в каталоге
[**UpdateOfferMappings**](BusinessOfferMappingsAPI.md#UpdateOfferMappings) | **Post** /v2/businesses/{businessId}/offer-mappings/update | Добавление товаров в каталог и изменение информации о них



## AddOffersToArchive

> AddOffersToArchiveResponse AddOffersToArchive(ctx, businessId).AddOffersToArchiveRequest(addOffersToArchiveRequest).Execute()

Добавление товаров в архив



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
	addOffersToArchiveRequest := *openapiclient.NewAddOffersToArchiveRequest([]string{"OfferIds_example"}) // AddOffersToArchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessOfferMappingsAPI.AddOffersToArchive(context.Background(), businessId).AddOffersToArchiveRequest(addOffersToArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessOfferMappingsAPI.AddOffersToArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOffersToArchive`: AddOffersToArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `BusinessOfferMappingsAPI.AddOffersToArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOffersToArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addOffersToArchiveRequest** | [**AddOffersToArchiveRequest**](AddOffersToArchiveRequest.md) |  | 

### Return type

[**AddOffersToArchiveResponse**](AddOffersToArchiveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffers

> DeleteOffersResponse DeleteOffers(ctx, businessId).DeleteOffersRequest(deleteOffersRequest).Execute()

Удаление товаров из каталога



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
	deleteOffersRequest := *openapiclient.NewDeleteOffersRequest([]string{"OfferIds_example"}) // DeleteOffersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessOfferMappingsAPI.DeleteOffers(context.Background(), businessId).DeleteOffersRequest(deleteOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessOfferMappingsAPI.DeleteOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOffers`: DeleteOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `BusinessOfferMappingsAPI.DeleteOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteOffersRequest** | [**DeleteOffersRequest**](DeleteOffersRequest.md) |  | 

### Return type

[**DeleteOffersResponse**](DeleteOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffersFromArchive

> DeleteOffersFromArchiveResponse DeleteOffersFromArchive(ctx, businessId).DeleteOffersFromArchiveRequest(deleteOffersFromArchiveRequest).Execute()

Удаление товаров из архива



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
	deleteOffersFromArchiveRequest := *openapiclient.NewDeleteOffersFromArchiveRequest([]string{"OfferIds_example"}) // DeleteOffersFromArchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessOfferMappingsAPI.DeleteOffersFromArchive(context.Background(), businessId).DeleteOffersFromArchiveRequest(deleteOffersFromArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessOfferMappingsAPI.DeleteOffersFromArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOffersFromArchive`: DeleteOffersFromArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `BusinessOfferMappingsAPI.DeleteOffersFromArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOffersFromArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteOffersFromArchiveRequest** | [**DeleteOffersFromArchiveRequest**](DeleteOffersFromArchiveRequest.md) |  | 

### Return type

[**DeleteOffersFromArchiveResponse**](DeleteOffersFromArchiveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOfferBarcodes

> GenerateOfferBarcodesResponse GenerateOfferBarcodes(ctx, businessId).GenerateOfferBarcodesRequest(generateOfferBarcodesRequest).Execute()

Генерация штрихкодов



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
	generateOfferBarcodesRequest := *openapiclient.NewGenerateOfferBarcodesRequest([]string{"OfferIds_example"}) // GenerateOfferBarcodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessOfferMappingsAPI.GenerateOfferBarcodes(context.Background(), businessId).GenerateOfferBarcodesRequest(generateOfferBarcodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessOfferMappingsAPI.GenerateOfferBarcodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateOfferBarcodes`: GenerateOfferBarcodesResponse
	fmt.Fprintf(os.Stdout, "Response from `BusinessOfferMappingsAPI.GenerateOfferBarcodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOfferBarcodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateOfferBarcodesRequest** | [**GenerateOfferBarcodesRequest**](GenerateOfferBarcodesRequest.md) |  | 

### Return type

[**GenerateOfferBarcodesResponse**](GenerateOfferBarcodesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferMappings

> GetOfferMappingsResponse GetOfferMappings(ctx, businessId).PageToken(pageToken).Limit(limit).Language(language).GetOfferMappingsRequest(getOfferMappingsRequest).Execute()

Информация о товарах в каталоге



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
	limit := int32(56) // int32 | {{ limit-truncate-param-description }}  (optional) (default to 50)
	language := openapiclient.CatalogLanguageType("RU") // CatalogLanguageType | Язык, на котором принимаются и возвращаются значения в параметрах `name` и `description`.  Значение по умолчанию: `RU`.  (optional)
	getOfferMappingsRequest := *openapiclient.NewGetOfferMappingsRequest() // GetOfferMappingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessOfferMappingsAPI.GetOfferMappings(context.Background(), businessId).PageToken(pageToken).Limit(limit).Language(language).GetOfferMappingsRequest(getOfferMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessOfferMappingsAPI.GetOfferMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferMappings`: GetOfferMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `BusinessOfferMappingsAPI.GetOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-truncate-param-description }}  | [default to 50]
 **language** | [**CatalogLanguageType**](CatalogLanguageType.md) | Язык, на котором принимаются и возвращаются значения в параметрах &#x60;name&#x60; и &#x60;description&#x60;.  Значение по умолчанию: &#x60;RU&#x60;.  | 
 **getOfferMappingsRequest** | [**GetOfferMappingsRequest**](GetOfferMappingsRequest.md) |  | 

### Return type

[**GetOfferMappingsResponse**](GetOfferMappingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfferMappings

> UpdateOfferMappingsResponse UpdateOfferMappings(ctx, businessId).UpdateOfferMappingsRequest(updateOfferMappingsRequest).Language(language).Execute()

Добавление товаров в каталог и изменение информации о них



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
	updateOfferMappingsRequest := *openapiclient.NewUpdateOfferMappingsRequest([]openapiclient.UpdateOfferMappingDTO{*openapiclient.NewUpdateOfferMappingDTO(*openapiclient.NewUpdateOfferDTO("OfferId_example"))}) // UpdateOfferMappingsRequest | 
	language := openapiclient.CatalogLanguageType("RU") // CatalogLanguageType | Язык, на котором принимаются и возвращаются значения в параметрах `name` и `description`.  Значение по умолчанию: `RU`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessOfferMappingsAPI.UpdateOfferMappings(context.Background(), businessId).UpdateOfferMappingsRequest(updateOfferMappingsRequest).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessOfferMappingsAPI.UpdateOfferMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOfferMappings`: UpdateOfferMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `BusinessOfferMappingsAPI.UpdateOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOfferMappingsRequest** | [**UpdateOfferMappingsRequest**](UpdateOfferMappingsRequest.md) |  | 
 **language** | [**CatalogLanguageType**](CatalogLanguageType.md) | Язык, на котором принимаются и возвращаются значения в параметрах &#x60;name&#x60; и &#x60;description&#x60;.  Значение по умолчанию: &#x60;RU&#x60;.  | 

### Return type

[**UpdateOfferMappingsResponse**](UpdateOfferMappingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

