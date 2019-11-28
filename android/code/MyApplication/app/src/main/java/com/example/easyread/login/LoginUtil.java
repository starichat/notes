package com.example.easyread.login;

import android.util.Log;

import com.example.easyread.http.OkHttpUtil;
import com.example.easyread.http.ResponseResult;
import com.example.easyread.pojo.ResultUtil;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.Callback;
import okhttp3.Response;

public class LoginUtil {


    public static ResponseResult res = null;

    // 登录验证
    public static ResponseResult doLogin(String url, String userJson) {


        OkHttpUtil.sendOkhttpPostRequest(url, userJson,new Callback() {
            @Override
            public void onFailure(Call call, IOException e) {
                Log.d("wxl","fail");
                ResultUtil.fail();
                e.printStackTrace();
                return ;
            }

            @Override
            public void onResponse(Call call, Response response) throws IOException {
                try {
                    Log.d("wxl","success");
                    Log.d("wxl",response.body().string());

                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
        });


        return ResultUtil.ok("成功");
    }





}
