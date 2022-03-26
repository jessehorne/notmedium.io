$(document).ready(function() {
  if (isUserAuthed()) {
    window.location.href = "/";
  }


  $("#login").submit(function(e) {
    e.preventDefault();

    var formData = getFormData($("#login"));

    (async () => {
      console.log("fetching...");
      const raw = await fetch('http://localhost:8080/api/login', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: formData
      });

      const res = await raw.json();

      if (raw.status === 200) {
        var data = JSON.stringify(res.data);

        window.localStorage.setItem("user", data);

        window.location.href = "http://localhost:8080/";
      } else {
        console.log("Login didn't work!");
      }

    })();
  });
});
