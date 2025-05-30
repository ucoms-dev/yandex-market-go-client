type FbsApiAddHiddenOffersRequest struct {
func (r FbsApiAddHiddenOffersRequest) AddHiddenOffersRequest(addHiddenOffersRequest AddHiddenOffersRequest) FbsApiAddHiddenOffersRequest {
func (r FbsApiAddHiddenOffersRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiAddHiddenOffersRequest
func (a *FbsAPIService) AddHiddenOffers(ctx context.Context, campaignId int64) FbsApiAddHiddenOffersRequest {
	return FbsApiAddHiddenOffersRequest{
func (a *FbsAPIService) AddHiddenOffersExecute(r FbsApiAddHiddenOffersRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiAddOffersToArchiveRequest struct {
func (r FbsApiAddOffersToArchiveRequest) AddOffersToArchiveRequest(addOffersToArchiveRequest AddOffersToArchiveRequest) FbsApiAddOffersToArchiveRequest {
func (r FbsApiAddOffersToArchiveRequest) Execute() (*AddOffersToArchiveResponse, *http.Response, error) {
	@return FbsApiAddOffersToArchiveRequest
func (a *FbsAPIService) AddOffersToArchive(ctx context.Context, businessId int64) FbsApiAddOffersToArchiveRequest {
	return FbsApiAddOffersToArchiveRequest{
func (a *FbsAPIService) AddOffersToArchiveExecute(r FbsApiAddOffersToArchiveRequest) (*AddOffersToArchiveResponse, *http.Response, error) {
type FbsApiCalculateTariffsRequest struct {
func (r FbsApiCalculateTariffsRequest) CalculateTariffsRequest(calculateTariffsRequest CalculateTariffsRequest) FbsApiCalculateTariffsRequest {
func (r FbsApiCalculateTariffsRequest) Execute() (*CalculateTariffsResponse, *http.Response, error) {
	@return FbsApiCalculateTariffsRequest
func (a *FbsAPIService) CalculateTariffs(ctx context.Context) FbsApiCalculateTariffsRequest {
	return FbsApiCalculateTariffsRequest{
func (a *FbsAPIService) CalculateTariffsExecute(r FbsApiCalculateTariffsRequest) (*CalculateTariffsResponse, *http.Response, error) {
type FbsApiConfirmBusinessPricesRequest struct {
func (r FbsApiConfirmBusinessPricesRequest) ConfirmPricesRequest(confirmPricesRequest ConfirmPricesRequest) FbsApiConfirmBusinessPricesRequest {
func (r FbsApiConfirmBusinessPricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiConfirmBusinessPricesRequest
func (a *FbsAPIService) ConfirmBusinessPrices(ctx context.Context, businessId int64) FbsApiConfirmBusinessPricesRequest {
	return FbsApiConfirmBusinessPricesRequest{
func (a *FbsAPIService) ConfirmBusinessPricesExecute(r FbsApiConfirmBusinessPricesRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiConfirmCampaignPricesRequest struct {
func (r FbsApiConfirmCampaignPricesRequest) ConfirmPricesRequest(confirmPricesRequest ConfirmPricesRequest) FbsApiConfirmCampaignPricesRequest {
func (r FbsApiConfirmCampaignPricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiConfirmCampaignPricesRequest
func (a *FbsAPIService) ConfirmCampaignPrices(ctx context.Context, campaignId int64) FbsApiConfirmCampaignPricesRequest {
	return FbsApiConfirmCampaignPricesRequest{
func (a *FbsAPIService) ConfirmCampaignPricesExecute(r FbsApiConfirmCampaignPricesRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiConfirmShipmentRequest struct {
func (r FbsApiConfirmShipmentRequest) ConfirmShipmentRequest(confirmShipmentRequest ConfirmShipmentRequest) FbsApiConfirmShipmentRequest {
func (r FbsApiConfirmShipmentRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiConfirmShipmentRequest
func (a *FbsAPIService) ConfirmShipment(ctx context.Context, campaignId int64, shipmentId int64) FbsApiConfirmShipmentRequest {
	return FbsApiConfirmShipmentRequest{
func (a *FbsAPIService) ConfirmShipmentExecute(r FbsApiConfirmShipmentRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiCreateChatRequest struct {
func (r FbsApiCreateChatRequest) CreateChatRequest(createChatRequest CreateChatRequest) FbsApiCreateChatRequest {
func (r FbsApiCreateChatRequest) Execute() (*CreateChatResponse, *http.Response, error) {
	@return FbsApiCreateChatRequest
func (a *FbsAPIService) CreateChat(ctx context.Context, businessId int64) FbsApiCreateChatRequest {
	return FbsApiCreateChatRequest{
func (a *FbsAPIService) CreateChatExecute(r FbsApiCreateChatRequest) (*CreateChatResponse, *http.Response, error) {
type FbsApiDeleteCampaignOffersRequest struct {
func (r FbsApiDeleteCampaignOffersRequest) DeleteCampaignOffersRequest(deleteCampaignOffersRequest DeleteCampaignOffersRequest) FbsApiDeleteCampaignOffersRequest {
func (r FbsApiDeleteCampaignOffersRequest) Execute() (*DeleteCampaignOffersResponse, *http.Response, error) {
	@return FbsApiDeleteCampaignOffersRequest
func (a *FbsAPIService) DeleteCampaignOffers(ctx context.Context, campaignId int64) FbsApiDeleteCampaignOffersRequest {
	return FbsApiDeleteCampaignOffersRequest{
func (a *FbsAPIService) DeleteCampaignOffersExecute(r FbsApiDeleteCampaignOffersRequest) (*DeleteCampaignOffersResponse, *http.Response, error) {
type FbsApiDeleteGoodsFeedbackCommentRequest struct {
func (r FbsApiDeleteGoodsFeedbackCommentRequest) DeleteGoodsFeedbackCommentRequest(deleteGoodsFeedbackCommentRequest DeleteGoodsFeedbackCommentRequest) FbsApiDeleteGoodsFeedbackCommentRequest {
func (r FbsApiDeleteGoodsFeedbackCommentRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiDeleteGoodsFeedbackCommentRequest
func (a *FbsAPIService) DeleteGoodsFeedbackComment(ctx context.Context, businessId int64) FbsApiDeleteGoodsFeedbackCommentRequest {
	return FbsApiDeleteGoodsFeedbackCommentRequest{
func (a *FbsAPIService) DeleteGoodsFeedbackCommentExecute(r FbsApiDeleteGoodsFeedbackCommentRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiDeleteHiddenOffersRequest struct {
func (r FbsApiDeleteHiddenOffersRequest) DeleteHiddenOffersRequest(deleteHiddenOffersRequest DeleteHiddenOffersRequest) FbsApiDeleteHiddenOffersRequest {
func (r FbsApiDeleteHiddenOffersRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiDeleteHiddenOffersRequest
func (a *FbsAPIService) DeleteHiddenOffers(ctx context.Context, campaignId int64) FbsApiDeleteHiddenOffersRequest {
	return FbsApiDeleteHiddenOffersRequest{
func (a *FbsAPIService) DeleteHiddenOffersExecute(r FbsApiDeleteHiddenOffersRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiDeleteOffersRequest struct {
func (r FbsApiDeleteOffersRequest) DeleteOffersRequest(deleteOffersRequest DeleteOffersRequest) FbsApiDeleteOffersRequest {
func (r FbsApiDeleteOffersRequest) Execute() (*DeleteOffersResponse, *http.Response, error) {
	@return FbsApiDeleteOffersRequest
func (a *FbsAPIService) DeleteOffers(ctx context.Context, businessId int64) FbsApiDeleteOffersRequest {
	return FbsApiDeleteOffersRequest{
func (a *FbsAPIService) DeleteOffersExecute(r FbsApiDeleteOffersRequest) (*DeleteOffersResponse, *http.Response, error) {
type FbsApiDeleteOffersFromArchiveRequest struct {
func (r FbsApiDeleteOffersFromArchiveRequest) DeleteOffersFromArchiveRequest(deleteOffersFromArchiveRequest DeleteOffersFromArchiveRequest) FbsApiDeleteOffersFromArchiveRequest {
func (r FbsApiDeleteOffersFromArchiveRequest) Execute() (*DeleteOffersFromArchiveResponse, *http.Response, error) {
	@return FbsApiDeleteOffersFromArchiveRequest
func (a *FbsAPIService) DeleteOffersFromArchive(ctx context.Context, businessId int64) FbsApiDeleteOffersFromArchiveRequest {
	return FbsApiDeleteOffersFromArchiveRequest{
func (a *FbsAPIService) DeleteOffersFromArchiveExecute(r FbsApiDeleteOffersFromArchiveRequest) (*DeleteOffersFromArchiveResponse, *http.Response, error) {
type FbsApiDeletePromoOffersRequest struct {
func (r FbsApiDeletePromoOffersRequest) DeletePromoOffersRequest(deletePromoOffersRequest DeletePromoOffersRequest) FbsApiDeletePromoOffersRequest {
func (r FbsApiDeletePromoOffersRequest) Execute() (*DeletePromoOffersResponse, *http.Response, error) {
	@return FbsApiDeletePromoOffersRequest
func (a *FbsAPIService) DeletePromoOffers(ctx context.Context, businessId int64) FbsApiDeletePromoOffersRequest {
	return FbsApiDeletePromoOffersRequest{
func (a *FbsAPIService) DeletePromoOffersExecute(r FbsApiDeletePromoOffersRequest) (*DeletePromoOffersResponse, *http.Response, error) {
type FbsApiDownloadShipmentActRequest struct {
func (r FbsApiDownloadShipmentActRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiDownloadShipmentActRequest
func (a *FbsAPIService) DownloadShipmentAct(ctx context.Context, campaignId int64, shipmentId int64) FbsApiDownloadShipmentActRequest {
	return FbsApiDownloadShipmentActRequest{
func (a *FbsAPIService) DownloadShipmentActExecute(r FbsApiDownloadShipmentActRequest) (*os.File, *http.Response, error) {
type FbsApiDownloadShipmentDiscrepancyActRequest struct {
func (r FbsApiDownloadShipmentDiscrepancyActRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiDownloadShipmentDiscrepancyActRequest
func (a *FbsAPIService) DownloadShipmentDiscrepancyAct(ctx context.Context, campaignId int64, shipmentId int64) FbsApiDownloadShipmentDiscrepancyActRequest {
	return FbsApiDownloadShipmentDiscrepancyActRequest{
func (a *FbsAPIService) DownloadShipmentDiscrepancyActExecute(r FbsApiDownloadShipmentDiscrepancyActRequest) (*os.File, *http.Response, error) {
type FbsApiDownloadShipmentInboundActRequest struct {
func (r FbsApiDownloadShipmentInboundActRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiDownloadShipmentInboundActRequest
func (a *FbsAPIService) DownloadShipmentInboundAct(ctx context.Context, campaignId int64, shipmentId int64) FbsApiDownloadShipmentInboundActRequest {
	return FbsApiDownloadShipmentInboundActRequest{
func (a *FbsAPIService) DownloadShipmentInboundActExecute(r FbsApiDownloadShipmentInboundActRequest) (*os.File, *http.Response, error) {
type FbsApiDownloadShipmentPalletLabelsRequest struct {
func (r FbsApiDownloadShipmentPalletLabelsRequest) Format(format ShipmentPalletLabelPageFormatType) FbsApiDownloadShipmentPalletLabelsRequest {
func (r FbsApiDownloadShipmentPalletLabelsRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiDownloadShipmentPalletLabelsRequest
func (a *FbsAPIService) DownloadShipmentPalletLabels(ctx context.Context, campaignId int64, shipmentId int64) FbsApiDownloadShipmentPalletLabelsRequest {
	return FbsApiDownloadShipmentPalletLabelsRequest{
func (a *FbsAPIService) DownloadShipmentPalletLabelsExecute(r FbsApiDownloadShipmentPalletLabelsRequest) (*os.File, *http.Response, error) {
type FbsApiDownloadShipmentReceptionTransferActRequest struct {
func (r FbsApiDownloadShipmentReceptionTransferActRequest) WarehouseId(warehouseId int32) FbsApiDownloadShipmentReceptionTransferActRequest {
func (r FbsApiDownloadShipmentReceptionTransferActRequest) Signatory(signatory string) FbsApiDownloadShipmentReceptionTransferActRequest {
func (r FbsApiDownloadShipmentReceptionTransferActRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiDownloadShipmentReceptionTransferActRequest
func (a *FbsAPIService) DownloadShipmentReceptionTransferAct(ctx context.Context, campaignId int64) FbsApiDownloadShipmentReceptionTransferActRequest {
	return FbsApiDownloadShipmentReceptionTransferActRequest{
func (a *FbsAPIService) DownloadShipmentReceptionTransferActExecute(r FbsApiDownloadShipmentReceptionTransferActRequest) (*os.File, *http.Response, error) {
type FbsApiDownloadShipmentTransportationWaybillRequest struct {
func (r FbsApiDownloadShipmentTransportationWaybillRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiDownloadShipmentTransportationWaybillRequest
func (a *FbsAPIService) DownloadShipmentTransportationWaybill(ctx context.Context, campaignId int64, shipmentId int64) FbsApiDownloadShipmentTransportationWaybillRequest {
	return FbsApiDownloadShipmentTransportationWaybillRequest{
func (a *FbsAPIService) DownloadShipmentTransportationWaybillExecute(r FbsApiDownloadShipmentTransportationWaybillRequest) (*os.File, *http.Response, error) {
type FbsApiGenerateBannersStatisticsReportRequest struct {
func (r FbsApiGenerateBannersStatisticsReportRequest) GenerateBannersStatisticsRequest(generateBannersStatisticsRequest GenerateBannersStatisticsRequest) FbsApiGenerateBannersStatisticsReportRequest {
func (r FbsApiGenerateBannersStatisticsReportRequest) Format(format ReportFormatType) FbsApiGenerateBannersStatisticsReportRequest {
func (r FbsApiGenerateBannersStatisticsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateBannersStatisticsReportRequest
func (a *FbsAPIService) GenerateBannersStatisticsReport(ctx context.Context) FbsApiGenerateBannersStatisticsReportRequest {
	return FbsApiGenerateBannersStatisticsReportRequest{
func (a *FbsAPIService) GenerateBannersStatisticsReportExecute(r FbsApiGenerateBannersStatisticsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateBoostConsolidatedReportRequest struct {
func (r FbsApiGenerateBoostConsolidatedReportRequest) GenerateBoostConsolidatedRequest(generateBoostConsolidatedRequest GenerateBoostConsolidatedRequest) FbsApiGenerateBoostConsolidatedReportRequest {
func (r FbsApiGenerateBoostConsolidatedReportRequest) Format(format ReportFormatType) FbsApiGenerateBoostConsolidatedReportRequest {
func (r FbsApiGenerateBoostConsolidatedReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateBoostConsolidatedReportRequest
func (a *FbsAPIService) GenerateBoostConsolidatedReport(ctx context.Context) FbsApiGenerateBoostConsolidatedReportRequest {
	return FbsApiGenerateBoostConsolidatedReportRequest{
func (a *FbsAPIService) GenerateBoostConsolidatedReportExecute(r FbsApiGenerateBoostConsolidatedReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateCompetitorsPositionReportRequest struct {
func (r FbsApiGenerateCompetitorsPositionReportRequest) GenerateCompetitorsPositionReportRequest(generateCompetitorsPositionReportRequest GenerateCompetitorsPositionReportRequest) FbsApiGenerateCompetitorsPositionReportRequest {
func (r FbsApiGenerateCompetitorsPositionReportRequest) Format(format ReportFormatType) FbsApiGenerateCompetitorsPositionReportRequest {
func (r FbsApiGenerateCompetitorsPositionReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateCompetitorsPositionReportRequest
func (a *FbsAPIService) GenerateCompetitorsPositionReport(ctx context.Context) FbsApiGenerateCompetitorsPositionReportRequest {
	return FbsApiGenerateCompetitorsPositionReportRequest{
func (a *FbsAPIService) GenerateCompetitorsPositionReportExecute(r FbsApiGenerateCompetitorsPositionReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateGoodsFeedbackReportRequest struct {
func (r FbsApiGenerateGoodsFeedbackReportRequest) GenerateGoodsFeedbackRequest(generateGoodsFeedbackRequest GenerateGoodsFeedbackRequest) FbsApiGenerateGoodsFeedbackReportRequest {
func (r FbsApiGenerateGoodsFeedbackReportRequest) Format(format ReportFormatType) FbsApiGenerateGoodsFeedbackReportRequest {
func (r FbsApiGenerateGoodsFeedbackReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateGoodsFeedbackReportRequest
func (a *FbsAPIService) GenerateGoodsFeedbackReport(ctx context.Context) FbsApiGenerateGoodsFeedbackReportRequest {
	return FbsApiGenerateGoodsFeedbackReportRequest{
func (a *FbsAPIService) GenerateGoodsFeedbackReportExecute(r FbsApiGenerateGoodsFeedbackReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateGoodsRealizationReportRequest struct {
func (r FbsApiGenerateGoodsRealizationReportRequest) GenerateGoodsRealizationReportRequest(generateGoodsRealizationReportRequest GenerateGoodsRealizationReportRequest) FbsApiGenerateGoodsRealizationReportRequest {
func (r FbsApiGenerateGoodsRealizationReportRequest) Format(format ReportFormatType) FbsApiGenerateGoodsRealizationReportRequest {
func (r FbsApiGenerateGoodsRealizationReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateGoodsRealizationReportRequest
func (a *FbsAPIService) GenerateGoodsRealizationReport(ctx context.Context) FbsApiGenerateGoodsRealizationReportRequest {
	return FbsApiGenerateGoodsRealizationReportRequest{
func (a *FbsAPIService) GenerateGoodsRealizationReportExecute(r FbsApiGenerateGoodsRealizationReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateJewelryFiscalReportRequest struct {
func (r FbsApiGenerateJewelryFiscalReportRequest) GenerateJewelryFiscalReportRequest(generateJewelryFiscalReportRequest GenerateJewelryFiscalReportRequest) FbsApiGenerateJewelryFiscalReportRequest {
func (r FbsApiGenerateJewelryFiscalReportRequest) Format(format ReportFormatType) FbsApiGenerateJewelryFiscalReportRequest {
func (r FbsApiGenerateJewelryFiscalReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateJewelryFiscalReportRequest
func (a *FbsAPIService) GenerateJewelryFiscalReport(ctx context.Context) FbsApiGenerateJewelryFiscalReportRequest {
	return FbsApiGenerateJewelryFiscalReportRequest{
func (a *FbsAPIService) GenerateJewelryFiscalReportExecute(r FbsApiGenerateJewelryFiscalReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateMassOrderLabelsReportRequest struct {
func (r FbsApiGenerateMassOrderLabelsReportRequest) GenerateMassOrderLabelsRequest(generateMassOrderLabelsRequest GenerateMassOrderLabelsRequest) FbsApiGenerateMassOrderLabelsReportRequest {
func (r FbsApiGenerateMassOrderLabelsReportRequest) Format(format PageFormatType) FbsApiGenerateMassOrderLabelsReportRequest {
func (r FbsApiGenerateMassOrderLabelsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateMassOrderLabelsReportRequest
func (a *FbsAPIService) GenerateMassOrderLabelsReport(ctx context.Context) FbsApiGenerateMassOrderLabelsReportRequest {
	return FbsApiGenerateMassOrderLabelsReportRequest{
func (a *FbsAPIService) GenerateMassOrderLabelsReportExecute(r FbsApiGenerateMassOrderLabelsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateOrderLabelRequest struct {
func (r FbsApiGenerateOrderLabelRequest) Format(format PageFormatType) FbsApiGenerateOrderLabelRequest {
func (r FbsApiGenerateOrderLabelRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiGenerateOrderLabelRequest
func (a *FbsAPIService) GenerateOrderLabel(ctx context.Context, campaignId int64, orderId int64, shipmentId int64, boxId int64) FbsApiGenerateOrderLabelRequest {
	return FbsApiGenerateOrderLabelRequest{
func (a *FbsAPIService) GenerateOrderLabelExecute(r FbsApiGenerateOrderLabelRequest) (*os.File, *http.Response, error) {
type FbsApiGenerateOrderLabelsRequest struct {
func (r FbsApiGenerateOrderLabelsRequest) Format(format PageFormatType) FbsApiGenerateOrderLabelsRequest {
func (r FbsApiGenerateOrderLabelsRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiGenerateOrderLabelsRequest
func (a *FbsAPIService) GenerateOrderLabels(ctx context.Context, campaignId int64, orderId int64) FbsApiGenerateOrderLabelsRequest {
	return FbsApiGenerateOrderLabelsRequest{
func (a *FbsAPIService) GenerateOrderLabelsExecute(r FbsApiGenerateOrderLabelsRequest) (*os.File, *http.Response, error) {
type FbsApiGeneratePricesReportRequest struct {
func (r FbsApiGeneratePricesReportRequest) GeneratePricesReportRequest(generatePricesReportRequest GeneratePricesReportRequest) FbsApiGeneratePricesReportRequest {
func (r FbsApiGeneratePricesReportRequest) Format(format ReportFormatType) FbsApiGeneratePricesReportRequest {
func (r FbsApiGeneratePricesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGeneratePricesReportRequest
func (a *FbsAPIService) GeneratePricesReport(ctx context.Context) FbsApiGeneratePricesReportRequest {
	return FbsApiGeneratePricesReportRequest{
func (a *FbsAPIService) GeneratePricesReportExecute(r FbsApiGeneratePricesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateSalesGeographyReportRequest struct {
func (r FbsApiGenerateSalesGeographyReportRequest) GenerateSalesGeographyRequest(generateSalesGeographyRequest GenerateSalesGeographyRequest) FbsApiGenerateSalesGeographyReportRequest {
func (r FbsApiGenerateSalesGeographyReportRequest) Format(format ReportFormatType) FbsApiGenerateSalesGeographyReportRequest {
func (r FbsApiGenerateSalesGeographyReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateSalesGeographyReportRequest
func (a *FbsAPIService) GenerateSalesGeographyReport(ctx context.Context) FbsApiGenerateSalesGeographyReportRequest {
	return FbsApiGenerateSalesGeographyReportRequest{
func (a *FbsAPIService) GenerateSalesGeographyReportExecute(r FbsApiGenerateSalesGeographyReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateShelfsStatisticsReportRequest struct {
func (r FbsApiGenerateShelfsStatisticsReportRequest) GenerateShelfsStatisticsRequest(generateShelfsStatisticsRequest GenerateShelfsStatisticsRequest) FbsApiGenerateShelfsStatisticsReportRequest {
func (r FbsApiGenerateShelfsStatisticsReportRequest) Format(format ReportFormatType) FbsApiGenerateShelfsStatisticsReportRequest {
func (r FbsApiGenerateShelfsStatisticsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateShelfsStatisticsReportRequest
func (a *FbsAPIService) GenerateShelfsStatisticsReport(ctx context.Context) FbsApiGenerateShelfsStatisticsReportRequest {
	return FbsApiGenerateShelfsStatisticsReportRequest{
func (a *FbsAPIService) GenerateShelfsStatisticsReportExecute(r FbsApiGenerateShelfsStatisticsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateShipmentListDocumentReportRequest struct {
func (r FbsApiGenerateShipmentListDocumentReportRequest) GenerateShipmentListDocumentReportRequest(generateShipmentListDocumentReportRequest GenerateShipmentListDocumentReportRequest) FbsApiGenerateShipmentListDocumentReportRequest {
func (r FbsApiGenerateShipmentListDocumentReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateShipmentListDocumentReportRequest
func (a *FbsAPIService) GenerateShipmentListDocumentReport(ctx context.Context) FbsApiGenerateShipmentListDocumentReportRequest {
	return FbsApiGenerateShipmentListDocumentReportRequest{
func (a *FbsAPIService) GenerateShipmentListDocumentReportExecute(r FbsApiGenerateShipmentListDocumentReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateShowsBoostReportRequest struct {
func (r FbsApiGenerateShowsBoostReportRequest) GenerateShowsBoostRequest(generateShowsBoostRequest GenerateShowsBoostRequest) FbsApiGenerateShowsBoostReportRequest {
func (r FbsApiGenerateShowsBoostReportRequest) Format(format ReportFormatType) FbsApiGenerateShowsBoostReportRequest {
func (r FbsApiGenerateShowsBoostReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateShowsBoostReportRequest
func (a *FbsAPIService) GenerateShowsBoostReport(ctx context.Context) FbsApiGenerateShowsBoostReportRequest {
	return FbsApiGenerateShowsBoostReportRequest{
func (a *FbsAPIService) GenerateShowsBoostReportExecute(r FbsApiGenerateShowsBoostReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateShowsSalesReportRequest struct {
func (r FbsApiGenerateShowsSalesReportRequest) GenerateShowsSalesReportRequest(generateShowsSalesReportRequest GenerateShowsSalesReportRequest) FbsApiGenerateShowsSalesReportRequest {
func (r FbsApiGenerateShowsSalesReportRequest) Format(format ReportFormatType) FbsApiGenerateShowsSalesReportRequest {
func (r FbsApiGenerateShowsSalesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateShowsSalesReportRequest
func (a *FbsAPIService) GenerateShowsSalesReport(ctx context.Context) FbsApiGenerateShowsSalesReportRequest {
	return FbsApiGenerateShowsSalesReportRequest{
func (a *FbsAPIService) GenerateShowsSalesReportExecute(r FbsApiGenerateShowsSalesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateStocksOnWarehousesReportRequest struct {
func (r FbsApiGenerateStocksOnWarehousesReportRequest) GenerateStocksOnWarehousesReportRequest(generateStocksOnWarehousesReportRequest GenerateStocksOnWarehousesReportRequest) FbsApiGenerateStocksOnWarehousesReportRequest {
func (r FbsApiGenerateStocksOnWarehousesReportRequest) Format(format ReportFormatType) FbsApiGenerateStocksOnWarehousesReportRequest {
func (r FbsApiGenerateStocksOnWarehousesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateStocksOnWarehousesReportRequest
func (a *FbsAPIService) GenerateStocksOnWarehousesReport(ctx context.Context) FbsApiGenerateStocksOnWarehousesReportRequest {
	return FbsApiGenerateStocksOnWarehousesReportRequest{
func (a *FbsAPIService) GenerateStocksOnWarehousesReportExecute(r FbsApiGenerateStocksOnWarehousesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateUnitedMarketplaceServicesReportRequest struct {
func (r FbsApiGenerateUnitedMarketplaceServicesReportRequest) GenerateUnitedMarketplaceServicesReportRequest(generateUnitedMarketplaceServicesReportRequest GenerateUnitedMarketplaceServicesReportRequest) FbsApiGenerateUnitedMarketplaceServicesReportRequest {
func (r FbsApiGenerateUnitedMarketplaceServicesReportRequest) Format(format ReportFormatType) FbsApiGenerateUnitedMarketplaceServicesReportRequest {
func (r FbsApiGenerateUnitedMarketplaceServicesReportRequest) Language(language ReportLanguageType) FbsApiGenerateUnitedMarketplaceServicesReportRequest {
func (r FbsApiGenerateUnitedMarketplaceServicesReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateUnitedMarketplaceServicesReportRequest
func (a *FbsAPIService) GenerateUnitedMarketplaceServicesReport(ctx context.Context) FbsApiGenerateUnitedMarketplaceServicesReportRequest {
	return FbsApiGenerateUnitedMarketplaceServicesReportRequest{
func (a *FbsAPIService) GenerateUnitedMarketplaceServicesReportExecute(r FbsApiGenerateUnitedMarketplaceServicesReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateUnitedNettingReportRequest struct {
func (r FbsApiGenerateUnitedNettingReportRequest) GenerateUnitedNettingReportRequest(generateUnitedNettingReportRequest GenerateUnitedNettingReportRequest) FbsApiGenerateUnitedNettingReportRequest {
func (r FbsApiGenerateUnitedNettingReportRequest) Format(format ReportFormatType) FbsApiGenerateUnitedNettingReportRequest {
func (r FbsApiGenerateUnitedNettingReportRequest) Language(language ReportLanguageType) FbsApiGenerateUnitedNettingReportRequest {
func (r FbsApiGenerateUnitedNettingReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateUnitedNettingReportRequest
func (a *FbsAPIService) GenerateUnitedNettingReport(ctx context.Context) FbsApiGenerateUnitedNettingReportRequest {
	return FbsApiGenerateUnitedNettingReportRequest{
func (a *FbsAPIService) GenerateUnitedNettingReportExecute(r FbsApiGenerateUnitedNettingReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateUnitedOrdersReportRequest struct {
func (r FbsApiGenerateUnitedOrdersReportRequest) GenerateUnitedOrdersRequest(generateUnitedOrdersRequest GenerateUnitedOrdersRequest) FbsApiGenerateUnitedOrdersReportRequest {
func (r FbsApiGenerateUnitedOrdersReportRequest) Format(format ReportFormatType) FbsApiGenerateUnitedOrdersReportRequest {
func (r FbsApiGenerateUnitedOrdersReportRequest) Language(language ReportLanguageType) FbsApiGenerateUnitedOrdersReportRequest {
func (r FbsApiGenerateUnitedOrdersReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateUnitedOrdersReportRequest
func (a *FbsAPIService) GenerateUnitedOrdersReport(ctx context.Context) FbsApiGenerateUnitedOrdersReportRequest {
	return FbsApiGenerateUnitedOrdersReportRequest{
func (a *FbsAPIService) GenerateUnitedOrdersReportExecute(r FbsApiGenerateUnitedOrdersReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGenerateUnitedReturnsReportRequest struct {
func (r FbsApiGenerateUnitedReturnsReportRequest) GenerateUnitedReturnsRequest(generateUnitedReturnsRequest GenerateUnitedReturnsRequest) FbsApiGenerateUnitedReturnsReportRequest {
func (r FbsApiGenerateUnitedReturnsReportRequest) Format(format ReportFormatType) FbsApiGenerateUnitedReturnsReportRequest {
func (r FbsApiGenerateUnitedReturnsReportRequest) Execute() (*GenerateReportResponse, *http.Response, error) {
	@return FbsApiGenerateUnitedReturnsReportRequest
func (a *FbsAPIService) GenerateUnitedReturnsReport(ctx context.Context) FbsApiGenerateUnitedReturnsReportRequest {
	return FbsApiGenerateUnitedReturnsReportRequest{
func (a *FbsAPIService) GenerateUnitedReturnsReportExecute(r FbsApiGenerateUnitedReturnsReportRequest) (*GenerateReportResponse, *http.Response, error) {
type FbsApiGetAuthTokenInfoRequest struct {
func (r FbsApiGetAuthTokenInfoRequest) Execute() (*GetTokenInfoResponse, *http.Response, error) {
	@return FbsApiGetAuthTokenInfoRequest
func (a *FbsAPIService) GetAuthTokenInfo(ctx context.Context) FbsApiGetAuthTokenInfoRequest {
	return FbsApiGetAuthTokenInfoRequest{
func (a *FbsAPIService) GetAuthTokenInfoExecute(r FbsApiGetAuthTokenInfoRequest) (*GetTokenInfoResponse, *http.Response, error) {
type FbsApiGetBidsInfoForBusinessRequest struct {
func (r FbsApiGetBidsInfoForBusinessRequest) PageToken(pageToken string) FbsApiGetBidsInfoForBusinessRequest {
func (r FbsApiGetBidsInfoForBusinessRequest) Limit(limit int32) FbsApiGetBidsInfoForBusinessRequest {
func (r FbsApiGetBidsInfoForBusinessRequest) GetBidsInfoRequest(getBidsInfoRequest GetBidsInfoRequest) FbsApiGetBidsInfoForBusinessRequest {
func (r FbsApiGetBidsInfoForBusinessRequest) Execute() (*GetBidsInfoResponse, *http.Response, error) {
	@return FbsApiGetBidsInfoForBusinessRequest
func (a *FbsAPIService) GetBidsInfoForBusiness(ctx context.Context, businessId int64) FbsApiGetBidsInfoForBusinessRequest {
	return FbsApiGetBidsInfoForBusinessRequest{
func (a *FbsAPIService) GetBidsInfoForBusinessExecute(r FbsApiGetBidsInfoForBusinessRequest) (*GetBidsInfoResponse, *http.Response, error) {
type FbsApiGetBidsRecommendationsRequest struct {
func (r FbsApiGetBidsRecommendationsRequest) GetBidsRecommendationsRequest(getBidsRecommendationsRequest GetBidsRecommendationsRequest) FbsApiGetBidsRecommendationsRequest {
func (r FbsApiGetBidsRecommendationsRequest) Execute() (*GetBidsRecommendationsResponse, *http.Response, error) {
	@return FbsApiGetBidsRecommendationsRequest
func (a *FbsAPIService) GetBidsRecommendations(ctx context.Context, businessId int64) FbsApiGetBidsRecommendationsRequest {
	return FbsApiGetBidsRecommendationsRequest{
func (a *FbsAPIService) GetBidsRecommendationsExecute(r FbsApiGetBidsRecommendationsRequest) (*GetBidsRecommendationsResponse, *http.Response, error) {
type FbsApiGetBusinessQuarantineOffersRequest struct {
func (r FbsApiGetBusinessQuarantineOffersRequest) GetQuarantineOffersRequest(getQuarantineOffersRequest GetQuarantineOffersRequest) FbsApiGetBusinessQuarantineOffersRequest {
func (r FbsApiGetBusinessQuarantineOffersRequest) PageToken(pageToken string) FbsApiGetBusinessQuarantineOffersRequest {
func (r FbsApiGetBusinessQuarantineOffersRequest) Limit(limit int32) FbsApiGetBusinessQuarantineOffersRequest {
func (r FbsApiGetBusinessQuarantineOffersRequest) Execute() (*GetQuarantineOffersResponse, *http.Response, error) {
	@return FbsApiGetBusinessQuarantineOffersRequest
func (a *FbsAPIService) GetBusinessQuarantineOffers(ctx context.Context, businessId int64) FbsApiGetBusinessQuarantineOffersRequest {
	return FbsApiGetBusinessQuarantineOffersRequest{
func (a *FbsAPIService) GetBusinessQuarantineOffersExecute(r FbsApiGetBusinessQuarantineOffersRequest) (*GetQuarantineOffersResponse, *http.Response, error) {
type FbsApiGetBusinessSettingsRequest struct {
func (r FbsApiGetBusinessSettingsRequest) Execute() (*GetBusinessSettingsResponse, *http.Response, error) {
	@return FbsApiGetBusinessSettingsRequest
func (a *FbsAPIService) GetBusinessSettings(ctx context.Context, businessId int64) FbsApiGetBusinessSettingsRequest {
	return FbsApiGetBusinessSettingsRequest{
func (a *FbsAPIService) GetBusinessSettingsExecute(r FbsApiGetBusinessSettingsRequest) (*GetBusinessSettingsResponse, *http.Response, error) {
type FbsApiGetCampaignRequest struct {
func (r FbsApiGetCampaignRequest) Execute() (*GetCampaignResponse, *http.Response, error) {
	@return FbsApiGetCampaignRequest
func (a *FbsAPIService) GetCampaign(ctx context.Context, campaignId int64) FbsApiGetCampaignRequest {
	return FbsApiGetCampaignRequest{
func (a *FbsAPIService) GetCampaignExecute(r FbsApiGetCampaignRequest) (*GetCampaignResponse, *http.Response, error) {
type FbsApiGetCampaignOffersRequest struct {
func (r FbsApiGetCampaignOffersRequest) GetCampaignOffersRequest(getCampaignOffersRequest GetCampaignOffersRequest) FbsApiGetCampaignOffersRequest {
func (r FbsApiGetCampaignOffersRequest) PageToken(pageToken string) FbsApiGetCampaignOffersRequest {
func (r FbsApiGetCampaignOffersRequest) Limit(limit int32) FbsApiGetCampaignOffersRequest {
func (r FbsApiGetCampaignOffersRequest) Execute() (*GetCampaignOffersResponse, *http.Response, error) {
	@return FbsApiGetCampaignOffersRequest
func (a *FbsAPIService) GetCampaignOffers(ctx context.Context, campaignId int64) FbsApiGetCampaignOffersRequest {
	return FbsApiGetCampaignOffersRequest{
func (a *FbsAPIService) GetCampaignOffersExecute(r FbsApiGetCampaignOffersRequest) (*GetCampaignOffersResponse, *http.Response, error) {
type FbsApiGetCampaignQuarantineOffersRequest struct {
func (r FbsApiGetCampaignQuarantineOffersRequest) GetQuarantineOffersRequest(getQuarantineOffersRequest GetQuarantineOffersRequest) FbsApiGetCampaignQuarantineOffersRequest {
func (r FbsApiGetCampaignQuarantineOffersRequest) PageToken(pageToken string) FbsApiGetCampaignQuarantineOffersRequest {
func (r FbsApiGetCampaignQuarantineOffersRequest) Limit(limit int32) FbsApiGetCampaignQuarantineOffersRequest {
func (r FbsApiGetCampaignQuarantineOffersRequest) Execute() (*GetQuarantineOffersResponse, *http.Response, error) {
	@return FbsApiGetCampaignQuarantineOffersRequest
func (a *FbsAPIService) GetCampaignQuarantineOffers(ctx context.Context, campaignId int64) FbsApiGetCampaignQuarantineOffersRequest {
	return FbsApiGetCampaignQuarantineOffersRequest{
func (a *FbsAPIService) GetCampaignQuarantineOffersExecute(r FbsApiGetCampaignQuarantineOffersRequest) (*GetQuarantineOffersResponse, *http.Response, error) {
type FbsApiGetCampaignRegionRequest struct {
func (r FbsApiGetCampaignRegionRequest) Execute() (*GetCampaignRegionResponse, *http.Response, error) {
	@return FbsApiGetCampaignRegionRequest
func (a *FbsAPIService) GetCampaignRegion(ctx context.Context, campaignId int64) FbsApiGetCampaignRegionRequest {
	return FbsApiGetCampaignRegionRequest{
func (a *FbsAPIService) GetCampaignRegionExecute(r FbsApiGetCampaignRegionRequest) (*GetCampaignRegionResponse, *http.Response, error) {
type FbsApiGetCampaignSettingsRequest struct {
func (r FbsApiGetCampaignSettingsRequest) Execute() (*GetCampaignSettingsResponse, *http.Response, error) {
	@return FbsApiGetCampaignSettingsRequest
func (a *FbsAPIService) GetCampaignSettings(ctx context.Context, campaignId int64) FbsApiGetCampaignSettingsRequest {
	return FbsApiGetCampaignSettingsRequest{
func (a *FbsAPIService) GetCampaignSettingsExecute(r FbsApiGetCampaignSettingsRequest) (*GetCampaignSettingsResponse, *http.Response, error) {
type FbsApiGetCampaignsRequest struct {
func (r FbsApiGetCampaignsRequest) Page(page int32) FbsApiGetCampaignsRequest {
func (r FbsApiGetCampaignsRequest) PageSize(pageSize int32) FbsApiGetCampaignsRequest {
func (r FbsApiGetCampaignsRequest) Execute() (*GetCampaignsResponse, *http.Response, error) {
	@return FbsApiGetCampaignsRequest
func (a *FbsAPIService) GetCampaigns(ctx context.Context) FbsApiGetCampaignsRequest {
	return FbsApiGetCampaignsRequest{
func (a *FbsAPIService) GetCampaignsExecute(r FbsApiGetCampaignsRequest) (*GetCampaignsResponse, *http.Response, error) {
type FbsApiGetCategoriesMaxSaleQuantumRequest struct {
func (r FbsApiGetCategoriesMaxSaleQuantumRequest) GetCategoriesMaxSaleQuantumRequest(getCategoriesMaxSaleQuantumRequest GetCategoriesMaxSaleQuantumRequest) FbsApiGetCategoriesMaxSaleQuantumRequest {
func (r FbsApiGetCategoriesMaxSaleQuantumRequest) Execute() (*GetCategoriesMaxSaleQuantumResponse, *http.Response, error) {
	@return FbsApiGetCategoriesMaxSaleQuantumRequest
func (a *FbsAPIService) GetCategoriesMaxSaleQuantum(ctx context.Context) FbsApiGetCategoriesMaxSaleQuantumRequest {
	return FbsApiGetCategoriesMaxSaleQuantumRequest{
func (a *FbsAPIService) GetCategoriesMaxSaleQuantumExecute(r FbsApiGetCategoriesMaxSaleQuantumRequest) (*GetCategoriesMaxSaleQuantumResponse, *http.Response, error) {
type FbsApiGetCategoriesTreeRequest struct {
func (r FbsApiGetCategoriesTreeRequest) GetCategoriesRequest(getCategoriesRequest GetCategoriesRequest) FbsApiGetCategoriesTreeRequest {
func (r FbsApiGetCategoriesTreeRequest) Execute() (*GetCategoriesResponse, *http.Response, error) {
	@return FbsApiGetCategoriesTreeRequest
func (a *FbsAPIService) GetCategoriesTree(ctx context.Context) FbsApiGetCategoriesTreeRequest {
	return FbsApiGetCategoriesTreeRequest{
func (a *FbsAPIService) GetCategoriesTreeExecute(r FbsApiGetCategoriesTreeRequest) (*GetCategoriesResponse, *http.Response, error) {
type FbsApiGetCategoryContentParametersRequest struct {
func (r FbsApiGetCategoryContentParametersRequest) Execute() (*GetCategoryContentParametersResponse, *http.Response, error) {
	@return FbsApiGetCategoryContentParametersRequest
func (a *FbsAPIService) GetCategoryContentParameters(ctx context.Context, categoryId int64) FbsApiGetCategoryContentParametersRequest {
	return FbsApiGetCategoryContentParametersRequest{
func (a *FbsAPIService) GetCategoryContentParametersExecute(r FbsApiGetCategoryContentParametersRequest) (*GetCategoryContentParametersResponse, *http.Response, error) {
type FbsApiGetChatRequest struct {
func (r FbsApiGetChatRequest) ChatId(chatId int64) FbsApiGetChatRequest {
func (r FbsApiGetChatRequest) Execute() (*GetChatResponse, *http.Response, error) {
	@return FbsApiGetChatRequest
func (a *FbsAPIService) GetChat(ctx context.Context, businessId int64) FbsApiGetChatRequest {
	return FbsApiGetChatRequest{
func (a *FbsAPIService) GetChatExecute(r FbsApiGetChatRequest) (*GetChatResponse, *http.Response, error) {
type FbsApiGetChatHistoryRequest struct {
func (r FbsApiGetChatHistoryRequest) ChatId(chatId int64) FbsApiGetChatHistoryRequest {
func (r FbsApiGetChatHistoryRequest) GetChatHistoryRequest(getChatHistoryRequest GetChatHistoryRequest) FbsApiGetChatHistoryRequest {
func (r FbsApiGetChatHistoryRequest) PageToken(pageToken string) FbsApiGetChatHistoryRequest {
func (r FbsApiGetChatHistoryRequest) Limit(limit int32) FbsApiGetChatHistoryRequest {
func (r FbsApiGetChatHistoryRequest) Execute() (*GetChatHistoryResponse, *http.Response, error) {
	@return FbsApiGetChatHistoryRequest
func (a *FbsAPIService) GetChatHistory(ctx context.Context, businessId int64) FbsApiGetChatHistoryRequest {
	return FbsApiGetChatHistoryRequest{
func (a *FbsAPIService) GetChatHistoryExecute(r FbsApiGetChatHistoryRequest) (*GetChatHistoryResponse, *http.Response, error) {
type FbsApiGetChatMessageRequest struct {
func (r FbsApiGetChatMessageRequest) ChatId(chatId int64) FbsApiGetChatMessageRequest {
func (r FbsApiGetChatMessageRequest) MessageId(messageId int64) FbsApiGetChatMessageRequest {
func (r FbsApiGetChatMessageRequest) Execute() (*GetChatMessageResponse, *http.Response, error) {
	@return FbsApiGetChatMessageRequest
func (a *FbsAPIService) GetChatMessage(ctx context.Context, businessId int64) FbsApiGetChatMessageRequest {
	return FbsApiGetChatMessageRequest{
func (a *FbsAPIService) GetChatMessageExecute(r FbsApiGetChatMessageRequest) (*GetChatMessageResponse, *http.Response, error) {
type FbsApiGetChatsRequest struct {
func (r FbsApiGetChatsRequest) GetChatsRequest(getChatsRequest GetChatsRequest) FbsApiGetChatsRequest {
func (r FbsApiGetChatsRequest) PageToken(pageToken string) FbsApiGetChatsRequest {
func (r FbsApiGetChatsRequest) Limit(limit int32) FbsApiGetChatsRequest {
func (r FbsApiGetChatsRequest) Execute() (*GetChatsResponse, *http.Response, error) {
	@return FbsApiGetChatsRequest
func (a *FbsAPIService) GetChats(ctx context.Context, businessId int64) FbsApiGetChatsRequest {
	return FbsApiGetChatsRequest{
func (a *FbsAPIService) GetChatsExecute(r FbsApiGetChatsRequest) (*GetChatsResponse, *http.Response, error) {
type FbsApiGetDeliveryServicesRequest struct {
func (r FbsApiGetDeliveryServicesRequest) Execute() (*GetDeliveryServicesResponse, *http.Response, error) {
	@return FbsApiGetDeliveryServicesRequest
func (a *FbsAPIService) GetDeliveryServices(ctx context.Context) FbsApiGetDeliveryServicesRequest {
	return FbsApiGetDeliveryServicesRequest{
func (a *FbsAPIService) GetDeliveryServicesExecute(r FbsApiGetDeliveryServicesRequest) (*GetDeliveryServicesResponse, *http.Response, error) {
type FbsApiGetGoodsFeedbackCommentsRequest struct {
func (r FbsApiGetGoodsFeedbackCommentsRequest) GetGoodsFeedbackCommentsRequest(getGoodsFeedbackCommentsRequest GetGoodsFeedbackCommentsRequest) FbsApiGetGoodsFeedbackCommentsRequest {
func (r FbsApiGetGoodsFeedbackCommentsRequest) PageToken(pageToken string) FbsApiGetGoodsFeedbackCommentsRequest {
func (r FbsApiGetGoodsFeedbackCommentsRequest) Limit(limit int32) FbsApiGetGoodsFeedbackCommentsRequest {
func (r FbsApiGetGoodsFeedbackCommentsRequest) Execute() (*GetGoodsFeedbackCommentsResponse, *http.Response, error) {
	@return FbsApiGetGoodsFeedbackCommentsRequest
func (a *FbsAPIService) GetGoodsFeedbackComments(ctx context.Context, businessId int64) FbsApiGetGoodsFeedbackCommentsRequest {
	return FbsApiGetGoodsFeedbackCommentsRequest{
func (a *FbsAPIService) GetGoodsFeedbackCommentsExecute(r FbsApiGetGoodsFeedbackCommentsRequest) (*GetGoodsFeedbackCommentsResponse, *http.Response, error) {
type FbsApiGetGoodsFeedbacksRequest struct {
func (r FbsApiGetGoodsFeedbacksRequest) PageToken(pageToken string) FbsApiGetGoodsFeedbacksRequest {
func (r FbsApiGetGoodsFeedbacksRequest) Limit(limit int32) FbsApiGetGoodsFeedbacksRequest {
func (r FbsApiGetGoodsFeedbacksRequest) GetGoodsFeedbackRequest(getGoodsFeedbackRequest GetGoodsFeedbackRequest) FbsApiGetGoodsFeedbacksRequest {
func (r FbsApiGetGoodsFeedbacksRequest) Execute() (*GetGoodsFeedbackResponse, *http.Response, error) {
	@return FbsApiGetGoodsFeedbacksRequest
func (a *FbsAPIService) GetGoodsFeedbacks(ctx context.Context, businessId int64) FbsApiGetGoodsFeedbacksRequest {
	return FbsApiGetGoodsFeedbacksRequest{
func (a *FbsAPIService) GetGoodsFeedbacksExecute(r FbsApiGetGoodsFeedbacksRequest) (*GetGoodsFeedbackResponse, *http.Response, error) {
type FbsApiGetGoodsStatsRequest struct {
func (r FbsApiGetGoodsStatsRequest) GetGoodsStatsRequest(getGoodsStatsRequest GetGoodsStatsRequest) FbsApiGetGoodsStatsRequest {
func (r FbsApiGetGoodsStatsRequest) Execute() (*GetGoodsStatsResponse, *http.Response, error) {
	@return FbsApiGetGoodsStatsRequest
func (a *FbsAPIService) GetGoodsStats(ctx context.Context, campaignId int64) FbsApiGetGoodsStatsRequest {
	return FbsApiGetGoodsStatsRequest{
func (a *FbsAPIService) GetGoodsStatsExecute(r FbsApiGetGoodsStatsRequest) (*GetGoodsStatsResponse, *http.Response, error) {
type FbsApiGetHiddenOffersRequest struct {
func (r FbsApiGetHiddenOffersRequest) OfferId(offerId []string) FbsApiGetHiddenOffersRequest {
func (r FbsApiGetHiddenOffersRequest) PageToken(pageToken string) FbsApiGetHiddenOffersRequest {
func (r FbsApiGetHiddenOffersRequest) Limit(limit int32) FbsApiGetHiddenOffersRequest {
func (r FbsApiGetHiddenOffersRequest) Execute() (*GetHiddenOffersResponse, *http.Response, error) {
	@return FbsApiGetHiddenOffersRequest
func (a *FbsAPIService) GetHiddenOffers(ctx context.Context, campaignId int64) FbsApiGetHiddenOffersRequest {
	return FbsApiGetHiddenOffersRequest{
func (a *FbsAPIService) GetHiddenOffersExecute(r FbsApiGetHiddenOffersRequest) (*GetHiddenOffersResponse, *http.Response, error) {
type FbsApiGetOfferCardsContentStatusRequest struct {
func (r FbsApiGetOfferCardsContentStatusRequest) PageToken(pageToken string) FbsApiGetOfferCardsContentStatusRequest {
func (r FbsApiGetOfferCardsContentStatusRequest) Limit(limit int32) FbsApiGetOfferCardsContentStatusRequest {
func (r FbsApiGetOfferCardsContentStatusRequest) GetOfferCardsContentStatusRequest(getOfferCardsContentStatusRequest GetOfferCardsContentStatusRequest) FbsApiGetOfferCardsContentStatusRequest {
func (r FbsApiGetOfferCardsContentStatusRequest) Execute() (*GetOfferCardsContentStatusResponse, *http.Response, error) {
	@return FbsApiGetOfferCardsContentStatusRequest
func (a *FbsAPIService) GetOfferCardsContentStatus(ctx context.Context, businessId int64) FbsApiGetOfferCardsContentStatusRequest {
	return FbsApiGetOfferCardsContentStatusRequest{
func (a *FbsAPIService) GetOfferCardsContentStatusExecute(r FbsApiGetOfferCardsContentStatusRequest) (*GetOfferCardsContentStatusResponse, *http.Response, error) {
type FbsApiGetOfferMappingEntriesRequest struct {
func (r FbsApiGetOfferMappingEntriesRequest) OfferId(offerId []string) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) ShopSku(shopSku []string) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) MappingKind(mappingKind OfferMappingKindType) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) Status(status []OfferProcessingStatusType) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) Availability(availability []OfferAvailabilityStatusType) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) CategoryId(categoryId []int32) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) Vendor(vendor []string) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) PageToken(pageToken string) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) Limit(limit int32) FbsApiGetOfferMappingEntriesRequest {
func (r FbsApiGetOfferMappingEntriesRequest) Execute() (*GetOfferMappingEntriesResponse, *http.Response, error) {
	@return FbsApiGetOfferMappingEntriesRequest
func (a *FbsAPIService) GetOfferMappingEntries(ctx context.Context, campaignId int64) FbsApiGetOfferMappingEntriesRequest {
	return FbsApiGetOfferMappingEntriesRequest{
func (a *FbsAPIService) GetOfferMappingEntriesExecute(r FbsApiGetOfferMappingEntriesRequest) (*GetOfferMappingEntriesResponse, *http.Response, error) {
type FbsApiGetOfferMappingsRequest struct {
func (r FbsApiGetOfferMappingsRequest) PageToken(pageToken string) FbsApiGetOfferMappingsRequest {
func (r FbsApiGetOfferMappingsRequest) Limit(limit int32) FbsApiGetOfferMappingsRequest {
func (r FbsApiGetOfferMappingsRequest) Language(language CatalogLanguageType) FbsApiGetOfferMappingsRequest {
func (r FbsApiGetOfferMappingsRequest) GetOfferMappingsRequest(getOfferMappingsRequest GetOfferMappingsRequest) FbsApiGetOfferMappingsRequest {
func (r FbsApiGetOfferMappingsRequest) Execute() (*GetOfferMappingsResponse, *http.Response, error) {
	@return FbsApiGetOfferMappingsRequest
func (a *FbsAPIService) GetOfferMappings(ctx context.Context, businessId int64) FbsApiGetOfferMappingsRequest {
	return FbsApiGetOfferMappingsRequest{
func (a *FbsAPIService) GetOfferMappingsExecute(r FbsApiGetOfferMappingsRequest) (*GetOfferMappingsResponse, *http.Response, error) {
type FbsApiGetOfferRecommendationsRequest struct {
func (r FbsApiGetOfferRecommendationsRequest) GetOfferRecommendationsRequest(getOfferRecommendationsRequest GetOfferRecommendationsRequest) FbsApiGetOfferRecommendationsRequest {
func (r FbsApiGetOfferRecommendationsRequest) PageToken(pageToken string) FbsApiGetOfferRecommendationsRequest {
func (r FbsApiGetOfferRecommendationsRequest) Limit(limit int32) FbsApiGetOfferRecommendationsRequest {
func (r FbsApiGetOfferRecommendationsRequest) Execute() (*GetOfferRecommendationsResponse, *http.Response, error) {
	@return FbsApiGetOfferRecommendationsRequest
func (a *FbsAPIService) GetOfferRecommendations(ctx context.Context, businessId int64) FbsApiGetOfferRecommendationsRequest {
	return FbsApiGetOfferRecommendationsRequest{
func (a *FbsAPIService) GetOfferRecommendationsExecute(r FbsApiGetOfferRecommendationsRequest) (*GetOfferRecommendationsResponse, *http.Response, error) {
type FbsApiGetOrderRequest struct {
func (r FbsApiGetOrderRequest) Execute() (*GetOrderResponse, *http.Response, error) {
	@return FbsApiGetOrderRequest
func (a *FbsAPIService) GetOrder(ctx context.Context, campaignId int64, orderId int64) FbsApiGetOrderRequest {
	return FbsApiGetOrderRequest{
func (a *FbsAPIService) GetOrderExecute(r FbsApiGetOrderRequest) (*GetOrderResponse, *http.Response, error) {
type FbsApiGetOrderBusinessBuyerInfoRequest struct {
func (r FbsApiGetOrderBusinessBuyerInfoRequest) Execute() (*GetBusinessBuyerInfoResponse, *http.Response, error) {
	@return FbsApiGetOrderBusinessBuyerInfoRequest
func (a *FbsAPIService) GetOrderBusinessBuyerInfo(ctx context.Context, campaignId int64, orderId int64) FbsApiGetOrderBusinessBuyerInfoRequest {
	return FbsApiGetOrderBusinessBuyerInfoRequest{
func (a *FbsAPIService) GetOrderBusinessBuyerInfoExecute(r FbsApiGetOrderBusinessBuyerInfoRequest) (*GetBusinessBuyerInfoResponse, *http.Response, error) {
type FbsApiGetOrderBusinessDocumentsInfoRequest struct {
func (r FbsApiGetOrderBusinessDocumentsInfoRequest) Execute() (*GetBusinessDocumentsInfoResponse, *http.Response, error) {
	@return FbsApiGetOrderBusinessDocumentsInfoRequest
func (a *FbsAPIService) GetOrderBusinessDocumentsInfo(ctx context.Context, campaignId int64, orderId int64) FbsApiGetOrderBusinessDocumentsInfoRequest {
	return FbsApiGetOrderBusinessDocumentsInfoRequest{
func (a *FbsAPIService) GetOrderBusinessDocumentsInfoExecute(r FbsApiGetOrderBusinessDocumentsInfoRequest) (*GetBusinessDocumentsInfoResponse, *http.Response, error) {
type FbsApiGetOrderIdentifiersStatusRequest struct {
func (r FbsApiGetOrderIdentifiersStatusRequest) Execute() (*GetOrderIdentifiersStatusResponse, *http.Response, error) {
	@return FbsApiGetOrderIdentifiersStatusRequest
func (a *FbsAPIService) GetOrderIdentifiersStatus(ctx context.Context, campaignId int64, orderId int64) FbsApiGetOrderIdentifiersStatusRequest {
	return FbsApiGetOrderIdentifiersStatusRequest{
func (a *FbsAPIService) GetOrderIdentifiersStatusExecute(r FbsApiGetOrderIdentifiersStatusRequest) (*GetOrderIdentifiersStatusResponse, *http.Response, error) {
type FbsApiGetOrderLabelsDataRequest struct {
func (r FbsApiGetOrderLabelsDataRequest) Execute() (*GetOrderLabelsDataResponse, *http.Response, error) {
	@return FbsApiGetOrderLabelsDataRequest
func (a *FbsAPIService) GetOrderLabelsData(ctx context.Context, campaignId int64, orderId int64) FbsApiGetOrderLabelsDataRequest {
	return FbsApiGetOrderLabelsDataRequest{
func (a *FbsAPIService) GetOrderLabelsDataExecute(r FbsApiGetOrderLabelsDataRequest) (*GetOrderLabelsDataResponse, *http.Response, error) {
type FbsApiGetOrdersRequest struct {
func (r FbsApiGetOrdersRequest) OrderIds(orderIds []int64) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) Status(status []OrderStatusType) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) Substatus(substatus []OrderSubstatusType) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) FromDate(fromDate string) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) ToDate(toDate string) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) SupplierShipmentDateFrom(supplierShipmentDateFrom string) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) SupplierShipmentDateTo(supplierShipmentDateTo string) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) UpdatedAtFrom(updatedAtFrom time.Time) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) UpdatedAtTo(updatedAtTo time.Time) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) DispatchType(dispatchType OrderDeliveryDispatchType) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) Fake(fake bool) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) HasCis(hasCis bool) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) OnlyWaitingForCancellationApprove(onlyWaitingForCancellationApprove bool) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) OnlyEstimatedDelivery(onlyEstimatedDelivery bool) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) BuyerType(buyerType OrderBuyerType) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) Page(page int32) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) PageSize(pageSize int32) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) PageToken(pageToken string) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) Limit(limit int32) FbsApiGetOrdersRequest {
func (r FbsApiGetOrdersRequest) Execute() (*GetOrdersResponse, *http.Response, error) {
	@return FbsApiGetOrdersRequest
func (a *FbsAPIService) GetOrders(ctx context.Context, campaignId int64) FbsApiGetOrdersRequest {
	return FbsApiGetOrdersRequest{
func (a *FbsAPIService) GetOrdersExecute(r FbsApiGetOrdersRequest) (*GetOrdersResponse, *http.Response, error) {
type FbsApiGetOrdersStatsRequest struct {
func (r FbsApiGetOrdersStatsRequest) PageToken(pageToken string) FbsApiGetOrdersStatsRequest {
func (r FbsApiGetOrdersStatsRequest) Limit(limit int32) FbsApiGetOrdersStatsRequest {
func (r FbsApiGetOrdersStatsRequest) GetOrdersStatsRequest(getOrdersStatsRequest GetOrdersStatsRequest) FbsApiGetOrdersStatsRequest {
func (r FbsApiGetOrdersStatsRequest) Execute() (*GetOrdersStatsResponse, *http.Response, error) {
	@return FbsApiGetOrdersStatsRequest
func (a *FbsAPIService) GetOrdersStats(ctx context.Context, campaignId int64) FbsApiGetOrdersStatsRequest {
	return FbsApiGetOrdersStatsRequest{
func (a *FbsAPIService) GetOrdersStatsExecute(r FbsApiGetOrdersStatsRequest) (*GetOrdersStatsResponse, *http.Response, error) {
type FbsApiGetPagedWarehousesRequest struct {
func (r FbsApiGetPagedWarehousesRequest) PageToken(pageToken string) FbsApiGetPagedWarehousesRequest {
func (r FbsApiGetPagedWarehousesRequest) Limit(limit int32) FbsApiGetPagedWarehousesRequest {
func (r FbsApiGetPagedWarehousesRequest) GetPagedWarehousesRequest(getPagedWarehousesRequest GetPagedWarehousesRequest) FbsApiGetPagedWarehousesRequest {
func (r FbsApiGetPagedWarehousesRequest) Execute() (*GetPagedWarehousesResponse, *http.Response, error) {
	@return FbsApiGetPagedWarehousesRequest
func (a *FbsAPIService) GetPagedWarehouses(ctx context.Context, businessId int64) FbsApiGetPagedWarehousesRequest {
	return FbsApiGetPagedWarehousesRequest{
func (a *FbsAPIService) GetPagedWarehousesExecute(r FbsApiGetPagedWarehousesRequest) (*GetPagedWarehousesResponse, *http.Response, error) {
type FbsApiGetPricesRequest struct {
func (r FbsApiGetPricesRequest) PageToken(pageToken string) FbsApiGetPricesRequest {
func (r FbsApiGetPricesRequest) Limit(limit int32) FbsApiGetPricesRequest {
func (r FbsApiGetPricesRequest) Archived(archived bool) FbsApiGetPricesRequest {
func (r FbsApiGetPricesRequest) Execute() (*GetPricesResponse, *http.Response, error) {
	@return FbsApiGetPricesRequest
func (a *FbsAPIService) GetPrices(ctx context.Context, campaignId int64) FbsApiGetPricesRequest {
	return FbsApiGetPricesRequest{
func (a *FbsAPIService) GetPricesExecute(r FbsApiGetPricesRequest) (*GetPricesResponse, *http.Response, error) {
type FbsApiGetPricesByOfferIdsRequest struct {
func (r FbsApiGetPricesByOfferIdsRequest) PageToken(pageToken string) FbsApiGetPricesByOfferIdsRequest {
func (r FbsApiGetPricesByOfferIdsRequest) Limit(limit int32) FbsApiGetPricesByOfferIdsRequest {
func (r FbsApiGetPricesByOfferIdsRequest) GetPricesByOfferIdsRequest(getPricesByOfferIdsRequest GetPricesByOfferIdsRequest) FbsApiGetPricesByOfferIdsRequest {
func (r FbsApiGetPricesByOfferIdsRequest) Execute() (*GetPricesByOfferIdsResponse, *http.Response, error) {
	@return FbsApiGetPricesByOfferIdsRequest
func (a *FbsAPIService) GetPricesByOfferIds(ctx context.Context, campaignId int64) FbsApiGetPricesByOfferIdsRequest {
	return FbsApiGetPricesByOfferIdsRequest{
func (a *FbsAPIService) GetPricesByOfferIdsExecute(r FbsApiGetPricesByOfferIdsRequest) (*GetPricesByOfferIdsResponse, *http.Response, error) {
type FbsApiGetPromoOffersRequest struct {
func (r FbsApiGetPromoOffersRequest) GetPromoOffersRequest(getPromoOffersRequest GetPromoOffersRequest) FbsApiGetPromoOffersRequest {
func (r FbsApiGetPromoOffersRequest) PageToken(pageToken string) FbsApiGetPromoOffersRequest {
func (r FbsApiGetPromoOffersRequest) Limit(limit int32) FbsApiGetPromoOffersRequest {
func (r FbsApiGetPromoOffersRequest) Execute() (*GetPromoOffersResponse, *http.Response, error) {
	@return FbsApiGetPromoOffersRequest
func (a *FbsAPIService) GetPromoOffers(ctx context.Context, businessId int64) FbsApiGetPromoOffersRequest {
	return FbsApiGetPromoOffersRequest{
func (a *FbsAPIService) GetPromoOffersExecute(r FbsApiGetPromoOffersRequest) (*GetPromoOffersResponse, *http.Response, error) {
type FbsApiGetPromosRequest struct {
func (r FbsApiGetPromosRequest) GetPromosRequest(getPromosRequest GetPromosRequest) FbsApiGetPromosRequest {
func (r FbsApiGetPromosRequest) Execute() (*GetPromosResponse, *http.Response, error) {
	@return FbsApiGetPromosRequest
func (a *FbsAPIService) GetPromos(ctx context.Context, businessId int64) FbsApiGetPromosRequest {
	return FbsApiGetPromosRequest{
func (a *FbsAPIService) GetPromosExecute(r FbsApiGetPromosRequest) (*GetPromosResponse, *http.Response, error) {
type FbsApiGetQualityRatingDetailsRequest struct {
func (r FbsApiGetQualityRatingDetailsRequest) Execute() (*GetQualityRatingDetailsResponse, *http.Response, error) {
	@return FbsApiGetQualityRatingDetailsRequest
func (a *FbsAPIService) GetQualityRatingDetails(ctx context.Context, campaignId int64) FbsApiGetQualityRatingDetailsRequest {
	return FbsApiGetQualityRatingDetailsRequest{
func (a *FbsAPIService) GetQualityRatingDetailsExecute(r FbsApiGetQualityRatingDetailsRequest) (*GetQualityRatingDetailsResponse, *http.Response, error) {
type FbsApiGetQualityRatingsRequest struct {
func (r FbsApiGetQualityRatingsRequest) GetQualityRatingRequest(getQualityRatingRequest GetQualityRatingRequest) FbsApiGetQualityRatingsRequest {
func (r FbsApiGetQualityRatingsRequest) Execute() (*GetQualityRatingResponse, *http.Response, error) {
	@return FbsApiGetQualityRatingsRequest
func (a *FbsAPIService) GetQualityRatings(ctx context.Context, businessId int64) FbsApiGetQualityRatingsRequest {
	return FbsApiGetQualityRatingsRequest{
func (a *FbsAPIService) GetQualityRatingsExecute(r FbsApiGetQualityRatingsRequest) (*GetQualityRatingResponse, *http.Response, error) {
type FbsApiGetRegionsCodesRequest struct {
func (r FbsApiGetRegionsCodesRequest) Execute() (*GetRegionsCodesResponse, *http.Response, error) {
	@return FbsApiGetRegionsCodesRequest
func (a *FbsAPIService) GetRegionsCodes(ctx context.Context) FbsApiGetRegionsCodesRequest {
	return FbsApiGetRegionsCodesRequest{
func (a *FbsAPIService) GetRegionsCodesExecute(r FbsApiGetRegionsCodesRequest) (*GetRegionsCodesResponse, *http.Response, error) {
type FbsApiGetReportInfoRequest struct {
func (r FbsApiGetReportInfoRequest) Execute() (*GetReportInfoResponse, *http.Response, error) {
	@return FbsApiGetReportInfoRequest
func (a *FbsAPIService) GetReportInfo(ctx context.Context, reportId string) FbsApiGetReportInfoRequest {
	return FbsApiGetReportInfoRequest{
func (a *FbsAPIService) GetReportInfoExecute(r FbsApiGetReportInfoRequest) (*GetReportInfoResponse, *http.Response, error) {
type FbsApiGetReturnRequest struct {
func (r FbsApiGetReturnRequest) Execute() (*GetReturnResponse, *http.Response, error) {
	@return FbsApiGetReturnRequest
func (a *FbsAPIService) GetReturn(ctx context.Context, campaignId int64, orderId int64, returnId int64) FbsApiGetReturnRequest {
	return FbsApiGetReturnRequest{
func (a *FbsAPIService) GetReturnExecute(r FbsApiGetReturnRequest) (*GetReturnResponse, *http.Response, error) {
type FbsApiGetReturnApplicationRequest struct {
func (r FbsApiGetReturnApplicationRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiGetReturnApplicationRequest
func (a *FbsAPIService) GetReturnApplication(ctx context.Context, campaignId int64, orderId int64, returnId int64) FbsApiGetReturnApplicationRequest {
	return FbsApiGetReturnApplicationRequest{
func (a *FbsAPIService) GetReturnApplicationExecute(r FbsApiGetReturnApplicationRequest) (*os.File, *http.Response, error) {
type FbsApiGetReturnPhotoRequest struct {
func (r FbsApiGetReturnPhotoRequest) Execute() (*os.File, *http.Response, error) {
	@return FbsApiGetReturnPhotoRequest
func (a *FbsAPIService) GetReturnPhoto(ctx context.Context, campaignId int64, orderId int64, returnId int64, itemId int64, imageHash string) FbsApiGetReturnPhotoRequest {
	return FbsApiGetReturnPhotoRequest{
func (a *FbsAPIService) GetReturnPhotoExecute(r FbsApiGetReturnPhotoRequest) (*os.File, *http.Response, error) {
type FbsApiGetReturnsRequest struct {
func (r FbsApiGetReturnsRequest) PageToken(pageToken string) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) Limit(limit int32) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) OrderIds(orderIds []int64) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) Statuses(statuses []RefundStatusType) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) Type_(type_ ReturnType) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) FromDate(fromDate string) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) ToDate(toDate string) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) FromDate2(fromDate2 string) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) ToDate2(toDate2 string) FbsApiGetReturnsRequest {
func (r FbsApiGetReturnsRequest) Execute() (*GetReturnsResponse, *http.Response, error) {
	@return FbsApiGetReturnsRequest
func (a *FbsAPIService) GetReturns(ctx context.Context, campaignId int64) FbsApiGetReturnsRequest {
	return FbsApiGetReturnsRequest{
func (a *FbsAPIService) GetReturnsExecute(r FbsApiGetReturnsRequest) (*GetReturnsResponse, *http.Response, error) {
type FbsApiGetShipmentRequest struct {
func (r FbsApiGetShipmentRequest) CancelledOrders(cancelledOrders bool) FbsApiGetShipmentRequest {
func (r FbsApiGetShipmentRequest) Execute() (*GetShipmentResponse, *http.Response, error) {
	@return FbsApiGetShipmentRequest
func (a *FbsAPIService) GetShipment(ctx context.Context, campaignId int64, shipmentId int64) FbsApiGetShipmentRequest {
	return FbsApiGetShipmentRequest{
func (a *FbsAPIService) GetShipmentExecute(r FbsApiGetShipmentRequest) (*GetShipmentResponse, *http.Response, error) {
type FbsApiGetShipmentOrdersInfoRequest struct {
func (r FbsApiGetShipmentOrdersInfoRequest) Execute() (*GetShipmentOrdersInfoResponse, *http.Response, error) {
	@return FbsApiGetShipmentOrdersInfoRequest
func (a *FbsAPIService) GetShipmentOrdersInfo(ctx context.Context, campaignId int64, shipmentId int64) FbsApiGetShipmentOrdersInfoRequest {
	return FbsApiGetShipmentOrdersInfoRequest{
func (a *FbsAPIService) GetShipmentOrdersInfoExecute(r FbsApiGetShipmentOrdersInfoRequest) (*GetShipmentOrdersInfoResponse, *http.Response, error) {
type FbsApiGetStocksRequest struct {
func (r FbsApiGetStocksRequest) PageToken(pageToken string) FbsApiGetStocksRequest {
func (r FbsApiGetStocksRequest) Limit(limit int32) FbsApiGetStocksRequest {
func (r FbsApiGetStocksRequest) GetWarehouseStocksRequest(getWarehouseStocksRequest GetWarehouseStocksRequest) FbsApiGetStocksRequest {
func (r FbsApiGetStocksRequest) Execute() (*GetWarehouseStocksResponse, *http.Response, error) {
	@return FbsApiGetStocksRequest
func (a *FbsAPIService) GetStocks(ctx context.Context, campaignId int64) FbsApiGetStocksRequest {
	return FbsApiGetStocksRequest{
func (a *FbsAPIService) GetStocksExecute(r FbsApiGetStocksRequest) (*GetWarehouseStocksResponse, *http.Response, error) {
type FbsApiGetSuggestedOfferMappingEntriesRequest struct {
func (r FbsApiGetSuggestedOfferMappingEntriesRequest) GetSuggestedOfferMappingEntriesRequest(getSuggestedOfferMappingEntriesRequest GetSuggestedOfferMappingEntriesRequest) FbsApiGetSuggestedOfferMappingEntriesRequest {
func (r FbsApiGetSuggestedOfferMappingEntriesRequest) Execute() (*GetSuggestedOfferMappingEntriesResponse, *http.Response, error) {
	@return FbsApiGetSuggestedOfferMappingEntriesRequest
func (a *FbsAPIService) GetSuggestedOfferMappingEntries(ctx context.Context, campaignId int64) FbsApiGetSuggestedOfferMappingEntriesRequest {
	return FbsApiGetSuggestedOfferMappingEntriesRequest{
func (a *FbsAPIService) GetSuggestedOfferMappingEntriesExecute(r FbsApiGetSuggestedOfferMappingEntriesRequest) (*GetSuggestedOfferMappingEntriesResponse, *http.Response, error) {
type FbsApiGetSuggestedOfferMappingsRequest struct {
func (r FbsApiGetSuggestedOfferMappingsRequest) GetSuggestedOfferMappingsRequest(getSuggestedOfferMappingsRequest GetSuggestedOfferMappingsRequest) FbsApiGetSuggestedOfferMappingsRequest {
func (r FbsApiGetSuggestedOfferMappingsRequest) Execute() (*GetSuggestedOfferMappingsResponse, *http.Response, error) {
	@return FbsApiGetSuggestedOfferMappingsRequest
func (a *FbsAPIService) GetSuggestedOfferMappings(ctx context.Context, businessId int64) FbsApiGetSuggestedOfferMappingsRequest {
	return FbsApiGetSuggestedOfferMappingsRequest{
func (a *FbsAPIService) GetSuggestedOfferMappingsExecute(r FbsApiGetSuggestedOfferMappingsRequest) (*GetSuggestedOfferMappingsResponse, *http.Response, error) {
type FbsApiGetSuggestedPricesRequest struct {
func (r FbsApiGetSuggestedPricesRequest) SuggestPricesRequest(suggestPricesRequest SuggestPricesRequest) FbsApiGetSuggestedPricesRequest {
func (r FbsApiGetSuggestedPricesRequest) Execute() (*SuggestPricesResponse, *http.Response, error) {
	@return FbsApiGetSuggestedPricesRequest
func (a *FbsAPIService) GetSuggestedPrices(ctx context.Context, campaignId int64) FbsApiGetSuggestedPricesRequest {
	return FbsApiGetSuggestedPricesRequest{
func (a *FbsAPIService) GetSuggestedPricesExecute(r FbsApiGetSuggestedPricesRequest) (*SuggestPricesResponse, *http.Response, error) {
type FbsApiGetWarehousesRequest struct {
func (r FbsApiGetWarehousesRequest) Execute() (*GetWarehousesResponse, *http.Response, error) {
	@return FbsApiGetWarehousesRequest
func (a *FbsAPIService) GetWarehouses(ctx context.Context, businessId int64) FbsApiGetWarehousesRequest {
	return FbsApiGetWarehousesRequest{
func (a *FbsAPIService) GetWarehousesExecute(r FbsApiGetWarehousesRequest) (*GetWarehousesResponse, *http.Response, error) {
type FbsApiProvideOrderItemIdentifiersRequest struct {
func (r FbsApiProvideOrderItemIdentifiersRequest) ProvideOrderItemIdentifiersRequest(provideOrderItemIdentifiersRequest ProvideOrderItemIdentifiersRequest) FbsApiProvideOrderItemIdentifiersRequest {
func (r FbsApiProvideOrderItemIdentifiersRequest) Execute() (*ProvideOrderItemIdentifiersResponse, *http.Response, error) {
	@return FbsApiProvideOrderItemIdentifiersRequest
func (a *FbsAPIService) ProvideOrderItemIdentifiers(ctx context.Context, campaignId int64, orderId int64) FbsApiProvideOrderItemIdentifiersRequest {
	return FbsApiProvideOrderItemIdentifiersRequest{
func (a *FbsAPIService) ProvideOrderItemIdentifiersExecute(r FbsApiProvideOrderItemIdentifiersRequest) (*ProvideOrderItemIdentifiersResponse, *http.Response, error) {
type FbsApiPutBidsForBusinessRequest struct {
func (r FbsApiPutBidsForBusinessRequest) PutSkuBidsRequest(putSkuBidsRequest PutSkuBidsRequest) FbsApiPutBidsForBusinessRequest {
func (r FbsApiPutBidsForBusinessRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiPutBidsForBusinessRequest
func (a *FbsAPIService) PutBidsForBusiness(ctx context.Context, businessId int64) FbsApiPutBidsForBusinessRequest {
	return FbsApiPutBidsForBusinessRequest{
func (a *FbsAPIService) PutBidsForBusinessExecute(r FbsApiPutBidsForBusinessRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiPutBidsForCampaignRequest struct {
func (r FbsApiPutBidsForCampaignRequest) PutSkuBidsRequest(putSkuBidsRequest PutSkuBidsRequest) FbsApiPutBidsForCampaignRequest {
func (r FbsApiPutBidsForCampaignRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiPutBidsForCampaignRequest
func (a *FbsAPIService) PutBidsForCampaign(ctx context.Context, campaignId int64) FbsApiPutBidsForCampaignRequest {
	return FbsApiPutBidsForCampaignRequest{
func (a *FbsAPIService) PutBidsForCampaignExecute(r FbsApiPutBidsForCampaignRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiSearchRegionChildrenRequest struct {
func (r FbsApiSearchRegionChildrenRequest) Page(page int32) FbsApiSearchRegionChildrenRequest {
func (r FbsApiSearchRegionChildrenRequest) PageSize(pageSize int32) FbsApiSearchRegionChildrenRequest {
func (r FbsApiSearchRegionChildrenRequest) Execute() (*GetRegionWithChildrenResponse, *http.Response, error) {
	@return FbsApiSearchRegionChildrenRequest
func (a *FbsAPIService) SearchRegionChildren(ctx context.Context, regionId int64) FbsApiSearchRegionChildrenRequest {
	return FbsApiSearchRegionChildrenRequest{
func (a *FbsAPIService) SearchRegionChildrenExecute(r FbsApiSearchRegionChildrenRequest) (*GetRegionWithChildrenResponse, *http.Response, error) {
type FbsApiSearchRegionsByIdRequest struct {
func (r FbsApiSearchRegionsByIdRequest) Execute() (*GetRegionsResponse, *http.Response, error) {
	@return FbsApiSearchRegionsByIdRequest
func (a *FbsAPIService) SearchRegionsById(ctx context.Context, regionId int64) FbsApiSearchRegionsByIdRequest {
	return FbsApiSearchRegionsByIdRequest{
func (a *FbsAPIService) SearchRegionsByIdExecute(r FbsApiSearchRegionsByIdRequest) (*GetRegionsResponse, *http.Response, error) {
type FbsApiSearchRegionsByNameRequest struct {
func (r FbsApiSearchRegionsByNameRequest) Name(name string) FbsApiSearchRegionsByNameRequest {
func (r FbsApiSearchRegionsByNameRequest) PageToken(pageToken string) FbsApiSearchRegionsByNameRequest {
func (r FbsApiSearchRegionsByNameRequest) Limit(limit int32) FbsApiSearchRegionsByNameRequest {
func (r FbsApiSearchRegionsByNameRequest) Execute() (*GetRegionsResponse, *http.Response, error) {
	@return FbsApiSearchRegionsByNameRequest
func (a *FbsAPIService) SearchRegionsByName(ctx context.Context) FbsApiSearchRegionsByNameRequest {
	return FbsApiSearchRegionsByNameRequest{
func (a *FbsAPIService) SearchRegionsByNameExecute(r FbsApiSearchRegionsByNameRequest) (*GetRegionsResponse, *http.Response, error) {
type FbsApiSearchShipmentsRequest struct {
func (r FbsApiSearchShipmentsRequest) SearchShipmentsRequest(searchShipmentsRequest SearchShipmentsRequest) FbsApiSearchShipmentsRequest {
func (r FbsApiSearchShipmentsRequest) PageToken(pageToken string) FbsApiSearchShipmentsRequest {
func (r FbsApiSearchShipmentsRequest) Limit(limit int32) FbsApiSearchShipmentsRequest {
func (r FbsApiSearchShipmentsRequest) Execute() (*SearchShipmentsResponse, *http.Response, error) {
	@return FbsApiSearchShipmentsRequest
func (a *FbsAPIService) SearchShipments(ctx context.Context, campaignId int64) FbsApiSearchShipmentsRequest {
	return FbsApiSearchShipmentsRequest{
func (a *FbsAPIService) SearchShipmentsExecute(r FbsApiSearchShipmentsRequest) (*SearchShipmentsResponse, *http.Response, error) {
type FbsApiSendFileToChatRequest struct {
func (r FbsApiSendFileToChatRequest) ChatId(chatId int64) FbsApiSendFileToChatRequest {
func (r FbsApiSendFileToChatRequest) File(file *os.File) FbsApiSendFileToChatRequest {
func (r FbsApiSendFileToChatRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiSendFileToChatRequest
func (a *FbsAPIService) SendFileToChat(ctx context.Context, businessId int64) FbsApiSendFileToChatRequest {
	return FbsApiSendFileToChatRequest{
func (a *FbsAPIService) SendFileToChatExecute(r FbsApiSendFileToChatRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiSendMessageToChatRequest struct {
func (r FbsApiSendMessageToChatRequest) ChatId(chatId int64) FbsApiSendMessageToChatRequest {
func (r FbsApiSendMessageToChatRequest) SendMessageToChatRequest(sendMessageToChatRequest SendMessageToChatRequest) FbsApiSendMessageToChatRequest {
func (r FbsApiSendMessageToChatRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiSendMessageToChatRequest
func (a *FbsAPIService) SendMessageToChat(ctx context.Context, businessId int64) FbsApiSendMessageToChatRequest {
	return FbsApiSendMessageToChatRequest{
func (a *FbsAPIService) SendMessageToChatExecute(r FbsApiSendMessageToChatRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiSetOrderBoxLayoutRequest struct {
func (r FbsApiSetOrderBoxLayoutRequest) SetOrderBoxLayoutRequest(setOrderBoxLayoutRequest SetOrderBoxLayoutRequest) FbsApiSetOrderBoxLayoutRequest {
func (r FbsApiSetOrderBoxLayoutRequest) Execute() (*SetOrderBoxLayoutResponse, *http.Response, error) {
	@return FbsApiSetOrderBoxLayoutRequest
func (a *FbsAPIService) SetOrderBoxLayout(ctx context.Context, campaignId int64, orderId int64) FbsApiSetOrderBoxLayoutRequest {
	return FbsApiSetOrderBoxLayoutRequest{
func (a *FbsAPIService) SetOrderBoxLayoutExecute(r FbsApiSetOrderBoxLayoutRequest) (*SetOrderBoxLayoutResponse, *http.Response, error) {
type FbsApiSetOrderShipmentBoxesRequest struct {
func (r FbsApiSetOrderShipmentBoxesRequest) SetOrderShipmentBoxesRequest(setOrderShipmentBoxesRequest SetOrderShipmentBoxesRequest) FbsApiSetOrderShipmentBoxesRequest {
func (r FbsApiSetOrderShipmentBoxesRequest) Execute() (*SetOrderShipmentBoxesResponse, *http.Response, error) {
	@return FbsApiSetOrderShipmentBoxesRequest
func (a *FbsAPIService) SetOrderShipmentBoxes(ctx context.Context, campaignId int64, orderId int64, shipmentId int64) FbsApiSetOrderShipmentBoxesRequest {
	return FbsApiSetOrderShipmentBoxesRequest{
func (a *FbsAPIService) SetOrderShipmentBoxesExecute(r FbsApiSetOrderShipmentBoxesRequest) (*SetOrderShipmentBoxesResponse, *http.Response, error) {
type FbsApiSetShipmentPalletsCountRequest struct {
func (r FbsApiSetShipmentPalletsCountRequest) SetShipmentPalletsCountRequest(setShipmentPalletsCountRequest SetShipmentPalletsCountRequest) FbsApiSetShipmentPalletsCountRequest {
func (r FbsApiSetShipmentPalletsCountRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiSetShipmentPalletsCountRequest
func (a *FbsAPIService) SetShipmentPalletsCount(ctx context.Context, campaignId int64, shipmentId int64) FbsApiSetShipmentPalletsCountRequest {
	return FbsApiSetShipmentPalletsCountRequest{
func (a *FbsAPIService) SetShipmentPalletsCountExecute(r FbsApiSetShipmentPalletsCountRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiSkipGoodsFeedbacksReactionRequest struct {
func (r FbsApiSkipGoodsFeedbacksReactionRequest) SkipGoodsFeedbackReactionRequest(skipGoodsFeedbackReactionRequest SkipGoodsFeedbackReactionRequest) FbsApiSkipGoodsFeedbacksReactionRequest {
func (r FbsApiSkipGoodsFeedbacksReactionRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiSkipGoodsFeedbacksReactionRequest
func (a *FbsAPIService) SkipGoodsFeedbacksReaction(ctx context.Context, businessId int64) FbsApiSkipGoodsFeedbacksReactionRequest {
	return FbsApiSkipGoodsFeedbacksReactionRequest{
func (a *FbsAPIService) SkipGoodsFeedbacksReactionExecute(r FbsApiSkipGoodsFeedbacksReactionRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiSubmitReturnDecisionRequest struct {
func (r FbsApiSubmitReturnDecisionRequest) Body(body map[string]interface{}) FbsApiSubmitReturnDecisionRequest {
func (r FbsApiSubmitReturnDecisionRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiSubmitReturnDecisionRequest
func (a *FbsAPIService) SubmitReturnDecision(ctx context.Context, campaignId int64, orderId int64, returnId int64) FbsApiSubmitReturnDecisionRequest {
	return FbsApiSubmitReturnDecisionRequest{
func (a *FbsAPIService) SubmitReturnDecisionExecute(r FbsApiSubmitReturnDecisionRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiTransferOrdersFromShipmentRequest struct {
func (r FbsApiTransferOrdersFromShipmentRequest) TransferOrdersFromShipmentRequest(transferOrdersFromShipmentRequest TransferOrdersFromShipmentRequest) FbsApiTransferOrdersFromShipmentRequest {
func (r FbsApiTransferOrdersFromShipmentRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiTransferOrdersFromShipmentRequest
func (a *FbsAPIService) TransferOrdersFromShipment(ctx context.Context, campaignId int64, shipmentId int64) FbsApiTransferOrdersFromShipmentRequest {
	return FbsApiTransferOrdersFromShipmentRequest{
func (a *FbsAPIService) TransferOrdersFromShipmentExecute(r FbsApiTransferOrdersFromShipmentRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiUpdateBusinessPricesRequest struct {
func (r FbsApiUpdateBusinessPricesRequest) UpdateBusinessPricesRequest(updateBusinessPricesRequest UpdateBusinessPricesRequest) FbsApiUpdateBusinessPricesRequest {
func (r FbsApiUpdateBusinessPricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiUpdateBusinessPricesRequest
func (a *FbsAPIService) UpdateBusinessPrices(ctx context.Context, businessId int64) FbsApiUpdateBusinessPricesRequest {
	return FbsApiUpdateBusinessPricesRequest{
func (a *FbsAPIService) UpdateBusinessPricesExecute(r FbsApiUpdateBusinessPricesRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiUpdateCampaignOffersRequest struct {
func (r FbsApiUpdateCampaignOffersRequest) UpdateCampaignOffersRequest(updateCampaignOffersRequest UpdateCampaignOffersRequest) FbsApiUpdateCampaignOffersRequest {
func (r FbsApiUpdateCampaignOffersRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiUpdateCampaignOffersRequest
func (a *FbsAPIService) UpdateCampaignOffers(ctx context.Context, campaignId int64) FbsApiUpdateCampaignOffersRequest {
	return FbsApiUpdateCampaignOffersRequest{
func (a *FbsAPIService) UpdateCampaignOffersExecute(r FbsApiUpdateCampaignOffersRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiUpdateExternalOrderIdRequest struct {
func (r FbsApiUpdateExternalOrderIdRequest) UpdateExternalOrderIdRequest(updateExternalOrderIdRequest UpdateExternalOrderIdRequest) FbsApiUpdateExternalOrderIdRequest {
func (r FbsApiUpdateExternalOrderIdRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiUpdateExternalOrderIdRequest
func (a *FbsAPIService) UpdateExternalOrderId(ctx context.Context, campaignId int64, orderId int64) FbsApiUpdateExternalOrderIdRequest {
	return FbsApiUpdateExternalOrderIdRequest{
func (a *FbsAPIService) UpdateExternalOrderIdExecute(r FbsApiUpdateExternalOrderIdRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiUpdateGoodsFeedbackCommentRequest struct {
func (r FbsApiUpdateGoodsFeedbackCommentRequest) UpdateGoodsFeedbackCommentRequest(updateGoodsFeedbackCommentRequest UpdateGoodsFeedbackCommentRequest) FbsApiUpdateGoodsFeedbackCommentRequest {
func (r FbsApiUpdateGoodsFeedbackCommentRequest) Execute() (*UpdateGoodsFeedbackCommentResponse, *http.Response, error) {
	@return FbsApiUpdateGoodsFeedbackCommentRequest
func (a *FbsAPIService) UpdateGoodsFeedbackComment(ctx context.Context, businessId int64) FbsApiUpdateGoodsFeedbackCommentRequest {
	return FbsApiUpdateGoodsFeedbackCommentRequest{
func (a *FbsAPIService) UpdateGoodsFeedbackCommentExecute(r FbsApiUpdateGoodsFeedbackCommentRequest) (*UpdateGoodsFeedbackCommentResponse, *http.Response, error) {
type FbsApiUpdateOfferContentRequest struct {
func (r FbsApiUpdateOfferContentRequest) UpdateOfferContentRequest(updateOfferContentRequest UpdateOfferContentRequest) FbsApiUpdateOfferContentRequest {
func (r FbsApiUpdateOfferContentRequest) Execute() (*UpdateOfferContentResponse, *http.Response, error) {
	@return FbsApiUpdateOfferContentRequest
func (a *FbsAPIService) UpdateOfferContent(ctx context.Context, businessId int64) FbsApiUpdateOfferContentRequest {
	return FbsApiUpdateOfferContentRequest{
func (a *FbsAPIService) UpdateOfferContentExecute(r FbsApiUpdateOfferContentRequest) (*UpdateOfferContentResponse, *http.Response, error) {
type FbsApiUpdateOfferMappingEntriesRequest struct {
func (r FbsApiUpdateOfferMappingEntriesRequest) UpdateOfferMappingEntryRequest(updateOfferMappingEntryRequest UpdateOfferMappingEntryRequest) FbsApiUpdateOfferMappingEntriesRequest {
func (r FbsApiUpdateOfferMappingEntriesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiUpdateOfferMappingEntriesRequest
func (a *FbsAPIService) UpdateOfferMappingEntries(ctx context.Context, campaignId int64) FbsApiUpdateOfferMappingEntriesRequest {
	return FbsApiUpdateOfferMappingEntriesRequest{
func (a *FbsAPIService) UpdateOfferMappingEntriesExecute(r FbsApiUpdateOfferMappingEntriesRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiUpdateOfferMappingsRequest struct {
func (r FbsApiUpdateOfferMappingsRequest) UpdateOfferMappingsRequest(updateOfferMappingsRequest UpdateOfferMappingsRequest) FbsApiUpdateOfferMappingsRequest {
func (r FbsApiUpdateOfferMappingsRequest) Language(language CatalogLanguageType) FbsApiUpdateOfferMappingsRequest {
func (r FbsApiUpdateOfferMappingsRequest) Execute() (*UpdateOfferMappingsResponse, *http.Response, error) {
	@return FbsApiUpdateOfferMappingsRequest
func (a *FbsAPIService) UpdateOfferMappings(ctx context.Context, businessId int64) FbsApiUpdateOfferMappingsRequest {
	return FbsApiUpdateOfferMappingsRequest{
func (a *FbsAPIService) UpdateOfferMappingsExecute(r FbsApiUpdateOfferMappingsRequest) (*UpdateOfferMappingsResponse, *http.Response, error) {
type FbsApiUpdateOrderItemsRequest struct {
func (r FbsApiUpdateOrderItemsRequest) UpdateOrderItemRequest(updateOrderItemRequest UpdateOrderItemRequest) FbsApiUpdateOrderItemsRequest {
func (r FbsApiUpdateOrderItemsRequest) Execute() (*http.Response, error) {
	@return FbsApiUpdateOrderItemsRequest
func (a *FbsAPIService) UpdateOrderItems(ctx context.Context, campaignId int64, orderId int64) FbsApiUpdateOrderItemsRequest {
	return FbsApiUpdateOrderItemsRequest{
func (a *FbsAPIService) UpdateOrderItemsExecute(r FbsApiUpdateOrderItemsRequest) (*http.Response, error) {
type FbsApiUpdateOrderStatusRequest struct {
func (r FbsApiUpdateOrderStatusRequest) UpdateOrderStatusRequest(updateOrderStatusRequest UpdateOrderStatusRequest) FbsApiUpdateOrderStatusRequest {
func (r FbsApiUpdateOrderStatusRequest) Execute() (*UpdateOrderStatusResponse, *http.Response, error) {
	@return FbsApiUpdateOrderStatusRequest
func (a *FbsAPIService) UpdateOrderStatus(ctx context.Context, campaignId int64, orderId int64) FbsApiUpdateOrderStatusRequest {
	return FbsApiUpdateOrderStatusRequest{
func (a *FbsAPIService) UpdateOrderStatusExecute(r FbsApiUpdateOrderStatusRequest) (*UpdateOrderStatusResponse, *http.Response, error) {
type FbsApiUpdateOrderStatusesRequest struct {
func (r FbsApiUpdateOrderStatusesRequest) UpdateOrderStatusesRequest(updateOrderStatusesRequest UpdateOrderStatusesRequest) FbsApiUpdateOrderStatusesRequest {
func (r FbsApiUpdateOrderStatusesRequest) Execute() (*UpdateOrderStatusesResponse, *http.Response, error) {
	@return FbsApiUpdateOrderStatusesRequest
func (a *FbsAPIService) UpdateOrderStatuses(ctx context.Context, campaignId int64) FbsApiUpdateOrderStatusesRequest {
	return FbsApiUpdateOrderStatusesRequest{
func (a *FbsAPIService) UpdateOrderStatusesExecute(r FbsApiUpdateOrderStatusesRequest) (*UpdateOrderStatusesResponse, *http.Response, error) {
type FbsApiUpdatePricesRequest struct {
func (r FbsApiUpdatePricesRequest) UpdatePricesRequest(updatePricesRequest UpdatePricesRequest) FbsApiUpdatePricesRequest {
func (r FbsApiUpdatePricesRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiUpdatePricesRequest
func (a *FbsAPIService) UpdatePrices(ctx context.Context, campaignId int64) FbsApiUpdatePricesRequest {
	return FbsApiUpdatePricesRequest{
func (a *FbsAPIService) UpdatePricesExecute(r FbsApiUpdatePricesRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiUpdatePromoOffersRequest struct {
func (r FbsApiUpdatePromoOffersRequest) UpdatePromoOffersRequest(updatePromoOffersRequest UpdatePromoOffersRequest) FbsApiUpdatePromoOffersRequest {
func (r FbsApiUpdatePromoOffersRequest) Execute() (*UpdatePromoOffersResponse, *http.Response, error) {
	@return FbsApiUpdatePromoOffersRequest
func (a *FbsAPIService) UpdatePromoOffers(ctx context.Context, businessId int64) FbsApiUpdatePromoOffersRequest {
	return FbsApiUpdatePromoOffersRequest{
func (a *FbsAPIService) UpdatePromoOffersExecute(r FbsApiUpdatePromoOffersRequest) (*UpdatePromoOffersResponse, *http.Response, error) {
type FbsApiUpdateStocksRequest struct {
func (r FbsApiUpdateStocksRequest) UpdateStocksRequest(updateStocksRequest UpdateStocksRequest) FbsApiUpdateStocksRequest {
func (r FbsApiUpdateStocksRequest) Execute() (*EmptyApiResponse, *http.Response, error) {
	@return FbsApiUpdateStocksRequest
func (a *FbsAPIService) UpdateStocks(ctx context.Context, campaignId int64) FbsApiUpdateStocksRequest {
	return FbsApiUpdateStocksRequest{
func (a *FbsAPIService) UpdateStocksExecute(r FbsApiUpdateStocksRequest) (*EmptyApiResponse, *http.Response, error) {
type FbsApiUpdateWarehouseStatusRequest struct {
func (r FbsApiUpdateWarehouseStatusRequest) UpdateWarehouseStatusRequest(updateWarehouseStatusRequest UpdateWarehouseStatusRequest) FbsApiUpdateWarehouseStatusRequest {
func (r FbsApiUpdateWarehouseStatusRequest) Execute() (*UpdateWarehouseStatusResponse, *http.Response, error) {
	@return FbsApiUpdateWarehouseStatusRequest
func (a *FbsAPIService) UpdateWarehouseStatus(ctx context.Context, campaignId int64) FbsApiUpdateWarehouseStatusRequest {
	return FbsApiUpdateWarehouseStatusRequest{
func (a *FbsAPIService) UpdateWarehouseStatusExecute(r FbsApiUpdateWarehouseStatusRequest) (*UpdateWarehouseStatusResponse, *http.Response, error) {
