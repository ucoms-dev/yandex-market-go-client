type DbsApiGetAuthTokenInfoRequest struct {
func (r DbsApiGetAuthTokenInfoRequest) Execute() (*GetTokenInfoResponse, *http.Response, error) {
 @return DbsApiGetAuthTokenInfoRequest
func (a *DbsAPIService) GetAuthTokenInfo(ctx context.Context) DbsApiGetAuthTokenInfoRequest {
	return DbsApiGetAuthTokenInfoRequest{
func (a *DbsAPIService) GetAuthTokenInfoExecute(r DbsApiGetAuthTokenInfoRequest) (*GetTokenInfoResponse, *http.Response, error) {
