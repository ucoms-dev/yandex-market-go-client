# CalculateTariffsParametersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** | Идентификатор кампании (магазина) — технический идентификатор, который представляет ваш магазин в системе Яндекс Маркета при работе через API. Он однозначно связывается с вашим магазином, но предназначен только для автоматизированного взаимодействия.  Его можно узнать с помощью запроса [GET v2/campaigns](../../reference/campaigns/getCampaigns.md) или найти в кабинете продавца на Маркете. Нажмите на иконку вашего аккаунта → **Настройки** и в меню слева выберите **API и модули**:  * блок **Идентификатор кампании**; * вкладка **Лог запросов** → выпадающий список в блоке **Показывать логи**.  ⚠️ Не путайте его с: - идентификатором магазина, который отображается в личном кабинете продавца; - рекламными кампаниями.  | [optional] 
**SellingProgram** | Pointer to [**SellingProgramType**](SellingProgramType.md) |  | [optional] 
**Frequency** | Pointer to [**PaymentFrequencyType**](PaymentFrequencyType.md) |  | [optional] 
**PaymentDelayWeeks** | Pointer to **int32** | Отсрочка выплат при еженедельном графике — сколько недель назад были доставлены заказы, за которые приходит выплата.  Допустимые значения: 0, 1, 2 или 4.  Значения параметра &#x60;paymentDelayWeeks&#x60;, отличные от 0, допускаются только вместе с параметром &#x60;frequency&#x60; равным &#39;WEEKLY&#39;. Использование других значений параметра &#x60;frequency&#x60; совместно с &#x60;paymentDelayWeeks&#x60; приведет к ошибке.  | [optional] 
**Currency** | Pointer to [**CurrencyType**](CurrencyType.md) |  | [optional] 

## Methods

### NewCalculateTariffsParametersDTO

`func NewCalculateTariffsParametersDTO() *CalculateTariffsParametersDTO`

NewCalculateTariffsParametersDTO instantiates a new CalculateTariffsParametersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculateTariffsParametersDTOWithDefaults

`func NewCalculateTariffsParametersDTOWithDefaults() *CalculateTariffsParametersDTO`

NewCalculateTariffsParametersDTOWithDefaults instantiates a new CalculateTariffsParametersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *CalculateTariffsParametersDTO) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CalculateTariffsParametersDTO) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CalculateTariffsParametersDTO) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CalculateTariffsParametersDTO) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetSellingProgram

`func (o *CalculateTariffsParametersDTO) GetSellingProgram() SellingProgramType`

GetSellingProgram returns the SellingProgram field if non-nil, zero value otherwise.

### GetSellingProgramOk

`func (o *CalculateTariffsParametersDTO) GetSellingProgramOk() (*SellingProgramType, bool)`

GetSellingProgramOk returns a tuple with the SellingProgram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingProgram

`func (o *CalculateTariffsParametersDTO) SetSellingProgram(v SellingProgramType)`

SetSellingProgram sets SellingProgram field to given value.

### HasSellingProgram

`func (o *CalculateTariffsParametersDTO) HasSellingProgram() bool`

HasSellingProgram returns a boolean if a field has been set.

### GetFrequency

`func (o *CalculateTariffsParametersDTO) GetFrequency() PaymentFrequencyType`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CalculateTariffsParametersDTO) GetFrequencyOk() (*PaymentFrequencyType, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CalculateTariffsParametersDTO) SetFrequency(v PaymentFrequencyType)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CalculateTariffsParametersDTO) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetPaymentDelayWeeks

`func (o *CalculateTariffsParametersDTO) GetPaymentDelayWeeks() int32`

GetPaymentDelayWeeks returns the PaymentDelayWeeks field if non-nil, zero value otherwise.

### GetPaymentDelayWeeksOk

`func (o *CalculateTariffsParametersDTO) GetPaymentDelayWeeksOk() (*int32, bool)`

GetPaymentDelayWeeksOk returns a tuple with the PaymentDelayWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDelayWeeks

`func (o *CalculateTariffsParametersDTO) SetPaymentDelayWeeks(v int32)`

SetPaymentDelayWeeks sets PaymentDelayWeeks field to given value.

### HasPaymentDelayWeeks

`func (o *CalculateTariffsParametersDTO) HasPaymentDelayWeeks() bool`

HasPaymentDelayWeeks returns a boolean if a field has been set.

### GetCurrency

`func (o *CalculateTariffsParametersDTO) GetCurrency() CurrencyType`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CalculateTariffsParametersDTO) GetCurrencyOk() (*CurrencyType, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CalculateTariffsParametersDTO) SetCurrency(v CurrencyType)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CalculateTariffsParametersDTO) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


