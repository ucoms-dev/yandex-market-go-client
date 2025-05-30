# \ReportsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateBannersStatisticsReport**](ReportsAPI.md#GenerateBannersStatisticsReport) | **Post** /reports/banners-statistics/generate | Отчет по охватному продвижению
[**GenerateBoostConsolidatedReport**](ReportsAPI.md#GenerateBoostConsolidatedReport) | **Post** /reports/boost-consolidated/generate | Отчет по бусту продаж
[**GenerateCompetitorsPositionReport**](ReportsAPI.md#GenerateCompetitorsPositionReport) | **Post** /reports/competitors-position/generate | Отчет «Конкурентная позиция»
[**GenerateGoodsFeedbackReport**](ReportsAPI.md#GenerateGoodsFeedbackReport) | **Post** /reports/goods-feedback/generate | Отчет по отзывам о товарах
[**GenerateGoodsMovementReport**](ReportsAPI.md#GenerateGoodsMovementReport) | **Post** /reports/goods-movement/generate | Отчет по движению товаров
[**GenerateGoodsRealizationReport**](ReportsAPI.md#GenerateGoodsRealizationReport) | **Post** /reports/goods-realization/generate | Отчет по реализации
[**GenerateGoodsTurnoverReport**](ReportsAPI.md#GenerateGoodsTurnoverReport) | **Post** /reports/goods-turnover/generate | Отчет по оборачиваемости
[**GenerateJewelryFiscalReport**](ReportsAPI.md#GenerateJewelryFiscalReport) | **Post** /reports/jewelry-fiscal/generate | Отчет по заказам с ювелирными изделиями
[**GenerateMassOrderLabelsReport**](ReportsAPI.md#GenerateMassOrderLabelsReport) | **Post** /reports/documents/labels/generate | Готовые ярлыки‑наклейки на все коробки в нескольких заказах
[**GeneratePricesReport**](ReportsAPI.md#GeneratePricesReport) | **Post** /reports/prices/generate | Отчет «Цены на рынке»
[**GenerateSalesGeographyReport**](ReportsAPI.md#GenerateSalesGeographyReport) | **Post** /reports/sales-geography/generate | Отчет по географии продаж
[**GenerateShelfsStatisticsReport**](ReportsAPI.md#GenerateShelfsStatisticsReport) | **Post** /reports/shelf-statistics/generate | Отчет по полкам
[**GenerateShipmentListDocumentReport**](ReportsAPI.md#GenerateShipmentListDocumentReport) | **Post** /reports/documents/shipment-list/generate | Получение листа сборки
[**GenerateShowsBoostReport**](ReportsAPI.md#GenerateShowsBoostReport) | **Post** /reports/shows-boost/generate | Отчет по бусту показов
[**GenerateShowsSalesReport**](ReportsAPI.md#GenerateShowsSalesReport) | **Post** /reports/shows-sales/generate | Отчет «Аналитика продаж»
[**GenerateStocksOnWarehousesReport**](ReportsAPI.md#GenerateStocksOnWarehousesReport) | **Post** /reports/stocks-on-warehouses/generate | Отчет по остаткам на складах
[**GenerateUnitedMarketplaceServicesReport**](ReportsAPI.md#GenerateUnitedMarketplaceServicesReport) | **Post** /reports/united-marketplace-services/generate | Отчет по стоимости услуг
[**GenerateUnitedNettingReport**](ReportsAPI.md#GenerateUnitedNettingReport) | **Post** /reports/united-netting/generate | Отчет по платежам
[**GenerateUnitedOrdersReport**](ReportsAPI.md#GenerateUnitedOrdersReport) | **Post** /reports/united-orders/generate | Отчет по заказам
[**GenerateUnitedReturnsReport**](ReportsAPI.md#GenerateUnitedReturnsReport) | **Post** /reports/united-returns/generate | Отчет по невыкупам и возвратам
[**GetReportInfo**](ReportsAPI.md#GetReportInfo) | **Get** /reports/info/{reportId} | Получение заданного отчета



## GenerateBannersStatisticsReport

> GenerateReportResponse GenerateBannersStatisticsReport(ctx).GenerateBannersStatisticsRequest(generateBannersStatisticsRequest).Format(format).Execute()

Отчет по охватному продвижению



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateBannersStatisticsRequest := *openapiclient.NewGenerateBannersStatisticsRequest(int64(123), time.Now(), time.Now()) // GenerateBannersStatisticsRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateBannersStatisticsReport(context.Background()).GenerateBannersStatisticsRequest(generateBannersStatisticsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateBannersStatisticsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBannersStatisticsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateBannersStatisticsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBannersStatisticsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateBannersStatisticsRequest** | [**GenerateBannersStatisticsRequest**](GenerateBannersStatisticsRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateBoostConsolidatedReport

> GenerateReportResponse GenerateBoostConsolidatedReport(ctx).GenerateBoostConsolidatedRequest(generateBoostConsolidatedRequest).Format(format).Execute()

Отчет по бусту продаж



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateBoostConsolidatedRequest := *openapiclient.NewGenerateBoostConsolidatedRequest(int64(123), time.Now(), time.Now()) // GenerateBoostConsolidatedRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateBoostConsolidatedReport(context.Background()).GenerateBoostConsolidatedRequest(generateBoostConsolidatedRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateBoostConsolidatedReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBoostConsolidatedReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateBoostConsolidatedReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateBoostConsolidatedReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateBoostConsolidatedRequest** | [**GenerateBoostConsolidatedRequest**](GenerateBoostConsolidatedRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCompetitorsPositionReport

> GenerateReportResponse GenerateCompetitorsPositionReport(ctx).GenerateCompetitorsPositionReportRequest(generateCompetitorsPositionReportRequest).Format(format).Execute()

Отчет «Конкурентная позиция»



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateCompetitorsPositionReportRequest := *openapiclient.NewGenerateCompetitorsPositionReportRequest(int64(123), int64(123), time.Now(), time.Now()) // GenerateCompetitorsPositionReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateCompetitorsPositionReport(context.Background()).GenerateCompetitorsPositionReportRequest(generateCompetitorsPositionReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateCompetitorsPositionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCompetitorsPositionReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateCompetitorsPositionReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCompetitorsPositionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateCompetitorsPositionReportRequest** | [**GenerateCompetitorsPositionReportRequest**](GenerateCompetitorsPositionReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateGoodsFeedbackReport

> GenerateReportResponse GenerateGoodsFeedbackReport(ctx).GenerateGoodsFeedbackRequest(generateGoodsFeedbackRequest).Format(format).Execute()

Отчет по отзывам о товарах



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
	generateGoodsFeedbackRequest := *openapiclient.NewGenerateGoodsFeedbackRequest(int64(123)) // GenerateGoodsFeedbackRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateGoodsFeedbackReport(context.Background()).GenerateGoodsFeedbackRequest(generateGoodsFeedbackRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateGoodsFeedbackReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateGoodsFeedbackReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateGoodsFeedbackReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGoodsFeedbackReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateGoodsFeedbackRequest** | [**GenerateGoodsFeedbackRequest**](GenerateGoodsFeedbackRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateGoodsMovementReport

> GenerateReportResponse GenerateGoodsMovementReport(ctx).GenerateGoodsMovementReportRequest(generateGoodsMovementReportRequest).Format(format).Execute()

Отчет по движению товаров



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateGoodsMovementReportRequest := *openapiclient.NewGenerateGoodsMovementReportRequest(int64(123), time.Now(), time.Now()) // GenerateGoodsMovementReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateGoodsMovementReport(context.Background()).GenerateGoodsMovementReportRequest(generateGoodsMovementReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateGoodsMovementReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateGoodsMovementReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateGoodsMovementReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGoodsMovementReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateGoodsMovementReportRequest** | [**GenerateGoodsMovementReportRequest**](GenerateGoodsMovementReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateGoodsRealizationReport

> GenerateReportResponse GenerateGoodsRealizationReport(ctx).GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest).Format(format).Execute()

Отчет по реализации



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
	generateGoodsRealizationReportRequest := *openapiclient.NewGenerateGoodsRealizationReportRequest(int64(123), int32(2025), int32(12)) // GenerateGoodsRealizationReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateGoodsRealizationReport(context.Background()).GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateGoodsRealizationReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateGoodsRealizationReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateGoodsRealizationReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGoodsRealizationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateGoodsRealizationReportRequest** | [**GenerateGoodsRealizationReportRequest**](GenerateGoodsRealizationReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateGoodsTurnoverReport

> GenerateReportResponse GenerateGoodsTurnoverReport(ctx).GenerateGoodsTurnoverRequest(generateGoodsTurnoverRequest).Format(format).Execute()

Отчет по оборачиваемости



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
	generateGoodsTurnoverRequest := *openapiclient.NewGenerateGoodsTurnoverRequest(int64(123)) // GenerateGoodsTurnoverRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateGoodsTurnoverReport(context.Background()).GenerateGoodsTurnoverRequest(generateGoodsTurnoverRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateGoodsTurnoverReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateGoodsTurnoverReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateGoodsTurnoverReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateGoodsTurnoverReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateGoodsTurnoverRequest** | [**GenerateGoodsTurnoverRequest**](GenerateGoodsTurnoverRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateJewelryFiscalReport

> GenerateReportResponse GenerateJewelryFiscalReport(ctx).GenerateJewelryFiscalReportRequest(generateJewelryFiscalReportRequest).Format(format).Execute()

Отчет по заказам с ювелирными изделиями



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateJewelryFiscalReportRequest := *openapiclient.NewGenerateJewelryFiscalReportRequest(int64(123), time.Now(), time.Now()) // GenerateJewelryFiscalReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateJewelryFiscalReport(context.Background()).GenerateJewelryFiscalReportRequest(generateJewelryFiscalReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateJewelryFiscalReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateJewelryFiscalReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateJewelryFiscalReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateJewelryFiscalReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateJewelryFiscalReportRequest** | [**GenerateJewelryFiscalReportRequest**](GenerateJewelryFiscalReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateMassOrderLabelsReport

> GenerateReportResponse GenerateMassOrderLabelsReport(ctx).GenerateMassOrderLabelsRequest(generateMassOrderLabelsRequest).Format(format).Execute()

Готовые ярлыки‑наклейки на все коробки в нескольких заказах



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
	generateMassOrderLabelsRequest := *openapiclient.NewGenerateMassOrderLabelsRequest(int64(123), []int64{int64(123)}) // GenerateMassOrderLabelsRequest | 
	format := openapiclient.PageFormatType("A9_HORIZONTALLY") // PageFormatType | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A7. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateMassOrderLabelsReport(context.Background()).GenerateMassOrderLabelsRequest(generateMassOrderLabelsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateMassOrderLabelsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateMassOrderLabelsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateMassOrderLabelsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMassOrderLabelsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateMassOrderLabelsRequest** | [**GenerateMassOrderLabelsRequest**](GenerateMassOrderLabelsRequest.md) |  | 
 **format** | [**PageFormatType**](PageFormatType.md) | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A7. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePricesReport

> GenerateReportResponse GeneratePricesReport(ctx).GeneratePricesReportRequest(generatePricesReportRequest).Format(format).Execute()

Отчет «Цены на рынке»



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
	generatePricesReportRequest := *openapiclient.NewGeneratePricesReportRequest() // GeneratePricesReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GeneratePricesReport(context.Background()).GeneratePricesReportRequest(generatePricesReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GeneratePricesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeneratePricesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GeneratePricesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePricesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generatePricesReportRequest** | [**GeneratePricesReportRequest**](GeneratePricesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSalesGeographyReport

> GenerateReportResponse GenerateSalesGeographyReport(ctx).GenerateSalesGeographyRequest(generateSalesGeographyRequest).Format(format).Execute()

Отчет по географии продаж



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateSalesGeographyRequest := *openapiclient.NewGenerateSalesGeographyRequest(int64(123), time.Now(), time.Now()) // GenerateSalesGeographyRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateSalesGeographyReport(context.Background()).GenerateSalesGeographyRequest(generateSalesGeographyRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateSalesGeographyReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSalesGeographyReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateSalesGeographyReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSalesGeographyReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateSalesGeographyRequest** | [**GenerateSalesGeographyRequest**](GenerateSalesGeographyRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateShelfsStatisticsReport

> GenerateReportResponse GenerateShelfsStatisticsReport(ctx).GenerateShelfsStatisticsRequest(generateShelfsStatisticsRequest).Format(format).Execute()

Отчет по полкам



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateShelfsStatisticsRequest := *openapiclient.NewGenerateShelfsStatisticsRequest(int64(123), time.Now(), time.Now(), openapiclient.StatisticsAttributionType("CLICKS")) // GenerateShelfsStatisticsRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateShelfsStatisticsReport(context.Background()).GenerateShelfsStatisticsRequest(generateShelfsStatisticsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateShelfsStatisticsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateShelfsStatisticsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateShelfsStatisticsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateShelfsStatisticsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateShelfsStatisticsRequest** | [**GenerateShelfsStatisticsRequest**](GenerateShelfsStatisticsRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateShipmentListDocumentReport

> GenerateReportResponse GenerateShipmentListDocumentReport(ctx).GenerateShipmentListDocumentReportRequest(generateShipmentListDocumentReportRequest).Execute()

Получение листа сборки



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
	generateShipmentListDocumentReportRequest := *openapiclient.NewGenerateShipmentListDocumentReportRequest(int64(123)) // GenerateShipmentListDocumentReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateShipmentListDocumentReport(context.Background()).GenerateShipmentListDocumentReportRequest(generateShipmentListDocumentReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateShipmentListDocumentReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateShipmentListDocumentReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateShipmentListDocumentReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateShipmentListDocumentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateShipmentListDocumentReportRequest** | [**GenerateShipmentListDocumentReportRequest**](GenerateShipmentListDocumentReportRequest.md) |  | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateShowsBoostReport

> GenerateReportResponse GenerateShowsBoostReport(ctx).GenerateShowsBoostRequest(generateShowsBoostRequest).Format(format).Execute()

Отчет по бусту показов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateShowsBoostRequest := *openapiclient.NewGenerateShowsBoostRequest(int64(123), time.Now(), time.Now(), openapiclient.StatisticsAttributionType("CLICKS")) // GenerateShowsBoostRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateShowsBoostReport(context.Background()).GenerateShowsBoostRequest(generateShowsBoostRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateShowsBoostReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateShowsBoostReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateShowsBoostReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateShowsBoostReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateShowsBoostRequest** | [**GenerateShowsBoostRequest**](GenerateShowsBoostRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateShowsSalesReport

> GenerateReportResponse GenerateShowsSalesReport(ctx).GenerateShowsSalesReportRequest(generateShowsSalesReportRequest).Format(format).Execute()

Отчет «Аналитика продаж»



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateShowsSalesReportRequest := *openapiclient.NewGenerateShowsSalesReportRequest(time.Now(), time.Now(), openapiclient.ShowsSalesGroupingType("CATEGORIES")) // GenerateShowsSalesReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateShowsSalesReport(context.Background()).GenerateShowsSalesReportRequest(generateShowsSalesReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateShowsSalesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateShowsSalesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateShowsSalesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateShowsSalesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateShowsSalesReportRequest** | [**GenerateShowsSalesReportRequest**](GenerateShowsSalesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateStocksOnWarehousesReport

> GenerateReportResponse GenerateStocksOnWarehousesReport(ctx).GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest).Format(format).Execute()

Отчет по остаткам на складах



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
	generateStocksOnWarehousesReportRequest := *openapiclient.NewGenerateStocksOnWarehousesReportRequest(int64(123)) // GenerateStocksOnWarehousesReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateStocksOnWarehousesReport(context.Background()).GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateStocksOnWarehousesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateStocksOnWarehousesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateStocksOnWarehousesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateStocksOnWarehousesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateStocksOnWarehousesReportRequest** | [**GenerateStocksOnWarehousesReportRequest**](GenerateStocksOnWarehousesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUnitedMarketplaceServicesReport

> GenerateReportResponse GenerateUnitedMarketplaceServicesReport(ctx).GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest).Format(format).Language(language).Execute()

Отчет по стоимости услуг



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
	generateUnitedMarketplaceServicesReportRequest := *openapiclient.NewGenerateUnitedMarketplaceServicesReportRequest(int64(123)) // GenerateUnitedMarketplaceServicesReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")
	language := *openapiclient.NewReportLanguageType() // ReportLanguageType | Язык отчета. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateUnitedMarketplaceServicesReport(context.Background()).GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest).Format(format).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateUnitedMarketplaceServicesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedMarketplaceServicesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateUnitedMarketplaceServicesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUnitedMarketplaceServicesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateUnitedMarketplaceServicesReportRequest** | [**GenerateUnitedMarketplaceServicesReportRequest**](GenerateUnitedMarketplaceServicesReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]
 **language** | [**ReportLanguageType**](ReportLanguageType.md) | Язык отчета. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUnitedNettingReport

> GenerateReportResponse GenerateUnitedNettingReport(ctx).GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest).Format(format).Language(language).Execute()

Отчет по платежам



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
	generateUnitedNettingReportRequest := *openapiclient.NewGenerateUnitedNettingReportRequest(int64(123)) // GenerateUnitedNettingReportRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")
	language := *openapiclient.NewReportLanguageType() // ReportLanguageType | Язык отчета. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateUnitedNettingReport(context.Background()).GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest).Format(format).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateUnitedNettingReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedNettingReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateUnitedNettingReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUnitedNettingReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateUnitedNettingReportRequest** | [**GenerateUnitedNettingReportRequest**](GenerateUnitedNettingReportRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]
 **language** | [**ReportLanguageType**](ReportLanguageType.md) | Язык отчета. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUnitedOrdersReport

> GenerateReportResponse GenerateUnitedOrdersReport(ctx).GenerateUnitedOrdersRequest(generateUnitedOrdersRequest).Format(format).Language(language).Execute()

Отчет по заказам



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateUnitedOrdersRequest := *openapiclient.NewGenerateUnitedOrdersRequest(int64(123), time.Now(), time.Now()) // GenerateUnitedOrdersRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")
	language := *openapiclient.NewReportLanguageType() // ReportLanguageType | Язык отчета. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateUnitedOrdersReport(context.Background()).GenerateUnitedOrdersRequest(generateUnitedOrdersRequest).Format(format).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateUnitedOrdersReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedOrdersReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateUnitedOrdersReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUnitedOrdersReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateUnitedOrdersRequest** | [**GenerateUnitedOrdersRequest**](GenerateUnitedOrdersRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]
 **language** | [**ReportLanguageType**](ReportLanguageType.md) | Язык отчета. | 

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUnitedReturnsReport

> GenerateReportResponse GenerateUnitedReturnsReport(ctx).GenerateUnitedReturnsRequest(generateUnitedReturnsRequest).Format(format).Execute()

Отчет по невыкупам и возвратам



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	generateUnitedReturnsRequest := *openapiclient.NewGenerateUnitedReturnsRequest(int64(123), time.Now(), time.Now()) // GenerateUnitedReturnsRequest | 
	format := openapiclient.ReportFormatType("FILE") // ReportFormatType | Формат отчета. (optional) (default to "FILE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateUnitedReturnsReport(context.Background()).GenerateUnitedReturnsRequest(generateUnitedReturnsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateUnitedReturnsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedReturnsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateUnitedReturnsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUnitedReturnsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateUnitedReturnsRequest** | [**GenerateUnitedReturnsRequest**](GenerateUnitedReturnsRequest.md) |  | 
 **format** | [**ReportFormatType**](ReportFormatType.md) | Формат отчета. | [default to &quot;FILE&quot;]

### Return type

[**GenerateReportResponse**](GenerateReportResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportInfo

> GetReportInfoResponse GetReportInfo(ctx, reportId).Execute()

Получение заданного отчета



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
	reportId := "reportId_example" // string | Идентификатор отчета, который вы получили после запуска генерации. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportInfo(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportInfo`: GetReportInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Идентификатор отчета, который вы получили после запуска генерации.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReportInfoResponse**](GetReportInfoResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

