# Terminal-Health-Checkup
Simple terminal health monitoring system. Implemented in Golang

This app lets users add Terminals(URLs) with following configurations: <br>
=>URL        : to hit. <br>
=>TimeOut    : timeout setting for GET request. <br>
=>freqency   : how much time app has to wait to resend GET request after a non 200 Status. <br>
=>Threshold  : Maximum number of tries to mark a terminal Healty or Unhealthy. <br>

<br>
<br>

URL Processing is done by cron jobs. <br>
Crons run by their own go routine and process each terminal. <br>


<br>
Main file is : base/main.go. <br> <br> <br> <br>









Feel free to use this for reference <br>
If you feel something is wrong or can be updated, please crete an issue :) <br>

Thank you <br>




