type DbsApiAcceptOrderCancellationRequest struct {
func (r DbsApiAcceptOrderCancellationRequest) AcceptOrderCancellationRequest(acceptOrderCancellationRequest AcceptOrderCancellationRequest) DbsApiAcceptOrderCancellationRequest {
func (r DbsApiAcceptOrderCancellationRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiAcceptOrderCancellationRequest
func (a *DbsAPIService) AcceptOrderCancellation(ctx context.Context, campaignId int64, orderId int64) DbsApiAcceptOrderCancellationRequest {
	return DbsApiAcceptOrderCancellationRequest{
func (a *DbsAPIService) AcceptOrderCancellationExecute(r DbsApiAcceptOrderCancellationRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiAddHiddenOffersRequest struct {
func (r DbsApiAddHiddenOffersRequest) AddHiddenOffersRequest(addHiddenOffersRequest AddHiddenOffersRequest) DbsApiAddHiddenOffersRequest {
func (r DbsApiAddHiddenOffersRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiAddHiddenOffersRequest
func (a *DbsAPIService) AddHiddenOffers(ctx context.Context, campaignId int64) DbsApiAddHiddenOffersRequest {
	return DbsApiAddHiddenOffersRequest{
func (a *DbsAPIService) AddHiddenOffersExecute(r DbsApiAddHiddenOffersRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiAddOffersToArchiveRequest struct {
func (r DbsApiAddOffersToArchiveRequest) AddOffersToArchiveRequest(addOffersToArchiveRequest AddOffersToArchiveRequest) DbsApiAddOffersToArchiveRequest {
func (r DbsApiAddOffersToArchiveRequest) Execute() (*AddOffersToArchiveResponse, *http.Response, error) {
	@return DbsApiAddOffersToArchiveRequest
func (a *DbsAPIService) AddOffersToArchive(ctx context.Context, businessId int64) DbsApiAddOffersToArchiveRequest {
	return DbsApiAddOffersToArchiveRequest{
func (a *DbsAPIService) AddOffersToArchiveExecute(r DbsApiAddOffersToArchiveRequest) (*AddOffersToArchiveResponse, *http.Response, error) {
type DbsApiCalculateTariffsRequest struct {
func (r DbsApiCalculateTariffsRequest) CalculateTariffsRequest(calculateTariffsRequest CalculateTariffsRequest) DbsApiCalculateTariffsRequest {
func (r DbsApiCalculateTariffsRequest) Execute() (*CalculateTariffsResponse, *http.Response, error) {
	@return DbsApiCalculateTariffsRequest
func (a *DbsAPIService) CalculateTariffs(ctx context.Context) DbsApiCalculateTariffsRequest {
	return DbsApiCalculateTariffsRequest{
func (a *DbsAPIService) CalculateTariffsExecute(r DbsApiCalculateTariffsRequest) (*CalculateTariffsResponse, *http.Response, error) {
type DbsApiConfirmBusinessPricesRequest struct {
func (r DbsApiConfirmBusinessPricesRequest) ConfirmPricesRequest(confirmPricesRequest ConfirmPricesRequest) DbsApiConfirmBusinessPricesRequest {
func (r DbsApiConfirmBusinessPricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiConfirmBusinessPricesRequest
func (a *DbsAPIService) ConfirmBusinessPrices(ctx context.Context, businessId int64) DbsApiConfirmBusinessPricesRequest {
	return DbsApiConfirmBusinessPricesRequest{
func (a *DbsAPIService) ConfirmBusinessPricesExecute(r DbsApiConfirmBusinessPricesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiConfirmCampaignPricesRequest struct {
func (r DbsApiConfirmCampaignPricesRequest) ConfirmPricesRequest(confirmPricesRequest ConfirmPricesRequest) DbsApiConfirmCampaignPricesRequest {
func (r DbsApiConfirmCampaignPricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiConfirmCampaignPricesRequest
func (a *DbsAPIService) ConfirmCampaignPrices(ctx context.Context, campaignId int64) DbsApiConfirmCampaignPricesRequest {
	return DbsApiConfirmCampaignPricesRequest{
func (a *DbsAPIService) ConfirmCampaignPricesExecute(r DbsApiConfirmCampaignPricesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiCreateChatRequest struct {
func (r DbsApiCreateChatRequest) CreateChatRequest(createChatRequest CreateChatRequest) DbsApiCreateChatRequest {
func (r DbsApiCreateChatRequest) Execute() (*CreateChatResponse, *http.Response, error) {
	@return DbsApiCreateChatRequest
func (a *DbsAPIService) CreateChat(ctx context.Context, businessId int64) DbsApiCreateChatRequest {
	return DbsApiCreateChatRequest{
func (a *DbsAPIService) CreateChatExecute(r DbsApiCreateChatRequest) (*CreateChatResponse, *http.Response, error) {
type DbsApiCreateOutletRequest struct {
func (r DbsApiCreateOutletRequest) ChangeOutletRequest(changeOutletRequest ChangeOutletRequest) DbsApiCreateOutletRequest {
func (r DbsApiCreateOutletRequest) Execute() (*CreateOutletResponse, *http.Response, error) {
	@return DbsApiCreateOutletRequest
func (a *DbsAPIService) CreateOutlet(ctx context.Context, campaignId int64) DbsApiCreateOutletRequest {
	return DbsApiCreateOutletRequest{
func (a *DbsAPIService) CreateOutletExecute(r DbsApiCreateOutletRequest) (*CreateOutletResponse, *http.Response, error) {
type DbsApiDeleteCampaignOffersRequest struct {
func (r DbsApiDeleteCampaignOffersRequest) DeleteCampaignOffersRequest(deleteCampaignOffersRequest DeleteCampaignOffersRequest) DbsApiDeleteCampaignOffersRequest {
func (r DbsApiDeleteCampaignOffersRequest) Execute() (*DeleteCampaignOffersResponse, *http.Response, error) {
	@return DbsApiDeleteCampaignOffersRequest
func (a *DbsAPIService) DeleteCampaignOffers(ctx context.Context, campaignId int64) DbsApiDeleteCampaignOffersRequest {
	return DbsApiDeleteCampaignOffersRequest{
func (a *DbsAPIService) DeleteCampaignOffersExecute(r DbsApiDeleteCampaignOffersRequest) (*DeleteCampaignOffersResponse, *http.Response, error) {
type DbsApiDeleteGoodsFeedbackCommentRequest struct {
func (r DbsApiDeleteGoodsFeedbackCommentRequest) DeleteGoodsFeedbackCommentRequest(deleteGoodsFeedbackCommentRequest DeleteGoodsFeedbackCommentRequest) DbsApiDeleteGoodsFeedbackCommentRequest {
func (r DbsApiDeleteGoodsFeedbackCommentRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiDeleteGoodsFeedbackCommentRequest
func (a *DbsAPIService) DeleteGoodsFeedbackComment(ctx context.Context, businessId int64) DbsApiDeleteGoodsFeedbackCommentRequest {
	return DbsApiDeleteGoodsFeedbackCommentRequest{
func (a *DbsAPIService) DeleteGoodsFeedbackCommentExecute(r DbsApiDeleteGoodsFeedbackCommentRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiDeleteHiddenOffersRequest struct {
func (r DbsApiDeleteHiddenOffersRequest) DeleteHiddenOffersRequest(deleteHiddenOffersRequest DeleteHiddenOffersRequest) DbsApiDeleteHiddenOffersRequest {
func (r DbsApiDeleteHiddenOffersRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiDeleteHiddenOffersRequest
func (a *DbsAPIService) DeleteHiddenOffers(ctx context.Context, campaignId int64) DbsApiDeleteHiddenOffersRequest {
	return DbsApiDeleteHiddenOffersRequest{
func (a *DbsAPIService) DeleteHiddenOffersExecute(r DbsApiDeleteHiddenOffersRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiDeleteOffersRequest struct {
func (r DbsApiDeleteOffersRequest) DeleteOffersRequest(deleteOffersRequest DeleteOffersRequest) DbsApiDeleteOffersRequest {
func (r DbsApiDeleteOffersRequest) Execute() (*DeleteOffersResponse, *http.Response, error) {
	@return DbsApiDeleteOffersRequest
func (a *DbsAPIService) DeleteOffers(ctx context.Context, businessId int64) DbsApiDeleteOffersRequest {
	return DbsApiDeleteOffersRequest{
func (a *DbsAPIService) DeleteOffersExecute(r DbsApiDeleteOffersRequest) (*DeleteOffersResponse, *http.Response, error) {
type DbsApiDeleteOffersFromArchiveRequest struct {
func (r DbsApiDeleteOffersFromArchiveRequest) DeleteOffersFromArchiveRequest(deleteOffersFromArchiveRequest DeleteOffersFromArchiveRequest) DbsApiDeleteOffersFromArchiveRequest {
func (r DbsApiDeleteOffersFromArchiveRequest) Execute() (*DeleteOffersFromArchiveResponse, *http.Response, error) {
	@return DbsApiDeleteOffersFromArchiveRequest
func (a *DbsAPIService) DeleteOffersFromArchive(ctx context.Context, businessId int64) DbsApiDeleteOffersFromArchiveRequest {
	return DbsApiDeleteOffersFromArchiveRequest{
func (a *DbsAPIService) DeleteOffersFromArchiveExecute(r DbsApiDeleteOffersFromArchiveRequest) (*DeleteOffersFromArchiveResponse, *http.Response, error) {
type DbsApiDeleteOutletRequest struct {
func (r DbsApiDeleteOutletRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiDeleteOutletRequest
func (a *DbsAPIService) DeleteOutlet(ctx context.Context, campaignId int64, outletId int64) DbsApiDeleteOutletRequest {
	return DbsApiDeleteOutletRequest{
func (a *DbsAPIService) DeleteOutletExecute(r DbsApiDeleteOutletRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiDeleteOutletLicensesRequest struct {
func (r DbsApiDeleteOutletLicensesRequest) Ids(ids []int64) DbsApiDeleteOutletLicensesRequest {
func (r DbsApiDeleteOutletLicensesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiDeleteOutletLicensesRequest
func (a *DbsAPIService) DeleteOutletLicenses(ctx context.Context, campaignId int64) DbsApiDeleteOutletLicensesRequest {
	return DbsApiDeleteOutletLicensesRequest{
func (a *DbsAPIService) DeleteOutletLicensesExecute(r DbsApiDeleteOutletLicensesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiDeletePromoOffersRequest struct {
func (r DbsApiDeletePromoOffersRequest) DeletePromoOffersRequest(deletePromoOffersRequest DeletePromoOffersRequest) DbsApiDeletePromoOffersRequest {
func (r DbsApiDeletePromoOffersRequest) Execute() (*DeletePromoOffersResponse, *http.Response, error) {
	@return DbsApiDeletePromoOffersRequest
func (a *DbsAPIService) DeletePromoOffers(ctx context.Context, businessId int64) DbsApiDeletePromoOffersRequest {
	return DbsApiDeletePromoOffersRequest{
func (a *DbsAPIService) DeletePromoOffersExecute(r DbsApiDeletePromoOffersRequest) (*DeletePromoOffersResponse, *http.Response, error) {
type DbsApiGenerateBannersStatisticsReportRequest struct {
func (r DbsApiGenerateBannersStatisticsReportRequest) GenerateBannersStatisticsRequest(generateBannersStatisticsRequest GenerateBannersStatisticsRequest) DbsApiGenerateBannersStatisticsReportRequest {
func (r DbsApiGenerateBannersStatisticsReportRequest) Format(format ReportFormatType) DbsApiGenerateBannersStatisticsReportRequest {
func (r DbsApiGenerateBannersStatisticsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateBannersStatisticsReportRequest
func (a *DbsAPIService) GenerateBannersStatisticsReport(ctx context.Context) DbsApiGenerateBannersStatisticsReportRequest {
	return DbsApiGenerateBannersStatisticsReportRequest{
func (a *DbsAPIService) GenerateBannersStatisticsReportExecute(r DbsApiGenerateBannersStatisticsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateBoostConsolidatedReportRequest struct {
func (r DbsApiGenerateBoostConsolidatedReportRequest) GenerateBoostConsolidatedRequest(generateBoostConsolidatedRequest GenerateBoostConsolidatedRequest) DbsApiGenerateBoostConsolidatedReportRequest {
func (r DbsApiGenerateBoostConsolidatedReportRequest) Format(format ReportFormatType) DbsApiGenerateBoostConsolidatedReportRequest {
func (r DbsApiGenerateBoostConsolidatedReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateBoostConsolidatedReportRequest
func (a *DbsAPIService) GenerateBoostConsolidatedReport(ctx context.Context) DbsApiGenerateBoostConsolidatedReportRequest {
	return DbsApiGenerateBoostConsolidatedReportRequest{
func (a *DbsAPIService) GenerateBoostConsolidatedReportExecute(r DbsApiGenerateBoostConsolidatedReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateCompetitorsPositionReportRequest struct {
func (r DbsApiGenerateCompetitorsPositionReportRequest) GenerateCompetitorsPositionReportRequest(generateCompetitorsPositionReportRequest GenerateCompetitorsPositionReportRequest) DbsApiGenerateCompetitorsPositionReportRequest {
func (r DbsApiGenerateCompetitorsPositionReportRequest) Format(format ReportFormatType) DbsApiGenerateCompetitorsPositionReportRequest {
func (r DbsApiGenerateCompetitorsPositionReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateCompetitorsPositionReportRequest
func (a *DbsAPIService) GenerateCompetitorsPositionReport(ctx context.Context) DbsApiGenerateCompetitorsPositionReportRequest {
	return DbsApiGenerateCompetitorsPositionReportRequest{
func (a *DbsAPIService) GenerateCompetitorsPositionReportExecute(r DbsApiGenerateCompetitorsPositionReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateGoodsFeedbackReportRequest struct {
func (r DbsApiGenerateGoodsFeedbackReportRequest) GenerateGoodsFeedbackRequest(generateGoodsFeedbackRequest GenerateGoodsFeedbackRequest) DbsApiGenerateGoodsFeedbackReportRequest {
func (r DbsApiGenerateGoodsFeedbackReportRequest) Format(format ReportFormatType) DbsApiGenerateGoodsFeedbackReportRequest {
func (r DbsApiGenerateGoodsFeedbackReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateGoodsFeedbackReportRequest
func (a *DbsAPIService) GenerateGoodsFeedbackReport(ctx context.Context) DbsApiGenerateGoodsFeedbackReportRequest {
	return DbsApiGenerateGoodsFeedbackReportRequest{
func (a *DbsAPIService) GenerateGoodsFeedbackReportExecute(r DbsApiGenerateGoodsFeedbackReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateGoodsRealizationReportRequest struct {
func (r DbsApiGenerateGoodsRealizationReportRequest) GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest GenerateGoodsRealizationReportRequest) DbsApiGenerateGoodsRealizationReportRequest {
func (r DbsApiGenerateGoodsRealizationReportRequest) Format(format ReportFormatType) DbsApiGenerateGoodsRealizationReportRequest {
func (r DbsApiGenerateGoodsRealizationReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateGoodsRealizationReportRequest
func (a *DbsAPIService) GenerateGoodsRealizationReport(ctx context.Context) DbsApiGenerateGoodsRealizationReportRequest {
	return DbsApiGenerateGoodsRealizationReportRequest{
func (a *DbsAPIService) GenerateGoodsRealizationReportExecute(r DbsApiGenerateGoodsRealizationReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateJewelryFiscalReportRequest struct {
func (r DbsApiGenerateJewelryFiscalReportRequest) GenerateJewelryFiscalReportRequest(generateJewelryFiscalReportRequest GenerateJewelryFiscalReportRequest) DbsApiGenerateJewelryFiscalReportRequest {
func (r DbsApiGenerateJewelryFiscalReportRequest) Format(format ReportFormatType) DbsApiGenerateJewelryFiscalReportRequest {
func (r DbsApiGenerateJewelryFiscalReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateJewelryFiscalReportRequest
func (a *DbsAPIService) GenerateJewelryFiscalReport(ctx context.Context) DbsApiGenerateJewelryFiscalReportRequest {
	return DbsApiGenerateJewelryFiscalReportRequest{
func (a *DbsAPIService) GenerateJewelryFiscalReportExecute(r DbsApiGenerateJewelryFiscalReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateMassOrderLabelsReportRequest struct {
func (r DbsApiGenerateMassOrderLabelsReportRequest) GenerateMassOrderLabelsRequest(generateMassOrderLabelsRequest GenerateMassOrderLabelsRequest) DbsApiGenerateMassOrderLabelsReportRequest {
func (r DbsApiGenerateMassOrderLabelsReportRequest) Format(format PageFormatType) DbsApiGenerateMassOrderLabelsReportRequest {
func (r DbsApiGenerateMassOrderLabelsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateMassOrderLabelsReportRequest
func (a *DbsAPIService) GenerateMassOrderLabelsReport(ctx context.Context) DbsApiGenerateMassOrderLabelsReportRequest {
	return DbsApiGenerateMassOrderLabelsReportRequest{
func (a *DbsAPIService) GenerateMassOrderLabelsReportExecute(r DbsApiGenerateMassOrderLabelsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateOrderLabelRequest struct {
func (r DbsApiGenerateOrderLabelRequest) Format(format PageFormatType) DbsApiGenerateOrderLabelRequest {
func (r DbsApiGenerateOrderLabelRequest) Execute() (*os.File, *http.Response, error) {
	@return DbsApiGenerateOrderLabelRequest
func (a *DbsAPIService) GenerateOrderLabel(ctx context.Context, campaignId int64, orderId int64, shipmentId int64, boxId int64) DbsApiGenerateOrderLabelRequest {
	return DbsApiGenerateOrderLabelRequest{
func (a *DbsAPIService) GenerateOrderLabelExecute(r DbsApiGenerateOrderLabelRequest) (*os.File, *http.Response, error) {
type DbsApiGenerateOrderLabelsRequest struct {
func (r DbsApiGenerateOrderLabelsRequest) Format(format PageFormatType) DbsApiGenerateOrderLabelsRequest {
func (r DbsApiGenerateOrderLabelsRequest) Execute() (*os.File, *http.Response, error) {
	@return DbsApiGenerateOrderLabelsRequest
func (a *DbsAPIService) GenerateOrderLabels(ctx context.Context, campaignId int64, orderId int64) DbsApiGenerateOrderLabelsRequest {
	return DbsApiGenerateOrderLabelsRequest{
func (a *DbsAPIService) GenerateOrderLabelsExecute(r DbsApiGenerateOrderLabelsRequest) (*os.File, *http.Response, error) {
type DbsApiGeneratePricesReportRequest struct {
func (r DbsApiGeneratePricesReportRequest) GeneratePricesReportRequest(generatePricesReportRequest GeneratePricesReportRequest) DbsApiGeneratePricesReportRequest {
func (r DbsApiGeneratePricesReportRequest) Format(format ReportFormatType) DbsApiGeneratePricesReportRequest {
func (r DbsApiGeneratePricesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGeneratePricesReportRequest
func (a *DbsAPIService) GeneratePricesReport(ctx context.Context) DbsApiGeneratePricesReportRequest {
	return DbsApiGeneratePricesReportRequest{
func (a *DbsAPIService) GeneratePricesReportExecute(r DbsApiGeneratePricesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateSalesGeographyReportRequest struct {
func (r DbsApiGenerateSalesGeographyReportRequest) GenerateSalesGeographyRequest(generateSalesGeographyRequest GenerateSalesGeographyRequest) DbsApiGenerateSalesGeographyReportRequest {
func (r DbsApiGenerateSalesGeographyReportRequest) Format(format ReportFormatType) DbsApiGenerateSalesGeographyReportRequest {
func (r DbsApiGenerateSalesGeographyReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateSalesGeographyReportRequest
func (a *DbsAPIService) GenerateSalesGeographyReport(ctx context.Context) DbsApiGenerateSalesGeographyReportRequest {
	return DbsApiGenerateSalesGeographyReportRequest{
func (a *DbsAPIService) GenerateSalesGeographyReportExecute(r DbsApiGenerateSalesGeographyReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateShelfsStatisticsReportRequest struct {
func (r DbsApiGenerateShelfsStatisticsReportRequest) GenerateShelfsStatisticsRequest(generateShelfsStatisticsRequest GenerateShelfsStatisticsRequest) DbsApiGenerateShelfsStatisticsReportRequest {
func (r DbsApiGenerateShelfsStatisticsReportRequest) Format(format ReportFormatType) DbsApiGenerateShelfsStatisticsReportRequest {
func (r DbsApiGenerateShelfsStatisticsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateShelfsStatisticsReportRequest
func (a *DbsAPIService) GenerateShelfsStatisticsReport(ctx context.Context) DbsApiGenerateShelfsStatisticsReportRequest {
	return DbsApiGenerateShelfsStatisticsReportRequest{
func (a *DbsAPIService) GenerateShelfsStatisticsReportExecute(r DbsApiGenerateShelfsStatisticsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateShowsBoostReportRequest struct {
func (r DbsApiGenerateShowsBoostReportRequest) GenerateShowsBoostRequest(generateShowsBoostRequest GenerateShowsBoostRequest) DbsApiGenerateShowsBoostReportRequest {
func (r DbsApiGenerateShowsBoostReportRequest) Format(format ReportFormatType) DbsApiGenerateShowsBoostReportRequest {
func (r DbsApiGenerateShowsBoostReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateShowsBoostReportRequest
func (a *DbsAPIService) GenerateShowsBoostReport(ctx context.Context) DbsApiGenerateShowsBoostReportRequest {
	return DbsApiGenerateShowsBoostReportRequest{
func (a *DbsAPIService) GenerateShowsBoostReportExecute(r DbsApiGenerateShowsBoostReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateShowsSalesReportRequest struct {
func (r DbsApiGenerateShowsSalesReportRequest) GenerateShowsSalesReportRequest(generateShowsSalesReportRequest GenerateShowsSalesReportRequest) DbsApiGenerateShowsSalesReportRequest {
func (r DbsApiGenerateShowsSalesReportRequest) Format(format ReportFormatType) DbsApiGenerateShowsSalesReportRequest {
func (r DbsApiGenerateShowsSalesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateShowsSalesReportRequest
func (a *DbsAPIService) GenerateShowsSalesReport(ctx context.Context) DbsApiGenerateShowsSalesReportRequest {
	return DbsApiGenerateShowsSalesReportRequest{
func (a *DbsAPIService) GenerateShowsSalesReportExecute(r DbsApiGenerateShowsSalesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateStocksOnWarehousesReportRequest struct {
func (r DbsApiGenerateStocksOnWarehousesReportRequest) GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest GenerateStocksOnWarehousesReportRequest) DbsApiGenerateStocksOnWarehousesReportRequest {
func (r DbsApiGenerateStocksOnWarehousesReportRequest) Format(format ReportFormatType) DbsApiGenerateStocksOnWarehousesReportRequest {
func (r DbsApiGenerateStocksOnWarehousesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateStocksOnWarehousesReportRequest
func (a *DbsAPIService) GenerateStocksOnWarehousesReport(ctx context.Context) DbsApiGenerateStocksOnWarehousesReportRequest {
	return DbsApiGenerateStocksOnWarehousesReportRequest{
func (a *DbsAPIService) GenerateStocksOnWarehousesReportExecute(r DbsApiGenerateStocksOnWarehousesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateUnitedMarketplaceServicesReportRequest struct {
func (r DbsApiGenerateUnitedMarketplaceServicesReportRequest) GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest GenerateUnitedMarketplaceServicesReportRequest) DbsApiGenerateUnitedMarketplaceServicesReportRequest {
func (r DbsApiGenerateUnitedMarketplaceServicesReportRequest) Format(format ReportFormatType) DbsApiGenerateUnitedMarketplaceServicesReportRequest {
func (r DbsApiGenerateUnitedMarketplaceServicesReportRequest) Language(language ReportLanguageType) DbsApiGenerateUnitedMarketplaceServicesReportRequest {
func (r DbsApiGenerateUnitedMarketplaceServicesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateUnitedMarketplaceServicesReportRequest
func (a *DbsAPIService) GenerateUnitedMarketplaceServicesReport(ctx context.Context) DbsApiGenerateUnitedMarketplaceServicesReportRequest {
	return DbsApiGenerateUnitedMarketplaceServicesReportRequest{
func (a *DbsAPIService) GenerateUnitedMarketplaceServicesReportExecute(r DbsApiGenerateUnitedMarketplaceServicesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateUnitedNettingReportRequest struct {
func (r DbsApiGenerateUnitedNettingReportRequest) GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest GenerateUnitedNettingReportRequest) DbsApiGenerateUnitedNettingReportRequest {
func (r DbsApiGenerateUnitedNettingReportRequest) Format(format ReportFormatType) DbsApiGenerateUnitedNettingReportRequest {
func (r DbsApiGenerateUnitedNettingReportRequest) Language(language ReportLanguageType) DbsApiGenerateUnitedNettingReportRequest {
func (r DbsApiGenerateUnitedNettingReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateUnitedNettingReportRequest
func (a *DbsAPIService) GenerateUnitedNettingReport(ctx context.Context) DbsApiGenerateUnitedNettingReportRequest {
	return DbsApiGenerateUnitedNettingReportRequest{
func (a *DbsAPIService) GenerateUnitedNettingReportExecute(r DbsApiGenerateUnitedNettingReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateUnitedOrdersReportRequest struct {
func (r DbsApiGenerateUnitedOrdersReportRequest) GenerateUnitedOrdersRequest(generateUnitedOrdersRequest GenerateUnitedOrdersRequest) DbsApiGenerateUnitedOrdersReportRequest {
func (r DbsApiGenerateUnitedOrdersReportRequest) Format(format ReportFormatType) DbsApiGenerateUnitedOrdersReportRequest {
func (r DbsApiGenerateUnitedOrdersReportRequest) Language(language ReportLanguageType) DbsApiGenerateUnitedOrdersReportRequest {
func (r DbsApiGenerateUnitedOrdersReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateUnitedOrdersReportRequest
func (a *DbsAPIService) GenerateUnitedOrdersReport(ctx context.Context) DbsApiGenerateUnitedOrdersReportRequest {
	return DbsApiGenerateUnitedOrdersReportRequest{
func (a *DbsAPIService) GenerateUnitedOrdersReportExecute(r DbsApiGenerateUnitedOrdersReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGenerateUnitedReturnsReportRequest struct {
func (r DbsApiGenerateUnitedReturnsReportRequest) GenerateUnitedReturnsRequest(generateUnitedReturnsRequest GenerateUnitedReturnsRequest) DbsApiGenerateUnitedReturnsReportRequest {
func (r DbsApiGenerateUnitedReturnsReportRequest) Format(format ReportFormatType) DbsApiGenerateUnitedReturnsReportRequest {
func (r DbsApiGenerateUnitedReturnsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return DbsApiGenerateUnitedReturnsReportRequest
func (a *DbsAPIService) GenerateUnitedReturnsReport(ctx context.Context) DbsApiGenerateUnitedReturnsReportRequest {
	return DbsApiGenerateUnitedReturnsReportRequest{
func (a *DbsAPIService) GenerateUnitedReturnsReportExecute(r DbsApiGenerateUnitedReturnsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type DbsApiGetAuthTokenInfoRequest struct {
func (r DbsApiGetAuthTokenInfoRequest) Execute() (*GetTokenInfoResponse, *http.Response, error) {
	@return DbsApiGetAuthTokenInfoRequest
func (a *DbsAPIService) GetAuthTokenInfo(ctx context.Context) DbsApiGetAuthTokenInfoRequest {
	return DbsApiGetAuthTokenInfoRequest{
func (a *DbsAPIService) GetAuthTokenInfoExecute(r DbsApiGetAuthTokenInfoRequest) (*GetTokenInfoResponse, *http.Response, error) {
type DbsApiGetBidsInfoForBusinessRequest struct {
func (r DbsApiGetBidsInfoForBusinessRequest) PageToken(pageToken string) DbsApiGetBidsInfoForBusinessRequest {
func (r DbsApiGetBidsInfoForBusinessRequest) Limit(limit int32) DbsApiGetBidsInfoForBusinessRequest {
func (r DbsApiGetBidsInfoForBusinessRequest) GetBidsInfoRequest(getBidsInfoRequest GetBidsInfoRequest) DbsApiGetBidsInfoForBusinessRequest {
func (r DbsApiGetBidsInfoForBusinessRequest) Execute() (*GetBidsInfoResponse, *http.Response, error) {
	@return DbsApiGetBidsInfoForBusinessRequest
func (a *DbsAPIService) GetBidsInfoForBusiness(ctx context.Context, businessId int64) DbsApiGetBidsInfoForBusinessRequest {
	return DbsApiGetBidsInfoForBusinessRequest{
func (a *DbsAPIService) GetBidsInfoForBusinessExecute(r DbsApiGetBidsInfoForBusinessRequest) (*GetBidsInfoResponse, *http.Response, error) {
type DbsApiGetBidsRecommendationsRequest struct {
func (r DbsApiGetBidsRecommendationsRequest) GetBidsRecommendationsRequest(getBidsRecommendationsRequest GetBidsRecommendationsRequest) DbsApiGetBidsRecommendationsRequest {
func (r DbsApiGetBidsRecommendationsRequest) Execute() (*GetBidsRecommendationsResponse, *http.Response, error) {
	@return DbsApiGetBidsRecommendationsRequest
func (a *DbsAPIService) GetBidsRecommendations(ctx context.Context, businessId int64) DbsApiGetBidsRecommendationsRequest {
	return DbsApiGetBidsRecommendationsRequest{
func (a *DbsAPIService) GetBidsRecommendationsExecute(r DbsApiGetBidsRecommendationsRequest) (*GetBidsRecommendationsResponse, *http.Response, error) {
type DbsApiGetBusinessQuarantineOffersRequest struct {
func (r DbsApiGetBusinessQuarantineOffersRequest) GetQuarantineOffersRequest(getQuarantineOffersRequest GetQuarantineOffersRequest) DbsApiGetBusinessQuarantineOffersRequest {
func (r DbsApiGetBusinessQuarantineOffersRequest) PageToken(pageToken string) DbsApiGetBusinessQuarantineOffersRequest {
func (r DbsApiGetBusinessQuarantineOffersRequest) Limit(limit int32) DbsApiGetBusinessQuarantineOffersRequest {
func (r DbsApiGetBusinessQuarantineOffersRequest) Execute() (*GetQuarantineOffersResponse, *http.Response, error) {
	@return DbsApiGetBusinessQuarantineOffersRequest
func (a *DbsAPIService) GetBusinessQuarantineOffers(ctx context.Context, businessId int64) DbsApiGetBusinessQuarantineOffersRequest {
	return DbsApiGetBusinessQuarantineOffersRequest{
func (a *DbsAPIService) GetBusinessQuarantineOffersExecute(r DbsApiGetBusinessQuarantineOffersRequest) (*GetQuarantineOffersResponse, *http.Response, error) {
type DbsApiGetBusinessSettingsRequest struct {
func (r DbsApiGetBusinessSettingsRequest) Execute() (*GetBusinessSettingsResponse, *http.Response, error) {
	@return DbsApiGetBusinessSettingsRequest
func (a *DbsAPIService) GetBusinessSettings(ctx context.Context, businessId int64) DbsApiGetBusinessSettingsRequest {
	return DbsApiGetBusinessSettingsRequest{
func (a *DbsAPIService) GetBusinessSettingsExecute(r DbsApiGetBusinessSettingsRequest) (*GetBusinessSettingsResponse, *http.Response, error) {
type DbsApiGetCampaignRequest struct {
func (r DbsApiGetCampaignRequest) Execute() (*GetCampaignResponse, *http.Response, error) {
	@return DbsApiGetCampaignRequest
func (a *DbsAPIService) GetCampaign(ctx context.Context, campaignId int64) DbsApiGetCampaignRequest {
	return DbsApiGetCampaignRequest{
func (a *DbsAPIService) GetCampaignExecute(r DbsApiGetCampaignRequest) (*GetCampaignResponse, *http.Response, error) {
type DbsApiGetCampaignOffersRequest struct {
func (r DbsApiGetCampaignOffersRequest) GetCampaignOffersRequest(getCampaignOffersRequest GetCampaignOffersRequest) DbsApiGetCampaignOffersRequest {
func (r DbsApiGetCampaignOffersRequest) PageToken(pageToken string) DbsApiGetCampaignOffersRequest {
func (r DbsApiGetCampaignOffersRequest) Limit(limit int32) DbsApiGetCampaignOffersRequest {
func (r DbsApiGetCampaignOffersRequest) Execute() (*GetCampaignOffersResponse, *http.Response, error) {
	@return DbsApiGetCampaignOffersRequest
func (a *DbsAPIService) GetCampaignOffers(ctx context.Context, campaignId int64) DbsApiGetCampaignOffersRequest {
	return DbsApiGetCampaignOffersRequest{
func (a *DbsAPIService) GetCampaignOffersExecute(r DbsApiGetCampaignOffersRequest) (*GetCampaignOffersResponse, *http.Response, error) {
type DbsApiGetCampaignQuarantineOffersRequest struct {
func (r DbsApiGetCampaignQuarantineOffersRequest) GetQuarantineOffersRequest(getQuarantineOffersRequest GetQuarantineOffersRequest) DbsApiGetCampaignQuarantineOffersRequest {
func (r DbsApiGetCampaignQuarantineOffersRequest) PageToken(pageToken string) DbsApiGetCampaignQuarantineOffersRequest {
func (r DbsApiGetCampaignQuarantineOffersRequest) Limit(limit int32) DbsApiGetCampaignQuarantineOffersRequest {
func (r DbsApiGetCampaignQuarantineOffersRequest) Execute() (*GetQuarantineOffersResponse, *http.Response, error) {
	@return DbsApiGetCampaignQuarantineOffersRequest
func (a *DbsAPIService) GetCampaignQuarantineOffers(ctx context.Context, campaignId int64) DbsApiGetCampaignQuarantineOffersRequest {
	return DbsApiGetCampaignQuarantineOffersRequest{
func (a *DbsAPIService) GetCampaignQuarantineOffersExecute(r DbsApiGetCampaignQuarantineOffersRequest) (*GetQuarantineOffersResponse, *http.Response, error) {
type DbsApiGetCampaignRegionRequest struct {
func (r DbsApiGetCampaignRegionRequest) Execute() (*GetCampaignRegionResponse, *http.Response, error) {
	@return DbsApiGetCampaignRegionRequest
func (a *DbsAPIService) GetCampaignRegion(ctx context.Context, campaignId int64) DbsApiGetCampaignRegionRequest {
	return DbsApiGetCampaignRegionRequest{
func (a *DbsAPIService) GetCampaignRegionExecute(r DbsApiGetCampaignRegionRequest) (*GetCampaignRegionResponse, *http.Response, error) {
type DbsApiGetCampaignSettingsRequest struct {
func (r DbsApiGetCampaignSettingsRequest) Execute() (*GetCampaignSettingsResponse, *http.Response, error) {
	@return DbsApiGetCampaignSettingsRequest
func (a *DbsAPIService) GetCampaignSettings(ctx context.Context, campaignId int64) DbsApiGetCampaignSettingsRequest {
	return DbsApiGetCampaignSettingsRequest{
func (a *DbsAPIService) GetCampaignSettingsExecute(r DbsApiGetCampaignSettingsRequest) (*GetCampaignSettingsResponse, *http.Response, error) {
type DbsApiGetCampaignsRequest struct {
func (r DbsApiGetCampaignsRequest) Page(page int32) DbsApiGetCampaignsRequest {
func (r DbsApiGetCampaignsRequest) PageSize(pageSize int32) DbsApiGetCampaignsRequest {
func (r DbsApiGetCampaignsRequest) Execute() (*GetCampaignsResponse, *http.Response, error) {
	@return DbsApiGetCampaignsRequest
func (a *DbsAPIService) GetCampaigns(ctx context.Context) DbsApiGetCampaignsRequest {
	return DbsApiGetCampaignsRequest{
func (a *DbsAPIService) GetCampaignsExecute(r DbsApiGetCampaignsRequest) (*GetCampaignsResponse, *http.Response, error) {
type DbsApiGetCategoriesMaxSaleQuantumRequest struct {
func (r DbsApiGetCategoriesMaxSaleQuantumRequest) GetCategoriesMaxSaleQuantumRequest(getCategoriesMaxSaleQuantumRequest GetCategoriesMaxSaleQuantumRequest) DbsApiGetCategoriesMaxSaleQuantumRequest {
func (r DbsApiGetCategoriesMaxSaleQuantumRequest) Execute() (*GetCategoriesMaxSaleQuantumResponse, *http.Response, error) {
	@return DbsApiGetCategoriesMaxSaleQuantumRequest
func (a *DbsAPIService) GetCategoriesMaxSaleQuantum(ctx context.Context) DbsApiGetCategoriesMaxSaleQuantumRequest {
	return DbsApiGetCategoriesMaxSaleQuantumRequest{
func (a *DbsAPIService) GetCategoriesMaxSaleQuantumExecute(r DbsApiGetCategoriesMaxSaleQuantumRequest) (*GetCategoriesMaxSaleQuantumResponse, *http.Response, error) {
type DbsApiGetCategoriesTreeRequest struct {
func (r DbsApiGetCategoriesTreeRequest) GetCategoriesRequest(getCategoriesRequest GetCategoriesRequest) DbsApiGetCategoriesTreeRequest {
func (r DbsApiGetCategoriesTreeRequest) Execute() (*GetCategoriesResponse, *http.Response, error) {
	@return DbsApiGetCategoriesTreeRequest
func (a *DbsAPIService) GetCategoriesTree(ctx context.Context) DbsApiGetCategoriesTreeRequest {
	return DbsApiGetCategoriesTreeRequest{
func (a *DbsAPIService) GetCategoriesTreeExecute(r DbsApiGetCategoriesTreeRequest) (*GetCategoriesResponse, *http.Response, error) {
type DbsApiGetCategoryContentParametersRequest struct {
func (r DbsApiGetCategoryContentParametersRequest) Execute() (*GetCategoryContentParametersResponse, *http.Response, error) {
	@return DbsApiGetCategoryContentParametersRequest
func (a *DbsAPIService) GetCategoryContentParameters(ctx context.Context, categoryId int64) DbsApiGetCategoryContentParametersRequest {
	return DbsApiGetCategoryContentParametersRequest{
func (a *DbsAPIService) GetCategoryContentParametersExecute(r DbsApiGetCategoryContentParametersRequest) (*GetCategoryContentParametersResponse, *http.Response, error) {
type DbsApiGetChatRequest struct {
func (r DbsApiGetChatRequest) ChatId(chatId int64) DbsApiGetChatRequest {
func (r DbsApiGetChatRequest) Execute() (*GetChatResponse, *http.Response, error) {
	@return DbsApiGetChatRequest
func (a *DbsAPIService) GetChat(ctx context.Context, businessId int64) DbsApiGetChatRequest {
	return DbsApiGetChatRequest{
func (a *DbsAPIService) GetChatExecute(r DbsApiGetChatRequest) (*GetChatResponse, *http.Response, error) {
type DbsApiGetChatHistoryRequest struct {
func (r DbsApiGetChatHistoryRequest) ChatId(chatId int64) DbsApiGetChatHistoryRequest {
func (r DbsApiGetChatHistoryRequest) GetChatHistoryRequest(getChatHistoryRequest GetChatHistoryRequest) DbsApiGetChatHistoryRequest {
func (r DbsApiGetChatHistoryRequest) PageToken(pageToken string) DbsApiGetChatHistoryRequest {
func (r DbsApiGetChatHistoryRequest) Limit(limit int32) DbsApiGetChatHistoryRequest {
func (r DbsApiGetChatHistoryRequest) Execute() (*GetChatHistoryResponse, *http.Response, error) {
	@return DbsApiGetChatHistoryRequest
func (a *DbsAPIService) GetChatHistory(ctx context.Context, businessId int64) DbsApiGetChatHistoryRequest {
	return DbsApiGetChatHistoryRequest{
func (a *DbsAPIService) GetChatHistoryExecute(r DbsApiGetChatHistoryRequest) (*GetChatHistoryResponse, *http.Response, error) {
type DbsApiGetChatMessageRequest struct {
func (r DbsApiGetChatMessageRequest) ChatId(chatId int64) DbsApiGetChatMessageRequest {
func (r DbsApiGetChatMessageRequest) MessageId(messageId int64) DbsApiGetChatMessageRequest {
func (r DbsApiGetChatMessageRequest) Execute() (*GetChatMessageResponse, *http.Response, error) {
	@return DbsApiGetChatMessageRequest
func (a *DbsAPIService) GetChatMessage(ctx context.Context, businessId int64) DbsApiGetChatMessageRequest {
	return DbsApiGetChatMessageRequest{
func (a *DbsAPIService) GetChatMessageExecute(r DbsApiGetChatMessageRequest) (*GetChatMessageResponse, *http.Response, error) {
type DbsApiGetChatsRequest struct {
func (r DbsApiGetChatsRequest) GetChatsRequest(getChatsRequest GetChatsRequest) DbsApiGetChatsRequest {
func (r DbsApiGetChatsRequest) PageToken(pageToken string) DbsApiGetChatsRequest {
func (r DbsApiGetChatsRequest) Limit(limit int32) DbsApiGetChatsRequest {
func (r DbsApiGetChatsRequest) Execute() (*GetChatsResponse, *http.Response, error) {
	@return DbsApiGetChatsRequest
func (a *DbsAPIService) GetChats(ctx context.Context, businessId int64) DbsApiGetChatsRequest {
	return DbsApiGetChatsRequest{
func (a *DbsAPIService) GetChatsExecute(r DbsApiGetChatsRequest) (*GetChatsResponse, *http.Response, error) {
type DbsApiGetDeliveryServicesRequest struct {
func (r DbsApiGetDeliveryServicesRequest) Execute() (*GetDeliveryServicesResponse, *http.Response, error) {
	@return DbsApiGetDeliveryServicesRequest
func (a *DbsAPIService) GetDeliveryServices(ctx context.Context) DbsApiGetDeliveryServicesRequest {
	return DbsApiGetDeliveryServicesRequest{
func (a *DbsAPIService) GetDeliveryServicesExecute(r DbsApiGetDeliveryServicesRequest) (*GetDeliveryServicesResponse, *http.Response, error) {
type DbsApiGetGoodsFeedbackCommentsRequest struct {
func (r DbsApiGetGoodsFeedbackCommentsRequest) GetGoodsFeedbackCommentsRequest(getGoodsFeedbackCommentsRequest GetGoodsFeedbackCommentsRequest) DbsApiGetGoodsFeedbackCommentsRequest {
func (r DbsApiGetGoodsFeedbackCommentsRequest) PageToken(pageToken string) DbsApiGetGoodsFeedbackCommentsRequest {
func (r DbsApiGetGoodsFeedbackCommentsRequest) Limit(limit int32) DbsApiGetGoodsFeedbackCommentsRequest {
func (r DbsApiGetGoodsFeedbackCommentsRequest) Execute() (*GetGoodsFeedbackCommentsResponse, *http.Response, error) {
	@return DbsApiGetGoodsFeedbackCommentsRequest
func (a *DbsAPIService) GetGoodsFeedbackComments(ctx context.Context, businessId int64) DbsApiGetGoodsFeedbackCommentsRequest {
	return DbsApiGetGoodsFeedbackCommentsRequest{
func (a *DbsAPIService) GetGoodsFeedbackCommentsExecute(r DbsApiGetGoodsFeedbackCommentsRequest) (*GetGoodsFeedbackCommentsResponse, *http.Response, error) {
type DbsApiGetGoodsFeedbacksRequest struct {
func (r DbsApiGetGoodsFeedbacksRequest) PageToken(pageToken string) DbsApiGetGoodsFeedbacksRequest {
func (r DbsApiGetGoodsFeedbacksRequest) Limit(limit int32) DbsApiGetGoodsFeedbacksRequest {
func (r DbsApiGetGoodsFeedbacksRequest) GetGoodsFeedbackRequest(getGoodsFeedbackRequest GetGoodsFeedbackRequest) DbsApiGetGoodsFeedbacksRequest {
func (r DbsApiGetGoodsFeedbacksRequest) Execute() (*GetGoodsFeedbackResponse, *http.Response, error) {
	@return DbsApiGetGoodsFeedbacksRequest
func (a *DbsAPIService) GetGoodsFeedbacks(ctx context.Context, businessId int64) DbsApiGetGoodsFeedbacksRequest {
	return DbsApiGetGoodsFeedbacksRequest{
func (a *DbsAPIService) GetGoodsFeedbacksExecute(r DbsApiGetGoodsFeedbacksRequest) (*GetGoodsFeedbackResponse, *http.Response, error) {
type DbsApiGetGoodsStatsRequest struct {
func (r DbsApiGetGoodsStatsRequest) GetGoodsStatsRequest(getGoodsStatsRequest GetGoodsStatsRequest) DbsApiGetGoodsStatsRequest {
func (r DbsApiGetGoodsStatsRequest) Execute() (*GetGoodsStatsResponse, *http.Response, error) {
	@return DbsApiGetGoodsStatsRequest
func (a *DbsAPIService) GetGoodsStats(ctx context.Context, campaignId int64) DbsApiGetGoodsStatsRequest {
	return DbsApiGetGoodsStatsRequest{
func (a *DbsAPIService) GetGoodsStatsExecute(r DbsApiGetGoodsStatsRequest) (*GetGoodsStatsResponse, *http.Response, error) {
type DbsApiGetHiddenOffersRequest struct {
func (r DbsApiGetHiddenOffersRequest) OfferId(offerId []string) DbsApiGetHiddenOffersRequest {
func (r DbsApiGetHiddenOffersRequest) PageToken(pageToken string) DbsApiGetHiddenOffersRequest {
func (r DbsApiGetHiddenOffersRequest) Limit(limit int32) DbsApiGetHiddenOffersRequest {
func (r DbsApiGetHiddenOffersRequest) Execute() (*GetHiddenOffersResponse, *http.Response, error) {
	@return DbsApiGetHiddenOffersRequest
func (a *DbsAPIService) GetHiddenOffers(ctx context.Context, campaignId int64) DbsApiGetHiddenOffersRequest {
	return DbsApiGetHiddenOffersRequest{
func (a *DbsAPIService) GetHiddenOffersExecute(r DbsApiGetHiddenOffersRequest) (*GetHiddenOffersResponse, *http.Response, error) {
type DbsApiGetModelRequest struct {
func (r DbsApiGetModelRequest) RegionId(regionId int64) DbsApiGetModelRequest {
func (r DbsApiGetModelRequest) Currency(currency CurrencyType) DbsApiGetModelRequest {
func (r DbsApiGetModelRequest) Execute() (*GetModelsResponse, *http.Response, error) {
	@return DbsApiGetModelRequest
func (a *DbsAPIService) GetModel(ctx context.Context, modelId int64) DbsApiGetModelRequest {
	return DbsApiGetModelRequest{
func (a *DbsAPIService) GetModelExecute(r DbsApiGetModelRequest) (*GetModelsResponse, *http.Response, error) {
type DbsApiGetModelOffersRequest struct {
func (r DbsApiGetModelOffersRequest) RegionId(regionId int64) DbsApiGetModelOffersRequest {
func (r DbsApiGetModelOffersRequest) Currency(currency CurrencyType) DbsApiGetModelOffersRequest {
func (r DbsApiGetModelOffersRequest) OrderByPrice(orderByPrice SortOrderType) DbsApiGetModelOffersRequest {
func (r DbsApiGetModelOffersRequest) Count(count int32) DbsApiGetModelOffersRequest {
func (r DbsApiGetModelOffersRequest) Page(page int32) DbsApiGetModelOffersRequest {
func (r DbsApiGetModelOffersRequest) Execute() (*GetModelsOffersResponse, *http.Response, error) {
	@return DbsApiGetModelOffersRequest
func (a *DbsAPIService) GetModelOffers(ctx context.Context, modelId int64) DbsApiGetModelOffersRequest {
	return DbsApiGetModelOffersRequest{
func (a *DbsAPIService) GetModelOffersExecute(r DbsApiGetModelOffersRequest) (*GetModelsOffersResponse, *http.Response, error) {
type DbsApiGetModelsRequest struct {
func (r DbsApiGetModelsRequest) RegionId(regionId int64) DbsApiGetModelsRequest {
func (r DbsApiGetModelsRequest) GetModelsRequest(getModelsRequest GetModelsRequest) DbsApiGetModelsRequest {
func (r DbsApiGetModelsRequest) Currency(currency CurrencyType) DbsApiGetModelsRequest {
func (r DbsApiGetModelsRequest) Execute() (*GetModelsResponse, *http.Response, error) {
	@return DbsApiGetModelsRequest
func (a *DbsAPIService) GetModels(ctx context.Context) DbsApiGetModelsRequest {
	return DbsApiGetModelsRequest{
func (a *DbsAPIService) GetModelsExecute(r DbsApiGetModelsRequest) (*GetModelsResponse, *http.Response, error) {
type DbsApiGetModelsOffersRequest struct {
func (r DbsApiGetModelsOffersRequest) RegionId(regionId int64) DbsApiGetModelsOffersRequest {
func (r DbsApiGetModelsOffersRequest) GetModelsRequest(getModelsRequest GetModelsRequest) DbsApiGetModelsOffersRequest {
func (r DbsApiGetModelsOffersRequest) Currency(currency CurrencyType) DbsApiGetModelsOffersRequest {
func (r DbsApiGetModelsOffersRequest) OrderByPrice(orderByPrice SortOrderType) DbsApiGetModelsOffersRequest {
func (r DbsApiGetModelsOffersRequest) Execute() (*GetModelsOffersResponse, *http.Response, error) {
	@return DbsApiGetModelsOffersRequest
func (a *DbsAPIService) GetModelsOffers(ctx context.Context) DbsApiGetModelsOffersRequest {
	return DbsApiGetModelsOffersRequest{
func (a *DbsAPIService) GetModelsOffersExecute(r DbsApiGetModelsOffersRequest) (*GetModelsOffersResponse, *http.Response, error) {
type DbsApiGetOfferCardsContentStatusRequest struct {
func (r DbsApiGetOfferCardsContentStatusRequest) PageToken(pageToken string) DbsApiGetOfferCardsContentStatusRequest {
func (r DbsApiGetOfferCardsContentStatusRequest) Limit(limit int32) DbsApiGetOfferCardsContentStatusRequest {
func (r DbsApiGetOfferCardsContentStatusRequest) GetOfferCardsContentStatusRequest(getOfferCardsContentStatusRequest GetOfferCardsContentStatusRequest) DbsApiGetOfferCardsContentStatusRequest {
func (r DbsApiGetOfferCardsContentStatusRequest) Execute() (*GetOfferCardsContentStatusResponse, *http.Response, error) {
	@return DbsApiGetOfferCardsContentStatusRequest
func (a *DbsAPIService) GetOfferCardsContentStatus(ctx context.Context, businessId int64) DbsApiGetOfferCardsContentStatusRequest {
	return DbsApiGetOfferCardsContentStatusRequest{
func (a *DbsAPIService) GetOfferCardsContentStatusExecute(r DbsApiGetOfferCardsContentStatusRequest) (*GetOfferCardsContentStatusResponse, *http.Response, error) {
type DbsApiGetOfferMappingEntriesRequest struct {
func (r DbsApiGetOfferMappingEntriesRequest) OfferId(offerId []string) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) ShopSku(shopSku []string) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) MappingKind(mappingKind OfferMappingKindType) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) Status(status []OfferProcessingStatusType) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) Availability(availability []OfferAvailabilityStatusType) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) CategoryId(categoryId []int32) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) Vendor(vendor []string) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) PageToken(pageToken string) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) Limit(limit int32) DbsApiGetOfferMappingEntriesRequest {
func (r DbsApiGetOfferMappingEntriesRequest) Execute() (*GetOfferMappingEntriesResponse, *http.Response, error) {
	@return DbsApiGetOfferMappingEntriesRequest
func (a *DbsAPIService) GetOfferMappingEntries(ctx context.Context, campaignId int64) DbsApiGetOfferMappingEntriesRequest {
	return DbsApiGetOfferMappingEntriesRequest{
func (a *DbsAPIService) GetOfferMappingEntriesExecute(r DbsApiGetOfferMappingEntriesRequest) (*GetOfferMappingEntriesResponse, *http.Response, error) {
type DbsApiGetOfferMappingsRequest struct {
func (r DbsApiGetOfferMappingsRequest) PageToken(pageToken string) DbsApiGetOfferMappingsRequest {
func (r DbsApiGetOfferMappingsRequest) Limit(limit int32) DbsApiGetOfferMappingsRequest {
func (r DbsApiGetOfferMappingsRequest) Language(language CatalogLanguageType) DbsApiGetOfferMappingsRequest {
func (r DbsApiGetOfferMappingsRequest) GetOfferMappingsRequest(getOfferMappingsRequest GetOfferMappingsRequest) DbsApiGetOfferMappingsRequest {
func (r DbsApiGetOfferMappingsRequest) Execute() (*GetOfferMappingsResponse, *http.Response, error) {
	@return DbsApiGetOfferMappingsRequest
func (a *DbsAPIService) GetOfferMappings(ctx context.Context, businessId int64) DbsApiGetOfferMappingsRequest {
	return DbsApiGetOfferMappingsRequest{
func (a *DbsAPIService) GetOfferMappingsExecute(r DbsApiGetOfferMappingsRequest) (*GetOfferMappingsResponse, *http.Response, error) {
type DbsApiGetOfferRecommendationsRequest struct {
func (r DbsApiGetOfferRecommendationsRequest) GetOfferRecommendationsRequest(getOfferRecommendationsRequest GetOfferRecommendationsRequest) DbsApiGetOfferRecommendationsRequest {
func (r DbsApiGetOfferRecommendationsRequest) PageToken(pageToken string) DbsApiGetOfferRecommendationsRequest {
func (r DbsApiGetOfferRecommendationsRequest) Limit(limit int32) DbsApiGetOfferRecommendationsRequest {
func (r DbsApiGetOfferRecommendationsRequest) Execute() (*GetOfferRecommendationsResponse, *http.Response, error) {
	@return DbsApiGetOfferRecommendationsRequest
func (a *DbsAPIService) GetOfferRecommendations(ctx context.Context, businessId int64) DbsApiGetOfferRecommendationsRequest {
	return DbsApiGetOfferRecommendationsRequest{
func (a *DbsAPIService) GetOfferRecommendationsExecute(r DbsApiGetOfferRecommendationsRequest) (*GetOfferRecommendationsResponse, *http.Response, error) {
type DbsApiGetOrderRequest struct {
func (r DbsApiGetOrderRequest) Execute() (*GetOrderResponse, *http.Response, error) {
	@return DbsApiGetOrderRequest
func (a *DbsAPIService) GetOrder(ctx context.Context, campaignId int64, orderId int64) DbsApiGetOrderRequest {
	return DbsApiGetOrderRequest{
func (a *DbsAPIService) GetOrderExecute(r DbsApiGetOrderRequest) (*GetOrderResponse, *http.Response, error) {
type DbsApiGetOrderBusinessBuyerInfoRequest struct {
func (r DbsApiGetOrderBusinessBuyerInfoRequest) Execute() (*GetBusinessBuyerInfoResponse, *http.Response, error) {
	@return DbsApiGetOrderBusinessBuyerInfoRequest
func (a *DbsAPIService) GetOrderBusinessBuyerInfo(ctx context.Context, campaignId int64, orderId int64) DbsApiGetOrderBusinessBuyerInfoRequest {
	return DbsApiGetOrderBusinessBuyerInfoRequest{
func (a *DbsAPIService) GetOrderBusinessBuyerInfoExecute(r DbsApiGetOrderBusinessBuyerInfoRequest) (*GetBusinessBuyerInfoResponse, *http.Response, error) {
type DbsApiGetOrderBusinessDocumentsInfoRequest struct {
func (r DbsApiGetOrderBusinessDocumentsInfoRequest) Execute() (*GetBusinessDocumentsInfoResponse, *http.Response, error) {
	@return DbsApiGetOrderBusinessDocumentsInfoRequest
func (a *DbsAPIService) GetOrderBusinessDocumentsInfo(ctx context.Context, campaignId int64, orderId int64) DbsApiGetOrderBusinessDocumentsInfoRequest {
	return DbsApiGetOrderBusinessDocumentsInfoRequest{
func (a *DbsAPIService) GetOrderBusinessDocumentsInfoExecute(r DbsApiGetOrderBusinessDocumentsInfoRequest) (*GetBusinessDocumentsInfoResponse, *http.Response, error) {
type DbsApiGetOrderBuyerInfoRequest struct {
func (r DbsApiGetOrderBuyerInfoRequest) Execute() (*GetOrderBuyerInfoResponse, *http.Response, error) {
	@return DbsApiGetOrderBuyerInfoRequest
func (a *DbsAPIService) GetOrderBuyerInfo(ctx context.Context, campaignId int64, orderId int64) DbsApiGetOrderBuyerInfoRequest {
	return DbsApiGetOrderBuyerInfoRequest{
func (a *DbsAPIService) GetOrderBuyerInfoExecute(r DbsApiGetOrderBuyerInfoRequest) (*GetOrderBuyerInfoResponse, *http.Response, error) {
type DbsApiGetOrderLabelsDataRequest struct {
func (r DbsApiGetOrderLabelsDataRequest) Execute() (*GetOrderLabelsDataResponse, *http.Response, error) {
	@return DbsApiGetOrderLabelsDataRequest
func (a *DbsAPIService) GetOrderLabelsData(ctx context.Context, campaignId int64, orderId int64) DbsApiGetOrderLabelsDataRequest {
	return DbsApiGetOrderLabelsDataRequest{
func (a *DbsAPIService) GetOrderLabelsDataExecute(r DbsApiGetOrderLabelsDataRequest) (*GetOrderLabelsDataResponse, *http.Response, error) {
type DbsApiGetOrdersRequest struct {
func (r DbsApiGetOrdersRequest) OrderIds(orderIds []int64) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) Status(status []OrderStatusType) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) Substatus(substatus []OrderSubstatusType) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) FromDate(fromDate string) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) ToDate(toDate string) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) SupplierShipmentDateFrom(supplierShipmentDateFrom string) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) SupplierShipmentDateTo(supplierShipmentDateTo string) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) UpdatedAtFrom(updatedAtFrom time.Time) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) UpdatedAtTo(updatedAtTo time.Time) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) DispatchType(dispatchType OrderDeliveryDispatchType) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) Fake(fake bool) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) HasCis(hasCis bool) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) OnlyWaitingForCancellationApprove(onlyWaitingForCancellationApprove bool) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) OnlyEstimatedDelivery(onlyEstimatedDelivery bool) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) BuyerType(buyerType OrderBuyerType) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) Page(page int32) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) PageSize(pageSize int32) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) PageToken(pageToken string) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) Limit(limit int32) DbsApiGetOrdersRequest {
func (r DbsApiGetOrdersRequest) Execute() (*GetOrdersResponse, *http.Response, error) {
	@return DbsApiGetOrdersRequest
func (a *DbsAPIService) GetOrders(ctx context.Context, campaignId int64) DbsApiGetOrdersRequest {
	return DbsApiGetOrdersRequest{
func (a *DbsAPIService) GetOrdersExecute(r DbsApiGetOrdersRequest) (*GetOrdersResponse, *http.Response, error) {
type DbsApiGetOrdersStatsRequest struct {
func (r DbsApiGetOrdersStatsRequest) PageToken(pageToken string) DbsApiGetOrdersStatsRequest {
func (r DbsApiGetOrdersStatsRequest) Limit(limit int32) DbsApiGetOrdersStatsRequest {
func (r DbsApiGetOrdersStatsRequest) GetOrdersStatsRequest(getOrdersStatsRequest GetOrdersStatsRequest) DbsApiGetOrdersStatsRequest {
func (r DbsApiGetOrdersStatsRequest) Execute() (*GetOrdersStatsResponse, *http.Response, error) {
	@return DbsApiGetOrdersStatsRequest
func (a *DbsAPIService) GetOrdersStats(ctx context.Context, campaignId int64) DbsApiGetOrdersStatsRequest {
	return DbsApiGetOrdersStatsRequest{
func (a *DbsAPIService) GetOrdersStatsExecute(r DbsApiGetOrdersStatsRequest) (*GetOrdersStatsResponse, *http.Response, error) {
type DbsApiGetOutletRequest struct {
func (r DbsApiGetOutletRequest) Execute() (*GetOutletResponse, *http.Response, error) {
	@return DbsApiGetOutletRequest
func (a *DbsAPIService) GetOutlet(ctx context.Context, campaignId int64, outletId int64) DbsApiGetOutletRequest {
	return DbsApiGetOutletRequest{
func (a *DbsAPIService) GetOutletExecute(r DbsApiGetOutletRequest) (*GetOutletResponse, *http.Response, error) {
type DbsApiGetOutletLicensesRequest struct {
func (r DbsApiGetOutletLicensesRequest) OutletIds(outletIds []int64) DbsApiGetOutletLicensesRequest {
func (r DbsApiGetOutletLicensesRequest) Ids(ids []int64) DbsApiGetOutletLicensesRequest {
func (r DbsApiGetOutletLicensesRequest) Execute() (*GetOutletLicensesResponse, *http.Response, error) {
	@return DbsApiGetOutletLicensesRequest
func (a *DbsAPIService) GetOutletLicenses(ctx context.Context, campaignId int64) DbsApiGetOutletLicensesRequest {
	return DbsApiGetOutletLicensesRequest{
func (a *DbsAPIService) GetOutletLicensesExecute(r DbsApiGetOutletLicensesRequest) (*GetOutletLicensesResponse, *http.Response, error) {
type DbsApiGetOutletsRequest struct {
func (r DbsApiGetOutletsRequest) PageToken(pageToken string) DbsApiGetOutletsRequest {
func (r DbsApiGetOutletsRequest) RegionId(regionId int64) DbsApiGetOutletsRequest {
func (r DbsApiGetOutletsRequest) ShopOutletCode(shopOutletCode string) DbsApiGetOutletsRequest {
func (r DbsApiGetOutletsRequest) RegionId2(regionId2 int64) DbsApiGetOutletsRequest {
func (r DbsApiGetOutletsRequest) Execute() (*GetOutletsResponse, *http.Response, error) {
	@return DbsApiGetOutletsRequest
func (a *DbsAPIService) GetOutlets(ctx context.Context, campaignId int64) DbsApiGetOutletsRequest {
	return DbsApiGetOutletsRequest{
func (a *DbsAPIService) GetOutletsExecute(r DbsApiGetOutletsRequest) (*GetOutletsResponse, *http.Response, error) {
type DbsApiGetPagedWarehousesRequest struct {
func (r DbsApiGetPagedWarehousesRequest) PageToken(pageToken string) DbsApiGetPagedWarehousesRequest {
func (r DbsApiGetPagedWarehousesRequest) Limit(limit int32) DbsApiGetPagedWarehousesRequest {
func (r DbsApiGetPagedWarehousesRequest) GetPagedWarehousesRequest(getPagedWarehousesRequest GetPagedWarehousesRequest) DbsApiGetPagedWarehousesRequest {
func (r DbsApiGetPagedWarehousesRequest) Execute() (*GetPagedWarehousesResponse, *http.Response, error) {
	@return DbsApiGetPagedWarehousesRequest
func (a *DbsAPIService) GetPagedWarehouses(ctx context.Context, businessId int64) DbsApiGetPagedWarehousesRequest {
	return DbsApiGetPagedWarehousesRequest{
func (a *DbsAPIService) GetPagedWarehousesExecute(r DbsApiGetPagedWarehousesRequest) (*GetPagedWarehousesResponse, *http.Response, error) {
type DbsApiGetPricesRequest struct {
func (r DbsApiGetPricesRequest) PageToken(pageToken string) DbsApiGetPricesRequest {
func (r DbsApiGetPricesRequest) Limit(limit int32) DbsApiGetPricesRequest {
func (r DbsApiGetPricesRequest) Archived(archived bool) DbsApiGetPricesRequest {
func (r DbsApiGetPricesRequest) Execute() (*GetPricesResponse, *http.Response, error) {
	@return DbsApiGetPricesRequest
func (a *DbsAPIService) GetPrices(ctx context.Context, campaignId int64) DbsApiGetPricesRequest {
	return DbsApiGetPricesRequest{
func (a *DbsAPIService) GetPricesExecute(r DbsApiGetPricesRequest) (*GetPricesResponse, *http.Response, error) {
type DbsApiGetPricesByOfferIdsRequest struct {
func (r DbsApiGetPricesByOfferIdsRequest) PageToken(pageToken string) DbsApiGetPricesByOfferIdsRequest {
func (r DbsApiGetPricesByOfferIdsRequest) Limit(limit int32) DbsApiGetPricesByOfferIdsRequest {
func (r DbsApiGetPricesByOfferIdsRequest) GetPricesByOfferIdsRequest(getPricesByOfferIdsRequest GetPricesByOfferIdsRequest) DbsApiGetPricesByOfferIdsRequest {
func (r DbsApiGetPricesByOfferIdsRequest) Execute() (*GetPricesByOfferIdsResponse, *http.Response, error) {
	@return DbsApiGetPricesByOfferIdsRequest
func (a *DbsAPIService) GetPricesByOfferIds(ctx context.Context, campaignId int64) DbsApiGetPricesByOfferIdsRequest {
	return DbsApiGetPricesByOfferIdsRequest{
func (a *DbsAPIService) GetPricesByOfferIdsExecute(r DbsApiGetPricesByOfferIdsRequest) (*GetPricesByOfferIdsResponse, *http.Response, error) {
type DbsApiGetPromoOffersRequest struct {
func (r DbsApiGetPromoOffersRequest) GetPromoOffersRequest(getPromoOffersRequest GetPromoOffersRequest) DbsApiGetPromoOffersRequest {
func (r DbsApiGetPromoOffersRequest) PageToken(pageToken string) DbsApiGetPromoOffersRequest {
func (r DbsApiGetPromoOffersRequest) Limit(limit int32) DbsApiGetPromoOffersRequest {
func (r DbsApiGetPromoOffersRequest) Execute() (*GetPromoOffersResponse, *http.Response, error) {
	@return DbsApiGetPromoOffersRequest
func (a *DbsAPIService) GetPromoOffers(ctx context.Context, businessId int64) DbsApiGetPromoOffersRequest {
	return DbsApiGetPromoOffersRequest{
func (a *DbsAPIService) GetPromoOffersExecute(r DbsApiGetPromoOffersRequest) (*GetPromoOffersResponse, *http.Response, error) {
type DbsApiGetPromosRequest struct {
func (r DbsApiGetPromosRequest) GetPromosRequest(getPromosRequest GetPromosRequest) DbsApiGetPromosRequest {
func (r DbsApiGetPromosRequest) Execute() (*GetPromosResponse, *http.Response, error) {
	@return DbsApiGetPromosRequest
func (a *DbsAPIService) GetPromos(ctx context.Context, businessId int64) DbsApiGetPromosRequest {
	return DbsApiGetPromosRequest{
func (a *DbsAPIService) GetPromosExecute(r DbsApiGetPromosRequest) (*GetPromosResponse, *http.Response, error) {
type DbsApiGetQualityRatingDetailsRequest struct {
func (r DbsApiGetQualityRatingDetailsRequest) Execute() (*GetQualityRatingDetailsResponse, *http.Response, error) {
	@return DbsApiGetQualityRatingDetailsRequest
func (a *DbsAPIService) GetQualityRatingDetails(ctx context.Context, campaignId int64) DbsApiGetQualityRatingDetailsRequest {
	return DbsApiGetQualityRatingDetailsRequest{
func (a *DbsAPIService) GetQualityRatingDetailsExecute(r DbsApiGetQualityRatingDetailsRequest) (*GetQualityRatingDetailsResponse, *http.Response, error) {
type DbsApiGetQualityRatingsRequest struct {
func (r DbsApiGetQualityRatingsRequest) GetQualityRatingRequest(getQualityRatingRequest GetQualityRatingRequest) DbsApiGetQualityRatingsRequest {
func (r DbsApiGetQualityRatingsRequest) Execute() (*GetQualityRatingResponse, *http.Response, error) {
	@return DbsApiGetQualityRatingsRequest
func (a *DbsAPIService) GetQualityRatings(ctx context.Context, businessId int64) DbsApiGetQualityRatingsRequest {
	return DbsApiGetQualityRatingsRequest{
func (a *DbsAPIService) GetQualityRatingsExecute(r DbsApiGetQualityRatingsRequest) (*GetQualityRatingResponse, *http.Response, error) {
type DbsApiGetRegionsCodesRequest struct {
func (r DbsApiGetRegionsCodesRequest) Execute() (*GetRegionsCodesResponse, *http.Response, error) {
	@return DbsApiGetRegionsCodesRequest
func (a *DbsAPIService) GetRegionsCodes(ctx context.Context) DbsApiGetRegionsCodesRequest {
	return DbsApiGetRegionsCodesRequest{
func (a *DbsAPIService) GetRegionsCodesExecute(r DbsApiGetRegionsCodesRequest) (*GetRegionsCodesResponse, *http.Response, error) {
type DbsApiGetReportInfoRequest struct {
func (r DbsApiGetReportInfoRequest) Execute() (*GetReportInfoResponse, *http.Response, error) {
	@return DbsApiGetReportInfoRequest
func (a *DbsAPIService) GetReportInfo(ctx context.Context, reportId string) DbsApiGetReportInfoRequest {
	return DbsApiGetReportInfoRequest{
func (a *DbsAPIService) GetReportInfoExecute(r DbsApiGetReportInfoRequest) (*GetReportInfoResponse, *http.Response, error) {
type DbsApiGetReturnRequest struct {
func (r DbsApiGetReturnRequest) Execute() (*GetReturnResponse, *http.Response, error) {
	@return DbsApiGetReturnRequest
func (a *DbsAPIService) GetReturn(ctx context.Context, campaignId int64, orderId int64, returnId int64) DbsApiGetReturnRequest {
	return DbsApiGetReturnRequest{
func (a *DbsAPIService) GetReturnExecute(r DbsApiGetReturnRequest) (*GetReturnResponse, *http.Response, error) {
type DbsApiGetReturnApplicationRequest struct {
func (r DbsApiGetReturnApplicationRequest) Execute() (*os.File, *http.Response, error) {
	@return DbsApiGetReturnApplicationRequest
func (a *DbsAPIService) GetReturnApplication(ctx context.Context, campaignId int64, orderId int64, returnId int64) DbsApiGetReturnApplicationRequest {
	return DbsApiGetReturnApplicationRequest{
func (a *DbsAPIService) GetReturnApplicationExecute(r DbsApiGetReturnApplicationRequest) (*os.File, *http.Response, error) {
type DbsApiGetReturnPhotoRequest struct {
func (r DbsApiGetReturnPhotoRequest) Execute() (*os.File, *http.Response, error) {
	@return DbsApiGetReturnPhotoRequest
func (a *DbsAPIService) GetReturnPhoto(ctx context.Context, campaignId int64, orderId int64, returnId int64, itemId int64, imageHash string) DbsApiGetReturnPhotoRequest {
	return DbsApiGetReturnPhotoRequest{
func (a *DbsAPIService) GetReturnPhotoExecute(r DbsApiGetReturnPhotoRequest) (*os.File, *http.Response, error) {
type DbsApiGetReturnsRequest struct {
func (r DbsApiGetReturnsRequest) PageToken(pageToken string) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) Limit(limit int32) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) OrderIds(orderIds []int64) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) Statuses(statuses []RefundStatusType) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) Type_(type_ ReturnType) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) FromDate(fromDate string) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) ToDate(toDate string) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) FromDate2(fromDate2 string) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) ToDate2(toDate2 string) DbsApiGetReturnsRequest {
func (r DbsApiGetReturnsRequest) Execute() (*GetReturnsResponse, *http.Response, error) {
	@return DbsApiGetReturnsRequest
func (a *DbsAPIService) GetReturns(ctx context.Context, campaignId int64) DbsApiGetReturnsRequest {
	return DbsApiGetReturnsRequest{
func (a *DbsAPIService) GetReturnsExecute(r DbsApiGetReturnsRequest) (*GetReturnsResponse, *http.Response, error) {
type DbsApiGetStocksRequest struct {
func (r DbsApiGetStocksRequest) PageToken(pageToken string) DbsApiGetStocksRequest {
func (r DbsApiGetStocksRequest) Limit(limit int32) DbsApiGetStocksRequest {
func (r DbsApiGetStocksRequest) GetWarehouseStocksRequest(getWarehouseStocksRequest GetWarehouseStocksRequest) DbsApiGetStocksRequest {
func (r DbsApiGetStocksRequest) Execute() (*GetWarehouseStocksResponse, *http.Response, error) {
	@return DbsApiGetStocksRequest
func (a *DbsAPIService) GetStocks(ctx context.Context, campaignId int64) DbsApiGetStocksRequest {
	return DbsApiGetStocksRequest{
func (a *DbsAPIService) GetStocksExecute(r DbsApiGetStocksRequest) (*GetWarehouseStocksResponse, *http.Response, error) {
type DbsApiGetSuggestedOfferMappingEntriesRequest struct {
func (r DbsApiGetSuggestedOfferMappingEntriesRequest) GetSuggestedOfferMappingEntriesRequest(getSuggestedOfferMappingEntriesRequest GetSuggestedOfferMappingEntriesRequest) DbsApiGetSuggestedOfferMappingEntriesRequest {
func (r DbsApiGetSuggestedOfferMappingEntriesRequest) Execute() (*GetSuggestedOfferMappingEntriesResponse, *http.Response, error) {
	@return DbsApiGetSuggestedOfferMappingEntriesRequest
func (a *DbsAPIService) GetSuggestedOfferMappingEntries(ctx context.Context, campaignId int64) DbsApiGetSuggestedOfferMappingEntriesRequest {
	return DbsApiGetSuggestedOfferMappingEntriesRequest{
func (a *DbsAPIService) GetSuggestedOfferMappingEntriesExecute(r DbsApiGetSuggestedOfferMappingEntriesRequest) (*GetSuggestedOfferMappingEntriesResponse, *http.Response, error) {
type DbsApiGetSuggestedOfferMappingsRequest struct {
func (r DbsApiGetSuggestedOfferMappingsRequest) GetSuggestedOfferMappingsRequest(getSuggestedOfferMappingsRequest GetSuggestedOfferMappingsRequest) DbsApiGetSuggestedOfferMappingsRequest {
func (r DbsApiGetSuggestedOfferMappingsRequest) Execute() (*GetSuggestedOfferMappingsResponse, *http.Response, error) {
	@return DbsApiGetSuggestedOfferMappingsRequest
func (a *DbsAPIService) GetSuggestedOfferMappings(ctx context.Context, businessId int64) DbsApiGetSuggestedOfferMappingsRequest {
	return DbsApiGetSuggestedOfferMappingsRequest{
func (a *DbsAPIService) GetSuggestedOfferMappingsExecute(r DbsApiGetSuggestedOfferMappingsRequest) (*GetSuggestedOfferMappingsResponse, *http.Response, error) {
type DbsApiGetSuggestedPricesRequest struct {
func (r DbsApiGetSuggestedPricesRequest) SuggestPricesRequest(suggestPricesRequest SuggestPricesRequest) DbsApiGetSuggestedPricesRequest {
func (r DbsApiGetSuggestedPricesRequest) Execute() (*SuggestPricesResponse, *http.Response, error) {
	@return DbsApiGetSuggestedPricesRequest
func (a *DbsAPIService) GetSuggestedPrices(ctx context.Context, campaignId int64) DbsApiGetSuggestedPricesRequest {
	return DbsApiGetSuggestedPricesRequest{
func (a *DbsAPIService) GetSuggestedPricesExecute(r DbsApiGetSuggestedPricesRequest) (*SuggestPricesResponse, *http.Response, error) {
type DbsApiGetWarehousesRequest struct {
func (r DbsApiGetWarehousesRequest) Execute() (*GetWarehousesResponse, *http.Response, error) {
	@return DbsApiGetWarehousesRequest
func (a *DbsAPIService) GetWarehouses(ctx context.Context, businessId int64) DbsApiGetWarehousesRequest {
	return DbsApiGetWarehousesRequest{
func (a *DbsAPIService) GetWarehousesExecute(r DbsApiGetWarehousesRequest) (*GetWarehousesResponse, *http.Response, error) {
type DbsApiProvideOrderDigitalCodesRequest struct {
func (r DbsApiProvideOrderDigitalCodesRequest) ProvideOrderDigitalCodesRequest(provideOrderDigitalCodesRequest ProvideOrderDigitalCodesRequest) DbsApiProvideOrderDigitalCodesRequest {
func (r DbsApiProvideOrderDigitalCodesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiProvideOrderDigitalCodesRequest
func (a *DbsAPIService) ProvideOrderDigitalCodes(ctx context.Context, campaignId int64, orderId int64) DbsApiProvideOrderDigitalCodesRequest {
	return DbsApiProvideOrderDigitalCodesRequest{
func (a *DbsAPIService) ProvideOrderDigitalCodesExecute(r DbsApiProvideOrderDigitalCodesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiProvideOrderItemIdentifiersRequest struct {
func (r DbsApiProvideOrderItemIdentifiersRequest) ProvideOrderItemIdentifiersRequest(provideOrderItemIdentifiersRequest ProvideOrderItemIdentifiersRequest) DbsApiProvideOrderItemIdentifiersRequest {
func (r DbsApiProvideOrderItemIdentifiersRequest) Execute() (*ProvideOrderItemIdentifiersResponse, *http.Response, error) {
	@return DbsApiProvideOrderItemIdentifiersRequest
func (a *DbsAPIService) ProvideOrderItemIdentifiers(ctx context.Context, campaignId int64, orderId int64) DbsApiProvideOrderItemIdentifiersRequest {
	return DbsApiProvideOrderItemIdentifiersRequest{
func (a *DbsAPIService) ProvideOrderItemIdentifiersExecute(r DbsApiProvideOrderItemIdentifiersRequest) (*ProvideOrderItemIdentifiersResponse, *http.Response, error) {
type DbsApiPutBidsForBusinessRequest struct {
func (r DbsApiPutBidsForBusinessRequest) PutSkuBidsRequest(putSkuBidsRequest PutSkuBidsRequest) DbsApiPutBidsForBusinessRequest {
func (r DbsApiPutBidsForBusinessRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiPutBidsForBusinessRequest
func (a *DbsAPIService) PutBidsForBusiness(ctx context.Context, businessId int64) DbsApiPutBidsForBusinessRequest {
	return DbsApiPutBidsForBusinessRequest{
func (a *DbsAPIService) PutBidsForBusinessExecute(r DbsApiPutBidsForBusinessRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiPutBidsForCampaignRequest struct {
func (r DbsApiPutBidsForCampaignRequest) PutSkuBidsRequest(putSkuBidsRequest PutSkuBidsRequest) DbsApiPutBidsForCampaignRequest {
func (r DbsApiPutBidsForCampaignRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiPutBidsForCampaignRequest
func (a *DbsAPIService) PutBidsForCampaign(ctx context.Context, campaignId int64) DbsApiPutBidsForCampaignRequest {
	return DbsApiPutBidsForCampaignRequest{
func (a *DbsAPIService) PutBidsForCampaignExecute(r DbsApiPutBidsForCampaignRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiSearchModelsRequest struct {
func (r DbsApiSearchModelsRequest) Query(query string) DbsApiSearchModelsRequest {
func (r DbsApiSearchModelsRequest) RegionId(regionId int64) DbsApiSearchModelsRequest {
func (r DbsApiSearchModelsRequest) Currency(currency CurrencyType) DbsApiSearchModelsRequest {
func (r DbsApiSearchModelsRequest) Page(page int32) DbsApiSearchModelsRequest {
func (r DbsApiSearchModelsRequest) PageSize(pageSize int32) DbsApiSearchModelsRequest {
func (r DbsApiSearchModelsRequest) Execute() (*SearchModelsResponse, *http.Response, error) {
	@return DbsApiSearchModelsRequest
func (a *DbsAPIService) SearchModels(ctx context.Context) DbsApiSearchModelsRequest {
	return DbsApiSearchModelsRequest{
func (a *DbsAPIService) SearchModelsExecute(r DbsApiSearchModelsRequest) (*SearchModelsResponse, *http.Response, error) {
type DbsApiSearchRegionChildrenRequest struct {
func (r DbsApiSearchRegionChildrenRequest) Page(page int32) DbsApiSearchRegionChildrenRequest {
func (r DbsApiSearchRegionChildrenRequest) PageSize(pageSize int32) DbsApiSearchRegionChildrenRequest {
func (r DbsApiSearchRegionChildrenRequest) Execute() (*GetRegionWithChildrenResponse, *http.Response, error) {
	@return DbsApiSearchRegionChildrenRequest
func (a *DbsAPIService) SearchRegionChildren(ctx context.Context, regionId int64) DbsApiSearchRegionChildrenRequest {
	return DbsApiSearchRegionChildrenRequest{
func (a *DbsAPIService) SearchRegionChildrenExecute(r DbsApiSearchRegionChildrenRequest) (*GetRegionWithChildrenResponse, *http.Response, error) {
type DbsApiSearchRegionsByIdRequest struct {
func (r DbsApiSearchRegionsByIdRequest) Execute() (*GetRegionsResponse, *http.Response, error) {
	@return DbsApiSearchRegionsByIdRequest
func (a *DbsAPIService) SearchRegionsById(ctx context.Context, regionId int64) DbsApiSearchRegionsByIdRequest {
	return DbsApiSearchRegionsByIdRequest{
func (a *DbsAPIService) SearchRegionsByIdExecute(r DbsApiSearchRegionsByIdRequest) (*GetRegionsResponse, *http.Response, error) {
type DbsApiSearchRegionsByNameRequest struct {
func (r DbsApiSearchRegionsByNameRequest) Name(name string) DbsApiSearchRegionsByNameRequest {
func (r DbsApiSearchRegionsByNameRequest) PageToken(pageToken string) DbsApiSearchRegionsByNameRequest {
func (r DbsApiSearchRegionsByNameRequest) Limit(limit int32) DbsApiSearchRegionsByNameRequest {
func (r DbsApiSearchRegionsByNameRequest) Execute() (*GetRegionsResponse, *http.Response, error) {
	@return DbsApiSearchRegionsByNameRequest
func (a *DbsAPIService) SearchRegionsByName(ctx context.Context) DbsApiSearchRegionsByNameRequest {
	return DbsApiSearchRegionsByNameRequest{
func (a *DbsAPIService) SearchRegionsByNameExecute(r DbsApiSearchRegionsByNameRequest) (*GetRegionsResponse, *http.Response, error) {
type DbsApiSendFileToChatRequest struct {
func (r DbsApiSendFileToChatRequest) ChatId(chatId int64) DbsApiSendFileToChatRequest {
func (r DbsApiSendFileToChatRequest) File(file *os.File) DbsApiSendFileToChatRequest {
func (r DbsApiSendFileToChatRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiSendFileToChatRequest
func (a *DbsAPIService) SendFileToChat(ctx context.Context, businessId int64) DbsApiSendFileToChatRequest {
	return DbsApiSendFileToChatRequest{
func (a *DbsAPIService) SendFileToChatExecute(r DbsApiSendFileToChatRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiSendMessageToChatRequest struct {
func (r DbsApiSendMessageToChatRequest) ChatId(chatId int64) DbsApiSendMessageToChatRequest {
func (r DbsApiSendMessageToChatRequest) SendMessageToChatRequest(sendMessageToChatRequest SendMessageToChatRequest) DbsApiSendMessageToChatRequest {
func (r DbsApiSendMessageToChatRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiSendMessageToChatRequest
func (a *DbsAPIService) SendMessageToChat(ctx context.Context, businessId int64) DbsApiSendMessageToChatRequest {
	return DbsApiSendMessageToChatRequest{
func (a *DbsAPIService) SendMessageToChatExecute(r DbsApiSendMessageToChatRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiSetOrderBoxLayoutRequest struct {
func (r DbsApiSetOrderBoxLayoutRequest) SetOrderBoxLayoutRequest(setOrderBoxLayoutRequest SetOrderBoxLayoutRequest) DbsApiSetOrderBoxLayoutRequest {
func (r DbsApiSetOrderBoxLayoutRequest) Execute() (*SetOrderBoxLayoutResponse, *http.Response, error) {
	@return DbsApiSetOrderBoxLayoutRequest
func (a *DbsAPIService) SetOrderBoxLayout(ctx context.Context, campaignId int64, orderId int64) DbsApiSetOrderBoxLayoutRequest {
	return DbsApiSetOrderBoxLayoutRequest{
func (a *DbsAPIService) SetOrderBoxLayoutExecute(r DbsApiSetOrderBoxLayoutRequest) (*SetOrderBoxLayoutResponse, *http.Response, error) {
type DbsApiSetOrderDeliveryDateRequest struct {
func (r DbsApiSetOrderDeliveryDateRequest) SetOrderDeliveryDateRequest(setOrderDeliveryDateRequest SetOrderDeliveryDateRequest) DbsApiSetOrderDeliveryDateRequest {
func (r DbsApiSetOrderDeliveryDateRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiSetOrderDeliveryDateRequest
func (a *DbsAPIService) SetOrderDeliveryDate(ctx context.Context, campaignId int64, orderId int64) DbsApiSetOrderDeliveryDateRequest {
	return DbsApiSetOrderDeliveryDateRequest{
func (a *DbsAPIService) SetOrderDeliveryDateExecute(r DbsApiSetOrderDeliveryDateRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiSetOrderDeliveryTrackCodeRequest struct {
func (r DbsApiSetOrderDeliveryTrackCodeRequest) SetOrderDeliveryTrackCodeRequest(setOrderDeliveryTrackCodeRequest SetOrderDeliveryTrackCodeRequest) DbsApiSetOrderDeliveryTrackCodeRequest {
func (r DbsApiSetOrderDeliveryTrackCodeRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiSetOrderDeliveryTrackCodeRequest
func (a *DbsAPIService) SetOrderDeliveryTrackCode(ctx context.Context, campaignId int64, orderId int64) DbsApiSetOrderDeliveryTrackCodeRequest {
	return DbsApiSetOrderDeliveryTrackCodeRequest{
func (a *DbsAPIService) SetOrderDeliveryTrackCodeExecute(r DbsApiSetOrderDeliveryTrackCodeRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiSetOrderShipmentBoxesRequest struct {
func (r DbsApiSetOrderShipmentBoxesRequest) SetOrderShipmentBoxesRequest(setOrderShipmentBoxesRequest SetOrderShipmentBoxesRequest) DbsApiSetOrderShipmentBoxesRequest {
func (r DbsApiSetOrderShipmentBoxesRequest) Execute() (*SetOrderShipmentBoxesResponse, *http.Response, error) {
	@return DbsApiSetOrderShipmentBoxesRequest
func (a *DbsAPIService) SetOrderShipmentBoxes(ctx context.Context, campaignId int64, orderId int64, shipmentId int64) DbsApiSetOrderShipmentBoxesRequest {
	return DbsApiSetOrderShipmentBoxesRequest{
func (a *DbsAPIService) SetOrderShipmentBoxesExecute(r DbsApiSetOrderShipmentBoxesRequest) (*SetOrderShipmentBoxesResponse, *http.Response, error) {
type DbsApiSetReturnDecisionRequest struct {
func (r DbsApiSetReturnDecisionRequest) SetReturnDecisionRequest(setReturnDecisionRequest SetReturnDecisionRequest) DbsApiSetReturnDecisionRequest {
func (r DbsApiSetReturnDecisionRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiSetReturnDecisionRequest
func (a *DbsAPIService) SetReturnDecision(ctx context.Context, campaignId int64, orderId int64, returnId int64) DbsApiSetReturnDecisionRequest {
	return DbsApiSetReturnDecisionRequest{
func (a *DbsAPIService) SetReturnDecisionExecute(r DbsApiSetReturnDecisionRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiSkipGoodsFeedbacksReactionRequest struct {
func (r DbsApiSkipGoodsFeedbacksReactionRequest) SkipGoodsFeedbackReactionRequest(skipGoodsFeedbackReactionRequest SkipGoodsFeedbackReactionRequest) DbsApiSkipGoodsFeedbacksReactionRequest {
func (r DbsApiSkipGoodsFeedbacksReactionRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiSkipGoodsFeedbacksReactionRequest
func (a *DbsAPIService) SkipGoodsFeedbacksReaction(ctx context.Context, businessId int64) DbsApiSkipGoodsFeedbacksReactionRequest {
	return DbsApiSkipGoodsFeedbacksReactionRequest{
func (a *DbsAPIService) SkipGoodsFeedbacksReactionExecute(r DbsApiSkipGoodsFeedbacksReactionRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiSubmitReturnDecisionRequest struct {
func (r DbsApiSubmitReturnDecisionRequest) Body(body map[string]interface{}) DbsApiSubmitReturnDecisionRequest {
func (r DbsApiSubmitReturnDecisionRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiSubmitReturnDecisionRequest
func (a *DbsAPIService) SubmitReturnDecision(ctx context.Context, campaignId int64, orderId int64, returnId int64) DbsApiSubmitReturnDecisionRequest {
	return DbsApiSubmitReturnDecisionRequest{
func (a *DbsAPIService) SubmitReturnDecisionExecute(r DbsApiSubmitReturnDecisionRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateBusinessPricesRequest struct {
func (r DbsApiUpdateBusinessPricesRequest) UpdateBusinessPricesRequest(updateBusinessPricesRequest UpdateBusinessPricesRequest) DbsApiUpdateBusinessPricesRequest {
func (r DbsApiUpdateBusinessPricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateBusinessPricesRequest
func (a *DbsAPIService) UpdateBusinessPrices(ctx context.Context, businessId int64) DbsApiUpdateBusinessPricesRequest {
	return DbsApiUpdateBusinessPricesRequest{
func (a *DbsAPIService) UpdateBusinessPricesExecute(r DbsApiUpdateBusinessPricesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateCampaignOffersRequest struct {
func (r DbsApiUpdateCampaignOffersRequest) UpdateCampaignOffersRequest(updateCampaignOffersRequest UpdateCampaignOffersRequest) DbsApiUpdateCampaignOffersRequest {
func (r DbsApiUpdateCampaignOffersRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateCampaignOffersRequest
func (a *DbsAPIService) UpdateCampaignOffers(ctx context.Context, campaignId int64) DbsApiUpdateCampaignOffersRequest {
	return DbsApiUpdateCampaignOffersRequest{
func (a *DbsAPIService) UpdateCampaignOffersExecute(r DbsApiUpdateCampaignOffersRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateExternalOrderIdRequest struct {
func (r DbsApiUpdateExternalOrderIdRequest) UpdateExternalOrderIdRequest(updateExternalOrderIdRequest UpdateExternalOrderIdRequest) DbsApiUpdateExternalOrderIdRequest {
func (r DbsApiUpdateExternalOrderIdRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateExternalOrderIdRequest
func (a *DbsAPIService) UpdateExternalOrderId(ctx context.Context, campaignId int64, orderId int64) DbsApiUpdateExternalOrderIdRequest {
	return DbsApiUpdateExternalOrderIdRequest{
func (a *DbsAPIService) UpdateExternalOrderIdExecute(r DbsApiUpdateExternalOrderIdRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateGoodsFeedbackCommentRequest struct {
func (r DbsApiUpdateGoodsFeedbackCommentRequest) UpdateGoodsFeedbackCommentRequest(updateGoodsFeedbackCommentRequest UpdateGoodsFeedbackCommentRequest) DbsApiUpdateGoodsFeedbackCommentRequest {
func (r DbsApiUpdateGoodsFeedbackCommentRequest) Execute() (*UpdateGoodsFeedbackCommentResponse, *http.Response, error) {
	@return DbsApiUpdateGoodsFeedbackCommentRequest
func (a *DbsAPIService) UpdateGoodsFeedbackComment(ctx context.Context, businessId int64) DbsApiUpdateGoodsFeedbackCommentRequest {
	return DbsApiUpdateGoodsFeedbackCommentRequest{
func (a *DbsAPIService) UpdateGoodsFeedbackCommentExecute(r DbsApiUpdateGoodsFeedbackCommentRequest) (*UpdateGoodsFeedbackCommentResponse, *http.Response, error) {
type DbsApiUpdateOfferContentRequest struct {
func (r DbsApiUpdateOfferContentRequest) UpdateOfferContentRequest(updateOfferContentRequest UpdateOfferContentRequest) DbsApiUpdateOfferContentRequest {
func (r DbsApiUpdateOfferContentRequest) Execute() (*UpdateOfferContentResponse, *http.Response, error) {
	@return DbsApiUpdateOfferContentRequest
func (a *DbsAPIService) UpdateOfferContent(ctx context.Context, businessId int64) DbsApiUpdateOfferContentRequest {
	return DbsApiUpdateOfferContentRequest{
func (a *DbsAPIService) UpdateOfferContentExecute(r DbsApiUpdateOfferContentRequest) (*UpdateOfferContentResponse, *http.Response, error) {
type DbsApiUpdateOfferMappingEntriesRequest struct {
func (r DbsApiUpdateOfferMappingEntriesRequest) UpdateOfferMappingEntryRequest(updateOfferMappingEntryRequest UpdateOfferMappingEntryRequest) DbsApiUpdateOfferMappingEntriesRequest {
func (r DbsApiUpdateOfferMappingEntriesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateOfferMappingEntriesRequest
func (a *DbsAPIService) UpdateOfferMappingEntries(ctx context.Context, campaignId int64) DbsApiUpdateOfferMappingEntriesRequest {
	return DbsApiUpdateOfferMappingEntriesRequest{
func (a *DbsAPIService) UpdateOfferMappingEntriesExecute(r DbsApiUpdateOfferMappingEntriesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateOfferMappingsRequest struct {
func (r DbsApiUpdateOfferMappingsRequest) UpdateOfferMappingsRequest(updateOfferMappingsRequest UpdateOfferMappingsRequest) DbsApiUpdateOfferMappingsRequest {
func (r DbsApiUpdateOfferMappingsRequest) Language(language CatalogLanguageType) DbsApiUpdateOfferMappingsRequest {
func (r DbsApiUpdateOfferMappingsRequest) Execute() (*UpdateOfferMappingsResponse, *http.Response, error) {
	@return DbsApiUpdateOfferMappingsRequest
func (a *DbsAPIService) UpdateOfferMappings(ctx context.Context, businessId int64) DbsApiUpdateOfferMappingsRequest {
	return DbsApiUpdateOfferMappingsRequest{
func (a *DbsAPIService) UpdateOfferMappingsExecute(r DbsApiUpdateOfferMappingsRequest) (*UpdateOfferMappingsResponse, *http.Response, error) {
type DbsApiUpdateOrderItemsRequest struct {
func (r DbsApiUpdateOrderItemsRequest) UpdateOrderItemRequest(updateOrderItemRequest UpdateOrderItemRequest) DbsApiUpdateOrderItemsRequest {
func (r DbsApiUpdateOrderItemsRequest) Execute() (*http.Response, error) {
	@return DbsApiUpdateOrderItemsRequest
func (a *DbsAPIService) UpdateOrderItems(ctx context.Context, campaignId int64, orderId int64) DbsApiUpdateOrderItemsRequest {
	return DbsApiUpdateOrderItemsRequest{
func (a *DbsAPIService) UpdateOrderItemsExecute(r DbsApiUpdateOrderItemsRequest) (*http.Response, error) {
type DbsApiUpdateOrderStatusRequest struct {
func (r DbsApiUpdateOrderStatusRequest) UpdateOrderStatusRequest(updateOrderStatusRequest UpdateOrderStatusRequest) DbsApiUpdateOrderStatusRequest {
func (r DbsApiUpdateOrderStatusRequest) Execute() (*UpdateOrderStatusResponse, *http.Response, error) {
	@return DbsApiUpdateOrderStatusRequest
func (a *DbsAPIService) UpdateOrderStatus(ctx context.Context, campaignId int64, orderId int64) DbsApiUpdateOrderStatusRequest {
	return DbsApiUpdateOrderStatusRequest{
func (a *DbsAPIService) UpdateOrderStatusExecute(r DbsApiUpdateOrderStatusRequest) (*UpdateOrderStatusResponse, *http.Response, error) {
type DbsApiUpdateOrderStatusesRequest struct {
func (r DbsApiUpdateOrderStatusesRequest) UpdateOrderStatusesRequest(updateOrderStatusesRequest UpdateOrderStatusesRequest) DbsApiUpdateOrderStatusesRequest {
func (r DbsApiUpdateOrderStatusesRequest) Execute() (*UpdateOrderStatusesResponse, *http.Response, error) {
	@return DbsApiUpdateOrderStatusesRequest
func (a *DbsAPIService) UpdateOrderStatuses(ctx context.Context, campaignId int64) DbsApiUpdateOrderStatusesRequest {
	return DbsApiUpdateOrderStatusesRequest{
func (a *DbsAPIService) UpdateOrderStatusesExecute(r DbsApiUpdateOrderStatusesRequest) (*UpdateOrderStatusesResponse, *http.Response, error) {
type DbsApiUpdateOrderStorageLimitRequest struct {
func (r DbsApiUpdateOrderStorageLimitRequest) UpdateOrderStorageLimitRequest(updateOrderStorageLimitRequest UpdateOrderStorageLimitRequest) DbsApiUpdateOrderStorageLimitRequest {
func (r DbsApiUpdateOrderStorageLimitRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateOrderStorageLimitRequest
func (a *DbsAPIService) UpdateOrderStorageLimit(ctx context.Context, campaignId int64, orderId int64) DbsApiUpdateOrderStorageLimitRequest {
	return DbsApiUpdateOrderStorageLimitRequest{
func (a *DbsAPIService) UpdateOrderStorageLimitExecute(r DbsApiUpdateOrderStorageLimitRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateOutletRequest struct {
func (r DbsApiUpdateOutletRequest) ChangeOutletRequest(changeOutletRequest ChangeOutletRequest) DbsApiUpdateOutletRequest {
func (r DbsApiUpdateOutletRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateOutletRequest
func (a *DbsAPIService) UpdateOutlet(ctx context.Context, campaignId int64, outletId int64) DbsApiUpdateOutletRequest {
	return DbsApiUpdateOutletRequest{
func (a *DbsAPIService) UpdateOutletExecute(r DbsApiUpdateOutletRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateOutletLicensesRequest struct {
func (r DbsApiUpdateOutletLicensesRequest) UpdateOutletLicenseRequest(updateOutletLicenseRequest UpdateOutletLicenseRequest) DbsApiUpdateOutletLicensesRequest {
func (r DbsApiUpdateOutletLicensesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateOutletLicensesRequest
func (a *DbsAPIService) UpdateOutletLicenses(ctx context.Context, campaignId int64) DbsApiUpdateOutletLicensesRequest {
	return DbsApiUpdateOutletLicensesRequest{
func (a *DbsAPIService) UpdateOutletLicensesExecute(r DbsApiUpdateOutletLicensesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdatePricesRequest struct {
func (r DbsApiUpdatePricesRequest) UpdatePricesRequest(updatePricesRequest UpdatePricesRequest) DbsApiUpdatePricesRequest {
func (r DbsApiUpdatePricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdatePricesRequest
func (a *DbsAPIService) UpdatePrices(ctx context.Context, campaignId int64) DbsApiUpdatePricesRequest {
	return DbsApiUpdatePricesRequest{
func (a *DbsAPIService) UpdatePricesExecute(r DbsApiUpdatePricesRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdatePromoOffersRequest struct {
func (r DbsApiUpdatePromoOffersRequest) UpdatePromoOffersRequest(updatePromoOffersRequest UpdatePromoOffersRequest) DbsApiUpdatePromoOffersRequest {
func (r DbsApiUpdatePromoOffersRequest) Execute() (*UpdatePromoOffersResponse, *http.Response, error) {
	@return DbsApiUpdatePromoOffersRequest
func (a *DbsAPIService) UpdatePromoOffers(ctx context.Context, businessId int64) DbsApiUpdatePromoOffersRequest {
	return DbsApiUpdatePromoOffersRequest{
func (a *DbsAPIService) UpdatePromoOffersExecute(r DbsApiUpdatePromoOffersRequest) (*UpdatePromoOffersResponse, *http.Response, error) {
type DbsApiUpdateStocksRequest struct {
func (r DbsApiUpdateStocksRequest) UpdateStocksRequest(updateStocksRequest UpdateStocksRequest) DbsApiUpdateStocksRequest {
func (r DbsApiUpdateStocksRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return DbsApiUpdateStocksRequest
func (a *DbsAPIService) UpdateStocks(ctx context.Context, campaignId int64) DbsApiUpdateStocksRequest {
	return DbsApiUpdateStocksRequest{
func (a *DbsAPIService) UpdateStocksExecute(r DbsApiUpdateStocksRequest) (*EmptyApiResponse, *http.Response, error) {
type DbsApiUpdateWarehouseStatusRequest struct {
func (r DbsApiUpdateWarehouseStatusRequest) UpdateWarehouseStatusRequest(updateWarehouseStatusRequest UpdateWarehouseStatusRequest) DbsApiUpdateWarehouseStatusRequest {
func (r DbsApiUpdateWarehouseStatusRequest) Execute() (*UpdateWarehouseStatusResponse, *http.Response, error) {
	@return DbsApiUpdateWarehouseStatusRequest
func (a *DbsAPIService) UpdateWarehouseStatus(ctx context.Context, campaignId int64) DbsApiUpdateWarehouseStatusRequest {
	return DbsApiUpdateWarehouseStatusRequest{
func (a *DbsAPIService) UpdateWarehouseStatusExecute(r DbsApiUpdateWarehouseStatusRequest) (*UpdateWarehouseStatusResponse, *http.Response, error) {
