# TariffParameterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | {% cut \&quot;Название параметра.\&quot; %}  Параметр &#x60;name&#x60; может принимать следующие значения:  **Основные параметры:**   - &#x60;value&#x60; — значение тарифа.   - &#x60;billingUnit&#x60; — единицы измерения для товара.   - &#x60;valueType&#x60; — тип значения тарифа: &#x60;absolute&#x60; (денежное значение) или &#x60;relative&#x60; (процент).   - &#x60;priceDependence&#x60; — зависимость тарифа от цены товара.  **Параметры диапазонов цен товара:** - &#x60;priceFrom&#x60; — минимальная цена товара, от которой применяется тариф. - &#x60;priceTo&#x60; — максимальная цена товара, до которой применяется тариф.  **Параметры ограничений значения тарифа:** - &#x60;minValue&#x60; — минимальное значение тарифа. - &#x60;maxValue&#x60; — максимальное значение тарифа.  **Параметры диапазонов дней:** - &#x60;dayFrom&#x60; — день, с которого начинает действовать тариф (для хранения). - &#x60;dayTo&#x60; — день, на который тариф прекращает действие (для хранения).  **Параметры маршрута транзитной поставки:** - &#x60;xdocFrom&#x60; — начальная точка маршрута. - &#x60;xdocTo&#x60; — конечная точка маршрута.  **Параметры частоты:** - &#x60;frequency&#x60; — частота выплат. - &#x60;paymentDelayWeeks&#x60; — отсрочка выплат при еженедельном графике — сколько недель назад были доставлены заказы, за которые приходит выплата.  **Параметры складов и логистики:** - &#x60;transitWarehouseType&#x60; — тип места привоза товаров продавцом. - &#x60;orderCargoType&#x60; — тип груза заказа.  **Параметры оборачиваемости:** - &#x60;turnoverFrom&#x60; — нижняя граница диапазона оборачиваемости, начиная с которой применяется тариф. - &#x60;turnoverTo&#x60; — верхняя граница диапазона оборачиваемости, исключая которую применяется тариф.  **Параметры рейтинга:** - &#x60;ratingLowerBound&#x60; — нижняя граница диапазона рейтинга, по которому применяется тариф. - &#x60;ratingUpperBound&#x60; — верхняя граница диапазона рейтинга, по которому применяется тариф.  **Параметры количества заказов:** - &#x60;fromOrders&#x60; — количество заказов, начиная с которого применяется тариф (включительно). - &#x60;toOrders&#x60; — количество заказов, до которого применяется тариф (не включительно).  **Параметры доставки из-за рубежа:** - &#x60;shippingPrice&#x60; — стоимость отправления при доставке из-за рубежа. - &#x60;totalSurchargeForWeight&#x60; — сумма наценки за вес отправления при доставке из-за рубежа.  {% note info \&quot;Не все параметры возвращаются для каждого тарифа.\&quot; %}  Набор параметров зависит от типа услуги и условий применения тарифа.  {% endnote %}  {% endcut %}  | 
**Value** | **string** | Значение параметра. | 

## Methods

### NewTariffParameterDTO

`func NewTariffParameterDTO(name string, value string, ) *TariffParameterDTO`

NewTariffParameterDTO instantiates a new TariffParameterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffParameterDTOWithDefaults

`func NewTariffParameterDTOWithDefaults() *TariffParameterDTO`

NewTariffParameterDTOWithDefaults instantiates a new TariffParameterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TariffParameterDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TariffParameterDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TariffParameterDTO) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *TariffParameterDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TariffParameterDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TariffParameterDTO) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


