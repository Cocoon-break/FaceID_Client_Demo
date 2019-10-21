package com.megvii.demo;

import org.apache.commons.httpclient.methods.multipart.FilePart;
import org.apache.commons.httpclient.methods.multipart.Part;
import org.apache.commons.httpclient.methods.multipart.StringPart;

import java.io.File;
import java.io.FileNotFoundException;

/**
 * Created by wushengping on 2019/10/21.
 */
public class FMP {
    public static void main(String[] args) {
        FMP("http://10.201.102.247:9001/faceid/v1/liveness_cfg");
    }
    public static void FMP(String url){
        try {
            //todo: 更多请求参数请查看api文档
            FilePart bestPart = new FilePart("img", new File("a.jpg"));
            StringPart stringPart = new StringPart("rotate", "1");
            Part[] parts = new Part[]{bestPart, stringPart};
            String result=HttpUtils.postWithFormData(url, parts);
            System.out.print(result);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
    }
}
