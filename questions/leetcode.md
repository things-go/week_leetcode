# LEETCODE

2021.05.14

1,多维数组变单维（可指定解几层）。

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

