package com.example.easyread.service;

import android.app.Service;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.IBinder;

import com.example.easyread.http.OkHttpUtil;
import com.google.gson.Gson;

import java.io.IOException;

/**
 * 用户处理
 * 1. 注册登录,基于oauth2 进行token校验
 * 2.
 */
public class UserService extends Service {

    private SharedPreferences sharedPreferences = getSharedPreferences("auth", MODE_PRIVATE);
    private SharedPreferences.Editor editor = sharedPreferences.edit();

    public UserService() {
    }

    @Override
    public IBinder onBind(Intent intent) {
        // TODO: Return the communication channel to the service.
        throw new UnsupportedOperationException("Not yet implemented");
    }

    public String dologin(String json) throws IOException {
        Gson gson = new Gson();
        return OkHttpUtil.auth("/login",json);
    }

    public String refreshToken(String token) throws IOException {


        if(token == null) {
            //TODO login
        }
        // verified token if valid
        // 过期则请求新的token
        token = OkHttpUtil.auth("",token);// 请求新的token
        // 保存新的token
        editor.putString("token", token);
        editor.apply();

        return token;
    }

    public String getToken(){
        return sharedPreferences.getString("token","");
    }




}
