package com.megvii.demo;

import org.apache.commons.httpclient.HttpClient;
import org.apache.commons.httpclient.HttpException;
import org.apache.commons.httpclient.methods.PostMethod;
import org.apache.commons.httpclient.methods.StringRequestEntity;
import org.apache.commons.httpclient.methods.multipart.MultipartRequestEntity;
import org.apache.commons.httpclient.methods.multipart.Part;
import org.apache.commons.httpclient.params.HttpMethodParams;
import org.json.JSONObject;

import java.io.*;
import java.util.HashMap;
import java.util.Map;

/**
 * Created by wushengping on 2019/10/21.
 */
public class HttpUtils {
    public static String postWithFormData(String url, Part[] parts) {
        HttpClient client = new HttpClient();
        PostMethod postMethod = new PostMethod(url);
        try {
            MultipartRequestEntity multipartRequestEntity = new MultipartRequestEntity(parts, new HttpMethodParams());
            postMethod.setRequestEntity(multipartRequestEntity);
            int statusCode = client.executeMethod(postMethod);
            if (statusCode != 200) {
                Map map1 = new HashMap();
                map1.put("msg", "http 请求错误");
                return new JSONObject(map1).toString();
            }
            InputStream inputStream = postMethod.getResponseBodyAsStream();
            BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream));
            String line = "";
            StringBuilder sb = new StringBuilder();
            while ((line = reader.readLine()) != null) {
                sb.append(line + "\n");
            }
            inputStream.close();
            return sb.toString();
        } catch (IOException e) {
            e.printStackTrace();
            return "";
        }
    }

    public static String postJson(String url, String json) {
        StringRequestEntity requestEntity = null;
        try {
            HttpClient client = new HttpClient();
            requestEntity = new StringRequestEntity(json, "application/json", "UTF-8");
            PostMethod postMethod = new PostMethod(url);
            postMethod.setRequestEntity(requestEntity);

            int statusCode = client.executeMethod(postMethod);
            if (statusCode != 200) {
                Map map1 = new HashMap();
                map1.put("msg", "http 请求错误");
                return new JSONObject(map1).toString();
            }
            InputStream inputStream = postMethod.getResponseBodyAsStream();
            BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream));
            String line = "";
            StringBuilder sb = new StringBuilder();
            while ((line = reader.readLine()) != null) {
                sb.append(line + "\n");
            }
            inputStream.close();
            return sb.toString();
        } catch (IOException e) {
            e.printStackTrace();
            return "";
        }
    }
}
