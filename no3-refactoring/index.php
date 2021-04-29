<?php

    function firstBracket($firstbracket)
    {
        if($firstbracket){
            $firstbracket = ltrim($firstbracket, '(');
            return strstr($firstbracket, ')', true);
        }
        throw new Exception('Failed');
    }
 
    function isHasBracket($str){
        $res = strstr($str, '(');
        return $res;
    }
   
    function isNotNull($str){
        $res = strlen($str) ? true : false;
        return $res;
    }
  

    function findFirstStringInBracket($str){
       
        if(isNotNull($str) && isHasBracket($str)){
            $newWord = isHasBracket($str); 
            return firstBracket($newWord);
        }else{
            return '';
        }
    }

//Fungsi sebelum refactor
    function oldfindFirstStringInBracket($str){
        if(strlen($str) > 0){ //return 1
            $firstbracket = strstr($str, '('); //return kuat dan sangat berani
            if($firstbracket){
                $firstbracket = ltrim($firstbracket, '(');
                return strstr($firstbracket, ')', true);
            }else{
                return '';
            }
        }else{
        return '';
        }
    }

//Result
$word = 'Saya tidak (pernah) melakukan perbuatan tercela (itu)';
$kata = "Before Refactor :". oldfindFirstStringInBracket($word);
$kata2 = "After Refactor :".findFirstStringInBracket($word);
print_r($kata); //return pernah
print_r($kata2); //return pernah

/**Result */
/** pernah */