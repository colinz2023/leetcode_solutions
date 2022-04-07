
# [格雷编码](https://leetcode-cn.com/problems/gray-code)

## 题目描述

<strong>n 位格雷码序列</strong> 是一个由 <code>2<sup>n</sup></code> 个整数组成的序列，其中：
<ul>
	<li>每个整数都在范围 <code>[0, 2<sup>n</sup> - 1]</code> 内（含 <code>0</code> 和 <code>2<sup>n</sup> - 1</code>）</li>
	<li>第一个整数是 <code>0</code></li>
	<li>一个整数在序列中出现 <strong>不超过一次</strong></li>
	<li>每对 <strong>相邻</strong> 整数的二进制表示 <strong>恰好一位不同</strong> ，且</li>
	<li><strong>第一个</strong> 和 <strong>最后一个</strong> 整数的二进制表示 <strong>恰好一位不同</strong></li>
</ul>

<p>给你一个整数 <code>n</code> ，返回任一有效的 <strong>n 位格雷码序列</strong> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>n = 2
<strong>输出：</strong>[0,1,3,2]
<strong>解释：</strong>
[0,1,3,2] 的二进制表示是 [00,01,11,10] 。
- 0<strong><em>0</em></strong> 和 0<em><strong>1</strong></em> 有一位不同
- <em><strong>0</strong></em>1 和 <em><strong>1</strong></em>1 有一位不同
- 1<em><strong>1</strong></em> 和 1<em><strong>0</strong></em> 有一位不同
- <em><strong>1</strong></em>0 和 <em><strong>0</strong></em>0 有一位不同
[0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
- <em><strong>0</strong></em>0 和 <em><strong>1</strong></em>0 有一位不同
- 1<em><strong>0</strong></em> 和 1<em><strong>1</strong></em> 有一位不同
- <em><strong>1</strong></em>1 和 <em><strong>0</strong></em>1 有一位不同
- 0<em><strong>1</strong></em> 和 0<em><strong>0</strong></em> 有一位不同
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = 1
<strong>输出：</strong>[0,1]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 16</code></li>
</ul>


## 题解

### python3 [🔗](gray-code.py) 
```python3
class Solution:
    def grayCode(self, n: int) -> List[int]: 
        return [i^(i >> 1) for i in range(1 << n)]
```


## 相关话题

- [位运算](https://leetcode-cn.com/tag/bit-manipulation) 
- [数学](https://leetcode-cn.com/tag/math) 
- [回溯](https://leetcode-cn.com/tag/backtracking) 


## 相似题目

- [1 比特与 2 比特字符](../1-bit-and-2-bit-characters/README.md)  [Easy] 


## Links

- [Prev](../merge-sorted-array/README.md) 
- [Next](../subsets-ii/README.md) 
