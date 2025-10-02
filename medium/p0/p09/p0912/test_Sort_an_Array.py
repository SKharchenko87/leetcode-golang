from typing import List
from unittest import TestCase

from medium.p0.p09.p0912.Sort_an_Array import Solution


class TestCase0912:
    nums: [int]
    res: [int]

    def __init__(self, nums, res):
        self.nums = nums
        self.res = res


s = Solution()

e1 = TestCase0912(nums=[5, 2, 3, 1], res=[1, 2, 3, 5])
e2 = TestCase0912([5, 1, 1, 2, 0, 0], [0, 0, 1, 1, 2, 5])


class TestSolution(TestCase):

    def test_sort_array(self):
        x = s.sortArray(e2.nums)
        self.assertEqual(x, e2.res)
        print(x)
    def test_sort_array1(self):
        self.assertEqual(s.sortArray(e1.nums), e1.res)

    def test_partition1(self):
        x = s.partition(e1.nums, 0, 3)
        self.assertEqual(x, 0)
        self.assertEqual(e1.nums, [1, 2, 3, 5])

    def test_partition2(self):
        x = s.partition(e2.nums, 0, 5)
        self.assertEqual(x, 0)
        self.assertEqual(e1.nums, [0, 1, 1, 2, 0, 6])