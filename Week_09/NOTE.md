Rabin-Karp 算法的思想

1，假设子串的长度为M(pat),目标字符串的长度为N(txt)

2. 计算子串的hash值hash_pat

3. 计算目标字符串txt中每个长度为M的子串的hash值（共需要计算 N-M+1次）

4. 比较hash值：如果hash值不同，字符串必然不匹配；如果hash值相同，还需要使用朴素算法再次判断
