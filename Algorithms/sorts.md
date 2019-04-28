# 排序算法
计算机领域中经常要用到排序，这应该是最基础且最重要的一个算法了。在平时项目甚至大型项目中，排序算法也是经常被用到。下面系统的阐述一下排序算法。

排序算法有很多种，这里按照时间复杂度将排序算法系统地分为三类。如下表所示：
|算法|时间复杂度|是否基于比较排序|
|:--:|:--:|:--:|
|冒泡、插入、选择|O(n^2)|是|
|快排、归并|O(nlogn)|是|
|桶、计数、基数|O(n)|否|

在学习各种排序算法之前，先了解一下如何去衡量一个排序算法，这样有助于我们再合适的场景选取合适的算法。了解用法可以加深我们的学习兴趣和学习效率。

算法无非涉及到效率、资源占用，但是排序算法还有一点就是算法的稳定性。有人会想都是排序，为什么需要稳定性。比如我们想给一串用户订单排序，我们想根据金额排序同时保证订单时间是从早到晚的。这里可以先采用排序算法按照订单时间排序，再将排序结果按照金额大小排序，因为排序算法是稳定的话，当有相同的两组数据不会发生交换（因为是基于比较排序），这样排下来金额相同的数据就还是按照之前订单顺序排列的。

进入正题，如何从上面三个方面去衡量排序算法
## 排序算法的执行效率
对于执行效率一般有下面三个方面：
> 1.最好、最坏、平均情况的时间复杂度

> 2.时间复杂度的系数

有的算法在数据规模很大的情况下会有一个增长趋势，因而一般我们都容易忽略掉系数，但是实际开发过程中有可能碰到一些可量化甚至是数据量很少的数据，这个时候就有必要考虑一下系数了。
> 3.比较次数和交换次数

## 排序算法的内存消耗
每一个算法都有一个内存消耗的衡量标准，排序算法当然也有。内存消耗一般从空间复杂度进行分析。

## 排序算法的稳定性
如果待排序的序列中存在值相等的元素，经过排序之后，相等元素之间原有的先后顺序不变。

## 冒泡排序
望文生义，冒泡排序就是将每一次都将最大（或最小）的元素放到最后面（或最前面）去，每次找到一个当前序列的最大（或最小）元素。
```
    public void bubbleSort(int[] a,int n){
        if(n <= 1) return ;
        
        for(int i = 0;i < n; ++i){
            // 提前退出冒泡循环的标志位
            boolean flag = false;
            for(int j = 0; j<n-i-1;++j){
                if(a[j]<a[i]){
                    int temp = a[i];
                    a[i]=a[j];
                    a[j]=temp;
                    flag = true; //有数据交换
                }
               
            }
             if(!flag) break;  //没有数据交换，证明后面已经拍好序了，不用再继续遍历了。
        }
    }
```

根据冒泡排序的原理，我们也可以分析出，冒泡排序的时间复杂度最坏情况是O(n^2),最好的情况是O(n),空间复杂度是O（1），冒泡排序也是稳定的排序算法，执行的前后交换元素这种措施。
## 插入排序
对一个待排序数组，从第0个元素开始，后面的都是未排序数组，将未排序区间的第一个数通过先比较找到应该插入的位置，在执行数据交换执行插入操作，然后已排序区间变大了，不断地这样执行下去，已达到整个数组全部排序完成。
```
    public void insertionSort(int [] a,int n){
        if(n <= 1) return ;

        for(int i = 1; i<n ;++i){
            int value = a[i];
            int j = i-1;
            // 查找要排序的位置
            for(;j>=0;--j){
                if(a[j]>value) {
                    a[j+1] = a[j];  //数据移动
                } else {
                    break;
                }
            }
            a[j+1] = value; //执行插入
        }
    }
```
## 选择排序

```
    public void selectSort(int[] a,int n){
        if(n <= 1) return ;
        
    }
```

## 快速排序
public class quickSort {

    // 快速排序
    public static void quickSort(int[] a, int start,int end){
        int p = Partition(a,start,end);
        quickSort(a,start,p-1);
        quickSort(a,p+1,end);

    }



    // 分区
    public static int Partition(int[] a,int start,int end){
        int pivot = a[end];  //为了方便，将最后一个元素作为开始元素
        int i = start;
        for(int j=start;j<end;++j){
            if(a[j]<pivot){
                if(i==j){
                    ++i;
                } else {
                    // 交换i,j两个位置
                    int tmp = a[i];
                    a[i++] = a[j];
                    a[j] = tmp;
                }

            }
        }
        int tmp = a[i];
        a[i] = pivot;
        a[end] = tmp;
    return i;
    }

}
## 归并排序
public class MergeSort {

    /**
     * 归并排序思想采用分治的思想:先分后治
     * 将待排序序列折半分解，直到分解为最小单元为止，这是分
     * 对最小单元排序不断向上递推，只是治。
     * 根据这种思路我们可以发现可以用递归来实现。分治是一种思想，递归是一种编程方法。
     */

    // 递归一定要先写出递推公式
    public void merge(int[] a,int start, int end){
        // 递归终止条件
        if(start >= end) return;

        // 找到中间位置
        int middle = start + (end - start)/2;
        // 分
        merge(a,start,middle);
        merge(a,middle+1,end);

        // 合,将有序的两个数据排序
        mergeSort(a,start,middle,end);
    }
    // 因为两个序列都是有序的，所以进行归并时，可以参考有序序列的合并
    public void mergeSort(int[] a,int start,int middle,int end){
       // 引入两个索引
        int i = start;
        int j = middle+1;
        // 申请一个和a[start~end] 一样大小的空间
        int[] temp = new int[end-start+1];
        int k = 0;
        while(i<=middle&&j<=end){
            if(a[i]>a[j]){
                // 说明a[j]是目前的最小的数
                temp[k] = a[j];
                k++;
                j++;
            }
            else {
                // 说明a[i]是最小的数了
                temp[k] = a[i];
                k++;
                i++;
            }
        }
        // 可能某一个序列的数全部放到temp里了，但是另外一个序列还有数据，应该全部放到temp后面去
        // 判断哪个序列还有剩余的数据
        int begin = i;
        int tend = middle;
        if(j<=end){
            begin = j;
            tend = end;
        }
        // 将剩余数组放到temp中去
        while(start <= end){
            temp[k++] = a[begin++];
        }
        // 拷贝会原数组
        for(i = 0;i<end-middle;++i){
            a[middle+i] = temp[i];
        }

    }


}

