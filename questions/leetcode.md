[TOC]

# LEETCODE


### 1. 多维数组变单维（可指定解几层）- 2021.05.14

```js
// 示例1
输入:[1, 2, [3, 4], [5, 6]]
输出:[1,2,3,4,5,6]
let arr1 = [1, 2, [3, 4, [5, 6]]];
flat(arr1,1); //输出 [1, 2, 3, 4, [5, 6]]
let arr2 = [1, 2, [3, 4, [5, 6]]];
flat(arr2,2); //输出 [1, 2, 3, 4, 5, 6]

// 示例2
输入:[
    {
        id:1,
        pid:null,
        name:'a',
        children:[
            {
                id:2,
                pid:1,
                name:'aa'
            },
            {
                id:3,
                pid:1,
                name:'ab'
            }
        ]
    },
    {
        id:4,
        pid:null,
        name:'b',
        children:[
            {
                id:5,
                pid:4,
                name:'ba'，
                    children:[{
                        id:7,
                        pid:5,
                        name:'baa'
                    },{
                        id:8,
                        pid:5,
                        name:'bab'
                    },{
                        id:9,
                        pid:5,
                        name:'bac'
                    }]
},
{
    id:6,
        pid:4,
    name:'bb'
}
]
}
]
输出:[
    { id: 1, pid: null, name: 'a' }, { id: 2, pid: 1, name: 'aa' },
    { id: 3, pid: 1, name: 'ab' },{ id: 4, pid: null, name: 'b' },
    { id: 5, pid: 4, name: 'ba' }, { id: 6, pid: 4, name: 'bb' },
    { id: 7, pid: 5, name: 'baa' }, { id: 8, pid: 5, name: 'bab' },
    { id: 9, pid: 5, name: 'bac' }
]
```

### 2. 求两个数组之间的差异,  增加/删除的元素

```js
输入: s1 = [1,2,3,4,5], s2 = [2, 3, 7, 8]
输出: add = [ 7, 8] , del = [1, 4, 5]
以s1为基础
```

### 3.  扁平数组转化树形结构, 要求遍历最大次数为O(2n), 空间复杂度和内存利用率越低越好

```js
输入: arr = [
    { id: 1, pid: 0, name: '超然科技' }, 
    { id: 2, pid: 0, name: '低速科技' },
    { id: 3, pid: 1, name: '科研中心' },
    { id: 4, pid: 1, name: '运营中心' },
    { id: 5, pid: 2, name: '吃喝院' }, 
    { id: 6, pid: 2, name: '研究院' },
    { id: 7, pid: 3, name: 'aa' }, 
    { id: 8, pid: 3, name: 'bb' },
    { id: 9, pid: 4, name: 'cc' },
    { id: 10, pid: 5, name: 'dd' },
    { id: 11, pid: 6, name: 'ee' }
    ],
输出: tree = [
    { 
        id: 1, 
        pid: 0, 
        name: '超然科技',
        children: [
            { 
                id: 3,
                pid: 1, 
                name: '科研中心',
                children: [
                    { 
                        id: 7, 
                        pid: 3, 
                        name: 'aa',
                        children:[]
                    },
                    { 
                        id: 8,
                        pid: 3, 
                        name: 'bb', 
                        children:[]
                    },
                ]
            },
            { 
                id: 4, 
                pid: 1, 
                name: '运营中心',
                children: [
                    { 
                        id: 9, 
                        pid: 4, 
                        name: 'cc', 
                        children:[]
                    },
                ]
            },
        ]
    },
    { 
        id: 2, 
        pid: 0, 
        name: '低速科技',
        children: [
            { 
                id: 5, 
                pid: 2, 
                name: '吃喝院',
                children: [
                    { 
                        id: 10, 
                        pid: 5, 
                        name: 'dd', 
                        children:[]
                    },
                ]
            },
            { 
                id: 6, 
                pid: 2, 
                name: '研究院',
                children: [
                    { 
                        id: 11, 
                        pid: 6, 
                        name: 'ee', 
                        children:[],
                    }
                ]
            },
        ]
    },
]
```

### 4，修改字符串，按指定格式输出

```
输入：str = '2021年6月4日'

输出：str = '2021-06-04'
```

### 5，数组去重并排序

```
输入：arr = [1,3,2,5,6,3,2,7]

输出：arr = [1,2,3,5,6,7]
```


### 6. 只出现一次的数字

leetcode: [136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/)

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
```
输入: [2,2,1]
输出: 1
```

示例2:

```
输入: [4,1,2,1,2]
输出: 4
```
示例3:

```bazaar
输入: [6, 1, 2, 1, 2, 5, 7, 5, 3, 7, 6]
输出: 3
```

### 7. 求一个数是否为2的幂次方数并且获得大于等于这个数的最小2的幂次数

示例1:
```
输入: 10
输出: false, 16
```
示例2:
```
输入: 1024
输出: true, 1024
```
示例2:
```
输入: 30
输出: false, 32
```