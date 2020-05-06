class Solution:
    def isValid(self, s: str) -> bool:
        if not s:
            return True
        # 括号映射
        s_map = {
            ")": "(",
            "}": "{",
            "]": "[",
        }
        # 如果是尾部开头则肯定不是
        if s_map.get(s[0]):
            return False
        # 创建一个栈
        stack = []
        # 遍历字符串， 如果当前字串与栈顶成对， 则弹出
        for sub_s in s:
            if stack and stack[-1] == s_map.get(sub_s, ""):
                stack.pop()
            else:
                stack.append(sub_s)
        # 空栈为有效 否则为无效
        return True if not stack else False



class Solution:
    """利用Python replace 函数"""
    def isValid(self, s: str) -> bool:
        if not s:
            return True
        while "()" in s or "[]" in s or "{}" in s:
            s = s.replace("()", "")
            s = s.replace("[]", "")
            s = s.replace("{}", "")
        return s == ""