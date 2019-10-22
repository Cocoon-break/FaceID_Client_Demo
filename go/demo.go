package main

import (
	"io"
	"net/http"
	"bytes"
	"mime/multipart"
	"strings"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/tidwall/gjson"
)

func postFormData(url string, params map[string]io.Reader) (*http.Response, error) {
	client := &http.Client{}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	var err error
	for key, r := range params {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if _, ok := r.(*strings.Reader); ok {
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		} else {
			if fw, err = w.CreateFormFile(key, key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}
	w.Close()

	req, err := http.NewRequest(http.MethodPost, url, &b)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return client.Do(req)
}

func postJson(url string, jsonBytes []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}

func minionsVerify(url string) {
	image_idcard, err := os.Open("a.jpg")
	if err != nil {
		fmt.Printf("open file err:%s", err.Error())
		return
	}
	image_best, err := os.Open("a.jpg")
	if err != nil {
		fmt.Printf("open file err:%s", err.Error())
		return
	}
	//todo:更多参数请参考API文档
	params := map[string]io.Reader{
		"image_idcard": image_idcard,
		"image_best":   image_best,
	}
	resp, err := postFormData(url, params)
	if err != nil {
		fmt.Printf("request err:%s\n", err.Error())
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("result:%s", body)
}
func minionsExtractWidthDetect(url string) (feat string, err error) {
	image, err := os.Open("a.jpg")
	if err != nil {
		fmt.Printf("open file err:%s", err.Error())
		return "", err
	}
	//todo:更多参数请参考API文档
	params := map[string]io.Reader{
		"image":  image,
		"cutimg": strings.NewReader("1"),
	}
	resp, err := postFormData(url, params)
	if err != nil {
		fmt.Printf("request err:%s \n", err.Error())
		return "", nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := gjson.Get(string(body), "faces.0.Feature")
	return result.Str, nil
}
func minionsCompare(url string, feat string) {
	//todo:更多参数请参考API文档
	json, _ := json.Marshal(map[string]string{
		"feat1": feat,
		"feat2": feat,
	})
	resp, err := postJson(url, json)
	if err != nil {
		fmt.Printf("request err:%s", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("result:%s \n", body)
}

func ocr(url string) {
	image, err := os.Open("idcard_front.jpg")
	if err != nil {
		fmt.Printf("open file err:%s", err.Error())
		return
	}
	//todo:更多参数请参考API文档
	params := map[string]io.Reader{
		"img":      image,
		"legality": strings.NewReader("1"),
	}
	resp, err := postFormData(url, params)
	if err != nil {
		fmt.Printf("request err:%s \n", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("result:%s \n", body)
}

func FMP(url string){
	image, err := os.Open("a.jpg")
	if err != nil {
		fmt.Printf("open file err:%s", err.Error())
		return
	}
	//todo:更多参数请参考API文档
	params := map[string]io.Reader{
		"img":      image,
		"rotate": strings.NewReader("1"),
	}
	resp, err := postFormData(url, params)
	if err != nil {
		fmt.Printf("request err:%s \n", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("result:%s \n", body)
}

func groupInit(url string){
	//todo:更多参数请参考API文档
	params := map[string]io.Reader{
		"groupname": strings.NewReader("test"),
	}
	resp, err := postFormData(url, params)
	if err != nil {
		fmt.Printf("request err:%s \n", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("result:%s \n", body)
}
func groupAdd(url string,feat string){
	//todo:更多参数请参考API文档
	params := map[string]io.Reader{
		"groupname": strings.NewReader("test"),
		"feature": strings.NewReader(feat),
	}
	resp, err := postFormData(url, params)
	if err != nil {
		fmt.Printf("request err:%s \n", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("result:%s \n", body)
}
func groupSearch(url string,feat string){
	//todo:更多参数请参考API文档
	params := map[string]io.Reader{
		"groupname": strings.NewReader("test"),
		"feature": strings.NewReader(feat),
	}
	resp, err := postFormData(url, params)
	if err != nil {
		fmt.Printf("request err:%s \n", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("result:%s \n", body)
}

func main() {
	minionsVerify("http://10.201.102.247:9000/verify")
	feat,_:=minionsExtractWidthDetect("http://10.201.102.247:9000/extract_with_detect")
	minionsCompare("http://10.201.102.247:9000/compare",feat)

	ocr("http://10.201.102.247:9002/faceid/v1/idcard")
	FMP("http://10.201.102.247:9001/faceid/v1/liveness_cfg")

	groupInit("http://10.201.102.247:9090/group/init")
	groupAdd("http://10.201.102.247:9090/group/add",feat)
	groupSearch("http://10.201.102.247:9090/group/search",feat)
}
