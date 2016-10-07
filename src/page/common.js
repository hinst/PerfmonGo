function PerfmonGoApp() {
};

var perfmonGoApp = new PerfmonGoApp();

PerfmonGoApp.prototype.getURLArgument = function (param) {
	var vars = {};
	window.location.href.replace(location.hash, '').replace(
		/[?&]+([^=&]+)=?([^&]*)?/gi, // regexp
		function (m, key, value) { // callback
			vars[key] = value !== undefined ? value : '';
		}
	);
	if (param) {
		return vars[param] ? decodeURIComponent(vars[param]) : null;
	}
	return vars;
};