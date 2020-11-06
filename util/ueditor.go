package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/sipt/GoJsoner"
	"github.com/spf13/viper"
)

type Config struct {
	ImageActionName         string   `json:"imageActionName"`
	ImageFieldName          string   `json:"imageFieldName"`
	ImageMaxSize            int      `json:"imageMaxSize"`
	ImageAllowFiles         []string `json:"imageAllowFiles"`
	ImageCompressEnable     bool     `json:"imageCompressEnable"`
	ImageCompressBorder     int      `json:"imageCompressBorder"`
	ImageInsertAlign        string   `json:"imageInsertAlign"`
	ImageURLPrefix          string   `json:"imageUrlPrefix"`
	ImagePathFormat         string   `json:"imagePathFormat"`
	ScrawlActionName        string   `json:"scrawlActionName"`
	ScrawlFieldName         string   `json:"scrawlFieldName"`
	ScrawlPathFormat        string   `json:"scrawlPathFormat"`
	ScrawlMaxSize           int      `json:"scrawlMaxSize"`
	ScrawlURLPrefix         string   `json:"scrawlUrlPrefix"`
	ScrawlInsertAlign       string   `json:"scrawlInsertAlign"`
	SnapscreenActionName    string   `json:"snapscreenActionName"`
	SnapscreenPathFormat    string   `json:"snapscreenPathFormat"`
	SnapscreenURLPrefix     string   `json:"snapscreenUrlPrefix"`
	SnapscreenInsertAlign   string   `json:"snapscreenInsertAlign"`
	CatcherLocalDomain      []string `json:"catcherLocalDomain"`
	CatcherActionName       string   `json:"catcherActionName"`
	CatcherFieldName        string   `json:"catcherFieldName"`
	CatcherPathFormat       string   `json:"catcherPathFormat"`
	CatcherURLPrefix        string   `json:"catcherUrlPrefix"`
	CatcherMaxSize          int      `json:"catcherMaxSize"`
	CatcherAllowFiles       []string `json:"catcherAllowFiles"`
	VideoActionName         string   `json:"videoActionName"`
	VideoFieldName          string   `json:"videoFieldName"`
	VideoPathFormat         string   `json:"videoPathFormat"`
	VideoURLPrefix          string   `json:"videoUrlPrefix"`
	VideoMaxSize            int      `json:"videoMaxSize"`
	VideoAllowFiles         []string `json:"videoAllowFiles"`
	FileActionName          string   `json:"fileActionName"`
	FileFieldName           string   `json:"fileFieldName"`
	FilePathFormat          string   `json:"filePathFormat"`
	FileURLPrefix           string   `json:"fileUrlPrefix"`
	FileMaxSize             int      `json:"fileMaxSize"`
	FileAllowFiles          []string `json:"fileAllowFiles"`
	ImageManagerActionName  string   `json:"imageManagerActionName"`
	ImageManagerListPath    string   `json:"imageManagerListPath"`
	ImageManagerListSize    int      `json:"imageManagerListSize"`
	ImageManagerURLPrefix   string   `json:"imageManagerUrlPrefix"`
	ImageManagerInsertAlign string   `json:"imageManagerInsertAlign"`
	ImageManagerAllowFiles  []string `json:"imageManagerAllowFiles"`
	FileManagerActionName   string   `json:"fileManagerActionName"`
	FileManagerListPath     string   `json:"fileManagerListPath"`
	FileManagerURLPrefix    string   `json:"fileManagerUrlPrefix"`
	FileManagerListSize     int      `json:"fileManagerListSize"`
	FileManagerAllowFiles   []string `json:"fileManagerAllowFiles"`
}

type ResFileInfo struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Original string `json:"original"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
}

type ResFileInfoWithState struct {
	State string `json:"state"`
	ResFileInfo
}

type ResFilesInfoWithStates struct {
	State string         `json:"state"`
	List  []*ResFileInfo `json:"list"`
}

type fileInfo struct {
	URL   string `json:"url"`
	Mtime int    `json:"mtime"`
}
type ResFileList struct {
	State string      `json:"state"`
	Start int         `json:"start"`
	Total int         `json:"total"`
	List  []*fileInfo `json:"list"`
}

func (resState *ResFileInfoWithState) fromResFileInfo(res *ResFileInfo) {
	resState.URL = res.URL
	resState.Title = res.Title
	resState.Original = res.Original
	resState.Type = res.Type
	resState.Size = res.Size
}

// 上传文件的参数
type UploaderParams struct {
	PathFormat       string   /* 上传保存路径,可以自定义保存路径和文件名格式 */
	MaxSize          int      /* 上传大小限制，单位B */
	AllowFiles       []string /* 上传格式限制 */
	OriName          string   /* 原始文件名 */
	CloudStoragePath string   /*云存储路径*/
}

type UploadParams struct {
}

type LocalStoreHandler func(savePath string, name string, reader io.Reader)
type Uploader struct {
	RootPath   string // 项目根目录
	params     *UploaderParams
	cfg        *Config
	localStore LocalStoreHandler
	localPath  string
	up         *OssUploader
}

var (
	globalCfg = &Config{}
	once      = new(sync.Once)
)

func loadConfig() {
	//curDir, _ := os.Getwd()
	//filePath := path.Join(curDir, viper.GetString("ueditorConfig"))
	filePath := viper.GetString("ueditorConfig")
	log.Printf("ueditor config file path:%s\n", filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	strData, err := GoJsoner.Discard(string(data))
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(strData), globalCfg); err != nil {
		panic(err)
	}
}

func GetUEditorConfig() interface{} {
	once.Do(loadConfig)
	return globalCfg
}

func NewUploader(localStore LocalStoreHandler, projectTag string) *Uploader {
	once.Do(loadConfig)
	uploaderObj := &Uploader{cfg: globalCfg}
	//uploaderObj.SetParams(upParams)
	uploaderObj.localStore = localStore
	uploaderObj.up = NewOssUploader(projectTag)
	return uploaderObj
}

func (up *Uploader) SetParams(params *UploaderParams) (err error) {
	up.params = params
	return
}

func (up *Uploader) SetRootPath(path string) error {
	up.RootPath = path
	return nil
}

func listFiles() {

}

func isExtInWhiteList(name string, whiteList []string) bool {
	for _, w := range whiteList {
		if name == w {
			return true
		}
	}
	return false
}
func walkDir(dirPath string, fileInfoList []*fileInfo, extWhiteList []string) ([]*fileInfo, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			fileInfoList, _ = walkDir(path.Join(dirPath, file.Name()), fileInfoList, extWhiteList)
			continue
		}

		if !isExtInWhiteList(getFileExt(fileName), extWhiteList) {
			continue
		}
		fileInfoList = append(fileInfoList, &fileInfo{
			URL:   "http://192.168.137.1:8040/" + path.Join(dirPath, file.Name()),
			Mtime: int(file.ModTime().Unix()),
		})
	}
	return fileInfoList, nil
}

func (up *Uploader) ListImages(start, num int) (*ResFileList, error) {
	fileInfoList := make([]*fileInfo, 0)

	fileInfoList, err := walkDir(up.cfg.ImageManagerListPath, fileInfoList, up.cfg.ImageManagerAllowFiles)
	if err != nil {
		return nil, err
	}
	fileList := &ResFileList{
		State: "SUCCESS",
		Start: start,
		Total: len(fileInfoList),
		List:  fileInfoList,
	}
	return fileList, nil
}

func (up *Uploader) ListFiles(start, num int) (*ResFileList, error) {
	fileInfoList := make([]*fileInfo, 0)
	fileInfoList, err := walkDir(up.cfg.FileManagerListPath, fileInfoList, up.cfg.FileManagerAllowFiles)
	if err != nil {
		return nil, err
	}
	fileList := &ResFileList{
		State: "SUCCESS",
		Start: start,
		Total: len(fileInfoList),
		List:  fileInfoList,
	}
	return fileList, nil

}

func (up *Uploader) UpScrawl(r *http.Request, savePath string) (*ResFileInfoWithState, error) {
	cfg := up.cfg
	file, header, err := r.FormFile(up.cfg.ScrawlFieldName)
	if err != nil {
		return nil, err
	}
	if header.Size > int64(cfg.ScrawlMaxSize) {
		return nil, fmt.Errorf("scraw file:%s size:%d beyond limit:%d\n",
			header.Filename, header.Size, cfg.ScrawlMaxSize)
	}

	//TODO: check type
	return up.upFile(savePath, file, header)
}

func (up *Uploader) UpImage(r *http.Request, savePath string) (*ResFileInfoWithState, error) {
	cfg := up.cfg
	file, header, err := r.FormFile(up.cfg.ImageFieldName)
	up.localPath = up.cfg.ImagePathFormat
	if err != nil {
		return nil, err
	}
	if header.Size > int64(cfg.ImageMaxSize) {
		return nil, fmt.Errorf("image file:%s size:%d beyond limit:%d\n",
			header.Filename, header.Size, cfg.ImageMaxSize)
	}

	//TODO: check type
	return up.upFile(savePath, file, header)
}

func (up *Uploader) UpVideo(r *http.Request, savePath string) (*ResFileInfoWithState, error) {
	cfg := up.cfg
	file, header, err := r.FormFile(up.cfg.VideoFieldName)
	if err != nil {
		return nil, err
	}
	if header.Size > int64(cfg.VideoMaxSize) {
		return nil, fmt.Errorf("video file:%s size:%d beyond limit:%d\n",
			header.Filename, header.Size, cfg.VideoMaxSize)
	}

	//TODO: check type
	return up.upFile(savePath, file, header)
}

func (up *Uploader) UpFile(r *http.Request, savePath string) (*ResFileInfoWithState, error) {
	cfg := up.cfg
	file, header, err := r.FormFile(up.cfg.FileFieldName)
	if err != nil {
		return nil, err
	}
	up.localPath = cfg.FilePathFormat
	if header.Size > int64(cfg.FileMaxSize) {
		return nil, fmt.Errorf("file file:%s size:%d beyond limit:%d\n",
			header.Filename, header.Size, cfg.FileMaxSize)
	}

	//TODO: check type
	return up.upFile(savePath, file, header)
}

func (up *Uploader) upFile(savePath string, file multipart.File, fileHeader *multipart.FileHeader) (*ResFileInfoWithState, error) {
	data := make([]byte, fileHeader.Size)
	if _, err := file.Read(data); err != nil {
		return nil, err
	}
	fileName := fileHeader.Filename
	fileUrl, err := up.up.UploadData(data, savePath, fileName)
	//fileUrl, err := uploadFile(savePath, fileName, data)
	if err != nil {
		return nil, err
	}
	if up.localStore != nil && up.localPath != "" {
		up.localStore(up.localPath, fileName, bytes.NewReader(data))
	}

	fileInfo := &ResFileInfoWithState{}
	ext := getFileExt(fileName)
	fileInfo.Size = fileHeader.Size
	fileInfo.Type = ext
	//fileInfo.Title = filepath.Base(fileAbsPath)
	fileInfo.Title = "ok"
	fileInfo.Original = fileName
	fileInfo.URL = fileUrl
	fileInfo.State = "SUCCESS"
	return fileInfo, nil
}

func LocalStore(savePath string, fileName string, reader io.Reader) {
	curDir, err := os.Getwd()
	if err != nil {
		log.Printf("LocalStore %v\n", err)
		return
	}
	relativePath := genPath(savePath, fileName)
	if runtime.GOOS == "windows" {
		curDir = strings.Replace(curDir, "\\", "/", -1)
	}
	relativeDir := path.Dir(relativePath)
	absPath := path.Join(curDir, relativePath)
	dstDir := path.Join(curDir, relativeDir)

	_, err = os.Stat(dstDir)
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dstDir, 0766); err != nil {
			log.Printf("LocalStore MkDirAll failed. %v\n", err)
			return
		}
	}
	file, err := os.Create(absPath)
	if err != nil {
		log.Printf("LocalStore os.Create failed. %v\n", err)
		return
	}
	defer file.Close()
	io.Copy(file, reader)

}
func getFileExt(name string) string {
	ext := path.Ext(name)
	return ext
	//return ext[1:]
}

/*
func (up *Uploader) UpBase64(fileName, base64data string) (fileInfo *ResFileInfo, err error) {

	imgData, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		err = errors.New(common.ERROR_BASE64_DATA)
		return
	}

	fileSize := len(imgData)
	// 校验文件大小
	err = up.checkSize(int64(fileSize))
	if err != nil {
		return
	}

	ext := filepath.Ext(fileName)
	err = up.checkType(ext)
	if err != nil {
		return
	}

	fullName := up.getFullName(fileName)
	fileAbsPath := up.getFilePath(fullName)

	fileUrl, err := up.Storage.SaveData(imgData, fileAbsPath, fullName)
	if err != nil {
		return
	}

	fileInfo = &ResFileInfo{}
	fileInfo.Size = int64(fileSize)
	fileInfo.Type = ext
	fileInfo.Title = filepath.Base(fileAbsPath)
	fileInfo.Original = up.params.OriName
	fileInfo.URL = fileUrl

	return
}

//拉取远程文件并保存
func (up *Uploader) SaveRemote(remoteUrl string) (fileInfo *ResFileInfo, err error) {
	urlObj, err := url.Parse(remoteUrl)
	if err != nil {
		err = errors.New(common.INVALID_URL)
		return
	}

	scheme := strings.ToLower(urlObj.Scheme)
	if scheme != "http" && scheme != "https" {
		err = errors.New(common.ERROR_HTTP_LINK)
		return
	}

	// 校验文件类型
	ext := filepath.Ext(urlObj.Path)
	err = up.checkType(ext)
	if err != nil {
		return
	}

	fileName := filepath.Base(urlObj.Path)
	fullName := up.getFullName(fileName)
	fileAbsPath := up.getFilePath(fullName)

	client := http.Client{Timeout: 5 * time.Second}
	// 校验是否是可用的链接
	headResp, err := client.Head(remoteUrl)
	if err == nil {
		defer func() {
			headResp.Body.Close()
		}()
	}
	if err != nil || headResp.StatusCode != http.StatusOK {
		err = errors.New(common.ERROR_DEAD_LINK)
		return
	}
	// 校验content-type
	contentType := headResp.Header.Get("Content-Type")
	if !strings.Contains(strings.ToLower(contentType), "image") {
		err = errors.New(common.ERROR_HTTP_CONTENTTYPE)
		return
	}

	resp, err := client.Get(remoteUrl)
	if err == nil {
		defer func() {
			resp.Body.Close()
		}()
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		err = errors.New(common.ERROR_DEAD_LINK)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.New(common.ERRPR_READ_REMOTE_DATA)
		return
	}

	fileUrl, err := up.Storage.SaveData(data, fileAbsPath, fullName)

	fileInfo = &ResFileInfo{}
	fileInfo.Size = int64(len(data))
	fileInfo.Type = ext
	fileInfo.Title = filepath.Base(fileAbsPath)
	fileInfo.Original = fileName
	fileInfo.URL = fileUrl

	return
}

//根据原始文件名生成新文件名
func (up *Uploader) getFullName(oriName string) string {
	timeNow := time.Now()
	timeNowFormat := time.Now().Format("2006_01_02_15_04_05")
	timeArr := strings.Split(timeNowFormat, "_")

	format := up.params.PathFormat

	format = strings.Replace(format, "{yyyy}", timeArr[0], 1)
	format = strings.Replace(format, "{mm}", timeArr[1], 1)
	format = strings.Replace(format, "{dd}", timeArr[2], 1)
	format = strings.Replace(format, "{hh}", timeArr[3], 1)
	format = strings.Replace(format, "{ii}", timeArr[4], 1)
	format = strings.Replace(format, "{ss}", timeArr[5], 1)

	timestamp := strconv.FormatInt(timeNow.UnixNano(), 10)
	format = strings.Replace(format, "{time}", string(timestamp), 1)

	pattern := "{rand:(\\d)+}"
	if ok, _ := regexp.MatchString(pattern, format); ok {
		// 生成随机字符串
		exp, _ := regexp.Compile(pattern)
		randLenStr := exp.FindSubmatch([]byte(format))

		randLen, _ := strconv.Atoi(string(randLenStr[1]))
		randStr := strconv.Itoa(rand.Int())
		randStrLen := len(randStr)
		if randStrLen > randLen {
			randStr = randStr[randStrLen-randLen:]
		}
		// 将随机传替换到format中
		format = exp.ReplaceAllString(format, randStr)
	}

	ext := filepath.Ext(oriName)

	return format + ext
}

func (up *Uploader) getFilePath(fullName string) string {
	return filepath.Join(up.RootPath, fullName)
}

*/

/*
//设置文件存储对象
func (up *Uploader) SetStorage(storageObj storage.BaseInterface) (err error) {
	up.Storage = storageObj
	return
}

*/
