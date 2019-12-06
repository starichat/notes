package com.example.easyread;

import android.os.Bundle;
import android.widget.TextView;

import androidx.appcompat.app.AppCompatActivity;

public class ArticleDetailActivity extends AppCompatActivity {

    TextView title, content ;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_article_detail);
        title = findViewById(R.id.title);
        content = findViewById(R.id.content);
        title.setText("aaaa");
        content.setText("hadksjssssssdhhhcnalkjdlllllllllllllllllllllllllllllllllllllllllllllllllllll");
    }
}
