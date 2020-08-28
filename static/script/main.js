let main = document.querySelector("#demo");

// demo
httpGET("./char?usr=chris&name=dan", function(response) {
	main.innerHTML = response.responseText
	});


function httpGET(url, callback) {
	var xmlHttp = new XMLHttpRequest();
	xmlHttp.onreadystatechange = function() { 
		if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
			callback(xmlHttp);
	}
	xmlHttp.open("GET", url, true); // true for asynchronous 
	xmlHttp.send(null);
}
