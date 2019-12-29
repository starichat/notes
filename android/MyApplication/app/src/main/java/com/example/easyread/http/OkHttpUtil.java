package com.example.easyread.http;

import android.util.Log;

import java.io.IOException;

import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;

public class OkHttpUtil {

    public static final MediaType JSON = MediaType.parse("application/json; charset=utf-8");

    public static  OkHttpClient client = new OkHttpClient();

    public static String sendPostjson(String api,String token,String json) throws IOException {

        RequestBody body = RequestBody.create(JSON, json);
        Request request = new Request.Builder().url(api).header("Authorization",token).post(body).build();
        Response response = client.newCall(request).execute();
        if (response.isSuccessful()) {
            return response.body().string();
        } else {
            throw new IOException("Unexpected code " + response);
        }

    }
    public static String auth(String api,String json) throws IOException {

        RequestBody body = RequestBody.create(JSON, json);
        Request request = new Request.Builder().url(api).post(body).build();
        Response response = client.newCall(request).execute();
        if (response.isSuccessful()) {
            return response.body().string();
        } else {
            Log.e("wxl","response err");

        }
        return null;

    }

    public static String sendGetjson(String api,String token,String json) throws IOException {
        String url = api + json; // 将参数放到api中
        RequestBody body = RequestBody.create(JSON, json);
        Request request = new Request.Builder().url(url).header("Authorization",token).get().build();
        Response response = client.newCall(request).execute();
        if (response.isSuccessful()) {
            return response.body().string();
        } else {
            throw new IOException("Unexpected code " + response);
        }
    }



}
