$(document).ready(function() {
  if (!isUserAuthed()) {
    window.location.href = "/";
  }

  var articleID = $("#articleID").val();

  console.log("Article ID", articleID)

  var userID = getUser().ID;
  $("#userID").val(userID);

  $("#savedraft").click(function(e) {
    e.preventDefault();

    $("#published").val("false");

    var formData = getFormData($("#form"));

    (async (articleID) => {
      const raw = await fetch('/api/articles/' + articleID, {
        method: 'PUT',
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
      } else if (raw.status == 401) {
        window.location.href = "/login";
      } else {
        $("#error").removeClass("hidden");
      }

    })(articleID);

  });

  // on form submit
  $("#form").submit(function(e) {
    e.preventDefault();

    $("#published").val("true");

    var formData = getFormData($("#form"));

    (async (articleID) => {
      console.log(articleID)
      const raw = await fetch('/api/articles/' + articleID, {
        method: 'PUT',
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
      } else if (raw.status == 401) {
        window.location.href = "/login";
      } else {
        $("#error").removeClass("hidden");
      }

    })(articleID);
  });

});
