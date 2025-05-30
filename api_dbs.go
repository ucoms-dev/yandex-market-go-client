type DbsApiCreateChatRequest struct {
func (r DbsApiCreateChatRequest) CreateChatRequest(createChatRequest CreateChatRequest) DbsApiCreateChatRequest {
func (r DbsApiCreateChatRequest) Execute() (*CreateChatResponse, *http.Response, error) {
	@return DbsApiCreateChatRequest
func (a *DbsAPIService) CreateChat(ctx context.Context, businessId int64) DbsApiCreateChatRequest {
	return DbsApiCreateChatRequest{
func (a *DbsAPIService) CreateChatExecute(r DbsApiCreateChatRequest) (*CreateChatResponse, *http.Response, error) {
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
