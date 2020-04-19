<?php
// 542题 矩阵
function updateMatrix($matrix) {
    $row = count($matrix);
    $col = count($matrix[0]);

    $range = $row * $col;
    $dis = [];
    for($i = 0; $i < $row; $i++) {
        for($j = 0; $j < $col; $j++) {
            if ($matrix[$i][$j] == 0){
                $dis[$i][$j] = 0;
                continue;
            } 
            $up = ($i > 0) ? $dis[$i - 1][$j] : $range;
            $left = ($j > 0) ? $dis[$i][$j - 1] : $range; 
            $dis[$i][$j] = 1 + min($up, $left);
        }
    } 

    for($i = $row - 1; $i >= 0; $i--) {
        for($j = $col - 1; $j >= 0; $j--) {
            if ($matrix[$i][$j] == 0){
                $dis[$i][$j] = 0;
                continue;
            } 
            $down = ($i < $row - 1) ? $dis[$i + 1][$j] : $range;
            $right = ($j < $col - 1) ? $dis[$i][$j + 1] : $range; 
            $dis[$i][$j] = min(min($down, $right) + 1, $dis[$i][$j]);
        }
    }

    return $dis;
}