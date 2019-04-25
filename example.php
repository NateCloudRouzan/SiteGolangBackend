<?php

//Was the request (post or get) parameter data supplied?
if(!empty($_REQUEST['data']) {
    $file = 'log.txt';
    $data = $_REQUEST['data']."\n";

    // using the FILE_APPEND flag to append the content to the end of the file
    // and the LOCK_EX flag to prevent anyone else writing to the file at the same time
    file_put_contents($file, $data, FILE_APPEND | LOCK_EX);

}

?>