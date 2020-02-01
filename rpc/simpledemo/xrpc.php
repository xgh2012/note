<?php

$client = new RpcClient();

$client->init();

class RpcClient{
    function init(){
        $url = "127.0.0.1";
        $port = 1234;

        $fk = fsockopen($url,$port);

        fwrite($fk,json_encode(array(
            //'method' => "HelloService.Wait", //map[string]string
            //'params' => [["name"=>"xghWait"]],

            //'method' => "HelloService.Say",  //string
            //'params' => ["xghWait1"],//'params' => ["xghWait1","xghWait2"] 这种写法 只会获取到第一个值

            //'method' => "HelloService.Slice",  //slice []string
            //'params' => [["xghSlice1","xghSlice2"]],

            'method' => "HelloService.SliceInt",  //slice []int
            'params' => [[1,2]],

            'id'     => 0)));
        stream_set_timeout($fk, 0, 3000);
        $res = fgets($fk,1024);
        echo $res."\n";
        fclose($fk);

    }
}