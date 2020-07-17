package leetcode

import org.junit.jupiter.api.Test
import xyz.hoper.leetcode.Solution
import kotlin.system.measureTimeMillis

class Solution {
  @Test
  fun lengthOfLongestSubstring() {
    val time1 = measureTimeMillis {
      repeat(100000) {
        lengthOfLongestSubstring("abcabcbb")
      }

    }
    val time2 = measureTimeMillis {
      repeat(100000) {
        Solution.lengthOfLongestSubstringV2("abcabcbb")
      }
    }
    //kotlin还快一点
    println("time:$time1")
    println("time:$time2")
  }

  @Test
  fun reverse() {
    println(reverse(1534236469))
  }

  @Test
  fun findMedianSortedArrays() {
    val nums1 = intArrayOf(3)
    val nums2 = intArrayOf(-2, -1)
    println(findMedianSortedArrays(nums1, nums2))
  }

  @Test
  fun threeSum() {
    val nums = intArrayOf(1, 1, -2)
    println(threeSum(nums))
  }

  @Test
  fun mergeTwoLists() {
    val node1 = ListNode(1).apply { next = ListNode(2).apply { next = ListNode(4) } }
    val node2 = ListNode(1).apply { next = ListNode(3).apply { next = ListNode(4) } }
    println(mergeTwoLists(node1, node2))
  }

  @Test
  fun isValid() {
    println(isValidV2("{[]}"))
  }

  @Test
  fun generateParenthesis() {
    println(generateParenthesis(3))
  }

  @Test
  fun binaryTree() {
    val arr = intArrayOf(5, 3, 12, 36, 728, 333, 128)
    val bt = BinaryTree<Int>()
    for (i in arr.indices) {
      bt.insert(arr[i])
    }
    bt.prevRecursive()
    println("\n")
    print("${bt.midIterator()}\n")
    print("${bt.prevIterator()}\n")
    bt.subRecursive()
    bt.sequence().forEach { println(it) }
  }

  @Test
  fun mergeKLists() {
    val node1 = ListNode(1).apply { next = ListNode(4).apply { next = ListNode(5) } }
    val node2 = ListNode(1).apply { next = ListNode(3).apply { next = ListNode(4) } }
    val node3 = ListNode(2).apply { next = ListNode(6) }
    val list: Array<ListNode?> = arrayOf(node1, node2, node3)
    println(mergeKListsV2(list))
  }

  @Test
  fun fourSum() {
    val arr = intArrayOf(1, 0, -1, 0, -2, 2)
    println(fourSum(arr, 0))
  }

  @Test
  fun removeNthFromEnd() {
    val node = ListNode(1).apply {
      next = ListNode(2).apply {
        next = ListNode(3).apply {
          next = ListNode(4).apply { next = ListNode(5) }
        }
      }
    }
    println(removeNthFromEnd(node, 2))
  }


  @Test
  fun combinationSum() {
    val arr = intArrayOf(1, 2)
    println(combinationSum(arr, 4))
  }

  @Test
  fun strStr() {
    println(strStr("hello", "llo"))
  }

  @Test
  fun serialize() {
    val node = TreeNode(0).apply {
      left = TreeNode(0).apply {
        left = TreeNode(0)
      }
      right = TreeNode(0).apply {
        right = TreeNode(1).apply {
          right = TreeNode(2)
        }
      }
    }
    println(serialize(deserialize("0,0,0,0,null,null,1,null,null,null,2")))
  }

  @Test
  fun maxScoreSightseeingPair() {
    println(maxScoreSightseeingPair(intArrayOf(8, 1, 5, 2, 6)))
  }

  @Test
  fun recoverFromPreorder() {
    println(serialize(recoverFromPreorder("7-6--2--10---1----7-----4----10---4")))
  }

  @Test
  fun isPalindrome() {
    println(isPalindrome("0P"))
  }

  @Test
  fun rotate() {
    val arr = arrayOf(
      intArrayOf(5, 1, 9, 11),
      intArrayOf(2, 4, 8, 10),
      intArrayOf(13, 3, 6, 7),
      intArrayOf(15, 14, 12, 16)
    )
    rotate(arr)
    println(arr)
  }

  @Test
  fun firstMissingPositive() {
    println(firstMissingPositive(intArrayOf(0, -1, 3, 1)))
  }

  @Test
  fun addBinary() {
    println(addBinary("1010", "1011"))
  }

  @Test
  fun threeSumClosest() {
    println(threeSumClosest(intArrayOf(1, 6, 9, 14, 16, 70), 81))
  }

  @Test
  fun myPow() {
    println(Int.MIN_VALUE)
    println(Int.MAX_VALUE)
    println(myPow(2.00000, -2147483648))
  }

  @Test
  fun searchRange() {
    val arr = searchRange(intArrayOf(2, 2, 2), 2)
    println("${arr[0]},${arr[1]}")
  }

  @Test
  fun minSubArrayLen() {
    println(minSubArrayLen(7, intArrayOf(2, 3, 1, 2, 4, 3)))
  }

  @Test
  fun longestValidParentheses() {
    println(longestValidParentheses("(()("))
  }

  @Test
  fun isValidSudoku() {
    for (i in 0 until 9) {
      println("---$i---")
      for (j in 0 until 9) {
        val x = j / 3 + (i / 3) * 3
        val y = j % 3 + (i % 3) * 3
        println("($x,$y)")
      }
    }
  }
  @Test
  fun findKthLargest(){
    println(findKthLargest(intArrayOf(3,2,1,5,6,4),2))
  }
  @Test
  fun findLength(){
    println(findLength(intArrayOf(3,2,1,5),intArrayOf(1,5)))
  }
  @Test
  fun sortedArrayToBST(){
   sortedArrayToBST(intArrayOf(0,1,2,3,4,5))
  }
  @Test
  fun permute(){
    println(permute(intArrayOf(1,2,3,4)))
  }
  @Test
  fun hasPathSum(){
    println(hasPathSum(deserialize("5,4,8,11,null,13,4,7,2,null,null,null,1"),22))
  }
  @Test
  fun divingBoard(){
    printArr(divingBoard(1,2,5))
  }
  @Test
  fun trieTree(){
    println(respace(arrayOf("looked","just","ju","like","her","brother"),"jesslookedjustliketimherbrother"))
  }
  @Test
  fun intersect(){
    printArr(intersect(intArrayOf(4,9,5),intArrayOf(9,4,8,4)))
  }
  @Test
  fun subsets(){
    println(subsets(intArrayOf(1,2,3,4)))
  }
  @Test
  fun lru(){
    val cache = LRUCache(10)
    cache.put(10, 13)
    cache.put(3, 17)
    cache.put(6, 11)
    cache.put(10, 5)
    cache.get(13)
    cache.put(2, 19)
    cache.get(2)
    println(cache.get(3))      // 返回  1
    cache.put(5, 25)    // 该操作会使得关键字 2 作废
    println(cache.get(8))       // 返回 -1 (未找到)
    cache.put(9, 22)
    cache.put(5, 5)
    cache.put(1, 30)
    println(cache.get(11))       // 返回  3
    cache.put(9, 12)
    println(cache.get(7))
    println(cache.get(5))
    println(cache.get(8))
    println(cache.get(9))
    cache.put(4, 30)
    cache.put(9, 3)
    println(cache.get(9))
    println(cache.get(10))
    println(cache.get(10))
    cache.put(6, 14)
    cache.put(3, 1)
    println(cache.get(3))
    cache.put(10, 11)
    println(cache.get(8))
    cache.put(2, 14)
    println(cache.get(1))
    println(cache.get(5))
    println(cache.get(4))
    cache.put(11, 4)
    cache.put(12, 24)
    cache.put(5, 18)
    println(cache.get(13))
    cache.put(7, 23)
    println(cache.get(8))
    println(cache.get(12))
    cache.put(3, 27)
    cache.put(2, 12)
    println(cache.get(5))
    cache.put(2, 9)
    cache.put(13, 4)
    cache.put(8, 18)
    cache.put(1, 7)
    println(cache.get(6))
  }
}
