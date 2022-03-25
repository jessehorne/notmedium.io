$(document).ready(function() {
  $("#register").submit(function(e) {
    e.preventDefault();

    var formData = getFormData($("#register"));

    console.log(formData);

    (async () => {
      console.log("fetching...");
      const raw = await fetch('http://localhost:8080/api/register', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: formData
      });

      if (raw.status === 200) {
        window.location.href = "http://localhost:8080/login";
      } else {
        console.log("Register didn't work!");
      }

      const res = await raw.json();
    })();
  });
});
