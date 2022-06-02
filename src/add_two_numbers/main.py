# https://leetcode.com/problems/add-two-numbers/

from typing import Optional, List, Union


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def unpack(l: List[ListNode]) -> List[int]:
    _l = l
    unpacked = []
    while _l is not None:
        value = _l.val
        unpacked.append(str(value))
        _l = _l.next
    return unpacked


def pack(l: List[int]) -> Optional[ListNode]:
    if not len(l):
        return None
    nodes = [ListNode(val, None) for val in l]
    while len(nodes) > 1:
        nodes[-2].next = nodes[-1]
        nodes.pop(-1)
    return nodes[0]


def get_num(l: List[ListNode]) -> Union[int, str]:
    unpacked = unpack(l)
    print(unpacked)
    numString = ''.join(unpacked)[::-1]
    while numString[0] == '0' and len(numString) != 1:
        numString = numString[1:]
    if len(numString):
        num = int(numString)
    return num


class Solution:
    def addTwoNumbers(
        self,
        l1: Optional[ListNode],
        l2: Optional[ListNode],
    ) -> Optional[ListNode]:
        return pack([int(i) for i in str(get_num(l1) + get_num(l2))][::-1])
