 <?php
 // 合并区间力扣第56道题 
 function merge($intervals) {
    $mergeArr = [];
    $count = count($intervals);
    for($i = 0; $i < $count; $i++){
        $sussessFlag = false;
        for($j = $i + 1;$j < $count; $j++) {
            if ($intervals[$i][1] >= $intervals[$j][0] 
                && ($intervals[$i][0] <= $intervals[$j][0] 
                || $intervals[$i][0] <= $intervals[$j][1])
            ){   
                $intervals[$j][0] = min($intervals[$i][0], $intervals[$j][0]);
                $intervals[$j][1] = max($intervals[$j][1], $intervals[$i][1]);
                $sussessFlag = true; 
                break;
            }
        }

        if (!$sussessFlag) {
            $mergeArr[] = $intervals[$i];
        }
    }
    return $mergeArr;
}

