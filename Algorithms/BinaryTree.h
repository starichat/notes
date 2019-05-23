//
// Created by wangxinlong-PC on 2019/5/23.
//

#ifndef UNTITLED_BINARYTREE_H
#define UNTITLED_BINARYTREE_H

// .h 文件用来声明类及其成员函数和成员变量
// 一般放置函数原型
// 使用的#define const 定义的符号常量


// 基于链表实现一颗二叉树
class BinaryTree {
    typedef struct _treenode
    {
        int data;
        struct _treenode *lchild;
        struct _treenode *rchild;
    }Tnode,Tree;
private:
    int node;

public:
    bool binarytree_create(Tree &root); // 创建一颗二叉树
    bool binarytree_destory(Tree &root);  //销毁二叉树
    void binarytree_preorder(Tree &root);  // 先序遍历
    void binarytree_inorder(Tree &root);  // 中序遍历
    void binarytree_postorder(Tree &root);  // 后序遍历
    void binarytree_levelorder(Tree &root);  // 层次遍历
    void binarytree_printfleaf(Tree &root);  // 打印节点
    void binarytree_printlevel(Tree &root);  // 打印树的高度


};


#endif //UNTITLED_BINARYTREE_H
