class MyQueue:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.in_stack = []
        self.out_stack = []

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """
        self.in_stack.append(x)

    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        if not self.out_stack and not self.in_stack:
            return 0
        if not self.out_stack:
            self.out_stack.extend(self.in_stack[::-1])
            self.in_stack = []
        return self.out_stack.pop()


    def peek(self) -> int:
        """
        Get the front element.
        """
        if not self.out_stack and not self.in_stack:
            return 0
        if not self.out_stack:
            # 这里注意 将所有入栈元素弹出到出栈后 清空入栈
            self.out_stack.extend(self.in_stack[::-1])
            self.in_stack = []
        return self.out_stack[-1]

    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        return len(self.in_stack) == 0 and len(self.out_stack) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
