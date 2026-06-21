# \GoodsFeedbackAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGoodsFeedbackComment**](GoodsFeedbackAPI.md#DeleteGoodsFeedbackComment) | **Post** /v2/businesses/{businessId}/goods-feedback/comments/delete | Удаление комментария к отзыву
[**GetGoodsFeedbackComments**](GoodsFeedbackAPI.md#GetGoodsFeedbackComments) | **Post** /v2/businesses/{businessId}/goods-feedback/comments | Получение комментариев к отзыву
[**GetGoodsFeedbacks**](GoodsFeedbackAPI.md#GetGoodsFeedbacks) | **Post** /v2/businesses/{businessId}/goods-feedback | Получение отзывов о товарах продавца
[**GetGoodsFeedbacksUrbanads**](GoodsFeedbackAPI.md#GetGoodsFeedbacksUrbanads) | **Post** /v1/businesses/{businessId}/goods-feedback-advertiser | Получение отзывов о товарах для рекламодателей
[**SkipGoodsFeedbacksReaction**](GoodsFeedbackAPI.md#SkipGoodsFeedbacksReaction) | **Post** /v2/businesses/{businessId}/goods-feedback/skip-reaction | Пропуск реакции на отзывы
[**UpdateGoodsFeedbackComment**](GoodsFeedbackAPI.md#UpdateGoodsFeedbackComment) | **Post** /v2/businesses/{businessId}/goods-feedback/comments/update | Добавление нового или изменение созданного комментария



## DeleteGoodsFeedbackComment

> EmptyApiResponse DeleteGoodsFeedbackComment(ctx, businessId).DeleteGoodsFeedbackCommentRequest(deleteGoodsFeedbackCommentRequest).SourceType(sourceType).Execute()

Удаление комментария к отзыву



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
	deleteGoodsFeedbackCommentRequest := *openapiclient.NewDeleteGoodsFeedbackCommentRequest(int64(123)) // DeleteGoodsFeedbackCommentRequest | 
	sourceType := openapiclient.SourceType("SELLER") // SourceType | Признак типа кабинета, от имени которого вызывается метод: {% if audience == \"partner\" %}  - `SELLER` — продавец.  {% endif %}  - `ADVERTISER` — рекламодатель.  {% if audience == \"advertiser\" %}  {% note info \"Обязательно указывайте sourceType=ADVERTISER в каждом запросе.\" %}     {% endnote %}  {% endif %}  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsFeedbackAPI.DeleteGoodsFeedbackComment(context.Background(), businessId).DeleteGoodsFeedbackCommentRequest(deleteGoodsFeedbackCommentRequest).SourceType(sourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsFeedbackAPI.DeleteGoodsFeedbackComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGoodsFeedbackComment`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsFeedbackAPI.DeleteGoodsFeedbackComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGoodsFeedbackCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteGoodsFeedbackCommentRequest** | [**DeleteGoodsFeedbackCommentRequest**](DeleteGoodsFeedbackCommentRequest.md) |  | 
 **sourceType** | [**SourceType**](SourceType.md) | Признак типа кабинета, от имени которого вызывается метод: {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  - &#x60;SELLER&#x60; — продавец.  {% endif %}  - &#x60;ADVERTISER&#x60; — рекламодатель.  {% if audience &#x3D;&#x3D; \&quot;advertiser\&quot; %}  {% note info \&quot;Обязательно указывайте sourceType&#x3D;ADVERTISER в каждом запросе.\&quot; %}     {% endnote %}  {% endif %}  | 

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


## GetGoodsFeedbackComments

> GetGoodsFeedbackCommentsResponse GetGoodsFeedbackComments(ctx, businessId).GetGoodsFeedbackCommentsRequest(getGoodsFeedbackCommentsRequest).PageToken(pageToken).Limit(limit).SourceType(sourceType).Execute()

Получение комментариев к отзыву



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
	getGoodsFeedbackCommentsRequest := *openapiclient.NewGetGoodsFeedbackCommentsRequest() // GetGoodsFeedbackCommentsRequest | 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 25)
	sourceType := openapiclient.SourceType("SELLER") // SourceType | Признак типа кабинета, от имени которого вызывается метод: {% if audience == \"partner\" %}  - `SELLER` — продавец.  {% endif %}  - `ADVERTISER` — рекламодатель.  {% if audience == \"advertiser\" %}  {% note info \"Обязательно указывайте sourceType=ADVERTISER в каждом запросе.\" %}     {% endnote %}  {% endif %}  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsFeedbackAPI.GetGoodsFeedbackComments(context.Background(), businessId).GetGoodsFeedbackCommentsRequest(getGoodsFeedbackCommentsRequest).PageToken(pageToken).Limit(limit).SourceType(sourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsFeedbackAPI.GetGoodsFeedbackComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsFeedbackComments`: GetGoodsFeedbackCommentsResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsFeedbackAPI.GetGoodsFeedbackComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsFeedbackCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getGoodsFeedbackCommentsRequest** | [**GetGoodsFeedbackCommentsRequest**](GetGoodsFeedbackCommentsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 25]
 **sourceType** | [**SourceType**](SourceType.md) | Признак типа кабинета, от имени которого вызывается метод: {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  - &#x60;SELLER&#x60; — продавец.  {% endif %}  - &#x60;ADVERTISER&#x60; — рекламодатель.  {% if audience &#x3D;&#x3D; \&quot;advertiser\&quot; %}  {% note info \&quot;Обязательно указывайте sourceType&#x3D;ADVERTISER в каждом запросе.\&quot; %}     {% endnote %}  {% endif %}  | 

### Return type

[**GetGoodsFeedbackCommentsResponse**](GetGoodsFeedbackCommentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoodsFeedbacks

> GetGoodsFeedbackResponse GetGoodsFeedbacks(ctx, businessId).PageToken(pageToken).Limit(limit).GetGoodsFeedbackRequest(getGoodsFeedbackRequest).Execute()

Получение отзывов о товарах продавца



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
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 25)
	getGoodsFeedbackRequest := *openapiclient.NewGetGoodsFeedbackRequest() // GetGoodsFeedbackRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsFeedbackAPI.GetGoodsFeedbacks(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetGoodsFeedbackRequest(getGoodsFeedbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsFeedbackAPI.GetGoodsFeedbacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsFeedbacks`: GetGoodsFeedbackResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsFeedbackAPI.GetGoodsFeedbacks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsFeedbacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 25]
 **getGoodsFeedbackRequest** | [**GetGoodsFeedbackRequest**](GetGoodsFeedbackRequest.md) |  | 

### Return type

[**GetGoodsFeedbackResponse**](GetGoodsFeedbackResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoodsFeedbacksUrbanads

> GetGoodsFeedbackUrbanadsResponse GetGoodsFeedbacksUrbanads(ctx, businessId).PageToken(pageToken).Limit(limit).SourceType(sourceType).GetGoodsFeedbackUrbanadsRequest(getGoodsFeedbackUrbanadsRequest).Execute()

Получение отзывов о товарах для рекламодателей



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
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 25)
	sourceType := openapiclient.SourceType("SELLER") // SourceType | Признак типа кабинета, от имени которого вызывается метод: {% if audience == \"partner\" %}  - `SELLER` — продавец.  {% endif %}  - `ADVERTISER` — рекламодатель.  {% if audience == \"advertiser\" %}  {% note info \"Обязательно указывайте sourceType=ADVERTISER в каждом запросе.\" %}     {% endnote %}  {% endif %}  (optional)
	getGoodsFeedbackUrbanadsRequest := *openapiclient.NewGetGoodsFeedbackUrbanadsRequest() // GetGoodsFeedbackUrbanadsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsFeedbackAPI.GetGoodsFeedbacksUrbanads(context.Background(), businessId).PageToken(pageToken).Limit(limit).SourceType(sourceType).GetGoodsFeedbackUrbanadsRequest(getGoodsFeedbackUrbanadsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsFeedbackAPI.GetGoodsFeedbacksUrbanads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsFeedbacksUrbanads`: GetGoodsFeedbackUrbanadsResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsFeedbackAPI.GetGoodsFeedbacksUrbanads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsFeedbacksUrbanadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 25]
 **sourceType** | [**SourceType**](SourceType.md) | Признак типа кабинета, от имени которого вызывается метод: {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  - &#x60;SELLER&#x60; — продавец.  {% endif %}  - &#x60;ADVERTISER&#x60; — рекламодатель.  {% if audience &#x3D;&#x3D; \&quot;advertiser\&quot; %}  {% note info \&quot;Обязательно указывайте sourceType&#x3D;ADVERTISER в каждом запросе.\&quot; %}     {% endnote %}  {% endif %}  | 
 **getGoodsFeedbackUrbanadsRequest** | [**GetGoodsFeedbackUrbanadsRequest**](GetGoodsFeedbackUrbanadsRequest.md) |  | 

### Return type

[**GetGoodsFeedbackUrbanadsResponse**](GetGoodsFeedbackUrbanadsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipGoodsFeedbacksReaction

> EmptyApiResponse SkipGoodsFeedbacksReaction(ctx, businessId).SkipGoodsFeedbackReactionRequest(skipGoodsFeedbackReactionRequest).SourceType(sourceType).Execute()

Пропуск реакции на отзывы



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
	skipGoodsFeedbackReactionRequest := *openapiclient.NewSkipGoodsFeedbackReactionRequest([]int64{int64(123)}) // SkipGoodsFeedbackReactionRequest | 
	sourceType := openapiclient.SourceType("SELLER") // SourceType | Признак типа кабинета, от имени которого вызывается метод: {% if audience == \"partner\" %}  - `SELLER` — продавец.  {% endif %}  - `ADVERTISER` — рекламодатель.  {% if audience == \"advertiser\" %}  {% note info \"Обязательно указывайте sourceType=ADVERTISER в каждом запросе.\" %}     {% endnote %}  {% endif %}  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsFeedbackAPI.SkipGoodsFeedbacksReaction(context.Background(), businessId).SkipGoodsFeedbackReactionRequest(skipGoodsFeedbackReactionRequest).SourceType(sourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsFeedbackAPI.SkipGoodsFeedbacksReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkipGoodsFeedbacksReaction`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsFeedbackAPI.SkipGoodsFeedbacksReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkipGoodsFeedbacksReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipGoodsFeedbackReactionRequest** | [**SkipGoodsFeedbackReactionRequest**](SkipGoodsFeedbackReactionRequest.md) |  | 
 **sourceType** | [**SourceType**](SourceType.md) | Признак типа кабинета, от имени которого вызывается метод: {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  - &#x60;SELLER&#x60; — продавец.  {% endif %}  - &#x60;ADVERTISER&#x60; — рекламодатель.  {% if audience &#x3D;&#x3D; \&quot;advertiser\&quot; %}  {% note info \&quot;Обязательно указывайте sourceType&#x3D;ADVERTISER в каждом запросе.\&quot; %}     {% endnote %}  {% endif %}  | 

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


## UpdateGoodsFeedbackComment

> UpdateGoodsFeedbackCommentResponse UpdateGoodsFeedbackComment(ctx, businessId).UpdateGoodsFeedbackCommentRequest(updateGoodsFeedbackCommentRequest).SourceType(sourceType).Execute()

Добавление нового или изменение созданного комментария



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
	updateGoodsFeedbackCommentRequest := *openapiclient.NewUpdateGoodsFeedbackCommentRequest(int64(123), *openapiclient.NewUpdateGoodsFeedbackCommentDTO("Text_example")) // UpdateGoodsFeedbackCommentRequest | 
	sourceType := openapiclient.SourceType("SELLER") // SourceType | Признак типа кабинета, от имени которого вызывается метод: {% if audience == \"partner\" %}  - `SELLER` — продавец.  {% endif %}  - `ADVERTISER` — рекламодатель.  {% if audience == \"advertiser\" %}  {% note info \"Обязательно указывайте sourceType=ADVERTISER в каждом запросе.\" %}     {% endnote %}  {% endif %}  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoodsFeedbackAPI.UpdateGoodsFeedbackComment(context.Background(), businessId).UpdateGoodsFeedbackCommentRequest(updateGoodsFeedbackCommentRequest).SourceType(sourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoodsFeedbackAPI.UpdateGoodsFeedbackComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoodsFeedbackComment`: UpdateGoodsFeedbackCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `GoodsFeedbackAPI.UpdateGoodsFeedbackComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoodsFeedbackCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoodsFeedbackCommentRequest** | [**UpdateGoodsFeedbackCommentRequest**](UpdateGoodsFeedbackCommentRequest.md) |  | 
 **sourceType** | [**SourceType**](SourceType.md) | Признак типа кабинета, от имени которого вызывается метод: {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  - &#x60;SELLER&#x60; — продавец.  {% endif %}  - &#x60;ADVERTISER&#x60; — рекламодатель.  {% if audience &#x3D;&#x3D; \&quot;advertiser\&quot; %}  {% note info \&quot;Обязательно указывайте sourceType&#x3D;ADVERTISER в каждом запросе.\&quot; %}     {% endnote %}  {% endif %}  | 

### Return type

[**UpdateGoodsFeedbackCommentResponse**](UpdateGoodsFeedbackCommentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

