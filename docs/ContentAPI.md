# \ContentAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContentSizeTables**](ContentAPI.md#CreateContentSizeTables) | **Post** /v2/businesses/{businessId}/size-tables/create | Создание таблиц размеров
[**GetCategoryContentParameters**](ContentAPI.md#GetCategoryContentParameters) | **Post** /v2/category/{categoryId}/parameters | Списки характеристик товаров по категориям
[**GetContentSizeCharts**](ContentAPI.md#GetContentSizeCharts) | **Post** /v2/businesses/{businessId}/size-tables/charts | Размерные сетки и обмеры для таблиц размеров
[**GetContentSizeTables**](ContentAPI.md#GetContentSizeTables) | **Post** /v2/businesses/{businessId}/size-tables | Таблицы размеров по идентификаторам
[**GetOfferCardsContentStatus**](ContentAPI.md#GetOfferCardsContentStatus) | **Post** /v2/businesses/{businessId}/offer-cards | Получение информации о заполненности карточек магазина
[**SuggestContentSizeTables**](ContentAPI.md#SuggestContentSizeTables) | **Post** /v2/businesses/{businessId}/size-tables/suggest | Подбор таблиц размеров
[**UpdateContentSizeTables**](ContentAPI.md#UpdateContentSizeTables) | **Post** /v2/businesses/{businessId}/size-tables/update | Обновление таблиц размеров
[**UpdateOfferContent**](ContentAPI.md#UpdateOfferContent) | **Post** /v2/businesses/{businessId}/offer-cards/update | Редактирование категорийных характеристик товара



## CreateContentSizeTables

> CreateSizeTablesResponse CreateContentSizeTables(ctx, businessId).CreateSizeTablesRequest(createSizeTablesRequest).Execute()

Создание таблиц размеров



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
	createSizeTablesRequest := *openapiclient.NewCreateSizeTablesRequest([]openapiclient.CreateSizeTableDTO{*openapiclient.NewCreateSizeTableDTO(int64(123), "SizeTableName_example", []openapiclient.SizeTableRowDTO{*openapiclient.NewSizeTableRowDTO([]openapiclient.SizeTableCellDTO{*openapiclient.NewSizeTableCellDTO()})})}) // CreateSizeTablesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.CreateContentSizeTables(context.Background(), businessId).CreateSizeTablesRequest(createSizeTablesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.CreateContentSizeTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContentSizeTables`: CreateSizeTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.CreateContentSizeTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContentSizeTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSizeTablesRequest** | [**CreateSizeTablesRequest**](CreateSizeTablesRequest.md) |  | 

### Return type

[**CreateSizeTablesResponse**](CreateSizeTablesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryContentParameters

> GetCategoryContentParametersResponse GetCategoryContentParameters(ctx, categoryId).BusinessId(businessId).Execute()

Списки характеристик товаров по категориям



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
	categoryId := int64(789) // int64 | Идентификатор категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится интересующий вас товар, воспользуйтесь запросом [POST v2/categories/tree](../../reference/categories/getCategoriesTree.md). 
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  Передайте параметр, чтобы получить характеристики, которые являются особенностями варианта товара в данном кабинете.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetCategoryContentParameters(context.Background(), categoryId).BusinessId(businessId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetCategoryContentParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryContentParameters`: GetCategoryContentParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetCategoryContentParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Идентификатор категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится интересующий вас товар, воспользуйтесь запросом [POST v2/categories/tree](../../reference/categories/getCategoriesTree.md).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryContentParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  Передайте параметр, чтобы получить характеристики, которые являются особенностями варианта товара в данном кабинете.  | 

### Return type

[**GetCategoryContentParametersResponse**](GetCategoryContentParametersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentSizeCharts

> GetSizeChartsResponse GetContentSizeCharts(ctx, businessId).GetSizeChartsRequest(getSizeChartsRequest).Execute()

Размерные сетки и обмеры для таблиц размеров



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
	getSizeChartsRequest := *openapiclient.NewGetSizeChartsRequest(int64(123)) // GetSizeChartsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContentSizeCharts(context.Background(), businessId).GetSizeChartsRequest(getSizeChartsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetContentSizeCharts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentSizeCharts`: GetSizeChartsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetContentSizeCharts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentSizeChartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSizeChartsRequest** | [**GetSizeChartsRequest**](GetSizeChartsRequest.md) |  | 

### Return type

[**GetSizeChartsResponse**](GetSizeChartsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentSizeTables

> GetSizeTablesResponse GetContentSizeTables(ctx, businessId).GetSizeTablesRequest(getSizeTablesRequest).PageToken(pageToken).Limit(limit).Execute()

Таблицы размеров по идентификаторам



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
	getSizeTablesRequest := *openapiclient.NewGetSizeTablesRequest() // GetSizeTablesRequest | 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContentSizeTables(context.Background(), businessId).GetSizeTablesRequest(getSizeTablesRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetContentSizeTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentSizeTables`: GetSizeTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetContentSizeTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentSizeTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSizeTablesRequest** | [**GetSizeTablesRequest**](GetSizeTablesRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 25]

### Return type

[**GetSizeTablesResponse**](GetSizeTablesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfferCardsContentStatus

> GetOfferCardsContentStatusResponse GetOfferCardsContentStatus(ctx, businessId).PageToken(pageToken).Limit(limit).GetOfferCardsContentStatusRequest(getOfferCardsContentStatusRequest).Execute()

Получение информации о заполненности карточек магазина



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
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 100)
	getOfferCardsContentStatusRequest := *openapiclient.NewGetOfferCardsContentStatusRequest() // GetOfferCardsContentStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetOfferCardsContentStatus(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetOfferCardsContentStatusRequest(getOfferCardsContentStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetOfferCardsContentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferCardsContentStatus`: GetOfferCardsContentStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetOfferCardsContentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferCardsContentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 100]
 **getOfferCardsContentStatusRequest** | [**GetOfferCardsContentStatusRequest**](GetOfferCardsContentStatusRequest.md) |  | 

### Return type

[**GetOfferCardsContentStatusResponse**](GetOfferCardsContentStatusResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuggestContentSizeTables

> SuggestSizeTablesResponse SuggestContentSizeTables(ctx, businessId).SuggestSizeTablesRequest(suggestSizeTablesRequest).PageToken(pageToken).Limit(limit).Execute()

Подбор таблиц размеров



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
	suggestSizeTablesRequest := *openapiclient.NewSuggestSizeTablesRequest(int32(123)) // SuggestSizeTablesRequest | 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.SuggestContentSizeTables(context.Background(), businessId).SuggestSizeTablesRequest(suggestSizeTablesRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.SuggestContentSizeTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuggestContentSizeTables`: SuggestSizeTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.SuggestContentSizeTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuggestContentSizeTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suggestSizeTablesRequest** | [**SuggestSizeTablesRequest**](SuggestSizeTablesRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 25]

### Return type

[**SuggestSizeTablesResponse**](SuggestSizeTablesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContentSizeTables

> UpdateSizeTablesResponse UpdateContentSizeTables(ctx, businessId).UpdateSizeTablesRequest(updateSizeTablesRequest).Execute()

Обновление таблиц размеров



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
	updateSizeTablesRequest := *openapiclient.NewUpdateSizeTablesRequest([]openapiclient.UpdateSizeTableDTO{*openapiclient.NewUpdateSizeTableDTO(int64(123), "SizeTableName_example", []openapiclient.SizeTableRowDTO{*openapiclient.NewSizeTableRowDTO([]openapiclient.SizeTableCellDTO{*openapiclient.NewSizeTableCellDTO()})})}) // UpdateSizeTablesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.UpdateContentSizeTables(context.Background(), businessId).UpdateSizeTablesRequest(updateSizeTablesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.UpdateContentSizeTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContentSizeTables`: UpdateSizeTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.UpdateContentSizeTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContentSizeTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSizeTablesRequest** | [**UpdateSizeTablesRequest**](UpdateSizeTablesRequest.md) |  | 

### Return type

[**UpdateSizeTablesResponse**](UpdateSizeTablesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfferContent

> UpdateOfferContentResponse UpdateOfferContent(ctx, businessId).UpdateOfferContentRequest(updateOfferContentRequest).Execute()

Редактирование категорийных характеристик товара



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
	updateOfferContentRequest := *openapiclient.NewUpdateOfferContentRequest([]openapiclient.OfferContentDTO{*openapiclient.NewOfferContentDTO("OfferId_example", int32(123), []openapiclient.ParameterValueDTO{*openapiclient.NewParameterValueDTO(int64(123))})}) // UpdateOfferContentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.UpdateOfferContent(context.Background(), businessId).UpdateOfferContentRequest(updateOfferContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.UpdateOfferContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOfferContent`: UpdateOfferContentResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.UpdateOfferContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOfferContentRequest** | [**UpdateOfferContentRequest**](UpdateOfferContentRequest.md) |  | 

### Return type

[**UpdateOfferContentResponse**](UpdateOfferContentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

