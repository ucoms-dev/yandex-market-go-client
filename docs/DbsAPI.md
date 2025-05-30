# \DbsAPI

All URIs are relative to *https://api.partner.market.yandex.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptOrderCancellation**](DbsAPI.md#AcceptOrderCancellation) | **Put** /campaigns/{campaignId}/orders/{orderId}/cancellation/accept | Отмена заказа покупателем
[**AddHiddenOffers**](DbsAPI.md#AddHiddenOffers) | **Post** /campaigns/{campaignId}/hidden-offers | Скрытие товаров и настройки скрытия
[**AddOffersToArchive**](DbsAPI.md#AddOffersToArchive) | **Post** /businesses/{businessId}/offer-mappings/archive | Добавление товаров в архив
[**CalculateTariffs**](DbsAPI.md#CalculateTariffs) | **Post** /tariffs/calculate | Калькулятор стоимости услуг
[**ConfirmBusinessPrices**](DbsAPI.md#ConfirmBusinessPrices) | **Post** /businesses/{businessId}/price-quarantine/confirm | Удаление товара из карантина по цене в кабинете
[**ConfirmCampaignPrices**](DbsAPI.md#ConfirmCampaignPrices) | **Post** /campaigns/{campaignId}/price-quarantine/confirm | Удаление товара из карантина по цене в магазине
[**CreateChat**](DbsAPI.md#CreateChat) | **Post** /businesses/{businessId}/chats/new | Создание нового чата с покупателем
[**CreateOutlet**](DbsAPI.md#CreateOutlet) | **Post** /campaigns/{campaignId}/outlets | Создание точки продаж
[**DeleteCampaignOffers**](DbsAPI.md#DeleteCampaignOffers) | **Post** /campaigns/{campaignId}/offers/delete | Удаление товаров из ассортимента магазина
[**DeleteGoodsFeedbackComment**](DbsAPI.md#DeleteGoodsFeedbackComment) | **Post** /businesses/{businessId}/goods-feedback/comments/delete | Удаление комментария к отзыву
[**DeleteHiddenOffers**](DbsAPI.md#DeleteHiddenOffers) | **Post** /campaigns/{campaignId}/hidden-offers/delete | Возобновление показа товаров
[**DeleteOffers**](DbsAPI.md#DeleteOffers) | **Post** /businesses/{businessId}/offer-mappings/delete | Удаление товаров из каталога
[**DeleteOffersFromArchive**](DbsAPI.md#DeleteOffersFromArchive) | **Post** /businesses/{businessId}/offer-mappings/unarchive | Удаление товаров из архива
[**DeleteOutlet**](DbsAPI.md#DeleteOutlet) | **Delete** /campaigns/{campaignId}/outlets/{outletId} | Удаление точки продаж
[**DeleteOutletLicenses**](DbsAPI.md#DeleteOutletLicenses) | **Delete** /campaigns/{campaignId}/outlets/licenses | Удаление лицензий для точек продаж
[**DeletePromoOffers**](DbsAPI.md#DeletePromoOffers) | **Post** /businesses/{businessId}/promos/offers/delete | Удаление товаров из акции
[**GenerateBannersStatisticsReport**](DbsAPI.md#GenerateBannersStatisticsReport) | **Post** /reports/banners-statistics/generate | Отчет по охватному продвижению
[**GenerateBoostConsolidatedReport**](DbsAPI.md#GenerateBoostConsolidatedReport) | **Post** /reports/boost-consolidated/generate | Отчет по бусту продаж
[**GenerateCompetitorsPositionReport**](DbsAPI.md#GenerateCompetitorsPositionReport) | **Post** /reports/competitors-position/generate | Отчет «Конкурентная позиция»
[**GenerateGoodsFeedbackReport**](DbsAPI.md#GenerateGoodsFeedbackReport) | **Post** /reports/goods-feedback/generate | Отчет по отзывам о товарах
[**GenerateGoodsRealizationReport**](DbsAPI.md#GenerateGoodsRealizationReport) | **Post** /reports/goods-realization/generate | Отчет по реализации
[**GenerateJewelryFiscalReport**](DbsAPI.md#GenerateJewelryFiscalReport) | **Post** /reports/jewelry-fiscal/generate | Отчет по заказам с ювелирными изделиями
[**GenerateMassOrderLabelsReport**](DbsAPI.md#GenerateMassOrderLabelsReport) | **Post** /reports/documents/labels/generate | Готовые ярлыки‑наклейки на все коробки в нескольких заказах
[**GenerateOrderLabel**](DbsAPI.md#GenerateOrderLabel) | **Get** /campaigns/{campaignId}/orders/{orderId}/delivery/shipments/{shipmentId}/boxes/{boxId}/label | Готовый ярлык‑наклейка для коробки в заказе
[**GenerateOrderLabels**](DbsAPI.md#GenerateOrderLabels) | **Get** /campaigns/{campaignId}/orders/{orderId}/delivery/labels | Готовые ярлыки‑наклейки на все коробки в одном заказе
[**GeneratePricesReport**](DbsAPI.md#GeneratePricesReport) | **Post** /reports/prices/generate | Отчет «Цены на рынке»
[**GenerateSalesGeographyReport**](DbsAPI.md#GenerateSalesGeographyReport) | **Post** /reports/sales-geography/generate | Отчет по географии продаж
[**GenerateShelfsStatisticsReport**](DbsAPI.md#GenerateShelfsStatisticsReport) | **Post** /reports/shelf-statistics/generate | Отчет по полкам
[**GenerateShowsBoostReport**](DbsAPI.md#GenerateShowsBoostReport) | **Post** /reports/shows-boost/generate | Отчет по бусту показов
[**GenerateShowsSalesReport**](DbsAPI.md#GenerateShowsSalesReport) | **Post** /reports/shows-sales/generate | Отчет «Аналитика продаж»
[**GenerateStocksOnWarehousesReport**](DbsAPI.md#GenerateStocksOnWarehousesReport) | **Post** /reports/stocks-on-warehouses/generate | Отчет по остаткам на складах
[**GenerateUnitedMarketplaceServicesReport**](DbsAPI.md#GenerateUnitedMarketplaceServicesReport) | **Post** /reports/united-marketplace-services/generate | Отчет по стоимости услуг
[**GenerateUnitedNettingReport**](DbsAPI.md#GenerateUnitedNettingReport) | **Post** /reports/united-netting/generate | Отчет по платежам
[**GenerateUnitedOrdersReport**](DbsAPI.md#GenerateUnitedOrdersReport) | **Post** /reports/united-orders/generate | Отчет по заказам
[**GenerateUnitedReturnsReport**](DbsAPI.md#GenerateUnitedReturnsReport) | **Post** /reports/united-returns/generate | Отчет по невыкупам и возвратам
[**GetAuthTokenInfo**](DbsAPI.md#GetAuthTokenInfo) | **Post** /auth/token | Получение информации об авторизационном токене
[**GetBidsInfoForBusiness**](DbsAPI.md#GetBidsInfoForBusiness) | **Post** /businesses/{businessId}/bids/info | Информация об установленных ставках
[**GetBidsRecommendations**](DbsAPI.md#GetBidsRecommendations) | **Post** /businesses/{businessId}/bids/recommendations | Рекомендованные ставки для заданных товаров
[**GetBusinessQuarantineOffers**](DbsAPI.md#GetBusinessQuarantineOffers) | **Post** /businesses/{businessId}/price-quarantine | Список товаров, находящихся в карантине по цене в кабинете
[**GetBusinessSettings**](DbsAPI.md#GetBusinessSettings) | **Post** /businesses/{businessId}/settings | Настройки кабинета
[**GetCampaign**](DbsAPI.md#GetCampaign) | **Get** /campaigns/{campaignId} | Информация о магазине
[**GetCampaignOffers**](DbsAPI.md#GetCampaignOffers) | **Post** /campaigns/{campaignId}/offers | Информация о товарах, которые размещены в заданном магазине
[**GetCampaignQuarantineOffers**](DbsAPI.md#GetCampaignQuarantineOffers) | **Post** /campaigns/{campaignId}/price-quarantine | Список товаров, находящихся в карантине по цене в магазине
[**GetCampaignRegion**](DbsAPI.md#GetCampaignRegion) | **Get** /campaigns/{campaignId}/region | Регион магазина
[**GetCampaignSettings**](DbsAPI.md#GetCampaignSettings) | **Get** /campaigns/{campaignId}/settings | Настройки магазина
[**GetCampaigns**](DbsAPI.md#GetCampaigns) | **Get** /campaigns | Список магазинов пользователя
[**GetCategoriesMaxSaleQuantum**](DbsAPI.md#GetCategoriesMaxSaleQuantum) | **Post** /categories/max-sale-quantum | Лимит на установку кванта продажи и минимального количества товаров в заказе
[**GetCategoriesTree**](DbsAPI.md#GetCategoriesTree) | **Post** /categories/tree | Дерево категорий
[**GetCategoryContentParameters**](DbsAPI.md#GetCategoryContentParameters) | **Post** /category/{categoryId}/parameters | Списки характеристик товаров по категориям
[**GetChat**](DbsAPI.md#GetChat) | **Get** /businesses/{businessId}/chat | Получение чата по идентификатору
[**GetChatHistory**](DbsAPI.md#GetChatHistory) | **Post** /businesses/{businessId}/chats/history | Получение истории сообщений в чате
[**GetChatMessage**](DbsAPI.md#GetChatMessage) | **Get** /businesses/{businessId}/chats/message | Получение сообщения в чате
[**GetChats**](DbsAPI.md#GetChats) | **Post** /businesses/{businessId}/chats | Получение доступных чатов
[**GetDeliveryServices**](DbsAPI.md#GetDeliveryServices) | **Get** /delivery/services | Справочник служб доставки
[**GetGoodsFeedbackComments**](DbsAPI.md#GetGoodsFeedbackComments) | **Post** /businesses/{businessId}/goods-feedback/comments | Получение комментариев к отзыву
[**GetGoodsFeedbacks**](DbsAPI.md#GetGoodsFeedbacks) | **Post** /businesses/{businessId}/goods-feedback | Получение отзывов о товарах продавца
[**GetGoodsStats**](DbsAPI.md#GetGoodsStats) | **Post** /campaigns/{campaignId}/stats/skus | Отчет по товарам
[**GetHiddenOffers**](DbsAPI.md#GetHiddenOffers) | **Get** /campaigns/{campaignId}/hidden-offers | Информация о скрытых вами товарах
[**GetModel**](DbsAPI.md#GetModel) | **Get** /models/{modelId} | Информация об одной модели
[**GetModelOffers**](DbsAPI.md#GetModelOffers) | **Get** /models/{modelId}/offers | Список предложений для одной модели
[**GetModels**](DbsAPI.md#GetModels) | **Post** /models | Информация о нескольких моделях
[**GetModelsOffers**](DbsAPI.md#GetModelsOffers) | **Post** /models/offers | Список предложений для нескольких моделей
[**GetOfferCardsContentStatus**](DbsAPI.md#GetOfferCardsContentStatus) | **Post** /businesses/{businessId}/offer-cards | Получение информации о заполненности карточек магазина
[**GetOfferMappingEntries**](DbsAPI.md#GetOfferMappingEntries) | **Get** /campaigns/{campaignId}/offer-mapping-entries | Список товаров в каталоге
[**GetOfferMappings**](DbsAPI.md#GetOfferMappings) | **Post** /businesses/{businessId}/offer-mappings | Информация о товарах в каталоге
[**GetOfferRecommendations**](DbsAPI.md#GetOfferRecommendations) | **Post** /businesses/{businessId}/offers/recommendations | Рекомендации Маркета, касающиеся цен
[**GetOrder**](DbsAPI.md#GetOrder) | **Get** /campaigns/{campaignId}/orders/{orderId} | Информация об одном заказе
[**GetOrderBusinessBuyerInfo**](DbsAPI.md#GetOrderBusinessBuyerInfo) | **Post** /campaigns/{campaignId}/orders/{orderId}/business-buyer | Информация о покупателе — юридическом лице
[**GetOrderBusinessDocumentsInfo**](DbsAPI.md#GetOrderBusinessDocumentsInfo) | **Post** /campaigns/{campaignId}/orders/{orderId}/documents | Информация о документах
[**GetOrderBuyerInfo**](DbsAPI.md#GetOrderBuyerInfo) | **Get** /campaigns/{campaignId}/orders/{orderId}/buyer | Информация о покупателе — физическом лице
[**GetOrderLabelsData**](DbsAPI.md#GetOrderLabelsData) | **Get** /campaigns/{campaignId}/orders/{orderId}/delivery/labels/data | Данные для самостоятельного изготовления ярлыков
[**GetOrders**](DbsAPI.md#GetOrders) | **Get** /campaigns/{campaignId}/orders | Информация о нескольких заказах
[**GetOrdersStats**](DbsAPI.md#GetOrdersStats) | **Post** /campaigns/{campaignId}/stats/orders | Детальная информация по заказам
[**GetOutlet**](DbsAPI.md#GetOutlet) | **Get** /campaigns/{campaignId}/outlets/{outletId} | Информация об одной точке продаж
[**GetOutletLicenses**](DbsAPI.md#GetOutletLicenses) | **Get** /campaigns/{campaignId}/outlets/licenses | Информация о лицензиях для точек продаж
[**GetOutlets**](DbsAPI.md#GetOutlets) | **Get** /campaigns/{campaignId}/outlets | Информация о нескольких точках продаж
[**GetPagedWarehouses**](DbsAPI.md#GetPagedWarehouses) | **Post** /businesses/{businessId}/warehouses | Список складов
[**GetPrices**](DbsAPI.md#GetPrices) | **Get** /campaigns/{campaignId}/offer-prices | Список цен
[**GetPricesByOfferIds**](DbsAPI.md#GetPricesByOfferIds) | **Post** /campaigns/{campaignId}/offer-prices | Просмотр цен на указанные товары в магазине
[**GetPromoOffers**](DbsAPI.md#GetPromoOffers) | **Post** /businesses/{businessId}/promos/offers | Получение списка товаров, которые участвуют или могут участвовать в акции
[**GetPromos**](DbsAPI.md#GetPromos) | **Post** /businesses/{businessId}/promos | Получение списка акций
[**GetQualityRatingDetails**](DbsAPI.md#GetQualityRatingDetails) | **Post** /campaigns/{campaignId}/ratings/quality/details | Заказы, которые повлияли на индекс качества
[**GetQualityRatings**](DbsAPI.md#GetQualityRatings) | **Post** /businesses/{businessId}/ratings/quality | Индекс качества магазинов
[**GetRegionsCodes**](DbsAPI.md#GetRegionsCodes) | **Post** /regions/countries | Список допустимых кодов стран
[**GetReportInfo**](DbsAPI.md#GetReportInfo) | **Get** /reports/info/{reportId} | Получение заданного отчета
[**GetReturn**](DbsAPI.md#GetReturn) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId} | Информация о невыкупе или возврате
[**GetReturnApplication**](DbsAPI.md#GetReturnApplication) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/application | Получение заявления на возврат
[**GetReturnPhoto**](DbsAPI.md#GetReturnPhoto) | **Get** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/{itemId}/image/{imageHash} | Получение фотографий товаров в возврате
[**GetReturns**](DbsAPI.md#GetReturns) | **Get** /campaigns/{campaignId}/returns | Список невыкупов и возвратов
[**GetStocks**](DbsAPI.md#GetStocks) | **Post** /campaigns/{campaignId}/offers/stocks | Информация об остатках и оборачиваемости
[**GetSuggestedOfferMappingEntries**](DbsAPI.md#GetSuggestedOfferMappingEntries) | **Post** /campaigns/{campaignId}/offer-mapping-entries/suggestions | Рекомендованные карточки для товаров
[**GetSuggestedOfferMappings**](DbsAPI.md#GetSuggestedOfferMappings) | **Post** /businesses/{businessId}/offer-mappings/suggestions | Просмотр карточек на Маркете, которые подходят вашим товарам
[**GetSuggestedPrices**](DbsAPI.md#GetSuggestedPrices) | **Post** /campaigns/{campaignId}/offer-prices/suggestions | Цены для продвижения товаров
[**GetWarehouses**](DbsAPI.md#GetWarehouses) | **Get** /businesses/{businessId}/warehouses | Список складов и групп складов
[**ProvideOrderDigitalCodes**](DbsAPI.md#ProvideOrderDigitalCodes) | **Post** /campaigns/{campaignId}/orders/{orderId}/deliverDigitalGoods | Передача ключей цифровых товаров
[**ProvideOrderItemIdentifiers**](DbsAPI.md#ProvideOrderItemIdentifiers) | **Put** /campaigns/{campaignId}/orders/{orderId}/identifiers | Передача кодов маркировки единиц товара
[**PutBidsForBusiness**](DbsAPI.md#PutBidsForBusiness) | **Put** /businesses/{businessId}/bids | Включение буста продаж и установка ставок
[**PutBidsForCampaign**](DbsAPI.md#PutBidsForCampaign) | **Put** /campaigns/{campaignId}/bids | Включение буста продаж и установка ставок для магазина
[**SearchModels**](DbsAPI.md#SearchModels) | **Get** /models | Поиск модели товара
[**SearchRegionChildren**](DbsAPI.md#SearchRegionChildren) | **Get** /regions/{regionId}/children | Информация о дочерних регионах
[**SearchRegionsById**](DbsAPI.md#SearchRegionsById) | **Get** /regions/{regionId} | Информация о регионе
[**SearchRegionsByName**](DbsAPI.md#SearchRegionsByName) | **Get** /regions | Поиск регионов по их имени
[**SendFileToChat**](DbsAPI.md#SendFileToChat) | **Post** /businesses/{businessId}/chats/file/send | Отправка файла в чат
[**SendMessageToChat**](DbsAPI.md#SendMessageToChat) | **Post** /businesses/{businessId}/chats/message | Отправка сообщения в чат
[**SetOrderBoxLayout**](DbsAPI.md#SetOrderBoxLayout) | **Put** /campaigns/{campaignId}/orders/{orderId}/boxes | Подготовка заказа
[**SetOrderDeliveryDate**](DbsAPI.md#SetOrderDeliveryDate) | **Put** /campaigns/{campaignId}/orders/{orderId}/delivery/date | Изменение даты доставки заказа
[**SetOrderDeliveryTrackCode**](DbsAPI.md#SetOrderDeliveryTrackCode) | **Post** /campaigns/{campaignId}/orders/{orderId}/delivery/track | Передача трек‑номера посылки
[**SetOrderShipmentBoxes**](DbsAPI.md#SetOrderShipmentBoxes) | **Put** /campaigns/{campaignId}/orders/{orderId}/delivery/shipments/{shipmentId}/boxes | Передача количества грузовых мест в заказе
[**SetReturnDecision**](DbsAPI.md#SetReturnDecision) | **Post** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision | Принятие или изменение решения по возврату
[**SkipGoodsFeedbacksReaction**](DbsAPI.md#SkipGoodsFeedbacksReaction) | **Post** /businesses/{businessId}/goods-feedback/skip-reaction | Пропуск реакции на отзывы
[**SubmitReturnDecision**](DbsAPI.md#SubmitReturnDecision) | **Post** /campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/submit | Подтверждение решения по возврату
[**UpdateBusinessPrices**](DbsAPI.md#UpdateBusinessPrices) | **Post** /businesses/{businessId}/offer-prices/updates | Установка цен на товары для всех магазинов
[**UpdateCampaignOffers**](DbsAPI.md#UpdateCampaignOffers) | **Post** /campaigns/{campaignId}/offers/update | Изменение условий продажи товаров в магазине
[**UpdateExternalOrderId**](DbsAPI.md#UpdateExternalOrderId) | **Post** /campaigns/{campaignId}/orders/{orderId}/external-id | Передача или изменение дополнительного идентификатора заказа
[**UpdateGoodsFeedbackComment**](DbsAPI.md#UpdateGoodsFeedbackComment) | **Post** /businesses/{businessId}/goods-feedback/comments/update | Добавление нового или изменение созданного комментария
[**UpdateOfferContent**](DbsAPI.md#UpdateOfferContent) | **Post** /businesses/{businessId}/offer-cards/update | Редактирование категорийных характеристик товара
[**UpdateOfferMappingEntries**](DbsAPI.md#UpdateOfferMappingEntries) | **Post** /campaigns/{campaignId}/offer-mapping-entries/updates | Добавление и редактирование товаров в каталоге
[**UpdateOfferMappings**](DbsAPI.md#UpdateOfferMappings) | **Post** /businesses/{businessId}/offer-mappings/update | Добавление товаров в каталог и изменение информации о них
[**UpdateOrderItems**](DbsAPI.md#UpdateOrderItems) | **Put** /campaigns/{campaignId}/orders/{orderId}/items | Удаление товара из заказа или уменьшение числа единиц
[**UpdateOrderStatus**](DbsAPI.md#UpdateOrderStatus) | **Put** /campaigns/{campaignId}/orders/{orderId}/status | Изменение статуса одного заказа
[**UpdateOrderStatuses**](DbsAPI.md#UpdateOrderStatuses) | **Post** /campaigns/{campaignId}/orders/status-update | Изменение статусов нескольких заказов
[**UpdateOrderStorageLimit**](DbsAPI.md#UpdateOrderStorageLimit) | **Put** /campaigns/{campaignId}/orders/{orderId}/delivery/storage-limit | Продление срока хранения заказа
[**UpdateOutlet**](DbsAPI.md#UpdateOutlet) | **Put** /campaigns/{campaignId}/outlets/{outletId} | Изменение информации о точке продаж
[**UpdateOutletLicenses**](DbsAPI.md#UpdateOutletLicenses) | **Post** /campaigns/{campaignId}/outlets/licenses | Создание и изменение лицензий для точек продаж
[**UpdatePrices**](DbsAPI.md#UpdatePrices) | **Post** /campaigns/{campaignId}/offer-prices/updates | Установка цен на товары в конкретном магазине
[**UpdatePromoOffers**](DbsAPI.md#UpdatePromoOffers) | **Post** /businesses/{businessId}/promos/offers/update | Добавление товаров в акцию или изменение их цен
[**UpdateStocks**](DbsAPI.md#UpdateStocks) | **Put** /campaigns/{campaignId}/offers/stocks | Передача информации об остатках
[**UpdateWarehouseStatus**](DbsAPI.md#UpdateWarehouseStatus) | **Post** /campaigns/{campaignId}/warehouse/status | Изменение статуса склада



## AcceptOrderCancellation

> EmptyApiResponse AcceptOrderCancellation(ctx, campaignId, orderId).AcceptOrderCancellationRequest(acceptOrderCancellationRequest).Execute()

Отмена заказа покупателем



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	acceptOrderCancellationRequest := *openapiclient.NewAcceptOrderCancellationRequest(false) // AcceptOrderCancellationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.AcceptOrderCancellation(context.Background(), campaignId, orderId).AcceptOrderCancellationRequest(acceptOrderCancellationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.AcceptOrderCancellation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptOrderCancellation`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.AcceptOrderCancellation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptOrderCancellationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptOrderCancellationRequest** | [**AcceptOrderCancellationRequest**](AcceptOrderCancellationRequest.md) |  | 

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


## AddHiddenOffers

> EmptyApiResponse AddHiddenOffers(ctx, campaignId).AddHiddenOffersRequest(addHiddenOffersRequest).Execute()

Скрытие товаров и настройки скрытия



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
	addHiddenOffersRequest := *openapiclient.NewAddHiddenOffersRequest([]openapiclient.HiddenOfferDTO{*openapiclient.NewHiddenOfferDTO("OfferId_example")}) // AddHiddenOffersRequest | Запрос на скрытие оферов.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.AddHiddenOffers(context.Background(), campaignId).AddHiddenOffersRequest(addHiddenOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.AddHiddenOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddHiddenOffers`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.AddHiddenOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddHiddenOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addHiddenOffersRequest** | [**AddHiddenOffersRequest**](AddHiddenOffersRequest.md) | Запрос на скрытие оферов. | 

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	addOffersToArchiveRequest := *openapiclient.NewAddOffersToArchiveRequest([]string{"OfferIds_example"}) // AddOffersToArchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.AddOffersToArchive(context.Background(), businessId).AddOffersToArchiveRequest(addOffersToArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.AddOffersToArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOffersToArchive`: AddOffersToArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.AddOffersToArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

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


## CalculateTariffs

> CalculateTariffsResponse CalculateTariffs(ctx).CalculateTariffsRequest(calculateTariffsRequest).Execute()

Калькулятор стоимости услуг



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
	calculateTariffsRequest := *openapiclient.NewCalculateTariffsRequest(*openapiclient.NewCalculateTariffsParametersDTO(), []openapiclient.CalculateTariffsOfferDTO{*openapiclient.NewCalculateTariffsOfferDTO(int64(123), float32(123), float32(123), float32(123), float32(123), float32(123))}) // CalculateTariffsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.CalculateTariffs(context.Background()).CalculateTariffsRequest(calculateTariffsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.CalculateTariffs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateTariffs`: CalculateTariffsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.CalculateTariffs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculateTariffsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculateTariffsRequest** | [**CalculateTariffsRequest**](CalculateTariffsRequest.md) |  | 

### Return type

[**CalculateTariffsResponse**](CalculateTariffsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmBusinessPrices

> EmptyApiResponse ConfirmBusinessPrices(ctx, businessId).ConfirmPricesRequest(confirmPricesRequest).Execute()

Удаление товара из карантина по цене в кабинете



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	confirmPricesRequest := *openapiclient.NewConfirmPricesRequest([]string{"OfferIds_example"}) // ConfirmPricesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.ConfirmBusinessPrices(context.Background(), businessId).ConfirmPricesRequest(confirmPricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.ConfirmBusinessPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmBusinessPrices`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.ConfirmBusinessPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmBusinessPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmPricesRequest** | [**ConfirmPricesRequest**](ConfirmPricesRequest.md) |  | 

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


## ConfirmCampaignPrices

> EmptyApiResponse ConfirmCampaignPrices(ctx, campaignId).ConfirmPricesRequest(confirmPricesRequest).Execute()

Удаление товара из карантина по цене в магазине



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
	confirmPricesRequest := *openapiclient.NewConfirmPricesRequest([]string{"OfferIds_example"}) // ConfirmPricesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.ConfirmCampaignPrices(context.Background(), campaignId).ConfirmPricesRequest(confirmPricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.ConfirmCampaignPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmCampaignPrices`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.ConfirmCampaignPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmCampaignPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmPricesRequest** | [**ConfirmPricesRequest**](ConfirmPricesRequest.md) |  | 

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


## CreateChat

> CreateChatResponse CreateChat(ctx, businessId).CreateChatRequest(createChatRequest).Execute()

Создание нового чата с покупателем



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	createChatRequest := *openapiclient.NewCreateChatRequest(int64(123)) // CreateChatRequest | description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.CreateChat(context.Background(), businessId).CreateChatRequest(createChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.CreateChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChat`: CreateChatResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.CreateChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createChatRequest** | [**CreateChatRequest**](CreateChatRequest.md) | description | 

### Return type

[**CreateChatResponse**](CreateChatResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOutlet

> CreateOutletResponse CreateOutlet(ctx, campaignId).ChangeOutletRequest(changeOutletRequest).Execute()

Создание точки продаж



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
	changeOutletRequest := *openapiclient.NewChangeOutletRequest("Name_example", openapiclient.OutletType("DEPOT"), *openapiclient.NewOutletAddressDTO(int64(123)), []string{"Phones_example"}, *openapiclient.NewOutletWorkingScheduleDTO([]openapiclient.OutletWorkingScheduleItemDTO{*openapiclient.NewOutletWorkingScheduleItemDTO(openapiclient.DayOfWeekType("MONDAY"), openapiclient.DayOfWeekType("MONDAY"), "09:59", "23:59")})) // ChangeOutletRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.CreateOutlet(context.Background(), campaignId).ChangeOutletRequest(changeOutletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.CreateOutlet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOutlet`: CreateOutletResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.CreateOutlet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOutletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeOutletRequest** | [**ChangeOutletRequest**](ChangeOutletRequest.md) |  | 

### Return type

[**CreateOutletResponse**](CreateOutletResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignOffers

> DeleteCampaignOffersResponse DeleteCampaignOffers(ctx, campaignId).DeleteCampaignOffersRequest(deleteCampaignOffersRequest).Execute()

Удаление товаров из ассортимента магазина



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
	deleteCampaignOffersRequest := *openapiclient.NewDeleteCampaignOffersRequest([]string{"OfferIds_example"}) // DeleteCampaignOffersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeleteCampaignOffers(context.Background(), campaignId).DeleteCampaignOffersRequest(deleteCampaignOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeleteCampaignOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCampaignOffers`: DeleteCampaignOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeleteCampaignOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteCampaignOffersRequest** | [**DeleteCampaignOffersRequest**](DeleteCampaignOffersRequest.md) |  | 

### Return type

[**DeleteCampaignOffersResponse**](DeleteCampaignOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGoodsFeedbackComment

> EmptyApiResponse DeleteGoodsFeedbackComment(ctx, businessId).DeleteGoodsFeedbackCommentRequest(deleteGoodsFeedbackCommentRequest).Execute()

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	deleteGoodsFeedbackCommentRequest := *openapiclient.NewDeleteGoodsFeedbackCommentRequest(int64(123)) // DeleteGoodsFeedbackCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeleteGoodsFeedbackComment(context.Background(), businessId).DeleteGoodsFeedbackCommentRequest(deleteGoodsFeedbackCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeleteGoodsFeedbackComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGoodsFeedbackComment`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeleteGoodsFeedbackComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGoodsFeedbackCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteGoodsFeedbackCommentRequest** | [**DeleteGoodsFeedbackCommentRequest**](DeleteGoodsFeedbackCommentRequest.md) |  | 

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


## DeleteHiddenOffers

> EmptyApiResponse DeleteHiddenOffers(ctx, campaignId).DeleteHiddenOffersRequest(deleteHiddenOffersRequest).Execute()

Возобновление показа товаров



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
	deleteHiddenOffersRequest := *openapiclient.NewDeleteHiddenOffersRequest([]openapiclient.HiddenOfferDTO{*openapiclient.NewHiddenOfferDTO("OfferId_example")}) // DeleteHiddenOffersRequest | Запрос на возобновление показа оферов.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeleteHiddenOffers(context.Background(), campaignId).DeleteHiddenOffersRequest(deleteHiddenOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeleteHiddenOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteHiddenOffers`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeleteHiddenOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHiddenOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteHiddenOffersRequest** | [**DeleteHiddenOffersRequest**](DeleteHiddenOffersRequest.md) | Запрос на возобновление показа оферов. | 

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	deleteOffersRequest := *openapiclient.NewDeleteOffersRequest([]string{"OfferIds_example"}) // DeleteOffersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeleteOffers(context.Background(), businessId).DeleteOffersRequest(deleteOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeleteOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOffers`: DeleteOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeleteOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	deleteOffersFromArchiveRequest := *openapiclient.NewDeleteOffersFromArchiveRequest([]string{"OfferIds_example"}) // DeleteOffersFromArchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeleteOffersFromArchive(context.Background(), businessId).DeleteOffersFromArchiveRequest(deleteOffersFromArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeleteOffersFromArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOffersFromArchive`: DeleteOffersFromArchiveResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeleteOffersFromArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

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


## DeleteOutlet

> EmptyApiResponse DeleteOutlet(ctx, campaignId, outletId).Execute()

Удаление точки продаж



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
	outletId := int64(789) // int64 | Идентификатор точки продаж.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeleteOutlet(context.Background(), campaignId, outletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeleteOutlet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOutlet`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeleteOutlet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**outletId** | **int64** | Идентификатор точки продаж. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutletLicenses

> EmptyApiResponse DeleteOutletLicenses(ctx, campaignId).Ids(ids).Execute()

Удаление лицензий для точек продаж



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
	ids := []int64{int64(123)} // []int64 | Список идентификаторов лицензий для удаления.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeleteOutletLicenses(context.Background(), campaignId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeleteOutletLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOutletLicenses`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeleteOutletLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutletLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]int64** | Список идентификаторов лицензий для удаления. | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePromoOffers

> DeletePromoOffersResponse DeletePromoOffers(ctx, businessId).DeletePromoOffersRequest(deletePromoOffersRequest).Execute()

Удаление товаров из акции



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	deletePromoOffersRequest := *openapiclient.NewDeletePromoOffersRequest("PromoId_example") // DeletePromoOffersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.DeletePromoOffers(context.Background(), businessId).DeletePromoOffersRequest(deletePromoOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.DeletePromoOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePromoOffers`: DeletePromoOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.DeletePromoOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePromoOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deletePromoOffersRequest** | [**DeletePromoOffersRequest**](DeletePromoOffersRequest.md) |  | 

### Return type

[**DeletePromoOffersResponse**](DeletePromoOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.DbsAPI.GenerateBannersStatisticsReport(context.Background()).GenerateBannersStatisticsRequest(generateBannersStatisticsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateBannersStatisticsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBannersStatisticsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateBannersStatisticsReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateBoostConsolidatedReport(context.Background()).GenerateBoostConsolidatedRequest(generateBoostConsolidatedRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateBoostConsolidatedReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateBoostConsolidatedReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateBoostConsolidatedReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateCompetitorsPositionReport(context.Background()).GenerateCompetitorsPositionReportRequest(generateCompetitorsPositionReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateCompetitorsPositionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCompetitorsPositionReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateCompetitorsPositionReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateGoodsFeedbackReport(context.Background()).GenerateGoodsFeedbackRequest(generateGoodsFeedbackRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateGoodsFeedbackReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateGoodsFeedbackReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateGoodsFeedbackReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateGoodsRealizationReport(context.Background()).GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateGoodsRealizationReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateGoodsRealizationReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateGoodsRealizationReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateJewelryFiscalReport(context.Background()).GenerateJewelryFiscalReportRequest(generateJewelryFiscalReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateJewelryFiscalReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateJewelryFiscalReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateJewelryFiscalReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateMassOrderLabelsReport(context.Background()).GenerateMassOrderLabelsRequest(generateMassOrderLabelsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateMassOrderLabelsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateMassOrderLabelsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateMassOrderLabelsReport`: %v\n", resp)
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


## GenerateOrderLabel

> *os.File GenerateOrderLabel(ctx, campaignId, orderId, shipmentId, boxId).Format(format).Execute()

Готовый ярлык‑наклейка для коробки в заказе



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	shipmentId := int64(789) // int64 | Идентификатор грузоместа.
	boxId := int64(789) // int64 | Идентификатор коробки.
	format := openapiclient.PageFormatType("A9_HORIZONTALLY") // PageFormatType | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A7. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GenerateOrderLabel(context.Background(), campaignId, orderId, shipmentId, boxId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateOrderLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateOrderLabel`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateOrderLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 
**shipmentId** | **int64** | Идентификатор грузоместа. | 
**boxId** | **int64** | Идентификатор коробки. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOrderLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | [**PageFormatType**](PageFormatType.md) | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A7. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOrderLabels

> *os.File GenerateOrderLabels(ctx, campaignId, orderId).Format(format).Execute()

Готовые ярлыки‑наклейки на все коробки в одном заказе



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	format := openapiclient.PageFormatType("A9_HORIZONTALLY") // PageFormatType | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A7. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GenerateOrderLabels(context.Background(), campaignId, orderId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateOrderLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateOrderLabels`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateOrderLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOrderLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | [**PageFormatType**](PageFormatType.md) | Настройка размещения ярлыков на странице. Если параметра нет, возвращается PDF с ярлыками формата A7. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

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
	resp, r, err := apiClient.DbsAPI.GeneratePricesReport(context.Background()).GeneratePricesReportRequest(generatePricesReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GeneratePricesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeneratePricesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GeneratePricesReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateSalesGeographyReport(context.Background()).GenerateSalesGeographyRequest(generateSalesGeographyRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateSalesGeographyReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSalesGeographyReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateSalesGeographyReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateShelfsStatisticsReport(context.Background()).GenerateShelfsStatisticsRequest(generateShelfsStatisticsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateShelfsStatisticsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateShelfsStatisticsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateShelfsStatisticsReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateShowsBoostReport(context.Background()).GenerateShowsBoostRequest(generateShowsBoostRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateShowsBoostReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateShowsBoostReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateShowsBoostReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateShowsSalesReport(context.Background()).GenerateShowsSalesReportRequest(generateShowsSalesReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateShowsSalesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateShowsSalesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateShowsSalesReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateStocksOnWarehousesReport(context.Background()).GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateStocksOnWarehousesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateStocksOnWarehousesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateStocksOnWarehousesReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateUnitedMarketplaceServicesReport(context.Background()).GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest).Format(format).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateUnitedMarketplaceServicesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedMarketplaceServicesReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateUnitedMarketplaceServicesReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateUnitedNettingReport(context.Background()).GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest).Format(format).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateUnitedNettingReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedNettingReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateUnitedNettingReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateUnitedOrdersReport(context.Background()).GenerateUnitedOrdersRequest(generateUnitedOrdersRequest).Format(format).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateUnitedOrdersReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedOrdersReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateUnitedOrdersReport`: %v\n", resp)
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
	resp, r, err := apiClient.DbsAPI.GenerateUnitedReturnsReport(context.Background()).GenerateUnitedReturnsRequest(generateUnitedReturnsRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GenerateUnitedReturnsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateUnitedReturnsReport`: GenerateReportResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GenerateUnitedReturnsReport`: %v\n", resp)
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


## GetAuthTokenInfo

> GetTokenInfoResponse GetAuthTokenInfo(ctx).Execute()

Получение информации об авторизационном токене



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetAuthTokenInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetAuthTokenInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthTokenInfo`: GetTokenInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetAuthTokenInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthTokenInfoRequest struct via the builder pattern


### Return type

[**GetTokenInfoResponse**](GetTokenInfoResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBidsInfoForBusiness

> GetBidsInfoResponse GetBidsInfoForBusiness(ctx, businessId).PageToken(pageToken).Limit(limit).GetBidsInfoRequest(getBidsInfoRequest).Execute()

Информация об установленных ставках



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)
	getBidsInfoRequest := *openapiclient.NewGetBidsInfoRequest() // GetBidsInfoRequest | description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetBidsInfoForBusiness(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetBidsInfoRequest(getBidsInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetBidsInfoForBusiness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBidsInfoForBusiness`: GetBidsInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetBidsInfoForBusiness`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidsInfoForBusinessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
 **getBidsInfoRequest** | [**GetBidsInfoRequest**](GetBidsInfoRequest.md) | description | 

### Return type

[**GetBidsInfoResponse**](GetBidsInfoResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBidsRecommendations

> GetBidsRecommendationsResponse GetBidsRecommendations(ctx, businessId).GetBidsRecommendationsRequest(getBidsRecommendationsRequest).Execute()

Рекомендованные ставки для заданных товаров



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getBidsRecommendationsRequest := *openapiclient.NewGetBidsRecommendationsRequest([]string{"Skus_example"}) // GetBidsRecommendationsRequest | description.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetBidsRecommendations(context.Background(), businessId).GetBidsRecommendationsRequest(getBidsRecommendationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetBidsRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBidsRecommendations`: GetBidsRecommendationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetBidsRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidsRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getBidsRecommendationsRequest** | [**GetBidsRecommendationsRequest**](GetBidsRecommendationsRequest.md) | description. | 

### Return type

[**GetBidsRecommendationsResponse**](GetBidsRecommendationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessQuarantineOffers

> GetQuarantineOffersResponse GetBusinessQuarantineOffers(ctx, businessId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()

Список товаров, находящихся в карантине по цене в кабинете



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getQuarantineOffersRequest := *openapiclient.NewGetQuarantineOffersRequest() // GetQuarantineOffersRequest | 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetBusinessQuarantineOffers(context.Background(), businessId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetBusinessQuarantineOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessQuarantineOffers`: GetQuarantineOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetBusinessQuarantineOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessQuarantineOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getQuarantineOffersRequest** | [**GetQuarantineOffersRequest**](GetQuarantineOffersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetQuarantineOffersResponse**](GetQuarantineOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessSettings

> GetBusinessSettingsResponse GetBusinessSettings(ctx, businessId).Execute()

Настройки кабинета



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetBusinessSettings(context.Background(), businessId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetBusinessSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessSettings`: GetBusinessSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetBusinessSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBusinessSettingsResponse**](GetBusinessSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> GetCampaignResponse GetCampaign(ctx, campaignId).Execute()

Информация о магазине



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: GetCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignResponse**](GetCampaignResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignOffers

> GetCampaignOffersResponse GetCampaignOffers(ctx, campaignId).GetCampaignOffersRequest(getCampaignOffersRequest).PageToken(pageToken).Limit(limit).Execute()

Информация о товарах, которые размещены в заданном магазине



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
	getCampaignOffersRequest := *openapiclient.NewGetCampaignOffersRequest() // GetCampaignOffersRequest | 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCampaignOffers(context.Background(), campaignId).GetCampaignOffersRequest(getCampaignOffersRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCampaignOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignOffers`: GetCampaignOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCampaignOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getCampaignOffersRequest** | [**GetCampaignOffersRequest**](GetCampaignOffersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetCampaignOffersResponse**](GetCampaignOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignQuarantineOffers

> GetQuarantineOffersResponse GetCampaignQuarantineOffers(ctx, campaignId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()

Список товаров, находящихся в карантине по цене в магазине



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
	getQuarantineOffersRequest := *openapiclient.NewGetQuarantineOffersRequest() // GetQuarantineOffersRequest | 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCampaignQuarantineOffers(context.Background(), campaignId).GetQuarantineOffersRequest(getQuarantineOffersRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCampaignQuarantineOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignQuarantineOffers`: GetQuarantineOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCampaignQuarantineOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignQuarantineOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getQuarantineOffersRequest** | [**GetQuarantineOffersRequest**](GetQuarantineOffersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetQuarantineOffersResponse**](GetQuarantineOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignRegion

> GetCampaignRegionResponse GetCampaignRegion(ctx, campaignId).Execute()

Регион магазина



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCampaignRegion(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCampaignRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignRegion`: GetCampaignRegionResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCampaignRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignRegionResponse**](GetCampaignRegionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignSettings

> GetCampaignSettingsResponse GetCampaignSettings(ctx, campaignId).Execute()

Настройки магазина



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCampaignSettings(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCampaignSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignSettings`: GetCampaignSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCampaignSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignSettingsResponse**](GetCampaignSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> GetCampaignsResponse GetCampaigns(ctx).Page(page).PageSize(pageSize).Execute()

Список магазинов пользователя



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
	page := int32(56) // int32 | {% note warning \"Если в методе есть `page_token`\" %}  Используйте его вместо параметра `page`.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром `page_size`.  `page_number` игнорируется, если задан `page_token` или `limit`.  (optional) (default to 1)
	pageSize := int32(56) // int32 | Размер страницы.  Используется вместе с параметром `page_number`.  `page_size` игнорируется, если задан `page_token` или `limit`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCampaigns(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaigns`: GetCampaignsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | {% note warning \&quot;Если в методе есть &#x60;page_token&#x60;\&quot; %}  Используйте его вместо параметра &#x60;page&#x60;.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром &#x60;page_size&#x60;.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | [default to 1]
 **pageSize** | **int32** | Размер страницы.  Используется вместе с параметром &#x60;page_number&#x60;.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | 

### Return type

[**GetCampaignsResponse**](GetCampaignsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoriesMaxSaleQuantum

> GetCategoriesMaxSaleQuantumResponse GetCategoriesMaxSaleQuantum(ctx).GetCategoriesMaxSaleQuantumRequest(getCategoriesMaxSaleQuantumRequest).Execute()

Лимит на установку кванта продажи и минимального количества товаров в заказе



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
	getCategoriesMaxSaleQuantumRequest := *openapiclient.NewGetCategoriesMaxSaleQuantumRequest([]int64{int64(123)}) // GetCategoriesMaxSaleQuantumRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCategoriesMaxSaleQuantum(context.Background()).GetCategoriesMaxSaleQuantumRequest(getCategoriesMaxSaleQuantumRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCategoriesMaxSaleQuantum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoriesMaxSaleQuantum`: GetCategoriesMaxSaleQuantumResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCategoriesMaxSaleQuantum`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesMaxSaleQuantumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCategoriesMaxSaleQuantumRequest** | [**GetCategoriesMaxSaleQuantumRequest**](GetCategoriesMaxSaleQuantumRequest.md) |  | 

### Return type

[**GetCategoriesMaxSaleQuantumResponse**](GetCategoriesMaxSaleQuantumResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoriesTree

> GetCategoriesResponse GetCategoriesTree(ctx).GetCategoriesRequest(getCategoriesRequest).Execute()

Дерево категорий



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
	getCategoriesRequest := *openapiclient.NewGetCategoriesRequest() // GetCategoriesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCategoriesTree(context.Background()).GetCategoriesRequest(getCategoriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCategoriesTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoriesTree`: GetCategoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCategoriesTree`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCategoriesRequest** | [**GetCategoriesRequest**](GetCategoriesRequest.md) |  | 

### Return type

[**GetCategoriesResponse**](GetCategoriesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryContentParameters

> GetCategoryContentParametersResponse GetCategoryContentParameters(ctx, categoryId).Execute()

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
	categoryId := int64(789) // int64 | Идентификатор категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится интересующий вас товар, воспользуйтесь запросом [POST categories/tree](../../reference/categories/getCategoriesTree.md). 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetCategoryContentParameters(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetCategoryContentParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryContentParameters`: GetCategoryContentParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetCategoryContentParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int64** | Идентификатор категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится интересующий вас товар, воспользуйтесь запросом [POST categories/tree](../../reference/categories/getCategoriesTree.md).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryContentParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetChat

> GetChatResponse GetChat(ctx, businessId).ChatId(chatId).Execute()

Получение чата по идентификатору



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	chatId := int64(789) // int64 | Идентификатор чата.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetChat(context.Background(), businessId).ChatId(chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChat`: GetChatResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatId** | **int64** | Идентификатор чата. | 

### Return type

[**GetChatResponse**](GetChatResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatHistory

> GetChatHistoryResponse GetChatHistory(ctx, businessId).ChatId(chatId).GetChatHistoryRequest(getChatHistoryRequest).PageToken(pageToken).Limit(limit).Execute()

Получение истории сообщений в чате



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	chatId := int64(789) // int64 | Идентификатор чата.
	getChatHistoryRequest := *openapiclient.NewGetChatHistoryRequest() // GetChatHistoryRequest | description
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetChatHistory(context.Background(), businessId).ChatId(chatId).GetChatHistoryRequest(getChatHistoryRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetChatHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatHistory`: GetChatHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetChatHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatId** | **int64** | Идентификатор чата. | 
 **getChatHistoryRequest** | [**GetChatHistoryRequest**](GetChatHistoryRequest.md) | description | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetChatHistoryResponse**](GetChatHistoryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatMessage

> GetChatMessageResponse GetChatMessage(ctx, businessId).ChatId(chatId).MessageId(messageId).Execute()

Получение сообщения в чате



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	chatId := int64(789) // int64 | Идентификатор чата.
	messageId := int64(789) // int64 | Идентификатор сообщения.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetChatMessage(context.Background(), businessId).ChatId(chatId).MessageId(messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetChatMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatMessage`: GetChatMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetChatMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatId** | **int64** | Идентификатор чата. | 
 **messageId** | **int64** | Идентификатор сообщения. | 

### Return type

[**GetChatMessageResponse**](GetChatMessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChats

> GetChatsResponse GetChats(ctx, businessId).GetChatsRequest(getChatsRequest).PageToken(pageToken).Limit(limit).Execute()

Получение доступных чатов



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getChatsRequest := *openapiclient.NewGetChatsRequest() // GetChatsRequest | description
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetChats(context.Background(), businessId).GetChatsRequest(getChatsRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChats`: GetChatsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetChats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getChatsRequest** | [**GetChatsRequest**](GetChatsRequest.md) | description | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetChatsResponse**](GetChatsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeliveryServices

> GetDeliveryServicesResponse GetDeliveryServices(ctx).Execute()

Справочник служб доставки



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetDeliveryServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetDeliveryServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeliveryServices`: GetDeliveryServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetDeliveryServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryServicesRequest struct via the builder pattern


### Return type

[**GetDeliveryServicesResponse**](GetDeliveryServicesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGoodsFeedbackComments

> GetGoodsFeedbackCommentsResponse GetGoodsFeedbackComments(ctx, businessId).GetGoodsFeedbackCommentsRequest(getGoodsFeedbackCommentsRequest).PageToken(pageToken).Limit(limit).Execute()

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getGoodsFeedbackCommentsRequest := *openapiclient.NewGetGoodsFeedbackCommentsRequest() // GetGoodsFeedbackCommentsRequest | 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetGoodsFeedbackComments(context.Background(), businessId).GetGoodsFeedbackCommentsRequest(getGoodsFeedbackCommentsRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetGoodsFeedbackComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsFeedbackComments`: GetGoodsFeedbackCommentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetGoodsFeedbackComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsFeedbackCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getGoodsFeedbackCommentsRequest** | [**GetGoodsFeedbackCommentsRequest**](GetGoodsFeedbackCommentsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)
	getGoodsFeedbackRequest := *openapiclient.NewGetGoodsFeedbackRequest() // GetGoodsFeedbackRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetGoodsFeedbacks(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetGoodsFeedbackRequest(getGoodsFeedbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetGoodsFeedbacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsFeedbacks`: GetGoodsFeedbackResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetGoodsFeedbacks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsFeedbacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
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


## GetGoodsStats

> GetGoodsStatsResponse GetGoodsStats(ctx, campaignId).GetGoodsStatsRequest(getGoodsStatsRequest).Execute()

Отчет по товарам



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
	getGoodsStatsRequest := *openapiclient.NewGetGoodsStatsRequest([]string{"ShopSkus_example"}) // GetGoodsStatsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetGoodsStats(context.Background(), campaignId).GetGoodsStatsRequest(getGoodsStatsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetGoodsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGoodsStats`: GetGoodsStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetGoodsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGoodsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getGoodsStatsRequest** | [**GetGoodsStatsRequest**](GetGoodsStatsRequest.md) |  | 

### Return type

[**GetGoodsStatsResponse**](GetGoodsStatsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHiddenOffers

> GetHiddenOffersResponse GetHiddenOffers(ctx, campaignId).OfferId(offerId).PageToken(pageToken).Limit(limit).Execute()

Информация о скрытых вами товарах



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
	offerId := []string{"Inner_example"} // []string | Идентификатор скрытого предложения.  (optional)
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetHiddenOffers(context.Background(), campaignId).OfferId(offerId).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetHiddenOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHiddenOffers`: GetHiddenOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetHiddenOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHiddenOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offerId** | **[]string** | Идентификатор скрытого предложения.  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetHiddenOffersResponse**](GetHiddenOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModel

> GetModelsResponse GetModel(ctx, modelId).RegionId(regionId).Currency(currency).Execute()

Информация об одной модели



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
	modelId := int64(789) // int64 | Идентификатор модели товара.
	regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md). 
	currency := openapiclient.CurrencyType("RUR") // CurrencyType | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * `BYN` — белорусский рубль.  * `KZT` — казахстанский тенге.  * `RUR` — российский рубль.  * `UAH` — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetModel(context.Background(), modelId).RegionId(regionId).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModel`: GetModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **int64** | Идентификатор модели товара. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md).  | 
 **currency** | [**CurrencyType**](CurrencyType.md) | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * &#x60;BYN&#x60; — белорусский рубль.  * &#x60;KZT&#x60; — казахстанский тенге.  * &#x60;RUR&#x60; — российский рубль.  * &#x60;UAH&#x60; — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  | 

### Return type

[**GetModelsResponse**](GetModelsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelOffers

> GetModelsOffersResponse GetModelOffers(ctx, modelId).RegionId(regionId).Currency(currency).OrderByPrice(orderByPrice).Count(count).Page(page).Execute()

Список предложений для одной модели



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
	modelId := int64(789) // int64 | Идентификатор модели товара.
	regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md). 
	currency := openapiclient.CurrencyType("RUR") // CurrencyType | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * `BYN` — белорусский рубль.  * `KZT` — казахстанский тенге.  * `RUR` — российский рубль.  * `UAH` — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  (optional)
	orderByPrice := openapiclient.SortOrderType("ASC") // SortOrderType | Направление сортировки по цене.  Возможные значения: * `ASC` — сортировка по возрастанию. * `DESC` — сортировка по убыванию.  Значение по умолчанию: предложения выводятся в произвольном порядке.  (optional)
	count := int32(56) // int32 | Количество предложений на странице ответа. (optional) (default to 10)
	page := int32(56) // int32 | {% note warning \"Если в методе есть `page_token`\" %}  Используйте его вместо параметра `page`.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром `page_size`.  `page_number` игнорируется, если задан `page_token` или `limit`.  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetModelOffers(context.Background(), modelId).RegionId(regionId).Currency(currency).OrderByPrice(orderByPrice).Count(count).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetModelOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelOffers`: GetModelsOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetModelOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **int64** | Идентификатор модели товара. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md).  | 
 **currency** | [**CurrencyType**](CurrencyType.md) | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * &#x60;BYN&#x60; — белорусский рубль.  * &#x60;KZT&#x60; — казахстанский тенге.  * &#x60;RUR&#x60; — российский рубль.  * &#x60;UAH&#x60; — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  | 
 **orderByPrice** | [**SortOrderType**](SortOrderType.md) | Направление сортировки по цене.  Возможные значения: * &#x60;ASC&#x60; — сортировка по возрастанию. * &#x60;DESC&#x60; — сортировка по убыванию.  Значение по умолчанию: предложения выводятся в произвольном порядке.  | 
 **count** | **int32** | Количество предложений на странице ответа. | [default to 10]
 **page** | **int32** | {% note warning \&quot;Если в методе есть &#x60;page_token&#x60;\&quot; %}  Используйте его вместо параметра &#x60;page&#x60;.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром &#x60;page_size&#x60;.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | [default to 1]

### Return type

[**GetModelsOffersResponse**](GetModelsOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModels

> GetModelsResponse GetModels(ctx).RegionId(regionId).GetModelsRequest(getModelsRequest).Currency(currency).Execute()

Информация о нескольких моделях



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
	regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md). 
	getModelsRequest := *openapiclient.NewGetModelsRequest([]int64{int64(123)}) // GetModelsRequest | 
	currency := openapiclient.CurrencyType("RUR") // CurrencyType | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * `BYN` — белорусский рубль.  * `KZT` — казахстанский тенге.  * `RUR` — российский рубль.  * `UAH` — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetModels(context.Background()).RegionId(regionId).GetModelsRequest(getModelsRequest).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModels`: GetModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md).  | 
 **getModelsRequest** | [**GetModelsRequest**](GetModelsRequest.md) |  | 
 **currency** | [**CurrencyType**](CurrencyType.md) | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * &#x60;BYN&#x60; — белорусский рубль.  * &#x60;KZT&#x60; — казахстанский тенге.  * &#x60;RUR&#x60; — российский рубль.  * &#x60;UAH&#x60; — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  | 

### Return type

[**GetModelsResponse**](GetModelsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsOffers

> GetModelsOffersResponse GetModelsOffers(ctx).RegionId(regionId).GetModelsRequest(getModelsRequest).Currency(currency).OrderByPrice(orderByPrice).Execute()

Список предложений для нескольких моделей



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
	regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md). 
	getModelsRequest := *openapiclient.NewGetModelsRequest([]int64{int64(123)}) // GetModelsRequest | 
	currency := openapiclient.CurrencyType("RUR") // CurrencyType | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * `BYN` — белорусский рубль.  * `KZT` — казахстанский тенге.  * `RUR` — российский рубль.  * `UAH` — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  (optional)
	orderByPrice := openapiclient.SortOrderType("ASC") // SortOrderType | Направление сортировки по цене.  Возможные значения: * `ASC` — сортировка по возрастанию. * `DESC` — сортировка по убыванию.  Значение по умолчанию: предложения выводятся в произвольном порядке.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetModelsOffers(context.Background()).RegionId(regionId).GetModelsRequest(getModelsRequest).Currency(currency).OrderByPrice(orderByPrice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetModelsOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelsOffers`: GetModelsOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetModelsOffers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md).  | 
 **getModelsRequest** | [**GetModelsRequest**](GetModelsRequest.md) |  | 
 **currency** | [**CurrencyType**](CurrencyType.md) | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * &#x60;BYN&#x60; — белорусский рубль.  * &#x60;KZT&#x60; — казахстанский тенге.  * &#x60;RUR&#x60; — российский рубль.  * &#x60;UAH&#x60; — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  | 
 **orderByPrice** | [**SortOrderType**](SortOrderType.md) | Направление сортировки по цене.  Возможные значения: * &#x60;ASC&#x60; — сортировка по возрастанию. * &#x60;DESC&#x60; — сортировка по убыванию.  Значение по умолчанию: предложения выводятся в произвольном порядке.  | 

### Return type

[**GetModelsOffersResponse**](GetModelsOffersResponse.md)

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)
	getOfferCardsContentStatusRequest := *openapiclient.NewGetOfferCardsContentStatusRequest() // GetOfferCardsContentStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOfferCardsContentStatus(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetOfferCardsContentStatusRequest(getOfferCardsContentStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOfferCardsContentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferCardsContentStatus`: GetOfferCardsContentStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOfferCardsContentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferCardsContentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
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


## GetOfferMappingEntries

> GetOfferMappingEntriesResponse GetOfferMappingEntries(ctx, campaignId).OfferId(offerId).ShopSku(shopSku).MappingKind(mappingKind).Status(status).Availability(availability).CategoryId(categoryId).Vendor(vendor).PageToken(pageToken).Limit(limit).Execute()

Список товаров в каталоге



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
	offerId := []string{"Inner_example"} // []string | Идентификатор товара в каталоге. (optional)
	shopSku := []string{"Inner_example"} // []string | Ваш SKU товара.  Параметр может быть указан несколько раз, например:  ```text translate=no ...shop_sku=123&shop_sku=129&shop_sku=141... ```  В запросе можно указать либо параметр `shopSku`, либо любые параметры для фильтрации товаров. Совместное использование параметра `shopSku` и параметров для фильтрации приведет к ошибке.  (optional)
	mappingKind := openapiclient.OfferMappingKindType("ACTIVE") // OfferMappingKindType | Тип маппинга. (optional)
	status := []openapiclient.OfferProcessingStatusType{openapiclient.OfferProcessingStatusType("UNKNOWN")} // []OfferProcessingStatusType | Фильтрация по статусу публикации товара:  * `READY` — товар прошел модерацию. * `IN_WORK` — товар проходит модерацию. * `NEED_CONTENT` — для товара без SKU на Маркете marketSku нужно найти карточку самостоятельно или создать ее. * `NEED_INFO` — товар не прошел модерацию из-за ошибок или недостающих сведений в описании товара. * `REJECTED` — товар не прошел модерацию, так как Маркет не планирует размещать подобные товары. * `SUSPENDED` — товар не прошел модерацию, так как Маркет пока не размещает подобные товары. * `OTHER` — товар не прошел модерацию по другой причине.  Можно указать несколько статусов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ```text translate=no ...status=READY,IN_WORK... ...status=READY&status=IN_WORK... ```  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  (optional)
	availability := []openapiclient.OfferAvailabilityStatusType{openapiclient.OfferAvailabilityStatusType("ACTIVE")} // []OfferAvailabilityStatusType | Фильтрация по планам поставок товара:  * `ACTIVE` — поставки будут. * `INACTIVE` — поставок не будет: товар есть на складе, но вы больше не планируете его поставлять. * `DELISTED` — архив: товар закончился на складе, и его поставок больше не будет.  Можно указать несколько значений в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ```text translate=no ...availability=INACTIVE,DELISTED... ...availability=INACTIVE&availability=DELISTED... ```  В запросе можно указать либо параметр `shopSku`, либо любые параметры для фильтрации товаров. Совместное использование параметра `shopSku` и параметров для фильтрации приведет к ошибке.  (optional)
	categoryId := []int32{int32(123)} // []int32 | Фильтрация по идентификатору категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится товар, воспользуйтесь запросом [POST categories/tree](../../reference/categories/getCategoriesTree.md).  Можно указать несколько идентификаторов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ```text translate=no ...category_id=14727164,14382343... ...category_id=14727164&category_id=14382343... ```  В запросе можно указать либо параметр `shopSku`, либо любые параметры для фильтрации товаров. Совместное использование параметра `shopSku` и параметров для фильтрации приведет к ошибке.  (optional)
	vendor := []string{"Inner_example"} // []string | Фильтрация по бренду товара.  Можно указать несколько брендов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  ```text translate=no ...vendor=Aqua%20Minerale,Borjomi... ...vendor=Aqua%20Minerale&vendor=Borjomi... ```  Чтобы товар попал в результаты фильтрации, его бренд должен точно совпадать с одним из указанных в запросе. Например, если указан бренд Schwarzkopf, то в результатах не будет товаров Schwarzkopf Professional.  Если в названии бренда есть символы, которые не входят в таблицу ASCII (в том числе кириллические символы), используйте для них URL-кодирование. Например, пробел — %20, апостроф «'» — %27 и т. д. Подробнее см. в разделе [Кодирование URL русскоязычной Википедии](https://ru.wikipedia.org/wiki/URL#Кодирование_URL).  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  (optional)
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOfferMappingEntries(context.Background(), campaignId).OfferId(offerId).ShopSku(shopSku).MappingKind(mappingKind).Status(status).Availability(availability).CategoryId(categoryId).Vendor(vendor).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOfferMappingEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferMappingEntries`: GetOfferMappingEntriesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOfferMappingEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferMappingEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offerId** | **[]string** | Идентификатор товара в каталоге. | 
 **shopSku** | **[]string** | Ваш SKU товара.  Параметр может быть указан несколько раз, например:  &#x60;&#x60;&#x60;text translate&#x3D;no ...shop_sku&#x3D;123&amp;shop_sku&#x3D;129&amp;shop_sku&#x3D;141... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр &#x60;shopSku&#x60;, либо любые параметры для фильтрации товаров. Совместное использование параметра &#x60;shopSku&#x60; и параметров для фильтрации приведет к ошибке.  | 
 **mappingKind** | [**OfferMappingKindType**](OfferMappingKindType.md) | Тип маппинга. | 
 **status** | [**[]OfferProcessingStatusType**](OfferProcessingStatusType.md) | Фильтрация по статусу публикации товара:  * &#x60;READY&#x60; — товар прошел модерацию. * &#x60;IN_WORK&#x60; — товар проходит модерацию. * &#x60;NEED_CONTENT&#x60; — для товара без SKU на Маркете marketSku нужно найти карточку самостоятельно или создать ее. * &#x60;NEED_INFO&#x60; — товар не прошел модерацию из-за ошибок или недостающих сведений в описании товара. * &#x60;REJECTED&#x60; — товар не прошел модерацию, так как Маркет не планирует размещать подобные товары. * &#x60;SUSPENDED&#x60; — товар не прошел модерацию, так как Маркет пока не размещает подобные товары. * &#x60;OTHER&#x60; — товар не прошел модерацию по другой причине.  Можно указать несколько статусов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60;text translate&#x3D;no ...status&#x3D;READY,IN_WORK... ...status&#x3D;READY&amp;status&#x3D;IN_WORK... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  | 
 **availability** | [**[]OfferAvailabilityStatusType**](OfferAvailabilityStatusType.md) | Фильтрация по планам поставок товара:  * &#x60;ACTIVE&#x60; — поставки будут. * &#x60;INACTIVE&#x60; — поставок не будет: товар есть на складе, но вы больше не планируете его поставлять. * &#x60;DELISTED&#x60; — архив: товар закончился на складе, и его поставок больше не будет.  Можно указать несколько значений в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60;text translate&#x3D;no ...availability&#x3D;INACTIVE,DELISTED... ...availability&#x3D;INACTIVE&amp;availability&#x3D;DELISTED... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр &#x60;shopSku&#x60;, либо любые параметры для фильтрации товаров. Совместное использование параметра &#x60;shopSku&#x60; и параметров для фильтрации приведет к ошибке.  | 
 **categoryId** | **[]int32** | Фильтрация по идентификатору категории на Маркете.  Чтобы узнать идентификатор категории, к которой относится товар, воспользуйтесь запросом [POST categories/tree](../../reference/categories/getCategoriesTree.md).  Можно указать несколько идентификаторов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60;text translate&#x3D;no ...category_id&#x3D;14727164,14382343... ...category_id&#x3D;14727164&amp;category_id&#x3D;14382343... &#x60;&#x60;&#x60;  В запросе можно указать либо параметр &#x60;shopSku&#x60;, либо любые параметры для фильтрации товаров. Совместное использование параметра &#x60;shopSku&#x60; и параметров для фильтрации приведет к ошибке.  | 
 **vendor** | **[]string** | Фильтрация по бренду товара.  Можно указать несколько брендов в одном параметре, через запятую, или в нескольких одинаковых параметрах. Например:  &#x60;&#x60;&#x60;text translate&#x3D;no ...vendor&#x3D;Aqua%20Minerale,Borjomi... ...vendor&#x3D;Aqua%20Minerale&amp;vendor&#x3D;Borjomi... &#x60;&#x60;&#x60;  Чтобы товар попал в результаты фильтрации, его бренд должен точно совпадать с одним из указанных в запросе. Например, если указан бренд Schwarzkopf, то в результатах не будет товаров Schwarzkopf Professional.  Если в названии бренда есть символы, которые не входят в таблицу ASCII (в том числе кириллические символы), используйте для них URL-кодирование. Например, пробел — %20, апостроф «&#39;» — %27 и т. д. Подробнее см. в разделе [Кодирование URL русскоязычной Википедии](https://ru.wikipedia.org/wiki/URL#Кодирование_URL).  В запросе можно указать либо параметр shopSku, либо любые параметры для фильтрации товаров. Совместное использование параметра shopSku и параметров для фильтрации приведет к ошибке.  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetOfferMappingEntriesResponse**](GetOfferMappingEntriesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)
	language := openapiclient.CatalogLanguageType("RU") // CatalogLanguageType | Язык, на котором принимаются и возвращаются значения в параметрах `name` и `description`.  Значение по умолчанию: `RU`.  (optional)
	getOfferMappingsRequest := *openapiclient.NewGetOfferMappingsRequest() // GetOfferMappingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOfferMappings(context.Background(), businessId).PageToken(pageToken).Limit(limit).Language(language).GetOfferMappingsRequest(getOfferMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOfferMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferMappings`: GetOfferMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
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


## GetOfferRecommendations

> GetOfferRecommendationsResponse GetOfferRecommendations(ctx, businessId).GetOfferRecommendationsRequest(getOfferRecommendationsRequest).PageToken(pageToken).Limit(limit).Execute()

Рекомендации Маркета, касающиеся цен



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getOfferRecommendationsRequest := *openapiclient.NewGetOfferRecommendationsRequest() // GetOfferRecommendationsRequest | 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOfferRecommendations(context.Background(), businessId).GetOfferRecommendationsRequest(getOfferRecommendationsRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOfferRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfferRecommendations`: GetOfferRecommendationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOfferRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getOfferRecommendationsRequest** | [**GetOfferRecommendationsRequest**](GetOfferRecommendationsRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetOfferRecommendationsResponse**](GetOfferRecommendationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> GetOrderResponse GetOrder(ctx, campaignId, orderId).Execute()

Информация об одном заказе



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
	orderId := int64(789) // int64 | Идентификатор заказа.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOrder(context.Background(), campaignId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: GetOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderBusinessBuyerInfo

> GetBusinessBuyerInfoResponse GetOrderBusinessBuyerInfo(ctx, campaignId, orderId).Execute()

Информация о покупателе — юридическом лице



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
	orderId := int64(789) // int64 | Идентификатор заказа.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOrderBusinessBuyerInfo(context.Background(), campaignId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOrderBusinessBuyerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderBusinessBuyerInfo`: GetBusinessBuyerInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOrderBusinessBuyerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderBusinessBuyerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBusinessBuyerInfoResponse**](GetBusinessBuyerInfoResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderBusinessDocumentsInfo

> GetBusinessDocumentsInfoResponse GetOrderBusinessDocumentsInfo(ctx, campaignId, orderId).Execute()

Информация о документах



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
	orderId := int64(789) // int64 | Идентификатор заказа.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOrderBusinessDocumentsInfo(context.Background(), campaignId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOrderBusinessDocumentsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderBusinessDocumentsInfo`: GetBusinessDocumentsInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOrderBusinessDocumentsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderBusinessDocumentsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBusinessDocumentsInfoResponse**](GetBusinessDocumentsInfoResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderBuyerInfo

> GetOrderBuyerInfoResponse GetOrderBuyerInfo(ctx, campaignId, orderId).Execute()

Информация о покупателе — физическом лице



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
	orderId := int64(789) // int64 | Идентификатор заказа.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOrderBuyerInfo(context.Background(), campaignId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOrderBuyerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderBuyerInfo`: GetOrderBuyerInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOrderBuyerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderBuyerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrderBuyerInfoResponse**](GetOrderBuyerInfoResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderLabelsData

> GetOrderLabelsDataResponse GetOrderLabelsData(ctx, campaignId, orderId).Execute()

Данные для самостоятельного изготовления ярлыков



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
	orderId := int64(789) // int64 | Идентификатор заказа.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOrderLabelsData(context.Background(), campaignId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOrderLabelsData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderLabelsData`: GetOrderLabelsDataResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOrderLabelsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderLabelsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrderLabelsDataResponse**](GetOrderLabelsDataResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> GetOrdersResponse GetOrders(ctx, campaignId).OrderIds(orderIds).Status(status).Substatus(substatus).FromDate(fromDate).ToDate(toDate).SupplierShipmentDateFrom(supplierShipmentDateFrom).SupplierShipmentDateTo(supplierShipmentDateTo).UpdatedAtFrom(updatedAtFrom).UpdatedAtTo(updatedAtTo).DispatchType(dispatchType).Fake(fake).HasCis(hasCis).OnlyWaitingForCancellationApprove(onlyWaitingForCancellationApprove).OnlyEstimatedDelivery(onlyEstimatedDelivery).BuyerType(buyerType).Page(page).PageSize(pageSize).PageToken(pageToken).Limit(limit).Execute()

Информация о нескольких заказах



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderIds := []int64{int64(123)} // []int64 | Фильтрация заказов по идентификаторам. <br><br> ⚠️ Не используйте это поле одновременно с другими фильтрами. Если вы хотите воспользоваться ими, оставьте поле пустым.  (optional)
	status := []openapiclient.OrderStatusType{openapiclient.OrderStatusType("PLACING")} // []OrderStatusType | Статус заказа:  * `CANCELLED` — заказ отменен.  * `DELIVERED` — заказ получен покупателем.  * `DELIVERY` — заказ передан в службу доставки.  * `PICKUP` — заказ доставлен в пункт самовывоза.  * `PROCESSING` — заказ находится в обработке.  * `UNPAID` — заказ оформлен, но еще не оплачен (если выбрана оплата при оформлении).  Также могут возвращаться другие значения. Обрабатывать их не требуется.  (optional)
	substatus := []openapiclient.OrderSubstatusType{openapiclient.OrderSubstatusType("RESERVATION_EXPIRED")} // []OrderSubstatusType | Этап обработки заказа (если он имеет статус `PROCESSING`) или причина отмены заказа (если он имеет статус `CANCELLED`).  Возможные значения для заказа в статусе `PROCESSING`:  * `STARTED` — заказ подтвержден, его можно начать обрабатывать. * `READY_TO_SHIP` — заказ собран и готов к отправке. * `SHIPPED` — заказ передан службе доставки.  Возможные значения для заказа в статусе `CANCELLED`:  * `RESERVATION_EXPIRED` — покупатель не завершил оформление зарезервированного заказа в течение 10 минут.  * `USER_NOT_PAID` — покупатель не оплатил заказ (для типа оплаты `PREPAID`) в течение 30 минут.  * `USER_UNREACHABLE` — не удалось связаться с покупателем. Для отмены с этой причиной необходимо выполнить условия:    * не менее 3 звонков с 8 до 21 в часовом поясе покупателя;   * перерыв между первым и третьим звонком не менее 90 минут;   * соединение не короче 5 секунд.    Если хотя бы одно из этих условий не выполнено (кроме случая, когда номер недоступен), отменить заказ не получится. Вернется ответ с кодом ошибки 400  * `USER_CHANGED_MIND` — покупатель отменил заказ по личным причинам.  * `USER_REFUSED_DELIVERY` — покупателя не устроили условия доставки.  * `USER_REFUSED_PRODUCT` — покупателю не подошел товар.  * `SHOP_FAILED` — магазин не может выполнить заказ.  * `USER_REFUSED_QUALITY` — покупателя не устроило качество товара.  * `REPLACING_ORDER` — покупатель решил заменить товар другим по собственной инициативе.  * `PROCESSING_EXPIRED` — значение более не используется.  * `PICKUP_EXPIRED` — закончился срок хранения заказа в ПВЗ.  * `DELIVERY_SERVICE_UNDELIVERED` — служба доставки не смогла доставить заказ.  * `CANCELLED_COURIER_NOT_FOUND` — не удалось найти курьера.  * `USER_WANTS_TO_CHANGE_DELIVERY_DATE` — покупатель хочет получить заказ в другой день.  * `RESERVATION_FAILED` — Маркет не может продолжить дальнейшую обработку заказа.  Также могут возвращаться другие значения. Обрабатывать их не требуется.  (optional)
	fromDate := time.Now() // string | Начальная дата для фильтрации заказов по дате оформления.  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной и конечной датой (параметр `toDate`) должно быть не больше 30 дней.  Значение по умолчанию: 30 дней назад от текущей даты.  (optional)
	toDate := time.Now() // string | Конечная дата для фильтрации заказов по дате оформления.  Показываются заказы, созданные до 00:00 указанного дня.  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной (параметр `fromDate`) и конечной датой должно быть не больше 30 дней.  Значение по умолчанию: текущая дата.  Если промежуток времени между `toDate` и `fromDate` меньше суток, то `toDate` равен `fromDate` + сутки.  (optional)
	supplierShipmentDateFrom := time.Now() // string | Начальная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр `shipmentDate`).  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной и конечной датой (параметр `supplierShipmentDateTo`) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  (optional)
	supplierShipmentDateTo := time.Now() // string | Конечная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр `shipmentDate`).  Формат даты: `ДД-ММ-ГГГГ`.  Между начальной (параметр `supplierShipmentDateFrom`) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  Если промежуток времени между `supplierShipmentDateTo` и `supplierShipmentDateFrom` меньше суток, то `supplierShipmentDateTo` равен `supplierShipmentDateFrom` + сутки.  (optional)
	updatedAtFrom := time.Now() // time.Time | Начальная дата для фильтрации заказов по дате и времени обновления (параметр `updatedAt`).  Формат даты: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`.  Между начальной и конечной датой (параметр `updatedAtTo`) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  (optional)
	updatedAtTo := time.Now() // time.Time | Конечная дата для фильтрации заказов по дате и времени обновления (параметр `updatedAt`).  Формат даты: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`.  Между начальной (параметр `updatedAtFrom`) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  (optional)
	dispatchType := openapiclient.OrderDeliveryDispatchType("UNKNOWN") // OrderDeliveryDispatchType | Способ отгрузки (optional)
	fake := true // bool | Фильтрация заказов по типам:  * `false` — настоящий заказ покупателя.  * `true` — [тестовый](../../concepts/sandbox.md) заказ Маркета.  (optional) (default to false)
	hasCis := true // bool | Нужно ли вернуть только те заказы, в составе которых есть хотя бы один товар с кодом идентификации в системе [«Честный ЗНАК»](https://честныйзнак.рф/) или [«ASL BELGISI»](https://aslbelgisi.uz) (для продавцов Market Yandex Go):  * `true` — да.  * `false` — нет.  Такие коды присваиваются товарам, которые подлежат маркировке и относятся к определенным категориям.  (optional) (default to false)
	onlyWaitingForCancellationApprove := true // bool | **Только для модели DBS**  Фильтрация заказов по наличию запросов покупателей на отмену.  При значение `true` возвращаются только заказы, которые находятся в статусе `DELIVERY` или `PICKUP` и которые пользователи решили отменить.  Чтобы подтвердить или отклонить отмену, отправьте запрос [PUT campaigns/{campaignId}/orders/{orderId}/cancellation/accept](../../reference/orders/acceptOrderCancellation).  (optional) (default to false)
	onlyEstimatedDelivery := true // bool | Фильтрация заказов с долгой доставкой (31-60 дней) по подтвержденной дате доставки:  * `true` — возвращаются только заказы с неподтвержденной датой доставки. * `false` — фильтрация не применяется.  (optional) (default to false)
	buyerType := openapiclient.OrderBuyerType("PERSON") // OrderBuyerType | Фильтрация заказов по типу покупателя.  (optional)
	page := int32(56) // int32 | {% note warning \"Если в методе есть `page_token`\" %}  Используйте его вместо параметра `page`.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром `page_size`.  `page_number` игнорируется, если задан `page_token` или `limit`.  (optional) (default to 1)
	pageSize := int32(56) // int32 | Размер страницы.  Используется вместе с параметром `page_number`.  `page_size` игнорируется, если задан `page_token` или `limit`.  (optional)
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOrders(context.Background(), campaignId).OrderIds(orderIds).Status(status).Substatus(substatus).FromDate(fromDate).ToDate(toDate).SupplierShipmentDateFrom(supplierShipmentDateFrom).SupplierShipmentDateTo(supplierShipmentDateTo).UpdatedAtFrom(updatedAtFrom).UpdatedAtTo(updatedAtTo).DispatchType(dispatchType).Fake(fake).HasCis(hasCis).OnlyWaitingForCancellationApprove(onlyWaitingForCancellationApprove).OnlyEstimatedDelivery(onlyEstimatedDelivery).BuyerType(buyerType).Page(page).PageSize(pageSize).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrders`: GetOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderIds** | **[]int64** | Фильтрация заказов по идентификаторам. &lt;br&gt;&lt;br&gt; ⚠️ Не используйте это поле одновременно с другими фильтрами. Если вы хотите воспользоваться ими, оставьте поле пустым.  | 
 **status** | [**[]OrderStatusType**](OrderStatusType.md) | Статус заказа:  * &#x60;CANCELLED&#x60; — заказ отменен.  * &#x60;DELIVERED&#x60; — заказ получен покупателем.  * &#x60;DELIVERY&#x60; — заказ передан в службу доставки.  * &#x60;PICKUP&#x60; — заказ доставлен в пункт самовывоза.  * &#x60;PROCESSING&#x60; — заказ находится в обработке.  * &#x60;UNPAID&#x60; — заказ оформлен, но еще не оплачен (если выбрана оплата при оформлении).  Также могут возвращаться другие значения. Обрабатывать их не требуется.  | 
 **substatus** | [**[]OrderSubstatusType**](OrderSubstatusType.md) | Этап обработки заказа (если он имеет статус &#x60;PROCESSING&#x60;) или причина отмены заказа (если он имеет статус &#x60;CANCELLED&#x60;).  Возможные значения для заказа в статусе &#x60;PROCESSING&#x60;:  * &#x60;STARTED&#x60; — заказ подтвержден, его можно начать обрабатывать. * &#x60;READY_TO_SHIP&#x60; — заказ собран и готов к отправке. * &#x60;SHIPPED&#x60; — заказ передан службе доставки.  Возможные значения для заказа в статусе &#x60;CANCELLED&#x60;:  * &#x60;RESERVATION_EXPIRED&#x60; — покупатель не завершил оформление зарезервированного заказа в течение 10 минут.  * &#x60;USER_NOT_PAID&#x60; — покупатель не оплатил заказ (для типа оплаты &#x60;PREPAID&#x60;) в течение 30 минут.  * &#x60;USER_UNREACHABLE&#x60; — не удалось связаться с покупателем. Для отмены с этой причиной необходимо выполнить условия:    * не менее 3 звонков с 8 до 21 в часовом поясе покупателя;   * перерыв между первым и третьим звонком не менее 90 минут;   * соединение не короче 5 секунд.    Если хотя бы одно из этих условий не выполнено (кроме случая, когда номер недоступен), отменить заказ не получится. Вернется ответ с кодом ошибки 400  * &#x60;USER_CHANGED_MIND&#x60; — покупатель отменил заказ по личным причинам.  * &#x60;USER_REFUSED_DELIVERY&#x60; — покупателя не устроили условия доставки.  * &#x60;USER_REFUSED_PRODUCT&#x60; — покупателю не подошел товар.  * &#x60;SHOP_FAILED&#x60; — магазин не может выполнить заказ.  * &#x60;USER_REFUSED_QUALITY&#x60; — покупателя не устроило качество товара.  * &#x60;REPLACING_ORDER&#x60; — покупатель решил заменить товар другим по собственной инициативе.  * &#x60;PROCESSING_EXPIRED&#x60; — значение более не используется.  * &#x60;PICKUP_EXPIRED&#x60; — закончился срок хранения заказа в ПВЗ.  * &#x60;DELIVERY_SERVICE_UNDELIVERED&#x60; — служба доставки не смогла доставить заказ.  * &#x60;CANCELLED_COURIER_NOT_FOUND&#x60; — не удалось найти курьера.  * &#x60;USER_WANTS_TO_CHANGE_DELIVERY_DATE&#x60; — покупатель хочет получить заказ в другой день.  * &#x60;RESERVATION_FAILED&#x60; — Маркет не может продолжить дальнейшую обработку заказа.  Также могут возвращаться другие значения. Обрабатывать их не требуется.  | 
 **fromDate** | **string** | Начальная дата для фильтрации заказов по дате оформления.  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной и конечной датой (параметр &#x60;toDate&#x60;) должно быть не больше 30 дней.  Значение по умолчанию: 30 дней назад от текущей даты.  | 
 **toDate** | **string** | Конечная дата для фильтрации заказов по дате оформления.  Показываются заказы, созданные до 00:00 указанного дня.  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной (параметр &#x60;fromDate&#x60;) и конечной датой должно быть не больше 30 дней.  Значение по умолчанию: текущая дата.  Если промежуток времени между &#x60;toDate&#x60; и &#x60;fromDate&#x60; меньше суток, то &#x60;toDate&#x60; равен &#x60;fromDate&#x60; + сутки.  | 
 **supplierShipmentDateFrom** | **string** | Начальная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр &#x60;shipmentDate&#x60;).  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной и конечной датой (параметр &#x60;supplierShipmentDateTo&#x60;) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  | 
 **supplierShipmentDateTo** | **string** | Конечная дата для фильтрации заказов по дате отгрузки в службу доставки (параметр &#x60;shipmentDate&#x60;).  Формат даты: &#x60;ДД-ММ-ГГГГ&#x60;.  Между начальной (параметр &#x60;supplierShipmentDateFrom&#x60;) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  Если промежуток времени между &#x60;supplierShipmentDateTo&#x60; и &#x60;supplierShipmentDateFrom&#x60; меньше суток, то &#x60;supplierShipmentDateTo&#x60; равен &#x60;supplierShipmentDateFrom&#x60; + сутки.  | 
 **updatedAtFrom** | **time.Time** | Начальная дата для фильтрации заказов по дате и времени обновления (параметр &#x60;updatedAt&#x60;).  Формат даты: ISO 8601 со смещением относительно UTC. Например, &#x60;2017-11-21T00:42:42+03:00&#x60;.  Между начальной и конечной датой (параметр &#x60;updatedAtTo&#x60;) должно быть не больше 30 дней.  Начальная дата включается в интервал для фильтрации.  | 
 **updatedAtTo** | **time.Time** | Конечная дата для фильтрации заказов по дате и времени обновления (параметр &#x60;updatedAt&#x60;).  Формат даты: ISO 8601 со смещением относительно UTC. Например, &#x60;2017-11-21T00:42:42+03:00&#x60;.  Между начальной (параметр &#x60;updatedAtFrom&#x60;) и конечной датой должно быть не больше 30 дней.  Конечная дата не включается в интервал для фильтрации.  | 
 **dispatchType** | [**OrderDeliveryDispatchType**](OrderDeliveryDispatchType.md) | Способ отгрузки | 
 **fake** | **bool** | Фильтрация заказов по типам:  * &#x60;false&#x60; — настоящий заказ покупателя.  * &#x60;true&#x60; — [тестовый](../../concepts/sandbox.md) заказ Маркета.  | [default to false]
 **hasCis** | **bool** | Нужно ли вернуть только те заказы, в составе которых есть хотя бы один товар с кодом идентификации в системе [«Честный ЗНАК»](https://честныйзнак.рф/) или [«ASL BELGISI»](https://aslbelgisi.uz) (для продавцов Market Yandex Go):  * &#x60;true&#x60; — да.  * &#x60;false&#x60; — нет.  Такие коды присваиваются товарам, которые подлежат маркировке и относятся к определенным категориям.  | [default to false]
 **onlyWaitingForCancellationApprove** | **bool** | **Только для модели DBS**  Фильтрация заказов по наличию запросов покупателей на отмену.  При значение &#x60;true&#x60; возвращаются только заказы, которые находятся в статусе &#x60;DELIVERY&#x60; или &#x60;PICKUP&#x60; и которые пользователи решили отменить.  Чтобы подтвердить или отклонить отмену, отправьте запрос [PUT campaigns/{campaignId}/orders/{orderId}/cancellation/accept](../../reference/orders/acceptOrderCancellation).  | [default to false]
 **onlyEstimatedDelivery** | **bool** | Фильтрация заказов с долгой доставкой (31-60 дней) по подтвержденной дате доставки:  * &#x60;true&#x60; — возвращаются только заказы с неподтвержденной датой доставки. * &#x60;false&#x60; — фильтрация не применяется.  | [default to false]
 **buyerType** | [**OrderBuyerType**](OrderBuyerType.md) | Фильтрация заказов по типу покупателя.  | 
 **page** | **int32** | {% note warning \&quot;Если в методе есть &#x60;page_token&#x60;\&quot; %}  Используйте его вместо параметра &#x60;page&#x60;.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром &#x60;page_size&#x60;.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | [default to 1]
 **pageSize** | **int32** | Размер страницы.  Используется вместе с параметром &#x60;page_number&#x60;.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetOrdersResponse**](GetOrdersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersStats

> GetOrdersStatsResponse GetOrdersStats(ctx, campaignId).PageToken(pageToken).Limit(limit).GetOrdersStatsRequest(getOrdersStatsRequest).Execute()

Детальная информация по заказам



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
	getOrdersStatsRequest := *openapiclient.NewGetOrdersStatsRequest() // GetOrdersStatsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOrdersStats(context.Background(), campaignId).PageToken(pageToken).Limit(limit).GetOrdersStatsRequest(getOrdersStatsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOrdersStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersStats`: GetOrdersStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOrdersStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
 **getOrdersStatsRequest** | [**GetOrdersStatsRequest**](GetOrdersStatsRequest.md) |  | 

### Return type

[**GetOrdersStatsResponse**](GetOrdersStatsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutlet

> GetOutletResponse GetOutlet(ctx, campaignId, outletId).Execute()

Информация об одной точке продаж



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
	outletId := int64(789) // int64 | Идентификатор точки продаж.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOutlet(context.Background(), campaignId, outletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOutlet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutlet`: GetOutletResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOutlet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**outletId** | **int64** | Идентификатор точки продаж. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOutletResponse**](GetOutletResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutletLicenses

> GetOutletLicensesResponse GetOutletLicenses(ctx, campaignId).OutletIds(outletIds).Ids(ids).Execute()

Информация о лицензиях для точек продаж



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
	outletIds := []int64{int64(123)} // []int64 | Список идентификаторов точек продаж, для которых нужно получить информацию о лицензиях. Идентификаторы указываются через запятую.  В запросе должен быть либо параметр `outletIds`, либо параметр `ids`. Запрос с обоими параметрами или без них приведет к ошибке.  (optional)
	ids := []int64{int64(123)} // []int64 | Список идентификаторов лицензий. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOutletLicenses(context.Background(), campaignId).OutletIds(outletIds).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOutletLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutletLicenses`: GetOutletLicensesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOutletLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutletLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **outletIds** | **[]int64** | Список идентификаторов точек продаж, для которых нужно получить информацию о лицензиях. Идентификаторы указываются через запятую.  В запросе должен быть либо параметр &#x60;outletIds&#x60;, либо параметр &#x60;ids&#x60;. Запрос с обоими параметрами или без них приведет к ошибке.  | 
 **ids** | **[]int64** | Список идентификаторов лицензий. | 

### Return type

[**GetOutletLicensesResponse**](GetOutletLicensesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutlets

> GetOutletsResponse GetOutlets(ctx, campaignId).PageToken(pageToken).RegionId(regionId).ShopOutletCode(shopOutletCode).RegionId2(regionId2).Execute()

Информация о нескольких точках продаж



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
	regionId := int64(789) // int64 | Идентификатор региона. Если задать идентификатор родительского региона любого уровня, в выходных данных будут отображены точки продаж всех дочерних регионов. Идентификатор региона можно получить c помощью метода [GET regions](../../reference/regions/searchRegionsByName.md).  (optional)
	shopOutletCode := "shopOutletCode_example" // string | Идентификатор точки продаж, присвоенный магазином. (optional)
	regionId2 := int64(789) // int64 | {% note warning \"Вместо него используйте `region_id`.\" %}     {% endnote %}  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetOutlets(context.Background(), campaignId).PageToken(pageToken).RegionId(regionId).ShopOutletCode(shopOutletCode).RegionId2(regionId2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetOutlets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutlets`: GetOutletsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetOutlets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **regionId** | **int64** | Идентификатор региона. Если задать идентификатор родительского региона любого уровня, в выходных данных будут отображены точки продаж всех дочерних регионов. Идентификатор региона можно получить c помощью метода [GET regions](../../reference/regions/searchRegionsByName.md).  | 
 **shopOutletCode** | **string** | Идентификатор точки продаж, присвоенный магазином. | 
 **regionId2** | **int64** | {% note warning \&quot;Вместо него используйте &#x60;region_id&#x60;.\&quot; %}     {% endnote %}  | 

### Return type

[**GetOutletsResponse**](GetOutletsResponse.md)

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)
	getPagedWarehousesRequest := *openapiclient.NewGetPagedWarehousesRequest() // GetPagedWarehousesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetPagedWarehouses(context.Background(), businessId).PageToken(pageToken).Limit(limit).GetPagedWarehousesRequest(getPagedWarehousesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetPagedWarehouses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagedWarehouses`: GetPagedWarehousesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetPagedWarehouses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagedWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
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


## GetPrices

> GetPricesResponse GetPrices(ctx, campaignId).PageToken(pageToken).Limit(limit).Archived(archived).Execute()

Список цен



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
	archived := true // bool | Фильтр по нахождению в архиве. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetPrices(context.Background(), campaignId).PageToken(pageToken).Limit(limit).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrices`: GetPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
 **archived** | **bool** | Фильтр по нахождению в архиве. | [default to false]

### Return type

[**GetPricesResponse**](GetPricesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricesByOfferIds

> GetPricesByOfferIdsResponse GetPricesByOfferIds(ctx, campaignId).PageToken(pageToken).Limit(limit).GetPricesByOfferIdsRequest(getPricesByOfferIdsRequest).Execute()

Просмотр цен на указанные товары в магазине



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
	getPricesByOfferIdsRequest := *openapiclient.NewGetPricesByOfferIdsRequest() // GetPricesByOfferIdsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetPricesByOfferIds(context.Background(), campaignId).PageToken(pageToken).Limit(limit).GetPricesByOfferIdsRequest(getPricesByOfferIdsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetPricesByOfferIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricesByOfferIds`: GetPricesByOfferIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetPricesByOfferIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesByOfferIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
 **getPricesByOfferIdsRequest** | [**GetPricesByOfferIdsRequest**](GetPricesByOfferIdsRequest.md) |  | 

### Return type

[**GetPricesByOfferIdsResponse**](GetPricesByOfferIdsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromoOffers

> GetPromoOffersResponse GetPromoOffers(ctx, businessId).GetPromoOffersRequest(getPromoOffersRequest).PageToken(pageToken).Limit(limit).Execute()

Получение списка товаров, которые участвуют или могут участвовать в акции



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getPromoOffersRequest := *openapiclient.NewGetPromoOffersRequest("PromoId_example") // GetPromoOffersRequest | 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetPromoOffers(context.Background(), businessId).GetPromoOffersRequest(getPromoOffersRequest).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetPromoOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromoOffers`: GetPromoOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetPromoOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromoOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getPromoOffersRequest** | [**GetPromoOffersRequest**](GetPromoOffersRequest.md) |  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetPromoOffersResponse**](GetPromoOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPromos

> GetPromosResponse GetPromos(ctx, businessId).GetPromosRequest(getPromosRequest).Execute()

Получение списка акций



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getPromosRequest := *openapiclient.NewGetPromosRequest() // GetPromosRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetPromos(context.Background(), businessId).GetPromosRequest(getPromosRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetPromos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromos`: GetPromosResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetPromos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getPromosRequest** | [**GetPromosRequest**](GetPromosRequest.md) |  | 

### Return type

[**GetPromosResponse**](GetPromosResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQualityRatingDetails

> GetQualityRatingDetailsResponse GetQualityRatingDetails(ctx, campaignId).Execute()

Заказы, которые повлияли на индекс качества



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetQualityRatingDetails(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetQualityRatingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQualityRatingDetails`: GetQualityRatingDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetQualityRatingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityRatingDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetQualityRatingDetailsResponse**](GetQualityRatingDetailsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQualityRatings

> GetQualityRatingResponse GetQualityRatings(ctx, businessId).GetQualityRatingRequest(getQualityRatingRequest).Execute()

Индекс качества магазинов



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getQualityRatingRequest := *openapiclient.NewGetQualityRatingRequest([]int64{int64(123)}) // GetQualityRatingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetQualityRatings(context.Background(), businessId).GetQualityRatingRequest(getQualityRatingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetQualityRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQualityRatings`: GetQualityRatingResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetQualityRatings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQualityRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getQualityRatingRequest** | [**GetQualityRatingRequest**](GetQualityRatingRequest.md) |  | 

### Return type

[**GetQualityRatingResponse**](GetQualityRatingResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionsCodes

> GetRegionsCodesResponse GetRegionsCodes(ctx).Execute()

Список допустимых кодов стран



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetRegionsCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetRegionsCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegionsCodes`: GetRegionsCodesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetRegionsCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsCodesRequest struct via the builder pattern


### Return type

[**GetRegionsCodesResponse**](GetRegionsCodesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.DbsAPI.GetReportInfo(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetReportInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportInfo`: GetReportInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetReportInfo`: %v\n", resp)
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
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetReturn(context.Background(), campaignId, orderId, returnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturn`: GetReturnResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetReturn`: %v\n", resp)
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
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetReturnApplication(context.Background(), campaignId, orderId, returnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetReturnApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturnApplication`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetReturnApplication`: %v\n", resp)
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
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.
	itemId := int64(789) // int64 | Идентификатор товара в возврате.
	imageHash := "imageHash_example" // string | Хеш ссылки изображения для загрузки.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetReturnPhoto(context.Background(), campaignId, orderId, returnId, itemId, imageHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetReturnPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturnPhoto`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetReturnPhoto`: %v\n", resp)
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
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
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
	resp, r, err := apiClient.DbsAPI.GetReturns(context.Background(), campaignId).PageToken(pageToken).Limit(limit).OrderIds(orderIds).Statuses(statuses).Type_(type_).FromDate(fromDate).ToDate(toDate).FromDate2(fromDate2).ToDate2(toDate2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetReturns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReturns`: GetReturnsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetReturns`: %v\n", resp)
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


## GetStocks

> GetWarehouseStocksResponse GetStocks(ctx, campaignId).PageToken(pageToken).Limit(limit).GetWarehouseStocksRequest(getWarehouseStocksRequest).Execute()

Информация об остатках и оборачиваемости



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
	getWarehouseStocksRequest := *openapiclient.NewGetWarehouseStocksRequest() // GetWarehouseStocksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetStocks(context.Background(), campaignId).PageToken(pageToken).Limit(limit).GetWarehouseStocksRequest(getWarehouseStocksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetStocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStocks`: GetWarehouseStocksResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetStocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 
 **getWarehouseStocksRequest** | [**GetWarehouseStocksRequest**](GetWarehouseStocksRequest.md) |  | 

### Return type

[**GetWarehouseStocksResponse**](GetWarehouseStocksResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestedOfferMappingEntries

> GetSuggestedOfferMappingEntriesResponse GetSuggestedOfferMappingEntries(ctx, campaignId).GetSuggestedOfferMappingEntriesRequest(getSuggestedOfferMappingEntriesRequest).Execute()

Рекомендованные карточки для товаров



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
	getSuggestedOfferMappingEntriesRequest := *openapiclient.NewGetSuggestedOfferMappingEntriesRequest([]openapiclient.MappingsOfferDTO{*openapiclient.NewMappingsOfferDTO()}) // GetSuggestedOfferMappingEntriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetSuggestedOfferMappingEntries(context.Background(), campaignId).GetSuggestedOfferMappingEntriesRequest(getSuggestedOfferMappingEntriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetSuggestedOfferMappingEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuggestedOfferMappingEntries`: GetSuggestedOfferMappingEntriesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetSuggestedOfferMappingEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedOfferMappingEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSuggestedOfferMappingEntriesRequest** | [**GetSuggestedOfferMappingEntriesRequest**](GetSuggestedOfferMappingEntriesRequest.md) |  | 

### Return type

[**GetSuggestedOfferMappingEntriesResponse**](GetSuggestedOfferMappingEntriesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestedOfferMappings

> GetSuggestedOfferMappingsResponse GetSuggestedOfferMappings(ctx, businessId).GetSuggestedOfferMappingsRequest(getSuggestedOfferMappingsRequest).Execute()

Просмотр карточек на Маркете, которые подходят вашим товарам



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	getSuggestedOfferMappingsRequest := *openapiclient.NewGetSuggestedOfferMappingsRequest() // GetSuggestedOfferMappingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetSuggestedOfferMappings(context.Background(), businessId).GetSuggestedOfferMappingsRequest(getSuggestedOfferMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetSuggestedOfferMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuggestedOfferMappings`: GetSuggestedOfferMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetSuggestedOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedOfferMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSuggestedOfferMappingsRequest** | [**GetSuggestedOfferMappingsRequest**](GetSuggestedOfferMappingsRequest.md) |  | 

### Return type

[**GetSuggestedOfferMappingsResponse**](GetSuggestedOfferMappingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestedPrices

> SuggestPricesResponse GetSuggestedPrices(ctx, campaignId).SuggestPricesRequest(suggestPricesRequest).Execute()

Цены для продвижения товаров



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
	suggestPricesRequest := *openapiclient.NewSuggestPricesRequest([]openapiclient.SuggestOfferPriceDTO{*openapiclient.NewSuggestOfferPriceDTO()}) // SuggestPricesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetSuggestedPrices(context.Background(), campaignId).SuggestPricesRequest(suggestPricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetSuggestedPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuggestedPrices`: SuggestPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetSuggestedPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suggestPricesRequest** | [**SuggestPricesRequest**](SuggestPricesRequest.md) |  | 

### Return type

[**SuggestPricesResponse**](SuggestPricesResponse.md)

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.GetWarehouses(context.Background(), businessId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.GetWarehouses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarehouses`: GetWarehousesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.GetWarehouses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

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


## ProvideOrderDigitalCodes

> EmptyApiResponse ProvideOrderDigitalCodes(ctx, campaignId, orderId).ProvideOrderDigitalCodesRequest(provideOrderDigitalCodesRequest).Execute()

Передача ключей цифровых товаров



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	provideOrderDigitalCodesRequest := *openapiclient.NewProvideOrderDigitalCodesRequest([]openapiclient.OrderDigitalItemDTO{*openapiclient.NewOrderDigitalItemDTO(int64(123), "Slip_example", time.Now())}) // ProvideOrderDigitalCodesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.ProvideOrderDigitalCodes(context.Background(), campaignId, orderId).ProvideOrderDigitalCodesRequest(provideOrderDigitalCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.ProvideOrderDigitalCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvideOrderDigitalCodes`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.ProvideOrderDigitalCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvideOrderDigitalCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provideOrderDigitalCodesRequest** | [**ProvideOrderDigitalCodesRequest**](ProvideOrderDigitalCodesRequest.md) |  | 

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


## ProvideOrderItemIdentifiers

> ProvideOrderItemIdentifiersResponse ProvideOrderItemIdentifiers(ctx, campaignId, orderId).ProvideOrderItemIdentifiersRequest(provideOrderItemIdentifiersRequest).Execute()

Передача кодов маркировки единиц товара



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	provideOrderItemIdentifiersRequest := *openapiclient.NewProvideOrderItemIdentifiersRequest([]openapiclient.OrderItemInstanceModificationDTO{*openapiclient.NewOrderItemInstanceModificationDTO(int64(123), []openapiclient.BriefOrderItemInstanceDTO{*openapiclient.NewBriefOrderItemInstanceDTO()})}) // ProvideOrderItemIdentifiersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.ProvideOrderItemIdentifiers(context.Background(), campaignId, orderId).ProvideOrderItemIdentifiersRequest(provideOrderItemIdentifiersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.ProvideOrderItemIdentifiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvideOrderItemIdentifiers`: ProvideOrderItemIdentifiersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.ProvideOrderItemIdentifiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvideOrderItemIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provideOrderItemIdentifiersRequest** | [**ProvideOrderItemIdentifiersRequest**](ProvideOrderItemIdentifiersRequest.md) |  | 

### Return type

[**ProvideOrderItemIdentifiersResponse**](ProvideOrderItemIdentifiersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBidsForBusiness

> EmptyApiResponse PutBidsForBusiness(ctx, businessId).PutSkuBidsRequest(putSkuBidsRequest).Execute()

Включение буста продаж и установка ставок



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	putSkuBidsRequest := *openapiclient.NewPutSkuBidsRequest([]openapiclient.SkuBidItemDTO{*openapiclient.NewSkuBidItemDTO("Sku_example", int32(570))}) // PutSkuBidsRequest | description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.PutBidsForBusiness(context.Background(), businessId).PutSkuBidsRequest(putSkuBidsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.PutBidsForBusiness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutBidsForBusiness`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.PutBidsForBusiness`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBidsForBusinessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putSkuBidsRequest** | [**PutSkuBidsRequest**](PutSkuBidsRequest.md) | description | 

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


## PutBidsForCampaign

> EmptyApiResponse PutBidsForCampaign(ctx, campaignId).PutSkuBidsRequest(putSkuBidsRequest).Execute()

Включение буста продаж и установка ставок для магазина



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
	putSkuBidsRequest := *openapiclient.NewPutSkuBidsRequest([]openapiclient.SkuBidItemDTO{*openapiclient.NewSkuBidItemDTO("Sku_example", int32(570))}) // PutSkuBidsRequest | description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.PutBidsForCampaign(context.Background(), campaignId).PutSkuBidsRequest(putSkuBidsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.PutBidsForCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutBidsForCampaign`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.PutBidsForCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBidsForCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putSkuBidsRequest** | [**PutSkuBidsRequest**](PutSkuBidsRequest.md) | description | 

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


## SearchModels

> SearchModelsResponse SearchModels(ctx).Query(query).RegionId(regionId).Currency(currency).Page(page).PageSize(pageSize).Execute()

Поиск модели товара



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
	query := "query_example" // string | Поисковый запрос по названию модели товара.
	regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md). 
	currency := openapiclient.CurrencyType("RUR") // CurrencyType | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * `BYN` — белорусский рубль.  * `KZT` — казахстанский тенге.  * `RUR` — российский рубль.  * `UAH` — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  (optional)
	page := int32(56) // int32 | {% note warning \"Если в методе есть `page_token`\" %}  Используйте его вместо параметра `page`.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром `page_size`.  `page_number` игнорируется, если задан `page_token` или `limit`.  (optional) (default to 1)
	pageSize := int32(56) // int32 | Размер страницы.  Используется вместе с параметром `page_number`.  `page_size` игнорируется, если задан `page_token` или `limit`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SearchModels(context.Background()).Query(query).RegionId(regionId).Currency(currency).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SearchModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchModels`: SearchModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SearchModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Поисковый запрос по названию модели товара. | 
 **regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md).  | 
 **currency** | [**CurrencyType**](CurrencyType.md) | Валюта, в которой отображаются цены предложений на страницах с результатами поиска.  Возможные значения:  * &#x60;BYN&#x60; — белорусский рубль.  * &#x60;KZT&#x60; — казахстанский тенге.  * &#x60;RUR&#x60; — российский рубль.  * &#x60;UAH&#x60; — украинская гривна.  Значение по умолчанию: используется национальная валюта магазина (национальная валюта страны происхождения магазина).  | 
 **page** | **int32** | {% note warning \&quot;Если в методе есть &#x60;page_token&#x60;\&quot; %}  Используйте его вместо параметра &#x60;page&#x60;.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром &#x60;page_size&#x60;.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | [default to 1]
 **pageSize** | **int32** | Размер страницы.  Используется вместе с параметром &#x60;page_number&#x60;.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | 

### Return type

[**SearchModelsResponse**](SearchModelsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRegionChildren

> GetRegionWithChildrenResponse SearchRegionChildren(ctx, regionId).Page(page).PageSize(pageSize).Execute()

Информация о дочерних регионах



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
	regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md). 
	page := int32(56) // int32 | {% note warning \"Если в методе есть `page_token`\" %}  Используйте его вместо параметра `page`.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром `page_size`.  `page_number` игнорируется, если задан `page_token` или `limit`.  (optional) (default to 1)
	pageSize := int32(56) // int32 | Размер страницы.  Используется вместе с параметром `page_number`.  `page_size` игнорируется, если задан `page_token` или `limit`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SearchRegionChildren(context.Background(), regionId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SearchRegionChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRegionChildren`: GetRegionWithChildrenResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SearchRegionChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRegionChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | {% note warning \&quot;Если в методе есть &#x60;page_token&#x60;\&quot; %}  Используйте его вместо параметра &#x60;page&#x60;.  [Подробнее о типах пагинации и их использовании](../../concepts/pagination.md)  {% endnote %}  Номер страницы результатов.  Используется вместе с параметром &#x60;page_size&#x60;.  &#x60;page_number&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | [default to 1]
 **pageSize** | **int32** | Размер страницы.  Используется вместе с параметром &#x60;page_number&#x60;.  &#x60;page_size&#x60; игнорируется, если задан &#x60;page_token&#x60; или &#x60;limit&#x60;.  | 

### Return type

[**GetRegionWithChildrenResponse**](GetRegionWithChildrenResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRegionsById

> GetRegionsResponse SearchRegionsById(ctx, regionId).Execute()

Информация о регионе



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
	regionId := int64(789) // int64 | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md). 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SearchRegionsById(context.Background(), regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SearchRegionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRegionsById`: GetRegionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SearchRegionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regionId** | **int64** | Идентификатор региона.  Идентификатор региона можно получить c помощью запроса [GET regions](../../reference/regions/searchRegionsByName.md).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRegionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRegionsResponse**](GetRegionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRegionsByName

> GetRegionsResponse SearchRegionsByName(ctx).Name(name).PageToken(pageToken).Limit(limit).Execute()

Поиск регионов по их имени



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
	name := "name_example" // string | Название региона.  Важно учитывать регистр: первая буква должна быть заглавной, остальные — строчными. Например, `Москва`. 
	pageToken := "eyBuZXh0SWQ6IDIzNDIgfQ==" // string | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра `nextPageToken`, полученное при последнем запросе.  Если задан `page_token` и в запросе есть параметры `page_number` и `page_size`, они игнорируются.  (optional)
	limit := int32(20) // int32 | Количество значений на одной странице.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SearchRegionsByName(context.Background()).Name(name).PageToken(pageToken).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SearchRegionsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRegionsByName`: GetRegionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SearchRegionsByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRegionsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Название региона.  Важно учитывать регистр: первая буква должна быть заглавной, остальные — строчными. Например, &#x60;Москва&#x60;.  | 
 **pageToken** | **string** | Идентификатор страницы c результатами.  Если параметр не указан, возвращается первая страница.  Рекомендуем передавать значение выходного параметра &#x60;nextPageToken&#x60;, полученное при последнем запросе.  Если задан &#x60;page_token&#x60; и в запросе есть параметры &#x60;page_number&#x60; и &#x60;page_size&#x60;, они игнорируются.  | 
 **limit** | **int32** | Количество значений на одной странице.  | 

### Return type

[**GetRegionsResponse**](GetRegionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendFileToChat

> EmptyApiResponse SendFileToChat(ctx, businessId).ChatId(chatId).File(file).Execute()

Отправка файла в чат



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	chatId := int64(789) // int64 | Идентификатор чата.
	file := os.NewFile(1234, "some_file") // *os.File | Содержимое файла. Максимальный размер файла — 5 Мбайт.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SendFileToChat(context.Background(), businessId).ChatId(chatId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SendFileToChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFileToChat`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SendFileToChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendFileToChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatId** | **int64** | Идентификатор чата. | 
 **file** | ***os.File** | Содержимое файла. Максимальный размер файла — 5 Мбайт. | 

### Return type

[**EmptyApiResponse**](EmptyApiResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessageToChat

> EmptyApiResponse SendMessageToChat(ctx, businessId).ChatId(chatId).SendMessageToChatRequest(sendMessageToChatRequest).Execute()

Отправка сообщения в чат



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	chatId := int64(789) // int64 | Идентификатор чата.
	sendMessageToChatRequest := *openapiclient.NewSendMessageToChatRequest("Message_example") // SendMessageToChatRequest | description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SendMessageToChat(context.Background(), businessId).ChatId(chatId).SendMessageToChatRequest(sendMessageToChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SendMessageToChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMessageToChat`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SendMessageToChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageToChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatId** | **int64** | Идентификатор чата. | 
 **sendMessageToChatRequest** | [**SendMessageToChatRequest**](SendMessageToChatRequest.md) | description | 

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


## SetOrderBoxLayout

> SetOrderBoxLayoutResponse SetOrderBoxLayout(ctx, campaignId, orderId).SetOrderBoxLayoutRequest(setOrderBoxLayoutRequest).Execute()

Подготовка заказа



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	setOrderBoxLayoutRequest := *openapiclient.NewSetOrderBoxLayoutRequest([]openapiclient.OrderBoxLayoutDTO{*openapiclient.NewOrderBoxLayoutDTO([]openapiclient.OrderBoxLayoutItemDTO{*openapiclient.NewOrderBoxLayoutItemDTO(int64(123))})}) // SetOrderBoxLayoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SetOrderBoxLayout(context.Background(), campaignId, orderId).SetOrderBoxLayoutRequest(setOrderBoxLayoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SetOrderBoxLayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrderBoxLayout`: SetOrderBoxLayoutResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SetOrderBoxLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrderBoxLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setOrderBoxLayoutRequest** | [**SetOrderBoxLayoutRequest**](SetOrderBoxLayoutRequest.md) |  | 

### Return type

[**SetOrderBoxLayoutResponse**](SetOrderBoxLayoutResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrderDeliveryDate

> EmptyApiResponse SetOrderDeliveryDate(ctx, campaignId, orderId).SetOrderDeliveryDateRequest(setOrderDeliveryDateRequest).Execute()

Изменение даты доставки заказа



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	setOrderDeliveryDateRequest := *openapiclient.NewSetOrderDeliveryDateRequest(*openapiclient.NewOrderDeliveryDateDTO(time.Now()), openapiclient.OrderDeliveryDateReasonType("USER_MOVED_DELIVERY_DATES")) // SetOrderDeliveryDateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SetOrderDeliveryDate(context.Background(), campaignId, orderId).SetOrderDeliveryDateRequest(setOrderDeliveryDateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SetOrderDeliveryDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrderDeliveryDate`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SetOrderDeliveryDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrderDeliveryDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setOrderDeliveryDateRequest** | [**SetOrderDeliveryDateRequest**](SetOrderDeliveryDateRequest.md) |  | 

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


## SetOrderDeliveryTrackCode

> EmptyApiResponse SetOrderDeliveryTrackCode(ctx, campaignId, orderId).SetOrderDeliveryTrackCodeRequest(setOrderDeliveryTrackCodeRequest).Execute()

Передача трек‑номера посылки



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	setOrderDeliveryTrackCodeRequest := *openapiclient.NewSetOrderDeliveryTrackCodeRequest("TrackCode_example", int64(123)) // SetOrderDeliveryTrackCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SetOrderDeliveryTrackCode(context.Background(), campaignId, orderId).SetOrderDeliveryTrackCodeRequest(setOrderDeliveryTrackCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SetOrderDeliveryTrackCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrderDeliveryTrackCode`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SetOrderDeliveryTrackCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrderDeliveryTrackCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setOrderDeliveryTrackCodeRequest** | [**SetOrderDeliveryTrackCodeRequest**](SetOrderDeliveryTrackCodeRequest.md) |  | 

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


## SetOrderShipmentBoxes

> SetOrderShipmentBoxesResponse SetOrderShipmentBoxes(ctx, campaignId, orderId, shipmentId).SetOrderShipmentBoxesRequest(setOrderShipmentBoxesRequest).Execute()

Передача количества грузовых мест в заказе



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	shipmentId := int64(789) // int64 | Параметр больше не используется. Вставьте любое число — просто чтобы получился корректный URL. 
	setOrderShipmentBoxesRequest := *openapiclient.NewSetOrderShipmentBoxesRequest([]openapiclient.ParcelBoxRequestDTO{*openapiclient.NewParcelBoxRequestDTO()}) // SetOrderShipmentBoxesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SetOrderShipmentBoxes(context.Background(), campaignId, orderId, shipmentId).SetOrderShipmentBoxesRequest(setOrderShipmentBoxesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SetOrderShipmentBoxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrderShipmentBoxes`: SetOrderShipmentBoxesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SetOrderShipmentBoxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 
**shipmentId** | **int64** | Параметр больше не используется. Вставьте любое число — просто чтобы получился корректный URL.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrderShipmentBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setOrderShipmentBoxesRequest** | [**SetOrderShipmentBoxesRequest**](SetOrderShipmentBoxesRequest.md) |  | 

### Return type

[**SetOrderShipmentBoxesResponse**](SetOrderShipmentBoxesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.
	setReturnDecisionRequest := *openapiclient.NewSetReturnDecisionRequest(int64(123), openapiclient.ReturnRequestDecisionType("REFUND_MONEY")) // SetReturnDecisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SetReturnDecision(context.Background(), campaignId, orderId, returnId).SetReturnDecisionRequest(setReturnDecisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SetReturnDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetReturnDecision`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SetReturnDecision`: %v\n", resp)
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


## SkipGoodsFeedbacksReaction

> EmptyApiResponse SkipGoodsFeedbacksReaction(ctx, businessId).SkipGoodsFeedbackReactionRequest(skipGoodsFeedbackReactionRequest).Execute()

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	skipGoodsFeedbackReactionRequest := *openapiclient.NewSkipGoodsFeedbackReactionRequest([]int64{int64(123)}) // SkipGoodsFeedbackReactionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SkipGoodsFeedbacksReaction(context.Background(), businessId).SkipGoodsFeedbackReactionRequest(skipGoodsFeedbackReactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SkipGoodsFeedbacksReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SkipGoodsFeedbacksReaction`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SkipGoodsFeedbacksReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSkipGoodsFeedbacksReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipGoodsFeedbackReactionRequest** | [**SkipGoodsFeedbackReactionRequest**](SkipGoodsFeedbackReactionRequest.md) |  | 

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
	openapiclient "github.com/ucoms-dev/yandex-market-go-client"
)

func main() {
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	returnId := int64(789) // int64 | Идентификатор невыкупа или возврата.
	body := map[string]interface{}{ ... } // map[string]interface{} | description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.SubmitReturnDecision(context.Background(), campaignId, orderId, returnId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.SubmitReturnDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitReturnDecision`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.SubmitReturnDecision`: %v\n", resp)
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


## UpdateBusinessPrices

> EmptyApiResponse UpdateBusinessPrices(ctx, businessId).UpdateBusinessPricesRequest(updateBusinessPricesRequest).Execute()

Установка цен на товары для всех магазинов



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	updateBusinessPricesRequest := *openapiclient.NewUpdateBusinessPricesRequest([]openapiclient.UpdateBusinessOfferPriceDTO{*openapiclient.NewUpdateBusinessOfferPriceDTO("OfferId_example", *openapiclient.NewUpdateBusinessPricesDTO(float32(123), openapiclient.CurrencyType("RUR")))}) // UpdateBusinessPricesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateBusinessPrices(context.Background(), businessId).UpdateBusinessPricesRequest(updateBusinessPricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateBusinessPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBusinessPrices`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateBusinessPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBusinessPricesRequest** | [**UpdateBusinessPricesRequest**](UpdateBusinessPricesRequest.md) |  | 

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


## UpdateCampaignOffers

> EmptyApiResponse UpdateCampaignOffers(ctx, campaignId).UpdateCampaignOffersRequest(updateCampaignOffersRequest).Execute()

Изменение условий продажи товаров в магазине



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
	updateCampaignOffersRequest := *openapiclient.NewUpdateCampaignOffersRequest([]openapiclient.UpdateCampaignOfferDTO{*openapiclient.NewUpdateCampaignOfferDTO("OfferId_example")}) // UpdateCampaignOffersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateCampaignOffers(context.Background(), campaignId).UpdateCampaignOffersRequest(updateCampaignOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateCampaignOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCampaignOffers`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateCampaignOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignOffersRequest** | [**UpdateCampaignOffersRequest**](UpdateCampaignOffersRequest.md) |  | 

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


## UpdateExternalOrderId

> EmptyApiResponse UpdateExternalOrderId(ctx, campaignId, orderId).UpdateExternalOrderIdRequest(updateExternalOrderIdRequest).Execute()

Передача или изменение дополнительного идентификатора заказа



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	updateExternalOrderIdRequest := *openapiclient.NewUpdateExternalOrderIdRequest("ExternalOrderId_example") // UpdateExternalOrderIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateExternalOrderId(context.Background(), campaignId, orderId).UpdateExternalOrderIdRequest(updateExternalOrderIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateExternalOrderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalOrderId`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateExternalOrderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalOrderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateExternalOrderIdRequest** | [**UpdateExternalOrderIdRequest**](UpdateExternalOrderIdRequest.md) |  | 

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

> UpdateGoodsFeedbackCommentResponse UpdateGoodsFeedbackComment(ctx, businessId).UpdateGoodsFeedbackCommentRequest(updateGoodsFeedbackCommentRequest).Execute()

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	updateGoodsFeedbackCommentRequest := *openapiclient.NewUpdateGoodsFeedbackCommentRequest(int64(123), *openapiclient.NewUpdateGoodsFeedbackCommentDTO("Text_example")) // UpdateGoodsFeedbackCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateGoodsFeedbackComment(context.Background(), businessId).UpdateGoodsFeedbackCommentRequest(updateGoodsFeedbackCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateGoodsFeedbackComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGoodsFeedbackComment`: UpdateGoodsFeedbackCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateGoodsFeedbackComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGoodsFeedbackCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGoodsFeedbackCommentRequest** | [**UpdateGoodsFeedbackCommentRequest**](UpdateGoodsFeedbackCommentRequest.md) |  | 

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	updateOfferContentRequest := *openapiclient.NewUpdateOfferContentRequest([]openapiclient.OfferContentDTO{*openapiclient.NewOfferContentDTO("OfferId_example", int32(123), []openapiclient.ParameterValueDTO{*openapiclient.NewParameterValueDTO(int64(123))})}) // UpdateOfferContentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOfferContent(context.Background(), businessId).UpdateOfferContentRequest(updateOfferContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOfferContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOfferContent`: UpdateOfferContentResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOfferContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

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


## UpdateOfferMappingEntries

> EmptyApiResponse UpdateOfferMappingEntries(ctx, campaignId).UpdateOfferMappingEntryRequest(updateOfferMappingEntryRequest).Execute()

Добавление и редактирование товаров в каталоге



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
	updateOfferMappingEntryRequest := *openapiclient.NewUpdateOfferMappingEntryRequest([]openapiclient.UpdateOfferMappingEntryDTO{*openapiclient.NewUpdateOfferMappingEntryDTO()}) // UpdateOfferMappingEntryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOfferMappingEntries(context.Background(), campaignId).UpdateOfferMappingEntryRequest(updateOfferMappingEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOfferMappingEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOfferMappingEntries`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOfferMappingEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfferMappingEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOfferMappingEntryRequest** | [**UpdateOfferMappingEntryRequest**](UpdateOfferMappingEntryRequest.md) |  | 

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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	updateOfferMappingsRequest := *openapiclient.NewUpdateOfferMappingsRequest([]openapiclient.UpdateOfferMappingDTO{*openapiclient.NewUpdateOfferMappingDTO(*openapiclient.NewUpdateOfferDTO("OfferId_example"))}) // UpdateOfferMappingsRequest | 
	language := openapiclient.CatalogLanguageType("RU") // CatalogLanguageType | Язык, на котором принимаются и возвращаются значения в параметрах `name` и `description`.  Значение по умолчанию: `RU`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOfferMappings(context.Background(), businessId).UpdateOfferMappingsRequest(updateOfferMappingsRequest).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOfferMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOfferMappings`: UpdateOfferMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOfferMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

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


## UpdateOrderItems

> UpdateOrderItems(ctx, campaignId, orderId).UpdateOrderItemRequest(updateOrderItemRequest).Execute()

Удаление товара из заказа или уменьшение числа единиц



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	updateOrderItemRequest := *openapiclient.NewUpdateOrderItemRequest([]openapiclient.OrderItemModificationDTO{*openapiclient.NewOrderItemModificationDTO(int64(123), int32(123))}) // UpdateOrderItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DbsAPI.UpdateOrderItems(context.Background(), campaignId, orderId).UpdateOrderItemRequest(updateOrderItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOrderItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderItemRequest** | [**UpdateOrderItemRequest**](UpdateOrderItemRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderStatus

> UpdateOrderStatusResponse UpdateOrderStatus(ctx, campaignId, orderId).UpdateOrderStatusRequest(updateOrderStatusRequest).Execute()

Изменение статуса одного заказа



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
	orderId := int64(789) // int64 | Идентификатор заказа.
	updateOrderStatusRequest := *openapiclient.NewUpdateOrderStatusRequest(*openapiclient.NewOrderStatusChangeDTO(openapiclient.OrderStatusType("PLACING"))) // UpdateOrderStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOrderStatus(context.Background(), campaignId, orderId).UpdateOrderStatusRequest(updateOrderStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOrderStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrderStatus`: UpdateOrderStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOrderStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderStatusRequest** | [**UpdateOrderStatusRequest**](UpdateOrderStatusRequest.md) |  | 

### Return type

[**UpdateOrderStatusResponse**](UpdateOrderStatusResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderStatuses

> UpdateOrderStatusesResponse UpdateOrderStatuses(ctx, campaignId).UpdateOrderStatusesRequest(updateOrderStatusesRequest).Execute()

Изменение статусов нескольких заказов



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
	updateOrderStatusesRequest := *openapiclient.NewUpdateOrderStatusesRequest([]openapiclient.OrderStateDTO{*openapiclient.NewOrderStateDTO(int64(123), openapiclient.OrderStatusType("PLACING"))}) // UpdateOrderStatusesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOrderStatuses(context.Background(), campaignId).UpdateOrderStatusesRequest(updateOrderStatusesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOrderStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrderStatuses`: UpdateOrderStatusesResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOrderStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrderStatusesRequest** | [**UpdateOrderStatusesRequest**](UpdateOrderStatusesRequest.md) |  | 

### Return type

[**UpdateOrderStatusesResponse**](UpdateOrderStatusesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrderStorageLimit

> EmptyApiResponse UpdateOrderStorageLimit(ctx, campaignId, orderId).UpdateOrderStorageLimitRequest(updateOrderStorageLimitRequest).Execute()

Продление срока хранения заказа



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	orderId := int64(789) // int64 | Идентификатор заказа.
	updateOrderStorageLimitRequest := *openapiclient.NewUpdateOrderStorageLimitRequest(time.Now()) // UpdateOrderStorageLimitRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOrderStorageLimit(context.Background(), campaignId, orderId).UpdateOrderStorageLimitRequest(updateOrderStorageLimitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOrderStorageLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrderStorageLimit`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOrderStorageLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**orderId** | **int64** | Идентификатор заказа. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderStorageLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderStorageLimitRequest** | [**UpdateOrderStorageLimitRequest**](UpdateOrderStorageLimitRequest.md) |  | 

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


## UpdateOutlet

> EmptyApiResponse UpdateOutlet(ctx, campaignId, outletId).ChangeOutletRequest(changeOutletRequest).Execute()

Изменение информации о точке продаж



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
	outletId := int64(789) // int64 | Идентификатор точки продаж.
	changeOutletRequest := *openapiclient.NewChangeOutletRequest("Name_example", openapiclient.OutletType("DEPOT"), *openapiclient.NewOutletAddressDTO(int64(123)), []string{"Phones_example"}, *openapiclient.NewOutletWorkingScheduleDTO([]openapiclient.OutletWorkingScheduleItemDTO{*openapiclient.NewOutletWorkingScheduleItemDTO(openapiclient.DayOfWeekType("MONDAY"), openapiclient.DayOfWeekType("MONDAY"), "09:59", "23:59")})) // ChangeOutletRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOutlet(context.Background(), campaignId, outletId).ChangeOutletRequest(changeOutletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOutlet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOutlet`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOutlet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 
**outletId** | **int64** | Идентификатор точки продаж. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **changeOutletRequest** | [**ChangeOutletRequest**](ChangeOutletRequest.md) |  | 

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


## UpdateOutletLicenses

> EmptyApiResponse UpdateOutletLicenses(ctx, campaignId).UpdateOutletLicenseRequest(updateOutletLicenseRequest).Execute()

Создание и изменение лицензий для точек продаж



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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	updateOutletLicenseRequest := *openapiclient.NewUpdateOutletLicenseRequest([]openapiclient.OutletLicenseDTO{*openapiclient.NewOutletLicenseDTO(int64(123), openapiclient.LicenseType("ALCOHOL"), "Number_example", time.Now(), time.Now())}) // UpdateOutletLicenseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateOutletLicenses(context.Background(), campaignId).UpdateOutletLicenseRequest(updateOutletLicenseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateOutletLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOutletLicenses`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateOutletLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutletLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOutletLicenseRequest** | [**UpdateOutletLicenseRequest**](UpdateOutletLicenseRequest.md) |  | 

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


## UpdatePrices

> EmptyApiResponse UpdatePrices(ctx, campaignId).UpdatePricesRequest(updatePricesRequest).Execute()

Установка цен на товары в конкретном магазине



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
	updatePricesRequest := *openapiclient.NewUpdatePricesRequest([]openapiclient.OfferPriceDTO{*openapiclient.NewOfferPriceDTO()}) // UpdatePricesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdatePrices(context.Background(), campaignId).UpdatePricesRequest(updatePricesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdatePrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrices`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdatePrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePricesRequest** | [**UpdatePricesRequest**](UpdatePricesRequest.md) |  | 

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


## UpdatePromoOffers

> UpdatePromoOffersResponse UpdatePromoOffers(ctx, businessId).UpdatePromoOffersRequest(updatePromoOffersRequest).Execute()

Добавление товаров в акцию или изменение их цен



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
	businessId := int64(789) // int64 | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html) 
	updatePromoOffersRequest := *openapiclient.NewUpdatePromoOffersRequest("PromoId_example", []openapiclient.UpdatePromoOfferDTO{*openapiclient.NewUpdatePromoOfferDTO("OfferId_example")}) // UpdatePromoOffersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdatePromoOffers(context.Background(), businessId).UpdatePromoOffersRequest(updatePromoOffersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdatePromoOffers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePromoOffers`: UpdatePromoOffersResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdatePromoOffers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessId** | **int64** | Идентификатор кабинета. Чтобы его узнать, воспользуйтесь запросом [GET campaigns](../../reference/campaigns/getCampaigns.md).  ℹ️ [Что такое кабинет и магазин на Маркете](https://yandex.ru/support/marketplace/account/introduction.html)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePromoOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePromoOffersRequest** | [**UpdatePromoOffersRequest**](UpdatePromoOffersRequest.md) |  | 

### Return type

[**UpdatePromoOffersResponse**](UpdatePromoOffersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [OAuth](../README.md#OAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStocks

> EmptyApiResponse UpdateStocks(ctx, campaignId).UpdateStocksRequest(updateStocksRequest).Execute()

Передача информации об остатках



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
	updateStocksRequest := *openapiclient.NewUpdateStocksRequest([]openapiclient.UpdateStockDTO{*openapiclient.NewUpdateStockDTO("Sku_example", []openapiclient.UpdateStockItemDTO{*openapiclient.NewUpdateStockItemDTO(int64(123))})}) // UpdateStocksRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateStocks(context.Background(), campaignId).UpdateStocksRequest(updateStocksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateStocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStocks`: EmptyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateStocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStocksRequest** | [**UpdateStocksRequest**](UpdateStocksRequest.md) |  | 

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
	campaignId := int64(789) // int64 | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах. 
	updateWarehouseStatusRequest := *openapiclient.NewUpdateWarehouseStatusRequest(false) // UpdateWarehouseStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbsAPI.UpdateWarehouseStatus(context.Background(), campaignId).UpdateWarehouseStatusRequest(updateWarehouseStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbsAPI.UpdateWarehouseStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWarehouseStatus`: UpdateWarehouseStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DbsAPI.UpdateWarehouseStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int64** | Идентификатор кампании.  Его можно узнать с помощью запроса [GET campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете — нажмите на название своего бизнеса и перейдите на страницу:    * **Модули и API** → блок **Передача данных Маркету**.   * **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не передавайте вместо него идентификатор магазина, который указан в кабинете продавца на Маркете рядом с названием магазина и в некоторых отчетах.  | 

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

