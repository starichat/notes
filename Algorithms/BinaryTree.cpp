//
// Created by wangxinlong-PC on 2019/5/23.
//
/**
 * 头文件
 * 实现头文件中的声明
 */
#include <iostream>
#include "BinaryTree.h"

using namespace std;


bool BinaryTree::binarytree_create(Tree &root){

    return true;

}

void BinaryTree::binarytree_preorder(BinaryTree::Tree &root) {
    if(root.data!=0){
        return;
    }
    printf(root.data);
    // 遍历左子树
    binarytree_preorder(root.lchild);

    binarytree_preorder(root.rchild);

}

void BinaryTree::binarytree_inorder(BinaryTree::Tree &root) {

    binarytree_inorder(root.lchild);
    printf(root.data);
    binarytree_inorder(root.rchild);
}

void BinaryTree::binarytree_postorder(BinaryTree::Tree &root) {
    // 遍历左子树

}

void BinaryTree::binarytree_levelorder(BinaryTree::Tree &root) {
    // 遍历跟
    // 用栈先序遍历的节点，当读到下一层的时候就出栈
}
