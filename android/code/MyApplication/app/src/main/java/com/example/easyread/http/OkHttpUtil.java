package com.example.easyread.http;

import android.util.Log;

import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;

public class OkHttpUtil {
    public static OkHttpClient client = new OkHttpClient();

    public static final MediaType JSON = MediaType.parse("application/json; charset=utf-8");

    public static void sendOkHttpRequest(String url, okhttp3.Callback callback) {

        Request request = new Request.Builder().url(url).build();
        client.newCall(request).enqueue(callback);
    }

    public static void sendOkhttpPostRequest(String url, String json, okhttp3.Callback callback) {
        RequestBody body = RequestBody.create(JSON,json);
        Log.d("wxl",json);
        Request request = new Request.Builder()
                .url(url)
                .post(body)
                .build();

        client.newCall(request).enqueue(callback);
    }

}
