<?php
header('Content-Type: text/html; charset=UTF-8');
$connection = new mysqli("localhost","root","usbw","test",3307);

$command = "DROP TABLE cnpj";
$connection->query($command);

$command = "CREATE TABLE IF NOT EXISTS cnpj(
    `id` int(11) not null AUTO_INCREMENT,
    `cnpj` varchar(20) not null,
    primary key (`id`)
) engine myISAM";

$connection->query($command);

$dir  = "../files-to-read/";
$files = dir($dir);
$all_lines = 0;

// Register the start time
$start = microtime(true);

while($file = $files->read()){
    if(substr($file,0,1)!="."){
        $data = fopen($dir.$file, "r");
        $line = fgets($data);
            while ($line) {
            $all_lines++;
            $line = fgets($data);
            $lineClean = str_replace("\"","",$line);
            $parts = explode(";",$lineClean);
            $command  = "INSERT INTO cnpj (cnpj) VALUES ('".$parts[0].$parts[2].$parts[5]."') ";
            $connection->query($command);
            }
            fclose( $data );
        
    }
}
$files -> close();

// Register the end time
$end = microtime(true);
echo "Total time for processing files: " . number_format(($end - $start), 2) . " seconds.\n";

?>
