function CreatePost() {
	title = document.getElementById("title_tb").value;
	desc = document.getElementById("desc_tb").value;
	templ = document.getElementById("template_sel").value;
	post = document.getElementById("post_ta").value;

	json = { "Title": title, "Description": desc, "Template": templ, "Post": post };
	sendJSON(json);
}

function sendJSON(content) {
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "/api/v1/blog", true);
	xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
	xhr.send(JSON.stringify(content));
}