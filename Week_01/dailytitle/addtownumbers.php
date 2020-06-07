<?php
// 445. 两数相加 II
function addTwoNumbers($l1, $l2) {
    if ($l1 == null) return $l2;
    if ($l2 == null) return $l1;

    $stack1 = [];
    $stack2 = [];
    while($l1) {
        $stack1[] = $l1->val;
        $l1 = $l1->next;
    }
    while($l2) {
        $stack2[] = $l2->val;
        $l2 = $l2->next;
    }

    $carry = 0;
    $head = null;
    while(!empty($stack1) || !empty($stack2) || $carry > 0) {
        $sum = $carry;
        $sum += empty($stack1) ? 0 : array_pop($stack1);
        $sum += empty($stack2) ? 0 : array_pop($stack2);
        $node = new ListNode($sum % 10);
        $node->next = $head;
        $head = $node;
        $carry = (int) ($sum / 10);
    }

    return $head;
}

