package com.example.easyread.http;

import java.io.IOException;

import okhttp3.Authenticator;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.Route;

public class TokenAuthenticator implements Authenticator {




    public String refreshToken(){
        return "";
    }
    @Override
    public Request authenticate(Route route, Response response) throws IOException {

        String token = refreshToken();


        return response.request()
                .newBuilder()
                .header("Authorization",token)
                .build();
    }



}
