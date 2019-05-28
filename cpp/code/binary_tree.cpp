#include<iostream>
using namespace std;

// 二叉树结点类型
struct Node
{
    int data;
    Node *left,*right;
};

class Btree {
    static int node_number;
    static int leaf_number;
public:
    Node *root;
    Btree(){
        root = CreateNode(root);
    }
    Node *CreateNode(Node *bn); // 按照二叉搜索树插入节点

    void PrintNode(Node *node);  //打印节点
    int PrintLevelanddepth(Node *node);//输出层数和高度、深度
    void PreOrder(Node *root);//先序
    void InOrder(Node *root);//中序
    void LastOrder(Node *root);//后序
    void LevelOrder(Node *root);//层次


};
int Btree::node_number = 0;
int Btree::leaf_number = 0;

/**
 * 按照先序遍历的方式构建：根左右
 * @param data
 * @return
 */
Node *Btree::CreateNode(Node *bn)
{
    int data;
    // 开辟一个新结点
    bn = new Node;
    cout<<"please input the next node data:"<<endl;
    cin>>data;
    if(data==-1)
    {// 获取要插入的数据，如果data为-1,则该节点为空

        return NULL;

    } else {
        // 先序遍历插入结点
        cout<<"aaa"<<endl;
        bn->data=data;
        bn->left=CreateNode(bn->left);
        bn->right=CreateNode(bn->right);
    }
    // 构建完成，返回头结点
    return bn;
}



// 先序遍历
void Btree::PreOrder(Node *root)
{
    if(root==NULL){
     //   cout<<"the node is NULL"<<endl;
        return;
    }
    cout<<root->data<<endl;
    PreOrder(root->left);
    PreOrder(root->right);

}
// 中序遍历
void Btree::InOrder(Node *root)
{
    if(root==NULL){
        return ;
    }
    InOrder(root->left);
    cout<<root->data<<endl;
    InOrder(root->right);
}
//后续遍历
void Btree::LastOrder(Node *root)
{
    if(root==NULL){
        return;
    }
    LastOrder(root->left);
    LastOrder(root->right);
    cout<<root->data<<endl;
}
// 层次遍历需要借助一个队列来实现，根节点入队，在跟节点入队的同时，该节点的左右孩子入队，对每个节点都是这样操作（节点出对的时候，左右孩子入队）
void Btree::LevelOrder(Node *root) {

}
int Btree::PrintLevelanddepth(Node *node) {
    int l_level = 0;
    int r_level = 0;
    if(node==NULL){
        return 0;
    }
    return (l_level>r_level)?(l_level+1):(r_level+1);// 简单理解就是去掉根节点，递归子树层次（因为去掉了根，所以要加一）
}

int main()
{

    Btree b;
    cout<<"pre"<<endl;
    b.PreOrder(b.root);
    cout<<"IN"<<endl;
    b.InOrder(b.root);
    cout<<"last"<<endl;
    b.LastOrder(b.root);
    cout<<b.PrintLevelanddepth(b.root);
    return 0;


}