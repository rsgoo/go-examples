<?php

$arr = [3, 1, 2, 5, 4, 7, 2];

$len = count($arr);

for($i=0; $i<$len-1;$i++) {
    for($j=1; $j<$len;$j++) {
        if ($arr[$i] == $arr[$j]) {
            echo $arr[$i];
        }

    }
}