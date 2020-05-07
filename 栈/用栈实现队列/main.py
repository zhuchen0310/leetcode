class MyQueue:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.stack_in = []
        self.stack_out = []

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """
        return self.stack_in.append(x)

    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        if self.stack_out:
            return self.stack_out.pop()
        if self.stack_in:
            # 将所有入栈元素压入出栈
            self.stack_out.extend(self.stack_in[::-1])
            self.stack_in = []
            return self.stack_out.pop()
        return 0

    def peek(self) -> int:
        """
        Get the front element.
        """
        if self.stack_out:
            return self.stack_out[-1]
        if self.stack_in:
            # 将所有入栈元素压入出栈
            self.stack_out.extend(self.stack_in[::-1])
            self.stack_in = []
            return self.stack_out[-1]
        return 0

    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        return len(self.stack_out) == 0 and len(self.stack_in) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
