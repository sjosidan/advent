<!DOCTYPE html>
<html>

<body>

<link rel="stylesheet" type="text/css" href="index.css" media="screen" />

<?php

function cmp($a, $b) {
    return strcmp($a->name, $b->name);
}

$url = 'https://adventofcode.com/2019/leaderboard/private/view/673024.json';
$request_url = $url;

$curl = curl_init($request_url);

curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
curl_setopt($curl, CURLOPT_COOKIEFILE, "cookies.txt");

$response = curl_exec($curl);
curl_close($curl);


$json = json_decode($response, true);
echo "<table class=\"sortable\">";
echo "<thead>";
echo "<tr>";
echo "<th>Name</th>";
echo "<th>Stars</th>";
echo "</tr>";
echo "</thead>";

    foreach($json['members'] as $key => $value){

        //echo $value['name'].PHP_EOL;;
        //echo $value['stars'].PHP_EOL;;
        echo "<tr>";
        echo "<td>" . $value['name'] . "</td>";
        echo "<td>" . $value['stars'] . "</td>";
        echo "</tr>";

    };
  echo "</table>";

  echo "<table>";
     

    foreach($json['members'] as $key => $value){

        //echo $value['name'].PHP_EOL;;
        //echo $value['stars'].PHP_EOL;;
        echo "<tr>";
        echo "<td>" . $value['name'] . "</td>";
        echo "<td>" . $value['stars'] . "</td>";
        echo "</tr>";

    };
  echo "</table>";

?>

</body>
</html>
