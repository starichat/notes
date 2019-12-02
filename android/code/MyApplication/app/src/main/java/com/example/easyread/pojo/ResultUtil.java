package com.example.easyread.pojo;

import com.example.easyread.http.ResponseResult;

public class ResultUtil {




    public static ResponseResult ok(Object data) {
        System.out.println(data);
        return new ResponseResult(200, "success", data);
    }

    public static ResponseResult ok(String msg, Object data) {
        return new ResponseResult(200, msg, data);
    }

    public static ResponseResult ok(String msg) {
        return new ResponseResult(200,msg);
    }

    public static ResponseResult ok() {
        return new ResponseResult(200,"success");
    }


    public static ResponseResult fail() {
        return new ResponseResult(4001,"fail");
    }

    public static ResponseResult fail(String msg) {
        return new ResponseResult(4001,msg);
    }
    public static ResponseResult fail(Integer statusCode,String msg) {
        return new ResponseResult(statusCode,msg);
    }
}
