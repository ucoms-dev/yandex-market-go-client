# \DocumentsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocuments**](DocumentsAPI.md#CreateDocuments) | **Post** /v1/businesses/{businessId}/offers/documents/create | Создание документов
[**DeleteDocuments**](DocumentsAPI.md#DeleteDocuments) | **Post** /v1/businesses/{businessId}/offers/documents/delete | Удаление документов
[**GetDocuments**](DocumentsAPI.md#GetDocuments) | **Post** /v1/businesses/{businessId}/offers/documents | Получение документов
[**UpdateDocuments**](DocumentsAPI.md#UpdateDocuments) | **Post** /v1/businesses/{businessId}/offers/documents/update | Обновление документов



## CreateDocuments

> CreateDocumentsResponse CreateDocuments(ctx, businessId).CreateDocumentsRequest(createDocumentsRequest).Execute()

Создание документов



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
	createDocumentsRequest := *openapiclient.NewCreateDocumentsRequest([]openapiclient.DocumentWriteDTO{*openapiclient.NewDocumentWriteDTO("Number_example", openapiclient.DocumentType("CONFORMITY_DECLARATION"))}) // CreateDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.CreateDocuments(context.Background(), businessId).CreateDocumentsRequest(createDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CreateDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDocuments`: CreateDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CreateDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDocumentsRequest** | [**CreateDocumentsRequest**](CreateDocumentsRequest.md) |  | 

### Return type

[**CreateDocumentsResponse**](CreateDocumentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocuments

> EmptyApiResponse DeleteDocuments(ctx, businessId).DeleteDocumentsRequest(deleteDocumentsRequest).Execute()

Удаление документов



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
	deleteDocumentsRequest := *openapiclient.NewDeleteDocumentsRequest([]int64{int64(123)}) // DeleteDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DeleteDocuments(context.Background(), businessId).DeleteDocumentsRequest(deleteDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DeleteDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDocuments`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DeleteDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteDocumentsRequest** | [**DeleteDocumentsRequest**](DeleteDocumentsRequest.md) |  | 

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


## GetDocuments

> GetDocumentsResponse GetDocuments(ctx, businessId).GetDocumentsRequest(getDocumentsRequest).PageToken(pageToken).Limit(limit).Execute()

Получение документов



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
	getDocumentsRequest := *openapiclient.NewGetDocumentsRequest() // GetDocumentsRequest | 
	pageToken := "pageToken_example" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра `nextPageToken`, полученное при последнем запросе.  (optional)
	limit := int32(56) // int32 | {{ limit-param-description }}  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GetDocuments(context.Background(), businessId).GetDocumentsRequest(getDocumentsRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocuments`: GetDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getDocumentsRequest** | [**GetDocumentsRequest**](GetDocumentsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Передавайте значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  | 
 **limit** | **int32** | {{ limit-param-description }}  | [default to 50]

### Return type

[**GetDocumentsResponse**](GetDocumentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocuments

> UpdateDocumentsResponse UpdateDocuments(ctx, businessId).UpdateDocumentsRequest(updateDocumentsRequest).Execute()

Обновление документов



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
	updateDocumentsRequest := *openapiclient.NewUpdateDocumentsRequest([]openapiclient.UpdateDocumentDTO{*openapiclient.NewUpdateDocumentDTO("Number_example", openapiclient.DocumentType("CONFORMITY_DECLARATION"), int64(123))}) // UpdateDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.UpdateDocuments(context.Background(), businessId).UpdateDocumentsRequest(updateDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.UpdateDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDocuments`: UpdateDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.UpdateDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета.  {% if audience &#x3D;&#x3D; \&quot;partner\&quot; %}  Чтобы его узнать, воспользуйтесь запросом [GET v2/campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  {% endif %}  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDocumentsRequest** | [**UpdateDocumentsRequest**](UpdateDocumentsRequest.md) |  | 

### Return type

[**UpdateDocumentsResponse**](UpdateDocumentsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

