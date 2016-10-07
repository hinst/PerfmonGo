/// <reference path="common.js"/>
var coresMode = false;
if (perfmonGoApp.getURLArgument("cores") == "y") {
	coresMode = true;
};

var receiveData = function(data) {
	console.log("r");
	for (var i = 0; i < data.Series.length; i++)
		data.Series[i].type = "bar";
	var lastMoment = data.UnixNow * 1000;
	data.UnixNow = undefined;
	Plotly.newPlot("graph", data.Series, 
		{
			bargap: 0,
			margin: {
				l: 30,
				r: 30,
				t: 8,
				b: 40
			},
			xaxis: {
				type: "date",
				//range: [lastMoment - 60 * 1000, lastMoment]
				range: [lastMoment - 60 * 1000, lastMoment]
			},
			yaxis: {
				range: [0, 100]
			}
		}
	);
}
var requestData = function() {
	var url = "";
	var args = {seconds: 60};
	if (coresMode) {
		args.cores = 1;
	}
	$.get(appURL + "/latestCPU", args, receiveData, "json");	
}
requestData();
setInterval(requestData, 2000);

