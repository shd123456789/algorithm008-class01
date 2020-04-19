 <?php
 // 189. 旋转数组
function rotate(&$nums, $k) {
    $len = count($nums);
    $k = $k % $len;
    $count = 0;

    for($start = 0; $count < $len; $start++) {
        $current = $start;
        $prev = $nums[$start];

        do {
            $next = ($current+$k) % $len;
            $tmp = $nums[$next];
            $nums[$next] = $prev;
            $prev = $tmp;
            $current = $next;
            $count ++;
        } while($current != $start);
    }
}


function rotate(&$nums, $k) {
    $count = count($nums);
    $k = $k % $count;
    $this->reverse($nums, 0, $count - 1);
    $this->reverse($nums, 0, $k - 1);
    $this->reverse($nums, $k, $count - 1);
}   

function reverse(&$nums, $start, $end)
{
    while ($start < $end) {
        $tmp = $nums[$start];
        $nums[$start] = $nums[$end];
        $nums[$end] = $tmp;
        $start++;
        $end--;
    }
}

