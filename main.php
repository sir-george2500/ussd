<?php

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $sessionId   = $_POST["sessionId"];
    $serviceCode = $_POST["serviceCode"];
    $phoneNumber = $_POST["msisdndigits"];
    $text        = $_POST["ussdstring"];
} else if ($_SERVER['REQUEST_METHOD'] === 'GET') {
    $sessionId   = $_GET["dialogueID"];
    $serviceCode = $_GET["serviceCode"];
    $phoneNumber = $_GET["msisdndigits"];
    $text        = $_GET["ussdstring"];
}


// Specify the file path to save JSON data
$filePath = "sess/$sessionId.json";
$sessionData = file_get_contents($filePath);

if($sessionData != ""){
	$response = json_decode($sessionData, true);
	$level = $response['level'] + 1;
} else {
	$level = 1;
}

$response['level'] = $level; 
$json_en_data = json_encode($response);
file_put_contents($filePath, $json_en_data);

//print_r($response);

if ($text == "" or $text == "255" && $level == 1) {
    $response  = "Welcome to Winliberia Raffle Draw\n";
    $response .= "1. Winliberia Jackpot\n";
    $response .= "2. View Results\n";
    $response .= "3. Jackpot Policy\n";

} elseif ($text == "1" && $level == 2) {
    $response = "Select Currency\n";
    $response .= "1. LRD\n";
    $response .= "2. USD\n";

} elseif ($text == "2" && $level == 2) {
	$response = "Your phone number is ".$phoneNumber;

} elseif (($text == "1" || $text == "2") && $level == 3) { 
    $accountNumber  = "ACC1001";
    $response = "Processing payment in LRD You will receive a message shortly.";
} else {
	$response = "Invalid Input!!";
}

//header("FreeFlow: 2.50");
header("FreeFlow: FC");
echo $response;
?>

