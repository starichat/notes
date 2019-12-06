package com.example.easyread.http;

public class ResponseResult {


    private int status;

    private String msg;

    private Object data;




    public ResponseResult() {
    }

    public ResponseResult(int status, String msg, Object data) {
        this.status = status;
        this.msg = msg;
        this.data = data;
    }

    public ResponseResult(int status, String msg) {
        this.status = status;
        this.msg = msg;
    }


    public int getStatus() {
        return status;
    }


    public void setStatus(int status) {
        this.status = status;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }
}
