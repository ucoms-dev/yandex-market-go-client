type FbsApiGetAuthTokenInfoRequest struct {
func (r FbsApiGetAuthTokenInfoRequest) Execute() (*GetTokenInfoResponse, *http.Response, error) {
 @return FbsApiGetAuthTokenInfoRequest
func (a *FbsAPIService) GetAuthTokenInfo(ctx context.Context) FbsApiGetAuthTokenInfoRequest {
	return FbsApiGetAuthTokenInfoRequest{
func (a *FbsAPIService) GetAuthTokenInfoExecute(r FbsApiGetAuthTokenInfoRequest) (*GetTokenInfoResponse, *http.Response, error) {
