$(document).ready(function() {
  if (!isUserAuthed()) {
    window.location.href = "/";
  }

  var userID = getUser().ID;
  $("#userID").val(userID);

  $("#savedraft").click(function(e) {
    e.preventDefault();

    $("#published").val("false");

    var formData = getFormData($("#form"));

    (async () => {
      const raw = await fetch('/api/articles', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
          'Authorization': getUser().ApiToken
        },
        body: formData
      });

      const res = await raw.json();

      var articleID = res.data.ID;

      if (raw.status === 200) {
        window.location.href = "/a/" + articleID;
      } else {
        $("#error").removeClass("hidden");
      }

    })();

  });

  // on form submit
  $("#form").submit(function(e) {
    e.preventDefault();

    $("#published").val("true");

    var formData = getFormData($("#form"));

    (async () => {
      const raw = await fetch('/api/articles', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
          'Authorization': getUser().ApiToken
        },
        body: formData
      });

      const res = await raw.json();

      var articleID = res.data.ID;

      if (raw.status === 200) {
        window.location.href = "/a/" + articleID;
      } else {
        $("#error").removeClass("hidden");
      }

    })();
  });

});
