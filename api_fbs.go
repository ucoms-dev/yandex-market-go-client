type FbsApiCreateChatRequest struct {
func (r FbsApiCreateChatRequest) CreateChatRequest(createChatRequest CreateChatRequest) FbsApiCreateChatRequest {
func (r FbsApiCreateChatRequest) Execute() (*CreateChatResponse, *http.Response, error) {
	@return FbsApiCreateChatRequest
func (a *FbsAPIService) CreateChat(ctx context.Context, businessId int64) FbsApiCreateChatRequest {
	return FbsApiCreateChatRequest{
func (a *FbsAPIService) CreateChatExecute(r FbsApiCreateChatRequest) (*CreateChatResponse, *http.Response, error) {
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
