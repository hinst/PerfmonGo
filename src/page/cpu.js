/// <reference path="common.js"/>
var timeInterval = 60; // seconds
var coresMode = false;
if (perfmonGoApp.getURLArgument("cores") == "y") {
	coresMode = true;
};
var receiveData = function(data) {
	for (var i = 0; i < data.Series.length; i++)
		data.Series[i].type = "bar";
	var lastMoment = data.UnixNow * 1000;
	data.UnixNow = undefined;
	var plotSettings = {
		bargap: 0,
		barmode: "stack",
		showlegend: false,
		hovermode: false,
		margin: {
			l: 30,
			r: 30,
			t: 8,
			b: 40
		},
		xaxis: {
			type: "date",
			range: [lastMoment - timeInterval * 1000, lastMoment]
		},
		yaxis: {
			range: [0, 100]
		}
	};
	Plotly.newPlot("graph", data.Series, plotSettings);
};
var requestData = function() {
	var url = "";
	var args = {seconds: timeInterval};
	if (coresMode) {
		args.cores = 1;
	}
	$.get(appURL + "/latestCPU", args, receiveData, "json");	
}
setInterval(requestData, 2000);
var toggleTimeInterval = function(newTimeInterval, button) {
	$("#intervalTogglePanel").find("[data-group='button']").removeClass("bold");
	button.addClass("bold");
	timeInterval = newTimeInterval;
	var buttonId = button.attr("id")
	localStorage.setItem("chosenTimeIntervalButtonId", buttonId);
}
$("#button1m").on("click", function() { toggleTimeInterval(60, $("#button1m")); } );
$("#button5m").on("click", function() { toggleTimeInterval(60 * 5, $("#button5m")); } );
var chosenTimeIntervalButtonId = localStorage.getItem("chosenTimeIntervalButtonId");
if (chosenTimeIntervalButtonId == null)
	chosenTimeIntervalButtonId = "button1m";
var chosenTimeIntervalButton = document.getElementById(chosenTimeIntervalButtonId);
if (chosenTimeIntervalButton != null)
	chosenTimeIntervalButton.click();
var coresButton = $("#coresButton");
var updateCoresButtonStatus = function() {
	if (coresMode)
		coresButton.addClass("bold");
	else
		coresButton.removeClass("bold");
};
var toggleCoresMode = function() {
	coresMode = !coresMode;
	updateCoresButtonStatus();
};
coresButton.on("click", toggleCoresMode);
requestData();
