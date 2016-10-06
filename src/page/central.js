var receiveData = function(data) {
	console.log("r");
	data.type = "scatter";
	console.log(data);
	/*
  data = [
  {
    x: ['2013-10-04 22:23:00', '2013-11-04 22:23:00', '2013-12-04 22:23:00'],
    y: [1, 3, 6],
    type: 'scatter'
  }
];
*/
	console.log(data);
	Plotly.newPlot("graph", [data]);
	/*
	Plotly.plot( "graph", [{
	x: [1, 2, 3, 4, 5],
	y: [1, 2, 4, 8, 16] }], {
	margin: { t: 0 } });
	*/
}
var requestData = function() {
	$.get(appURL + "/latestCPU", {seconds: 10}, receiveData, "json");	
}
requestData();
//receiveData();
