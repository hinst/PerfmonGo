var receiveData = function(data) {
	console.log("r");
	data.type = "bar";
	var lastMoment = data.UnixNow * 1000;
	data.UnixNow = undefined;
	Plotly.newPlot("graph", [data], 
		{
			bargap: 0,
			margin: {
				l: 30,
				r: 0,
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
	$.get(appURL + "/latestCPU", {seconds: 60}, receiveData, "json");	
}
requestData();
setInterval(requestData, 2000);
