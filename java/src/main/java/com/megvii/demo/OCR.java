package com.megvii.demo;

import org.apache.commons.httpclient.methods.multipart.FilePart;
import org.apache.commons.httpclient.methods.multipart.Part;
import org.apache.commons.httpclient.methods.multipart.StringPart;

import java.io.File;
import java.io.FileNotFoundException;

/**
 * Created by wushengping on 2019/10/21.
 */
public class OCR {
    public static void main(String[] args) {
        OCR("http://10.201.102.247:9002/faceid/v1/idcard");
    }

    public static void OCR(String url){
        try {
            //todo: 更多请求参数请查看api文档
            FilePart bestPart = new FilePart("img", new File("idcard_front.jpg"));
            StringPart stringPart = new StringPart("legality", "1");
            Part[] parts = new Part[]{bestPart, stringPart};
            String result=HttpUtils.postWithFormData(url, parts);
            System.out.print(result);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
    }
}
