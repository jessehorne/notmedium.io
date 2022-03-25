$(document).ready(function() {
  $("#login").submit(function(e) {
    e.preventDefault();

    var formData = getFormData($("#login"));

    console.log(formData);

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
        window.localStorage.setItem("ApiToken", res.data.ApiToken);
        window.localStorage.setItem("ApiTokenExpiresAt", res.data.ApiTokenExpiresAt);
        window.location.href = "http://localhost:8080/";
      } else {
        console.log("Register didn't work!");
      }

    })();
  });
});
