package com.example.easyread.login;

import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;

import androidx.appcompat.app.AppCompatActivity;

import com.example.easyread.ListActivity;
import com.example.easyread.R;

import org.json.JSONException;
import org.json.JSONObject;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.Callback;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

public class LoginActivity extends AppCompatActivity {

    private EditText edit_account, edit_password;



    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_login);

        Button loginbtn = (Button)findViewById(R.id.login);

        loginbtn.setOnClickListener(new View.OnClickListener() {

            @Override
            public void onClick(View v){
                try {
                    // 校验登录是否成功



//                            ResponseResult responseItem = LoginUtil.doLogin("http://192.168.164.201:8080/login",getJsonFromInput());
//                             System.out.println(responseItem.getMsg());
//                            if (responseItem.getStatus()!=200){
//
//                                return;
//                            }
                            //test("http://192.168.164.201:8080/restricted");
                            Intent intent = new Intent(LoginActivity.this, ListActivity.class);
                            startActivity(intent);




                } catch (Exception e){
                    e.printStackTrace();
                }



            }
        });

    }


    private String getJsonFromInput() {

        edit_account = findViewById(R.id.username);

        edit_password = findViewById(R.id.password);

        JSONObject userInfo = new JSONObject();

        try {
            userInfo.put("username",edit_account.getText().toString());
            userInfo.put("password",edit_password.getText().toString());
        } catch (JSONException e) {
            e.printStackTrace();
        }

        return userInfo.toString();

    }


    // token
    public void test(String url) {
        String returnValue;

        //RequestBody body = RequestBody.create(JSON, json);
        Request request = new Request.Builder()
                .url(url)
                .get()
                .header("Authorization","Bearer eyJhbGciOiJIUzI1NiIsnR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTc1MTE3ODkwLCJuYW1lIjoiSm9uIFNub3cifQ.lLfzwLVFgvlBjxTNUBzSEjmtBX4eGQ3bGGaTL1geayo")
                .build();
        Request req = request;
        OkHttpClient mcl = new OkHttpClient();
        Call call = mcl.newCall(request);

        Callback callback = new Callback() {
            @Override
            public void onFailure(Call call, IOException e) {
                e.printStackTrace();
                Log.e("wxl","adada");

            }

            @Override
            public void onResponse(Call call, Response response)  {
                try {
                    Log.d("Wxl",response.body().string());
                } catch (IOException e) {
                    e.printStackTrace();
                }

            }
        };
        call.enqueue(callback);


    }


    public void readUsersInfo(){
        SharedPreferences sharedPreferences = getSharedPreferences("token",MODE_PRIVATE);
        String token = sharedPreferences.getString("token","");
    }
}
