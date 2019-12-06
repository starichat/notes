package com.example.easyread;

import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ListView;

import androidx.appcompat.app.AppCompatActivity;

import java.util.List;

public class ListActivity extends AppCompatActivity {

    private String[] data = {
            "ada"
    };

    private ListView mListView ;
    private AuthorListAdapter mListAdapter;
    private List<AuthorInfo> mAuthorInfoList;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_list);
        initData();
        initView();
    }






    private void initData() {
        mAuthorInfoList = AuthorInfo.createTestData();
    }

    private void initView() {
        mListView = (ListView) findViewById(R.id.listview);
        mListView.addHeaderView(View.inflate(this, R.layout.author_card_layout, null));
        mListView.addFooterView(View.inflate(this, R.layout.author_card_layout, null));
        mListView.setEmptyView(findViewById(R.id.empty_layout));//设置内容为空时显示的视图
        mListAdapter = new AuthorListAdapter(mAuthorInfoList);
        mListView.setAdapter(mListAdapter);
        mListView.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view, int position, long id) {
                Log.i("wxl", "position : " + position);
                Log.i("wxl", "id : " + id);
                Intent intent = new Intent(ListActivity.this, ArticleDetailActivity.class);
                startActivity(intent);
            }
        });
    }
}
