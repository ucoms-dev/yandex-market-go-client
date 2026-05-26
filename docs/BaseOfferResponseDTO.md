# BaseOfferResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfferId** | **string** | Ваш SKU — идентификатор товара в вашей системе.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  SKU товара можно изменить в кабинете продавца на Маркете. О том, как это сделать, читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/assortment/operations/edit-sku).  {% note warning %}  Пробельные символы в начале и конце значения автоматически удаляются. Например, &#x60;\&quot;  SKU123  \&quot;&#x60; и &#x60;\&quot;SKU123\&quot;&#x60; будут обработаны как одинаковые значения.  {% endnote %}  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields)  | 
**Name** | Pointer to **string** | Составляйте название по схеме: тип + бренд или производитель + модель + особенности, если есть (например, цвет, размер или вес) и количество в упаковке.  Не включайте в название условия продажи (например, «скидка», «бесплатная доставка» и т. д.), эмоциональные характеристики («хит», «супер» и т. д.). Не пишите слова большими буквами — кроме устоявшихся названий брендов и моделей.  Оптимальная длина — 50–60 символов.  [Рекомендации и правила](https://yandex.ru/support/marketplace/assortment/fields/title.html)  | [optional] 
**MarketCategoryId** | Pointer to **int64** | Идентификатор категории на Маркете, к которой вы относите свой товар.  {% note warning \&quot;Всегда указывайте, когда передаете &#x60;parameterValues&#x60;\&quot; %}  Если при изменении характеристик передать &#x60;parameterValues&#x60; и не указать &#x60;marketCategoryId&#x60;, характеристики обновятся, но в ответе придет предупреждение (параметр &#x60;warnings&#x60;).  Если не передать их оба, будет использована информация из устаревших параметров &#x60;params&#x60; и &#x60;category&#x60;, а &#x60;marketCategoryId&#x60; будет определен автоматически.  {% endnote %}  При изменении категории убедитесь, что характеристики товара и их значения в параметре &#x60;parameterValues&#x60; вы передаете для новой категории.  Список категорий Маркета можно получить с помощью запроса  [POST v2/categories/tree](../../reference/categories/getCategoriesTree.md).  | [optional] 
**Category** | Pointer to **string** | {% note warning \&quot;Вместо него используйте &#x60;marketCategoryId&#x60;.\&quot; %}     {% endnote %}  Категория товара в вашем магазине.  | [optional] 
**Pictures** | Pointer to **[]string** | Ссылки на изображения товара. Изображение по первой ссылке считается основным, остальные дополнительными.  **Требования к ссылкам**  * Ссылок может быть до 30. * Указывайте ссылку целиком, включая протокол http или https. * Максимальная длина — 512 символов. * Русские буквы в URL можно. * Можно использовать прямые ссылки на изображения и на Яндекс Диск. Ссылки на Яндекс Диске нужно копировать с помощью функции **Поделиться**. Относительные ссылки и ссылки на другие облачные хранилища — не работают.  ✅ &#x60;https://example-shop.ru/images/sku12345.jpg&#x60;  ✅ &#x60;https://yadi.sk/i/NaBoRsimVOLov&#x60;  ❌ &#x60;/images/sku12345.jpg&#x60;  ❌ &#x60;https://www.dropbox.com/s/818f/tovar.jpg&#x60;  Ссылки на изображение должны быть постоянными. Нельзя использовать динамические ссылки, меняющиеся от выгрузки к выгрузке.  Если нужно заменить изображение, выложите новое изображение по новой ссылке, а ссылку на старое удалите. Если просто заменить изображение по старой ссылке, оно не обновится.  [Требования к изображениям](https://yandex.ru/support/marketplace/assortment/fields/images.html)  | [optional] 
**Videos** | Pointer to **[]string** | Ссылки (URL) на видео товара.  **Требования к ссылке**  * Указывайте ссылку целиком, включая протокол http или https. * Максимальная длина — 512 символов. * Русские буквы в URL можно. * Можно использовать прямые ссылки на видео и на Яндекс Диск. Ссылки на Яндекс Диске нужно копировать с помощью функции **Поделиться**. Относительные ссылки и ссылки на другие облачные хранилища — не работают.  ✅ &#x60;https://example-shop.ru/video/sku12345.avi&#x60;  ✅ &#x60;https://yadi.sk/i/NaBoRsimVOLov&#x60;  ❌ &#x60;/video/sku12345.avi&#x60;  ❌ &#x60;https://www.dropbox.com/s/818f/super-tovar.avi&#x60;  Ссылки на видео должны быть постоянными. Нельзя использовать динамические ссылки, меняющиеся от выгрузки к выгрузке.  Если нужно заменить видео, выложите новое видео по новой ссылке, а ссылку на старое удалите. Если просто заменить видео по старой ссылке, оно не обновится.  [Требования к видео](https://yandex.ru/support/marketplace/assortment/fields/video.html)  | [optional] 
**Manuals** | Pointer to [**[]OfferManualDTO**](OfferManualDTO.md) | Список инструкций по использованию товара.  | [optional] 
**Vendor** | Pointer to **string** | Название бренда или производителя. Должно быть записано так, как его пишет сам бренд. | [optional] 
**Barcodes** | Pointer to **[]string** | Штрихкод.  Указывайте в виде последовательности цифр. Подойдут коды :no-translate[EAN-13, EAN-8, UPC-A, UPC-E] или :no-translate[Code 128]. Для книг — :no-translate[ISBN].  Для товаров [определенных категорий и торговых марок](https://yastatic.net/s3/doc-binary/src/support/market/ru/yandex-market-list-for-gtin.xlsx) штрихкод должен быть действительным кодом :no-translate[GTIN]. Обратите внимание: внутренние штрихкоды, начинающиеся на 2 или 02, и коды формата :no-translate[Code 128] не являются :no-translate[GTIN].  [Что такое :no-translate[GTIN]](*gtin)  | [optional] 
**Description** | Pointer to **string** | Подробное описание товара: например, его преимущества и особенности.  Не давайте в описании инструкций по установке и сборке. Не используйте слова «скидка», «распродажа», «дешевый», «подарок» (кроме подарочных категорий), «бесплатно», «акция», «специальная цена», «новинка», «new», «аналог», «заказ», «хит». Не указывайте никакой контактной информации и не давайте ссылок.  Для форматирования текста можно использовать теги HTML:  * \\&lt;h&gt;, \\&lt;h1&gt;, \\&lt;h2&gt; и так далее — для заголовков; * \\&lt;br&gt; и \\&lt;p&gt; — для переноса строки; * \\&lt;ol&gt; — для нумерованного списка; * \\&lt;ul&gt; — для маркированного списка; * \\&lt;li&gt; — для создания элементов списка (должен находиться внутри \\&lt;ol&gt; или \\&lt;ul&gt;); * \\&lt;div&gt; — поддерживается, но не влияет на отображение текста.  Оптимальная длина — 400–600 символов.  [Рекомендации и правила](https://yandex.ru/support/marketplace/assortment/fields/description.html)  | [optional] 
**ManufacturerCountries** | Pointer to **[]string** | Страна, где был произведен товар.  Записывайте названия стран так, как они записаны в [списке](https://yastatic.net/s3/doc-binary/src/support/market/ru/countries.xlsx).  | [optional] 
**WeightDimensions** | Pointer to [**OfferWeightDimensionsDTO**](OfferWeightDimensionsDTO.md) |  | [optional] 
**VendorCode** | Pointer to **string** | Артикул товара от производителя. | [optional] 
**Tags** | Pointer to **[]string** | Метки товара, которые использует магазин. Покупателям теги не видны. По тегам можно группировать и фильтровать разные товары в каталоге — например, товары одной серии, коллекции или линейки.  Максимальная длина тега — 20 символов. У одного товара может быть максимум 10 тегов.  | [optional] 
**ShelfLife** | Pointer to [**TimePeriodDTO**](TimePeriodDTO.md) |  | [optional] 
**LifeTime** | Pointer to [**TimePeriodDTO**](TimePeriodDTO.md) |  | [optional] 
**GuaranteePeriod** | Pointer to [**TimePeriodDTO**](TimePeriodDTO.md) |  | [optional] 
**CustomsCommodityCode** | Pointer to **string** | {% note warning \&quot;Вместо него используйте &#x60;commodityCodes&#x60; с типом &#x60;CUSTOMS_COMMODITY_CODE&#x60;.\&quot; %}     {% endnote %}  Код товара в единой Товарной номенклатуре внешнеэкономической деятельности (ТН ВЭД) — 10 или 14 цифр без пробелов.  Обязательно укажите, если он есть.  | [optional] 
**CommodityCodes** | Pointer to [**[]CommodityCodeDTO**](CommodityCodeDTO.md) | Товарные коды.  | [optional] 
**Certificates** | Pointer to **[]string** | Номера документов на товар: сертификата, декларации соответствия и т. п.  Передавать можно только номера документов, сканы которого загружены в кабинете продавца по [инструкции](https://yandex.ru/support/marketplace/assortment/restrictions/certificates.html).  | [optional] 
**BoxCount** | Pointer to **int32** | Количество грузовых мест.  Параметр используется, если товар представляет собой несколько коробок, упаковок и так далее. Например, кондиционер занимает два места — внешний и внутренний блоки в двух коробках.  Для товаров, занимающих одно место, не передавайте этот параметр.  | [optional] 
**Condition** | Pointer to [**OfferConditionDTO**](OfferConditionDTO.md) |  | [optional] 
**Type** | Pointer to [**OfferType**](OfferType.md) |  | [optional] 
**Downloadable** | Pointer to **bool** | Признак цифрового товара. Укажите &#x60;true&#x60;, если товар доставляется по электронной почте.  [Как работать с цифровыми товарами](../../step-by-step/digital.md)  | [optional] 
**Adult** | Pointer to **bool** | Параметр включает для товара пометку 18+. Устанавливайте ее только для товаров, которые относятся к удовлетворению сексуальных потребностей.  | [optional] 
**Age** | Pointer to [**AgeDTO**](AgeDTO.md) |  | [optional] 
**Params** | Pointer to [**[]OfferParamDTO**](OfferParamDTO.md) | {% note warning \&quot;При передаче характеристик используйте &#x60;parameterValues&#x60;.\&quot; %}     {% endnote %}  Характеристики, которые есть только у товаров конкретной категории — например, диаметр колес велосипеда или материал подошвы обуви.  | [optional] 

## Methods

### NewBaseOfferResponseDTO

`func NewBaseOfferResponseDTO(offerId string, ) *BaseOfferResponseDTO`

NewBaseOfferResponseDTO instantiates a new BaseOfferResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseOfferResponseDTOWithDefaults

`func NewBaseOfferResponseDTOWithDefaults() *BaseOfferResponseDTO`

NewBaseOfferResponseDTOWithDefaults instantiates a new BaseOfferResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfferId

`func (o *BaseOfferResponseDTO) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *BaseOfferResponseDTO) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *BaseOfferResponseDTO) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetName

`func (o *BaseOfferResponseDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseOfferResponseDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseOfferResponseDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseOfferResponseDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMarketCategoryId

`func (o *BaseOfferResponseDTO) GetMarketCategoryId() int64`

GetMarketCategoryId returns the MarketCategoryId field if non-nil, zero value otherwise.

### GetMarketCategoryIdOk

`func (o *BaseOfferResponseDTO) GetMarketCategoryIdOk() (*int64, bool)`

GetMarketCategoryIdOk returns a tuple with the MarketCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCategoryId

`func (o *BaseOfferResponseDTO) SetMarketCategoryId(v int64)`

SetMarketCategoryId sets MarketCategoryId field to given value.

### HasMarketCategoryId

`func (o *BaseOfferResponseDTO) HasMarketCategoryId() bool`

HasMarketCategoryId returns a boolean if a field has been set.

### GetCategory

`func (o *BaseOfferResponseDTO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BaseOfferResponseDTO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BaseOfferResponseDTO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BaseOfferResponseDTO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPictures

`func (o *BaseOfferResponseDTO) GetPictures() []string`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *BaseOfferResponseDTO) GetPicturesOk() (*[]string, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *BaseOfferResponseDTO) SetPictures(v []string)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *BaseOfferResponseDTO) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### SetPicturesNil

`func (o *BaseOfferResponseDTO) SetPicturesNil(b bool)`

 SetPicturesNil sets the value for Pictures to be an explicit nil

### UnsetPictures
`func (o *BaseOfferResponseDTO) UnsetPictures()`

UnsetPictures ensures that no value is present for Pictures, not even an explicit nil
### GetVideos

`func (o *BaseOfferResponseDTO) GetVideos() []string`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *BaseOfferResponseDTO) GetVideosOk() (*[]string, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *BaseOfferResponseDTO) SetVideos(v []string)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *BaseOfferResponseDTO) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### SetVideosNil

`func (o *BaseOfferResponseDTO) SetVideosNil(b bool)`

 SetVideosNil sets the value for Videos to be an explicit nil

### UnsetVideos
`func (o *BaseOfferResponseDTO) UnsetVideos()`

UnsetVideos ensures that no value is present for Videos, not even an explicit nil
### GetManuals

`func (o *BaseOfferResponseDTO) GetManuals() []OfferManualDTO`

GetManuals returns the Manuals field if non-nil, zero value otherwise.

### GetManualsOk

`func (o *BaseOfferResponseDTO) GetManualsOk() (*[]OfferManualDTO, bool)`

GetManualsOk returns a tuple with the Manuals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuals

`func (o *BaseOfferResponseDTO) SetManuals(v []OfferManualDTO)`

SetManuals sets Manuals field to given value.

### HasManuals

`func (o *BaseOfferResponseDTO) HasManuals() bool`

HasManuals returns a boolean if a field has been set.

### SetManualsNil

`func (o *BaseOfferResponseDTO) SetManualsNil(b bool)`

 SetManualsNil sets the value for Manuals to be an explicit nil

### UnsetManuals
`func (o *BaseOfferResponseDTO) UnsetManuals()`

UnsetManuals ensures that no value is present for Manuals, not even an explicit nil
### GetVendor

`func (o *BaseOfferResponseDTO) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BaseOfferResponseDTO) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BaseOfferResponseDTO) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *BaseOfferResponseDTO) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetBarcodes

`func (o *BaseOfferResponseDTO) GetBarcodes() []string`

GetBarcodes returns the Barcodes field if non-nil, zero value otherwise.

### GetBarcodesOk

`func (o *BaseOfferResponseDTO) GetBarcodesOk() (*[]string, bool)`

GetBarcodesOk returns a tuple with the Barcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodes

`func (o *BaseOfferResponseDTO) SetBarcodes(v []string)`

SetBarcodes sets Barcodes field to given value.

### HasBarcodes

`func (o *BaseOfferResponseDTO) HasBarcodes() bool`

HasBarcodes returns a boolean if a field has been set.

### SetBarcodesNil

`func (o *BaseOfferResponseDTO) SetBarcodesNil(b bool)`

 SetBarcodesNil sets the value for Barcodes to be an explicit nil

### UnsetBarcodes
`func (o *BaseOfferResponseDTO) UnsetBarcodes()`

UnsetBarcodes ensures that no value is present for Barcodes, not even an explicit nil
### GetDescription

`func (o *BaseOfferResponseDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseOfferResponseDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseOfferResponseDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseOfferResponseDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManufacturerCountries

`func (o *BaseOfferResponseDTO) GetManufacturerCountries() []string`

GetManufacturerCountries returns the ManufacturerCountries field if non-nil, zero value otherwise.

### GetManufacturerCountriesOk

`func (o *BaseOfferResponseDTO) GetManufacturerCountriesOk() (*[]string, bool)`

GetManufacturerCountriesOk returns a tuple with the ManufacturerCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerCountries

`func (o *BaseOfferResponseDTO) SetManufacturerCountries(v []string)`

SetManufacturerCountries sets ManufacturerCountries field to given value.

### HasManufacturerCountries

`func (o *BaseOfferResponseDTO) HasManufacturerCountries() bool`

HasManufacturerCountries returns a boolean if a field has been set.

### SetManufacturerCountriesNil

`func (o *BaseOfferResponseDTO) SetManufacturerCountriesNil(b bool)`

 SetManufacturerCountriesNil sets the value for ManufacturerCountries to be an explicit nil

### UnsetManufacturerCountries
`func (o *BaseOfferResponseDTO) UnsetManufacturerCountries()`

UnsetManufacturerCountries ensures that no value is present for ManufacturerCountries, not even an explicit nil
### GetWeightDimensions

`func (o *BaseOfferResponseDTO) GetWeightDimensions() OfferWeightDimensionsDTO`

GetWeightDimensions returns the WeightDimensions field if non-nil, zero value otherwise.

### GetWeightDimensionsOk

`func (o *BaseOfferResponseDTO) GetWeightDimensionsOk() (*OfferWeightDimensionsDTO, bool)`

GetWeightDimensionsOk returns a tuple with the WeightDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightDimensions

`func (o *BaseOfferResponseDTO) SetWeightDimensions(v OfferWeightDimensionsDTO)`

SetWeightDimensions sets WeightDimensions field to given value.

### HasWeightDimensions

`func (o *BaseOfferResponseDTO) HasWeightDimensions() bool`

HasWeightDimensions returns a boolean if a field has been set.

### GetVendorCode

`func (o *BaseOfferResponseDTO) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *BaseOfferResponseDTO) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *BaseOfferResponseDTO) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *BaseOfferResponseDTO) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.

### GetTags

`func (o *BaseOfferResponseDTO) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BaseOfferResponseDTO) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BaseOfferResponseDTO) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BaseOfferResponseDTO) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BaseOfferResponseDTO) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BaseOfferResponseDTO) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetShelfLife

`func (o *BaseOfferResponseDTO) GetShelfLife() TimePeriodDTO`

GetShelfLife returns the ShelfLife field if non-nil, zero value otherwise.

### GetShelfLifeOk

`func (o *BaseOfferResponseDTO) GetShelfLifeOk() (*TimePeriodDTO, bool)`

GetShelfLifeOk returns a tuple with the ShelfLife field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShelfLife

`func (o *BaseOfferResponseDTO) SetShelfLife(v TimePeriodDTO)`

SetShelfLife sets ShelfLife field to given value.

### HasShelfLife

`func (o *BaseOfferResponseDTO) HasShelfLife() bool`

HasShelfLife returns a boolean if a field has been set.

### GetLifeTime

`func (o *BaseOfferResponseDTO) GetLifeTime() TimePeriodDTO`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *BaseOfferResponseDTO) GetLifeTimeOk() (*TimePeriodDTO, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *BaseOfferResponseDTO) SetLifeTime(v TimePeriodDTO)`

SetLifeTime sets LifeTime field to given value.

### HasLifeTime

`func (o *BaseOfferResponseDTO) HasLifeTime() bool`

HasLifeTime returns a boolean if a field has been set.

### GetGuaranteePeriod

`func (o *BaseOfferResponseDTO) GetGuaranteePeriod() TimePeriodDTO`

GetGuaranteePeriod returns the GuaranteePeriod field if non-nil, zero value otherwise.

### GetGuaranteePeriodOk

`func (o *BaseOfferResponseDTO) GetGuaranteePeriodOk() (*TimePeriodDTO, bool)`

GetGuaranteePeriodOk returns a tuple with the GuaranteePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteePeriod

`func (o *BaseOfferResponseDTO) SetGuaranteePeriod(v TimePeriodDTO)`

SetGuaranteePeriod sets GuaranteePeriod field to given value.

### HasGuaranteePeriod

`func (o *BaseOfferResponseDTO) HasGuaranteePeriod() bool`

HasGuaranteePeriod returns a boolean if a field has been set.

### GetCustomsCommodityCode

`func (o *BaseOfferResponseDTO) GetCustomsCommodityCode() string`

GetCustomsCommodityCode returns the CustomsCommodityCode field if non-nil, zero value otherwise.

### GetCustomsCommodityCodeOk

`func (o *BaseOfferResponseDTO) GetCustomsCommodityCodeOk() (*string, bool)`

GetCustomsCommodityCodeOk returns a tuple with the CustomsCommodityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsCommodityCode

`func (o *BaseOfferResponseDTO) SetCustomsCommodityCode(v string)`

SetCustomsCommodityCode sets CustomsCommodityCode field to given value.

### HasCustomsCommodityCode

`func (o *BaseOfferResponseDTO) HasCustomsCommodityCode() bool`

HasCustomsCommodityCode returns a boolean if a field has been set.

### GetCommodityCodes

`func (o *BaseOfferResponseDTO) GetCommodityCodes() []CommodityCodeDTO`

GetCommodityCodes returns the CommodityCodes field if non-nil, zero value otherwise.

### GetCommodityCodesOk

`func (o *BaseOfferResponseDTO) GetCommodityCodesOk() (*[]CommodityCodeDTO, bool)`

GetCommodityCodesOk returns a tuple with the CommodityCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommodityCodes

`func (o *BaseOfferResponseDTO) SetCommodityCodes(v []CommodityCodeDTO)`

SetCommodityCodes sets CommodityCodes field to given value.

### HasCommodityCodes

`func (o *BaseOfferResponseDTO) HasCommodityCodes() bool`

HasCommodityCodes returns a boolean if a field has been set.

### SetCommodityCodesNil

`func (o *BaseOfferResponseDTO) SetCommodityCodesNil(b bool)`

 SetCommodityCodesNil sets the value for CommodityCodes to be an explicit nil

### UnsetCommodityCodes
`func (o *BaseOfferResponseDTO) UnsetCommodityCodes()`

UnsetCommodityCodes ensures that no value is present for CommodityCodes, not even an explicit nil
### GetCertificates

`func (o *BaseOfferResponseDTO) GetCertificates() []string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *BaseOfferResponseDTO) GetCertificatesOk() (*[]string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *BaseOfferResponseDTO) SetCertificates(v []string)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *BaseOfferResponseDTO) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *BaseOfferResponseDTO) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *BaseOfferResponseDTO) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetBoxCount

`func (o *BaseOfferResponseDTO) GetBoxCount() int32`

GetBoxCount returns the BoxCount field if non-nil, zero value otherwise.

### GetBoxCountOk

`func (o *BaseOfferResponseDTO) GetBoxCountOk() (*int32, bool)`

GetBoxCountOk returns a tuple with the BoxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoxCount

`func (o *BaseOfferResponseDTO) SetBoxCount(v int32)`

SetBoxCount sets BoxCount field to given value.

### HasBoxCount

`func (o *BaseOfferResponseDTO) HasBoxCount() bool`

HasBoxCount returns a boolean if a field has been set.

### GetCondition

`func (o *BaseOfferResponseDTO) GetCondition() OfferConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *BaseOfferResponseDTO) GetConditionOk() (*OfferConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *BaseOfferResponseDTO) SetCondition(v OfferConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *BaseOfferResponseDTO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetType

`func (o *BaseOfferResponseDTO) GetType() OfferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseOfferResponseDTO) GetTypeOk() (*OfferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseOfferResponseDTO) SetType(v OfferType)`

SetType sets Type field to given value.

### HasType

`func (o *BaseOfferResponseDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDownloadable

`func (o *BaseOfferResponseDTO) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *BaseOfferResponseDTO) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *BaseOfferResponseDTO) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *BaseOfferResponseDTO) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetAdult

`func (o *BaseOfferResponseDTO) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *BaseOfferResponseDTO) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *BaseOfferResponseDTO) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *BaseOfferResponseDTO) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetAge

`func (o *BaseOfferResponseDTO) GetAge() AgeDTO`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *BaseOfferResponseDTO) GetAgeOk() (*AgeDTO, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *BaseOfferResponseDTO) SetAge(v AgeDTO)`

SetAge sets Age field to given value.

### HasAge

`func (o *BaseOfferResponseDTO) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetParams

`func (o *BaseOfferResponseDTO) GetParams() []OfferParamDTO`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *BaseOfferResponseDTO) GetParamsOk() (*[]OfferParamDTO, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *BaseOfferResponseDTO) SetParams(v []OfferParamDTO)`

SetParams sets Params field to given value.

### HasParams

`func (o *BaseOfferResponseDTO) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *BaseOfferResponseDTO) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *BaseOfferResponseDTO) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


