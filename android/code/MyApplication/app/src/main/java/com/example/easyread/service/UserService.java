package com.example.easyread.service;

import android.app.Service;
import android.content.Intent;
import android.os.IBinder;

/**
 * 用户处理
 * 1. 注册登录,基于oauth2 进行token校验
 * 2.
 */
public class UserService extends Service {
    public UserService() {
    }

    @Override
    public IBinder onBind(Intent intent) {
        // TODO: Return the communication channel to the service.
        throw new UnsupportedOperationException("Not yet implemented");
    }


}
