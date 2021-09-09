package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author demon
 * @Description //TODO 常用函数库
 * @Date 2020-7-12 17:08:54
 **/

//注意client 本身是连接池，不要每次请求时创建client
var (
	HttpClient = &http.Client{
		Timeout: 30 * time.Second,
	}
)

func DelFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil //f为空 错误不为空 错误是文件不存在 可以忽略
		}
		if f.IsDir() {
			fmt.Printf("文件夹 继续递归 %s  \n", path)
			DelFilelist(path)
		} else {
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("删除文件 %s  \n", path)
			return nil
		}
		//	println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("walk 错误 err: %v\n", err)
	}
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func UploadFile(url string, params map[string]string, nameField, fileName string, file io.Reader) ([]byte, error) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func DistributeFile(url string, params map[string]string, nameField, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		writer.WriteField(key, val)
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// Replace 根据替换表执行批量替换
func Replace(table map[string]string, s string) string {
	for key, value := range table {
		s = strings.Replace(s, key, value, -1)
	}
	return s
}

// Int2Str int类型转string类型
func Int2Str(inter int) string {
	string := strconv.Itoa(inter)
	return string
}

// Int2Str int64类型转string类型
func Int642Str(inter int64) string {
	string := strconv.FormatInt(inter, 10)
	return string
}

// Str2Int string类型转Int类型
func Str2Int(inter string) int {
	int, _ := strconv.Atoi(inter)
	return int
}

// Str2Int64 string类型转Int64类型
func Str2Int64(inter string) int64 {
	int64, _ := strconv.ParseInt(inter, 10, 64)
	return int64
}

func Arr2Str(strings []string) string {
	b, _ := json.Marshal(strings)
	return fmt.Sprintf("%s", b)
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
