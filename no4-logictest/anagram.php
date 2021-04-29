<?php 
$map = [];

$data = ['kita', 'atik', 'tika', 'aku', 'kia', 'makan', 'kua'];

//split and sort string
foreach($data as $str){
    $strSplit = str_split($str);
    sort($strSplit);
  
    $strSplit = implode("",$strSplit);
    $map[$strSplit][] = $str;

}

$object = new stdClass();
foreach ($map as $key => $value)
{
  $anagram = json_encode($object->$key = $value). "<br>";
    
  print_r($anagram);
}

/** Output : */
/**
 * ["kita","atik","tika"]
 * ["aku","kua"]
 * ["kia"]
 * ["makan"]
 * 
 * 
 */



